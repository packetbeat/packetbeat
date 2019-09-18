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
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "docker", asset.ModuleFieldsPri, AssetDocker); err != nil {
		panic(err)
	}
}

// AssetDocker returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/docker.
func AssetDocker() string {
	return "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"
}
