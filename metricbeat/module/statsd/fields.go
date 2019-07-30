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

package statsd

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "statsd", asset.ModuleFieldsPri, AssetStatsd); err != nil {
		panic(err)
	}
}

// AssetStatsd returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/statsd.
func AssetStatsd() string {
	return "eJx0kE1OxTAQg/c5hZXlE48DZMElWCL0lDamBNImykyRentEf6IKeNnl88gz9hWfXBxEvUowgEZNdLDPK7AGCJS+xqIxTw5PBgA2EWMOc6IBKhO90KGjegMIVeM0iMOLFUn2AfZdtdhXA7xFpiBu9bli8iOP7Y+XFQK6FDrk7oO97mj73DYl5Llb9/5SbqMvJU7DPmYvdp/5J8I5BrXGvtF2FOsXa8PHWUPNcznRO94/749D62nwJ3o08h0AAP//kvxwfA=="
}
