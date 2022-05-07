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

//go:build mage
// +build mage

package main

import (
	"context"
	devtools "github.com/elastic/beats/v7/dev-tools/mage"
	"github.com/elastic/beats/v7/dev-tools/mage/target/build"
	"github.com/magefile/mage/mg"

	// mage:import
	_ "github.com/elastic/beats/v7/dev-tools/mage/target/common"
	// mage:import
	"github.com/elastic/beats/v7/dev-tools/mage/target/unittest"
	// mage:import
	_ "github.com/elastic/beats/v7/dev-tools/mage/target/integtest/containers"
	// mage:import
	_ "github.com/elastic/beats/v7/dev-tools/mage/target/test"
)

func init() {
	unittest.RegisterPythonTestDeps(Fields)
}

// Build builds the Beat binary.
func Build() error {
	return devtools.Build(devtools.DefaultBuildArgs())
}

// Builds the system test binary.
func BuildSystemTestBinary() {
	mg.Deps(devtools.BuildSystemTestBinary)
}

// Fields generates a fields.yml for the Beat.
func Fields() error {
	return devtools.GenerateFieldsYAML("processors")
}

// Config generates example and reference configuration for libbeat.
func Config() error {
	return devtools.Config(devtools.ShortConfigType|devtools.ReferenceConfigType, devtools.DefaultConfigFileParams(), ".")
}

// AssembleDarwinUniversal merges the darwin/amd64 and darwin/arm64 into a single
// universal binary using `lipo`. It assumes the darwin/amd64 and darwin/arm64
// were built and only performs the merge.
func AssembleDarwinUniversal() error {
	return build.AssembleDarwinUniversal()
}

// IntegTest executes integration tests (it uses Docker to run the tests).
func IntegTest() {
	mg.SerialDeps(GoIntegTest, PythonIntegTest)
}

// GoIntegTest starts the docker containers and executes the Go integration tests.
func GoIntegTest(ctx context.Context) error {
	mg.Deps(Fields)
	args := devtools.DefaultGoTestIntegrationArgs()
	// TODO: Tidy these up, put them in the default args?
	args.Env["ES_HOST"] = "localhost"
	args.Env["ES_USER"] = "admin"
	args.Env["ES_PASS"] = "testing"
	args.Env["ES_SUPERUSER_USER"] = "admin"
	args.Env["ES_SUPERUSER_PASS"] = "testing"
	args.Env["KIBANA_HOST"] = "localhost"
	args.Env["KIBANA_USER"] = "beats"
	args.Env["KIBANA_PASS"] = "testing"
	args.Env["REDIS_HOST"] = "localhost"
	args.Env["SREDIS_HOST"] = "localhost"
	args.Env["LS_HOST"] = "localhost"
	return devtools.GoIntegTest(ctx, args)
}

// PythonIntegTest starts the docker containers and executes the Python integration tests.
// Use GENERATE=true to generate expected log files.
// Use TESTING_FILEBEAT_MODULES=module[,module] to limit what modules to test.
// Use TESTING_FILEBEAT_FILESETS=fileset[,fileset] to limit what fileset to test.
func PythonIntegTest(ctx context.Context) error {
	mg.Deps(Fields, BuildSystemTestBinary)
	args := devtools.DefaultPythonTestIntegrationArgs()
	// TODO: Tidy these up?
	args.Env["INTEGRATION_TESTS"] = "1"
	args.Env["MODULES_PATH"] = devtools.CWD("module")
	args.Env["ES_HOST"] = "localhost"
	args.Env["ES_USER"] = "admin"
	args.Env["ES_PASS"] = "testing"
	args.Env["ES_SUPERUSER_USER"] = "admin"
	args.Env["ES_SUPERUSER_PASS"] = "testing"
	args.Env["KIBANA_HOST"] = "localhost"
	args.Env["KIBANA_USER"] = "beats"
	args.Env["KIBANA_PASS"] = "testing"
	args.Env["REDIS_HOST"] = "localhost"
	args.Env["SREDIS_HOST"] = "localhost"
	args.Env["LS_HOST"] = "localhost"
	return devtools.PythonIntegTest(args)
}
