// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package decode_cef

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "decode_cef", asset.ModuleFieldsPri, AssetDecodeCef); err != nil {
		panic(err)
	}
}

// AssetDecodeCef returns asset data.
// This is the base64 encoded gzipped contents of processors/decode_cef.
func AssetDecodeCef() string {
	return "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"
}
