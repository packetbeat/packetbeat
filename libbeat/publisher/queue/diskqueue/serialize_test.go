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

package diskqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/publisher"
	"github.com/elastic/elastic-agent-libs/mapstr"
	"github.com/elastic/elastic-agent-shipper-client/pkg/proto/messages"
)

// A test to make sure serialization works correctly on multi-byte characters.
func TestSerialize(t *testing.T) {
	testCases := []struct {
		name   string
		value  string
		format SerializationFormat
	}{
		{
			name:   "Ascii only, CBOR",
			value:  "{\"name\": \"Momotaro\"}",
			format: SerializationCBOR,
		},
		{
			name:   "Multi-byte, CBOR",
			value:  "{\"name\": \"桃太郎\"}",
			format: SerializationCBOR,
		},
		{
			name:   "Ascii only, Protobuf",
			value:  "{\"name\": \"Momotaro\"}",
			format: SerializationProtobuf,
		},
		{
			name:   "Multi-byte, Protobuf",
			value:  "{\"name\": \"桃太郎\"}",
			format: SerializationProtobuf,
		},
	}

	for _, test := range testCases {
		encoder := newEventEncoder(test.format)
		var event interface{}
		switch test.format {
		case SerializationCBOR:
			event = publisher.Event{
				Content: beat.Event{
					Fields: mapstr.M{
						"test_field": test.value,
					},
				},
			}
		case SerializationProtobuf:
			event = &messages.Event{
				Fields: &messages.Struct{
					Data: map[string]*messages.Value{
						"test_field": &messages.Value{
							Kind: &messages.Value_StringValue{
								StringValue: test.value,
							},
						},
					},
				},
			}
		}
		serialized, err := encoder.encode(event)
		if err != nil {
			t.Fatalf("[%v] Couldn't encode event: %v", test.name, err)
		}

		// Use decoder to decode the serialized bytes.
		decoder := newEventDecoder()
		decoder.serializationFormat = test.format
		buf := decoder.Buffer(len(serialized))
		copy(buf, serialized)
		decoded, err := decoder.Decode()
		if err != nil {
			t.Fatalf("[%v] Couldn't decode serialized data: %v", test.name, err)
		}

		switch test.format {
		case SerializationCBOR:
			event, ok := decoded.(publisher.Event)
			require.True(t, ok)
			decodedValue, err := event.Content.Fields.GetValue("test_field")
			if err != nil {
				t.Fatalf("[%v] Couldn't get field 'test_field': %v", test.name, err)
			}
			assert.Equal(t, test.value, decodedValue)
		case SerializationProtobuf:
			event, ok := decoded.(*messages.Event)
			require.True(t, ok)
			d := event.GetFields().GetData()
			decodedValue := d["test_field"].GetStringValue()
			assert.Equal(t, test.value, decodedValue)
		}
	}
}
