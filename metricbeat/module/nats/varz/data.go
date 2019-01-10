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

package varz

import (
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/elastic/beats/libbeat/common"
	s "github.com/elastic/beats/libbeat/common/schema"
	c "github.com/elastic/beats/libbeat/common/schema/mapstriface"
)

var (
	httpReqStatsSchema = s.Schema{
		"root_uri":   c.Int("/"),
		"connz_uri":  c.Int("/connz"),
		"routez_uri": c.Int("/routez"),
		"subsz_uri":  c.Int("/subsz"),
		"varz_uri":   c.Int("/varz"),
	}
	varzSchema = s.Schema{
		"server_id":         c.Str("server_id"),
		"now":               c.Str("now"),
		"uptime":            c.Str("uptime"),
		"mem":               c.Int("mem"),
		"cores":             c.Int("cores"),
		"cpu":               c.Int("cpu"),
		"total_connections": c.Int("total_connections"),
		"remotes":           c.Int("remotes"),
		"in.msgs":           c.Int("in_msgs"),
		"out.msgs":          c.Int("out_msgs"),
		"in.bytes":          c.Int("in_bytes"),
		"out.bytes":         c.Int("out_bytes"),
		"slow_consumers":    c.Int("slow_consumers"),
		"http_req_stats":    c.Dict("http_req_stats", httpReqStatsSchema),
	}
)

func eventMapping(content []byte) (common.MapStr, error) {
	var event common.MapStr
	var inInterface map[string]interface{}

	err := json.Unmarshal(content, &inInterface)
	if err != nil {
		err = errors.Wrap(err, "failure parsing Nats varz API response")
		return event, err
	}
	event, err = varzSchema.Apply(inInterface)
	if err != nil {
		err = errors.Wrap(err, "failure applying index schema")
		return event, err
	}
	// TODO: convert uptime field here to long (secs)

	d, _ := event.GetValue("http_req_stats")
	httpStats := d.(common.MapStr)
	event.Delete("http_req_stats")
	event["http_req_stats.root_uri"] = httpStats["root_uri"]
	event["http_req_stats.connz_uri"] = httpStats["connz_uri"]
	event["http_req_stats.routez_uri"] = httpStats["routez_uri"]
	event["http_req_stats.subsz_uri"] = httpStats["subsz_uri"]
	event["http_req_stats.varz_uri"] = httpStats["varz_uri"]

	return event, nil
}
