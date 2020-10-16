// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package virustotal

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "virustotal", asset.ModuleFieldsPri, AssetVirustotal); err != nil {
		panic(err)
	}
}

// AssetVirustotal returns asset data.
// This is the base64 encoded gzipped contents of module/virustotal.
func AssetVirustotal() string {
	return "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"
}
