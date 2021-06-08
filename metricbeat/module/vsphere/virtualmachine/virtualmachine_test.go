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

package virtualmachine

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vmware/govmomi/simulator"

	"github.com/elastic/beats/v7/libbeat/common"
	mbtest "github.com/elastic/beats/v7/metricbeat/mb/testing"
)

func TestFetchEventContents(t *testing.T) {
	model := simulator.ESX()
	if err := model.Create(); err != nil {
		t.Fatal(err)
	}

	ts := model.Service.NewServer()
	defer ts.Close()

	f := mbtest.NewReportingMetricSetV2WithContext(t, getConfig(ts))
	events, errs := mbtest.ReportingFetchV2WithContext(f)
	if len(errs) > 0 {
		t.Fatalf("Expected 0 error, had %d. %v\n", len(errs), errs)
	}

	assert.NotEmpty(t, events)
	event := events[0].MetricSetFields

	t.Logf("%s/%s event: %+v", f.Module().Name(), f.Name(), event.StringToPrint())

	assert.EqualValues(t, "ha-host", event["host.id"])
	assert.EqualValues(t, "localhost.localdomain", event["host.hostname"])
	assert.True(t, strings.Contains(event["name"].(string), "ha-host_VM"))

	cpu := event["cpu"].(common.MapStr)

	cpuUsed := cpu["used"].(common.MapStr)
	assert.EqualValues(t, 0, cpuUsed["mhz"])
	cpuTotal := cpu["total"].(common.MapStr)
	assert.EqualValues(t, 0, cpuTotal["mhz"])
	cpuFree := cpu["free"].(common.MapStr)
	assert.EqualValues(t, 0, cpuFree["mhz"])

	memory := event["memory"].(common.MapStr)

	memoryUsed := memory["used"].(common.MapStr)
	memoryUsedHost := memoryUsed["host"].(common.MapStr)
	memoryUsedGuest := memoryUsed["guest"].(common.MapStr)
	assert.EqualValues(t, 0, memoryUsedGuest["bytes"])
	assert.EqualValues(t, 0, memoryUsedHost["bytes"])

	memoryTotal := memory["total"].(common.MapStr)
	memoryTotalGuest := memoryTotal["guest"].(common.MapStr)
	assert.EqualValues(t, uint64(33554432), memoryTotalGuest["bytes"])

	memoryFree := memory["free"].(common.MapStr)
	memoryFreeGuest := memoryFree["guest"].(common.MapStr)
	assert.EqualValues(t, uint64(33554432), memoryFreeGuest["bytes"])
}

func TestData(t *testing.T) {
	model := simulator.ESX()
	if err := model.Create(); err != nil {
		t.Fatal(err)
	}

	ts := model.Service.NewServer()
	defer ts.Close()

	f := mbtest.NewReportingMetricSetV2WithContext(t, getConfig(ts))

	if err := mbtest.WriteEventsReporterV2WithContext(f, t, ""); err != nil {
		t.Fatal("write", err)
	}
}

func getConfig(ts *simulator.Server) map[string]interface{} {
	urlSimulator := ts.URL.Scheme + "://" + ts.URL.Host + ts.URL.Path

	return map[string]interface{}{
		"module":     "vsphere",
		"metricsets": []string{"virtualmachine"},
		"hosts":      []string{urlSimulator},
		"username":   "user",
		"password":   "pass",
		"insecure":   true,
	}
}
