// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package cloudflare

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "cloudflare", asset.ModuleFieldsPri, AssetCloudflare); err != nil {
		panic(err)
	}
}

// AssetCloudflare returns asset data.
// This is the base64 encoded gzipped contents of module/cloudflare.
func AssetCloudflare() string {
	return "eJxsjEEKwzAMBO9+xZJ7PqBDf9BHhEgpompkbPng3xeXgkvJHmeZWfGUTtjNGx+2FUlAaJgQlgmXBLDUvWgO9ZNwSwB+LNydmw35UDGu9PlXnNtL/upj0bMQHsVb/pKL+my9AwAA///9kzKA"
}
