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

package queuetest

import (
	"fmt"
	"sync"
	"testing"

	"github.com/elastic/beats/v7/libbeat/publisher/queue"
	"github.com/elastic/elastic-agent-libs/mapstr"
)

// QueueFactory is used to create a per test queue instance.
type QueueFactory func(t *testing.T) queue.Queue

type workerFactory func(*sync.WaitGroup, interface{}, *TestLogger, queue.Queue) func()

type testCase struct {
	name                 string
	producers, consumers workerFactory
}

func TestSingleProducerConsumer(
	t *testing.T,
	events, batchSize int,
	queueFactory QueueFactory,
) {
	tests := []testCase{
		{
			"single producer, consumer without ack, complete batches",
			makeProducer(events, false, countEvent),
			makeConsumer(events, -1),
		},
		{
			"single producer, consumer, without ack, limited batches",
			makeProducer(events, false, countEvent),
			makeConsumer(events, batchSize),
		},
		{

			"single producer, consumer, with ack, complete batches",
			makeProducer(events, true, countEvent),
			makeConsumer(events, -1),
		},
		{
			"single producer, consumer, with ack, limited batches",
			makeProducer(events, true, countEvent),
			makeConsumer(events, batchSize),
		},
	}

	runTestCases(t, tests, queueFactory)
}

func TestMultiProducerConsumer(
	t *testing.T,
	events, batchSize int,
	queueFactory QueueFactory,
) {
	tests := []testCase{
		/*{
			"2 producers, 1 consumer, without ack, complete batches",
			multiple(
				makeProducer(events, false, countEvent),
				makeProducer(events, false, countEvent),
			),
			makeConsumer(events*2, -1),
		},*/
		{
			"2 producers, 1 consumer, all ack, complete batches",
			multiple(
				makeProducer(events, true, countEvent),
				makeProducer(events, true, countEvent),
			),
			makeConsumer(events*2, -1),
		},
		/*{
			"2 producers, 1 consumer, 1 ack, complete batches",
			multiple(
				makeProducer(events, true, countEvent),
				makeProducer(events, false, countEvent),
			),
			makeConsumer(events*2, -1),
		},
		{
			"2 producers, 1 consumer, without ack, limited batches",
			multiple(
				makeProducer(events, false, countEvent),
				makeProducer(events, false, countEvent),
			),
			makeConsumer(events*2, batchSize),
		},
		{
			"2 producers, 1 consumer, all ack, limited batches",
			multiple(
				makeProducer(events, true, countEvent),
				makeProducer(events, true, countEvent),
			),
			makeConsumer(events*2, batchSize),
		},
		{
			"2 producers, 1 consumer, 1 ack, limited batches",
			multiple(
				makeProducer(events, true, countEvent),
				makeProducer(events, false, countEvent),
			),
			makeConsumer(events*2, batchSize),
		},

		{
			"1 producer, 2 consumers, without ack, complete batches",
			makeProducer(events, true, countEvent),
			multiConsumer(2, events, -1),
		},
		{
			"1 producer, 2 consumers, without ack, limited batches",
			makeProducer(events, true, countEvent),
			multiConsumer(2, events, -1),
		},

		{
			"2 producers, 2 consumer, without ack, complete batches",
			multiple(
				makeProducer(events, false, countEvent),
				makeProducer(events, false, countEvent),
			),
			multiConsumer(2, events*2, -1),
		},
		{
			"2 producers, 2 consumer, all ack, complete batches",
			multiple(
				makeProducer(events, true, countEvent),
				makeProducer(events, true, countEvent),
			),
			multiConsumer(2, events*2, -1),
		},
		{
			"2 producers, 2 consumer, 1 ack, complete batches",
			multiple(
				makeProducer(events, true, countEvent),
				makeProducer(events, false, countEvent),
			),
			multiConsumer(2, events*2, -1),
		},
		{
			"2 producers, 2 consumer, without ack, limited batches",
			multiple(
				makeProducer(events, false, countEvent),
				makeProducer(events, false, countEvent),
			),
			multiConsumer(2, events*2, batchSize),
		},
		{
			"2 producers, 2 consumer, all ack, limited batches",
			multiple(
				makeProducer(events, true, countEvent),
				makeProducer(events, true, countEvent),
			),
			multiConsumer(2, events*2, batchSize),
		},
		{
			"2 producers, 2 consumer, 1 ack, limited batches",
			multiple(
				makeProducer(events, true, countEvent),
				makeProducer(events, false, countEvent),
			),
			multiConsumer(2, events*2, batchSize),
		},*/
	}

	runTestCases(t, tests, queueFactory)
}

func runTestCases(t *testing.T, tests []testCase, queueFactory QueueFactory) {
	for _, test := range tests {
		verbose := testing.Verbose()
		t.Run(test.name, withOptLogOutput(verbose, func(t *testing.T) {
			if !verbose {
				t.Parallel()
			}

			log := NewTestLogger(t)
			log.Debug("run test: ", test.name)

			queue := queueFactory(t)
			defer func() {
				err := queue.Close()
				if err != nil {
					t.Error(err)
				}
			}()

			var wg sync.WaitGroup

			go test.producers(&wg, nil, log, queue)()
			go test.consumers(&wg, nil, log, queue)()

			fmt.Printf("waiting on wg\n")
			wg.Wait()
			fmt.Printf("done waiting\n")
		}))
	}
}

func multiple(
	fns ...workerFactory,
) workerFactory {
	return func(wg *sync.WaitGroup, info interface{}, log *TestLogger, queue queue.Queue) func() {
		runners := make([]func(), len(fns))
		for i, gen := range fns {
			runners[i] = gen(wg, info, log, queue)
		}

		return func() {
			for _, r := range runners {
				go r()
			}
		}
	}
}

func makeProducer(
	maxEvents int,
	waitACK bool,
	makeFields func(int) mapstr.M,
) func(*sync.WaitGroup, interface{}, *TestLogger, queue.Queue) func() {
	return func(wg *sync.WaitGroup, info interface{}, log *TestLogger, b queue.Queue) func() {
		wg.Add(1)
		return func() {
			defer wg.Done()

			fmt.Printf("start producer\n")
			defer fmt.Printf("stop producer\n")
			log.Debug("start producer")
			defer log.Debug("stop producer")

			var (
				ackWG sync.WaitGroup
				ackCB func(int)
			)

			if waitACK {
				fmt.Printf("waitACK is true\n")
				ackWG.Add(maxEvents)

				total := 0
				ackCB = func(N int) {
					fmt.Printf("ackCB(%v)\n", N)
					total += N
					log.Debugf("producer ACK: N=%v, total=%v\n", N, total)

					for i := 0; i < N; i++ {
						ackWG.Done()
					}
				}
			}

			producer := b.Producer(queue.ProducerConfig{
				ACK: ackCB,
			})
			for i := 0; i < maxEvents; i++ {
				log.Debug("publish event", i)
				producer.Publish(makeEvent(makeFields(i)))
			}

			ackWG.Wait()
		}
	}
}

func makeConsumer(maxEvents, batchSize int) workerFactory {
	return multiConsumer(1, maxEvents, batchSize)
}

func multiConsumer(numConsumers, maxEvents, batchSize int) workerFactory {
	return func(wg *sync.WaitGroup, info interface{}, log *TestLogger, b queue.Queue) func() {
		wg.Add(1)
		return func() {
			defer wg.Done()

			var events sync.WaitGroup

			log.Debugf("consumer: wait for %v events\n", maxEvents)
			events.Add(maxEvents)

			for i := 0; i < numConsumers; i++ {
				b := b

				go func() {
					fmt.Printf("start consumer\n")
					defer fmt.Printf("end consumer\n")
					for {
						batch, err := b.Get(batchSize)
						if err != nil {
							return
						}

						collected := batch.Events()
						fmt.Printf("got batch of size %d\n", len(collected))
						log.Debug("consumer: process batch", len(collected))

						for range collected {
							events.Done()
						}
						batch.ACK()
						fmt.Printf("batch acked\n")
					}
				}()
			}

			fmt.Printf("waiting on events wg\n")
			events.Wait()
			fmt.Printf("done waiting on events wg\n")
		}
	}
}

func countEvent(i int) mapstr.M {
	return mapstr.M{
		"count": i,
	}
}
