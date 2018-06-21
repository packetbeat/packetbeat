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

package schema

import (
	"strings"

	"github.com/elastic/beats/libbeat/logp"
)

type Errors []Error

func NewErrors() *Errors {
	return &Errors{}
}

func (errs *Errors) AddError(err *Error) {
	*errs = append(*errs, *err)
}

func (errs *Errors) AddErrors(errors *Errors) {
	if errors == nil {
		return
	}
	*errs = append(*errs, *errors...)
}

func (errs *Errors) HasRequiredErrors() bool {
	for _, err := range *errs {
		if err.IsType(RequiredType) {
			return true
		}
	}
	return false
}

func (errs *Errors) Error() string {
	error := "Required fields are missing: "
	for _, err := range *errs {
		if err.IsType(RequiredType) {
			error = error + "," + err.key
		}
	}
	return error
}

// Log logs all missing required and optional fields to the debug log.
func (errs *Errors) Log() {
	if len(*errs) == 0 {
		return
	}
	var optional, required []string

	for _, err := range *errs {
		if err.IsType(RequiredType) {
			required = append(required, err.key)
		} else {
			optional = append(optional, err.key)
		}
	}

	log := ""
	if len(required) > 0 {
		log = log + "required: " + strings.Join(required, ",") + "; "
	}

	if len(optional) > 0 {
		log = log + "optional: " + strings.Join(optional, ",") + ";"
	}

	logp.Debug("schema", "Fields missing - %s", log)
}
