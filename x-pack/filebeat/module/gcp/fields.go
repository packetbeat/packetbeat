// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package gcp

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "gcp", asset.ModuleFieldsPri, AssetGcp); err != nil {
		panic(err)
	}
}

// AssetGcp returns asset data.
// This is the base64 encoded gzipped contents of module/gcp.
func AssetGcp() string {
	return "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"
}
