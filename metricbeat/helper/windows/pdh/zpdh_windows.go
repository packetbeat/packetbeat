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

// Code generated by 'go generate'; DO NOT EDIT.

package pdh

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return nil
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modpdh = windows.NewLazySystemDLL("pdh.dll")

	procPdhOpenQueryW                = modpdh.NewProc("PdhOpenQueryW")
	procPdhAddEnglishCounterW        = modpdh.NewProc("PdhAddEnglishCounterW")
	procPdhAddCounterW               = modpdh.NewProc("PdhAddCounterW")
	procPdhRemoveCounter             = modpdh.NewProc("PdhRemoveCounter")
	procPdhCollectQueryData          = modpdh.NewProc("PdhCollectQueryData")
	procPdhGetFormattedCounterValue  = modpdh.NewProc("PdhGetFormattedCounterValue")
	procPdhCloseQuery                = modpdh.NewProc("PdhCloseQuery")
	procPdhExpandWildCardPathW       = modpdh.NewProc("PdhExpandWildCardPathW")
	procPdhExpandCounterPathW        = modpdh.NewProc("PdhExpandCounterPathW")
	procPdhGetCounterInfoW           = modpdh.NewProc("PdhGetCounterInfoW")
	procPdhEnumObjectItemsW          = modpdh.NewProc("PdhEnumObjectItemsW")
	procPdhGetFormattedCounterArrayW = modpdh.NewProc("PdhGetFormattedCounterArrayW")
)

func _PdhOpenQuery(dataSource *uint16, userData uintptr, query *PdhQueryHandle) (errcode error) {
	r0, _, _ := syscall.Syscall(procPdhOpenQueryW.Addr(), 3, uintptr(unsafe.Pointer(dataSource)), uintptr(userData), uintptr(unsafe.Pointer(query)))
	if r0 != 0 {
		errcode = syscall.Errno(r0)
	}
	return
}

func _PdhAddEnglishCounter(query PdhQueryHandle, counterPath string, userData uintptr, counter *PdhCounterHandle) (errcode error) {
	var _p0 *uint16
	_p0, errcode = syscall.UTF16PtrFromString(counterPath)
	if errcode != nil {
		return
	}
	return __PdhAddEnglishCounter(query, _p0, userData, counter)
}

func __PdhAddEnglishCounter(query PdhQueryHandle, counterPath *uint16, userData uintptr, counter *PdhCounterHandle) (errcode error) {
	r0, _, _ := syscall.Syscall6(procPdhAddEnglishCounterW.Addr(), 4, uintptr(query), uintptr(unsafe.Pointer(counterPath)), uintptr(userData), uintptr(unsafe.Pointer(counter)), 0, 0)
	if r0 != 0 {
		errcode = syscall.Errno(r0)
	}
	return
}

func _PdhAddCounter(query PdhQueryHandle, counterPath string, userData uintptr, counter *PdhCounterHandle) (errcode error) {
	var _p0 *uint16
	_p0, errcode = syscall.UTF16PtrFromString(counterPath)
	if errcode != nil {
		return
	}
	return __PdhAddCounter(query, _p0, userData, counter)
}

func __PdhAddCounter(query PdhQueryHandle, counterPath *uint16, userData uintptr, counter *PdhCounterHandle) (errcode error) {
	r0, _, _ := syscall.Syscall6(procPdhAddCounterW.Addr(), 4, uintptr(query), uintptr(unsafe.Pointer(counterPath)), uintptr(userData), uintptr(unsafe.Pointer(counter)), 0, 0)
	if r0 != 0 {
		errcode = syscall.Errno(r0)
	}
	return
}

func _PdhRemoveCounter(counter PdhCounterHandle) (errcode error) {
	r0, _, _ := syscall.Syscall(procPdhRemoveCounter.Addr(), 1, uintptr(counter), 0, 0)
	if r0 != 0 {
		errcode = syscall.Errno(r0)
	}
	return
}

func _PdhCollectQueryData(query PdhQueryHandle) (errcode error) {
	r0, _, _ := syscall.Syscall(procPdhCollectQueryData.Addr(), 1, uintptr(query), 0, 0)
	if r0 != 0 {
		errcode = syscall.Errno(r0)
	}
	return
}

func _PdhGetFormattedCounterValueDouble(counter PdhCounterHandle, format PdhCounterFormat, counterType *uint32, value *PdhCounterValueDouble) (errcode error) {
	r0, _, _ := syscall.Syscall6(procPdhGetFormattedCounterValue.Addr(), 4, uintptr(counter), uintptr(format), uintptr(unsafe.Pointer(counterType)), uintptr(unsafe.Pointer(value)), 0, 0)
	if r0 != 0 {
		errcode = syscall.Errno(r0)
	}
	return
}

func _PdhGetFormattedCounterValueLarge(counter PdhCounterHandle, format PdhCounterFormat, counterType *uint32, value *PdhCounterValueLarge) (errcode error) {
	r0, _, _ := syscall.Syscall6(procPdhGetFormattedCounterValue.Addr(), 4, uintptr(counter), uintptr(format), uintptr(unsafe.Pointer(counterType)), uintptr(unsafe.Pointer(value)), 0, 0)
	if r0 != 0 {
		errcode = syscall.Errno(r0)
	}
	return
}

func _PdhGetFormattedCounterValueLong(counter PdhCounterHandle, format PdhCounterFormat, counterType *uint32, value *PdhCounterValueLong) (errcode error) {
	r0, _, _ := syscall.Syscall6(procPdhGetFormattedCounterValue.Addr(), 4, uintptr(counter), uintptr(format), uintptr(unsafe.Pointer(counterType)), uintptr(unsafe.Pointer(value)), 0, 0)
	if r0 != 0 {
		errcode = syscall.Errno(r0)
	}
	return
}

func _PdhCloseQuery(query PdhQueryHandle) (errcode error) {
	r0, _, _ := syscall.Syscall(procPdhCloseQuery.Addr(), 1, uintptr(query), 0, 0)
	if r0 != 0 {
		errcode = syscall.Errno(r0)
	}
	return
}

func _PdhExpandWildCardPath(dataSource *uint16, wildcardPath *uint16, expandedPathList *uint16, pathListLength *uint32) (errcode error) {
	r0, _, _ := syscall.Syscall6(procPdhExpandWildCardPathW.Addr(), 4, uintptr(unsafe.Pointer(dataSource)), uintptr(unsafe.Pointer(wildcardPath)), uintptr(unsafe.Pointer(expandedPathList)), uintptr(unsafe.Pointer(pathListLength)), 0, 0)
	if r0 != 0 {
		errcode = syscall.Errno(r0)
	}
	return
}

func _PdhExpandCounterPath(wildcardPath *uint16, expandedPathList *uint16, pathListLength *uint32) (errcode error) {
	r0, _, _ := syscall.Syscall(procPdhExpandCounterPathW.Addr(), 3, uintptr(unsafe.Pointer(wildcardPath)), uintptr(unsafe.Pointer(expandedPathList)), uintptr(unsafe.Pointer(pathListLength)))
	if r0 != 0 {
		errcode = syscall.Errno(r0)
	}
	return
}

func _PdhGetCounterInfo(counter PdhCounterHandle, text uint16, size *uint32, lpBuffer *byte) (errcode error) {
	r0, _, _ := syscall.Syscall6(procPdhGetCounterInfoW.Addr(), 4, uintptr(counter), uintptr(text), uintptr(unsafe.Pointer(size)), uintptr(unsafe.Pointer(lpBuffer)), 0, 0)
	if r0 != 0 {
		errcode = syscall.Errno(r0)
	}
	return
}

func _PdhEnumObjectItems(dataSource uint16, machineName uint16, objectName *uint16, counterList *uint16, counterListSize *uint32, instanceList *uint16, instanceListSize *uint32, detailLevel uint32, flags uint32) (errcode error) {
	r0, _, _ := syscall.Syscall9(procPdhEnumObjectItemsW.Addr(), 9, uintptr(dataSource), uintptr(machineName), uintptr(unsafe.Pointer(objectName)), uintptr(unsafe.Pointer(counterList)), uintptr(unsafe.Pointer(counterListSize)), uintptr(unsafe.Pointer(instanceList)), uintptr(unsafe.Pointer(instanceListSize)), uintptr(detailLevel), uintptr(flags))
	if r0 != 0 {
		errcode = syscall.Errno(r0)
	}
	return
}

func _PdhGetFormattedCounterArray(counter PdhCounterHandle, format PdhCounterFormat, bufsize *uint32, bufcnt *uint32, item *PdhFmtCounterValueItem) (errcode error) {
	r0, _, _ := syscall.Syscall6(procPdhGetFormattedCounterArrayW.Addr(), 5, uintptr(counter), uintptr(format), uintptr(unsafe.Pointer(bufsize)), uintptr(unsafe.Pointer(bufcnt)), uintptr(unsafe.Pointer(item)), 0)
	if r0 != 0 {
		errcode = syscall.Errno(r0)
	}
	return
}
