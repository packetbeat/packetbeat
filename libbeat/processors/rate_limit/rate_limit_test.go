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
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
)

func TestNew(t *testing.T) {
	cases := map[string]struct {
		config common.MapStr
		err    string
	}{
		"default": {
			common.MapStr{},
			"",
		},
		"unknown_algo": {
			common.MapStr{
				"algorithm": common.MapStr{
					"foobar": common.MapStr{},
				},
			},
			"rate limiting algorithm 'foobar' not implemented",
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			config := common.MustNewConfigFrom(test.config)
			_, err := New(config)
			if test.err == "" {
				require.NoError(t, err)
			} else {
				require.Error(t, err, test.err)
			}
		})
	}
}

func TestRateLimit(t *testing.T) {
	inEvents := []beat.Event{
		{
			Timestamp: time.Now(),
			Fields: common.MapStr{
				"foo": "bar",
			},
		},
		{
			Timestamp: time.Now(),
			Fields: common.MapStr{
				"foo": "bar",
				"baz": "mosquito",
			},
		},
		{
			Timestamp: time.Now(),
			Fields: common.MapStr{
				"baz": "qux",
			},
		},
		{
			Timestamp: time.Now(),
			Fields: common.MapStr{
				"foo": "seger",
			},
		},
	}

	cases := map[string]struct {
		config    common.MapStr
		inEvents  []beat.Event
		delay     time.Duration
		outEvents []beat.Event
	}{
		"rate_0": {
			config:    common.MapStr{},
			inEvents:  inEvents,
			outEvents: []beat.Event{},
		},
		"rate_1_per_min": {
			config: common.MapStr{
				"limit": "1/m",
			},
			inEvents:  inEvents,
			outEvents: inEvents[0:1],
		},
		"rate_2_per_min": {
			config: common.MapStr{
				"limit": "2/m",
			},
			inEvents:  inEvents,
			outEvents: inEvents[0:2],
		},
		"rate_5_per_min": {
			config: common.MapStr{
				"limit": "5/m",
			},
			inEvents:  inEvents,
			outEvents: inEvents,
		},
		"rate_2_per_sec": {
			config: common.MapStr{
				"limit": "2/s",
			},
			delay:     200 * time.Millisecond,
			inEvents:  inEvents,
			outEvents: []beat.Event{inEvents[0], inEvents[1], inEvents[3]},
		},
		"with_fields": {
			config: common.MapStr{
				"limit":  "1/s",
				"fields": []string{"foo"},
			},
			delay:     400 * time.Millisecond,
			inEvents:  inEvents,
			outEvents: []beat.Event{inEvents[0], inEvents[2], inEvents[3]},
		},
		"with_burst": {
			config: common.MapStr{
				"limit":            "2/s",
				"burst_multiplier": 2,
			},
			delay:     400 * time.Millisecond,
			inEvents:  inEvents,
			outEvents: inEvents,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			p, err := New(common.MustNewConfigFrom(test.config))
			require.NoError(t, err)

			out := make([]beat.Event, 0)
			for _, in := range test.inEvents {
				o, err := p.Run(&in)
				require.NoError(t, err)
				if o != nil {
					out = append(out, *o)
				}
				time.Sleep(test.delay)
			}

			require.Equal(t, test.outEvents, out)
		})
	}
}
