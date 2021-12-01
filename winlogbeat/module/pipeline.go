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

package module

import (
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/joeshaw/multierror"
	"gopkg.in/yaml.v2"

	"github.com/elastic/beats/v7/filebeat/fileset"
	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/esleg/eslegclient"
	"github.com/elastic/beats/v7/libbeat/logp"
)

var PipelinesFS *embed.FS
var errNoFS = errors.New("no ebedded file system")

const logName = "pipeline"

type pipeline struct {
	id       string
	contents map[string]interface{}
}

// Upload reads all pipelines stored in winlogbeat executable adapts
// pipeline for a given ES version, converts to JSON if necessary and
// creates or updates ingest pipeline in ES
func UploadPipelines(info beat.Info, esClient *eslegclient.Connection, overwritePipelines bool) error {
	log := logp.NewLogger(logName)
	pipelines, err := readAll(info)
	if err != nil {
		return err
	}
	return load(esClient, pipelines, overwritePipelines, log)
}

// Export reads all pipelines stored in winlogbeat executable adapts
// pipelines for a given ES version and writes the converted pipelines
// to the given directory in JSON format
func ExportPipelines(info beat.Info, version common.Version, directory string) error {
	if !info.ElasticLicensed {
		return nil
	}
	log := logp.NewLogger(logName)
	pipelines, err := readAll(info)
	if err != nil {
		return err
	}

	for _, pipeline := range pipelines {
		if err := fileset.AdaptPipelineForCompatibility(version, pipeline.id, pipeline.contents, log); err != nil {
			return fmt.Errorf("failed to adapt pipeline with backwards compatibility changes for version %s: %w", version.String(), err)
		}
		f, err := os.Create(filepath.Join(directory, pipeline.id+".json"))
		if err != nil {
			return fmt.Errorf("unable to create file to export pipeline to: %+v", err)
		}
		enc := json.NewEncoder(f)
		enc.SetEscapeHTML(false)
		err = enc.Encode(pipeline.contents)
		if err != nil {
			return fmt.Errorf("unable to JSON encode pipeline %s: %+v", f.Name(), err)
		}
	}
	return nil
}

// readAll reads pipelines from the the embedded filesystem and
// returns a slice of pipelines suitable for sending to Elasticsearch
// with load.
func readAll(info beat.Info) (pipelines []pipeline, err error) {
	p, err := readDir(".", info)
	if err == errNoFS {
		return nil, nil
	}
	return p, err
}

func readDir(dir string, info beat.Info) (pipelines []pipeline, err error) {
	if PipelinesFS == nil {
		return nil, errNoFS
	}
	dirEntries, err := PipelinesFS.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, de := range dirEntries {
		if de.IsDir() {
			subPipelines, err := readDir(path.Join(dir, de.Name()), info)
			if err != nil {
				return nil, err
			}
			pipelines = append(pipelines, subPipelines...)
			continue
		}
		p, err := readFile(path.Join(dir, de.Name()), info)
		if err == errNoFS {
			continue
		}
		if err != nil {
			return nil, err
		}
		pipelines = append(pipelines, p)
	}
	return pipelines, nil
}

func readFile(filename string, info beat.Info) (p pipeline, err error) {
	if PipelinesFS == nil {
		return pipeline{}, errNoFS
	}
	contents, err := PipelinesFS.ReadFile(filename)
	if err != nil {
		return pipeline{}, err
	}
	updatedContent, err := applyTemplates(info.IndexPrefix, info.Version, filename, contents)
	if err != nil {
		return pipeline{}, err
	}
	p = pipeline{
		id:       fileset.FormatPipelineID(info.IndexPrefix, "default", "default", filename, info.Version),
		contents: updatedContent,
	}
	return p, nil
}

// Load uses esClient to load pipelines to Elasticsearch
// cluster.  Will only overwrite existing pipelines if
// overwritePipelines is true.
func load(esClient *eslegclient.Connection, pipelines []pipeline, overwritePipelines bool, log *logp.Logger) (err error) {
	var pipelineIDsLoaded []string

	for _, pipeline := range pipelines {
		err = fileset.LoadPipeline(esClient, pipeline.id, pipeline.contents, overwritePipelines, log)
		if err != nil {
			err = fmt.Errorf("error loading pipeline %s: %v", pipeline.id, err)
			break
		}
		pipelineIDsLoaded = append(pipelineIDsLoaded, pipeline.id)
	}

	if err != nil {
		errs := multierror.Errors{err}
		for _, pipelineID := range pipelineIDsLoaded {
			err = fileset.DeletePipeline(esClient, pipelineID)
			if err != nil {
				errs = append(errs, err)
			}
		}
		return errs.Err()
	}
	return nil
}

func applyTemplates(prefix string, version string, filename string, original []byte) (converted map[string]interface{}, err error) {
	vars := map[string]interface{}{
		"builtin": map[string]interface{}{
			"prefix":      prefix,
			"module":      "default",
			"fileset":     "default",
			"beatVersion": version,
		},
	}

	encodedString, err := fileset.ApplyTemplate(vars, string(original), true)

	var content map[string]interface{}
	switch extension := strings.ToLower(filepath.Ext(filename)); extension {
	case ".json":
		if err = json.Unmarshal([]byte(encodedString), &content); err != nil {
			return nil, fmt.Errorf("Error JSON decoding the pipeline file: %s: %v", filename, err)
		}
	case ".yaml", ".yml":
		if err = yaml.Unmarshal([]byte(encodedString), &content); err != nil {
			return nil, fmt.Errorf("Error YAML decoding the pipeline file: %s: %v", filename, err)
		}
		newContent, err := fileset.FixYAMLMaps(content)
		if err != nil {
			return nil, fmt.Errorf("Failed to sanitize the YAML pipeline file: %s: %v", filename, err)
		}
		content = newContent.(map[string]interface{})
	default:
		return nil, fmt.Errorf("Unsupported extension '%s' for pipeline file: %s", extension, filename)
	}
	return content, nil
}
