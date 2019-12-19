// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build mips64 mips64le

package cpu

<<<<<<< HEAD:vendor/golang.org/x/sys/cpu/cpu_mips64x.go
const cacheLineSize = 32

func doinit() {}
=======
import "runtime"

const cacheLineSize = 64

func init() {
	switch runtime.GOOS {
	case "darwin":
		// iOS does not seem to allow reading these registers
	case "android", "linux":
		doinit()
	default:
		readARM64Registers()
	}
}

func readARM64Registers() {
	Initialized = true

	// ID_AA64ISAR0_EL1
	isar0 := getisar0()

	switch extractBits(isar0, 4, 7) {
	case 1:
		ARM64.HasAES = true
	case 2:
		ARM64.HasAES = true
		ARM64.HasPMULL = true
	}

	switch extractBits(isar0, 8, 11) {
	case 1:
		ARM64.HasSHA1 = true
	}

	switch extractBits(isar0, 12, 15) {
	case 1:
		ARM64.HasSHA2 = true
	case 2:
		ARM64.HasSHA2 = true
		ARM64.HasSHA512 = true
	}

	switch extractBits(isar0, 16, 19) {
	case 1:
		ARM64.HasCRC32 = true
	}

	switch extractBits(isar0, 20, 23) {
	case 2:
		ARM64.HasATOMICS = true
	}

	switch extractBits(isar0, 28, 31) {
	case 1:
		ARM64.HasASIMDRDM = true
	}

	switch extractBits(isar0, 32, 35) {
	case 1:
		ARM64.HasSHA3 = true
	}

	switch extractBits(isar0, 36, 39) {
	case 1:
		ARM64.HasSM3 = true
	}

	switch extractBits(isar0, 40, 43) {
	case 1:
		ARM64.HasSM4 = true
	}

	switch extractBits(isar0, 44, 47) {
	case 1:
		ARM64.HasASIMDDP = true
	}

	// ID_AA64ISAR1_EL1
	isar1 := getisar1()

	switch extractBits(isar1, 0, 3) {
	case 1:
		ARM64.HasDCPOP = true
	}

	switch extractBits(isar1, 12, 15) {
	case 1:
		ARM64.HasJSCVT = true
	}

	switch extractBits(isar1, 16, 19) {
	case 1:
		ARM64.HasFCMA = true
	}

	switch extractBits(isar1, 20, 23) {
	case 1:
		ARM64.HasLRCPC = true
	}

	// ID_AA64PFR0_EL1
	pfr0 := getpfr0()

	switch extractBits(pfr0, 16, 19) {
	case 0:
		ARM64.HasFP = true
	case 1:
		ARM64.HasFP = true
		ARM64.HasFPHP = true
	}

	switch extractBits(pfr0, 20, 23) {
	case 0:
		ARM64.HasASIMD = true
	case 1:
		ARM64.HasASIMD = true
		ARM64.HasASIMDHP = true
	}

	switch extractBits(pfr0, 32, 35) {
	case 1:
		ARM64.HasSVE = true
	}
}

func extractBits(data uint64, start, end uint) uint {
	return (uint)(data>>start) & ((1 << (end - start + 1)) - 1)
}
>>>>>>> update golang.org/x/sys:vendor/golang.org/x/sys/cpu/cpu_arm64.go
