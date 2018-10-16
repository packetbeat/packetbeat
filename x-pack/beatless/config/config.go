// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import (
	"fmt"

	"github.com/elastic/beats/libbeat/common"
)

// ConfigOverrides overrides the defaults provided by libbeat.
var ConfigOverrides = common.MustNewConfigFrom(map[string]interface{}{
	"path.data":              "/tmp",
	"path.logs":              "/tmp/logs",
	"logging.to_stderr":      true,
	"logging.to_files":       false,
	"logging.level":          "debug",
	"setup.template.enabled": true,
	"queue.mem": map[string]interface{}{
		"events":           "${output.elasticsearch.bulk_max_size}",
		"flush.min_events": 10,
		"flush.timeout":    "0.01s",
	},
	"output.elasticsearch.bulk_max_size": 50,
})

// Config default configuration for Beatless.
type Config struct {
	Provider *common.ConfigNamespace `config:"provider" validate:"required"`
}

// ProviderConfig is a generic configured used by providers.
type ProviderConfig struct {
	Functions []*common.Config `config:"functions"`
}

// FunctionConfig minimal configuration from each function.
type FunctionConfig struct {
	Type    string `config:"type"`
	Name    string `config:"name"`
	Enabled bool   `config:"enabled"`
}

// DefaultConfig is the default configuration for Beatless.
var DefaultConfig = Config{}

// DefaultFunctionConfig is the default configuration for new function.
var DefaultFunctionConfig = FunctionConfig{
	Enabled: true,
}

// Validate enforces that function names are unique.
func (p *ProviderConfig) Validate() error {
	names := make(map[string]bool)
	for _, rawfn := range p.Functions {
		fc := FunctionConfig{}
		rawfn.Unpack(&fc)

		if !fc.Enabled {
			return nil
		}

		if _, found := names[fc.Name]; found {
			return fmt.Errorf("function name '%s' already exist, name must be unique", fc.Name)
		}

		names[fc.Name] = true
	}
	return nil
}
