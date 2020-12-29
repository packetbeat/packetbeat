// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package nomad

import (
	"time"

	"github.com/elastic/beats/v7/libbeat/autodiscover/template"
	"github.com/elastic/beats/v7/libbeat/common"
)

// Config for nomad autodiscover provider
type Config struct {
	Address        string        `config:"address"`
	Region         string        `config:"region"`
	Namespace      string        `config:"namespace"`
	SecretID       string        `config:"secret_id"`
	Node           string        `config:"node"`
	CleanupTimeout time.Duration `config:"cleanup_timeout" validate:"positive"`

	Prefix    string                  `config:"prefix"`
	Hints     *common.Config          `config:"hints"`
	Builders  []*common.Config        `config:"builders"`
	Appenders []*common.Config        `config:"appenders"`
	Templates template.MapperSettings `config:"templates"`

	waitTime   time.Duration
	syncPeriod time.Duration
	allowStale bool
}

func defaultConfig() *Config {
	return &Config{
		Address:        "http://127.0.0.1:4646",
		Region:         "",
		Namespace:      "",
		SecretID:       "",
		allowStale:     true,
		waitTime:       15 * time.Second,
		syncPeriod:     30 * time.Second,
		CleanupTimeout: 15 * time.Second,
		Prefix:         "co.elastic",
	}
}

// Validate ensures correctness of config
func (c *Config) Validate() {
	// Make sure that prefix doesn't ends with a '.'
	if c.Prefix[len(c.Prefix)-1] == '.' && c.Prefix != "." {
		c.Prefix = c.Prefix[:len(c.Prefix)-2]
	}
}
