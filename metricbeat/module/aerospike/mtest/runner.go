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

// +build integration

package mtest

import (
	"github.com/elastic/beats/libbeat/tests/compose"
)

var (
	// Runner is the compose test runner for aerospike
	Runner = compose.TestRunner{
		Service: "aerospike",
		Options: compose.RunnerOptions{
			"AEROSPIKE_VERSION": {
				"3.9.0",
				"3.13.0.11",
				"3.16.0.6",
				"4.3.0.2",
			},
		},
		Parallel: true,
	}
)
