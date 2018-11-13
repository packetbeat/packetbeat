// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package management

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/beats/x-pack/libbeat/management/api"
)

func TestConfigBlacklist(t *testing.T) {
	tests := []struct {
		name        string
		patterns    map[string]string
		blocks      api.ConfigBlocks
		blacklisted bool

		// only fill if blacklisted == true
		expected api.ConfigBlocks
	}{
		{
			name:        "No patterns",
			blacklisted: false,
			blocks: api.ConfigBlocks{
				api.ConfigBlocksWithType{
					Type: "output",
					Blocks: []*api.ConfigBlock{
						{
							Raw: map[string]interface{}{
								"output": "console",
							},
						},
					},
				},
			},
		},
		{
			name:        "Blacklisted dict key",
			blacklisted: true,
			patterns: map[string]string{
				"output": "^console$",
			},
			blocks: api.ConfigBlocks{
				api.ConfigBlocksWithType{
					Type: "output",
					Blocks: []*api.ConfigBlock{
						{
							Raw: map[string]interface{}{
								"console": map[string]interface{}{
									"pretty": "true",
								},
							},
						},
						{
							Raw: map[string]interface{}{
								"elasticsearch": map[string]interface{}{
									"host": "localhost",
								},
							},
						},
					},
				},
			},
			expected: api.ConfigBlocks{
				api.ConfigBlocksWithType{
					Type: "output",
					Blocks: []*api.ConfigBlock{
						{
							Raw: map[string]interface{}{
								"elasticsearch": map[string]interface{}{
									"host": "localhost",
								},
							},
						},
					},
				},
			},
		},
		{
			name:        "Blacklisted value key",
			blacklisted: true,
			patterns: map[string]string{
				"metricbeat.modules.module": "k.{8}s",
			},
			blocks: api.ConfigBlocks{
				api.ConfigBlocksWithType{
					Type: "metricbeat.modules",
					Blocks: []*api.ConfigBlock{
						{
							Raw: map[string]interface{}{
								"module": "kubernetes",
								"hosts":  "localhost:10255",
							},
						},
					},
				},
			},
		},
		{
			name:        "Blacklisted value in a list",
			blacklisted: true,
			patterns: map[string]string{
				"metricbeat.modules.metricsets": "event",
			},
			blocks: api.ConfigBlocks{
				api.ConfigBlocksWithType{
					Type: "metricbeat.modules",
					Blocks: []*api.ConfigBlock{
						{
							Raw: map[string]interface{}{
								"module": "kubernetes",
								"metricsets": []string{
									"event",
									"default",
								},
							},
						},
						{
							Raw: map[string]interface{}{
								"module": "kubernetes",
								"metricsets": []string{
									"default",
								},
							},
						},
					},
				},
			},
			expected: api.ConfigBlocks{
				api.ConfigBlocksWithType{
					Type: "metricbeat.modules",
					Blocks: []*api.ConfigBlock{
						{
							Raw: map[string]interface{}{
								"module": "kubernetes",
								"metricsets": []string{
									"default",
								},
							},
						},
					},
				},
			},
		},
		{
			name:        "Blacklisted value in a deep list",
			blacklisted: true,
			patterns: map[string]string{
				"filebeat.inputs.containers.ids": "1ffeb0dbd13",
			},
			blocks: api.ConfigBlocks{
				api.ConfigBlocksWithType{
					Type: "metricbeat.modules",
					Blocks: []*api.ConfigBlock{
						{
							Raw: map[string]interface{}{
								"module": "kubernetes",
								"metricsets": []string{
									"event",
									"default",
								},
							},
						},
					},
				},
				api.ConfigBlocksWithType{
					Type: "filebeat.inputs",
					Blocks: []*api.ConfigBlock{
						{
							Raw: map[string]interface{}{
								"type": "docker",
								"containers": map[string]interface{}{
									"ids": []string{
										"1ffeb0dbd13",
									},
								},
							},
						},
						{
							Raw: map[string]interface{}{
								"type": "docker",
								"containers": map[string]interface{}{
									"ids": []string{
										"256425931c2",
									},
								},
							},
						},
					},
				},
			},
			expected: api.ConfigBlocks{
				api.ConfigBlocksWithType{
					Type: "metricbeat.modules",
					Blocks: []*api.ConfigBlock{
						{
							Raw: map[string]interface{}{
								"module": "kubernetes",
								"metricsets": []string{
									"event",
									"default",
								},
							},
						},
					},
				},
				api.ConfigBlocksWithType{
					Type: "filebeat.inputs",
					Blocks: []*api.ConfigBlock{
						{
							Raw: map[string]interface{}{
								"type": "docker",
								"containers": map[string]interface{}{
									"ids": []string{
										"256425931c2",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name:        "Blacklisted dict key in a list",
			blacklisted: true,
			patterns: map[string]string{
				"list.of.elements":            "forbidden",
				"list.of.elements.disallowed": "yes",
			},
			blocks: api.ConfigBlocks{
				api.ConfigBlocksWithType{
					Type: "list",
					Blocks: []*api.ConfigBlock{
						{
							Raw: map[string]interface{}{
								"of": map[string]interface{}{
									"elements": []interface{}{
										map[string]interface{}{
											"forbidden": "yes",
										},
									},
								},
							},
						},
						{
							Raw: map[string]interface{}{
								"of": map[string]interface{}{
									"elements": []interface{}{
										map[string]interface{}{
											"allowed": "yes",
										},
									},
								},
							},
						},
						{
							Raw: map[string]interface{}{
								"of": map[string]interface{}{
									"elements": []interface{}{
										map[string]interface{}{
											"disallowed": "yes",
										},
									},
								},
							},
						},
					},
				},
			},
			expected: api.ConfigBlocks{
				api.ConfigBlocksWithType{
					Type: "list",
					Blocks: []*api.ConfigBlock{
						{
							Raw: map[string]interface{}{
								"of": map[string]interface{}{
									"elements": []interface{}{
										map[string]interface{}{
											"allowed": "yes",
										},
									},
								},
							},
						},
					},
				}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			bl, err := NewConfigBlacklist(test.patterns)
			if err != nil {
				t.Fatal(err)
			}

			result := bl.Filter(test.blocks)
			equal := api.ConfigBlocksEqual(result, test.blocks)
			assert.Equal(t, test.blacklisted, !equal)

			if test.blacklisted {
				assert.Equal(t, test.expected, result)
			}
		})
	}
}
