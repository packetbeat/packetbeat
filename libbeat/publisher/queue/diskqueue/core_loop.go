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

import "github.com/elastic/beats/v7/libbeat/publisher"

// A frame waiting to be written to disk
type writeFrame struct {
	// The original event provided by the client to diskQueueProducer
	event publisher.Event

	// The event, serialized for writing to disk and wrapped in a frame
	// header / footer.
	serialized []byte
}

// A frame that has been read from disk
type readFrame struct {
}

// A request sent from a producer to the core loop to add a frame to the queue.
//
type writeRequest struct {
	frame        *writeFrame
	shouldBlock  bool
	responseChan chan bool
}

// A readRequest is sent from the reader loop to the core loop when it
// needs a new segment file to read.
type readRequest struct {
	responseChan chan *readResponse
}

type readResponse struct {
}

type cancelRequest struct {
	producer *diskQueueProducer
	// If producer.config.DropOnCancel is true, then the core loop will respond
	// on responseChan with the number of dropped events.
	// Otherwise, this field may be nil.
	responseChan chan int
}

type coreLoop struct {
	// The queue that created this coreLoop. The core loop is the only one of
	// the main goroutines for the queue that has a pointer to the queue and
	// understands its logic / structure.
	// Possible TODO: the split between fields in coreLoop and fields in
	// diskQueue seems artificial. Maybe the values here should be promoted
	// to diskQueue fields, and the core loop should just be a function on
	// diskQueue.
	queue *diskQueue

	// writing is true if a writeRequest is currently being processed by the
	// writer loop, false otherwise.
	writing bool

	// reading is true if the reader loop is processing a readBlock, false
	// otherwise.
	reading bool

	// deleting is true if the segment-deletion loop is processing a deletion
	// request, false otherwise.
	deleting bool

	// pendingWrites is a list of all write requests that have been accepted
	// by the queue and are waiting to be written to disk.
	pendingWrites []*writeRequest

	// blockedWrites is a list of all write requests that are waiting for
	// free space in the queue.
	blockedWrites []*writeRequest

	// This value represents the oldest frame ID for a segment that has not
	// yet been moved to the acked list. It is used to detect when the oldest
	// outstanding segment has been fully acknowledged by the consumer.
	oldestFrameID frameID
}

func (cl *coreLoop) run() {
	dq := cl.queue

	for {
		select {
		///////////////////////////////////////
		// Endpoints used by the public API
		case writeRequest := <-dq.writeRequestChan:
			cl.handleProducerWriteRequest(writeRequest)

		case cancelRequest := <-dq.cancelRequestChan:
			cl.handleProducerCancelRequest(cancelRequest)

		case ackedUpTo := <-dq.consumerAckChan:
			cl.handleConsumerAck(ackedUpTo)

		///////////////////////////////////////
		// Writer loop handling
		case <-dq.writerLoop.finishedWriting:
			if len(cl.pendingWrites) > 0 {
				cl.forwardWriteRequest(cl.pendingWrites[0])
				cl.pendingWrites = cl.pendingWrites[1:]
			} else {
				cl.writing = false
			}

		///////////////////////////////////////
		// Reader loop handling
		case readResponse := <-dq.readerLoop.finishedReading:

		///////////////////////////////////////
		// Deleter loop handling
		case deleteResponse := <-dq.deleterLoop.response:
			if len(deleteResponse.deleted) > 0 {
				// One or more segments were deleted, recompute the outstanding list.
				newAckedSegments := []*queueSegment{}
				for _, segment := range dq.segments.acked {
					if !deleteResponse.deleted[segment] {
						// This segment wasn't deleted, so it goes in the new list.
						newAckedSegments = append(newAckedSegments, segment)
					}
				}
				dq.segments.acked = newAckedSegments
			}
			if len(deleteResponse.errors) > 0 {
				dq.settings.Logger.Errorw("Couldn't delete old segment files",
					"errors", deleteResponse.errors)
			}

			if len(dq.segments.acked) > 0 {
				// There are still (possibly new) segments to delete, send the
				// next batch.
				dq.deleterLoop.input <- &deleteRequest{segments: dq.segments.acked}
			} else {
				// Nothing more to delete for now, update the deleting flag.
				cl.deleting = false
			}
		}
	}
}

func (cl *coreLoop) forwardWriteRequest(request *writeRequest) {
	dq := cl.queue
	// First we must decide which segment the new frame should be written to.
	data := request.frame.serialized
	segment := dq.segments.writing

	if segment != nil &&
		segment.size+uint64(len(data)) > dq.settings.MaxSegmentSize {
		// The new frame is too big to fit in this segment, so close it and
		// move it to the read queue.
		segment.writer.Close()
		// TODO: make reasonable attempts to sync the closed file.
		dq.segments.reading = append(dq.segments.reading, segment)
		segment = nil
	}

	// If we don't have a segment, we need to create one.
	if segment == nil {
		segment = &queueSegment{id: dq.segments.nextID}
	}
}

func (cl *coreLoop) handleProducerWriteRequest(request *writeRequest) {
	if len(cl.blockedWrites) > 0 {
		// If other requests are still waiting for space, then there
		// definitely isn't enough for this one.
		if request.shouldBlock {
			cl.blockedWrites = append(cl.blockedWrites, request)
		} else {
			// If the request is non-blocking, send immediate failure and discard it.
			request.responseChan <- false
		}
		return
	}
	// We will accept this request if there is enough capacity left in
	// the queue (after accounting for the pending writes that were
	// already accepted).
	pendingBytes := uint64(0)
	for _, request := range cl.pendingWrites {
		pendingBytes += uint64(len(request.frame.serialized))
	}
	currentSize := pendingBytes + cl.queue.segments.sizeOnDisk()
	frameSize := uint64(len(request.frame.serialized))
	if currentSize+frameSize > cl.queue.settings.MaxBufferSize {
		// The queue is too full
	}
}

func (cl *coreLoop) handleProducerCancelRequest(request *cancelRequest) {
}

func (cl *coreLoop) handleConsumerAck(ackedUpTo frameID) {
	acking := cl.queue.segments.acking
	if len(acking) == 0 {
		return
	}
	segmentsAcked := 0
	startFrame := cl.oldestFrameID
	for ; segmentsAcked < len(acking); segmentsAcked++ {
		segment := acking[segmentsAcked]
		endFrame := startFrame + frameID(segment.framesRead)
		if endFrame > ackedUpTo {
			// This segment has not been fully read, we're done.
			break
		}
	}
	if segmentsAcked > 0 {
		// Move fully acked segments to the acked list and remove them
		// from the acking list.
		cl.queue.segments.acked =
			append(cl.queue.segments.acked, acking[:segmentsAcked]...)
		cl.queue.segments.acking = acking[segmentsAcked:]
		cl.maybeDeleteAcked()
	}
}

// If the acked list is nonempty, and there are no outstanding deletion
// requests, send one.
func (cl *coreLoop) maybeDeleteAcked() {
	if !cl.deleting && len(cl.queue.segments.acked) > 0 {
		cl.queue.deleterLoop.input <- &deleteRequest{segments: cl.queue.segments.acked}
		cl.deleting = true
	}
}
