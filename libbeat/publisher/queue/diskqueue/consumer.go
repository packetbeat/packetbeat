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

package diskqueue

import (
	"fmt"

	"github.com/elastic/beats/v7/libbeat/publisher"
	"github.com/elastic/beats/v7/libbeat/publisher/queue"
)

type diskQueueConsumer struct {
	queue  *diskQueue
	closed bool
}

type diskQueueBatch struct {
	queue  *diskQueue
	events []publisher.Event
	ackIDs []frameID
}

//
// diskQueueConsumer implementation of the queue.Consumer interface
//

func (consumer *diskQueueConsumer) Get(eventCount int) (queue.Batch, error) {
	if consumer.closed {
		return nil, fmt.Errorf("Tried to read from closed disk queue consumer")
	}
	// Read at least one frame. This is guaranteed to eventually
	// succeed unless the queue is closed.
	frame, ok := <-consumer.queue.readerLoop.output
	if !ok {
		return nil, fmt.Errorf("Tried to read from a closed disk queue")
	}
	frames := []*readFrame{frame}
eventLoop:
	for eventCount <= 0 || len(frames) < eventCount {
		select {
		case frame, ok = <-consumer.queue.readerLoop.output:
			if !ok {
				// The queue was closed while we were reading it, just send back
				// what we have so far.
				break eventLoop
			}
			frames = append(frames, frame)
		default:
			// We can't read any more frames without blocking, so send back
			// what we have now.
			break eventLoop
		}
	}

	var ackIDs []frameID
	var events []publisher.Event
	for _, frame := range frames {
		events = append(events, frame.event)
		ackIDs = append(ackIDs, frame.id)
	}

	return &diskQueueBatch{
		queue:  consumer.queue,
		events: events,
		ackIDs: ackIDs,
	}, nil
}

func (consumer *diskQueueConsumer) Close() error {
	consumer.closed = true
	return nil
}

//
// diskQueueBatch implementation of the queue.Batch interface
//

func (batch *diskQueueBatch) Events() []publisher.Event {
	return batch.events
}

// This is the only place that the queue state is changed from
// outside the core loop. This is because ACKs are messy and bursty
// and we don't want the core loop to bottleneck on manipulating
// a potentially large dictionary, so we use a lock and let
// consumer threads handle most of the processing themselves.
// TODO: this shouldn't really be a dictionary, use a bitfield or
// something more efficient.
func (batch *diskQueueBatch) ACK() {
	dq := batch.queue
	dq.ackLock.Lock()
	defer dq.ackLock.Unlock()
	for _, frameID := range batch.ackIDs {
		dq.acked[frameID] = true
	}
	if dq.acked[dq.ackedUpTo] {
		for ; dq.acked[dq.ackedUpTo]; dq.ackedUpTo++ {
			delete(dq.acked, dq.ackedUpTo)
		}
		// It would be considerate to send this less frequently, so
		// as not to bother the core loop with messages that are
		// usually no-ops.
		// TODO: only inform the core loop when we cross a segment
		// boundary.
		dq.consumerACKChan <- dq.ackedUpTo
	}
}
