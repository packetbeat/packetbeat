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
