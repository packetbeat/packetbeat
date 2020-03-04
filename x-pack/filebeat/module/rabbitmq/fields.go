// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package rabbitmq

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "rabbitmq", asset.ModuleFieldsPri, AssetRabbitmq); err != nil {
		panic(err)
	}
}

// AssetRabbitmq returns asset data.
// This is the base64 encoded gzipped contents of module/rabbitmq.
func AssetRabbitmq() string {
	return "eJx0kMFuhSAURPd8xcS9xrgkjbsuXbTpD6BcKRGFAqb17xttNeLz3eWc5Mzk5hho4fCibXUcvxgQdTTEkb1vUfOWMUBS6Lx2UduJo2YAsGM0Vs6GGNBrMjLwjeaYxEiJd724OOJQ3s7uP7kxp6azzVh1ZHeyp8K/OyYbq9BrQ+GEr53nXqdlku/dAy3f1l9ZsuDjk/DqjZgUnLcdhYAHGf2I0a0/fymLqqqKsma/AQAA///y5GyB"
}
