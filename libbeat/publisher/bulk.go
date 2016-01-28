package publisher

import (
	"sync"
	"time"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/outputs"
)

type bulkWorker struct {
	output worker
	wg     *sync.WaitGroup

	queue       chan message
	bulkQueue   chan message
	guaranteed  bool
	flushTicker *time.Ticker

	maxBatchSize int
	events       []common.MapStr    // batched events
	pending      []outputs.Signaler // pending signalers for batched events
}

func newBulkWorker(
	wg *sync.WaitGroup, hwm int, bulkHWM int,
	output worker,
	flushInterval time.Duration,
	maxBatchSize int,
) *bulkWorker {
	b := &bulkWorker{
		output:       output,
		wg:           wg,
		queue:        make(chan message, hwm),
		bulkQueue:    make(chan message, bulkHWM),
		flushTicker:  time.NewTicker(flushInterval),
		maxBatchSize: maxBatchSize,
		events:       make([]common.MapStr, 0, maxBatchSize),
		pending:      nil,
	}

	wg.Add(1)
	go b.run()
	return b
}

func (b *bulkWorker) send(m message) {
	if m.events == nil {
		b.wg.Add(1)
		b.queue <- m
	} else {
		b.wg.Add(len(m.events))
		b.bulkQueue <- m
	}
}

func (b *bulkWorker) run() {
	defer func() {
		b.wg.Done()
	}()

	var queueClosed bool
	var bulkQueueClosed bool

	for {
		select {
		case m, ok := <-b.queue:
			if !ok {
				queueClosed = true
			} else {
				b.onEvent(&m.context, m.event)
			}
		case m, ok := <-b.bulkQueue:
			if !ok {
				bulkQueueClosed = true
			} else {
				b.onEvents(&m.context, m.events)
			}
		case <-b.flushTicker.C:
			b.flush()
		}
		if queueClosed && bulkQueueClosed {
			b.flush()
			return
		}
	}
}

func (b *bulkWorker) flush() {
	if len(b.events) > 0 {
		b.publish()
	}
}

func (b *bulkWorker) onEvent(ctx *Context, event common.MapStr) {
	b.events = append(b.events, event)
	b.guaranteed = b.guaranteed || ctx.Guaranteed

	signal := ctx.Signal
	if signal != nil {
		b.pending = append(b.pending, signal)
	}

	if len(b.events) == cap(b.events) {
		b.publish()
	}
}

func (b *bulkWorker) onEvents(ctx *Context, events []common.MapStr) {
	for len(events) > 0 {
		// split up bulk to match required bulk sizes.
		// If input events have been split up bufferFull will be set and
		// bulk request will be published.
		spaceLeft := cap(b.events) - len(b.events)
		consume := len(events)
		bufferFull := spaceLeft <= consume
		signal := ctx.Signal
		b.guaranteed = b.guaranteed || ctx.Guaranteed
		if spaceLeft < consume {
			consume = spaceLeft
			if signal != nil {
				// creating cascading signaler chain for
				// subset of events being send
				signal = outputs.NewSplitSignaler(signal, 2)
			}
		}

		// buffer events
		b.events = append(b.events, events[:consume]...)
		events = events[consume:]
		if signal != nil {
			b.pending = append(b.pending, signal)
		}

		if bufferFull {
			b.publish()
		}
	}
}

func (b *bulkWorker) publish() {
	// TODO: remember/merge and forward context options to output worker
	b.output.send(message{
		context: Context{
			publishOptions: publishOptions{Guaranteed: b.guaranteed},
			Signal:         outputs.NewCompositeSignaler(b.pending...),
		},
		event:  nil,
		events: b.events,
	})

	b.wg.Add(-len(b.events))
	b.pending = nil
	b.guaranteed = false
	b.events = make([]common.MapStr, 0, b.maxBatchSize)
}

func (b *bulkWorker) shutdown() {
	b.flushTicker.Stop()
	close(b.queue)
	close(b.bulkQueue)
	b.wg.Wait()
}
