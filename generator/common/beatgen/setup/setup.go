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

package setup

import (
	"fmt"
	"os"
	"path/filepath"

	devtools "github.com/elastic/beats/dev-tools/mage"
	"github.com/magefile/mage/sh"
	"github.com/pkg/errors"
)

// RunSetup runs any remaining setup commands after the vendor directory has been setup
func RunSetup() error {
	vendorPath := "./vendor/github.com/"

	//Copy mage stuff
	err := os.MkdirAll(filepath.Join(vendorPath, "magefile"), 0755)
	if err != nil {
		return errors.Wrap(err, "error making mage dir")
	}

	err = sh.Run("cp", "-R", filepath.Join(vendorPath, "elastic/beats/vendor/github.com/magefile/mage"), filepath.Join(vendorPath, "magefile"))
	if err != nil {
		return errors.Wrap(err, "error copying vendored magefile")
	}

	//Copy the pkg helper
	err = sh.Run("cp", "-R", filepath.Join(vendorPath, "elastic/beats/vendor/github.com/pkg"), vendorPath)
	if err != nil {
		return errors.Wrap(err, "error copying pkg to vendor")
	}

	pwd, err := os.Getwd()
	if err != nil {
		return errors.Wrap(err, "error gettting current directory")
	}

	return sh.Run("ln", "-sf", filepath.Join(pwd, vendorPath, "elastic/beats/metricbeat/scripts/generate_imports_helper.py"), filepath.Join(pwd, vendorPath, "elastic/beats/script/generate_imports_helper.py"))

}

// CopyVendor copies a new version of beats into the vendor directory of PWD
// By default this uses git archive, meaning any uncommitted changes will not be copied.
// Set the NEWBEAT_DEV env variable to use a slow `cp` copy that will catch uncommited changes
func CopyVendor() error {
	vendorPath := "./vendor/github.com/elastic/"
	beatPath, err := devtools.ElasticBeatsDir()
	if err != nil {
		return errors.Wrap(err, "Could not find ElasticBeatsDir")
	}
	err = os.MkdirAll(vendorPath, 0755)
	if err != nil {
		return errors.Wrap(err, "error creating vendor dir")
	}
	if _, isDev := os.LookupEnv(CfgPrefix + "_DEV"); isDev {
		//Dev mode. Use CP.
		fmt.Printf("CopyVendor running in dev mode, elastic/beats will be copied into the vendor directory with cp\n")
		vendorPath = filepath.Join(vendorPath, "beats")

		err = sh.Run("cp", "-R", beatPath, vendorPath)
		if err != nil {
			return errors.Wrap(err, "error copying vendor dir")
		}
		err = sh.Rm(filepath.Join(vendorPath, ".git"))
		if err != nil {
			return errors.Wrap(err, "error removing vendor git directory")
		}
		err = sh.Rm(filepath.Join(vendorPath, "x-pack"))
		if err != nil {
			return errors.Wrap(err, "error removing x-pack directory")
		}
	} else {
		//not dev mode. Use git archive
		vendorPath = filepath.Join(vendorPath, "beats")
		err = os.MkdirAll(vendorPath, 0755)
		if err != nil {
			return errors.Wrap(err, "error creating vendor dir")
		}
		err = sh.Run("sh",
			"-c",
			"git archive --remote "+beatPath+" HEAD |  tar -x --exclude=x-pack -C "+vendorPath)
		if err != nil {
			return errors.Wrap(err, "error running git archive")
		}
	}

	return nil

}

// getRealBeatsRepo gets a non
func getRealBeatsRepo() {

}

// GitInit initializes a new git repo in the current directory
func GitInit() error {
	return sh.Run("git", "init")
}

// GitAdd adds the current working directory to git and does an initial commit
func GitAdd() error {
	err := sh.Run("git", "add", "-A")
	if err != nil {
		return errors.Wrap(err, "error running git add")
	}
	return sh.Run("git", "commit", "-q", "-m", "Initial commit, Add generated files")
}

// Update updates the generated files (aka make update).
func Update() error {
	return sh.Run("make", "update")
}
