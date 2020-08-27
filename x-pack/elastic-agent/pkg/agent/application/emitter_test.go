// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package application

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/agent/transpiler"
	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/composable"
)

func TestRenderInputs(t *testing.T) {
	testcases := map[string]struct {
		input     transpiler.Node
		expected  transpiler.Node
		varsArray []composable.Vars
		err       bool
	}{
		"inputs not list": {
			input: transpiler.NewKey("inputs", transpiler.NewStrVal("not list")),
			err:   true,
			varsArray: []composable.Vars{
				{
					Mapping: map[string]interface{}{},
				},
			},
		},
		"bad variable error": {
			input: transpiler.NewKey("inputs", transpiler.NewList([]transpiler.Node{
				transpiler.NewDict([]transpiler.Node{
					transpiler.NewKey("key", transpiler.NewStrVal("{{var1.name|'missing ending quote}}")),
				}),
			})),
			err: true,
			varsArray: []composable.Vars{
				{
					Mapping: map[string]interface{}{
						"var1": map[string]interface{}{
							"name": "value1",
						},
					},
				},
			},
		},
		"basic single var": {
			input: transpiler.NewKey("inputs", transpiler.NewList([]transpiler.Node{
				transpiler.NewDict([]transpiler.Node{
					transpiler.NewKey("key", transpiler.NewStrVal("{{var1.name}}")),
				}),
			})),
			expected: transpiler.NewList([]transpiler.Node{
				transpiler.NewDict([]transpiler.Node{
					transpiler.NewKey("key", transpiler.NewStrVal("value1")),
				}),
			}),
			varsArray: []composable.Vars{
				{
					Mapping: map[string]interface{}{
						"var1": map[string]interface{}{
							"name": "value1",
						},
					},
				},
			},
		},
		"duplicate result is removed": {
			input: transpiler.NewKey("inputs", transpiler.NewList([]transpiler.Node{
				transpiler.NewDict([]transpiler.Node{
					transpiler.NewKey("key", transpiler.NewStrVal("{{var1.name}}")),
				}),
				transpiler.NewDict([]transpiler.Node{
					transpiler.NewKey("key", transpiler.NewStrVal("{{var1.diff}}")),
				}),
			})),
			expected: transpiler.NewList([]transpiler.Node{
				transpiler.NewDict([]transpiler.Node{
					transpiler.NewKey("key", transpiler.NewStrVal("value1")),
				}),
			}),
			varsArray: []composable.Vars{
				{
					Mapping: map[string]interface{}{
						"var1": map[string]interface{}{
							"name": "value1",
							"diff": "value1",
						},
					},
				},
			},
		},
		"missing var removes input": {
			input: transpiler.NewKey("inputs", transpiler.NewList([]transpiler.Node{
				transpiler.NewDict([]transpiler.Node{
					transpiler.NewKey("key", transpiler.NewStrVal("{{var1.name}}")),
				}),
				transpiler.NewDict([]transpiler.Node{
					transpiler.NewKey("key", transpiler.NewStrVal("{{var1.missing|var1.diff}}")),
				}),
				transpiler.NewDict([]transpiler.Node{
					transpiler.NewKey("key", transpiler.NewStrVal("{{var1.removed}}")),
				}),
			})),
			expected: transpiler.NewList([]transpiler.Node{
				transpiler.NewDict([]transpiler.Node{
					transpiler.NewKey("key", transpiler.NewStrVal("value1")),
				}),
			}),
			varsArray: []composable.Vars{
				{
					Mapping: map[string]interface{}{
						"var1": map[string]interface{}{
							"name": "value1",
							"diff": "value1",
						},
					},
				},
			},
		},
		"duplicate var result but unique input not removed": {
			input: transpiler.NewKey("inputs", transpiler.NewList([]transpiler.Node{
				transpiler.NewDict([]transpiler.Node{
					transpiler.NewKey("key", transpiler.NewStrVal("{{var1.name}}")),
					transpiler.NewKey("unique", transpiler.NewStrVal("0")),
				}),
				transpiler.NewDict([]transpiler.Node{
					transpiler.NewKey("key", transpiler.NewStrVal("{{var1.diff}}")),
					transpiler.NewKey("unique", transpiler.NewStrVal("1")),
				}),
			})),
			expected: transpiler.NewList([]transpiler.Node{
				transpiler.NewDict([]transpiler.Node{
					transpiler.NewKey("key", transpiler.NewStrVal("value1")),
					transpiler.NewKey("unique", transpiler.NewStrVal("0")),
				}),
				transpiler.NewDict([]transpiler.Node{
					transpiler.NewKey("key", transpiler.NewStrVal("value1")),
					transpiler.NewKey("unique", transpiler.NewStrVal("1")),
				}),
			}),
			varsArray: []composable.Vars{
				{
					Mapping: map[string]interface{}{
						"var1": map[string]interface{}{
							"name": "value1",
							"diff": "value1",
						},
					},
				},
			},
		},
		"duplicates across vars array handled": {
			input: transpiler.NewKey("inputs", transpiler.NewList([]transpiler.Node{
				transpiler.NewDict([]transpiler.Node{
					transpiler.NewKey("key", transpiler.NewStrVal("{{var1.name}}")),
				}),
				transpiler.NewDict([]transpiler.Node{
					transpiler.NewKey("key", transpiler.NewStrVal("{{var1.diff}}")),
				}),
			})),
			expected: transpiler.NewList([]transpiler.Node{
				transpiler.NewDict([]transpiler.Node{
					transpiler.NewKey("key", transpiler.NewStrVal("value1")),
				}),
				transpiler.NewDict([]transpiler.Node{
					transpiler.NewKey("key", transpiler.NewStrVal("value2")),
				}),
				transpiler.NewDict([]transpiler.Node{
					transpiler.NewKey("key", transpiler.NewStrVal("value3")),
				}),
				transpiler.NewDict([]transpiler.Node{
					transpiler.NewKey("key", transpiler.NewStrVal("value4")),
				}),
			}),
			varsArray: []composable.Vars{
				{
					Mapping: map[string]interface{}{
						"var1": map[string]interface{}{
							"name": "value1",
							"diff": "value1",
						},
					},
				},
				{
					Mapping: map[string]interface{}{
						"var1": map[string]interface{}{
							"name": "value1",
							"diff": "value2",
						},
					},
				},
				{
					Mapping: map[string]interface{}{
						"var1": map[string]interface{}{
							"name": "value1",
							"diff": "value3",
						},
					},
				},
				{
					Mapping: map[string]interface{}{
						"var1": map[string]interface{}{
							"name": "value1",
							"diff": "value2",
						},
					},
				},
				{
					Mapping: map[string]interface{}{
						"var1": map[string]interface{}{
							"name": "value1",
							"diff": "value4",
						},
					},
				},
			},
		},
	}

	for name, test := range testcases {
		t.Run(name, func(t *testing.T) {
			v, err := renderInputs(test.input, test.varsArray)
			if test.err {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				if !assert.True(t, reflect.DeepEqual(test.expected, v)) {
					t.Logf(
						`received: %+v
					 	 expected: %+v`, v, test.expected)
				}
			}
		})
	}
}
