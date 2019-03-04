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

// Code generated by scripts/gocodegen.go - DO NOT EDIT.

package ecs

// Meta-information specific to ECS.
type ECS struct {
	// ECS version this event conforms to. `ecs.version` is a required field
	// and must exist in all events.
	// When querying across multiple indices -- which may conform to slightly
	// different ECS versions -- this field lets integrations adjust to the
	// schema version of the events.
	// The current version is 1.0.0-beta2 .
	Version string `ecs:"version"`
}
