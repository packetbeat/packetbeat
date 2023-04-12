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

package pipeline

import (
	"sync"
	"testing"
	"testing/quick"
	"time"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common/atomic"
	"github.com/elastic/beats/v7/libbeat/internal/testutil"
	"github.com/elastic/beats/v7/libbeat/outputs"
	"github.com/elastic/beats/v7/libbeat/publisher"
	"github.com/elastic/beats/v7/libbeat/publisher/queue/memqueue"
	conf "github.com/elastic/elastic-agent-libs/config"

	//"github.com/elastic/beats/v7/libbeat/tests/resources"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOutputReload(t *testing.T) {
	tests := map[string]func(mockPublishFn) outputs.Client{
		"client":         newMockClient,
		"network_client": newMockNetworkClient,
	}

	for name, ctor := range tests {
		t.Run(name, func(t *testing.T) {
			testutil.SeedPRNG(t)

			// Flaky check: https://github.com/elastic/beats/issues/21656
			//goroutines := resources.NewGoroutinesChecker()
			//defer goroutines.Check(t)

			err := quick.Check(func(q uint) bool {
				numEventsToPublish := 15000 + (q % 500) // 15000 to 19999
				numOutputReloads := 350 + (q % 150)     // 350 to 499

				queueConfig := QueueConfig{
					Type: memqueue.QueueType,
					UserConfig: conf.MustNewConfigFrom(
						map[string]interface{}{"events": numEventsToPublish},
					),
				}

				var publishedCount atomic.Uint
				countingPublishFn := func(batch publisher.Batch) error {
					publishedCount.Add(uint(len(batch.Events())))
					return nil
				}

				pipeline, err := New(
					beat.Info{},
					Monitors{},
					queueConfig,
					outputs.Group{},
					Settings{},
				)
				require.NoError(t, err)
				defer pipeline.Close()

				pipelineClient, err := pipeline.Connect()
				require.NoError(t, err)
				defer pipelineClient.Close()

				var wg sync.WaitGroup
				wg.Add(1)
				go func() {
					for i := uint(0); i < numEventsToPublish; i++ {
						pipelineClient.Publish(beat.Event{})
					}
					wg.Done()
				}()

				for i := uint(0); i < numOutputReloads; i++ {
					outputClient := ctor(countingPublishFn)
					out := outputs.Group{
						Clients: []outputs.Client{outputClient},
					}
					pipeline.outputController.Set(out)
				}

				wg.Wait()

				timeout := 20 * time.Second
				return waitUntilTrue(timeout, func() bool {
					return numEventsToPublish == publishedCount.Load()
				})
			}, &quick.Config{MaxCount: 25})

			if err != nil {
				t.Error(err)
			}
		})
	}
}

func TestSetEmptyOutputsSendsNilChannel(t *testing.T) {
	// Just fill out enough to confirm what's sent to the event consumer,
	// we don't want to start up real helper routines.
	controller := outputController{
		consumer: &eventConsumer{
			targetChan: make(chan consumerTarget, 2),
		},
	}
	controller.Set(outputs.Group{})

	// Two messages should be sent to eventConsumer's targetChan:
	// one to clear the old target while the state is updating,
	// and one with the new metadata after the state update is
	// complete. Since we're setting an empty output group, both
	// of these calls should have a nil target channel.
	target := <-controller.consumer.targetChan
	assert.Nil(t, target.ch, "consumerTarget should receive a nil channel to block batch assembly")
	target = <-controller.consumer.targetChan
	assert.Nil(t, target.ch, "consumerTarget should receive a nil channel to block batch assembly")
}
