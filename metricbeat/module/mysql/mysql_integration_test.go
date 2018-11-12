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

// +build integration

package mysql

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/beats/libbeat/tests/compose"
	"github.com/elastic/beats/metricbeat/module/mysql/mtest"
)

func TestNewDB(t *testing.T) {
	mtest.Runner.Run(t, compose.Suite{
		"NewDB": func(t *testing.T, r compose.R) {
			db, err := NewDB(mtest.GetDSN(r.Host()))
			assert.NoError(t, err)

			err = db.Ping()
			assert.NoError(t, err)
		},
	})
}
