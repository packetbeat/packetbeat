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

import (
	"strings"

	"k8s.io/apimachinery/pkg/api/meta"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/common/kubernetes"
	"github.com/elastic/beats/libbeat/common/safemapstr"
)

type resource struct {
	config *Config
}

func NewResourceMetadataGenerator(config *Config) *resource {
	return &resource{
		config: config,
	}
}

func (r *resource) Generate(obj kubernetes.Resource, options ...FieldOptions) common.MapStr {
	accessor, err := meta.Accessor(obj)
	if err != nil {
		return nil
	}

	labelMap := common.MapStr{}
	if len(r.config.IncludeLabels) == 0 {
		for k, v := range accessor.GetLabels() {
			if r.config.LabelsDedot {
				label := common.DeDot(k)
				labelMap.Put(label, v)
			} else {
				safemapstr.Put(labelMap, k, v)
			}
		}
	} else {
		labelMap = generateMapSubset(accessor.GetLabels(), r.config.IncludeLabels, g.LabelsDedot)
	}

	// Exclude any labels that are present in the exclude_labels config
	for _, label := range r.config.ExcludeLabels {
		labelMap.Delete(label)
	}

	annotationsMap := generateMapSubset(accessor.GetAnnotations(), r.config.IncludeAnnotations, g.AnnotationsDedot)
	meta := common.MapStr{}
	if accessor.GetNamespace() != "" {
		safemapstr.Put(meta, "namespace.name", accessor.GetNamespace())
	}

	// Add controller metadata if present
	if r.config.IncludeCreatorMetadata {
		for _, ref := range accessor.GetOwnerReferences() {
			if ref.Controller != nil && *ref.Controller {
				switch ref.Kind {
				// TODO grow this list as we keep adding more `state_*` metricsets
				case "Deployment",
					"ReplicaSet",
					"StatefulSet":
					safemapstr.Put(meta, strings.ToLower(ref.Kind)+".name", ref.Name)
				}
			}
		}
	}

	if len(labelMap) != 0 {
		meta["labels"] = labelMap
	}

	if len(annotationsMap) != 0 {
		meta["annotations"] = annotationsMap
	}

	return meta
}

func generateMapSubset(input map[string]string, keys []string, dedot bool) common.MapStr {
	output := common.MapStr{}
	if input == nil {
		return output
	}

	for _, key := range keys {
		value, ok := input[key]
		if ok {
			if dedot {
				dedotKey := common.DeDot(key)
				output.Put(dedotKey, value)
			} else {
				safemapstr.Put(output, key, value)
			}
		}
	}

	return output
}
