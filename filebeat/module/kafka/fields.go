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

package kafka

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "kafka", asset.ModuleFieldsPri, AssetKafka); err != nil {
		panic(err)
	}
}

// AssetKafka returns asset data.
// This is the base64 encoded gzipped contents of module/kafka.
func AssetKafka() string {
	return "eJykkjtuwzAQRHudYuDeOoCKNOkSpMsFFtZKJsQfyLUT3z6gGP8Y2paR7cQRZ55Gu8bEhw4TDRM1gCjR3GH1np5XDdBz3ATlRTnb4aUBgFmDcf1OcwMMinUfu1law5Lhs10aOXjuMAa3878nFc9rm0sr7cbTWc3spmGeDKvdCK0sx/ZCLBOvUnnP+ko5ZpNWFAvFk2xn0rZ2z6gxUGaTsONqnuEYaeQnE+u3buVVgzfOeGfZSjV64sOXC32h3ak7zevRErLluXoVU46yI4bgTFsH0RTLr/wHxBvtKXs+RSGBNvW/UO7cAgbgM9lB2RNDWsK2eLG2h4+aedTOIro/LUnGvdMTHu7smUz4u9yqhVjAR/aGpyBww5mubX4CAAD//40BKYM="
}
