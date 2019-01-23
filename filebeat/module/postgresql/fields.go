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

package postgresql

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "postgresql", asset.ModuleFieldsPri, AssetPostgresql); err != nil {
		panic(err)
	}
}

<<<<<<< HEAD
// Asset returns asset data
func Asset() string {
	return "eJyck0FP3DAQhe/5FU97rER+QCr1guAEbSnckdlMHKu2x3gc1O2vr7xhtcGbRKRztPXe9/ImvsJvOjQILElHkldbAckkSw12P8fDx4e7XQW0JPtoQjLsG3yrAOCe28ESOo4IKorxGqknnHWwrNEZS1JXgPQc0/OefWd0gxQHqoDOkG2lOfpdwStHRZo86RCogY48hPeTmTTj3B790EV2RZBjhjxT5BRrWU+MLpmr3I/ktRpOU8aYRknGkSTlwofbVXyep57O0nOSjLfGUz3L2nOkZ9MWZuPnW/Z6W4RrjoTC7URqVVIvSqgQ0R/lwvGPc4f2ZbfCu6B9V47A3bzzCfs6UDwsMh9v7m6un/AFt79+3GMQivJ1U4aHbA9JKpEjn84lL272L/uyg7FuZY2S4iao1Degt2y9IHZGRzUu5P1VzXD7SKpd2vMyOETek0gdLpSfoeY6NwKzpM7q/8BZeiO7kWdZ13O6z/AciSi9dZfzqpL3LwAA//+FnHBP"
=======
// AssetPostgresql returns asset data.
// This is the base64 encoded gzipped contents of module/postgresql.
func AssetPostgresql() string {
	return "eJyck0FP3DAQhe/5FU97rESE1Fsq9ULhBG0p3NGwmThW7YzxOKjbX185AXbJJivSOcZ67/sycc7wm3cVgmgykfXJFUCyyXGFzc/x4d3t9aYAatZttCFZ6Sp8LQDgRureMRqJCBTVdgapZexzcGLQWMdaFoC2EtPDVrrGmgop9lwAjWVXazX0naEjzxObPGkXuIKJ0oeXJzM241wNfWii+InI4JDnEHmIdWIOio6ZJ7nvyafW8DpTjUOVZD1rIh/enZ7E57lveR/dm2S8sx2Xs6ytRH6w9aRsfH0nnVmncCGRMWl7JdWU6JGUJyH+Qz4MN87v6sfNCd4R7Tt5hjTzzW/YPlKOz75h44TSotDn8/J8ldC3F1aWIjz1HHfzWx+OFrl3l9eXF/f4hKtfP27QK0f9ssrjNtdDEyX23KW9w+J1+yvd9MOMGyJnSScngVJbgZ9z9ULYWzOu4u1Xn+G2kaleunzL4BBly6plOEp+hJrXuRKYI2VO/wfO8TO7lTwnppzLfYTnWZXM2m85n5ry/gUAAP//D0KTig=="
>>>>>>> update fields.go
}
