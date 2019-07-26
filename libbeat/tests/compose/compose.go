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

package compose

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

// HostInfo exposes information about started scenario
type HostInfo interface {
	Host() string
	HostForPort(int) string
}

// EnsureUp starts all the requested services (must be defined in docker-compose.yml)
// with a default timeout of 300 seconds
func EnsureUp(t *testing.T, service string) HostInfo {
	t.Helper()
	return EnsureUpWithTimeout(t, 60, service)
}

// EnsureUpWithTimeout starts all the requested services (must be defined in docker-compose.yml)
// Wait for `timeout` seconds for health
func EnsureUpWithTimeout(t *testing.T, timeout int, service string) HostInfo {
	t.Helper()
	// The NO_COMPOSE env variables makes it possible to skip the starting of the environment.
	// This is useful if the service is already running locally.
	if noCompose, err := strconv.ParseBool(os.Getenv("NO_COMPOSE")); err == nil && noCompose {
		envVar := fmt.Sprintf("%s_HOST", strings.ToUpper(service))
		host := os.Getenv(envVar)
		if host == "" {
			t.Fatalf("%s environment variable must be set as the host:port where %s is running", envVar, service)
		}
		return &staticHostInfo{host: host}
	}

	compose, err := getComposeProject(os.Getenv("DOCKER_COMPOSE_PROJECT_NAME"))
	if err != nil {
		t.Fatal(err)
	}

	// Kill no longer used containers
	err = compose.KillOld([]string{service})
	if err != nil {
		t.Fatal(err)
	}

	// Start container
	err = compose.Start(service, RecreateOnUp(false))
	if err != nil {
		t.Fatal("failed to start service", service, err)
	}

	// Wait for health
	err = compose.Wait(timeout, service)
	if err != nil {
		t.Fatal(err)
	}

	// Get host information
	host, err := compose.HostInformation(service)
	if err != nil {
		t.Fatalf("getting host for %s", service)
	}

	return host
}

type staticHostInfo struct {
	host string
}

func (i *staticHostInfo) Host() string {
	return i.host
}

func (i *staticHostInfo) HostForPort(int) string {
	return i.host
}

func findComposePath() (string, error) {
	// find docker-compose
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if path == "/" {
			break
		}

		composePath := filepath.Join(path, "docker-compose.yml")
		if _, err = os.Stat(composePath); err == nil {
			return composePath, nil
		}
		path = filepath.Dir(path)
	}

	return "", errors.New("docker-compose.yml not found")
}

func getComposeProject(name string) (*Project, error) {
	path, err := findComposePath()
	if err != nil {
		return nil, err
	}

	return NewProject(
		name,
		[]string{
			path,
		},
	)
}
