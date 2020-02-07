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
