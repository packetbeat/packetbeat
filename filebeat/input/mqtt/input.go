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

package mqtt

import (
	"context"
	"sync"
	"time"

	"github.com/elastic/beats/filebeat/channel"
	"github.com/elastic/beats/filebeat/input"
	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/common/backoff"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/pkg/errors"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func init() {
	err := input.Register("mqtt", NewInput)
	if err != nil {
		panic(err)
	}
}

// Input contains the input and its config
type mqttInput struct {
	config        mqttInputConfig
	context       input.Context
	outlet        channel.Outleter
	log           *logp.Logger
	mqttWaitGroup sync.WaitGroup
	runOnce       sync.Once
	client        MQTT.Client
}

// NewInput creates a new mqtt input
func NewInput(
	cfg *common.Config,
	connector channel.Connector,
	inputContext input.Context,
) (input.Input, error) {

	config := defaultConfig()
	if err := cfg.Unpack(&config); err != nil {
		return nil, errors.Wrap(err, "reading mqtt input config")
	}

	out, err := connector.ConnectWith(cfg, beat.ClientConfig{
		Processing: beat.ProcessingConfig{
			DynamicFields: inputContext.DynamicFields,
		},
		//		ACKEvents: func(events []interface{}) {
		//			for _, event := range events {
		//				if meta, ok := event.(eventMeta); ok {
		//					meta.handler.ack(meta.message)
		//				}
		//			}
		//		},
		CloseRef:  doneChannelContext(inputContext.Done),
		WaitClose: config.WaitClose,
	})
	if err != nil {
		return nil, err
	}

	input := &mqttInput{
		config:  config,
		context: inputContext,
		outlet:  out,
		log:     logp.NewLogger("mqtt input").With("host", config.Host),
	}

	input.setupMqttClient()

	return input, nil
}

// Run starts the input by scanning for incoming messages and errors.
func (input *mqttInput) Run() {
	input.runOnce.Do(func() {
		go func() {

			context := doneChannelContext(input.context.Done)

			// If the consumer fails to connect, we use exponential backoff with
			// jitter up to 8 * the initial backoff interval.
			backoff := backoff.NewEqualJitterBackoff(
				input.context.Done,
				input.config.ConnectBackoff,
				8*input.config.ConnectBackoff)

			for context.Err() == nil {
				err := input.connect()
				if err != nil {
					logp.Error(err)
					backoff.Wait()
				} else {
					break
				}
			}

			// We've successfully connected, reset the backoff timer.
			backoff.Reset()

			//All the rest is working asynchronously within the MQTT client
		}()
	})
}

// Stop doesn't need to do anything because the kafka consumer group and the
// input's outlet both have a context based on input.context.Done and will
// shut themselves down, since the done channel is already closed as part of
// the shutdown process in Runner.Stop().
func (input *mqttInput) Stop() {
}

// Wait should shut down the input and wait for it to complete, however (see
// Stop above) we don't need to take actions to shut down as long as the
// input.config.Done channel is closed, so we just make a (currently no-op)
// call to Stop() and then wait for sarama to signal completion.
func (input *mqttInput) Wait() {
	input.Stop()
}

// A barebones implementation of context.Context wrapped around the done
// channels that are more common in the beats codebase.
// TODO(faec): Generalize this to a common utility in a shared library
// (https://github.com/elastic/beats/issues/13125).
type channelCtx <-chan struct{}

func doneChannelContext(ch <-chan struct{}) context.Context {
	return channelCtx(ch)
}

func (c channelCtx) Deadline() (deadline time.Time, ok bool) { return }
func (c channelCtx) Done() <-chan struct{} {
	return (<-chan struct{})(c)
}
func (c channelCtx) Err() error {
	select {
	case <-c:
		return context.Canceled
	default:
		return nil
	}
}
func (c channelCtx) Value(key interface{}) interface{} { return nil }
