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

package metrics

import (
	"bufio"
	"strings"

	"github.com/joeshaw/multierror"
	"github.com/pkg/errors"

	"github.com/elastic/beats/v7/libbeat/common"
)

// fillTicks is the FreeBSD implementation of FillTicks
func (self CPU) fillTicks(event *common.MapStr) {
	event.Put("user.ticks", self.user)
	event.Put("system.ticks", self.sys)
	event.Put("idle.ticks", self.idle)
	event.Put("nice.ticks", self.nice)
}

func fillCPUMetrics(event *common.MapStr, current, prev CPU, numCPU int, timeDelta uint64, pathPostfix string) {
	idleTime := cpuMetricTimeDelta(prev.idle, current.idle, timeDelta, numCPU)
	totalPct := common.Round(float64(numCPU)-idleTime, common.DefaultDecimalPlacesCount)

	event.Put("total"+pathPostfix, totalPct)
	event.Put("user"+pathPostfix, cpuMetricTimeDelta(prev.user, current.user, timeDelta, numCPU))
	event.Put("system"+pathPostfix, cpuMetricTimeDelta(prev.sys, current.sys, timeDelta, numCPU))
	event.Put("idle"+pathPostfix, cpuMetricTimeDelta(prev.idle, current.idle, timeDelta, numCPU))
	event.Put("nice"+pathPostfix, cpuMetricTimeDelta(prev.nice, current.nice, timeDelta, numCPU))
}

func scanStatFile(scanner *bufio.Scanner) (CPUMetrics, error) {
	cpuData, err := statScanner(scanner, parseCPULine)
	if err != nil {
		return CPUMetrics{}, errors.Wrap(err, "error scanning stat file")
	}
	return cpuData, nil
}

func parseCPULine(line string) (CPU, error) {
	cpuData := CPU{}
	fields := strings.Fields(line)
	var errs multierror.Errors
	var err error

	cpuData.user, err = touint(fields[1])
	if err != nil {
		errs = append(errs, err)
	}
	cpuData.nice, err = touint(fields[2])
	if err != nil {
		errs = append(errs, err)
	}
	cpuData.sys, err = touint(fields[3])
	if err != nil {
		errs = append(errs, err)
	}
	cpuData.idle, err = touint(fields[4])
	if err != nil {
		errs = append(errs, err)
	}

	return cpuData, errs.Err()
}
