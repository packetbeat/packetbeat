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

// Code generated by cgo -godefs - DO NOT EDIT.
// cgo -godefs defs_constants_linux.go

package seccomp

const prSetNoNewPrivs = 0x26

const (
	seccompSetModeFilter = 0x1
)

const x32SyscallMask = 0x40000000

const (
	ActionKillThread  Action = 0x0
	ActionKillProcess Action = 0x80000000
	ActionTrap        Action = 0x30000
	ActionErrno       Action = 0x50000
	ActionTrace       Action = 0x7ff00000
	ActionLog         Action = 0x7ffc0000
	ActionAllow       Action = 0x7fff0000
)

const errnoEPERM = 0x1

const (
	FilterFlagTSync FilterFlag = 0x1

	FilterFlagLog FilterFlag = 0x2
)
