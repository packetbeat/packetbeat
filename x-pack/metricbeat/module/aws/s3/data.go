// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package s3

import (
	"github.com/elastic/beats/libbeat/common"
	s "github.com/elastic/beats/libbeat/common/schema"
	c "github.com/elastic/beats/libbeat/common/schema/mapstrstr"
)

var (
	schemaMetricSetFields = s.Schema{
		"bucket": s.Object{
			"size": s.Object{
				"bytes": c.Float("bucket.size.bytes", s.Optional),
			},
		},
		"object": s.Object{
			"count": c.Float("object.count", s.Optional),
		},
	}
)

var (
	schemaRootFields = s.Schema{
		"service": s.Object{
			"name": c.Str("service.name", s.Optional),
		},
		"cloud": s.Object{
			"provider": c.Str("cloud.provider", s.Optional),
			"region": c.Str("cloud.region", s.Optional),
		},
	}
)

func eventMapping(input map[string]interface{}, schema s.Schema) (common.MapStr, error) {
	return schema.Apply(input)
}
