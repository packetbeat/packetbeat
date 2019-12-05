// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package aws

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "aws", asset.ModuleFieldsPri, AssetAws); err != nil {
		panic(err)
	}
}

// AssetAws returns asset data.
// This is the base64 encoded gzipped contents of module/aws.
func AssetAws() string {
	return "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"
}
