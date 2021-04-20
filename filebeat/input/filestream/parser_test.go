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

package filestream

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/reader"
	"github.com/elastic/beats/v7/libbeat/reader/multiline"
	"github.com/elastic/beats/v7/libbeat/reader/readfile"
	"github.com/elastic/beats/v7/libbeat/reader/readfile/encoding"
)

func TestParsersConfigAndReading(t *testing.T) {
	tests := map[string]struct {
		lines            string
		parsers          map[string]interface{}
		expectedMessages []string
		expectedError    string
	}{
		"no parser, no error": {
			lines: "line 1\nline 2\n",
			parsers: map[string]interface{}{
				"paths": []string{"dummy_path"},
			},
			expectedMessages: []string{"line 1\n", "line 2\n"},
		},
		"correct strip_newline parser": {
			lines: "line 1\nline 2\n",
			parsers: map[string]interface{}{
				"paths": []string{"dummy_path"},
				"parsers": []map[string]interface{}{
					map[string]interface{}{
						"strip_newline": nil,
					},
				},
			},
			expectedMessages: []string{"line 1", "line 2"},
		},
		"non existent parser configuration": {
			parsers: map[string]interface{}{
				"paths": []string{"dummy_path"},
				"parsers": []map[string]interface{}{
					map[string]interface{}{
						"no_such_parser": nil,
					},
				},
			},
			expectedError: ErrNoSuchParser.Error(),
		},
		"invalid multiline parser configuration is caught before parser creation": {
			parsers: map[string]interface{}{
				"paths": []string{"dummy_path"},
				"parsers": []map[string]interface{}{
					map[string]interface{}{
						"multiline": map[string]interface{}{
							"match": "after",
						},
					},
				},
			},
			expectedError: multiline.ErrMissingPattern.Error(),
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			cfg := defaultConfig()
			parsersConfig := common.MustNewConfigFrom(test.parsers)
			err := parsersConfig.Unpack(&cfg)
			if test.expectedError == "" {
				require.NoError(t, err)
			} else {
				require.Contains(t, err.Error(), test.expectedError)
				return
			}

			p, err := newParsers(testReader(test.lines), parserConfig{lineTerminator: readfile.AutoLineTerminator}, cfg.Reader.Parsers)

			i := 0
			msg, err := p.Next()
			for err == nil {
				require.Contains(t, test.expectedMessages[i], string(msg.Content))
				i++
				msg, err = p.Next()
			}
		})
	}
}

func TestPostProcessor(t *testing.T) {
	tests := map[string]struct {
		message         reader.Message
		postProcessors  map[string]interface{}
		expectedMessage reader.Message
	}{
		"no postprocesser, no processing": {
			message: reader.Message{
				Content: []byte("line 1"),
			},
			postProcessors: map[string]interface{}{
				"paths": []string{"dummy_path"},
			},
			expectedMessage: reader.Message{
				Content: []byte("line 1"),
			},
		},
		"JSON post processer with keys_under_root": {
			message: reader.Message{
				Fields: common.MapStr{
					"json": common.MapStr{
						"key": "value",
					},
				},
			},
			postProcessors: map[string]interface{}{
				"paths": []string{"dummy_path"},
				"parsers": []map[string]interface{}{
					map[string]interface{}{
						"ndjson": map[string]interface{}{
							"keys_under_root": true,
						},
					},
				},
			},
			expectedMessage: reader.Message{
				Fields: common.MapStr{
					"key": "value",
				},
			},
		},
		"JSON post processer with document ID": {
			message: reader.Message{
				Fields: common.MapStr{
					"json": common.MapStr{
						"key":         "value",
						"my-id-field": "my-id",
					},
				},
			},
			postProcessors: map[string]interface{}{
				"paths": []string{"dummy_path"},
				"parsers": []map[string]interface{}{
					map[string]interface{}{
						"ndjson": map[string]interface{}{
							"keys_under_root": true,
							"document_id":     "my-id-field",
						},
					},
				},
			},
			expectedMessage: reader.Message{
				Fields: common.MapStr{
					"key": "value",
				},
				Meta: common.MapStr{
					"_id": "my-id",
				},
			},
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			cfg := defaultConfig()
			common.MustNewConfigFrom(test.postProcessors).Unpack(&cfg)
			pp := newPostProcessors(cfg.Reader.Parsers)

			msg := test.message
			for _, p := range pp {
				p.PostProcess(&msg)
			}
			require.Equal(t, test.expectedMessage, msg)
		})
	}

}

func testReader(lines string) reader.Reader {
	encF, _ := encoding.FindEncoding("")
	reader := strings.NewReader(lines)
	enc, err := encF(reader)
	if err != nil {
		panic(err)
	}
	r, err := readfile.NewEncodeReader(ioutil.NopCloser(reader), readfile.Config{
		Codec:      enc,
		BufferSize: 64,
		Terminator: readfile.AutoLineTerminator,
		MaxBytes:   128,
	})
	if err != nil {
		panic(err)
	}

	return r
}
