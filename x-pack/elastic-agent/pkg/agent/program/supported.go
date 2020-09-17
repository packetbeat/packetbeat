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
