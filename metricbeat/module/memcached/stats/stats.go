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

package stats

import (
	"bufio"
	"net"
	"strings"

	"github.com/elastic/beats/libbeat/logp"

	"github.com/elastic/beats/metricbeat/mb"
)

var logger = logp.NewLogger("memcached.stats")

func init() {
	mb.Registry.MustAddMetricSet("memcached", "stats", New,
		mb.DefaultMetricSet(),
	)
}

type MetricSet struct {
	mb.BaseMetricSet
}

func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	return &MetricSet{
		BaseMetricSet: base,
	}, nil
}

// Fetch methods implements the data gathering and data conversion to the right
// format. It publishes the event which is then forwarded to the output. In case
// of an error set the Error field of mb.Event or simply call report.Error().
func (m *MetricSet) Fetch(reporter mb.ReporterV2) {
	conn, err := net.DialTimeout("tcp", m.Host(), m.Module().Config().Timeout)
	if err != nil {
		logger.Error(err)
		reporter.Error(err)
		return
	}
	defer conn.Close()

	_, err = conn.Write([]byte("stats\n"))
	if err != nil {
		logger.Error(err)
		reporter.Error(err)
		return
	}

	scanner := bufio.NewScanner(conn)

	data := map[string]interface{}{}

	for scanner.Scan() {
		text := scanner.Text()
		if text == "END" {
			break
		}

		// Split entries which look like: STAT time 1488291730
		entries := strings.Split(text, " ")
		if len(entries) == 3 {
			data[entries[1]] = entries[2]
		}
	}

	event, _ := schema.Apply(data)

	reporter.Event(mb.Event{MetricSetFields: event})

	return
}
