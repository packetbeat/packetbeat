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

package foo

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "foo", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "eJx8j0EOwiAQRfdzip/uewEW7jyFcYEyGFLoEEpje3vTooY21bf8E94LLTqeFawIAdllzwqNFWkIMDzck4vZSa9wIgDgKXJygfus/eVK62ZFEMSMngmwjr0Z1Hpo0evAH/lCniMrPJKM8b0cNLaSWnTT6bsdyX4KC9vn+0gd4kmHuP6npgQ7np+SzO72J7twLsISpVcAAAD//+HhYIk="
}
