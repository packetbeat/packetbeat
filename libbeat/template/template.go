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

package template

import (
	"fmt"
	"sync"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/common/cfgwarn"
	"github.com/elastic/beats/libbeat/common/fmtstr"
	"github.com/elastic/go-ucfg/yaml"
)

var (
	// Array to store dynamicTemplate parts in
	dynamicTemplates []common.MapStr

	defaultFields []string
)

//Template holds information about an ES template
type Template struct {
	sync.Mutex
	name        string
	pattern     string
	beatVersion common.Version
	esVersion   common.Version
	config      Config
}

// New creates a new template instance
func New(beatVersion string, beatName string, esVersion common.Version, config Config) (*Template, error) {
	bV, err := common.NewVersion(beatVersion)
	if err != nil {
		return nil, err
	}

	event := event(beatName, bV.String())
	name := config.Name
	if config.JSON.Enabled {
		name = config.JSON.Name
	}
	name, err = runFormatter(name, event)
	if err != nil {
		return nil, err
	}
	pattern, err := runFormatter(config.Pattern, event)
	if err != nil {
		return nil, err
	}

	// update lifecycle entries
	if config.Settings.Index != nil {
		if lifecycle, ok := config.Settings.Index["lifecycle"].(map[string]interface{}); ok {
			if alias, ok := lifecycle["rollover_alias"].(string); ok {
				alias, err = runFormatter(alias, event)
				if err != nil {
					return nil, err
				}
				config.Settings.Index["lifecycle"].(map[string]interface{})["rollover_alias"] = alias
			}
			if name, ok := lifecycle["name"].(string); ok {
				name, err = runFormatter(name, event)
				if err != nil {
					return nil, err
				}
				config.Settings.Index["lifecycle"].(map[string]interface{})["name"] = name
			}
		}
	}

	// In case no esVersion is set, it is assumed the same as beat version
	if !esVersion.IsValid() {
		esVersion = *bV
	}

	return &Template{
		pattern:     pattern,
		name:        name,
		beatVersion: *bV,
		esVersion:   esVersion,
		config:      config,
	}, nil
}

func event(name, version string) *beat.Event {
	return &beat.Event{
		Fields: common.MapStr{
			// beat object was left in for backward compatibility reason for older configs.
			"beat": common.MapStr{
				"name":    name,
				"version": version,
			},
			"agent": common.MapStr{
				"name":    name,
				"version": version,
			},
			// For the Beats that have an observer role
			"observer": common.MapStr{
				"name":    name,
				"version": version,
			},
		},
		Timestamp: time.Now(),
	}
}

func (t *Template) load(fields common.Fields) (common.MapStr, error) {

	// Locking to make sure dynamicTemplates and defaultFields is not accessed in parallel
	t.Lock()
	defer t.Unlock()

	dynamicTemplates = nil
	defaultFields = nil

	var err error
	if len(t.config.AppendFields) > 0 {
		cfgwarn.Experimental("append_fields is used.")
		fields, err = appendFields(fields, t.config.AppendFields)
		if err != nil {
			return nil, err
		}
	}

	// Start processing at the root
	properties := common.MapStr{}
	processor := Processor{EsVersion: t.esVersion}
	if err := processor.Process(fields, "", properties); err != nil {
		return nil, err
	}
	output := t.Generate(properties, dynamicTemplates)

	return output, nil
}

// LoadFile loads the the template from the given file path
func (t *Template) LoadFile(file string) (common.MapStr, error) {

	fields, err := common.LoadFieldsYaml(file)
	if err != nil {
		return nil, err
	}

	return t.load(fields)
}

// LoadBytes loads the the template from the given byte array
func (t *Template) LoadBytes(data []byte) (common.MapStr, error) {
	fields, err := loadYamlByte(data)
	if err != nil {
		return nil, err
	}

	return t.load(fields)
}

// GetName returns the name of the template
func (t *Template) GetName() string {
	return t.name
}

// GetPattern returns the pattern of the template
func (t *Template) GetPattern() string {
	return t.pattern
}

// Generate generates the full template
// The default values are taken from the default variable.
func (t *Template) Generate(properties common.MapStr, dynamicTemplates []common.MapStr) common.MapStr {
	// Add base dynamic template
	var dynamicTemplateBase = common.MapStr{
		"strings_as_keyword": common.MapStr{
			"mapping": common.MapStr{
				"ignore_above": 1024,
				"type":         "keyword",
			},
			"match_mapping_type": "string",
		},
	}

	if t.esVersion.IsMajor(2) {
		dynamicTemplateBase.Put("strings_as_keyword.mapping.type", "string")
		dynamicTemplateBase.Put("strings_as_keyword.mapping.index", "not_analyzed")
	}

	dynamicTemplates = append(dynamicTemplates, dynamicTemplateBase)

	indexSettings := common.MapStr{
		"refresh_interval": "5s",
		"mapping": common.MapStr{
			"total_fields": common.MapStr{
				"limit": defaultTotalFieldsLimit,
			},
		},
	}

	// number_of_routing shards is only supported for ES version >= 6.1
	version61, _ := common.NewVersion("6.1.0")
	if !t.esVersion.LessThan(version61) {
		indexSettings.Put("number_of_routing_shards", defaultNumberOfRoutingShards)
	}

	mappingName := "_doc"
	major := t.esVersion.Major
	switch {
	case major < 6:
		mappingName = "_default_"
	case major >= 7:
		defaultFields = append(defaultFields, "fields.*")
		indexSettings.Put("query.default_field", defaultFields)
	}

	indexSettings.DeepUpdate(t.config.Settings.Index)

	// Load basic structure
	basicStructure := common.MapStr{
		"mappings": common.MapStr{
			mappingName: common.MapStr{
				"_meta": common.MapStr{
					"version": t.beatVersion.String(),
				},
				"date_detection":    defaultDateDetection,
				"dynamic_templates": dynamicTemplates,
				"properties":        properties,
			},
		},
		"order": 1,
		"settings": common.MapStr{
			"index": indexSettings,
		},
	}

	if len(t.config.Settings.Source) > 0 {
		key := fmt.Sprintf("mappings.%s._source", mappingName)
		basicStructure.Put(key, t.config.Settings.Source)
	}

	// ES 6 moved from template to index_patterns: https://github.com/elastic/elasticsearch/pull/21009
	if t.esVersion.Major >= 6 {
		basicStructure.Put("index_patterns", []string{t.GetPattern()})
	} else {
		basicStructure.Put("template", t.GetPattern())
	}

	if t.esVersion.IsMajor(2) {
		basicStructure.Put("mappings._default_._all.norms.enabled", false)
	}

	return basicStructure
}

func appendFields(fields, appendFields common.Fields) (common.Fields, error) {
	if len(appendFields) > 0 {
		appendFieldKeys := appendFields.GetKeys()

		// Append is only allowed to add fields, not overwrite
		for _, key := range appendFieldKeys {
			if fields.HasNode(key) {
				return nil, fmt.Errorf("append_fields contains an already existing key: %s", key)
			}
		}
		// Appends fields to existing fields
		fields = append(fields, appendFields...)
	}
	return fields, nil
}

func loadYamlByte(data []byte) (common.Fields, error) {

	var keys []common.Field

	cfg, err := yaml.NewConfig(data)
	if err != nil {
		return nil, err
	}
	cfg.Unpack(&keys)

	fields := common.Fields{}

	for _, key := range keys {
		fields = append(fields, key.Fields...)
	}
	return fields, nil
}

func runFormatter(s string, event *beat.Event) (string, error) {
	formatter, err := fmtstr.CompileEvent(s)
	if err != nil {
		return s, err
	}
	return formatter.Run(event)
}
