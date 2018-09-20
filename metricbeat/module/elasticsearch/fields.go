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
	if err := asset.SetFields("metricbeat", "elasticsearch", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "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"
}
