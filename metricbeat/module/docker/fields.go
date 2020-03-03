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
// This is the base64 encoded gzipped contents of ../metricbeat/module/docker.
func AssetDocker() string {
	return "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"
}
