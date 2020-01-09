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

package syslog

import (
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/elastic/beats/filebeat/channel"
	"github.com/elastic/beats/filebeat/harvester"
	"github.com/elastic/beats/filebeat/input"
	"github.com/elastic/beats/filebeat/inputsource"
	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/common/cfgwarn"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/influxdata/go-syslog"
	"github.com/influxdata/go-syslog/rfc5424"
)

// RFC 3164 parser is generated from a ragel state machine using the following command:
//go:generate ragel -Z -G2 parser.rl -o parser.go
//go:generate go fmt parser.go

// Severity and Facility are derived from the priority, theses are the human readable terms
// defined in https://tools.ietf.org/html/rfc3164#section-4.1.1.
//
// Example:
// 2 => "Critical"
type mapper []string

var (
	severityLabels = mapper{
		"Emergency",
		"Alert",
		"Critical",
		"Error",
		"Warning",
		"Notice",
		"Informational",
		"Debug",
	}

	facilityLabels = mapper{
		"kernel",
		"user-level",
		"mail",
		"system",
		"security/authorization",
		"syslogd",
		"line printer",
		"network news",
		"UUCP",
		"clock",
		"security/authorization",
		"FTP",
		"NTP",
		"log audit",
		"log alert",
		"clock",
		"local0",
		"local1",
		"local2",
		"local3",
		"local4",
		"local5",
		"local6",
		"local7",
	}
)

func init() {
	err := input.Register("syslog", NewInput)
	if err != nil {
		panic(err)
	}
}

// Input define a syslog input
type Input struct {
	sync.Mutex
	started bool
	outlet  channel.Outleter
	server  inputsource.Network
	config  *config
	log     *logp.Logger
}

// NewInput creates a new syslog input
func NewInput(
	cfg *common.Config,
	outlet channel.Connector,
	context input.Context,
) (input.Input, error) {
	cfgwarn.Experimental("Syslog input type is used")

	log := logp.NewLogger("syslog")

	out, err := outlet.ConnectWith(cfg, beat.ClientConfig{
		Processing: beat.ProcessingConfig{
			DynamicFields: context.DynamicFields,
		},
	})
	if err != nil {
		return nil, err
	}

	config := defaultConfig
	if err = cfg.Unpack(&config); err != nil {
		return nil, err
	}

	var cb inputsource.NetworkFunc
	forwarder := harvester.NewForwarder(out)
	switch config.Format {
	case syslogFormatRFC3164:
		cb = networkFuncFor3164(log, forwarder)
	case syslogFormatRFC5424:
		cb = networkFuncFor5424(log, forwarder)
	}

	server, err := factory(cb, config.Protocol)
	if err != nil {
		return nil, err
	}

	return &Input{
		outlet:  out,
		started: false,
		server:  server,
		config:  &config,
		log:     log,
	}, nil
}

func networkFuncFor3164(
	log *logp.Logger, forwarder *harvester.Forwarder,
) inputsource.NetworkFunc {
	return (func(data []byte, metadata inputsource.NetworkMetadata) {
		ev := newEvent()
		Parse(data, ev)
		if !ev.IsValid() {
			log.Errorw("can't parse event as syslog rfc3164", "message", string(data))
			// On error revert to the raw bytes content, we need a better way to communicate this kind of
			// error upstream this should be a global effort.
			forwarder.Send(beat.Event{
				Timestamp: time.Now(),
				Meta: common.MapStr{
					"truncated": metadata.Truncated,
				},
				Fields: common.MapStr{
					"message": string(data),
				},
			})
		} else {
			forwarder.Send(createEvent3164(ev, metadata, time.Local, log))
		}
	})
}

// This is very similar to the network function for RFC 3164, but hard to
// combine because the underlying library API is different, so for now there is
// some duplication.
func networkFuncFor5424(
	log *logp.Logger, forwarder *harvester.Forwarder,
) inputsource.NetworkFunc {
	parser := rfc5424.NewParser()
	return (func(data []byte, metadata inputsource.NetworkMetadata) {
		message, err := parser.Parse(data)
		if err != nil {
			log.Errorw("can't parse event as syslog rfc5424",
				"message", string(data), "error", err)
			// On error revert to the raw bytes content, we need a better way to communicate this kind of
			// error upstream this should be a global effort.
			forwarder.Send(beat.Event{
				Timestamp: time.Now(),
				Meta: common.MapStr{
					"truncated": metadata.Truncated,
				},
				Fields: common.MapStr{
					"message": string(data),
				},
			})
		} else {
			forwarder.Send(createEvent5424(message, metadata, time.Local, log))
		}
	})
}

// Run starts listening for Syslog events over the network.
func (p *Input) Run() {
	p.Lock()
	defer p.Unlock()

	if !p.started {
		p.log.Infow("Starting Syslog input", "protocol", p.config.Protocol.Name())
		err := p.server.Start()
		if err != nil {
			p.log.Error("Error starting the server", "error", err)
			return
		}
		p.started = true
	}
}

// Stop stops the syslog input.
func (p *Input) Stop() {
	defer p.outlet.Close()
	p.Lock()
	defer p.Unlock()

	if !p.started {
		return
	}

	p.log.Info("Stopping Syslog input")
	p.server.Stop()
	p.started = false
}

// Wait stops the syslog input.
func (p *Input) Wait() {
	p.Stop()
}

func createEvent3164(ev *event, metadata inputsource.NetworkMetadata, timezone *time.Location, log *logp.Logger) beat.Event {
	f := common.MapStr{
		"message": strings.TrimRight(ev.Message(), "\n"),
		"log": common.MapStr{
			"source": common.MapStr{
				"address": metadata.RemoteAddr.String(),
			},
		},
	}

	syslog := common.MapStr{}
	event := common.MapStr{}
	process := common.MapStr{}

	if ev.Hostname() != "" {
		f["hostname"] = ev.Hostname()
	}

	if ev.HasPid() {
		process["pid"] = ev.Pid()
	}

	if ev.Program() != "" {
		process["program"] = ev.Program()
	}

	if ev.HasPriority() {
		syslog["priority"] = ev.Priority()

		event["severity"] = ev.Severity()
		v, err := mapValueToName(ev.Severity(), severityLabels)
		if err != nil {
			log.Debugw("could not find severity label", "error", err)
		} else {
			syslog["severity_label"] = v
		}

		syslog["facility"] = ev.Facility()
		v, err = mapValueToName(ev.Facility(), facilityLabels)
		if err != nil {
			log.Debugw("could not find facility label", "error", err)
		} else {
			syslog["facility_label"] = v
		}
	}

	f["syslog"] = syslog
	f["event"] = event
	if len(process) > 0 {
		f["process"] = process
	}

	if ev.Sequence() != -1 {
		f["event.sequence"] = ev.Sequence()
	}

	return beat.Event{
		Timestamp: ev.Timestamp(timezone),
		Meta: common.MapStr{
			"truncated": metadata.Truncated,
		},
		Fields: f,
	}
}

func createEvent5424(
	message syslog.Message,
	metadata inputsource.NetworkMetadata,
	timezone *time.Location,
	log *logp.Logger,
) beat.Event {
	if message.Message() == nil {
	}
	f := common.MapStr{
		"log": common.MapStr{
			"source": common.MapStr{
				"address": metadata.RemoteAddr.String(),
			},
		},
	}
	if message.Message() != nil {
		f["message"] = strings.TrimRight(*message.Message(), "\n")
	}

	syslog := common.MapStr{}
	event := common.MapStr{}
	process := common.MapStr{}

	if message.Hostname() != nil {
		hostname := *message.Hostname()
		if hostname != "" {
			f["hostname"] = hostname
		}
	}
	if message.ProcID() != nil {
		process["pid"] = *message.ProcID()
	}
	if message.Appname() != nil {
		process["program"] = *message.Appname()
	}
	if message.Priority() != nil {
		syslog["priority"] = *message.Priority()
	}
	if message.Severity() != nil {
		severity := *message.Severity()
		event["severity"] = severity
		// Severity values are the same in both syslog versions.
		v, err := mapValueToName(int(severity), severityLabels)
		if err != nil {
			log.Debugw("could not find severity label", "error", err)
		} else {
			syslog["severity_label"] = v
		}
	}
	if message.Facility() != nil {
		// Facility values are the same in both syslog versions.
		//facility := *message.Facility()

	}
	/*

		if ev.HasPriority() {
			syslog["priority"] = ev.Priority()

			event["severity"] = ev.Severity()
			v, err := mapValueToName(ev.Severity(), severityLabels)
			if err != nil {
				log.Debugw("could not find severity label", "error", err)
			} else {
				syslog["severity_label"] = v
			}

			syslog["facility"] = ev.Facility()
			v, err = mapValueToName(ev.Facility(), facilityLabels)
			if err != nil {
				log.Debugw("could not find facility label", "error", err)
			} else {
				syslog["facility_label"] = v
			}
		}

		f["syslog"] = syslog
		f["event"] = event
		if len(process) > 0 {
			f["process"] = process
		}

		if ev.Sequence() != -1 {
			f["event.sequence"] = ev.Sequence()
		}
	*/
	var timestamp time.Time
	timestampPtr := message.Timestamp()
	if timestampPtr != nil {
		timestamp = *timestampPtr
	} else {
		// Fall back to the current time if no timestamp is included.
		timestamp = time.Now()
	}
	return beat.Event{
		Timestamp: timestamp,
		Meta: common.MapStr{
			"truncated": metadata.Truncated,
		},
		Fields: f,
	}
}

func mapValueToName(v int, m mapper) (string, error) {
	if v < 0 || v >= len(m) {
		return "", errors.Errorf("value out of bound: %d", v)
	}
	return m[v], nil
}
