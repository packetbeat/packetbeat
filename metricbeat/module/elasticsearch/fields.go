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
// This is the base64 encoded zlib format compressed contents of module/elasticsearch.
func AssetElasticsearch() string {
	return "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"
}
