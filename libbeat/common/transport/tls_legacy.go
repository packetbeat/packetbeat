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

// +build !go1.15

package transport

import (
	"crypto/tls"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/elastic/beats/v7/libbeat/common/transport/tlscommon"
	"github.com/elastic/beats/v7/libbeat/testing"
)

func TestTLSDialer(
	d testing.Driver,
	forward Dialer,
	config *tlscommon.TLSConfig,
	timeout time.Duration,
) Dialer {
	var lastTLSConfig *tls.Config
	var lastNetwork string
	var lastAddress string
	var m sync.Mutex

	return DialerFunc(func(network, address string) (net.Conn, error) {
		switch network {
		case "tcp", "tcp4", "tcp6":
		default:
			return nil, fmt.Errorf("unsupported network type %v", network)
		}

		host, _, err := net.SplitHostPort(address)
		if err != nil {
			return nil, err
		}

		var tlsConfig *tls.Config
		m.Lock()
		if network == lastNetwork && address == lastAddress {
			tlsConfig = lastTLSConfig
		}
		if tlsConfig == nil {
			tlsConfig = config.BuildModuleConfig(host)
			lastNetwork = network
			lastAddress = address
			lastTLSConfig = tlsConfig
		}
		m.Unlock()

		return tlsDialWith(d, forward, network, address, timeout, tlsConfig, config)
	})
}
