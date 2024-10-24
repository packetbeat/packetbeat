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

package docker

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "docker", asset.ModuleFieldsPri, AssetDocker); err != nil {
		panic(err)
	}
}

// AssetDocker returns asset data.
// This is the base64 encoded zlib format compressed contents of module/docker.
func AssetDocker() string {
	return "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"
}
