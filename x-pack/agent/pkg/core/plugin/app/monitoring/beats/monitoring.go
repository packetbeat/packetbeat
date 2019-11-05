// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package beats

import (
	"fmt"
	"path/filepath"
)

const (
	// args: pipeline name, application name
	logFileFormat = "/var/log/%s/%s"
	// args: install path, pipeline name, application name
	logFileFormatWin = "%s\\logs\\%s\\%s"

	// args: pipeline name, application name
	mbEndpointFileFormat = "unix:///var/run/fleet/%s/%s/%s.sock"
	// args: pipeline name, application name
	mbEndpointFileFormatWin = `npipe:///%s-%s`
)

func getMonitoringEndpoint(program, operatingSystem, pipelineID string) string {
	format := mbEndpointFileFormat
	if operatingSystem == "windows" {
		format = mbEndpointFileFormatWin
	}

	return fmt.Sprintf(format, pipelineID, program, program)
}

func getLoggingFile(program, operatingSystem, installPath, pipelineID string) string {
	if operatingSystem == "windows" {
		return fmt.Sprintf(logFileFormatWin, installPath, pipelineID, program)
	}

	return fmt.Sprintf(logFileFormat, pipelineID, program)
}

func getLoggingFileDirectory(installPath, operatingSystem, pipelineID string) string {
	return filepath.Base(getLoggingFile("program", operatingSystem, installPath, pipelineID))
}
