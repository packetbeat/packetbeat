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

package icmp

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("packetbeat", "icmp", asset.ModuleFieldsPri, Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "eJy0kDFvgzAQhXd+xVP2MFUdGCpVnTpU6tC9cswBVozP9Rkq/n1FMBFKCgNSmPD53vee3xFnGgoY3foMiCZaKnB4f/v4PGRASaKD8dGwK/CSAcB4BfGkTWU0qCcXURmypeQZ0l9x2TzCqZau7PGLg6cCdeBuniwFS1FPQQy76/wmy1dD8wq4QmxoCuYDR9Zs84XQs4g5Wfrule1o4TX5Pd2cn7O7MIF+OpKYtySialoIpgedafjlUK6FfUXTtcohkCrVyRIqDu0ce2avu44ed5aWXb1VThJftjfYmsv97FH8L1s8O6EH1ZXgG747C5vU640l+s7KEj119hcAAP//NUH8pQ=="
}
