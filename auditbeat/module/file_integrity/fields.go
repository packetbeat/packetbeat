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

package file_integrity

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("auditbeat", "file_integrity", asset.ModuleFieldsPri, AssetFileIntegrity); err != nil {
		panic(err)
	}
}

// AssetFileIntegrity returns asset data.
// This is the base64 encoded gzipped contents of module/file_integrity.
func AssetFileIntegrity() string {
	return "eJykk81u6jAQhfd5inkBuIohEcriSqnaKlXbFV2wQw4eYgsnRrYD5O2ruCk/ElJsusxo8p3jmTMT2GGXwVZIXIvGYqWF7SIAK6zEDF6FRHi7qjM0Gy32Vqgmgy+OBoFqBMsRtgIlM1Bhg5paZFB2Q/2aDbVircRpBMMPWQQwgYbWmAGnhkcAALbbYwaVVu3efd/I/nclgIIajgbU9iwz7S31LzLOFZWV0sLy2uEN0Ia51gOVLbqWgdQXOZ4Am41iyICJCo0d+qaR67q4vfgtJd0hKdckSX9JzvgOu6PSbKjdmH/6yN9fSDkhSeqee2M/ukufLeah9Nli7ktPYhJKT2IyRq9ZMhDUAfVRC4sZWN2ir9bnczKmYTiN/yayLPLYQ4UQ7wUsi5yQ0dn3zHNkHvfukSHDaUB8lkXukZyeuQ6biev34/qf0g/XbwYhR+S4vnMIOB/H9bgdw+mF+ng6vJUCd5nE5J/fNh07aJ+OPb7R04mn3pZXq/Se2e8AAAD//yR/DPc="
}
