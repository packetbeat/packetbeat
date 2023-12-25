// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package salesforce

import (
	"time"

	"github.com/elastic/elastic-agent-libs/logp"
	"github.com/elastic/elastic-agent-libs/mapstr"
)

// state is the state of the salesforce module. It is used to watermark the state
// to avoid pulling duplicate data from Salesforce. The state is persisted separately
// for EventLogFile and Object.
type state struct {
	Object       dateTimeCursor `json:"object,omitempty"`
	EventLogFile dateTimeCursor `json:"event_log_file,omitempty"`
}

type dateTimeCursor struct {
	FirstEventTime string `struct:"first_event_time,omitempty"`
	LastEventTime  string `struct:"last_event_time,omitempty"`
}

// parseCursor parses the cursor from the configuration and executes the
// template. If cursor is nil, the default templated query is used else
// the value templated query is used. See QueryConfig struct for more.
func parseCursor(initialInterval *time.Duration, cfg *QueryConfig, cursor mapstr.M, log *logp.Logger) (string, error) {
	ctxTmpl := mapstr.M{"cursor": nil}

	if cursor != nil {
		ctxTmpl["cursor"] = cursor
		qr, err := cfg.Value.Execute(ctxTmpl, nil, log)
		if err != nil {
			return "", err
		}
		return qr, nil
	}

	return cfg.Default.Execute(ctxTmpl, nil, log)
}
