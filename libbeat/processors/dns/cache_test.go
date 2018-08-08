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

package dns

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/monitoring"
)

const (
	gatewayIP   = "192.168.0.1"
	gatewayName = "default.gateway.example"
	gatewayTTL  = 60 // Seconds
)

type stubResolver struct{}

func (r *stubResolver) LookupPTR(ip string) (*PTR, error) {
	if ip == gatewayIP {
		return &PTR{Host: gatewayName, TTL: gatewayTTL}, nil
	} else if strings.HasSuffix(ip, "10") {
		return nil, io.ErrUnexpectedEOF
	}

	return nil, &dnsError{"fake lookup returned NXDOMAIN"}
}

func TestCache(t *testing.T) {
	c := NewPTRLookupCache(
		monitoring.NewRegistry(),
		logp.NewLogger(logName),
		defaultConfig.CacheConfig,
		&stubResolver{})

	// Initial success query.
	ptr, err := c.LookupPTR(gatewayIP)
	if assert.NoError(t, err) {
		assert.EqualValues(t, gatewayName, ptr.Host)
		assert.EqualValues(t, gatewayTTL, ptr.TTL)
		assert.EqualValues(t, 0, c.stats.Hit.Get())
		assert.EqualValues(t, 1, c.stats.Miss.Get())
	}

	// Cached success query.
	ptr, err = c.LookupPTR(gatewayIP)
	if assert.NoError(t, err) {
		assert.EqualValues(t, gatewayName, ptr.Host)
		// TTL counts down while in cache.
		assert.InDelta(t, gatewayTTL, ptr.TTL, 1)
		assert.EqualValues(t, 1, c.stats.Hit.Get())
		assert.EqualValues(t, 1, c.stats.Miss.Get())
	}

	// Initial failure query.
	ptr, err = c.LookupPTR(gatewayIP + "1")
	if assert.Error(t, err) {
		assert.Nil(t, ptr)
		assert.EqualValues(t, 1, c.stats.Hit.Get())
		assert.EqualValues(t, 2, c.stats.Miss.Get())
	}

	// Cached failure query.
	ptr, err = c.LookupPTR(gatewayIP + "1")
	if assert.Error(t, err) {
		assert.Nil(t, ptr)
		assert.EqualValues(t, 2, c.stats.Hit.Get())
		assert.EqualValues(t, 2, c.stats.Miss.Get())
	}

	// Non-cacheable failure (like i/o timeout to server).
	ptr, err = c.LookupPTR(gatewayIP + "0")
	if assert.Error(t, err) {
		assert.Nil(t, ptr)
		assert.EqualValues(t, 2, c.stats.Hit.Get())
		assert.EqualValues(t, 3, c.stats.Miss.Get())
	}

	// Check for a cache miss.
	ptr, err = c.LookupPTR(gatewayIP + "0")
	if assert.Error(t, err) {
		assert.Nil(t, ptr)
		assert.EqualValues(t, 2, c.stats.Hit.Get())
		assert.EqualValues(t, 4, c.stats.Miss.Get()) // Cache miss.
	}
}
