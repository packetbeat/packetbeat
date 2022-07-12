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

package kubernetes

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "kubernetes", asset.ModuleFieldsPri, AssetKubernetes); err != nil {
		panic(err)
	}
}

// AssetKubernetes returns asset data.
// This is the base64 encoded zlib format compressed contents of module/kubernetes.
func AssetKubernetes() string {
	return "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"
}
