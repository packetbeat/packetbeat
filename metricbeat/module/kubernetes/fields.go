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

package kubernetes

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "kubernetes", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "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"
}
