// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package capabilities

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshal(t *testing.T) {
	t.Run("valid json", func(t *testing.T) {
		rr := make(ruleDefinitions, 0, 0)

		err := json.Unmarshal(jsonDefinitionValid, &rr)

		assert.Nil(t, err, "no error is expected")
		assert.Equal(t, 3, len(rr))
		assert.Equal(t, "*capabilities.upgradeCapability", reflect.TypeOf(rr[0]).String())
		assert.Equal(t, "*capabilities.inputCapability", reflect.TypeOf(rr[1]).String())
		assert.Equal(t, "*capabilities.outputCapability", reflect.TypeOf(rr[2]).String())
	})

	t.Run("invalid json", func(t *testing.T) {
		var rr ruleDefinitions

		err := json.Unmarshal(jsonDefinitionInvalid, &rr)

		assert.Error(t, err, "error is expected")
	})
}

var jsonDefinitionValid = []byte(`[{
	"upgrade": "${version} == '8.0.0'",
	"rule": "allow"
},
{
	"input": "system/metrics",
	"rule": "allow"
},
{
	"output": "elasticsearch",
	"rule": "allow"
}
]`)

var jsonDefinitionInvalid = []byte(`[{
	"upgrade": "${version} == '8.0.0'",
	"rule": "allow"
},
{
	"input": "system/metrics",
	"rule": "allow"
},
{
	"output": "elasticsearch",
	"rule": "allow"
},
{
	"ayay": "elasticsearch",
	"rule": "allow"
}
]`)
