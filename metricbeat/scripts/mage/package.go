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

package mage

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/pkg/errors"

	"github.com/elastic/beats/dev-tools/mage"
)

const (
	dirModulesDGenerated = "build/package/modules.d"
)

// CustomizePackaging modifies the package specs to add the modules.d directory.
// And for Windows it comments out the system/load metricset because it's
// not supported. You must declare a dependency on either
// PrepareModulePackagingOSS or PrepareModulePackagingXPack.
func CustomizePackaging() {
	var (
		modulesDTarget = "modules.d"
		modulesD       = mage.PackageFile{
			Mode:    0644,
			Source:  dirModulesDGenerated,
			Config:  true,
			Modules: true,
		}
		windowsModulesDir = mage.PackageFile{
			Mode:    0644,
			Source:  "{{.PackageDir}}/modules.d",
			Config:  true,
			Modules: true,
			Dep: func(spec mage.PackageSpec) error {
				if err := mage.Copy("modules.d", spec.MustExpand("{{.PackageDir}}/modules.d")); err != nil {
					return errors.Wrap(err, "failed to copy modules.d dir")
				}

				return mage.FindReplace(
					spec.MustExpand("{{.PackageDir}}/modules.d/system.yml"),
					regexp.MustCompile(`- load`), `#- load`)
			},
		}
		windowsReferenceConfig = mage.PackageFile{
			Mode:   0644,
			Source: "{{.PackageDir}}/metricbeat.reference.yml",
			Dep: func(spec mage.PackageSpec) error {
				err := mage.Copy("metricbeat.reference.yml",
					spec.MustExpand("{{.PackageDir}}/metricbeat.reference.yml"))
				if err != nil {
					return errors.Wrap(err, "failed to copy reference config")
				}

				return mage.FindReplace(
					spec.MustExpand("{{.PackageDir}}/metricbeat.reference.yml"),
					regexp.MustCompile(`- load`), `#- load`)
			},
		}
	)

	for _, args := range mage.Packages {
		switch args.OS {
		case "windows":
			args.Spec.Files[modulesDTarget] = windowsModulesDir
			args.Spec.ReplaceFile("{{.BeatName}}.reference.yml", windowsReferenceConfig)
		default:
			pkgType := args.Types[0]
			switch pkgType {
			case mage.TarGz, mage.Zip, mage.Docker:
				args.Spec.Files[modulesDTarget] = modulesD
			case mage.Deb, mage.RPM, mage.DMG:
				args.Spec.Files["/etc/{{.BeatName}}/"+modulesDTarget] = modulesD
			default:
				panic(errors.Errorf("unhandled package type: %v", pkgType))
			}
		}
	}
}

// PrepareModulePackagingOSS generates build/package/modules and
// build/package/modules.d directories for use in packaging.
func PrepareModulePackagingOSS() error {
	return prepareModulePackaging([]struct{ Src, Dst string }{
		{mage.OSSBeatDir("modules.d"), dirModulesDGenerated},
	}...)
}

// PrepareModulePackagingXPack generates build/package/modules and
// build/package/modules.d directories for use in packaging.
func PrepareModulePackagingXPack() error {
	return prepareModulePackaging([]struct{ Src, Dst string }{
		{mage.OSSBeatDir("modules.d"), dirModulesDGenerated},
		{"modules.d", dirModulesDGenerated},
	}...)
}

func prepareModulePackaging(files ...struct{ Src, Dst string }) error {
	mg.Deps(GenerateDirModulesD)

	err := mage.Clean([]string{
		dirModulesDGenerated,
	})
	if err != nil {
		return err
	}

	for _, copyAction := range files {
		err := (&mage.CopyTask{
			Source:  copyAction.Src,
			Dest:    copyAction.Dst,
			Mode:    0644,
			DirMode: 0755,
		}).Execute()
		if err != nil {
			return err
		}
	}
	return nil
}

// Generate
func GenerateDirModulesD() error {
	if err := os.RemoveAll("modules.d"); err != nil {
		return err
	}

	shortConfigs, err := filepath.Glob("module/*/_meta/config.yml")
	if err != nil {
		return err
	}

	for _, f := range shortConfigs {
		parts := strings.Split(filepath.ToSlash(f), "/")
		if len(parts) < 2 {
			continue
		}
		moduleName := parts[1]

		suffix := ".yml.disabled"
		if moduleName == "system" {
			suffix = ".yml"
		}

		cp := mage.CopyTask{
			Source: f,
			Dest:   filepath.Join("modules.d", moduleName+suffix),
			Mode:   0644,
		}
		if err = cp.Execute(); err != nil {
			return err
		}
	}
	return nil
}
