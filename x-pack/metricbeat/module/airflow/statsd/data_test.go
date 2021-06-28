// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package statsd

import (
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/beats/v7/auditbeat/core"
	_ "github.com/elastic/beats/v7/libbeat/processors/actions"
	"github.com/elastic/beats/v7/metricbeat/mb"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/statsd/server"

	mbtest "github.com/elastic/beats/v7/metricbeat/mb/testing"
)

func init() {
	mb.Registry.SetSecondarySource(mb.NewLightModulesSource("../../../module"))
}

const (
	STATSD_HOST = "localhost"
	STATSD_PORT = 8126
)

func getConfig() map[string]interface{} {
	return map[string]interface{}{
		"module":     "airflow",
		"metricsets": []string{"statsd"},
		"host":       STATSD_HOST,
		"port":       STATSD_PORT,
		"period":     "100ms",
	}
}

func createEvent(t *testing.T) {
	udpAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", STATSD_HOST, STATSD_PORT))
	assert.NoError(t, err)

	conn, err := net.DialUDP("udp", nil, udpAddr)
	assert.NoError(t, err)

	_, err = fmt.Fprint(conn, "dagrun.duration.failed.a_dagid:200|ms|#k1:v1,k2:v2")
	assert.NoError(t, err)
}

func TestData(t *testing.T) {
	ms := mbtest.NewPushMetricSetV2(t, getConfig())
	var events []mb.Event
	done := make(chan interface{})
	go func() {
		events = mbtest.RunPushMetricSetV2(10*time.Second, 1, ms)
		close(done)
	}()

	createEvent(t)
	<-done

	if len(events) == 0 {
		t.Fatal("received no events")
	}

	beatEvent := mbtest.StandardizeEvent(ms, events[0], core.AddDatasetToEvent)
	mbtest.WriteEventToDataJSON(t, beatEvent, "")
}
