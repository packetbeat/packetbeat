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

package fingerprint

import (
	"fmt"
	"hash"
	"sort"
	"time"

	"github.com/pkg/errors"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/processors"
	jsprocessor "github.com/elastic/beats/libbeat/processors/script/javascript/module/processor"
)

func init() {
	processors.RegisterPlugin("fingerprint", New)
	jsprocessor.RegisterPlugin("Fingerprint", New)
}

const processorName = "fingerprint"

type fingerprint struct {
	config Config
	hash   hash.Hash
}

// New constructs a new fingerprint processor.
func New(cfg *common.Config) (processors.Processor, error) {
	config := defaultConfig()
	if err := cfg.Unpack(&config); err != nil {
		return nil, errors.Wrapf(err, "failed to unpack %v processor configuration", processorName)
	}

	sort.Strings(config.Fields)

	p := &fingerprint{
		config: config,
		hash:   config.Method(),
	}

	return p, nil
}

// Run enriches the given event with fingerprint information
func (p *fingerprint) Run(event *beat.Event) (*beat.Event, error) {
	source, err := makeSource(p.config.Fields, event.Fields)
	if err != nil {
		return nil, makeComputeFingerprintError(err)
	}

	hashFn := p.hash
	hashFn.Reset()

	hashFn.Write([]byte(source))
	hash := hashFn.Sum(nil)

	encodedHash := p.config.Encoding(hash)

	if _, err = event.PutValue(p.config.TargetField, encodedHash); err != nil {
		return nil, makeComputeFingerprintError(err)
	}

	return event, nil
}

func (p *fingerprint) String() string {
	return fmt.Sprintf("%v=[method=[%v]]", processorName, p.config.Method)
}

func makeSource(sourceFields []string, eventFields common.MapStr) (string, error) {
	var str string
	for _, k := range sourceFields {
		v, err := eventFields.GetValue(k)
		if err == common.ErrKeyNotFound {
			return "", errors.Wrapf(err, "failed to find field [%v] in event", k)
		}
		if err != nil {
			return "", errors.Wrapf(err, "failed when finding field [%v] in event", k)
		}

		i := v
		switch vv := v.(type) {
		case map[string]interface{}, []interface{}:
			return "", errors.Errorf("cannot compute fingerprint using non-scalar field [%v]", k)
		case time.Time:
			// Ensure we consistently hash times in UTC.
			i = vv.UTC()
		}

		str += fmt.Sprintf("|%v|%v", k, i)
	}
	str += "|"
	return str, nil
}

func makeComputeFingerprintError(err error) error {
	return errors.Wrap(err, "failed to compute fingerprint")
}
