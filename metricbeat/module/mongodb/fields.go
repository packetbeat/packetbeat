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

package mongodb

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "mongodb", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "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"
}
