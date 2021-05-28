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

//+build !netbsd

package service

import (
	"testing"
	"time"

	"github.com/coreos/go-systemd/v22/dbus"

	"github.com/elastic/beats/v7/libbeat/common"

	//mbtest "github.com/elastic/beats/v7/metricbeat/mb/testing"
	"github.com/stretchr/testify/assert"
)

var exampleUnits = []dbus.UnitStatus{
	{
		Name:      "sshd.service",
		LoadState: "active",
	},
	{
		Name:      "metricbeat.service",
		LoadState: "active",
	},
	{
		Name: "filebeat.service",
	},
}

func TestFormProps(t *testing.T) {
	testUnit := dbus.UnitStatus{
		Name:        "test.service",
		LoadState:   "loaded",
		ActiveState: "active",
		SubState:    "running",
	}
	testprops := Properties{
		ExecMainPID:          0,
		ExecMainStatus:       0,
		ExecMainCode:         1,
		ActiveEnterTimestamp: 1571850129000000,
		IPAccounting:         true,
		IPEgressBytes:        100,
		IPIngressBytes:       50,
		IPEgressPackets:      100,
		IPIngressPackets:     50,
	}
	event, err := formProperties(testUnit, testprops)
	assert.NoError(t, err)

	testEvent := common.MapStr{
		"state":       "active",
		"exec_code":   "exited",
		"load_state":  "loaded",
		"name":        "test.service",
		"state_since": time.Unix(0, 1571850129000000*1000),
		"sub_state":   "running",
		"resources": common.MapStr{"network": common.MapStr{
			"in": common.MapStr{
				"bytes":   50,
				"packets": 50},
			"out": common.MapStr{
				"bytes":   100,
				"packets": 100},
		},
		},
	}

	assert.NotEmpty(t, event.MetricSetFields["resources"])
	assert.Equal(t, event.MetricSetFields["state_since"], testEvent["state_since"])
	assert.NotEmpty(t, event.RootFields)
}

func TestFilterEmpty(t *testing.T) {

	filtersBad := []string{
		"asdf",
	}
	shouldNotMatch, err := matchUnitPatterns(filtersBad, exampleUnits)
	assert.NoError(t, err)
	assert.Empty(t, shouldNotMatch)
}

func TestFilterMatches(t *testing.T) {
	filtersMatch := []string{
		"ssh*",
	}

	shouldMatch, err := matchUnitPatterns(filtersMatch, exampleUnits)
	assert.NoError(t, err)
	assert.Len(t, shouldMatch, 1)
}

func TestNoFilter(t *testing.T) {
	shouldReturnResults, err := matchUnitPatterns([]string{}, exampleUnits)
	assert.NoError(t, err)
	assert.Len(t, shouldReturnResults, 3)
}

func TestUnitStateFilter(t *testing.T) {
	stateFilter := []string{
		"active",
	}
	shouldReturnResults := matchUnitState(stateFilter, exampleUnits)
	assert.Len(t, shouldReturnResults, 2)

}

func TestUnitStateNoFilter(t *testing.T) {
	shouldReturnResults := matchUnitState([]string{}, exampleUnits)
	assert.Len(t, shouldReturnResults, 3)
}

// func TestFetch(t *testing.T) {
// 	f := mbtest.NewReportingMetricSetV2Error(t, getConfig())
// 	events, errs := mbtest.ReportingFetchV2Error(f)

// 	assert.Empty(t, errs)
// 	if !assert.NotEmpty(t, events) {
// 		t.FailNow()
// 	}
// 	t.Logf("Events: %d", len(events))

// 	for _, unit := range events {
// 		t.Logf("Unit: %s State: %s, SubState: %#v", unit.MetricSetFields["name"], unit.MetricSetFields["state"], unit.MetricSetFields["sub_state"])
// 	}
// }

// func getConfig() map[string]interface{} {
// 	return map[string]interface{}{
// 		"module":                     "system",
// 		"metricsets":                 []string{"service"},
// 		"service.track_non_service":  false,
// 		"service.track_not_found":    false,
// 		"service.track_instantiated": true,
// 	}
// }
