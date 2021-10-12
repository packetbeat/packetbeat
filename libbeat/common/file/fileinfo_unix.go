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

//go:build !windows
// +build !windows

package file

import (
	"errors"
	"os"
	"syscall"
)

func stat(name string, statFunc func(name string) (os.FileInfo, error)) (FileInfo, error) {
	info, err := statFunc(name)
	if err != nil {
		return nil, err
	}

	return wrap(info)
}

func wrap(info os.FileInfo) (FileInfo, error) {
	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return nil, errors.New("failed to get uid/gid")
	}

	uid := int(stat.Uid)
	gid := int(stat.Gid)
	return fileInfo{FileInfo: info, uid: &uid, gid: &gid}, nil
}
