// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

//go:build !integration
// +build !integration

package qmgr

import (
	"os"
	"testing"

	"github.com/elastic/beats/v8/libbeat/logp"
	"github.com/elastic/beats/v8/metricbeat/mb"
	mbtest "github.com/elastic/beats/v8/metricbeat/mb/testing"

	// Register input module and metricset
	_ "github.com/elastic/beats/v8/metricbeat/module/prometheus"
	_ "github.com/elastic/beats/v8/metricbeat/module/prometheus/collector"
)

func init() {
	// To be moved to some kind of helper
	os.Setenv("BEAT_STRICT_PERMS", "false")
	mb.Registry.SetSecondarySource(mb.NewLightModulesSource("../../../module"))
}

func TestEventMapping(t *testing.T) {
	logp.TestingSetup()

	mbtest.TestDataFiles(t, "ibmmq", "qmgr")
}
