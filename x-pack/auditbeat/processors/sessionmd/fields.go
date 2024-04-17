// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package sessionmd

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("auditbeat", "fields.yml", asset.BeatFieldsPri, AssetFieldsYml); err != nil {
		panic(err)
	}
}

// AssetFieldsYml returns asset data.
// This is the base64 encoded zlib format compressed contents of fields.yml.
func AssetFieldsYml() string {
	return "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"
}
