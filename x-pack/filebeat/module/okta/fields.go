// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package okta

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "okta", asset.ModuleFieldsPri, AssetOkta); err != nil {
		panic(err)
	}
}

// AssetOkta returns asset data.
// This is the base64 encoded gzipped contents of module/okta.
func AssetOkta() string {
	return "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"
}
