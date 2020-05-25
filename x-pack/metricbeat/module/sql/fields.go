// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package sql

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "sql", asset.ModuleFieldsPri, AssetSql); err != nil {
		panic(err)
	}
}

// AssetSql returns asset data.
// This is the base64 encoded gzipped contents of module/sql.
func AssetSql() string {
	return "eJykkkFuwjAQRfc5xRfLSnAAL7rqEiEhDlA59gdcnBjsMW1uX5HEKFWKVFQvv2fmvUy8xImdQrr4ChAnngqL3Xa9qIBIT52oUFN0BVgmE91ZXGgVXisA2G3XaILNnthTzJEJDSU6k7CPoYHuK6wWXevECtg7eptU37xEqxsW+O1Id6bCIYZ8HpNp/bTHRndlvMel9cTuM0Q7yX+RLuetn4GcaCEB/KLJQsiRuGTGbjWj9vH/oNvbiMLquSZ4TyNlcXNquWhzw+jM6mVmEOoPGpnEQ/A+3NqQa8+/6W0Gxv0vjnK0j7WSRNcenrZ6amub0C7Hz8dV+8wHZj+f7HcAAAD//yd20AA="
}
