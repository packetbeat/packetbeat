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

package filter

import (
	"io"
	"testing"

	"github.com/elastic/beats/v7/libbeat/reader"
	"github.com/elastic/elastic-agent-libs/config"
	"github.com/stretchr/testify/require"
)

func TestParser(t *testing.T) {
	tests := map[string]struct {
		config                 map[string]interface{}
		input                  []reader.Message
		expectedMessageContent [][]byte
	}{
		"keep all messages": {
			config: map[string]interface{}{
				"pattern": "this matches*",
			},
			input: []reader.Message{
				{
					Content: []byte("this matches"),
				},
				{
					Content: []byte("this matches again"),
				},
			},
			expectedMessageContent: [][]byte{
				[]byte("this matches"),
				[]byte("this matches again"),
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var c Config
			cfg := config.MustNewConfigFrom(test.config)
			err := cfg.Unpack(&c)
			r := NewFilterParser(newTestReader(test.input), &c)

			contents := make([][]byte, 0)
			msg, err := r.Next()
			for err == nil {
				contents = append(contents, msg.Content)
				msg, err = r.Next()
			}
			require.ElementsMatch(t, test.expectedMessageContent, contents)
		})

	}
}

type testReader struct {
	msg []reader.Message
	idx int
}

func newTestReader(input []reader.Message) reader.Reader {
	return &testReader{
		msg: input,
		idx: 0,
	}
}

func (r *testReader) Next() (reader.Message, error) {
	if r.idx == len(r.msg) {
		return reader.Message{}, io.EOF
	}

	m := r.msg[r.idx]
	r.idx += 1
	return m, nil
}

func (r *testReader) Close() error { return nil }
