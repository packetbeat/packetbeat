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

package traefik

import (
	"github.com/elastic/beats/v8/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "traefik", asset.ModuleFieldsPri, AssetTraefik); err != nil {
		panic(err)
	}
}

// AssetTraefik returns asset data.
// This is the base64 encoded zlib format compressed contents of module/traefik.
func AssetTraefik() string {
	return "eJyslTFv4zwMhvf8CqJ7DRTfN2W4pUCBG245dDdUi3aEyKKPoq7wvz/IsVvHkFzHjbZY4fu8IinqEc7YH0FYYW3OBwAxYvEID6+XLw8HAI2+YtOJIXeEHwcAgF+kg0WoiaFT7I1rQE4IYxBYaqA2Fn1xAKgNWu2PQ9wjONXinBeX9B0eoWEK3fglgYzrZZCCmqnN8+KaM+dcVVXo/cfnFHoFH9czOVHG+RExpGBu5UKIjj7MpAzNTQWPXBqNTkxtkK/+Mzk8Y/9OrBd7Kz7j+ukHa79fnuHp/6f/4MKQHqgeNipr0EnSE+OfgF7KisLiH5MjS665zc7rCcGF9g05GhgJPomvmZyg02X8eb+EDA5Ui1MCJkwsgU4aeVPVOfoIbHfaSJoIbCcPIwHeT8g4ZQXM0FnvivXMWNoh6b706KR46wV90qWyRi13OiWnI5xEuoLRd+Q8FlErKdOahtUlq8IBMy3TkmBpuhsteApcYaG05uu7uRU83J9so+TBMa5IxG1htignWpZ9Y7KHChdJhU3HzbTiykHZFsSmMU4tQ7cAo+3yL7I35PacOB26raUujVlWpG+t7nVne1ESfEpnm48amTOjeWO9Mxpb8KpZjultzV0OgbeWPn/H8j6WDyhknry5ZOLmrR8sebikyoSglMSNAPJFHaxNDZg5KL2/h7YOStdyD2mplKx7g5QZ6HtKXpET49DJd7I1vhcNUvGF3ic2OOG+NJ5SE2AX+AvFCW2pGu7Y95FZpc8h1Rhyd8rsmthHWo309ypkVmpxuvuVMCd4+BcAAP//8XRZtA=="
}
