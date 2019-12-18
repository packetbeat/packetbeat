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
