// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/agent/errors"
	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/core/logger"
	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/release"
	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/remote"
)

// Sender is an sender interface describing client behavior.
type Sender interface {
	Send(
		ctx context.Context,
		method string,
		path string,
		params url.Values,
		headers http.Header,
		body io.Reader,
	) (*http.Response, error)

	URI() string
}

var baseRoundTrippers = func(rt http.RoundTripper) (http.RoundTripper, error) {
	rt = NewFleetUserAgentRoundTripper(rt, release.Version())
	return rt, nil
}

func init() {
	val, ok := os.LookupEnv("DEBUG_AGENT")
	if ok && val == "1" {
		fn := baseRoundTrippers
		baseRoundTrippers = func(rt http.RoundTripper) (http.RoundTripper, error) {
			rt, err := fn(rt)
			if err != nil {
				return nil, err
			}

			l, err := logger.New("fleet_client")
			if err != nil {
				return nil, errors.New(err, "could not create the logger for debugging HTTP request")
			}

			return remote.NewDebugRoundTripper(rt, l), nil
		}
	}
}

// NewAuthWithConfig returns a Kibana client that will:
//
// - Send the API Key on every HTTP request.
// - Ensure a minimun version of Kibana is required.
// - Send the Fleet User Agent on every HTTP request.
func NewAuthWithConfig(log *logger.Logger, apiKey string, cfg remote.Config) (*remote.Client, error) {
	return remote.NewWithConfig(log, cfg, func(rt http.RoundTripper) (http.RoundTripper, error) {
		rt, err := baseRoundTrippers(rt)
		if err != nil {
			return nil, err
		}

		rt, err = NewFleetAuthRoundTripper(rt, apiKey)
		if err != nil {
			return nil, err
		}

		return rt, nil
	})
}

// NewWithConfig takes a Kibana configuration and create a kibana.client with the appropriate tripper.
func NewWithConfig(log *logger.Logger, cfg remote.Config) (*remote.Client, error) {
	return remote.NewWithConfig(log, cfg, baseRoundTrippers)
}

// ExtractError extracts error from a fleet response
func ExtractError(resp io.Reader) error {
	// Lets try to extract a high level Kibana error.
	e := &struct {
		StatusCode int    `json:"statusCode"`
		Error      string `json:"error"`
		Message    string `json:"message"`
	}{}

	data, err := ioutil.ReadAll(resp)
	if err != nil {
		return errors.New(err, "fail to read original error")
	}

	err = json.Unmarshal(data, e)
	if err == nil {
		// System errors doesn't return a message, fleet code can return a Message key which has more
		// information.
		if len(e.Message) == 0 {
			return fmt.Errorf("Status code: %d, Kibana returned an error: %s", e.StatusCode, e.Error)
		}
		return fmt.Errorf(
			"Status code: %d, Kibana returned an error: %s, message: %s",
			e.StatusCode,
			e.Error,
			e.Message,
		)
	}

	return fmt.Errorf("could not decode the response, raw response: %s", string(data))
}
