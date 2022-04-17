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

//go:build !integration
// +build !integration

package kafka

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/menderesk/beats/v7/libbeat/common"
	"github.com/menderesk/beats/v7/libbeat/tests/resources"
)

func TestNewInputDone(t *testing.T) {
	config := common.MustNewConfigFrom(common.MapStr{
		"hosts":    "localhost:9092",
		"topics":   "messages",
		"group_id": "filebeat",
	})

	AssertNotStartedInputCanBeDone(t, config)
}

// AssertNotStartedInputCanBeDone checks that the context of an input can be
// done before starting the input, and it doesn't leak goroutines. This is
// important to confirm that leaks don't happen with CheckConfig.
func AssertNotStartedInputCanBeDone(t *testing.T, configMap *common.Config) {
	goroutines := resources.NewGoroutinesChecker()
	defer goroutines.Check(t)

	config, err := common.NewConfigFrom(configMap)
	require.NoError(t, err)

	_, err = Plugin().Manager.Create(config)
	require.NoError(t, err)
}
