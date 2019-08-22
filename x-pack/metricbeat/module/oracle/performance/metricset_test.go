// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// +build integration

package performance

import (
	"testing"

	"github.com/elastic/beats/x-pack/metricbeat/module/oracle"

	"github.com/stretchr/testify/assert"

	_ "gopkg.in/goracle.v2"

	mbtest "github.com/elastic/beats/metricbeat/mb/testing"
)

func TestData(t *testing.T) {
	//t.Skip("Skip until a proper Docker image is setup for Metricbeat")
	f := mbtest.NewReportingMetricSetV2WithContext(t, getConfig())

	events, errs := mbtest.ReportingFetchV2WithContext(f)
	if len(errs) > 0 {
		t.Fatalf("Expected 0 error, had %d. %v\n", len(errs), errs)
	}
	assert.NotEmpty(t, events)

	if err := mbtest.WriteEventsReporterV2WithContext(f, t, ""); err != nil {
		t.Fatal("write", err)
	}
}

func getConfig() map[string]interface{} {
	return map[string]interface{}{
		"module":     "oracle",
		"metricsets": []string{"performance"},
		"hosts":      []string{oracle.GetOracleConnectionDetails("localhost")},
		"username":   "sys",
		"password":   "Oradoc_db1",
	}
}
