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
	"github.com/elastic/beats/v7/libbeat/common"
	"net/url"
	"strings"
)

func addToURL(urlStr string, params map[string]string) string {
	if strings.HasSuffix(urlStr, "/") {
		urlStr = strings.TrimSuffix(urlStr, "/")
	}
	if len(params) == 0 {
		return urlStr
	}
	values := url.Values{}
	for key, val := range params {
		values.Add(key, val)
	}
	return common.EncodeURLParams(urlStr, values)
}

func parseProxyURL(raw string) (*url.URL, error) {
	if raw == "" {
		return nil, nil
	}
	parsedUrl, err := url.Parse(raw)
	if err == nil && strings.HasPrefix(parsedUrl.Scheme, "http") {
		return parsedUrl, err
	}
	// Proxy was bogus. Try prepending "http://" to it and
	// see if that parses correctly.
	return url.Parse("http://" + raw)
}
