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
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/elastic/beats/v7/libbeat/common"
)

// RemoveIndexPattern removes the index pattern entry from a given dashboard export
func RemoveIndexPattern(data []byte) ([]byte, error) {
	result := make([]byte, 0)
	r := bufio.NewReader(bytes.NewReader(data))
	for {
		line, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				return result, nil
			}
			return data, err
		}

		var r common.MapStr
		// Full struct need to not loose any data
		err = json.Unmarshal(line, &r)
		if err != nil {
			return nil, err
		}
		v, err := r.GetValue("type")
		if err != nil {
			return nil, fmt.Errorf("type key not found or not string")
		}
		if v != "index-pattern" {
			result = append(result, line...)
		}
	}
}
