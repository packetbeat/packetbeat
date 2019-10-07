// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package azure

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "azure", asset.ModuleFieldsPri, AssetAzure); err != nil {
		panic(err)
	}
}

// AssetAzure returns asset data.
// This is the base64 encoded gzipped contents of module/azure.
func AssetAzure() string {
	return "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"
}
