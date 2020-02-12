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

package ceph

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "ceph", asset.ModuleFieldsPri, AssetCeph); err != nil {
		panic(err)
	}
}

// AssetCeph returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/ceph.
func AssetCeph() string {
	return "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"
}
