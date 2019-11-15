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

package elasticsearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidAddress(t *testing.T) {
	tests := map[string]struct {
		addr     []byte
		expected bool
	}{
		"nil": {
			nil,
			false,
		},
		"too_short": {
			[]byte{0xde, 0xad, 0xbe, 0xef},
			false,
		},
		"too_long": {
			[]byte{0xbe, 0xa7, 0x5a, 0x43, 0xda, 0xbe, 0x57},
			false,
		},
		"all_zeros": {
			[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			false,
		},
		"good": {
			[]byte{0xbe, 0xa7, 0x5a, 0x43, 0x90, 0x0d},
			true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			v := isValidAddress(test.addr)
			assert.Equal(t, test.expected, v)
		})
	}
}

// TODO: test constructDummyMulticastAddress (length, multicast bit)
// TODO: test getSecureMungedMACAddress (length)
// TODO: getMacAddress (not loopback, length)
