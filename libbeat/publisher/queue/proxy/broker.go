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

package proxyqueue

import (
	"io"
	"sync"

	"github.com/elastic/beats/v7/libbeat/publisher/queue"
	c "github.com/elastic/elastic-agent-libs/config"
	"github.com/elastic/elastic-agent-libs/logp"
)

type broker struct {
	done chan struct{}

	logger *logp.Logger

	// The maximum number of events in any pending batch
	batchSize int

	///////////////////////////
	// api channels

	// Producers send requests to pushChan to add events to the queue.
	pushChan chan pushRequest

	// Consumers send requests to getChan to read events from the queue.
	getChan chan getRequest

	// A listener that should be notified when ACKs are processed.
	// Right now this listener always points at the Pipeline associated with
	// this queue, and Pipeline.OnACK forwards the notification to
	// Pipeline.observer.queueACKed(), which updates the beats registry
	// if needed.
	ackListener queue.ACKListener

	// wait group for worker shutdown
	wg sync.WaitGroup
}

type Settings struct {
	ACKListener queue.ACKListener
	BatchSize   int
}

type queueEntry struct {
	event interface{}
	id    queue.EntryID

	producer *producer
}

func init() {
	// TODO: This queue is not currently registered so it doesn't become
	// accessible via configurations. The intent is for it to be enabled
	// automatically when the shipper output is in use, and not otherwise,
	// which will require some changes to beats initialization and/or
	// configuration handling.
	/*queue.RegisterQueueType(
	"proxy",
	create,
	feature.MakeDetails(
		"Proxy queue",
		"Pass through batched events to the Elastic Agent Shipper.",
		feature.Experimental))*/
}

func create(
	ackListener queue.ACKListener, logger *logp.Logger, cfg *c.C, inQueueSize int,
) (queue.Queue, error) {
	config := defaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, err
	}

	if logger == nil {
		logger = logp.L()
	}

	return NewQueue(logger, Settings{
		ACKListener: ackListener,
		BatchSize:   50,
	}), nil
}

// NewQueue creates a new broker based in-memory queue holding up to sz number of events.
// If waitOnClose is set to true, the broker will block on Close, until all internal
// workers handling incoming messages and ACKs have been shut down.
func NewQueue(
	logger *logp.Logger,
	settings Settings,
) *broker {
	if logger == nil {
		logger = logp.NewLogger("proxyqueue")
	}

	b := &broker{
		done:   make(chan struct{}),
		logger: logger,

		// broker API channels
		pushChan: make(chan pushRequest),
		getChan:  make(chan getRequest),

		ackListener: settings.ACKListener,
	}

	b.wg.Add(1)
	go func() {
		defer b.wg.Done()
		b.run()
	}()

	return b
}

func (b *broker) Close() error {
	close(b.done)
	return nil
}

func (b *broker) BufferConfig() queue.BufferConfig {
	return queue.BufferConfig{}
}

func (b *broker) Producer(cfg queue.ProducerConfig) queue.Producer {
	return newProducer(b, cfg.ACK)
}

func (b *broker) Get(count int) (queue.Batch, error) {
	responseChan := make(chan getResponse, 1)
	select {
	case <-b.done:
		return nil, io.EOF
	case b.getChan <- getRequest{
		entryCount: count, responseChan: responseChan}:
	}

	// if request has been sent, we have to wait for a response
	resp := <-responseChan
	return &batch{
		queue:    b,
		entries:  resp.entries,
		doneChan: resp.ackChan,
	}, nil
}

// Metrics returns an empty response because the proxy queue
// doesn't accumulate batches; for the real metadata, use either the
// Beats pipeline metrics, or the queue metrics in the shipper, which
// is where pending events are really queued when the proxy queue is
// in use.
func (b *broker) Metrics() (queue.Metrics, error) {
	return queue.Metrics{}, nil
}

func (b *broker) run() {
	var (
		nextEntryID  = queue.EntryID(0)
		pendingBatch = &batch{queue: b}
		pendingACKs  pendingACKsList
	)

	for {
		var pushChan chan pushRequest
		// Push requests are enabled if the pending batch isn't yet full.
		if len(pendingBatch.entries) < b.batchSize {
			pushChan = b.pushChan
		}

		var getChan chan getRequest
		// Get requests are enabled if the current pending batch is nonempty.
		if len(pendingBatch.entries) > 0 {
			getChan = b.getChan
		}

		select {
		case <-b.done:
			return

		case req := <-pushChan: // producer pushing new event
			req.responseChan <- nextEntryID
			req.producer.producedCount++
			pendingBatch.entries = append(pendingBatch.entries,
				queueEntry{
					event:    req.event,
					id:       nextEntryID,
					producer: req.producer,
				})
			nextEntryID++

		case req := <-getChan: // consumer asking for next batch
			acks := acksForBatch(pendingBatch)
			var doneChan chan struct{}
			if len(acks) > 0 {
				// We only actually save the ACK state and give it a done
				// channel if some producer in this batch has ACKs enabled.
				doneChan = make(chan struct{})
				pendingACKs.append(&batchACKState{
					doneChan: doneChan,
					acks:     acks,
				})
			}
			req.responseChan <- getResponse{
				ackChan: doneChan,
				entries: pendingBatch.entries,
			}
			// Reset the pending batch
			pendingBatch = &batch{queue: b}

		case <-pendingACKs.nextDoneChan(): // Previous batch being acknowledged
			batch := pendingACKs.pop()
			// Send any acknowledgments needed for this batch.
			total := 0
			for _, ack := range batch.acks {
				ack.producer.ackHandler(ack.count)
				total += ack.count
			}
			// Also update the ACKs in the metrics observer
			b.ackListener.OnACK(total)
		}
	}
}
