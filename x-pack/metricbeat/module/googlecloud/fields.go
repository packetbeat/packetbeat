// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package googlecloud

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "googlecloud", asset.ModuleFieldsPri, AssetGooglecloud); err != nil {
		panic(err)
	}
}

// AssetGooglecloud returns asset data.
// This is the base64 encoded gzipped contents of module/googlecloud.
func AssetGooglecloud() string {
	return "eJy8lr1u2zAQx3c/xW0BiiYPoKGLhzRD2wxNV4MiTzJrikeQRxvq0xeULEGyLcUppGaK+fH//+7L9CMcsM6glG4DwJoNZvDwTFQahK2hqODVCC7IVw8bAI8GRcAMcmSxAVAYpNeONdkMvmwAAJ63r1CRigY3AIVGo0LWbDyCFRVmUDbiMmk/GZGjCc02ANcOM6D8N0o+Lw3vDzViQP/0qV++eTf9tQu7dveA9Ym8uhKrkIUSLJYU9FqGxfRCHRirxeQkVS4yDq5dlnXsUnqKbrA6KvqoVbatcpeAwZ3LQg55tA0srMTR5pT5lNhQsNAeT8KYqwNzonPCQ3HlyTlUu7xmDDtJ0fLN852XIVtOHBgl8sVKqrQtoRHubCCvgfc4F9I1mhPygLwK3Fn6bry+51xcpRoeA/ojqp0kj2E2VkUxN5dNdjPa77HK0QMV0Kj2JkC2iXZPgdNu+n+ieceUkbXRf0RSXwjxZ0q6FzJ96lCEMSQFo4Lt6xvwXjDoADJ6j5ZNDdqm780uiPvAgyhxx7qagvoo91sShIJ8oj2nV1sIKMmq6/r1na3DYaX+EcuP8jYppaq0o5w8oPBUTYVxgUNuFZhk3rK8/ABy6Jt+vD0zHc/Ja8a185NMGC0wvZ+gFmjNDDUO76SoHw83MRqzACPzr3RqzjUz+esb7EWAHNGCj9ZqW36+Zzws8on8WhMiUR9XePCupqT1aSclZWM6qht0y795PV/35v0bYUDLa+cueQAd0X+M6j/kbJ7sceKX5+X63wAAAP//889N8w=="
}
