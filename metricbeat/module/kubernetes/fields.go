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
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "kubernetes", asset.ModuleFieldsPri, AssetKubernetes); err != nil {
		panic(err)
	}
}

// AssetKubernetes returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/kubernetes.
func AssetKubernetes() string {
	return "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"
}
