// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package paths

import (
	"os"
	"path/filepath"
)

var (
	homePath string
	dataPath string
)

func init() {
	exePath := retrieveExecutablePath()
	homePath = exePath
	dataPath = filepath.Join(exePath, "data")
}

// Home returns a directory where binary lives
// Executable is not supported on nacl.
func Home() string {
	return homePath
}

// Data returns a home directory of current user
func Data() string {
	return dataPath
}

func retrieveExecutablePath() string {
	execPath, err := os.Executable()
	if err != nil {
		panic(err)
	}

	return filepath.Dir(execPath)
}
