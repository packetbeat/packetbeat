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
// This is the base64 encoded gzipped contents of module/kubernetes.
func AssetKubernetes() string {
	return "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"
}
