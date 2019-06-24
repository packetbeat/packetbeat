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
