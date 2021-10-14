// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package gcp

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "gcp", asset.ModuleFieldsPri, AssetGcp); err != nil {
		panic(err)
	}
}

// AssetGcp returns asset data.
// This is the base64 encoded zlib format compressed contents of module/gcp.
func AssetGcp() string {
	return "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"
}
