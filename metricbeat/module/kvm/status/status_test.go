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

package status

import (
	"testing"

	"github.com/digitalocean/go-libvirt/libvirttest"
	"github.com/stretchr/testify/assert"

	mbtest "github.com/menderesk/beats/v7/metricbeat/mb/testing"
)

func TestFetchEventContents(t *testing.T) {
	conn := libvirttest.New()
	defer conn.Close()

	host := "test://" + conn.RemoteAddr().String() + ":123"
	f := mbtest.NewReportingMetricSetV2Error(t, getConfig(host))

	events, errs := mbtest.ReportingFetchV2Error(f)
	if len(errs) > 0 {
		t.Fatal(errs)
	}
	if len(events) == 0 {
		t.Fatal("no events received")
	}

	for _, e := range events {
		if e.Error != nil {
			t.Fatalf("received error: %+v", e.Error)
		}
	}
	if len(events) == 0 {
		t.Fatal("received no events")
	}

	e := events[0]

	t.Logf("%s/%s event: %+v", f.Module().Name(), f.Name(), e)

	statName, err := e.MetricSetFields.GetValue("state")
	if err == nil {
		assert.EqualValues(t, statName.(string), "running")
	} else {
		t.Errorf("error while getting value from event: %v", err)
	}
}

func getConfig(host string) map[string]interface{} {
	return map[string]interface{}{
		"module":     "kvm",
		"metricsets": []string{"status"},
		"hosts":      []string{host},
	}
}
