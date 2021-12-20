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

package useragent

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/elastic/beats/v7/libbeat/version"
)

// UserAgent takes the capitalized name of the current beat and returns
// an RFC compliant user agent string for that beat.
func UserAgent(beatNameCapitalized string, additional_agents ...string) string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "Elastic-%s/%s ", beatNameCapitalized, version.GetDefaultVersion())
	uaValues := []string{
		runtime.GOOS,
		runtime.GOARCH,
		version.Commit(),
		version.BuildTime().String(),
	}
	if additional_agents != nil {
		for _, val := range additional_agents {
			if val != "" {
				uaValues = append(uaValues, val)
			}
		}
	}
	fmt.Fprintf(&builder, "(%s)", strings.Join(uaValues, "; "))
	return builder.String()
}
