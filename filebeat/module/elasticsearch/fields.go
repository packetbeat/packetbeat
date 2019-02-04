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
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "elasticsearch", asset.ModuleFieldsPri, AssetElasticsearch); err != nil {
		panic(err)
	}
}

// AssetElasticsearch returns asset data.
// This is the base64 encoded gzipped contents of module/elasticsearch.
func AssetElasticsearch() string {
	return "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"
}
