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

package http

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "http", asset.ModuleFieldsPri, AssetHttp); err != nil {
		panic(err)
	}
}

// AssetHttp returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/http.
func AssetHttp() string {
	return "eJzMlDFvwjAQhff8iqfMRerAlKFzp6oDW9XBxA8wJHHqu9Dm31cBQklqkCit1Btt53ufz7En2LDNsFKtE0CdFsyQPs5mz2kCWEoeXK3OVxkeEgDoplB62xRMgMCCRphhaRJAqOqqpWR4SUWK9A5pB05fE2DhWFjJdowJKlPymNqVtnVHCb7pRyLZQ8opKfCtoehxPAY8C+1rt7UDCa5a+FCabuXJsnH+YDc0lkFG0L2Hn6+Z62jqogwwW3GvdABDWGk0ee5tG43dsH33wf4wtzZt4Y0d5n51XGpfCX+l5XvUf+x5YE63pY2m597yD/ouarSRGJ0fpqy7Gzq9n0aN6lUwco3TkfjkFQvfVDdKjwR6r7UMjvTqP6X7HiU1uFx4eqbDB6ivc8+EMGwZbhH5Rrgs8BkAAP//p5Fb6Q=="
}
