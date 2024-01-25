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

//go:build integration

package logstash_test

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/elastic/beats/v7/libbeat/tests/compose"
	"github.com/elastic/beats/v7/metricbeat/mb"
	mbtest "github.com/elastic/beats/v7/metricbeat/mb/testing"
	"github.com/elastic/beats/v7/metricbeat/module/logstash"
	_ "github.com/elastic/beats/v7/metricbeat/module/logstash/node"
	_ "github.com/elastic/beats/v7/metricbeat/module/logstash/node_stats"
)

var metricSets = []string{
	"node",
	"node_stats",
}

func TestFetch(t *testing.T) {
	service := compose.EnsureUpWithTimeout(t, 300, "logstash")

	for _, metricSet := range metricSets {
		t.Run(metricSet, func(t *testing.T) {
			config := getConfig(metricSet, service.Host())
			f := mbtest.NewReportingMetricSetV2Error(t, config)
			var (
				events []mb.Event
				errs   []error
			)

			require.Eventually(t, func() bool {
				events, errs = mbtest.ReportingFetchV2Error(f)
				return len(events) > 0 && len(errs) == 0
			}, 30*time.Second, time.Second, "events cannot be empty: %v, %v", events, errs)

			require.Empty(t, errs)
			require.NotEmpty(t, events)

			t.Logf("%s/%s event: %+v", f.Module().Name(), f.Name(),
				events[0].BeatEvent("logstash", metricSet).Fields.StringToPrint())
		})
	}
}

func TestData(t *testing.T) {
	service := compose.EnsureUpWithTimeout(t, 300, "logstash")

	for _, metricSet := range metricSets {
		t.Run(metricSet, func(t *testing.T) {
			config := getConfig(metricSet, service.Host())
			f := mbtest.NewReportingMetricSetV2Error(t, config)
			err := mbtest.WriteEventsReporterV2Error(f, t, metricSet)
			require.NoError(t, err)
		})
	}
}

func getConfig(metricSet string, host string) map[string]interface{} {
	return map[string]interface{}{
		"module":     logstash.ModuleName,
		"metricsets": []string{metricSet},
		"hosts":      []string{host},
	}
}

func getXPackConfig(host string) map[string]interface{} {
	return map[string]interface{}{
		"module":     logstash.ModuleName,
		"metricsets": metricSets,
		"hosts":      []string{host},
	}
}

func getESClusterUUID(t *testing.T, host string) string {
	resp, err := http.Get("http://" + host + "/")
	require.NoError(t, err)
	defer resp.Body.Close()

	var body struct {
		ClusterUUID string `json:"cluster_uuid"`
	}

	data, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	json.Unmarshal(data, &body)

	return body.ClusterUUID
}
