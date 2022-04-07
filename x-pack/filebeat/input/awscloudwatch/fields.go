// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package awscloudwatch

import (
	"github.com/elastic/beats/v8/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "awscloudwatch", asset.ModuleFieldsPri, AssetAwscloudwatch); err != nil {
		panic(err)
	}
}

// AssetAwscloudwatch returns asset data.
// This is the base64 encoded zlib format compressed contents of input/awscloudwatch.
func AssetAwscloudwatch() string {
	return "eJyskUFq7DAQRPc+RTF7zwG0+PAJ5AIJzNIoVtsSI6uN1I7w7YPkIWMnWSQhvVR31XugFldaFXRObe95MVlLbxtAnHhSOB0XpwYwlProZnEcFP41APDoyJuEIfKE/5cnPJTApQTgeUznBhjqiarnLYKe6AtoGUODXrx0NaAgcaHbRtaZFMbIy/x++0nlmzpl9kp7Lc9jt6fc2VdaM0ezez8YPFuqFeABYqkUbboQRrautxDrEuiVguCFPIe7zZGfJJKe/kJga/qZgQsjpVLZiZvoNxYlVxU2UtbpVkoGLnz4lXPzFgAA///o3McA"
}
