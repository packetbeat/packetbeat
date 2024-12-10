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

package lists

import (
	s "github.com/elastic/beats/v7/libbeat/common/schema"
	c "github.com/elastic/beats/v7/libbeat/common/schema/mapstrstr"
)

var schema = s.Schema{
	"databases":  c.Int("databases"),
	"users":      c.Int("users"),
	"peers":      c.Int("peers"),
	"pools":      c.Int("pools"),
	"peer_pools": c.Int("peer_pools"),
	"clients": s.Object{
		"free":  c.Int("free_clients"),
		"used":  c.Int("used_clients"),
		"login": c.Int("login_clients"),
	},
	"servers": s.Object{
		"free": c.Int("free_servers"),
		"used": c.Int("used_servers"),
	},
	"dns_names":   c.Int("dns_names"),
	"dns_zones":   c.Int("dns_zones"),
	"dns_queries": c.Int("dns_queries"),
}
