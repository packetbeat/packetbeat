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

package docker

import (
	"fmt"
	"os"
	"path"
	"strings"

	devtools "github.com/elastic/beats/v7/dev-tools/mage"
	"github.com/magefile/mage/mg"
)

type Docker mg.Namespace

// ComposeBuild builds the docker-compose containers.
func (Docker) ComposeBuild() error {
	return devtools.BuildIntegTestContainers()
}

// ComposeUp starts the docker-compose containers, waits until they are healthy, and leaves them in the background.
func (Docker) ComposeUp() error {
	return devtools.StartIntegTestContainers()
}

// ComposeDown stops the docker-compose containers started by ComposeUp.
func (Docker) ComposeDown() error {
	return devtools.StopIntegTestContainers()
}

// ComposeProject prints the docker-compose project name used when starting containers.
// Pass this to docker-compose with the -p option to interact with running containers.
func (Docker) ComposeProject() {
	fmt.Println(devtools.DockerComposeProjectName())
}

// ComposeEnvFile generates the environment variable file to pass to docker-compose with --env-file.
func (Docker) ComposeEnvFile() error {
	esBeatsDir, err := devtools.ElasticBeatsDir()
	if err != nil {
		return err
	}

	envFile := path.Join(esBeatsDir, "docker.env")
	envVars := []string{
		"# Environment variable file to pass to docker-compose with the --env-file option.",
		fmt.Sprintf("ES_BEATS=%s", esBeatsDir),
		fmt.Sprintf("STACK_ENVIRONMENT=%s", devtools.StackEnvironment),
		fmt.Sprintf("TESTING_ENVIRONMENT=%s", devtools.StackEnvironment),
	}

	err = os.WriteFile(
		envFile,
		[]byte(strings.Join(envVars, "\n")),
		0644,
	)
	if err != nil {
		return err
	}

	fmt.Printf("Created: %s\n", envFile)
	return nil
}
