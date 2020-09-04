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
	"bytes"
	"os"

	"github.com/elastic/beats/v7/libbeat/logp"
)

type frameWriteRequest struct {
	// The frame to be written to disk.
	frame *writeFrame

	// The segment to which this frame should be written.
	segment *queueSegment
}

// One block for the writer loop consists of a write request and the
// segment it should be written to.
type writeRequest struct {
	frames []frameWriteRequest
}

// A writeResponse reports the list of segments that have been
// completely written and can be moved to segments.reading.
// A segment is determined to have been completely written
// (and is then closed by the writer loop) when a frameWriteRequest
// targets a different segment than the previous ones.
type writeResponse struct {
	completedSegments []*queueSegment
}

type writerLoop struct {
	logger *logp.Logger

	// The writer loop listens on requestChan for write blocks, and
	// writes them to disk immediately (all queue capacity checking etc. is
	// done by the core loop before sending it to the writer).
	requestChan chan writeRequest

	// The writer loop sends to responseChan when it has finished handling a
	// request, to signal the core loop that it is ready for the next one.
	responseChan chan writeResponse

	// The most recent segment that has been written to, if there is one.
	// This segment
	currentSegment *queueSegment

	// The file handle corresponding to currentSegment. When currentSegment
	// changes, this handle is closed and a new one is created.
	outputFile *os.File
}

func (wl *writerLoop) run() {
	for {
		block, ok := <-wl.requestChan
		if !ok {
			// The input channel is closed, we are done
			return
		}
		completedSegments := wl.processRequest(block)
		wl.responseChan <- writeResponse{
			completedSegments: completedSegments,
		}
	}
}

// Write the given data to disk, returns the list of segments that were
// completed in the process.
func (wl *writerLoop) processRequest(request writeRequest) []*queueSegment {
	var completedSegments []*queueSegment
	for _, frameRequest := range request.frames {
		// If the new segment doesn't match the last one, we need to open a new
		// file handle and possibly clean up the old one.
		if wl.currentSegment != frameRequest.segment {
			if wl.outputFile != nil {
				completedSegments = append(completedSegments, wl.currentSegment)
				wl.outputFile.Close()
				wl.outputFile = nil
				// TODO: try to sync?
			}
			wl.currentSegment = frameRequest.segment
			file, err := wl.currentSegment.getWriter()
			if err != nil {
				// TODO: retry, etc
			}
			wl.outputFile = file
		}

		// We have the data and a file to write it to. We are now committed
		// to writing this block unless the queue is closed in the meantime.
		_, err := wl.outputFile.Write(frameRequest.frame.serialized)
		// TODO: retry forever if there is an error or n isn't the right
		// length.
		if err != nil {
			wl.logger.Errorf("Couldn't write pending data to disk: %w", err)
		}
	}
	return completedSegments
}

// frameForContent wraps the given content buffer in a
// frame header / footer and returns the resulting larger buffer.
func frameForContent(
	frameContent []byte, checksumType ChecksumType,
) bytes.Buffer {
	buf := bytes.Buffer{}
	//checksum := computeChecksum(frameContent, checksumType)
	/*buf
	frameLength := len(frameContent) + frameMetadataSize;
	frameBytes := make([]byte, frameLength)
	frameWriter :=
	binary.Write(reader.raw, binary.LittleEndian, &frameLength)*/
	return buf
}

/*type writerState struct {
	// The open file handle for the segment currently being written.
	// This should be non-nil if and only if diskQueue.segments.writing is.
	file         *os.File
	filePosition int64
}*/

/*func handleFrame(dq *diskQueue, state *writerState, frame bytes.Buffer) {
	dq.segments.Lock()
	defer dq.segments.Unlock()

	frameLen := uint64(frame.Len())
	// If there isn't enough space left in the current segment, close the
	// segment's file handle and move it to the reading list.
	if dq.segments.writing != nil &&
		dq.segments.writing.size+frameLen > dq.settings.MaxSegmentSize {
		// TODO: make reasonable efforts to sync to disk.
		state.file.Close()
		dq.segments.reading = append(dq.segments.reading, dq.segments.writing)
		dq.segments.writing = nil
	}

	if dq.segments.writing == nil {
		// There is no active writing segment, create one.
		// TODO: (actually create one)
	}

	// TODO: try to delete dq.segments.acked

	currentSize := dq.segments.sizeOnDiskWithLock()
	// Block (releasing the dq.segments lock) until
	// currentSize + frameLen <= dq.settings.MaxBufferSize
	for currentSize+frameLen > dq.settings.MaxBufferSize {
		// Wait for some space to be freed.
		dq.segments.segmentDeletedCond.Wait()
		if dq.closedForWrite.Load() {
			// The queue is closed, abort
		}
	}

	// We now have a frame we want to write to disk, and enough free capacity
	// to write it.
	writeAll(state.file, frame.Bytes())
}*/

// The writer loop's job is to continuously read a data frame from the
// queue's intake channel, if there is one, and write it to disk.
// It continues until the intake channel is empty or
/*func (dq *diskQueue) writerLoop() {
	defer dq.waitGroup.Done()
	//logger := dq.settings.Logger.Named("writerLoop")
	state := &writerState{}

	for {
		if dq.abort.Load() {
			// We are aborting, ignore any remaining buffered frames.
			return
		}
		select {
		case frameContent := <-dq.inChan:
			if frameContent == nil {
				// The channel has been drained, the writer loop should shut down.
				return
			}
			frameBuffer := frameForContent(frameContent, dq.settings.ChecksumType)
			handleFrame(dq, state, frameBuffer)
			if !dq.abort.Load() {
				// As long as we aren't aborting, continue processing any pending
				// frames.
				continue
			}
		case <-dq.done:
		}
		// We've processed
	}
}*/
/*
func writeAll(writer io.Writer, p []byte) (int, error) {
	var N int
	for len(p) > 0 {
		n, err := writer.Write(p)
		N, p = N+n, p[n:]
		if err != nil && isRetryErr(err) {
			return N, err
		}
	}
	return N, nil
}

func isRetryErr(err error) bool {
	return err == syscall.EINTR || err == syscall.EAGAIN
}
*/
