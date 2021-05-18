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

// +build openbsd

package metrics

/*
#include <sys/param.h>
#include <sys/types.h>
#include <sys/sysctl.h>
#include <sys/mount.h>
#include <sys/sched.h>
#include <sys/swap.h>
#include <stdlib.h>
#include <unistd.h>
*/
import "C"

import (
	"syscall"
	"unsafe"

	"github.com/elastic/beats/v7/libbeat/common"
)

// fillTicks is the OpenBSD implementation of FillTicks
func (self CPU) fillTicks(event *common.MapStr) {
	event.Put("user.ticks", self.user)
	event.Put("system.ticks", self.sys)
	event.Put("idle.ticks", self.idle)
	event.Put("nice.ticks", self.nice)
	event.Put("irq.ticks", self.irq)
}

func fillCPUMetrics(event *common.MapStr, current, prev CPU, numCPU int, timeDelta uint64, pathPostfix string) {
	idleTime := cpuMetricTimeDelta(prev.idle, current.idle, timeDelta, numCPU)
	totalPct := common.Round(float64(numCPU)-idleTime, common.DefaultDecimalPlacesCount)

	event.Put("total"+pathPostfix, totalPct)
	event.Put("user"+pathPostfix, cpuMetricTimeDelta(prev.user, current.user, timeDelta, numCPU))
	event.Put("system"+pathPostfix, cpuMetricTimeDelta(prev.sys, current.sys, timeDelta, numCPU))
	event.Put("idle"+pathPostfix, cpuMetricTimeDelta(prev.idle, current.idle, timeDelta, numCPU))
	event.Put("nice"+pathPostfix, cpuMetricTimeDelta(prev.nice, current.nice, timeDelta, numCPU))
	event.Put("irq"+pathPostfix, cpuMetricTimeDelta(prev.irq, current.irq, timeDelta, numCPU))
}

func Get(_ string) (CPUMetrics, error) {

	// see man 2 sysctl
	loadGlobal := [C.CPUSTATES]C.long{
		C.CP_USER,
		C.CP_NICE,
		C.CP_SYS,
		C.CP_INTR,
		C.CP_IDLE,
	}

	// first, fetch global CPU data
	err := sysctlGetCPUTimes(0, 0, &loadGlobal)
	if err != nil {
		return CPUMetrics{}, err
	}
	self := CPU{}

	self.user = uint64(loadGlobal[0])
	self.nice = uint64(loadGlobal[1])
	self.sys = uint64(loadGlobal[2])
	self.irq = uint64(loadGlobal[3])
	self.idle = uint64(loadGlobal[4])
	// Get count of available CPUs
	ncpuMIB := [2]int32{C.CTL_HW, C.HW_NCPU}
	callSize := uintptr(0)
	var ncpu int
	// The first call nulls out the retun pointer, so we instead fetch the expected size of expected return value.
	_, _, errno := syscall.Syscall6(syscall.SYS___SYSCTL, uintptr(unsafe.Pointer(&ncpuMIB[0])), 2, 0, uintptr(unsafe.Pointer(&callSize)), 0, 0)

	if errno != 0 || callSize == 0 {
		return CPUMetrics{}, errno
	}

	// Populate the cpu count
	_, _, errno = syscall.Syscall6(syscall.SYS___SYSCTL, uintptr(unsafe.Pointer(&ncpuMIB[0])), 2, uintptr(unsafe.Pointer(&ncpu)), uintptr(unsafe.Pointer(&callSize)), 0, 0)

	if errno != 0 || callSize == 0 {
		return CPUMetrics{}, errno
	}

	loadPerCPU := [C.CPUSTATES]C.long{
		C.CP_USER,
		C.CP_NICE,
		C.CP_SYS,
		C.CP_INTR,
		C.CP_IDLE,
	}

	perCPU := make([]CPU, ncpu)

	// iterate over metrics for each CPU
	for i := 0; i < ncpu; i++ {
		sysctlGetCPUTimes(ncpu, i, &loadPerCPU)
		perCPU[i].user = uint64(loadPerCPU[0])
		perCPU[i].nice = uint64(loadPerCPU[1])
		perCPU[i].sys = uint64(loadPerCPU[2])
		perCPU[i].irq = uint64(loadPerCPU[3])
		perCPU[i].idle = uint64(loadPerCPU[4])
	}

	metrics := CPUMetrics{totals: self, list: perCPU}

	return metrics, nil
}

// sysctlGetCPUTimes runs the CTL_KERN::KERN_CPTIME sysctl command against the host.
func sysctlGetCPUTimes(ncpu int, curcpu int, load *[C.CPUSTATES]C.long) error {
	var mib []int32

	// Use the correct mib based on the number of CPUs and fill out the
	// current CPU number in case of SMP. (0 indexed cf. self.List)
	if ncpu == 0 {
		mib = []int32{C.CTL_KERN, C.KERN_CPTIME}
	} else {
		mib = []int32{C.CTL_KERN, C.KERN_CPTIME2, int32(curcpu)}
	}

	len := len(mib)

	n := uintptr(0)
	// First we determine how much memory we'll need to pass later on (via `n`)
	_, _, errno := syscall.Syscall6(syscall.SYS___SYSCTL, uintptr(unsafe.Pointer(&mib[0])), uintptr(len), 0, uintptr(unsafe.Pointer(&n)), 0, 0)
	if errno != 0 || n == 0 {
		return errno
	}

	_, _, errno = syscall.Syscall6(syscall.SYS___SYSCTL, uintptr(unsafe.Pointer(&mib[0])), uintptr(len), uintptr(unsafe.Pointer(load)), uintptr(unsafe.Pointer(&n)), 0, 0)
	if errno != 0 || n == 0 {
		return errno
	}

	return nil
}
