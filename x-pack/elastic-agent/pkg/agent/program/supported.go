// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by x-pack/elastic-agent/dev-tools/cmd/buildspec/buildspec.go - DO NOT EDIT.

package program

import (
	"strings"

	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/packer"
)

var Supported []Spec
var SupportedMap map[string]Spec

func init() {
	// Packed Files
	// spec/apm-server.yml
	// spec/endpoint.yml
	// spec/filebeat.yml
	// spec/fleet-server.yml
	// spec/heartbeat.yml
	// spec/metricbeat.yml
	// spec/osquerybeat.yml
	// spec/packetbeat.yml
	unpacked := packer.MustUnpack("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")
	SupportedMap = make(map[string]Spec)

	for f, v := range unpacked {
		s, err := NewSpecFromBytes(v)
		if err != nil {
			panic("Cannot read spec from " + f)
		}
		Supported = append(Supported, s)
		SupportedMap[strings.ToLower(s.Cmd)] = s
	}
}
