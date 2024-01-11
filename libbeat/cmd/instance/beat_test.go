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

//go:build !integration

package instance

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/elastic/beats/v7/libbeat/cfgfile"
	"github.com/elastic/beats/v7/libbeat/common/reload"
	"github.com/elastic/beats/v7/libbeat/outputs"
	"github.com/elastic/beats/v7/libbeat/publisher/queue/memqueue"
	"github.com/elastic/elastic-agent-libs/config"
	"github.com/elastic/go-ucfg/yaml"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewInstance(t *testing.T) {
	b, err := NewBeat("testbeat", "testidx", "0.9", false)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "testbeat", b.Info.Beat)
	assert.Equal(t, "testidx", b.Info.IndexPrefix)
	assert.Equal(t, "0.9", b.Info.Version)

	// UUID4 should be 36 chars long
	assert.Equal(t, 16, len(b.Info.ID))
	assert.Equal(t, 36, len(b.Info.ID.String()))

	// indexPrefix set to name if empty
	b, err = NewBeat("testbeat", "", "0.9", false)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "testbeat", b.Info.Beat)
	assert.Equal(t, "testbeat", b.Info.IndexPrefix)

}

func TestNewInstanceUUID(t *testing.T) {
	b, err := NewBeat("testbeat", "", "0.9", false)
	if err != nil {
		panic(err)
	}

	// Make sure the ID's are different
	differentUUID, err := uuid.NewV4()
	if err != nil {
		t.Fatalf("error while generating ID: %v", err)
	}
	assert.NotEqual(t, b.Info.ID, differentUUID)
}

func TestInitKibanaConfig(t *testing.T) {
	b, err := NewBeat("filebeat", "testidx", "0.9", false)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "filebeat", b.Info.Beat)
	assert.Equal(t, "testidx", b.Info.IndexPrefix)
	assert.Equal(t, "0.9", b.Info.Version)

	const configPath = "../test/filebeat_test.yml"

	// Ensure that the config has owner-exclusive write permissions.
	// This is necessary on some systems which have a default umask
	// of 0o002, meaning that files are checked out by git with mode
	// 0o664. This would cause cfgfile.Load to fail.
	err = os.Chmod(configPath, 0o644)
	assert.NoError(t, err)

	cfg, err := cfgfile.Load(configPath, nil)
	assert.NoError(t, err)
	err = cfg.Unpack(&b.Config)
	assert.NoError(t, err)

	kibanaConfig := InitKibanaConfig(b.Config)
	username, err := kibanaConfig.String("username", -1)
	assert.NoError(t, err)
	password, err := kibanaConfig.String("password", -1)
	assert.NoError(t, err)
	api_key, err := kibanaConfig.String("api_key", -1)
	assert.NoError(t, err)
	protocol, err := kibanaConfig.String("protocol", -1)
	assert.NoError(t, err)
	host, err := kibanaConfig.String("host", -1)
	assert.NoError(t, err)

	assert.Equal(t, "elastic-test-username", username)
	assert.Equal(t, "elastic-test-password", password)
	assert.Equal(t, "elastic-test-api-key", api_key)
	assert.Equal(t, "https", protocol)
	assert.Equal(t, "127.0.0.1:5601", host)
}

func TestEmptyMetaJson(t *testing.T) {
	b, err := NewBeat("filebeat", "testidx", "0.9", false)
	if err != nil {
		panic(err)
	}

	// prepare empty meta file
	metaFile, err := ioutil.TempFile("../test", "meta.json")
	assert.Equal(t, nil, err, "Unable to create temporary meta file")

	metaPath := metaFile.Name()
	metaFile.Close()
	defer os.Remove(metaPath)

	// load metadata
	err = b.loadMeta(metaPath)

	assert.Equal(t, nil, err, "Unable to load meta file properly")
	assert.NotEqual(t, uuid.Nil, b.Info.ID, "Beats UUID is not set")
}

func TestMetaJsonWithTimestamp(t *testing.T) {
	firstBeat, err := NewBeat("filebeat", "testidx", "0.9", false)
	if err != nil {
		panic(err)
	}
	firstStart := firstBeat.Info.FirstStart

	metaFile, err := ioutil.TempFile("../test", "meta.json")
	assert.Equal(t, nil, err, "Unable to create temporary meta file")

	metaPath := metaFile.Name()
	metaFile.Close()
	defer os.Remove(metaPath)

	err = firstBeat.loadMeta(metaPath)
	assert.Equal(t, nil, err, "Unable to load meta file properly")

	secondBeat, err := NewBeat("filebeat", "testidx", "0.9", false)
	if err != nil {
		panic(err)
	}
	assert.False(t, firstStart.Equal(secondBeat.Info.FirstStart), "Before meta.json is loaded, first start must be different")
	err = secondBeat.loadMeta(metaPath)
	require.NoError(t, err)

	assert.Equal(t, nil, err, "Unable to load meta file properly")
	assert.True(t, firstStart.Equal(secondBeat.Info.FirstStart), "Cannot load first start")
}

func TestSanitizeIPs(t *testing.T) {
	cases := []struct {
		name        string
		ips         []string
		expectedIPs []string
	}{
		{
			name: "does not change valid IPs",
			ips: []string{
				"127.0.0.1",
				"::1",
				"fe80::1",
				"fe80::6ca6:cdff:fe6a:4f59",
				"192.168.1.101",
			},
			expectedIPs: []string{
				"127.0.0.1",
				"::1",
				"fe80::1",
				"fe80::6ca6:cdff:fe6a:4f59",
				"192.168.1.101",
			},
		},
		{
			name: "cuts the masks",
			ips: []string{
				"127.0.0.1/8",
				"::1/128",
				"fe80::1/64",
				"fe80::6ca6:cdff:fe6a:4f59/64",
				"192.168.1.101/24",
			},
			expectedIPs: []string{
				"127.0.0.1",
				"::1",
				"fe80::1",
				"fe80::6ca6:cdff:fe6a:4f59",
				"192.168.1.101",
			},
		},
		{
			name: "excludes invalid IPs",
			ips: []string{
				"",
				"fe80::6ca6:cdff:fe6a:4f59",
				"invalid",
				"192.168.1.101",
			},
			expectedIPs: []string{
				"fe80::6ca6:cdff:fe6a:4f59",
				"192.168.1.101",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expectedIPs, sanitizeIPs(tc.ips))
		})
	}
}

func TestReloader(t *testing.T) {
	t.Run("updates the output configuration on the beat", func(t *testing.T) {
		b, err := NewBeat("testbeat", "testidx", "0.9", false)
		require.NoError(t, err)

		cfg := `
elasticsearch:
  hosts: ["https://127.0.0.1:9200"]
  username: "elastic"
  allow_older_versions: false
`
		c, err := config.NewConfigWithYAML([]byte(cfg), cfg)
		require.NoError(t, err)
		outCfg, err := c.Child("elasticsearch", -1)
		require.NoError(t, err)

		update := &reload.ConfigWithMeta{Config: c}
		m := &outputReloaderMock{}
		reloader := b.makeOutputReloader(m)

		require.False(t, b.Config.Output.IsSet(), "the output should not be set yet")
		require.True(t, b.isConnectionToOlderVersionAllowed(), "allow_older_versions flag should be true from 8.11")
		err = reloader.Reload(update)
		require.NoError(t, err)
		require.True(t, b.Config.Output.IsSet(), "now the output should be set")
		require.Equal(t, outCfg, b.Config.Output.Config())
		require.Same(t, c, m.cfg.Config)
		require.False(t, b.isConnectionToOlderVersionAllowed(), "allow_older_versions flag should now be set to false")
	})
}

type outputReloaderMock struct {
	cfg *reload.ConfigWithMeta
}

func (r *outputReloaderMock) Reload(
	cfg *reload.ConfigWithMeta,
	factory func(o outputs.Observer, cfg config.Namespace) (outputs.Group, error),
) error {
	r.cfg = cfg
	return nil
}

func TestPromoteOutputQueueSettings(t *testing.T) {
	tests := map[string]struct {
		input     []byte
		memEvents int
	}{
		"blank": {
			input:     []byte(""),
			memEvents: 3200,
		},
		"defaults": {
			input: []byte(`
name: mockbeat
output:
  elasticsearch:
    hosts:
      - "localhost:9200"
`),
			memEvents: 3200,
		},
		"topLevelQueue": {
			input: []byte(`
name: mockbeat
queue:
  mem:
    events: 8096
output:
  elasticsearch:
    hosts:
      - "localhost:9200"
`),
			memEvents: 8096,
		},
		"outputLevelQueue": {
			input: []byte(`
name: mockbeat
output:
  elasticsearch:
    hosts:
      - "localhost:9200"
    queue:
      mem:
        events: 8096
`),
			memEvents: 8096,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			cfg, err := yaml.NewConfig(tc.input)
			require.NoError(t, err)

			config := beatConfig{}
			err = cfg.Unpack(&config)
			require.NoError(t, err)

			err = promoteOutputQueueSettings(&config)
			require.NoError(t, err)

			ms, err := memqueue.SettingsForUserConfig(config.Pipeline.Queue.Config())
			require.NoError(t, err)
			require.Equalf(t, tc.memEvents, ms.Events, "config was: %v", config.Pipeline.Queue.Config())
		})
	}
}

func TestValidateBeatConfig(t *testing.T) {
	tests := map[string]struct {
		input                 []byte
		expectValidationError string
	}{
		"blank": {
			input:                 []byte(""),
			expectValidationError: "",
		},
		"defaults": {
			input: []byte(`
name: mockbeat
output:
  elasticsearch:
    hosts:
      - "localhost:9200"
`),
			expectValidationError: "",
		},
		"topAndOutputLevelQueue": {
			input: []byte(`
name: mockbeat
queue:
  mem:
    events: 2048
output:
  elasticsearch:
    hosts:
      - "localhost:9200"
    queue:
      mem:
        events: 8096
`),
			expectValidationError: "top level queue and output level queue settings defined, only one is allowed accessing config",
		},
		"managementTopLevelDiskQueue": {
			input: []byte(`
name: mockbeat
management:
  enabled: true
queue:
  disk:
    max_size: 1G
output:
  elasticsearch:
    hosts:
      - "localhost:9200"
`),
			expectValidationError: "disk queue is not supported when management is enabled accessing config",
		},
		"managementOutputLevelDiskQueue": {
			input: []byte(`
name: mockbeat
management:
  enabled: true
output:
  elasticsearch:
    hosts:
      - "localhost:9200"
    queue:
      disk:
        max_size: 1G
`),
			expectValidationError: "disk queue is not supported when management is enabled accessing config",
		},
		"managementFalseOutputLevelDiskQueue": {
			input: []byte(`
name: mockbeat
management:
  enabled: false
output:
  elasticsearch:
    hosts:
      - "localhost:9200"
    queue:
      disk:
        max_size: 1G
`),
			expectValidationError: "",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			cfg, err := yaml.NewConfig(tc.input)
			require.NoError(t, err)
			config := beatConfig{}
			err = cfg.Unpack(&config)
			if tc.expectValidationError != "" {
				require.Error(t, err)
				require.Equal(t, tc.expectValidationError, err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestBeat_initConfigManager(t *testing.T) {
	b, err := NewBeat("testbeat", "", "0.9", false)
	require.NoError(t, err, "could not create a new beat")

	err = b.initConfigManager()
	require.NoError(t, err, "failed initialising config manager")

	assert.Empty(t, b.Info.Version, "config manager initialization"+
		"should have set b.Info.Version to empty string")
}
