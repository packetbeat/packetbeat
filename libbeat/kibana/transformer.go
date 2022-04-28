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

package kibana

import (
	"github.com/elastic/elastic-agent-libs/mapstr"
	"github.com/elastic/go-ucfg/yaml"
)

type transformer struct {
	entries []kibanaEntry
}
type kibanaEntry struct {
	Kibana struct {
		SourceFilters []string `config:"source_filters"`
	} `config:"kibana"`
}

func newTransformer(entries []kibanaEntry) *transformer {
	return &transformer{entries: entries}
}

func (t *transformer) transform() mapstr.M {
	transformed := mapstr.M{}

	var srcFilters []mapstr.M
	for _, entry := range t.entries {
		for _, sourceFilter := range entry.Kibana.SourceFilters {
			srcFilters = append(srcFilters, mapstr.M{"value": sourceFilter})
		}
	}
	if len(srcFilters) > 0 {
		transformed["sourceFilters"] = srcFilters
	}
	return transformed
}

func loadKibanaEntriesFromYaml(fields []byte) ([]kibanaEntry, error) {
	entries := []kibanaEntry{}
	cfg, err := yaml.NewConfig(fields)
	if err != nil {
		return nil, err
	}
	err = cfg.Unpack(&entries)
	if err != nil {
		return nil, err
	}
	return entries, nil
}
