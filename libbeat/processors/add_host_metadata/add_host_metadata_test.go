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

package add_host_metadata

import (
	"fmt"
	"github.com/elastic/elastic-agent-libs/monitoring"
	"net"
	"os"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/features"
	conf "github.com/elastic/elastic-agent-libs/config"
	"github.com/elastic/elastic-agent-libs/mapstr"
	"github.com/elastic/go-sysinfo/types"

	"github.com/foxcpp/go-mockdns"
)

var (
	hostName = "testHost"
	hostID   = "9C7FAB7B"
)

func deregisterMonitoringRegistry() {
	monitoring.Default.Remove(logName)
}

func TestConfigDefault(t *testing.T) {
	event := &beat.Event{
		Fields:    mapstr.M{},
		Timestamp: time.Now(),
	}
	testConfig, err := conf.NewConfigFrom(map[string]interface{}{})
	assert.NoError(t, err)

	p, err := New(testConfig)
	defer deregisterMonitoringRegistry()

	switch runtime.GOOS {
	case "windows", "darwin", "linux":
		assert.NoError(t, err)
	default:
		assert.IsType(t, types.ErrNotImplemented, err)
		return
	}

	newEvent, err := p.Run(event)
	assert.NoError(t, err)

	v, err := newEvent.GetValue("host.os.family")
	assert.NoError(t, err)
	assert.NotNil(t, v)

	v, err = newEvent.GetValue("host.os.kernel")
	assert.NoError(t, err)
	assert.NotNil(t, v)

	v, err = newEvent.GetValue("host.os.name")
	assert.NoError(t, err)
	assert.NotNil(t, v)

	v, err = newEvent.GetValue("host.ip")
	assert.NoError(t, err)
	assert.NotNil(t, v)

	v, err = newEvent.GetValue("host.mac")
	assert.NoError(t, err)
	assert.NotNil(t, v)

	v, err = newEvent.GetValue("host.os.type")
	assert.NoError(t, err)
	assert.NotNil(t, v)
}

func TestConfigNetInfoDisabled(t *testing.T) {
	event := &beat.Event{
		Fields:    mapstr.M{},
		Timestamp: time.Now(),
	}
	testConfig, err := conf.NewConfigFrom(map[string]interface{}{
		"netinfo.enabled": false,
	})
	assert.NoError(t, err)

	p, err := New(testConfig)
	defer deregisterMonitoringRegistry()

	switch runtime.GOOS {
	case "windows", "darwin", "linux":
		assert.NoError(t, err)
	default:
		assert.IsType(t, types.ErrNotImplemented, err)
		return
	}

	newEvent, err := p.Run(event)
	assert.NoError(t, err)

	v, err := newEvent.GetValue("host.os.family")
	assert.NoError(t, err)
	assert.NotNil(t, v)

	v, err = newEvent.GetValue("host.os.kernel")
	assert.NoError(t, err)
	assert.NotNil(t, v)

	v, err = newEvent.GetValue("host.os.name")
	assert.NoError(t, err)
	assert.NotNil(t, v)

	v, err = newEvent.GetValue("host.ip")
	assert.Error(t, err)
	assert.Nil(t, v)

	v, err = newEvent.GetValue("host.mac")
	assert.Error(t, err)
	assert.Nil(t, v)

	v, err = newEvent.GetValue("host.os.type")
	assert.NoError(t, err)
	assert.NotNil(t, v)
}

func TestConfigName(t *testing.T) {
	event := &beat.Event{
		Fields:    mapstr.M{},
		Timestamp: time.Now(),
	}

	config := map[string]interface{}{
		"name": "my-host",
	}

	testConfig, err := conf.NewConfigFrom(config)
	assert.NoError(t, err)

	p, err := New(testConfig)
	require.NoError(t, err)
	defer deregisterMonitoringRegistry()

	newEvent, err := p.Run(event)
	assert.NoError(t, err)

	for configKey, configValue := range config {
		t.Run(fmt.Sprintf("Check of %s", configKey), func(t *testing.T) {
			v, err := newEvent.GetValue(fmt.Sprintf("host.%s", configKey))
			assert.NoError(t, err)
			assert.Equal(t, configValue, v, "Could not find in %s", newEvent)
		})
	}
}

func TestConfigGeoEnabled(t *testing.T) {
	event := &beat.Event{
		Fields:    mapstr.M{},
		Timestamp: time.Now(),
	}

	config := map[string]interface{}{
		"geo.name":             "yerevan-am",
		"geo.location":         "40.177200, 44.503490",
		"geo.continent_name":   "Asia",
		"geo.country_name":     "Armenia",
		"geo.country_iso_code": "AM",
		"geo.region_name":      "Erevan",
		"geo.region_iso_code":  "AM-ER",
		"geo.city_name":        "Yerevan",
	}

	testConfig, err := conf.NewConfigFrom(config)
	assert.NoError(t, err)
	defer deregisterMonitoringRegistry()

	p, err := New(testConfig)
	require.NoError(t, err)

	newEvent, err := p.Run(event)
	assert.NoError(t, err)

	eventGeoField, err := newEvent.GetValue("host.geo")
	require.NoError(t, err)

	assert.Len(t, eventGeoField, len(config))
}

func TestConfigGeoDisabled(t *testing.T) {
	event := &beat.Event{
		Fields:    mapstr.M{},
		Timestamp: time.Now(),
	}

	config := map[string]interface{}{}

	testConfig, err := conf.NewConfigFrom(config)
	require.NoError(t, err)

	p, err := New(testConfig)
	require.NoError(t, err)
	defer deregisterMonitoringRegistry()

	newEvent, err := p.Run(event)

	require.NoError(t, err)

	eventGeoField, err := newEvent.GetValue("host.geo")
	assert.Error(t, err)
	assert.Equal(t, nil, eventGeoField)
}

func TestEventWithReplaceFieldsFalse(t *testing.T) {
	cfg := map[string]interface{}{}
	cfg["replace_fields"] = false
	testConfig, err := conf.NewConfigFrom(cfg)
	assert.NoError(t, err)

	p, err := New(testConfig)
	defer deregisterMonitoringRegistry()

	switch runtime.GOOS {
	case "windows", "darwin", "linux":
		assert.NoError(t, err)
	default:
		assert.IsType(t, types.ErrNotImplemented, err)
		return
	}

	cases := []struct {
		title                   string
		event                   beat.Event
		hostLengthLargerThanOne bool
		hostLengthEqualsToOne   bool
		expectedHostFieldLength int
	}{
		{
			"replace_fields=false with only host.name",
			beat.Event{
				Fields: mapstr.M{
					"host": mapstr.M{
						"name": hostName,
					},
				},
			},
			true,
			false,
			-1,
		},
		{
			"replace_fields=false with only host.id",
			beat.Event{
				Fields: mapstr.M{
					"host": mapstr.M{
						"id": hostID,
					},
				},
			},
			false,
			true,
			1,
		},
		{
			"replace_fields=false with host.name and host.id",
			beat.Event{
				Fields: mapstr.M{
					"host": mapstr.M{
						"name": hostName,
						"id":   hostID,
					},
				},
			},
			true,
			false,
			2,
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			newEvent, err := p.Run(&c.event)
			assert.NoError(t, err)

			v, err := newEvent.GetValue("host")
			assert.NoError(t, err)
			assert.Equal(t, c.hostLengthLargerThanOne, len(v.(mapstr.M)) > 1)
			assert.Equal(t, c.hostLengthEqualsToOne, len(v.(mapstr.M)) == 1)
			if c.expectedHostFieldLength != -1 {
				assert.Equal(t, c.expectedHostFieldLength, len(v.(mapstr.M)))
			}
		})
	}
}

func TestEventWithReplaceFieldsTrue(t *testing.T) {
	cfg := map[string]interface{}{}
	cfg["replace_fields"] = true
	testConfig, err := conf.NewConfigFrom(cfg)
	assert.NoError(t, err)

	p, err := New(testConfig)
	defer deregisterMonitoringRegistry()

	switch runtime.GOOS {
	case "windows", "darwin", "linux":
		assert.NoError(t, err)
	default:
		assert.IsType(t, types.ErrNotImplemented, err)
		return
	}

	cases := []struct {
		title                   string
		event                   beat.Event
		hostLengthLargerThanOne bool
		hostLengthEqualsToOne   bool
	}{
		{
			"replace_fields=true with host.name",
			beat.Event{
				Fields: mapstr.M{
					"host": mapstr.M{
						"name": hostName,
					},
				},
			},
			true,
			false,
		},
		{
			"replace_fields=true with host.id",
			beat.Event{
				Fields: mapstr.M{
					"host": mapstr.M{
						"id": hostID,
					},
				},
			},
			true,
			false,
		},
		{
			"replace_fields=true with host.name and host.id",
			beat.Event{
				Fields: mapstr.M{
					"host": mapstr.M{
						"name": hostName,
						"id":   hostID,
					},
				},
			},
			true,
			false,
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			newEvent, err := p.Run(&c.event)
			assert.NoError(t, err)

			v, err := newEvent.GetValue("host")
			assert.NoError(t, err)
			assert.Equal(t, c.hostLengthLargerThanOne, len(v.(mapstr.M)) > 1)
			assert.Equal(t, c.hostLengthEqualsToOne, len(v.(mapstr.M)) == 1)
		})
	}
}

func TestSkipAddingHostMetadata(t *testing.T) {
	hostIDMap := map[string]string{}
	hostIDMap["id"] = hostID

	hostNameMap := map[string]string{}
	hostNameMap["name"] = hostName

	hostIDNameMap := map[string]string{}
	hostIDNameMap["id"] = hostID
	hostIDNameMap["name"] = hostName

	cases := []struct {
		title        string
		event        beat.Event
		expectedSkip bool
	}{
		{
			"event only with host.name",
			beat.Event{
				Fields: mapstr.M{
					"host": mapstr.M{
						"name": hostName,
					},
				},
			},
			false,
		},
		{
			"event only with host.id",
			beat.Event{
				Fields: mapstr.M{
					"host": mapstr.M{
						"id": hostID,
					},
				},
			},
			true,
		},
		{
			"event with host.name and host.id",
			beat.Event{
				Fields: mapstr.M{
					"host": mapstr.M{
						"name": hostName,
						"id":   hostID,
					},
				},
			},
			true,
		},
		{
			"event without host field",
			beat.Event{
				Fields: mapstr.M{},
			},
			false,
		},
		{
			"event with field type map[string]string hostID",
			beat.Event{
				Fields: mapstr.M{
					"host": hostIDMap,
				},
			},
			true,
		},
		{
			"event with field type map[string]string host name",
			beat.Event{
				Fields: mapstr.M{
					"host": hostNameMap,
				},
			},
			false,
		},
		{
			"event with field type map[string]string host ID and name",
			beat.Event{
				Fields: mapstr.M{
					"host": hostIDNameMap,
				},
			},
			true,
		},
		{
			"event with field type string",
			beat.Event{
				Fields: mapstr.M{
					"host": "string",
				},
			},
			false,
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			skip := skipAddingHostMetadata(&c.event)
			assert.Equal(t, c.expectedSkip, skip)
		})
	}
}

func TestExpireCacheOnFQDNReportingChange(t *testing.T) {
	testConfig := conf.MustNewConfigFrom(map[string]interface{}{
		"cache.ttl": "5m",
	})

	p, err := New(testConfig)
	require.NoError(t, err)
	defer deregisterMonitoringRegistry()

	ahmP, ok := p.(*addHostMetadata)
	require.True(t, ok)

	// Call the expired() method once to prime the cache's
	// lastUpdated value
	ahmP.expired()

	// Since we just primed the cache's lastUpdated value, the
	// cache should no longer be expired.
	expired := ahmP.expired()
	require.False(t, expired)

	// Toggle the FQDN feature flag; this should cause the cache
	// to expire.
	features.UpdateFromConfig(conf.MustNewConfigFrom(map[string]interface{}{
		"features.fqdn.enabled": true,
	}))
	expired = ahmP.expired()
	require.True(t, expired)

	// Set the FQDN feature flag to the same value; this should NOT
	// cause the cache to expire.
	features.UpdateFromConfig(conf.MustNewConfigFrom(map[string]interface{}{
		"features.fqdn.enabled": true,
	}))
	expired = ahmP.expired()
	require.False(t, expired)
}

func TestFQDNLookup(t *testing.T) {
	hostname, err := os.Hostname()
	require.NoError(t, err)

	tests := map[string]struct {
		cnameLookupResult             string
		expectedHostName              string
		expectedFQDNLookupFailedCount int64
	}{
		"lookup_succeeds": {
			cnameLookupResult:             "foo.bar.baz.",
			expectedHostName:              "foo.bar.baz",
			expectedFQDNLookupFailedCount: 0,
		},
		"lookup_fails": {
			cnameLookupResult:             "",
			expectedHostName:              hostname,
			expectedFQDNLookupFailedCount: 1,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			// Mock CNAME resolution
			srv, _ := mockdns.NewServer(map[string]mockdns.Zone{
				hostname + ".": {
					CNAME: test.cnameLookupResult,
				},
				test.cnameLookupResult: {
					A: []string{"1.1.1.1"},
				},
			}, false)
			defer srv.Close()

			srv.PatchNet(net.DefaultResolver)
			defer mockdns.UnpatchNet(net.DefaultResolver)

			// Enable FQDN feature flag
			err = features.UpdateFromConfig(fqdnFeatureFlagConfig(true))
			require.NoError(t, err)
			defer func() {
				err = features.UpdateFromConfig(fqdnFeatureFlagConfig(true))
				require.NoError(t, err)
			}()

			// Create processor and check that FQDN lookup failed
			testConfig, err := conf.NewConfigFrom(map[string]interface{}{})
			require.NoError(t, err)

			p, err := New(testConfig)
			require.NoError(t, err)
			defer deregisterMonitoringRegistry()

			addHostMetadataP, ok := p.(*addHostMetadata)
			require.True(t, ok)
			require.Equal(t, test.expectedFQDNLookupFailedCount, addHostMetadataP.metrics.FQDNLookupFailed.Get())

			// Run event through processor and check that hostname reported
			// by processor is same as OS-reported hostname
			event := &beat.Event{
				Fields:    mapstr.M{},
				Timestamp: time.Now(),
			}
			newEvent, err := p.Run(event)
			require.NoError(t, err)

			v, err := newEvent.GetValue("host.name")
			require.NoError(t, err)
			require.Equal(t, test.expectedHostName, v)
		})
	}
}

func fqdnFeatureFlagConfig(fqdnEnabled bool) *conf.C {
	return conf.MustNewConfigFrom(map[string]interface{}{
		"features.fqdn.enabled": fqdnEnabled,
	})
}
