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

package http

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/elastic/beats/libbeat/logp"

	"github.com/elastic/beats/heartbeat/eventext"
	"github.com/elastic/beats/heartbeat/look"
	"github.com/elastic/beats/heartbeat/monitors"
	"github.com/elastic/beats/heartbeat/monitors/active/dialchain"
	"github.com/elastic/beats/heartbeat/monitors/jobs"
	"github.com/elastic/beats/heartbeat/reason"
	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/outputs/transport"
)

func newHTTPMonitorHostJob(
	addr string,
	config *Config,
	transport *http.Transport,
	enc contentEncoder,
	body []byte,
	validator RespCheck,
) (jobs.Job, error) {

	client := &http.Client{
		CheckRedirect: makeCheckRedirect(config.MaxRedirects),
		Transport:     transport,
		Timeout:       config.Timeout,
	}
	request, err := buildRequest(addr, config, enc)
	if err != nil {
		return nil, err
	}

	timeout := config.Timeout

	return jobs.MakeSimpleJob(func(event *beat.Event) error {
		_, _, err := execPing(event, client, request, body, timeout, validator, config.IncludeBody)
		return err
	}), nil
}

func newHTTPMonitorIPsJob(
	config *Config,
	addr string,
	tls *transport.TLSConfig,
	enc contentEncoder,
	body []byte,
	validator RespCheck,
) (jobs.Job, error) {

	req, err := buildRequest(addr, config, enc)
	if err != nil {
		return nil, err
	}

	hostname, port, err := splitHostnamePort(req)
	if err != nil {
		return nil, err
	}

	settings := monitors.MakeHostJobSettings(hostname, config.Mode)

	pingFactory := createPingFactory(config, port, tls, req, body, validator)
	job, err := monitors.MakeByHostJob(settings, pingFactory)

	return job, err
}

func createPingFactory(
	config *Config,
	port uint16,
	tls *transport.TLSConfig,
	request *http.Request,
	body []byte,
	validator RespCheck,
) func(*net.IPAddr) jobs.Job {
	timeout := config.Timeout
	isTLS := request.URL.Scheme == "https"
	checkRedirect := makeCheckRedirect(config.MaxRedirects)

	return monitors.MakePingIPFactory(func(event *beat.Event, ip *net.IPAddr) error {
		addr := net.JoinHostPort(ip.String(), strconv.Itoa(int(port)))
		d := &dialchain.DialerChain{
			Net: dialchain.MakeConstAddrDialer(addr, dialchain.TCPDialer(timeout)),
		}

		// TODO: add socks5 proxy?

		if isTLS {
			d.AddLayer(dialchain.TLSLayer(tls, timeout))
		}

		dialer, err := d.Build(event)
		if err != nil {
			return err
		}

		var (
			writeStart, readStart, writeEnd time.Time
		)
		// Ensure memory consistency for these callbacks.
		// It seems they can be invoked still sometime after the request is done
		cbMutex := sync.Mutex{}

		client := &http.Client{
			CheckRedirect: checkRedirect,
			Timeout:       timeout,
			Transport: &SimpleTransport{
				Dialer: dialer,
				OnStartWrite: func() {
					cbMutex.Lock()
					writeStart = time.Now()
					cbMutex.Unlock()
				},
				OnEndWrite: func() {
					cbMutex.Lock()
					writeEnd = time.Now()
					cbMutex.Unlock()
				},
				OnStartRead: func() {
					cbMutex.Lock()
					readStart = time.Now()
					cbMutex.Unlock()
				},
			},
		}

		_, end, err := execPing(event, client, request, body, timeout, validator)
		cbMutex.Lock()
		defer cbMutex.Unlock()

		if !readStart.IsZero() {
			eventext.MergeEventFields(event, common.MapStr{
				"http": common.MapStr{
					"rtt": common.MapStr{
						"write_request":   look.RTT(writeEnd.Sub(writeStart)),
						"response_header": look.RTT(readStart.Sub(writeStart)),
					},
				},
			})
		}
		if !writeStart.IsZero() {
			event.PutValue("http.rtt.validate", look.RTT(end.Sub(writeStart)))
			event.PutValue("http.rtt.content", look.RTT(end.Sub(readStart)))
		}

		return err
	})
}

func buildRequest(addr string, config *Config, enc contentEncoder) (*http.Request, error) {
	method := strings.ToUpper(config.Check.Request.Method)
	request, err := http.NewRequest(method, addr, nil)
	if err != nil {
		return nil, err
	}
	request.Close = true

	if config.Username != "" {
		request.SetBasicAuth(config.Username, config.Password)
	}
	for k, v := range config.Check.Request.SendHeaders {
		// defining the Host header isn't enough. See https://github.com/golang/go/issues/7682
		if k == "Host" {
			request.Host = v
		}

		request.Header.Add(k, v)
	}

	if enc != nil {
		enc.AddHeaders(&request.Header)
	}

	return request, nil
}

func execPing(
	event *beat.Event,
	client *http.Client,
	req *http.Request,
	reqBody []byte,
	timeout time.Duration,
	validator func(*http.Response) error,
	responseConfig responseConfig,
) (start, end time.Time, errReason reason.Reason) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req = attachRequestBody(&ctx, req, reqBody)

	start, end, resp, errReason := execRequest(client, req, validator)

	maxBodyBytes := responseConfig.IncludeBodyMaxBytes
	if responseConfig.IncludeBody == "never" {
		// Don't return the body if the config says not to
		maxBodyBytes = 0
	} else if errReason == nil && responseConfig.IncludeBody == "on_error" {
		// If configured to only return the body on error, don't return it on success
		maxBodyBytes = 0
	}
	respBody, respSize := readBodySampleAndClose(resp, maxBodyBytes)

	if errReason == nil || errReason.Type() != "io" {
		eventext.MergeEventFields(event, common.MapStr{"http": common.MapStr{
			"rtt": common.MapStr{
				"total": look.RTT(end.Sub(start)),
			},
		}})
	}

	if resp != nil {
		responseFields := common.MapStr{"status_code": resp.StatusCode}
		if respSize > -1 {
			responseFields["body.content"] = respBody
			responseFields["body.bytes"] = respSize
		}
		eventext.MergeEventFields(event, common.MapStr{"http": common.MapStr{
			"response": responseFields,
		}})
	}

	if errReason != nil {
		if resp != nil {
			return start, end, errReason
		}
		return start, end, errReason
	}

	return start, end, nil
}

// readBodySampleAndClose reads the first sampleSize bytes from the httpResponse,
// then closes the body (which closes the connection). It doesn't return any errors
// but does log them. During an error case the return values will be (nil, -1).
// The maxBytes params controls how many bytes will be returned in a string, not how many will be read.
// We always read the full response here since we want to time downloading the full thing.
func readBodySampleAndClose(resp *http.Response, maxBytes int) (bodySample *string, bodySize int64) {
	if resp == nil {
		return nil, -1
	}
	defer resp.Body.Close()

	// Function to lazily get the body of the response
	buf := make([]byte, maxBytes)
	respSize := int64(-1)
	if resp != nil {
		startSize, readErr := resp.Body.Read(buf)
		if startSize > 0 {
			buf = buf[0:startSize]
			// Read the entirety of the body. Otherwise, the stats for the check
			// don't include download time.
			restSize, _ := io.Copy(ioutil.Discard, resp.Body)
			respSize = int64(startSize) + restSize
		} else if readErr != nil {
			logp.Warn("could not read HTTP response body after ping: %s", readErr)
			buf = buf[:0]
		}
	}

	bodyStr := string(buf)
	return &bodyStr, respSize
}

func attachRequestBody(ctx *context.Context, req *http.Request, body []byte) *http.Request {
	req = req.WithContext(*ctx)
	if len(body) > 0 {
		req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		req.ContentLength = int64(len(body))
	}

	return req
}

// execute the request. Note that this does not close the resp body, which should be done by caller
func execRequest(client *http.Client, req *http.Request, validator func(*http.Response) error) (start time.Time, end time.Time, resp *http.Response, errReason reason.Reason) {
	start = time.Now()
	resp, err := client.Do(req)
	end = time.Now()

	if err != nil {
		return start, end, nil, reason.IOFailed(err)
	}

	err = validator(resp)
	if err != nil {
		return start, time.Now(), resp, reason.ValidateFailed(err)
	}

	return start, time.Now(), resp, nil
}

func splitHostnamePort(requ *http.Request) (string, uint16, error) {
	host := requ.URL.Host
	// Try to add a default port if needed
	if strings.LastIndex(host, ":") == -1 {
		switch requ.URL.Scheme {
		case urlSchemaHTTP:
			host += ":80"
		case urlSchemaHTTPS:
			host += ":443"
		}
	}
	host, port, err := net.SplitHostPort(host)
	if err != nil {
		return "", 0, err
	}
	p, err := strconv.ParseUint(port, 10, 16)
	if err != nil {
		return "", 0, fmt.Errorf("'%v' is no valid port number in '%v'", port, requ.URL.Host)
	}
	return host, uint16(p), nil
}

func makeCheckRedirect(max int) func(*http.Request, []*http.Request) error {
	if max == 0 {
		return func(_ *http.Request, _ []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}

	return func(_ *http.Request, via []*http.Request) error {
		if max == len(via) {
			return http.ErrUseLastResponse
		}
		return nil
	}
}
