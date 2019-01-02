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

package dropwizard

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "dropwizard", asset.ModuleFieldsPri, Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "eJxsjkGOwjAMRfc5xVfW0zmAF7OaG8xyhFBInDYiTaLYCJXTowICFeHls/z8Bhx5IYRe2zldXA8G0KSZCfb3Ca0BAovvqWmqhfBjAOBPnQp8zZm9ckDsdcbr6tsAnTM7YcKB1RlAWDWVUQj/ViTbL9hJtdnduptq172vJaaREF0WNkBMnIPQ7eGA4mZ+y11Hl8aEsddTe5APuVvX3bftuwYAAP//aMZUBA=="
}
