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
