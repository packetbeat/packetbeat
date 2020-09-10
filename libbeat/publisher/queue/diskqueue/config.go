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

package diskqueue

import (
	"errors"

	"github.com/elastic/beats/v7/libbeat/common"
)

// userConfig holds the parameters for a disk queue that are configurable
// by the end user in the beats yml file.
type userConfig struct {
	Path string `config:"path"`
}

func (c *userConfig) Validate() error {
	if false {
		return errors.New("something is wrong")
	}

	return nil
}

// DefaultSettings returns a Settings object with reasonable default values
// for all important fields.
func DefaultSettings() Settings {
	return Settings{
		ChecksumType:   ChecksumTypeCRC32,
		MaxSegmentSize: 100 * (1 << 20), // 100MB
		MaxBufferSize:  (1 << 30),       // 1GB
	}
}

// SettingsForUserConfig returns a Settings struct initialized with the
// end-user-configurable settings in the given config tree.
func SettingsForUserConfig(config *common.Config) (Settings, error) {
	userConfig := userConfig{}
	if err := config.Unpack(&userConfig); err != nil {
		// TODO: report error / decide what to do
		return Settings{}, err
	}
	settings := DefaultSettings()
	settings.Path = userConfig.Path

	return settings, nil
}
