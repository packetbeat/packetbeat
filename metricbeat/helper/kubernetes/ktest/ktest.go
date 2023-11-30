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

package ktest

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	p "github.com/elastic/beats/v7/metricbeat/helper/prometheus"
	"github.com/elastic/beats/v7/metricbeat/helper/prometheus/ptest"
)

func getFiles(folder string) ([]string, error) {
	var files []string //nolint:prealloc
	entries, err := os.ReadDir(folder)
	if err != nil {
		return nil, err
	}
	for _, e := range entries {
		files = append(files, filepath.Join(folder, e.Name()))
	}
	return files, nil
}

// GetTestCases Build test cases based on the files from folder, and the expected files in the expectedFolder
func GetTestCases(folder string, expectedFolder string) (ptest.TestCases, error) {
	var files []string
	var cases ptest.TestCases

	files, err := getFiles(folder)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		cases = append(cases,
			struct {
				MetricsFile  string
				ExpectedFile string
			}{
				// the metrics file is inside the parent directory, while the expected file is in the current directory
				// Example:
				//	Metricsfile: ../_meta/test/ksm.v2.4.2
				//	ExpectedFile: ./_meta/test/ksm.v2.4.2.expected
				MetricsFile:  file,
				ExpectedFile: filepath.Join(expectedFolder, filepath.Base(file)+".expected"),
			},
		)
	}
	return cases, nil
}

// TestMetricsFamilyFromFiles
// This function reads the metric files and checks if the resource fetched metrics exist in it.
// It only checks the family metric, because if the metric doesn't have any data, we don't have a way
// to know the labels from the file.
// The test fails if the metric does not exist in any of the files.
// A warning is printed if the metric is not present in all of them.
// Nothing happens, otherwise.
func TestMetricsFamilyFromFiles(t *testing.T, files []string, mapping *p.MetricsMapping) {
	metricsFiles := map[string][]string{}
	for i := 0; i < len(files); i++ {
		content, err := ioutil.ReadFile(files[i])
		if err != nil {
			t.Fatalf("Unknown file %s.", files[i])
		}
		text := string(content)
		for metric := range mapping.Metrics {
			// A space is needed to check if the metric exists, since there are metrics that can follow this logic:
			// some_metric and some_metric_total
			if !strings.Contains(text, "# TYPE "+metric+" ") {
				metricsFiles[metric] = append(metricsFiles[metric], files[i])
			}
		}
	}
	for metric, filesList := range metricsFiles {
		if len(filesList) != len(files) {
			t.Logf("Warning: metric %s is not present in all files.", metric)
		} else {
			t.Errorf("Unknown metric: %s", metric)
		}
	}

}

// TestMetricsFamilyFromFolder is the same as TestMetricsFamilyFromFiles, but for folder
func TestMetricsFamilyFromFolder(t *testing.T, folder string, mapping *p.MetricsMapping) {
	files, err := getFiles(folder)
	if err != nil {
		t.Fatalf(err.Error())
	}

	metricsFiles := map[string][]string{}
	for i := 0; i < len(files); i++ {
		content, err := ioutil.ReadFile(files[i])
		if err != nil {
			t.Fatalf("Unknown file %s.", files[i])
		}
		text := string(content)
		for metric := range mapping.Metrics {
			// A space is needed to check if the metric exists, since there are metrics that can follow this logic:
			// some_metric and some_metric_total
			if !strings.Contains(text, "# TYPE "+metric+" ") {
				metricsFiles[metric] = append(metricsFiles[metric], files[i])
			}
		}
	}
	for metric, filesList := range metricsFiles {
		if len(filesList) != len(files) {
			t.Logf("Warning: metric %s is not present in all files.", metric)
		} else {
			t.Errorf("Unknown metric: %s", metric)
		}
	}

}
