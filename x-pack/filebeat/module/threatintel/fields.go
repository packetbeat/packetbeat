// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package threatintel

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "threatintel", asset.ModuleFieldsPri, AssetThreatintel); err != nil {
		panic(err)
	}
}

// AssetThreatintel returns asset data.
// This is the base64 encoded zlib format compressed contents of module/threatintel.
func AssetThreatintel() string {
	return "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"
}
