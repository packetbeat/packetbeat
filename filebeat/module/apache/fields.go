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

package apache

import (
	"github.com/menderesk/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "apache", asset.ModuleFieldsPri, AssetApache); err != nil {
		panic(err)
	}
}

// AssetApache returns asset data.
// This is the base64 encoded zlib format compressed contents of module/apache.
func AssetApache() string {
	return "eJysksFq6zAQRff+ikv28Qdo8eBRKF20UEj2xUgTWUTWiJGckr8vcuTUadNAaWY5su455mqNPR0Vutjpnhogu+xJYfV/WqwawFDS4mJ2HBT+NQBwOsQLm9GXS6lnyW+aw85ZhSxjWe4ceZPUdGGN0A20wJTJx0gKVniMdXMFtcCdAtu6XsZfILSmlM7ra5gbqDIPHHLnQqoI7FiQe5o1nrbbV2xIDiQVBs/27HXNbemXkm+jcGbN/uKD2XRPx3cW8+Xshm+ZzeYZcyoOJMlx+DT6UUS72JPcV+OUOUHab+2QCMtfynmsnQgPy1Km3F/1MMyP906/vu2pZpYizKhdsJOhZ2vJYKCUOktt8xEAAP//tzvrEQ=="
}
