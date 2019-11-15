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
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBase64UUIDLen(t *testing.T) {
	id := GetBase64UUID()

	// Check that decoded ID is 15 bytes long
	decodedBytes, err := base64.RawURLEncoding.DecodeString(id)
	assert.NoError(t, err)
	assert.Len(t, decodedBytes, 15)
}

func TestGetBase64UUIDBytes(t *testing.T) {
	id := GetBase64UUID()

	// Check that bytes 7-12 are secure munged mac address
	decodedBytes, err := base64.RawURLEncoding.DecodeString(id)
	assert.NoError(t, err)
	assert.Equal(t, mac, decodedBytes[6:6+addrLen])
}

func TestGetBase64UUIDConsecutiveOrdering(t *testing.T) {
	prevID := GetBase64UUID()
	for i := 0; i < 10000; i++ {
		decodedPrevID, err := base64.RawURLEncoding.DecodeString(prevID)
		assert.NoError(t, err)

		currID := GetBase64UUID()
		decodedCurrID, err := base64.RawURLEncoding.DecodeString(currID)
		assert.NoError(t, err)

		// Check if current ID is greater than previous ID (accounting for
		// wrap around of first byte).
		if decodedCurrID[0] == 0x00 { // first byte wrapped around
			// Check that previous ID's first byte was max possible byte value (0xff)
			assert.EqualValues(t, decodedPrevID[0], 0xff)

			// Check that rest of current ID (after first byte) is greater than rest of
			// previous ID (after first byte)
			assert.True(t, isGreaterThan(decodedCurrID[1:], decodedPrevID[1:]))
		} else {
			// Check that current ID's first byte is exactly 1 more than previous ID's
			// first byte
			assert.Equal(t, decodedPrevID[0]+1, decodedCurrID[0])

			// Check that entire current ID is greater than entire previous ID
			assert.True(t, isGreaterThan(decodedCurrID, decodedPrevID))
		}

		prevID = currID
	}
}

func BenchmarkGetBase64UUID(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetBase64UUID()
	}
}

func isGreaterThan(b1, b2 []byte) bool {
	if len(b1) > len(b2) {
		return true
	}

	if len(b2) < len(b1) {
		return false
	}

	if len(b1) == 0 {
		return false
	}

	// Lengths are equal and at least 1, compare values

	if b1[0] < b2[0] {
		return false
	}

	if b1[0] > b2[0] {
		return true
	}

	return isGreaterThan(b1[1:], b2[1:])
}
