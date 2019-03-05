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

package subscriptions

import (
	"testing"

	"github.com/elastic/beats/libbeat/tests/compose"
	mbtest "github.com/elastic/beats/metricbeat/mb/testing"
)

func TestSubscriptions(t *testing.T) {
	runner := compose.TestRunner{Service: "nats"}
	runner.Run(t, compose.Suite{
		"Data":  testData,
		"Fetch": testFetch,
	})
}

func testData(t *testing.T, r compose.R) {
	metricSet := mbtest.NewReportingMetricSetV2(t, getConfig(r.Host()))
	err := mbtest.WriteEventsReporterV2(metricSet, t, "./test_data.json")
	if err != nil {
		t.Fatal("write", err)
	}
}

func testFetch(t *testing.T, r compose.R) {
	reporter := &mbtest.CapturingReporterV2{}

	metricSet := mbtest.NewReportingMetricSetV2(t, getConfig(r.Host()))
	metricSet.Fetch(reporter)

	e := mbtest.StandardizeEvent(metricSet, reporter.GetEvents()[0])
	t.Logf("%s/%s event: %+v", metricSet.Module().Name(), metricSet.Name(), e.Fields.StringToPrint())
}

func getConfig(host string) map[string]interface{} {
	return map[string]interface{}{
		"module":     "nats",
		"metricsets": []string{"subscriptions"},
		"hosts":      []string{host},
	}
}
