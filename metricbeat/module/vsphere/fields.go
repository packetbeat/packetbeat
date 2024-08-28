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
