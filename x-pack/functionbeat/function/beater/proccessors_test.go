// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package beater

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/beats/v7/libbeat/beat"
	_ "github.com/elastic/beats/v7/libbeat/processors/actions"
	conf "github.com/elastic/elastic-agent-libs/config"
	"github.com/elastic/elastic-agent-libs/mapstr"
)

func TestProcessorsForFunction(t *testing.T) {
	testCases := map[string]struct {
		beatInfo       beat.Info
		configStr      string
		event          beat.Event
		expectedFields map[string]string
	}{
		"Simple static index": {
			configStr: "index: 'test'",
			expectedFields: map[string]string{
				"@metadata.raw_index": "test",
			},
		},
		"Index with agent info + timestamp": {
			beatInfo:  beat.Info{Beat: "TestBeat", Version: "3.9.27"},
			configStr: "index: 'beat-%{[agent.name]}-%{[agent.version]}-%{+yyyy.MM.dd}'",
			event:     beat.Event{Timestamp: time.Date(1999, time.December, 31, 23, 0, 0, 0, time.UTC)},
			expectedFields: map[string]string{
				"@metadata.raw_index": "beat-TestBeat-3.9.27-1999.12.31",
			},
		},
		"Set field in input config": {
			configStr: `processors: [add_fields: {fields: {testField: inputConfig}}]`,
			expectedFields: map[string]string{
				"fields.testField": "inputConfig",
			},
		},
	}
	for description, test := range testCases {
		if test.event.Fields == nil {
			test.event.Fields = mapstr.M{}
		}
		config, err := functionConfigFromString(test.configStr)
		if err != nil {
			t.Errorf("[%s] %v", description, err)
			continue
		}
		processors, err := processorsForFunction(test.beatInfo, config)
		if err != nil {
			t.Errorf("[%s] %v", description, err)
			continue
		}
		processedEvent, err := processors.Run(&test.event)
		// We don't check if err != nil, because we are testing the final outcome
		// of running the processors, including when some of them fail.
		if processedEvent == nil {
			t.Errorf("[%s] Unexpected fatal error running processors: %v\n",
				description, err)
		}
		for key, value := range test.expectedFields {
			field, err := processedEvent.GetValue(key)
			if err != nil {
				t.Errorf("[%s] Couldn't get field %s from event: %v", description, key, err)
				continue
			}
			assert.Equal(t, field, value)
			fieldStr, ok := field.(string)
			if !ok {
				// Note that requiring a string here is just to simplify the test setup,
				// not a requirement of the underlying api.
				t.Errorf("[%s] Field [%s] should be a string", description, key)
				continue
			}
			if fieldStr != value {
				t.Errorf("[%s] Event field [%s]: expected [%s], got [%s]", description, key, value, fieldStr)
			}
		}
	}
}

func TestProcessorsForFunctionIsFlat(t *testing.T) {
	// This test is regrettable, and exists because of inconsistencies in
	// processor handling between processors.Processors and processing.group
	// (which implements beat.ProcessorList) -- see processorsForConfig for
	// details. The upshot is that, for now, if the function configuration specifies
	// processors, they must be returned as direct children of the resulting
	// processors.Processors (rather than being collected in additional tree
	// structure).
	// This test should be removed once we have a more consistent mechanism for
	// collecting and running processors.
	configStr := `processors:
- add_fields: {fields: {testField: value}}
- add_fields: {fields: {testField2: stuff}}`
	config, err := functionConfigFromString(configStr)
	if err != nil {
		t.Fatal(err)
	}
	processors, err := processorsForFunction(
		beat.Info{}, config)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 2, len(processors.List))
}

// Helper function to convert from YML input string to an unpacked
// fnExtraConfig
func functionConfigFromString(s string) (fnExtraConfig, error) {
	config := fnExtraConfig{}
	cfg, err := conf.NewConfigFrom(s)
	if err != nil {
		return config, err
	}
	err = cfg.Unpack(&config)
	return config, err
}
