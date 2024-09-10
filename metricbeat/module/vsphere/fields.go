// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package vsphere

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "vsphere", asset.ModuleFieldsPri, AssetVsphere); err != nil {
		panic(err)
	}
}

// AssetVsphere returns asset data.
// This is the base64 encoded zlib format compressed contents of module/vsphere.
func AssetVsphere() string {
	return "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"
}
