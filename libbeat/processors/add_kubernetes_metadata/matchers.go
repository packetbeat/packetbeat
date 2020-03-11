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

package add_kubernetes_metadata

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strconv"
	"sync"
	"time"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/common/fmtstr"
	"github.com/elastic/beats/v7/libbeat/logp"
	"github.com/elastic/beats/v7/libbeat/outputs/codec"
	"github.com/elastic/beats/v7/libbeat/outputs/codec/format"
)

const (
	FieldMatcherName       = "fields"
	FieldFormatMatcherName = "field_format"
	PidMatcherName         = "pid"
)

// Matcher takes a new event and returns the index
type Matcher interface {
	// MetadataIndex returns the index string to use in annotation lookups for the given
	// event. A previous indexer should have generated that index for this to work
	// This function can return "" if the event doesn't match
	MetadataIndex(event common.MapStr) string
}

type Matchers struct {
	sync.RWMutex
	matchers []Matcher
}

type MatcherConstructor func(config common.Config) (Matcher, error)

func NewMatchers(configs PluginConfig) *Matchers {
	matchers := []Matcher{}
	for _, pluginConfigs := range configs {
		for name, pluginConfig := range pluginConfigs {
			matchFunc := Indexing.GetMatcher(name)
			if matchFunc == nil {
				logp.Warn("Unable to find matcher plugin %s", name)
				continue
			}

			matcher, err := matchFunc(pluginConfig)
			if err != nil {
				logp.Warn("Unable to initialize matcher plugin %s due to error %v", name, err)
				continue
			}

			matchers = append(matchers, matcher)

		}
	}
	return &Matchers{
		matchers: matchers,
	}
}

// MetadataIndex returns the index string for the first matcher from the Registry returning one
func (m *Matchers) MetadataIndex(event common.MapStr) string {
	m.RLock()
	defer m.RUnlock()
	for _, matcher := range m.matchers {
		index := matcher.MetadataIndex(event)
		if index != "" {
			return index
		}
	}

	// No index returned
	return ""
}

func (m *Matchers) Empty() bool {
	m.RLock()
	defer m.RUnlock()
	if len(m.matchers) == 0 {
		return true
	}

	return false
}

type FieldMatcher struct {
	MatchFields []string
}

func NewFieldMatcher(cfg common.Config) (Matcher, error) {
	config := struct {
		LookupFields []string `config:"lookup_fields"`
	}{}

	err := cfg.Unpack(&config)
	if err != nil {
		return nil, fmt.Errorf("fail to unpack the `lookup_fields` configuration: %s", err)
	}

	if len(config.LookupFields) == 0 {
		return nil, fmt.Errorf("lookup_fields can not be empty")
	}

	return &FieldMatcher{MatchFields: config.LookupFields}, nil
}

func (f *FieldMatcher) MetadataIndex(event common.MapStr) string {
	for _, field := range f.MatchFields {
		keyIface, err := event.GetValue(field)
		if err == nil {
			key, ok := keyIface.(string)
			if ok {
				return key
			}
		}
	}

	return ""
}

type FieldFormatMatcher struct {
	Codec codec.Codec
}

func NewFieldFormatMatcher(cfg common.Config) (Matcher, error) {
	config := struct {
		Format string `config:"format"`
	}{}

	err := cfg.Unpack(&config)
	if err != nil {
		return nil, fmt.Errorf("fail to unpack the `format` configuration of `field_format` matcher: %s", err)
	}

	if config.Format == "" {
		return nil, fmt.Errorf("`format` of `field_format` matcher can't be empty")
	}

	return &FieldFormatMatcher{
		Codec: format.New(fmtstr.MustCompileEvent(config.Format)),
	}, nil

}

func (f *FieldFormatMatcher) MetadataIndex(event common.MapStr) string {
	bytes, err := f.Codec.Encode("", &beat.Event{
		Fields: event,
	})

	if err != nil {
		logp.Debug("kubernetes", "Unable to apply field format pattern on event")
	}

	if len(bytes) == 0 {
		return ""
	}

	return string(bytes)
}

// PidMatcher implements matcher that matches container id based on pid from cgroup file
type PidMatcher struct {
	matchRegex  string
	hostPath    string
	pidField    string
	ppidField   string
	pidCidCache *common.Cache
}

// NewPidMatcher initializes and returns a PidMatcher
func NewPidMatcher(cfg common.Config) (Matcher, error) {
	config := struct {
		RegexField string `config:"matcher_regex"`
		HostPath   string `config:"host_path"`
		PidField   string `config:"pid_field"`
		PpidField  string `config:"ppid_field"`
	}{}

	err := cfg.Unpack(&config)
	if err != nil {
		return nil, fmt.Errorf("fail to unpack the `matcher_regex` configuration: %s", err)
	}

	if len(config.RegexField) == 0 {
		// no matcher regex from conf, use default one
		config.RegexField = "^\\d+:\\w+:\\/.+\\/.+\\/.+\\/([0-9a-f]{64})"
	}

	if len(config.HostPath) == 0 {
		// no host path from conf, use default one
		config.HostPath = "/"
	}

	if len(config.PidField) == 0 {
		// no pid field from conf, use default one
		config.PidField = "process.pid"
	}

	if len(config.PpidField) == 0 {
		// no ppid field from conf, use default one
		config.PpidField = "process.ppid"
	}

	cache := common.NewCache(time.Minute, 1000)
	cache.StartJanitor(time.Minute)

	return &PidMatcher{
		matchRegex:  config.RegexField,
		hostPath:    config.HostPath,
		pidField:    config.PidField,
		ppidField:   config.PpidField,
		pidCidCache: cache}, nil
}

// MetadataIndex returns index for matching metadata indexed based on container id
func (p *PidMatcher) MetadataIndex(event common.MapStr) string {
	// As we are post processing here, actual process may be already exited.
	// So trying to find the container id by the parent process id first, hoping parent process lives a bit longer
	valPpid, errPpid := event.GetValue(p.ppidField)
	if errPpid == nil {
		ppid, ok := valPpid.(int)
		if ok {
			cid, err := p.getContainerIDFromCgroup(ppid)
			if err == nil {
				return cid
			}
			logp.Debug("kubernetes", "Unable to extract container id for ppid: %d", ppid)
		}
	}

	// No success, try with process id too, just in case
	valPid, errPid := event.GetValue(p.pidField)
	if errPid == nil {
		pid, ok := valPid.(int)
		if ok {
			cid, err := p.getContainerIDFromCgroup(pid)
			if err == nil {
				return cid
			}
			logp.Debug("kubernetes", "Unable to extract container id for pid: %d", pid)
		}
	}
	return ""
}

// For easy stubbing file read in unit tests
var readCgroupFile = func(hostPath string, pid int) ([]byte, error) {
	return ioutil.ReadFile(filepath.Join(hostPath, "proc", strconv.Itoa(pid), "cgroup"))
}

func (p *PidMatcher) getContainerIDFromCgroup(pid int) (string, error) {
	// check from from cache
	cid := p.pidCidCache.Get(pid)
	if cid != nil {
		return cid.(string), nil
	}

	// not found in cache, try to read form cgroup file
	data, err := readCgroupFile(p.hostPath, pid)
	if err != nil {
		// no cgroup file with given pid
		return "", err
	}

	// parse uuid from file contents
	re := regexp.MustCompile(p.matchRegex)
	rs := re.FindStringSubmatch(string(data))
	if rs != nil {
		// add pid and cid to cache
		cid := rs[1]
		p.pidCidCache.Put(pid, cid)
		return cid, nil
	}

	// no regex match, probably not a container process
	return "", nil
}
