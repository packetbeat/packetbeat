// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package memqueue

import (
	"context"
	"errors"
	"io"
	"time"

	"github.com/elastic/beats/v7/libbeat/common/fifo"
	"github.com/elastic/beats/v7/libbeat/publisher/queue"
	"github.com/elastic/elastic-agent-libs/logp"
)

// The string used to specify this queue in beats configurations.
const QueueType = "mem"

const (
	minInputQueueSize      = 20
	maxInputQueueSizeRatio = 0.1
)

// broker is the main implementation type for the memory queue. An active queue
// consists of two goroutines: runLoop, which handles all public API requests
// and owns the buffer state, and ackLoop, which listens for acknowledgments of
// consumed events and runs any appropriate completion handlers.
type broker struct {
	settings Settings
	logger   *logp.Logger

	ctx       context.Context
	ctxCancel context.CancelFunc

	// The factory used to create an event encoder when creating a producer
	encoderFactory queue.EncoderFactory

	///////////////////////////
	// api channels

	// Producers send requests to pushChan to add events to the queue.
	pushChan chan pushRequest

	// Consumers send requests to getChan to read events from the queue.
	getChan chan getRequest

	// Close triggers a queue close by sending to closeChan.
	closeChan chan struct{}

	///////////////////////////
	// internal channels

	// Batches sent to consumers are also collected and forwarded to ackLoop
	// through this channel so ackLoop can monitor them for acknowledgments.
	consumedChan chan batchList

	// When batches are acknowledged, ackLoop saves any metadata needed
	// for producer callbacks and such, then notifies runLoop that it's
	// safe to free these events and advance the queue by sending the
	// acknowledged event count to this channel.
	deleteChan chan int

	// closingChan is closed when the queue has processed a close request.
	// It's used to prevent producers from blocking on a closing queue.
	closingChan chan struct{}

	///////////////////////////////
	// internal goroutine state

	// The goroutine that manages the queue's core run state and owns its
	// backing buffer.
	runLoop *runLoop

	// The goroutine that manages ack notifications and callbacks
	ackLoop *ackLoop
}

type Settings struct {
	// The number of events and bytes the queue can hold. <= zero means no limit.
	// At least one must be greater than zero.
	Events int
	Bytes  int

	// The most events that will ever be returned from one Get request.
	MaxGetRequest int

	// If positive, the amount of time the queue will wait to fill up
	// a batch if a Get request asks for more events than we have.
	FlushTimeout time.Duration
}

type queueEntry struct {
	event     queue.Entry
	eventSize int

	producer   *ackProducer
	producerID producerID // The order of this entry within its producer
}

type batch struct {
	// The queue buffer (at the time that this batch was generated --
	// only the indices corresponding to this batch's events are guaranteed
	// to be valid).
	queueBuf circularBuffer

	// Position of the batch's events within the queue. This is an absolute
	// index over the lifetime of the queue, to get the position within the
	// queue's current circular buffer, use (start % len(queue.buf)).
	start entryIndex

	// Number of sequential events in this batch.
	count int

	// batch.Done() sends to doneChan, where ackLoop reads it and handles
	// acknowledgment / cleanup.
	doneChan chan batchDoneMsg
}

type batchList = fifo.FIFO[batch]

// FactoryForSettings is a simple wrapper around NewQueue so a concrete
// Settings object can be wrapped in a queue-agnostic interface for
// later use by the pipeline.
func FactoryForSettings(settings Settings) queue.QueueFactory {
	return func(
		logger *logp.Logger,
		observer queue.Observer,
		inputQueueSize int,
		encoderFactory queue.EncoderFactory,
	) (queue.Queue, error) {
		return NewQueue(logger, observer, settings, inputQueueSize, encoderFactory)
	}
}

// NewQueue creates a new broker based in-memory queue holding up to sz number of events.
// If waitOnClose is set to true, the broker will block on Close, until all internal
// workers handling incoming messages and ACKs have been shut down.
func NewQueue(
	logger *logp.Logger,
	observer queue.Observer,
	settings Settings,
	inputQueueSize int,
	encoderFactory queue.EncoderFactory,
) (*broker, error) {
	b, err := newQueue(logger, observer, settings, inputQueueSize, encoderFactory)

	if err == nil {
		// Start the queue workers
		go b.runLoop.run()
		go b.ackLoop.run()
	}

	return b, err
}

// newQueue does most of the work of creating a queue from the given
// parameters, but doesn't start the runLoop or ackLoop workers. This
// lets us perform more granular / deterministic tests by controlling
// when the workers are active.
func newQueue(
	logger *logp.Logger,
	observer queue.Observer,
	settings Settings,
	inputQueueSize int,
	encoderFactory queue.EncoderFactory,
) (*broker, error) {
	if observer == nil {
		observer = queue.NewQueueObserver(nil)
	}
	chanSize := AdjustInputQueueSize(inputQueueSize, settings.Events)

	// Backwards compatibility: an old way to select synchronous queue
	// behavior was to set "flush.min_events" to 0 or 1, in which case the
	// timeout was disabled and the max get request was half the queue.
	// (Otherwise, it would make sense to leave FlushTimeout unchanged here.)
	if settings.MaxGetRequest <= 1 {
		settings.FlushTimeout = 0
		settings.MaxGetRequest = (settings.Events + 1) / 2
	}

	if settings.Bytes > 0 && encoderFactory == nil {
		return nil, errors.New("queue.mem.bytes is set but the output doesn't support byte-based event buffers")
	}

	// Can't request more than the full queue
	if settings.Events > 0 && settings.MaxGetRequest > settings.Events {
		settings.MaxGetRequest = settings.Events
	}

	if logger == nil {
		logger = logp.NewLogger("memqueue")
	}

	b := &broker{
		settings: settings,
		logger:   logger,

		encoderFactory: encoderFactory,

		// broker API channels
		pushChan:  make(chan pushRequest, chanSize),
		getChan:   make(chan getRequest),
		closeChan: make(chan struct{}),

		// internal runLoop and ackLoop channels
		consumedChan: make(chan batchList),
		deleteChan:   make(chan int),
		closingChan:  make(chan struct{}),
	}
	b.ctx, b.ctxCancel = context.WithCancel(context.Background())

	b.runLoop = newRunLoop(b, observer)
	b.ackLoop = newACKLoop(b)

	observer.MaxEvents(settings.Events)

	return b, nil
}

func (b *broker) Close() error {
	b.closeChan <- struct{}{}
	return nil
}

func (b *broker) Done() <-chan struct{} {
	return b.ctx.Done()
}

func (b *broker) QueueType() string {
	return QueueType
}

func (b *broker) BufferConfig() queue.BufferConfig {
	return queue.BufferConfig{
		MaxEvents: b.settings.Events,
	}
}

func (b *broker) Producer(cfg queue.ProducerConfig) queue.Producer {
	// If we were given an encoder factory to allow producers to encode
	// events for output before they entered the queue, then create an
	// encoder for the new producer.
	var encoder queue.Encoder
	if b.encoderFactory != nil {
		encoder = b.encoderFactory()
	}
	return newProducer(b, cfg.ACK, encoder)
}

func (b *broker) Get(count int, bytes int) (queue.Batch, error) {
	responseChan := make(chan batch, 1)
	select {
	case <-b.ctx.Done():
		return nil, io.EOF
	case b.getChan <- getRequest{
		entryCount: count, byteCount: bytes, responseChan: responseChan}:
	}

	// if request has been sent, we have to wait for a response
	resp := <-responseChan
	return resp, nil
}

func (b *broker) useByteLimits() bool {
	return b.settings.Bytes > 0
}

func newBatch(queueBuf circularBuffer, start entryIndex, count int) batch {
	return batch{
		doneChan: make(chan batchDoneMsg, 1),
		queueBuf: queueBuf,
		start:    start,
		count:    count,
	}
}

// AdjustInputQueueSize decides the size for the input queue.
func AdjustInputQueueSize(requested, mainQueueSize int) (actual int) {
	actual = requested
	if max := int(float64(mainQueueSize) * maxInputQueueSizeRatio); mainQueueSize > 0 && actual > max {
		actual = max
	}
	if actual < minInputQueueSize {
		actual = minInputQueueSize
	}
	return actual
}

func (b batch) Count() int {
	return b.count
}

// Return a pointer to the queueEntry for the i-th element of this batch
func (b batch) entry(i int) *queueEntry {
	entryIndex := b.start.plus(i)
	return b.queueBuf.entry(entryIndex)
}

// Return the event referenced by the i-th element of this batch
func (b batch) Entry(i int) queue.Entry {
	return b.entry(i).event
}

func (b batch) Done() {
	b.doneChan <- batchDoneMsg{}
}
