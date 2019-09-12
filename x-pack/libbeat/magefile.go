// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// +build mage

package main

import (
	"context"

	"github.com/magefile/mage/mg"

	devtools "github.com/elastic/beats/dev-tools/mage"

	// mage:import
	_ "github.com/elastic/beats/dev-tools/mage/target/common"
)

func init() {
	devtools.BeatLicense = "Elastic License"
}

// Build builds the Beat binary.
func Build() error {
	return devtools.Build(devtools.DefaultBuildArgs())
}

// Fields generates a fields.yml for the Beat.
func Fields() error {
	return devtools.GenerateFieldsYAML()
}

// GoTestUnit executes the Go unit tests.
// Use TEST_COVERAGE=true to enable code coverage profiling.
// Use RACE_DETECTOR=true to enable the race detector.
func GoTestUnit(ctx context.Context) error {
	return devtools.GoTest(ctx, devtools.DefaultGoTestUnitArgs())
}

// GoTestIntegration executes the Go integration tests.
// Use TEST_COVERAGE=true to enable code coverage profiling.
// Use RACE_DETECTOR=true to enable the race detector.
func GoTestIntegration(ctx context.Context) error {
	return devtools.GoTest(ctx, devtools.DefaultGoTestIntegrationArgs())
}

// Config generates example and reference configuration for libbeat.
func Config() error {
	return devtools.Config(devtools.ShortConfigType|devtools.ReferenceConfigType, devtools.ConfigFileParams{}, ".")
}

// AddLicenseHeaders adds license headers
func AddLicenseHeaders() {
	mg.Deps(devtools.AddLicenseHeaders)
}

// CheckLicenseHeaders checks license headers
func CheckLicenseHeaders() {
	mg.Deps(devtools.CheckLicenseHeaders)
}
