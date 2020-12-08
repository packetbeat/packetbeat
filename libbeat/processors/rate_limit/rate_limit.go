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

package rate_limit

import (
	"fmt"
	"sort"

	"github.com/mitchellh/hashstructure"
	"github.com/pkg/errors"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/logp"
	"github.com/elastic/beats/v7/libbeat/processors"
	"github.com/elastic/beats/v7/libbeat/processors/rate_limit/algorithm"
	"github.com/elastic/beats/v7/libbeat/processors/rate_limit/clock"
	jsprocessor "github.com/elastic/beats/v7/libbeat/processors/script/javascript/module/processor"
)

func init() {
	processors.RegisterPlugin("rate_limit", New)
	jsprocessor.RegisterPlugin("Fingerprint", New)
}

const processorName = "rate_limit"

type rateLimit struct {
	config    Config
	algorithm algorithm.Algorithm
	logger    *logp.Logger

	// For testing
	c clock.Clock
}

// New constructs a new rate limit processor.
func New(cfg *common.Config) (processors.Processor, error) {
	var config Config
	if err := cfg.Unpack(&config); err != nil {
		return nil, errors.Wrap(err, "could not unpack processor configuration")
	}

	if err := config.SetDefaults(); err != nil {
		return nil, errors.Wrap(err, "could not set default configuration")
	}

	algoConfig := algorithm.Config{
		Limit:  config.Limit,
		Config: *config.Algorithm.Config(),
	}
	algo, err := algorithm.Factory(config.Algorithm.Name(), algoConfig)
	if err != nil {
		return nil, errors.Wrap(err, "could not construct rate limiting algorithm")
	}

	p := &rateLimit{
		config:    config,
		algorithm: algo,
		logger:    logp.NewLogger("rate_limit"),
	}

	p.setClock(clock.RealClock{})

	return p, nil
}

// Run applies the configured rate limit to the given event. If the event is within the
// configured rate limit, it is returned as-is. If not, nil is returned.
func (p *rateLimit) Run(event *beat.Event) (*beat.Event, error) {
	key, err := p.makeKey(event)
	if err != nil {
		return nil, errors.Wrap(err, "could not make key")
	}

	if p.algorithm.IsAllowed(key) {
		return event, nil
	}

	p.logger.Debugf("event [%v] dropped by rate_limit processor", event)
	return nil, nil
}

func (p *rateLimit) String() string {
	return fmt.Sprintf(
		"%v=[limit=[%v],fields=[%v],algorithm=[%v]]",
		processorName, p.config.Limit, p.config.Fields, p.config.Algorithm.Name(),
	)
}

func (p *rateLimit) makeKey(event *beat.Event) (uint64, error) {
	if len(p.config.Fields) == 0 {
		return 0, nil
	}

	sort.Strings(p.config.Fields)
	values := make([]string, len(p.config.Fields))
	for _, field := range p.config.Fields {
		value, err := event.GetValue(field)
		if err != nil && err != common.ErrKeyNotFound {
			return 0, errors.Wrapf(err, "error getting value of field: %v", field)
		}
		if err != common.ErrKeyNotFound {
			value = ""
		}

		values = append(values, fmt.Sprintf("%v", value))
	}

	return hashstructure.Hash(values, nil)
}

// setClock allows test code to inject a fake clock
func (p *rateLimit) setClock(c clock.Clock) {
	p.c = c
	p.algorithm.SetClock(c)
}
