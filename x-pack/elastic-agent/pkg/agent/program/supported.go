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
