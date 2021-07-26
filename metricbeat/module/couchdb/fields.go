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

package couchdb

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "couchdb", asset.ModuleFieldsPri, AssetCouchdb); err != nil {
		panic(err)
	}
}

// AssetCouchdb returns asset data.
// This is the base64 encoded zlib format compressed contents of module/couchdb.
func AssetCouchdb() string {
	return "eJy8mFtvwjYUx9/5FEd5XqVweygPlVpK22m9oJVqmqYpMvaBWE3szD4pYp9+ciCQFChzqySPSXz+P/3POb5dwDuuR8B1zmMx7wCQpARHEIzdm9uboAMg0HIjM5JajeCqAwDl/5BqkSfYATCYILM4giXrAFgkkmppR/BXYG0S/AJBTJQFf3cAFhITYUdFmAtQLMWqvHtonblARufZ9k2NYFxqIxnJ7faXathqaIvmA83u9bHoBwpXlQ8AY62ISWVha8k2JFhiZCt/1i0on89cVTbniah9OYV3BtE9D7PZtGCSlva2fIVRRfmQuIoMMvF54J4p0Wp55OMZrOJ5ztM5GtCLQgfqOieZ5nnyHhn8J0dLTWDtqZwSHCidBOOJREW2ZJNqGfGYqSU2S7mVhYU2wLUiqXKdWzjQPslNmGbaMLOOGk33nngn6JX3VlJetMvXKa+1aZnsKEWK9YFvP2vbbWw4Hvtc745fpn8275VT8eiRh8n1bfNQTsUDavryOmseyql4QN1OHiezSfNYGx0PsPtJC2bdT3y8mr61kb+3c0j1mcEturmNuBYHC8DPV/PcQhH4+0t70AvDoHnXemEIL7+BQZtpZf/PchT0wm4rYF0YG2SEwpOu1wpdD645x8wXr9+Kef2wC0/6AwVM0aRMoaJk7ck5aIVzAM+a4EkLuZCeVg5aaZBBGMINE/D7drH3I2wj2YOwC2+K5RRrI//1NrHfCmIf7rSZSyFQefI1UoeHgJtCvNO58jVw2IqBQ3gqtpkF5nWS6JV3pi9bAb10R+9FIrlnr3TbmLcH3R5MDXKthHTD4I7JxNPIYSvzzjAM4VdFaBRL4HVzdzExRptzrMdvZ+qU39jalHdH397QCEZszixGKyOp4TM3yRQtsJ0mrJiFzYFbnM+wzlBF5dBmQZ0UHEqdRHOzfMQZjzFKpW2azqmhIslZ0SuFLnzSPYlanrxdLhqgfES1pLig3B3EpbJS4O7GbyUp1rnbXfBY/oHz88y7Gm3hkqUsUc3zFBUVJepkYWF0Wqldr6KIZcMXMEdLoqb6dVtpGy1k0nDdOoXdCG32d8AxswVG578AAAD//2DQe0E="
}
