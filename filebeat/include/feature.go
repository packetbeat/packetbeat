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

package include

import (
	"github.com/elastic/beats/filebeat/input/docker"
	"github.com/elastic/beats/filebeat/input/log"
	"github.com/elastic/beats/filebeat/input/redis"
	"github.com/elastic/beats/filebeat/input/stdin"
	"github.com/elastic/beats/filebeat/input/syslog"
	"github.com/elastic/beats/filebeat/input/tcp"
	"github.com/elastic/beats/filebeat/input/udp"
	"github.com/elastic/beats/filebeat/processor/add_kubernetes_metadata"
	"github.com/elastic/beats/libbeat/feature"
)

var bundle = feature.MustBundle(
	// inputs
	feature.MustBundle(
		docker.Feature,
		log.Feature,
		redis.Feature,
		stdin.Feature,
		syslog.Feature,
		tcp.Feature,
		udp.Feature,
	),

	// processors
	feature.MustBundle(
		add_kubernetes_metadata.Feature,
	),
)

func init() {
	feature.MustOverwriteBundle(bundle)
}
