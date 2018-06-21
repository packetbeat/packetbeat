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

package kubernetes

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/beats/libbeat/common"
)

func TestPodMetadataDeDot(t *testing.T) {
	withPodUID, _ := common.NewConfigFrom(map[string]interface{}{"include_pod_uid": true})

	tests := []struct {
		pod    *Pod
		meta   common.MapStr
		config *common.Config
	}{
		{
			pod: &Pod{
				Metadata: ObjectMeta{
					Labels: map[string]string{"a.key": "foo", "a": "bar"},
					UID:    "005f3b90-4b9d-12f8-acf0-31020a840133",
				},
			},
			meta: common.MapStr{
				"pod":       common.MapStr{"name": ""},
				"namespace": "",
				"node":      common.MapStr{"name": ""},
				"labels":    common.MapStr{"a": common.MapStr{"value": "bar", "key": "foo"}},
			},
			config: common.NewConfig(),
		},
		{
			pod: &Pod{
				Metadata: ObjectMeta{
					Labels: map[string]string{"a.key": "foo", "a": "bar"},
					UID:    "005f3b90-4b9d-12f8-acf0-31020a840133",
				},
			},
			meta: common.MapStr{
				"pod":       common.MapStr{"name": "", "uid": "005f3b90-4b9d-12f8-acf0-31020a840133"},
				"namespace": "",
				"node":      common.MapStr{"name": ""},
				"labels":    common.MapStr{"a": common.MapStr{"value": "bar", "key": "foo"}},
			},
			config: withPodUID,
		},
		{
			pod: &Pod{
				Metadata: ObjectMeta{
					Labels: map[string]string{"a.key": "foo", "a": "bar"},
					UID:    "005f3b90-4b9d-12f8-acf0-31020a840133",
					OwnerReferences: []struct {
						APIVersion string `json:"apiVersion"`
						Controller bool   `json:"controller"`
						Kind       string `json:"kind"`
						Name       string `json:"name"`
						UID        string `json:"uid"`
					}{
						{
							Kind:       "Deployment",
							Name:       "test",
							Controller: true,
						},
						{
							Kind:       "Replicaset",
							Name:       "replicaset",
							Controller: false,
						},
					},
				},
			},
			meta: common.MapStr{
				"pod":        common.MapStr{"name": ""},
				"namespace":  "",
				"node":       common.MapStr{"name": ""},
				"labels":     common.MapStr{"a": common.MapStr{"value": "bar", "key": "foo"}},
				"deployment": common.MapStr{"name": "test"},
			},
			config: common.NewConfig(),
		},
	}

	for _, test := range tests {
		metaGen, err := NewMetaGenerator(test.config)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, metaGen.PodMetadata(test.pod), test.meta)
	}
}
