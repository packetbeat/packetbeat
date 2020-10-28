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
	// spec/endpoint.yml
	// spec/filebeat.yml
	// spec/heartbeat.yml
	// spec/metricbeat.yml
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
