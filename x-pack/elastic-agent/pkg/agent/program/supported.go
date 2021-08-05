// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by x-pack/dev-tools/cmd/buildspec/buildspec.go - DO NOT EDIT.

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
