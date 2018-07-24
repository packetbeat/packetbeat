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

package system

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "system", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "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"
}
