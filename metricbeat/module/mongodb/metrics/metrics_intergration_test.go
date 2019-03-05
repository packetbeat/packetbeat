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

package metrics

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/beats/libbeat/tests/compose"
	mbtest "github.com/elastic/beats/metricbeat/mb/testing"
	"github.com/elastic/beats/metricbeat/module/mongodb/mtest"
)

func TestMetrics(t *testing.T) {
	mtest.Runner.Run(t, compose.Suite{
		"Fetch": testFetch,
		"Data":  testData,
	})
}

func testFetch(t *testing.T, r compose.R) {
	f := mbtest.NewReportingMetricSetV2(t, mtest.GetConfig("metrics", r.Host()))
	events, errs := mbtest.ReportingFetchV2(f)
	if len(errs) > 0 {
		t.Fatalf("Expected 0 error, had %d. %v\n", len(errs), errs)
	}
	assert.NotEmpty(t, events)
	event := events[0].MetricSetFields

	// Check a few event Fields
	findCount, err := event.GetValue("commands.find.total")
	assert.NoError(t, err)
	assert.True(t, findCount.(int64) >= 0)

	deletedDocuments, err := event.GetValue("document.deleted")
	assert.NoError(t, err)
	assert.True(t, deletedDocuments.(int64) >= 0)
}

func testData(t *testing.T, r compose.R) {
	f := mbtest.NewReportingMetricSetV2(t, mtest.GetConfig("metrics", r.Host()))
	if err := mbtest.WriteEventsReporterV2(f, t, ""); err != nil {
		t.Fatal("write", err)
	}
}
