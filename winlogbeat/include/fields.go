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

package include

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("winlogbeat", "fields.yml", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "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"
}
