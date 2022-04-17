// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package istio

import (
	"github.com/menderesk/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "istio", asset.ModuleFieldsPri, AssetIstio); err != nil {
		panic(err)
	}
}

// AssetIstio returns asset data.
// This is the base64 encoded zlib format compressed contents of module/istio.
func AssetIstio() string {
	return "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"
}
