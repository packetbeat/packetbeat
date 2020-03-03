// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package aws

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "aws", asset.ModuleFieldsPri, AssetAws); err != nil {
		panic(err)
	}
}

// AssetAws returns asset data.
// This is the base64 encoded gzipped contents of module/aws.
func AssetAws() string {
	return "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"
}
