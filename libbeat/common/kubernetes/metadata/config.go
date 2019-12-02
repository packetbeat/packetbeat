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

package metadata

import "github.com/elastic/beats/libbeat/common"

// Config declares supported configuration for metadata generation
type Config struct {
	IncludeLabels      []string `config:"include_labels"`
	ExcludeLabels      []string `config:"exclude_labels"`
	IncludeAnnotations []string `config:"include_annotations"`

	LabelsDedot      bool `config:"labels.dedot"`
	AnnotationsDedot bool `config:"annotations.dedot"`

	// Undocumented settings, to be deprecated in favor of `drop_fields` processor:
	IncludeCreatorMetadata bool `config:"include_creator_metadata"`
}

type AddResourceMetadataConfig struct {
	Node      *common.Config `config:"config"`
	Namespace *common.Config `config:"config"`
}

func defaultConfig() Config {
	return Config{
		IncludeCreatorMetadata: true,
		LabelsDedot:            true,
		AnnotationsDedot:       true,
	}
}

func (c *Config) Unmarshal(cfg *common.Config) error {
	return cfg.Unpack(c)
}
