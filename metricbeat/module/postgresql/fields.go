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

package postgresql

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "postgresql", asset.ModuleFieldsPri, AssetPostgresql); err != nil {
		panic(err)
	}
}

// AssetPostgresql returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/postgresql.
func AssetPostgresql() string {
	return "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"
}
