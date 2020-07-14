// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package gsuite

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "gsuite", asset.ModuleFieldsPri, AssetGsuite); err != nil {
		panic(err)
	}
}

// AssetGsuite returns asset data.
// This is the base64 encoded gzipped contents of module/gsuite.
func AssetGsuite() string {
	return "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"
}
