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

package rabbitmq

import (
	"github.com/elastic/beats/v8/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "rabbitmq", asset.ModuleFieldsPri, AssetRabbitmq); err != nil {
		panic(err)
	}
}

// AssetRabbitmq returns asset data.
// This is the base64 encoded zlib format compressed contents of module/rabbitmq.
func AssetRabbitmq() string {
	return "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"
}
