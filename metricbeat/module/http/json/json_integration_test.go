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

package json

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/beats/libbeat/tests/compose"
	mbtest "github.com/elastic/beats/metricbeat/mb/testing"
)

func TestFetchObject(t *testing.T) {
	compose.EnsureUp(t, "http")

	f := mbtest.NewReportingMetricSetV2Error(t, getConfig("object"))
	events, errs := mbtest.ReportingFetchV2Error(f)
	if len(errs) > 0 {
		t.Fatalf("Expected 0 error, had %d. %v\n", len(errs), errs)
	}
	assert.NotEmpty(t, events)

	t.Logf("%s/%s event: %#v", f.Module().Name(), f.Name(), events)
}

func TestFetchArray(t *testing.T) {
	compose.EnsureUp(t, "http")

	f := mbtest.NewReportingMetricSetV2Error(t, getConfig("array"))
	events, errs := mbtest.ReportingFetchV2Error(f)
	if len(errs) > 0 {
		t.Fatalf("Expected 0 error, had %d. %v\n", len(errs), errs)
	}
	assert.NotEmpty(t, events)

	t.Logf("%s/%s event: %+v", f.Module().Name(), f.Name(), events[0])
}
func TestData(t *testing.T) {
	compose.EnsureUp(t, "http")

	f := mbtest.NewReportingMetricSetV2Error(t, getConfig("object"))
	if err := mbtest.WriteEventsReporterV2Error(f, t, ""); err != nil {
		t.Fatal("write", err)
	}
}

func getConfig(jsonType string) map[string]interface{} {
	var path string
	var responseIsArray bool
	switch jsonType {
	case "object":
		path = "/jsonobj"
		responseIsArray = false
	case "array":
		path = "/jsonarr"
		responseIsArray = true
	}

	return map[string]interface{}{
		"module":        "http",
		"metricsets":    []string{"json"},
		"hosts":         []string{getEnvHost() + ":" + getEnvPort()},
		"path":          path,
		"namespace":     "testnamespace",
		"json.is_array": responseIsArray,
	}
}

func getEnvHost() string {
	host := os.Getenv("HTTP_HOST")

	if len(host) == 0 {
		host = "127.0.0.1"
	}
	return host
}

func getEnvPort() string {
	port := os.Getenv("HTTP_PORT")

	if len(port) == 0 {
		port = "8080"
	}
	return port
}
