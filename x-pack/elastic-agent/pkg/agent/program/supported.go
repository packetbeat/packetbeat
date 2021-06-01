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
<<<<<<< HEAD
	unpacked := packer.MustUnpack("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")
=======
	unpacked := packer.MustUnpack("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")
>>>>>>> 77c0973dd (Osquerybeat - Linux AARCH64 support (#25925))
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
