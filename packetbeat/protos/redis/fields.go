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

package redis

import (
	"github.com/menderesk/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("packetbeat", "redis", asset.ModuleFieldsPri, AssetRedis); err != nil {
		panic(err)
	}
}

// AssetRedis returns asset data.
// This is the base64 encoded zlib format compressed contents of protos/redis.
func AssetRedis() string {
	return "eJyMj71uwzAMhHc9xSFzkwfw0L1r0b1grHMs1JIMkg7gty/8kzZBM5STRPK7Ox7xxbmBMiYLgCcf2ODwvvwPAYi0VtPoqZYGrwEA1tnRRrapSy14ZXF0iUO0U8D+atbVI4pk/sov5fPIBhet07h37olHyictn1cZJv4Mn0a61UfPncJKoXbwnltktDVnKRGpQNBPWQqUEuU8EF3VLH4Kf0JQter/3N+emfViUNo0ODfnskm+wPtk2+1oa3FJxRb+QXJdRaaZXG6nMeI83zkZ9Uo9hfAdAAD//8E/imI="
}
