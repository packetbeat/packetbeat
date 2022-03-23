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

package linux

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "linux", asset.ModuleFieldsPri, AssetLinux); err != nil {
		panic(err)
	}
}

// AssetLinux returns asset data.
// This is the base64 encoded zlib format compressed contents of module/linux.
func AssetLinux() string {
	return "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"
}
