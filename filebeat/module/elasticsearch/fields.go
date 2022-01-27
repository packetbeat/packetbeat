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
	if err := asset.SetFields("filebeat", "elasticsearch", asset.ModuleFieldsPri, AssetElasticsearch); err != nil {
		panic(err)
	}
}

// AssetElasticsearch returns asset data.
// This is the base64 encoded zlib format compressed contents of module/elasticsearch.
func AssetElasticsearch() string {
	return "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"
}
