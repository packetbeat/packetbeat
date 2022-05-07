// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package azure

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "azure", asset.ModuleFieldsPri, AssetAzure); err != nil {
		panic(err)
	}
}

// AssetAzure returns asset data.
// This is the base64 encoded zlib format compressed contents of module/azure.
func AssetAzure() string {
	return "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"
}
