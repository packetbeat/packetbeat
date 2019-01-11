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

	"strconv"
	"strings"

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
		"in": s.Object{
			"messages": c.Int("in_msgs"),
			"bytes":    c.Int("in_bytes"),
		},
		"out": s.Object{
			"messages": c.Int("out_msgs"),
			"bytes":    c.Int("out_bytes"),
		},
		"slow_consumers": c.Int("slow_consumers"),
		"http_req_stats": c.Dict("http_req_stats", httpReqStatsSchema),
	}
)

// Converts uptime from formatted string to seconds
// input: "1y20d22h3m30s", output: 33343410
func convertUptime(uptime string) (seconds int64, err error) {

	var split []string
	var years, days, hours, minutes, secs int64
	if strings.Contains(uptime, "y") {
		split = strings.Split(uptime, "y")
		uptime = split[1]
		years, err = strconv.ParseInt(split[0], 10, 64)
		if err != nil {
			return
		}
		seconds += years * 31536000
	}

	if strings.Contains(uptime, "d") {
		split = strings.Split(uptime, "d")
		uptime = split[1]
		days, err = strconv.ParseInt(split[0], 10, 64)
		if err != nil {
			return
		}
		seconds += days * 86400
	}

	if strings.Contains(uptime, "h") {
		split = strings.Split(uptime, "h")
		uptime = split[1]
		hours, err = strconv.ParseInt(split[0], 10, 64)
		if err != nil {
			return
		}
		seconds += hours * 3600
	}

	if strings.Contains(uptime, "m") {
		split = strings.Split(uptime, "m")
		uptime = split[1]
		minutes, err = strconv.ParseInt(split[0], 10, 64)
		if err != nil {
			return
		}
		seconds += minutes * 60
	}

	if strings.Contains(uptime, "s") {
		split = strings.Split(uptime, "s")
		uptime = split[1]
		secs, err = strconv.ParseInt(split[0], 10, 64)
		if err != nil {
			return
		}
		seconds += secs
	}
	return
}

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

	uptime, err := event.GetValue("uptime")
	if err != nil {
		err = errors.Wrap(err, "failure retrieving uptime key")
		return event, err
	}
	uptime, err = convertUptime(uptime.(string))
	if err != nil {
		err = errors.Wrap(err, "failure converting uptime from string to integer")
		return event, err
	}
	_, err = event.Put("uptime", uptime)
	if err != nil {
		err = errors.Wrap(err, "failure updating uptime key")
		return event, err
	}

	d, err := event.GetValue("http_req_stats")
	if err != nil {
		err = errors.Wrap(err, "failure retrieving http_req_stats key")
		return event, err
	}
	httpStats, ok := d.(common.MapStr)
	if !ok {
		err = errors.Wrap(err, "failure casting http_req_stats to common.Mapstr")
		return event, err
	}
	err = event.Delete("http_req_stats")
	if err != nil {
		err = errors.Wrap(err, "failure deleting http_req_stats key")
		return event, err
	}
	event["http"] = common.MapStr{
		"req_stats": common.MapStr{
			"uri": common.MapStr{
				"root":   httpStats["root_uri"],
				"connz":  httpStats["connz_uri"],
				"routez": httpStats["routez_uri"],
				"subsz":  httpStats["subsz_uri"],
				"varz":   httpStats["varz_uri"],
			},
		},
	}
	return event, nil
}
