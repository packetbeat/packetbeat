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

package kvm

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "kvm", asset.ModuleFieldsPri, AssetKvm); err != nil {
		panic(err)
	}
}

// AssetKvm returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/kvm.
func AssetKvm() string {
	return "eJyskLFuwyAQhnee4lf2vABDp659gqoDKtcIwXEWxm799lXcUCXOBXcogwcOfd/nOyLSYhFnNkANNZHFIc58MICn8b2EoQbJFk8GAOhroBKYcnXp9c2sd3FmsPgpkQE+AiU/2nVwRHZMDX4+dRnI4lRkGi43iuMWcg3ywkw8Vld/RxrzIfcy0ihb5a1487gn3tX/nBdiKYsO1luue85fZdyaIi2fUrz6YrdsU/fI1VJml6ZeS5J8+p8QzdQqwv3PdvS76mdhF/I9tbv//u7/6lzR3wEAAP//jl7YkQ=="
}
