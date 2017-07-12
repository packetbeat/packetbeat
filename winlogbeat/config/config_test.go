// +build !integration

package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type validationTestCase struct {
	config Validator
	errMsg string
}

func (v validationTestCase) run(t *testing.T) {
	if v.errMsg == "" {
		assert.NoError(t, v.config.Validate())
	} else {
		err := v.config.Validate()
		if err != nil {
			assert.Contains(t, err.Error(), v.errMsg)
		} else {
			t.Errorf("expected error with '%s'", v.errMsg)
		}
	}
}

func TestConfigValidate(t *testing.T) {
	testCases := []validationTestCase{
		// Top-level config
		{
			WinlogbeatConfig{
				EventLogs: []map[string]interface{}{
					{"Name": "App"},
				},
			},
			"", // No Error
		},
		{
			WinlogbeatConfig{},
			"1 error: At least one event log must be configured as part of " +
				"event_logs",
		},
	}

	for _, test := range testCases {
		test.run(t)
	}
}
