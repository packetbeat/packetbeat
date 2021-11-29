// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by dev-tools/cmd/buildfleetcfg/buildfleetcfg.go - DO NOT EDIT.

package application

import "github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/packer"

// DefaultAgentFleetConfig is the content of the default configuration when we enroll a beat, the elastic-agent.yml
// will be replaced with this variables.
var DefaultAgentFleetConfig []byte

func init() {
	// Packed File
	// _meta/elastic-agent.fleet.yml
	unpacked := packer.MustUnpack("eJycVUuTo7oV3udnzD43PIauITsLN0KModO4jZA2KSS5BbaEqWsDhlT+e0rgtrtv3aSmsmAjpPP4Huf869s/9f5S/m2vyvOl5n8t5b65/Pau9vvLb6NW3/7+DY0r/x9v1q99WxBR7CmuwzNa/+Kbx/cTjSCgOLT4CFo+ggNzPF1iYXOdH1GUdxTmI3qmZ4pzaxMAj7k7SZywQzBTVKuObsGFuamFolSJKGuZFlNQJ5I2cc+2wCrhTm4dVRPsVSjM400N6D1GlA8Uhhad76kJQdoyuJMC/pDE8TuqVSOKWKEoO9EtmGiRWSX2Gj6a+6Fl8hDzH3q9CMAkCjAwN7YofpWkOcpMXxUtXrtAL/W/yJNEMO8Ijs8Uvz6hKBsF3v2c42ChWPPaUei7TF97gl+fzDkKVgcUpb0o4gPdgp7WS0+kiK0S04q42bQJwEiL0C6LWPERKAbDSUB1QPDals5OctfkSa25Fxh2NAAdwXbLNJfMIVLASiEYt0yHY4nz6Z4X7jo6rnrhqLkuPtrfX9ak22s1LHfAxNx8JE6e747HJxTEJxFlA59O/cYJRwGVJji1+Ojda9rotN84WS8c78yc8MhHv174/1E/+gWVgHLmYakta7nODwL64347a8QiWHWGCwQfsQzuHIaH0gkbWiQ/0biSKMoqroUSoX80+BA363ljakXHvTvn6YwGSGH1D+yzCcnTAQUriQzX2FbMzS0Eac9rcGSu6Bj0KwpVR0cwEJz+bvi7YXK//1KvrskBfMJz0ReHoVUGwGIjqFAEbKKvLRnB77Q4SlGkigZAMxdJofOxxNT7sz4/8/+Ir/RNi7rEV4WMpiOjjdRoS2b4WnE3M7kAdfJOBMAt8fVsdLjfzvXM8f9Y081LXzhEkVCkSVsBd92t74G6S96XelXH0WUgRXba6KwSMPQ3Tn4mRWqVOOk4FMMnfFvWpBbB13Ox8CEpzCfixi3R+XTz3pk5aSWg6lk9e+D+5g+c9zyKe+rGFSuSDj0ra8YAhmfu/FiwXd3zzhhznd955NpvuA4vKEoHitOWanUgRdYyx5s2t3pL7E0ChmcW+nN/f96r6ribVQwOdWDewbCh2LM2DRiZkyruJg9vQ+NDdSF4kBz6owhAJ7Bd0wJJqv3R1EIK0c19RnFPnHzizodHwaXE3vv8bn2SydtqSNarRfdBIhkOv5fYttl2mUsIejbDseL1rXeDt+MZvZneD7RIp7tvoF3tQ39+91KvpmS9GpLVf/XEweAoFhw7WvAvtW5q4JYw74wWb/3Lskgk076Z2zaZz1It8PVs/L3Eyo8Ixp7h39wVUF0o9m3xqMHwYPDqOczHhafvc1wzu7mjjrSITW8VgrQq8dXmenebCQumhZu1DOe9KF6f0DoZ+HAy+6havOGNtMjG/dvpjifFXkXMTA/AkRaZ4tpwZzQqTrTIToYzrnOLN8e5Dg79iTuqZnDXoWezr0KbRa/zP6qNJu/eUUyHNYP58aUGFm9y9ZiFsRJR3NIifSeObxsfcxco4ihtZvUnTo0v5prN+aYGzxTTSuDr4oF10t1jhnNMD8G0F9izEPQ1Wj9LhvOKNYabVykcdWYBqOn2HvOIoN/NuxLbZo9oYXbzcqYJVmezXz58MuOgleK2f2Bu3s36HO7cHSm+VvstmDjMDyWmLTWYQDMjjrLE3yUzPOrcEk6u2P/uy/A2x729f/p8hkZwIUX1/hH7K3Z2tYfqIrD1f+dOoi9+Wfp6O8kUW3fdZEU1MMdTrMla43MKc4sWsfGu2dfGG1Y5Y5mZGestczsbzV5cZn0i6dcYn7gUrZmZyQSmx3uz41fTQzvJjMkmWDlJMMjkDX3G52tPkFYsSj92mdpHoGc67wRUFVufPnT789u///KfAAAA//+w0n9F")
	raw, ok := unpacked["_meta/elastic-agent.fleet.yml"]
	if !ok {
		// ensure we have something loaded.
		panic("elastic-agent.fleet.yml is not included in the binary")
	}
	DefaultAgentFleetConfig = raw
}
