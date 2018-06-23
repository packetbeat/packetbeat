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
	if err := asset.SetFields("auditbeat", "fields.yml", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "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"
}
