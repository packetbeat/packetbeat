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

package cfgfile

import (
	"errors"
	"fmt"
	"testing"

	"github.com/elastic/beats/v7/libbeat/common"
)

func TestIsReloadable(t *testing.T) {
	testCases := map[string]struct {
		err      error
		expected bool
	}{
		"simple error": {
			err:      errors.New("a generic error"),
			expected: true,
		},
		"common.ErrNonReloadable": {
			err:      common.ErrNonReloadable{},
			expected: false,
		},
		"wrapped common.ErrNonReloadable": {
			err:      fmt.Errorf("wrapping %w", common.ErrNonReloadable{}),
			expected: false,
		},
		"errors.Join, all errors are common.ErrNonReloadable": {
			err:      errors.Join(common.ErrNonReloadable{}, common.ErrNonReloadable{}),
			expected: false,
		},
		"errors.Join, only one is common.ErrNonReloadable": {
			err:      errors.Join(errors.New("generic reloadable error"), common.ErrNonReloadable{}),
			expected: true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			reloadable := isReloadable(tc.err)
			if reloadable != tc.expected {
				t.Errorf(
					"expecting isReloadable to return %t, but got %t for: '%s'",
					tc.expected,
					reloadable,
					tc.err,
				)
			}
		})
	}
}
