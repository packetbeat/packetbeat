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
