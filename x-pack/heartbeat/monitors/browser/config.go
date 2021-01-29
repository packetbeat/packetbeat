// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package browser

import (
	"fmt"

	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/x-pack/heartbeat/monitors/browser/source"
)

type Config struct {
	Schedule  string                 `config:"schedule"`
	Params    map[string]interface{} `config:"params"`
	RawConfig *common.Config
	Source    *source.Source `config:"source"`
	// Name is optional for lightweight checks but required for browsers
	Name string `config:"id"`
	// Id is optional for lightweight checks but required for browsers
	Id string `config:"name"`
}

func (c *Config) Validate() error {
	if c.Name == "" {
		return fmt.Errorf("config 'name' must be specified for this moniotr")
	}
	if c.Id == "" {
		return fmt.Errorf("config 'id' must be specified for this monitor")
	}

	return nil
}
