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

package monitors

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/plugin"
)

type pluginBuilder struct {
	name    string
	typ     Type
	builder PluginBuilder
}

var pluginKey = "heartbeat.monitor"

func init() {
	plugin.MustRegisterLoader(pluginKey, func(ifc interface{}) error {
		p, ok := ifc.(pluginBuilder)
		if !ok {
			return errors.New("plugin does not match monitor plugin type")
		}

		return globalPluginsReg.register(pluginBuilder{p.name, p.typ, p.builder})
	})
}

// PluginBuilder is the signature of functions used to build active
// monitors
type PluginBuilder func(string, *common.Config) ([]Job, error)

// Type represents whether a plugin is active or passive.
type Type uint8

const (
	// ActiveMonitor represents monitors that reach across the network to do things.
	ActiveMonitor Type = iota + 1
	// PassiveMonitor represents monitors that receive inbound data.
	PassiveMonitor
)

// globalPluginsReg maintains the canonical list of valid Heartbeat monitors at runtime.
var globalPluginsReg = newPluginsReg()

type pluginsReg struct {
	monitors map[string]pluginBuilder
}

func newPluginsReg() *pluginsReg {
	return &pluginsReg{
		monitors: map[string]pluginBuilder{},
	}
}

// RegisterActive registers a new active (as opposed to passive) monitor.
func RegisterActive(name string, builder PluginBuilder) {
	if err := globalPluginsReg.add(pluginBuilder{name, ActiveMonitor, builder}); err != nil {
		panic(err)
	}
}

// ErrPluginAlreadyExists is returned when there is an attempt to register two plugins
// with the same name.
type ErrPluginAlreadyExists pluginBuilder

func (m ErrPluginAlreadyExists) Error() string {
	return fmt.Sprintf("monitor plugin '%s' already exists", m.name)
}

func (r *pluginsReg) add(plugin pluginBuilder) error {
	if _, exists := r.monitors[plugin.name]; exists {
		return ErrPluginAlreadyExists(plugin)
	}
	r.monitors[plugin.name] = plugin
	return nil
}

func (r *pluginsReg) register(plugin pluginBuilder) error {
	if _, found := r.monitors[plugin.name]; found {
		return fmt.Errorf("monitor type %v already exists", plugin.name)
	}

	r.monitors[plugin.name] = plugin

	return nil
}

func (r *pluginsReg) get(name string) (pluginBuilder, bool) {
	e, found := r.monitors[name]
	return e, found
}

func (r *pluginsReg) String() string {
	var monitors []string
	for m := range r.monitors {
		monitors = append(monitors, m)
	}
	sort.Strings(monitors)

	return fmt.Sprintf("globalPluginsReg, monitor: %v",
		strings.Join(monitors, ", "))
}
func (r *pluginsReg) monitorNames() []string {
	names := make([]string, 0, len(r.monitors))
	for k := range r.monitors {
		names = append(names, k)
	}
	return names
}

func (e *pluginBuilder) create(cfg *common.Config) ([]Job, error) {
	return e.builder(e.name, cfg)
}

func (t Type) String() string {
	switch t {
	case ActiveMonitor:
		return "active"
	case PassiveMonitor:
		return "passive"
	default:
		return "unknown type"
	}
}
