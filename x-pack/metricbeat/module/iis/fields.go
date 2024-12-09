// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package iis

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "iis", asset.ModuleFieldsPri, AssetIis); err != nil {
		panic(err)
	}
}

// AssetIis returns asset data.
// This is the base64 encoded zlib format compressed contents of module/iis.
func AssetIis() string {
	return "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"
}
