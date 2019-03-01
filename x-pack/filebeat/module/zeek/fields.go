// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package zeek

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "zeek", asset.ModuleFieldsPri, AssetZeek); err != nil {
		panic(err)
	}
}

// AssetZeek returns asset data.
// This is the base64 encoded gzipped contents of module/zeek.
func AssetZeek() string {
	return "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"
}
