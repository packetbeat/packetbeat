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

// Code generated by "stringer -type=Stability"; DO NOT EDIT.

package feature

import (
	"strconv"
)

const _Stability_name = "UndefinedStableBetaExperimental"

var _Stability_index = [...]uint8{0, 9, 15, 19, 31}

func (i Stability) String() string {
	if i < 0 || i >= Stability(len(_Stability_index)-1) {
		return "Stability(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Stability_name[_Stability_index[i]:_Stability_index[i+1]]
}
