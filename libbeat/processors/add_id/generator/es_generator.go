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

package generator

import (
	"encoding/base64"
	"math/rand"
	"sync"
	"time"
)

type esTimeBasedUUIDGenerator struct{}

// Singleton instance
var _esTimeBasedUUIDGenerator IDGenerator = (*esTimeBasedUUIDGenerator)(nil)

// ESTimeBasedUUIDGenerator returns the singleton instance for this generator
func ESTimeBasedUUIDGenerator() IDGenerator {
	return _esTimeBasedUUIDGenerator
}

var (
	sequenceNumber uint32
	lastTimestamp  time.Time
	once           sync.Once
	mac            []byte
	mu             sync.Mutex
)

// NextID returns a base64-encoded, randomly-generated, but roughly ordered (over time), unique
// ID. The algorithm used to generate the ID is the same as used by Elasticsearch.
// See https://github.com/elastic/elasticsearch/blob/a666fb2266/server/src/main/java/org/elasticsearch/common/TimeBasedUUIDGenerator.java
func (*esTimeBasedUUIDGenerator) NextID() string {
	// Initialize sequence number and mac address. We do this here instead of doing it in a package-level
	// init function to give the runtime time to generate enough entropy for randomization.
	initOnce()

	ts, seq := nextIDData()
	var uuidBytes [15]byte

	packID(uuidBytes[:], ts, seq)
	return base64.RawURLEncoding.EncodeToString(uuidBytes[:])
}

func initOnce() {
	once.Do(func() {
		sequenceNumber = rand.Uint32()
		m, err := getSecureMungedMACAddress()
		if err != nil {
			panic(err)
		}
		mac = m
	})
}

func nextIDData() (int64, uint32) {
	mu.Lock()
	defer mu.Unlock()

	sequenceNumber++

	// We only use bottom 3 bytes for the sequence number.
	s := sequenceNumber & 0xffffff

	timestamp := getTimestamp()
	if s == 0 {
		// Always force the clock to increment whenever sequence number is 0, in case we have a long time-slip backwards.
		timestamp.Add(1 * time.Millisecond)
	}
	lastTimestamp = timestamp

	t := timestamp.UnixNano() / 1000 // timestamp in ms-since-epoch
	return t, s
}

func getTimestamp() time.Time {
	// Don't let timestamp go backwards, at least "on our watch" (while this process is running).  We are still vulnerable if we are
	// shut down, clock goes backwards, and we restart... for this we randomize the sequenceNumber on init to decrease chance of
	// collision.
	now := time.Now()

	if lastTimestamp.IsZero() {
		return now
	}

	if lastTimestamp.After(now) {
		return lastTimestamp
	}

	return now
}

func packID(buf []byte, ts int64, seq uint32) {
	//// We have auto-generated ids, which are usually used for append-only workloads.
	//// So we try to optimize the order of bytes for indexing speed (by having quite
	//// unique bytes close to the beginning of the ids so that sorting is fast) and
	//// compression (by making sure we share common prefixes between enough ids).

	// We use the sequence number rather than the timestamp because the distribution of
	// the timestamp depends too much on the indexing rate, so it is less reliable.
	buf[0] = byte(seq)       // copy lowest-order byte from sequence number
	buf[1] = byte(seq >> 16) // copy 3rd lowest-order byte from sequence number

	// Now we start focusing on compression and put bytes that should not change too often.
	buf[2] = byte(ts >> 16) // 3rd lowest-order byte from timestamp; changes every ~65 secs
	buf[3] = byte(ts >> 24) // 4th lowest-order byte from timestamp; changes every ~4.5h
	buf[4] = byte(ts >> 32) // 5th lowest-order byte from timestamp; changes every ~50 days
	buf[5] = byte(ts >> 40) // 6th lowest-order byte from timestamp; changes every 35 years

	// Copy mac address bytes (6 bytes)
	copy(buf[6:6+addrLen], mac)

	// Finally we put the remaining bytes, which will likely not be compressed at all.
	buf[12] = byte(ts >> 8)  // 2nd lowest-order byte from timestamp
	buf[13] = byte(seq >> 8) // 2nd lowest-order byte from sequence number
	buf[14] = byte(ts)

	// See also: more detailed explanation of byte choices at
	// https://github.com/elastic/elasticsearch/blob/a666fb22664284d8e2114841ebb58ea4e1924691/server/src/main/java/org/elasticsearch/common/TimeBasedUUIDGenerator.java#L80-L95
}
