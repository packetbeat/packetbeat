// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package salesforce

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "salesforce", asset.ModuleFieldsPri, AssetSalesforce); err != nil {
		panic(err)
	}
}

// AssetSalesforce returns asset data.
// This is the base64 encoded zlib format compressed contents of module/salesforce.
func AssetSalesforce() string {
	return "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"
}
