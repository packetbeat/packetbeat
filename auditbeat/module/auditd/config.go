// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package auditd

import (
	"bufio"
	"bytes"
	"strings"
	"time"

	"github.com/joeshaw/multierror"
	"github.com/pkg/errors"

	"github.com/elastic/go-libaudit/rule"
	"github.com/elastic/go-libaudit/rule/flags"
)

const (
	moduleName    = "auditd"
	metricsetName = "auditd"
)

// Config defines the kernel metricset's possible configuration options.
type Config struct {
	ResolveIDs   bool   `config:"resolve_ids"`         // Resolve UID/GIDs to names.
	FailureMode  string `config:"failure_mode"`        // Failure mode for the kernel (silent, log, panic).
	BacklogLimit uint32 `config:"backlog_limit"`       // Max number of message to buffer in the auditd.
	RateLimit    uint32 `config:"rate_limit"`          // Rate limit in messages/sec of messages from auditd.
	RawMessage   bool   `config:"include_raw_message"` // Include the list of raw audit messages in the event.
	Warnings     bool   `config:"include_warnings"`    // Include warnings in the event (for dev/debug purposes only).
	RulesBlob    string `config:"audit_rules"`         // Audit rules. One rule per line.
	SocketType   string `config:"socket_type"`         // Socket type to use with the kernel (unicast or multicast).

	// Tuning options (advanced, use with care)
	ReassemblerMaxInFlight uint32        `config:"reassembler.max_in_flight"`
	ReassemblerTimeout     time.Duration `config:"reassembler.timeout"`
	StreamBufferQueueSize  uint32        `config:"reassembler.queue_size"`
	// BackpressureStrategy defines the strategy used to mitigate backpressure
	// propagating to the kernel causing audited processes to block until
	// Auditbeat can keep-up.
	// One of "user-space", "kernel", "both", "none", "auto" (default)
	BackpressureStrategy  string `config:"backpressure_strategy"`
	StreamBufferConsumers int    `config:"stream_buffer_consumers"`
}

type auditRule struct {
	flags string
	data  []byte
}

// Validate validates the rules specified in the config.
func (c *Config) Validate() error {
	var errs multierror.Errors
	_, err := c.rules()
	if err != nil {
		errs = append(errs, err)
	}
	_, err = c.failureMode()
	if err != nil {
		errs = append(errs, err)
	}

	c.SocketType = strings.ToLower(c.SocketType)
	switch c.SocketType {
	case "", "unicast", "multicast":
	default:
		errs = append(errs, errors.Errorf("invalid socket_type "+
			"'%v' (use unicast, multicast, or don't set a value)", c.SocketType))
	}

	return errs.Err()
}

// Rules returns a list of rules specified in the config.
func (c Config) rules() ([]auditRule, error) {
	var errs multierror.Errors
	var auditRules []auditRule
	ruleSet := map[string]auditRule{}
	s := bufio.NewScanner(bytes.NewBufferString(c.RulesBlob))
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if len(line) == 0 || line[0] == '#' {
			continue
		}

		// Parse the CLI flags into an intermediate rule specification.
		r, err := flags.Parse(line)
		if err != nil {
			errs = append(errs, errors.Wrapf(err, "failed on rule '%v'", line))
			continue
		}

		// Convert rule specification to a binary rule representation.
		data, err := rule.Build(r)
		if err != nil {
			errs = append(errs, errors.Wrapf(err, "failed on rule '%v'", line))
			continue
		}

		// Detect duplicates based on the normalized binary rule representation.
		existingRule, found := ruleSet[string(data)]
		if found {
			errs = append(errs, errors.Errorf("failed on rule '%v' because its a duplicate of '%v'", line, existingRule.flags))
			continue
		}
		auditRule := auditRule{flags: line, data: []byte(data)}
		ruleSet[string(data)] = auditRule

		auditRules = append(auditRules, auditRule)
	}

	if len(errs) > 0 {
		return nil, errors.Wrap(errs.Err(), "invalid audit_rules")
	}
	return auditRules, nil
}

func (c Config) failureMode() (uint32, error) {
	switch strings.ToLower(c.FailureMode) {
	case "silent":
		return 0, nil
	case "log":
		return 1, nil
	case "panic":
		return 2, nil
	default:
		return 0, errors.Errorf("invalid failure_mode '%v' (use silent, log, or panic)", c.FailureMode)
	}
}

var defaultConfig = Config{
	ResolveIDs:             true,
	FailureMode:            "silent",
	BacklogLimit:           8192,
	RateLimit:              0,
	RawMessage:             false,
	Warnings:               false,
	ReassemblerMaxInFlight: 50,
	ReassemblerTimeout:     2 * time.Second,
	StreamBufferQueueSize:  8192,
	StreamBufferConsumers:  0,
}
