// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package awsfargate

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "awsfargate", asset.ModuleFieldsPri, AssetAwsfargate); err != nil {
		panic(err)
	}
}

// AssetAwsfargate returns asset data.
// This is the base64 encoded zlib format compressed contents of module/awsfargate.
func AssetAwsfargate() string {
	return "eJx0kEFqwzAQRfc+xb9AfAAtCqE0u6666Hpqj4XoWGNGE0p6+mITGTXEf/mR5j3+Cd98C6CfMpFFcu4ATy4ccP78wGUvjYWpcMAXO3XAyGWwtHjSHPDSAcC7jldhTGoYVIQHTzli0OyUMhtEY8FkOuM8069mvL3uhL4DpsQylrDdOiHTzA9ia/y2cEA0vS735onJmst27Yi3yfT3xy24hYvGvXtGXvO4S82BVWumVsWq1P+t+uZXVfwLAAD//0tEf3I="
}
