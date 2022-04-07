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

package prometheus

import (
	"github.com/elastic/beats/v8/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "prometheus", asset.ModuleFieldsPri, AssetPrometheus); err != nil {
		panic(err)
	}
}

// AssetPrometheus returns asset data.
// This is the base64 encoded zlib format compressed contents of module/prometheus.
func AssetPrometheus() string {
	return "eJzMkk1u20AMhfc6xYO6C5wcQIueoEBbdFkUxlh6sqaZv5JUDN++kGU5im2gf5tyyTckP77hI555bFAkR9rAUSvAvAU2qD9dknUFdNRWfDGfU4P3FQB8MWcKbcUVduglRzi8VoGpK9kne6oAHbLYts2p9/sGvQvKChAGOmWDvZve0MynvTb4WquGeoN6MCv1twroPUOnzWnuI5KLvKKewo5l6iV5LOfMumyKd/goHQVe4WPJYi4ZBgo3CG7HoDj4EBCdtQN6L2ob2EAI1eCE6PK4C7z0W1Dm4qeHi7DA5N13trZKz4ntrD7zeMjSreQ7Ni+xcjbSxLfnqTcws/rnNFe7vVG30ZXi0/78tH6o/xL6hvbHSDn+b6wvLoynXx+DLbc9yZ8/rPjfXu+drW52Wp/mL2BODZav5NqHe2NvL30BEcZs3B7EG/+FZ+6DU58F69WXs21KeaH8NuvPAAAA///rUkpn"
}
