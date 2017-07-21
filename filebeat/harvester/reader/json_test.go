package reader

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/beats/libbeat/common"
)

func TestUnmarshal(t *testing.T) {
	tests := []struct {
		Name   string
		Input  string
		Output map[string]interface{}
	}{
		{
			Name:  "Top level int, float, string, bool",
			Input: `{"a": 3, "b": 2.0, "c": "hello", "d": true}`,
			Output: map[string]interface{}{
				"a": int64(3),
				"b": float64(2),
				"c": "hello",
				"d": true,
			},
		},
		{
			Name:  "Nested objects with ints",
			Input: `{"a": 3, "b": {"c": {"d": 5}}}`,
			Output: map[string]interface{}{
				"a": int64(3),
				"b": map[string]interface{}{
					"c": map[string]interface{}{
						"d": int64(5),
					},
				},
			},
		},
		{
			Name:  "Array of floats",
			Input: `{"a": 3, "b": {"c": [4.0, 4.1, 4.2]}}`,
			Output: map[string]interface{}{
				"a": int64(3),
				"b": map[string]interface{}{
					"c": []interface{}{
						float64(4.0), float64(4.1), float64(4.2),
					},
				},
			},
		},
		{
			Name:  "Array of mixed ints and floats",
			Input: `{"a": 3, "b": {"c": [4, 4.1, 4.2]}}`,
			Output: map[string]interface{}{
				"a": int64(3),
				"b": map[string]interface{}{
					"c": []interface{}{
						int64(4), float64(4.1), float64(4.2),
					},
				},
			},
		},
		{
			Name:  "Negative values",
			Input: `{"a": -3, "b": -1.0}`,
			Output: map[string]interface{}{
				"a": int64(-3),
				"b": float64(-1),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var output map[string]interface{}
			err := unmarshal([]byte(test.Input), &output)
			assert.NoError(t, err)
			assert.Equal(t, test.Output, output)
		})

	}
}

func TestDecodeJSON(t *testing.T) {
	var tests = []struct {
		Text         string
		Config       JSONConfig
		ExpectedText string
		ExpectedMap  common.MapStr
	}{
		{
			Text:         `{"message": "test", "value": 1}`,
			Config:       JSONConfig{MessageKey: "message"},
			ExpectedText: "test",
			ExpectedMap:  common.MapStr{"message": "test", "value": int64(1)},
		},
		{
			Text:         `{"message": "test", "value": 1}`,
			Config:       JSONConfig{MessageKey: "message1"},
			ExpectedText: "",
			ExpectedMap:  common.MapStr{"message": "test", "value": int64(1)},
		},
		{
			Text:         `{"message": "test", "value": 1}`,
			Config:       JSONConfig{MessageKey: "value"},
			ExpectedText: "",
			ExpectedMap:  common.MapStr{"message": "test", "value": int64(1)},
		},
		{
			Text:         `{"message": "test", "value": "1"}`,
			Config:       JSONConfig{MessageKey: "value"},
			ExpectedText: "1",
			ExpectedMap:  common.MapStr{"message": "test", "value": "1"},
		},
		{
			// in case of JSON decoding errors, the text is passed as is
			Text:         `{"message": "test", "value": "`,
			Config:       JSONConfig{MessageKey: "value"},
			ExpectedText: `{"message": "test", "value": "`,
			ExpectedMap:  nil,
		},
		{
			// in case the JSON is "null", we should just not panic
			Text:         `null`,
			Config:       JSONConfig{MessageKey: "value", AddErrorKey: true},
			ExpectedText: `null`,
			ExpectedMap:  common.MapStr{"error": common.MapStr{"message": "Error decoding JSON: <nil>", "type": "json"}},
		},
		{
			// Add key error helps debugging this
			Text:         `{"message": "test", "value": "`,
			Config:       JSONConfig{MessageKey: "value", AddErrorKey: true},
			ExpectedText: `{"message": "test", "value": "`,
			ExpectedMap:  common.MapStr{"error": common.MapStr{"message": "Error decoding JSON: unexpected EOF", "type": "json"}},
		},
		{
			// If the text key is not found, put an error
			Text:         `{"message": "test", "value": "1"}`,
			Config:       JSONConfig{MessageKey: "hello", AddErrorKey: true},
			ExpectedText: ``,
			ExpectedMap:  common.MapStr{"message": "test", "value": "1", "error": common.MapStr{"message": "Key 'hello' not found", "type": "json"}},
		},
		{
			// If the text key is found, but not a string, put an error
			Text:         `{"message": "test", "value": 1}`,
			Config:       JSONConfig{MessageKey: "value", AddErrorKey: true},
			ExpectedText: ``,
			ExpectedMap:  common.MapStr{"message": "test", "value": int64(1), "error": common.MapStr{"message": "Value of key 'value' is not a string", "type": "json"}},
		},
		{
			// Without a text key, simple return the json and an empty text
			Text:         `{"message": "test", "value": 1}`,
			Config:       JSONConfig{AddErrorKey: true},
			ExpectedText: ``,
			ExpectedMap:  common.MapStr{"message": "test", "value": int64(1)},
		},
	}

	for _, test := range tests {

		var p JSON
		p.cfg = &test.Config
		text, M := p.decodeJSON([]byte(test.Text))
		assert.Equal(t, test.ExpectedText, string(text))
		assert.Equal(t, test.ExpectedMap, M)
	}
}

func TestAddJSONFields(t *testing.T) {
	type io struct {
	}

	text := "hello"

	now := time.Now().UTC()

	tests := []struct {
		Name          string
		Data          common.MapStr
		Text          *string
		JSONConfig    JSONConfig
		ExpectedItems common.MapStr
	}{
		{
			// by default, don't overwrite keys
			Name:       "default: do not overwrite",
			Data:       common.MapStr{"type": "test_type", "json": common.MapStr{"type": "test", "text": "hello"}},
			Text:       &text,
			JSONConfig: JSONConfig{KeysUnderRoot: true},
			ExpectedItems: common.MapStr{
				"type": "test_type",
				"text": "hello",
			},
		},
		{
			// overwrite keys if asked
			Name:       "overwrite keys if configured",
			Data:       common.MapStr{"type": "test_type", "json": common.MapStr{"type": "test", "text": "hello"}},
			Text:       &text,
			JSONConfig: JSONConfig{KeysUnderRoot: true, OverwriteKeys: true},
			ExpectedItems: common.MapStr{
				"type": "test",
				"text": "hello",
			},
		},
		{
			// without keys_under_root, put everything in a json key
			Name:       "use json namespace w/o keys_under_root",
			Data:       common.MapStr{"type": "test_type", "json": common.MapStr{"type": "test", "text": "hello"}},
			Text:       &text,
			JSONConfig: JSONConfig{},
			ExpectedItems: common.MapStr{
				"json": common.MapStr{"type": "test", "text": "hello"},
			},
		},
		{
			// when MessageKey is defined, the Text overwrites the value of that key
			Name:       "write result to message_key field",
			Data:       common.MapStr{"type": "test_type", "json": common.MapStr{"type": "test", "text": "hi"}},
			Text:       &text,
			JSONConfig: JSONConfig{MessageKey: "text"},
			ExpectedItems: common.MapStr{
				"json": common.MapStr{"type": "test", "text": "hello"},
				"type": "test_type",
			},
		},
		{
			// when @timestamp is in JSON and overwrite_keys is true, parse it
			// in a common.Time
			Name:       "parse @timestamp",
			Data:       common.MapStr{"@timestamp": now, "type": "test_type", "json": common.MapStr{"type": "test", "@timestamp": "2016-04-05T18:47:18.444Z"}},
			Text:       &text,
			JSONConfig: JSONConfig{KeysUnderRoot: true, OverwriteKeys: true},
			ExpectedItems: common.MapStr{
				"@timestamp": common.MustParseTime("2016-04-05T18:47:18.444Z"),
				"type":       "test",
			},
		},
		{
			// when the parsing on @timestamp fails, leave the existing value and add an error key
			// in a common.Time
			Name:       "fail to parse @timestamp",
			Data:       common.MapStr{"@timestamp": common.Time(now), "type": "test_type", "json": common.MapStr{"type": "test", "@timestamp": "2016-04-05T18:47:18.44XX4Z"}},
			Text:       &text,
			JSONConfig: JSONConfig{KeysUnderRoot: true, OverwriteKeys: true},
			ExpectedItems: common.MapStr{
				"@timestamp": common.Time(now),
				"type":       "test",
				"error":      common.MapStr{"type": "json", "message": "@timestamp not overwritten (parse error on 2016-04-05T18:47:18.44XX4Z)"},
			},
		},
		{
			// when the @timestamp has the wrong type, leave the existing value and add an error key
			// in a common.Time
			Name:       "wrong @timestamp format",
			Data:       common.MapStr{"@timestamp": common.Time(now), "type": "test_type", "json": common.MapStr{"type": "test", "@timestamp": 42}},
			Text:       &text,
			JSONConfig: JSONConfig{KeysUnderRoot: true, OverwriteKeys: true},
			ExpectedItems: common.MapStr{
				"@timestamp": common.Time(now),
				"type":       "test",
				"error":      common.MapStr{"type": "json", "message": "@timestamp not overwritten (not string)"},
			},
		},
		{
			// if overwrite_keys is true, but the `type` key in json is not a string, ignore it
			Name:       "ignore non-string type field",
			Data:       common.MapStr{"type": "test_type", "json": common.MapStr{"type": 42}},
			Text:       &text,
			JSONConfig: JSONConfig{KeysUnderRoot: true, OverwriteKeys: true},
			ExpectedItems: common.MapStr{
				"type":  "test_type",
				"error": common.MapStr{"type": "json", "message": "type not overwritten (not string)"},
			},
		},
		{
			// if overwrite_keys is true, but the `type` key in json is empty, ignore it
			Name:       "ignore empty type field",
			Data:       common.MapStr{"type": "test_type", "json": common.MapStr{"type": ""}},
			Text:       &text,
			JSONConfig: JSONConfig{KeysUnderRoot: true, OverwriteKeys: true},
			ExpectedItems: common.MapStr{
				"type":  "test_type",
				"error": common.MapStr{"type": "json", "message": "type not overwritten (invalid value [])"},
			},
		},
		{
			// if overwrite_keys is true, but the `type` key in json starts with _, ignore it
			Name:       "ignore type names starting with underscore",
			Data:       common.MapStr{"@timestamp": common.Time(now), "type": "test_type", "json": common.MapStr{"type": "_type"}},
			Text:       &text,
			JSONConfig: JSONConfig{KeysUnderRoot: true, OverwriteKeys: true},
			ExpectedItems: common.MapStr{
				"type":  "test_type",
				"error": common.MapStr{"type": "json", "message": "type not overwritten (invalid value [_type])"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var jsonFields common.MapStr
			if fields, ok := test.Data["json"]; ok {
				jsonFields = fields.(common.MapStr)
			}

			MergeJSONFields(test.Data, jsonFields, test.Text, test.JSONConfig)

			t.Log("Executing test:", test)
			for k, v := range test.ExpectedItems {
				assert.Equal(t, v, test.Data[k])
			}
		})
	}
}
