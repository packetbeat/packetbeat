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

package beat

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "beat", asset.ModuleFieldsPri, AssetBeat); err != nil {
		panic(err)
	}
}

// AssetBeat returns asset data.
// This is the base64 encoded zlib format compressed contents of module/beat.
func AssetBeat() string {
	return "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"
}
