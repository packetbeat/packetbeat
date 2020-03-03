// Copyright 2019+ Klaus Post. All rights reserved.
// License information can be found in the LICENSE file.
// Based on work by Yann Collet, released under BSD License.

package zstd

import (
	"crypto/rand"
	"fmt"
	"io"
	rdebug "runtime/debug"
	"sync"

	"github.com/klauspost/compress/zstd/internal/xxhash"
)

// Encoder provides encoding to Zstandard.
// An Encoder can be used for either compressing a stream via the
// io.WriteCloser interface supported by the Encoder or as multiple independent
// tasks via the EncodeAll function.
// Smaller encodes are encouraged to use the EncodeAll function.
// Use NewWriter to create a new instance.
type Encoder struct {
	o        encoderOptions
	encoders chan encoder
	state    encoderState
	init     sync.Once
}

type encoder interface {
	Encode(blk *blockEnc, src []byte)
	Block() *blockEnc
	CRC() *xxhash.Digest
	AppendCRC([]byte) []byte
	WindowSize(size int) int32
	UseBlock(*blockEnc)
	Reset()
}

type encoderState struct {
	w             io.Writer
	filling       []byte
	current       []byte
	previous      []byte
	encoder       encoder
	writing       *blockEnc
	err           error
	writeErr      error
	nWritten      int64
	headerWritten bool
	eofWritten    bool

	// This waitgroup indicates an encode is running.
	wg sync.WaitGroup
	// This waitgroup indicates we have a block encoding/writing.
	wWg sync.WaitGroup
}

// NewWriter will create a new Zstandard encoder.
// If the encoder will be used for encoding blocks a nil writer can be used.
func NewWriter(w io.Writer, opts ...EOption) (*Encoder, error) {
	initPredefined()
	var e Encoder
	e.o.setDefault()
	for _, o := range opts {
		err := o(&e.o)
		if err != nil {
			return nil, err
		}
	}
	if w != nil {
		e.Reset(w)
	} else {
		e.init.Do(func() {
			e.initialize()
		})
	}
	return &e, nil
}

func (e *Encoder) initialize() {
	e.encoders = make(chan encoder, e.o.concurrent)
	for i := 0; i < e.o.concurrent; i++ {
		e.encoders <- e.o.encoder()
	}
}

// Reset will re-initialize the writer and new writes will encode to the supplied writer
// as a new, independent stream.
func (e *Encoder) Reset(w io.Writer) {
	e.init.Do(func() {
		e.initialize()
	})
	s := &e.state
	s.wg.Wait()
	s.wWg.Wait()
	if cap(s.filling) == 0 {
		s.filling = make([]byte, 0, e.o.blockSize)
	}
	if cap(s.current) == 0 {
		s.current = make([]byte, 0, e.o.blockSize)
	}
	if cap(s.previous) == 0 {
		s.previous = make([]byte, 0, e.o.blockSize)
	}
	if s.encoder == nil {
		s.encoder = e.o.encoder()
	}
	if s.writing == nil {
		s.writing = &blockEnc{}
		s.writing.init()
	}
	s.writing.initNewEncode()
	s.filling = s.filling[:0]
	s.current = s.current[:0]
	s.previous = s.previous[:0]
	s.encoder.Reset()
	s.headerWritten = false
	s.eofWritten = false
	s.w = w
	s.err = nil
	s.nWritten = 0
	s.writeErr = nil
}

// Write data to the encoder.
// Input data will be buffered and as the buffer fills up
// content will be compressed and written to the output.
// When done writing, use Close to flush the remaining output
// and write CRC if requested.
func (e *Encoder) Write(p []byte) (n int, err error) {
	s := &e.state
	for len(p) > 0 {
		if len(p)+len(s.filling) < e.o.blockSize {
			if e.o.crc {
				_, _ = s.encoder.CRC().Write(p)
			}
			s.filling = append(s.filling, p...)
			return n + len(p), nil
		}
		add := p
		if len(p)+len(s.filling) > e.o.blockSize {
			add = add[:e.o.blockSize-len(s.filling)]
		}
		if e.o.crc {
			_, _ = s.encoder.CRC().Write(add)
		}
		s.filling = append(s.filling, add...)
		p = p[len(add):]
		n += len(add)
		if len(s.filling) < e.o.blockSize {
			return n, nil
		}
		err := e.nextBlock(false)
		if err != nil {
			return n, err
		}
		if debug && len(s.filling) > 0 {
			panic(len(s.filling))
		}
	}
	return n, nil
}

// nextBlock will synchronize and start compressing input in e.state.filling.
// If an error has occurred during encoding it will be returned.
func (e *Encoder) nextBlock(final bool) error {
	s := &e.state
	// Wait for current block.
	s.wg.Wait()
	if s.err != nil {
		return s.err
	}
	if len(s.filling) > e.o.blockSize {
		return fmt.Errorf("block > maxStoreBlockSize")
	}
	if !s.headerWritten {
		var tmp [maxHeaderSize]byte
		fh := frameHeader{
			ContentSize:   0,
			WindowSize:    uint32(s.encoder.WindowSize(0)),
			SingleSegment: false,
			Checksum:      e.o.crc,
			DictID:        0,
		}
		dst, err := fh.appendTo(tmp[:0])
		if err != nil {
			return err
		}
		s.headerWritten = true
		s.wWg.Wait()
		var n2 int
		n2, s.err = s.w.Write(dst)
		if s.err != nil {
			return s.err
		}
		s.nWritten += int64(n2)
	}
	if s.eofWritten {
		// Ensure we only write it once.
		final = false
	}

	if len(s.filling) == 0 {
		// Final block, but no data.
		if final {
			enc := s.encoder
			blk := enc.Block()
			blk.reset(nil)
			blk.last = true
			blk.encodeRaw(nil)
			s.wWg.Wait()
			_, s.err = s.w.Write(blk.output)
			s.nWritten += int64(len(blk.output))
			s.eofWritten = true
		}
		return s.err
	}

	// Move blocks forward.
	s.filling, s.current, s.previous = s.previous[:0], s.filling, s.current
	s.wg.Add(1)
	go func(src []byte) {
		if debug {
			println("Adding block,", len(src), "bytes, final:", final)
		}
		defer func() {
			if r := recover(); r != nil {
				s.err = fmt.Errorf("panic while encoding: %v", r)
				rdebug.PrintStack()
			}
			s.wg.Done()
		}()
		enc := s.encoder
		blk := enc.Block()
		enc.Encode(blk, src)
		blk.last = final
		if final {
			s.eofWritten = true
		}
		// Wait for pending writes.
		s.wWg.Wait()
		if s.writeErr != nil {
			s.err = s.writeErr
			return
		}
		// Transfer encoders from previous write block.
		blk.swapEncoders(s.writing)
		// Transfer recent offsets to next.
		enc.UseBlock(s.writing)
		s.writing = blk
		s.wWg.Add(1)
		go func() {
			defer func() {
				if r := recover(); r != nil {
					s.writeErr = fmt.Errorf("panic while encoding/writing: %v", r)
					rdebug.PrintStack()
				}
				s.wWg.Done()
			}()
			err := errIncompressible
			// If we got the exact same number of literals as input,
			// assume the literals cannot be compressed.
			if len(src) != len(blk.literals) || len(src) != e.o.blockSize {
				err = blk.encode()
			}
			switch err {
			case errIncompressible:
				if debug {
					println("Storing incompressible block as raw")
				}
				blk.encodeRaw(src)
				// In fast mode, we do not transfer offsets, so we don't have to deal with changing the.
			case nil:
			default:
				s.writeErr = err
				return
			}
			_, s.writeErr = s.w.Write(blk.output)
			s.nWritten += int64(len(blk.output))
		}()
	}(s.current)
	return nil
}

// ReadFrom reads data from r until EOF or error.
// The return value n is the number of bytes read.
// Any error except io.EOF encountered during the read is also returned.
//
// The Copy function uses ReaderFrom if available.
func (e *Encoder) ReadFrom(r io.Reader) (n int64, err error) {
	if debug {
		println("Using ReadFrom")
	}
	// Maybe handle stuff queued?
	e.state.filling = e.state.filling[:e.o.blockSize]
	src := e.state.filling
	for {
		n2, err := r.Read(src)
		_, _ = e.state.encoder.CRC().Write(src[:n2])
		// src is now the unfilled part...
		src = src[n2:]
		n += int64(n2)
		switch err {
		case io.EOF:
			e.state.filling = e.state.filling[:len(e.state.filling)-len(src)]
			if debug {
				println("ReadFrom: got EOF final block:", len(e.state.filling))
			}
			return n, e.nextBlock(true)
		default:
			if debug {
				println("ReadFrom: got error:", err)
			}
			e.state.err = err
			return n, err
		case nil:
		}
		if len(src) > 0 {
			if debug {
				println("ReadFrom: got space left in source:", len(src))
			}
			continue
		}
		err = e.nextBlock(false)
		if err != nil {
			return n, err
		}
		e.state.filling = e.state.filling[:e.o.blockSize]
		src = e.state.filling
	}
}

// Flush will send the currently written data to output
// and block until everything has been written.
// This should only be used on rare occasions where pushing the currently queued data is critical.
func (e *Encoder) Flush() error {
	s := &e.state
	if len(s.filling) > 0 {
		err := e.nextBlock(false)
		if err != nil {
			return err
		}
	}
	s.wg.Wait()
	s.wWg.Wait()
	if s.err != nil {
		return s.err
	}
	return s.writeErr
}

// Close will flush the final output and close the stream.
// The function will block until everything has been written.
// The Encoder can still be re-used after calling this.
func (e *Encoder) Close() error {
	s := &e.state
	if s.encoder == nil {
		return nil
	}
	err := e.nextBlock(true)
	if err != nil {
		return err
	}
	s.wg.Wait()
	s.wWg.Wait()

	if s.err != nil {
		return s.err
	}
	if s.writeErr != nil {
		return s.writeErr
	}

	// Write CRC
	if e.o.crc && s.err == nil {
		// heap alloc.
		var tmp [4]byte
		_, s.err = s.w.Write(s.encoder.AppendCRC(tmp[:0]))
		s.nWritten += 4
	}

	// Add padding with content from crypto/rand.Reader
	if s.err == nil && e.o.pad > 0 {
		add := calcSkippableFrame(s.nWritten, int64(e.o.pad))
		frame, err := skippableFrame(s.filling[:0], add, rand.Reader)
		if err != nil {
			return err
		}
		_, s.err = s.w.Write(frame)
	}
	return s.err
}

// EncodeAll will encode all input in src and append it to dst.
// This function can be called concurrently, but each call will only run on a single goroutine.
// If empty input is given, nothing is returned, unless WithZeroFrames is specified.
// Encoded blocks can be concatenated and the result will be the combined input stream.
// Data compressed with EncodeAll can be decoded with the Decoder,
// using either a stream or DecodeAll.
func (e *Encoder) EncodeAll(src, dst []byte) []byte {
	if len(src) == 0 {
		if e.o.fullZero {
			// Add frame header.
			fh := frameHeader{
				ContentSize:   0,
				WindowSize:    MinWindowSize,
				SingleSegment: true,
				// Adding a checksum would be a waste of space.
				Checksum: false,
				DictID:   0,
			}
			dst, _ = fh.appendTo(dst)

			// Write raw block as last one only.
			var blk blockHeader
			blk.setSize(0)
			blk.setType(blockTypeRaw)
			blk.setLast(true)
			dst = blk.appendTo(dst)
		}
		return dst
	}
	e.init.Do(func() {
		e.o.setDefault()
		e.initialize()
	})
	enc := <-e.encoders
	defer func() {
		// Release encoder reference to last block.
		enc.Reset()
		e.encoders <- enc
	}()
	enc.Reset()
	blk := enc.Block()
	single := len(src) > 1<<20
	if e.o.single != nil {
		single = *e.o.single
	}
	fh := frameHeader{
		ContentSize:   uint64(len(src)),
		WindowSize:    uint32(enc.WindowSize(len(src))),
		SingleSegment: single,
		Checksum:      e.o.crc,
		DictID:        0,
	}

	// If less than 1MB, allocate a buffer up front.
	if len(dst) == 0 && cap(dst) == 0 && len(src) < 1<<20 {
		dst = make([]byte, 0, len(src))
	}
	dst, err := fh.appendTo(dst)
	if err != nil {
		panic(err)
	}

	for len(src) > 0 {
		todo := src
		if len(todo) > e.o.blockSize {
			todo = todo[:e.o.blockSize]
		}
		src = src[len(todo):]
		if e.o.crc {
			_, _ = enc.CRC().Write(todo)
		}
		blk.reset(nil)
		blk.pushOffsets()
		enc.Encode(blk, todo)
		if len(src) == 0 {
			blk.last = true
		}
		err := errIncompressible
		// If we got the exact same number of literals as input,
		// assume the literals cannot be compressed.
		if len(blk.literals) != len(todo) || len(todo) != e.o.blockSize {
			err = blk.encode()
		}

		switch err {
		case errIncompressible:
			if debug {
				println("Storing incompressible block as raw")
			}
			blk.encodeRaw(todo)
			blk.popOffsets()
		case nil:
		default:
			panic(err)
		}
		dst = append(dst, blk.output...)
	}
	if e.o.crc {
		dst = enc.AppendCRC(dst)
	}
	// Add padding with content from crypto/rand.Reader
	if e.o.pad > 0 {
		add := calcSkippableFrame(int64(len(dst)), int64(e.o.pad))
		dst, err = skippableFrame(dst, add, rand.Reader)
		if err != nil {
			panic(err)
		}
	}
	return dst
}
