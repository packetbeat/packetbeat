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

package monitors

import (
	"fmt"
	"testing"
	"time"

	"github.com/elastic/beats/heartbeat/eventext"
	"github.com/elastic/beats/heartbeat/monitors/jobs"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/common/mapval"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/elastic/beats/heartbeat/scheduler"
	"github.com/elastic/beats/libbeat/testing/mapvaltest"
)

func TestMonitor(t *testing.T) {
	serverMonConf := mockPluginConf(t, "", "@every 1ms", "http://example.net")
	reg := mockPluginsReg()
	pipelineConnector := &MockPipelineConnector{}

	sched := scheduler.New(1)
	err := sched.Start()
	require.NoError(t, err)
	defer sched.Stop()

	mon, err := newMonitor(serverMonConf, reg, pipelineConnector, sched, false)
	require.NoError(t, err)

	mon.Start()

	require.Equal(t, 1, len(pipelineConnector.clients))
	pcClient := pipelineConnector.clients[0]

	timeout := time.Second
	start := time.Now()
	success := false
	for time.Since(start) < timeout && !success {
		count := len(pcClient.Publishes())
		if count >= 1 {
			success = true

			mon.Stop()
			pcClient.Close()

			for _, event := range pcClient.Publishes() {
				mapvaltest.Test(t, mockEventMonitorValidator(""), event.Fields)
			}
		} else {
			// Let's yield this goroutine so we don't spin
			// This could (possibly?) lock on a single core system otherwise
			time.Sleep(time.Microsecond)
		}
	}

	if !success {
		t.Fatalf("No publishes detected!")
	}

	mon.Stop()
	assert.Equal(t, true, pcClient.closed)
}

func TestDuplicateMonitorIDs(t *testing.T) {
	serverMonConf := mockPluginConf(t, "custom", "@every 1ms", "http://example.net")
	reg := mockPluginsReg()
	pipelineConnector := &MockPipelineConnector{}

	sched := scheduler.New(1)
	err := sched.Start()
	require.NoError(t, err)
	defer sched.Stop()

	makeTestMon := func() (*Monitor, error) {
		return newMonitor(serverMonConf, reg, pipelineConnector, sched, false)
	}

	m1, m1Err := makeTestMon()
	assert.NoError(t, m1Err)
	_, m2Err := makeTestMon()
	assert.Error(t, m2Err)

	m1.Stop()
	_, m3Err := makeTestMon()
	assert.NoError(t, m3Err)
}

func TestMonitor_wrapCommon(t *testing.T) {
	var simpleJob jobs.Job = func(event *beat.Event) ([]jobs.Job, error) {
		eventext.MergeEventFields(event, common.MapStr{"simple": "job"})
		return nil, nil
	}
	simpleJobValidator := mapval.MustCompile(mapval.Map{"simple": "job"})

	var errorJob jobs.Job = func(event *beat.Event) ([]jobs.Job, error) {
		return nil, fmt.Errorf("myerror")
	}
	errorJobValidator := mapval.MustCompile(mapval.Map{
		"error": mapval.Map{
			"message": "myerror",
			"type":    "io",
		},
	})

	type fields struct {
		id   string
		name string
		typ  string
	}

	commonFieldsValidator := func(f fields, status string) mapval.Validator {
		return mapval.MustCompile(mapval.Map{
			"monitor": mapval.Map{
				"duration.us": mapval.IsDuration,
				"id":          f.id,
				"name":        f.name,
				"type":        f.typ,
				"status":      status,
			},
		})
	}

	testFields := fields{"myid", "myname", "mytyp"}

	tests := []struct {
		name   string
		fields fields
		jobs   []jobs.Job
		want   []mapval.Validator
	}{
		{
			"simple",
			testFields,
			[]jobs.Job{simpleJob},
			[]mapval.Validator{
				mapval.Strict(mapval.Compose(
					simpleJobValidator,
					commonFieldsValidator(testFields, "up"),
				)),
			},
		},
		{
			"job error",
			testFields,
			[]jobs.Job{errorJob},
			[]mapval.Validator{
				mapval.Strict(mapval.Compose(
					errorJobValidator,
					commonFieldsValidator(testFields, "down"),
				)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Monitor{
				id:   tt.fields.id,
				name: tt.fields.name,
				typ:  tt.fields.typ,
			}
			wrapped := jobs.WrapCommon(tt.jobs, m.id, m.name, m.typ)

			defer m.freeID()
			results, err := jobs.ExecJobsAndConts(t, wrapped)
			assert.NoError(t, err)

			for idx, r := range results {
				mapvaltest.Test(t, tt.want[idx], r.Fields)
			}
		})
	}
}
