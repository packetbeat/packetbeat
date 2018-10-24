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

// +build mage

package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/pkg/errors"

	"github.com/elastic/beats/dev-tools/mage"
)

func init() {
	mage.BeatDescription = "Journalbeat ships systemd journal entries to Elasticsearch or Logstash."

	mage.Platforms = mage.Platforms.Filter("linux !linux/ppc64 !linux/mips64")
}

var (
	deps = map[string]func() error{
		"linux/386":      installLinux386,
		"linux/amd64":    installDefaultLinux,
		"linux/arm64":    installLinuxArm64,
		"linux/armv5":    installLinuxArmLe,
		"linux/armv6":    installLinuxArmLe,
		"linux/armv7":    installLinuxArmHf,
		"linux/mips":     installLinuxMips,
		"linux/mipsle":   installLinuxMipsle,
		"linux/mips64le": installLinuxMips64le,
		"linux/ppc64le":  installLinuxPpc64le,
		"linux/s390x":    installLinuxs390x,

		// No deb packages
		//"linux/ppc64": installLinuxPpc64,
		//"linux/mips64": installLinuxMips64,
	}
)

// Build builds the Beat binary.
func Build() error {
	return mage.Build(mage.DefaultBuildArgs())
}

// GolangCrossBuild build the Beat binary inside of the golang-builder.
// Do not use directly, use crossBuild instead.
func GolangCrossBuild() error {
	if d, ok := deps[mage.Platform.Name]; ok {
		mg.Deps(d)
	}
	return mage.GolangCrossBuild(mage.DefaultGolangCrossBuildArgs())
}

func installDefaultLinux() error {
	return installDependencies("libsystemd-dev", "")
}

func installLinuxArm64() error {
	return installDependencies("libsystemd-dev:arm64", "arm64")
}

func installLinuxArmHf() error {
	return installDependencies("libsystemd-dev:armhf", "armhf")
}

func installLinuxArmLe() error {
	return installDependencies("libsystemd-dev:armel", "armel")
}

func installLinux386() error {
	return installDependencies("libsystemd-journal-dev:i386", "i386")
}

func installLinuxMips() error {
	return installDependencies("libsystemd-journal-dev:mips", "mips")
}

func installLinuxMips64le() error {
	return installDependencies("libsystemd-journal-dev:mips64el", "mips64el")
}

func installLinuxMipsle() error {
	return installDependencies("libsystemd-journal-dev:mipsel", "mipsel")
}

func installLinuxPpc64le() error {
	return installDependencies("libsystemd-journal-dev:ppc64el", "ppc64el")
}

func installLinuxs390x() error {
	return installDependencies("libsystemd-journal-dev:s390x", "s390x")
}

func installDependencies(pkg, arch string) error {
	if arch != "" {
		err := sh.Run("dpkg", "--add-architecture", arch)
		if err != nil {
			return errors.Wrap(err, "error while adding architecture")
		}
	}

	if err := sh.Run("apt-get", "update"); err != nil {
		return err
	}

	return sh.Run("apt-get", "install", "-y", "--no-install-recommends", pkg)
}

// BuildGoDaemon builds the go-daemon binary (use crossBuildGoDaemon).
func BuildGoDaemon() error {
	return mage.BuildGoDaemon()
}

// CrossBuild cross-builds the beat for all target platforms.
func CrossBuild() error {
	return mage.CrossBuild(mage.ImageSelector(selectImage))
}

func selectImage(platform string) (string, error) {
	tagSuffix := "main"

	switch {
	case strings.HasPrefix(platform, "darwin"):
		tagSuffix = "darwin"
	case strings.HasPrefix(platform, "linux/arm"):
		tagSuffix = "arm"
	case strings.HasPrefix(platform, "linux/mips"):
		tagSuffix = "mips"
	case strings.HasPrefix(platform, "linux/ppc"):
		tagSuffix = "ppc"
	case platform == "linux/s390x":
		tagSuffix = "s390x"
	case strings.HasPrefix(platform, "linux"):
		tagSuffix = "main-debian8"
	}

	goVersion, err := mage.GoVersion()
	if err != nil {
		return "", err
	}

	return mage.BeatsCrossBuildImage + ":" + goVersion + "-" + tagSuffix, nil
}

// CrossBuildXPack cross-builds the beat with XPack for all target platforms.
func CrossBuildXPack() error {
	return mage.CrossBuildXPack(mage.ImageSelector(selectImage))
}

// CrossBuildGoDaemon cross-builds the go-daemon binary using Docker.
func CrossBuildGoDaemon() error {
	return mage.CrossBuildGoDaemonWithSelectableImage(selectImage)
}

// Clean cleans all generated files and build artifacts.
func Clean() error {
	return mage.Clean()
}

// Package packages the Beat for distribution.
// Use SNAPSHOT=true to build snapshots.
// Use PLATFORMS to control the target platforms.
func Package() {
	start := time.Now()
	defer func() { fmt.Println("package ran for", time.Since(start)) }()

	mage.UseElasticBeatPackaging()
	mg.Deps(Update)
	mg.Deps(CrossBuild, CrossBuildXPack, CrossBuildGoDaemon)
	mg.SerialDeps(mage.Package, TestPackages)
}

// TestPackages tests the generated packages (i.e. file modes, owners, groups).
func TestPackages() error {
	return mage.TestPackages()
}

// Update updates the generated files (aka make update).
func Update() error {
	return sh.Run("make", "update")
}

// Fields generates a fields.yml for the Beat.
func Fields() error {
	return mage.GenerateFieldsYAML()
}

// GoTestUnit executes the Go unit tests.
// Use TEST_COVERAGE=true to enable code coverage profiling.
// Use RACE_DETECTOR=true to enable the race detector.
func GoTestUnit(ctx context.Context) error {
	return mage.GoTest(ctx, mage.DefaultGoTestUnitArgs())
}

// GoTestIntegration executes the Go integration tests.
// Use TEST_COVERAGE=true to enable code coverage profiling.
// Use RACE_DETECTOR=true to enable the race detector.
func GoTestIntegration(ctx context.Context) error {
	return mage.GoTest(ctx, mage.DefaultGoTestIntegrationArgs())
}
