// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package statsd

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "statsd", asset.ModuleFieldsPri, AssetStatsd); err != nil {
		panic(err)
	}
}

// AssetStatsd returns asset data.
// This is the base64 encoded gzipped contents of module/statsd.
func AssetStatsd() string {
	return "eJyUkktuwyAQhvec4pc3kSwlB2DRS/QAETETiwYziBmr8u0rP1KR1l5klv/D/gY440GThahT8QbQoJEsms9FaAzgSboSsgZOFh8GAFYTA/sxkgEKRXJCFr0zwD1Q9GKX5BnJDVR9fx6d8pwtPOZNqSt17dReOh6Tnn6dZ5tvX9RpJa/CdXUjp37fuw4u55D6LdjMyaaK7mz7nG3rBYiK7MG274LeI7sD8y9p+x7mQFpCJy+3EN2Nolxacwj4Aveg6ZuL/28doh1gbUjr76tzq9/NTwAAAP//fPG2zQ=="
}
