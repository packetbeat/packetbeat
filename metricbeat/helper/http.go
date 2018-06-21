package helper

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/elastic/beats/libbeat/common/transport/tlscommon"
	"github.com/elastic/beats/libbeat/outputs/transport"
	"github.com/elastic/beats/metricbeat/mb"
)

type HTTP struct {
	hostData  mb.HostData
	metricset string
	client    *http.Client // HTTP client that is reused across requests.
	headers   map[string]string
	uri       string
	method    string
	body      []byte
}

// HTTPConfig is used to configure the HTTP helper
type HTTPConfig struct {
	TLS     *tlscommon.Config `config:"ssl"`
	Timeout time.Duration     `config:"timeout"`
	Headers map[string]string `config:"headers"`
}

// NewHTTP creates new http helper
func NewHTTP(base mb.BaseMetricSet) (*HTTP, error) {
	config := HTTPConfig{}
	if err := base.Module().UnpackConfig(&config); err != nil {
		return nil, err
	}

	if config.Headers == nil {
		config.Headers = map[string]string{}
	}

	return NewHTTPFromConfig(&config, base.HostData(), base.Name())
}

// NewHTTPFromConfig builds and returns an HTTP helper using the given config parameters
func NewHTTPFromConfig(config *HTTPConfig, hostData mb.HostData, metricset string) (*HTTP, error) {
	tlsConfig, err := tlscommon.LoadTLSConfig(config.TLS)
	if err != nil {
		return nil, err
	}

	var dialer, tlsDialer transport.Dialer

	dialer = transport.NetDialer(config.Timeout)
	tlsDialer, err = transport.TLSDialer(dialer, tlsConfig, config.Timeout)
	if err != nil {
		return nil, err
	}

	return &HTTP{
		metricset: metricset,
		hostData:  hostData,
		client: &http.Client{
			Transport: &http.Transport{
				Dial:    dialer.Dial,
				DialTLS: tlsDialer.Dial,
			},
			Timeout: config.Timeout,
		},
		headers: config.Headers,
		method:  "GET",
		uri:     hostData.SanitizedURI,
		body:    nil,
	}, nil
}

// FetchResponse fetches a response for the http metricset.
// It's important that resp.Body has to be closed if this method is used. Before using this method
// check if one of the other Fetch* methods could be used as they ensure that the Body is properly closed.
func (h *HTTP) FetchResponse() (*http.Response, error) {
	// Create a fresh reader every time
	var reader io.Reader
	if h.body != nil {
		reader = bytes.NewReader(h.body)
	}

	req, err := http.NewRequest(h.method, h.uri, reader)
	if h.hostData.User != "" || h.hostData.Password != "" {
		req.SetBasicAuth(h.hostData.User, h.hostData.Password)
	}

	for k, v := range h.headers {
		req.Header.Set(k, v)
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making http request: %v", err)
	}

	return resp, nil
}

// SetHeader sets HTTP headers to use in requests
func (h *HTTP) SetHeader(key, value string) {
	h.headers[key] = value
}

// SetMethod sets HTTP method to use in requests
func (h *HTTP) SetMethod(method string) {
	h.method = method
}

// GetURI gets the URI used in requests
func (h *HTTP) GetURI() string {
	return h.uri
}

// SetURI sets URI to use in requests
func (h *HTTP) SetURI(uri string) {
	h.uri = uri
}

// SetBody sets the body of the requests
func (h *HTTP) SetBody(body []byte) {
	h.body = body
}

// FetchContent makes an HTTP request to the configured url and returns the body content.
func (h *HTTP) FetchContent() ([]byte, error) {
	resp, err := h.FetchResponse()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP error %d in %s: %s", resp.StatusCode, h.metricset, resp.Status)
	}

	return ioutil.ReadAll(resp.Body)
}

// FetchScanner returns a Scanner for the content.
func (h *HTTP) FetchScanner() (*bufio.Scanner, error) {
	content, err := h.FetchContent()
	if err != nil {
		return nil, err
	}

	return bufio.NewScanner(bytes.NewReader(content)), nil
}

// FetchJSON makes an HTTP request to the configured url and returns the JSON content.
// This only works if the JSON output needed is in map[string]interface format.
func (h *HTTP) FetchJSON() (map[string]interface{}, error) {
	body, err := h.FetchContent()
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
