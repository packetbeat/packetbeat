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

package gotool

import (
	"github.com/magefile/mage/sh"
)

type goLicenser func(opts ...ArgOpt) error

// Test runs `go-licenser` and provides optionals for adding command line arguments.
var Licenser goLicenser = runGoLicenser

func runGoLicenser(opts ...ArgOpt) error {
	args := buildArgs(opts).build()
	return sh.RunV("go-licenser", args...)
}

func (goLicenser) Check() ArgOpt                 { return flagBoolIf("-d", true) }
func (goLicenser) License(license string) ArgOpt { return flagArgIf("-license", license) }
func (goLicenser) Exclude(path string) ArgOpt    { return flagArgIf("-exclude", path) }
func (goLicenser) Path(path string) ArgOpt       { return posArg(path) }
