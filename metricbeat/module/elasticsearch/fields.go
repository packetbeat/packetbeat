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
