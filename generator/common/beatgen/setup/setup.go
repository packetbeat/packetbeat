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
	"bufio"
	"os"
	"path/filepath"

	"github.com/magefile/mage/sh"
	"github.com/pkg/errors"

	devtools "github.com/elastic/beats/v7/dev-tools/mage"
	"github.com/elastic/beats/v7/dev-tools/mage/gotool"
)

func InitModule() error {
	err := gotool.Mod.Init()
	if err != nil {
		return errors.Wrap(err, "error initializing a module for the Beat")
	}

	err = copyReplacedModules()
	if err != nil {
		return errors.Wrap(err, "error adding replaced modules to go.mod")
	}
	return gotool.Mod.Tidy()
}

func copyReplacedModules() error {
	const goModPath = "go.mod"

	beatPath, err := devtools.ElasticBeatsDir()
	if err != nil {
		return errors.Wrap(err, "error determining path to github.com/elastic/beats")
	}

	inMod, err := os.Open(filepath.Join(beatPath, goModPath))
	if err != nil {
		return errors.Wrap(err, "error opening go.mod file of github.com/elastic/beats")
	}
	defer inMod.Close()

	s := bufio.NewScanner(inMod)
	inReplaceSection := false
	replacedLines := []string{
		"",
		"replace (",
	}
	for s.Scan() {
		if err := s.Err(); err != nil {
			return errors.Wrap(err, "error reading go.mod file")
		}

		line := s.Text()
		if inReplaceSection {
			replacedLines = append(replacedLines, line)

			if line == ")" {
				break
			}
			continue
		}

		if line == "replace (" {
			inReplaceSection = true
		}
	}

	outMod, err := os.OpenFile(goModPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0640)
	if err != nil {
		return errors.Wrap(err, "error opening the go.mod file of the generated Beat")
	}
	defer outMod.Close()

	w := bufio.NewWriter(outMod)
	for _, rep := range replacedLines {
		_, err = w.WriteString(rep + "\n")
		if err != nil {
			return errors.Wrap(err, "error writing replace lines to go.mod file")
		}
	}

	return w.Flush()
}

// CopyVendor copies a new version of the dependencies to the vendor folder
func CopyVendor() error {
	err := gotool.Mod.Vendor()
	if err != nil {
		return errors.Wrapf(err, "error while running go mod vendor")
	}

	err = devtools.CopyFilesToVendor(
		"./vendor",
		[]devtools.CopyModule{
			devtools.CopyModule{
				Name: "github.com/elastic/beats/v7",
				FilesToCopy: []string{
					"dev-tools",
					"libbeat",
					"licenses",
					"metricbeat",
					"script",
					".go-version",
				},
			},
			devtools.CopyModule{
				Name: "github.com/tsg/go-daemon",
				FilesToCopy: []string{
					"src",
				},
			},
		},
	)
	if err != nil {
		return errors.Wrapf(err, "error while copying required files to vendor")
	}

	return nil
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
	err := sh.Run("make", "update")
	if err != nil {
		return err
	}
	return gotool.Mod.Tidy()
	return gotool.Mod.Download(func() gotool.ArgOpt {
		a := &gotool.Args{}
		return a.Add("all")
	})
}
