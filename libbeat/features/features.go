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

package features

import (
	"fmt"
	"sync"

	"github.com/elastic/elastic-agent-client/v7/pkg/proto"
	conf "github.com/elastic/elastic-agent-libs/config"
)

var (
	flags = fflags{}
)

type fflags struct {
	mu sync.RWMutex

	fqdnEnabled bool
}

// UpdateFromProto updates the feature flags configuration. If f is nil UpdateFromProto is no-op.
func UpdateFromProto(f *proto.Features) {
	if f == nil {
		return
	}

	flags.mu.Lock()
	defer flags.mu.Unlock()

	if f.Fqdn == nil {
		// By default, FQDN reporting is disabled.
		flags.fqdnEnabled = false
		return
	}

	flags.fqdnEnabled = f.Fqdn.Enabled
}

// UpdateFromConfig updates the feature flags configuration. If c is nil UpdateFromConfig is no-op.
func UpdateFromConfig(c *conf.C) error {
	if c == nil {
		return nil
	}

	type cfg struct {
		Features struct {
			FQDN *conf.C `json:"fqdn" yaml:"fqdn" config:"fqdn"`
		} `json:"features" yaml:"features" config:"features"`
	}

	parsedFlags := cfg{}
	if err := c.Unpack(&parsedFlags); err != nil {
		return fmt.Errorf("could not Unpack features config: %w", err)
	}

	flags.mu.Lock()
	defer flags.mu.Unlock()
	flags.fqdnEnabled = parsedFlags.Features.FQDN.Enabled()

	return nil
}

// FQDN reports if FQDN should be used instead of hostname for host.name.
// If it hasn't been set by UpdateFromConfig or UpdateFromProto, it returns false.
func FQDN() bool {
	flags.mu.RLock()
	defer flags.mu.RUnlock()
	return flags.fqdnEnabled
}
