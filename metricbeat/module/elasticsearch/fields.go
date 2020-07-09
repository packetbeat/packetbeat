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

package elasticsearch

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "elasticsearch", asset.ModuleFieldsPri, AssetElasticsearch); err != nil {
		panic(err)
	}
}

// AssetElasticsearch returns asset data.
// This is the base64 encoded gzipped contents of module/elasticsearch.
func AssetElasticsearch() string {
	return "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"
}
