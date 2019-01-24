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

package system

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "system", asset.ModuleFieldsPri, AssetSystem); err != nil {
		panic(err)
	}
}

// AssetSystem returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/system.
func AssetSystem() string {
	return "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"
}
