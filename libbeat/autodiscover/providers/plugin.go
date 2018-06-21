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

package providers

import (
	"errors"

	"github.com/elastic/beats/libbeat/autodiscover"
	p "github.com/elastic/beats/libbeat/plugin"
)

type providerPlugin struct {
	name     string
	provider autodiscover.ProviderBuilder
}

var pluginKey = "libbeat.autodiscover.provider"

// Plugin accepts a ProviderBuilder to be registered as a plugin
func Plugin(name string, provider autodiscover.ProviderBuilder) map[string][]interface{} {
	return p.MakePlugin(pluginKey, providerPlugin{name, provider})
}

func init() {
	p.MustRegisterLoader(pluginKey, func(ifc interface{}) error {
		prov, ok := ifc.(providerPlugin)
		if !ok {
			return errors.New("plugin does not match processor plugin type")
		}

		return autodiscover.Registry.AddProvider(prov.name, prov.provider)
	})
}
