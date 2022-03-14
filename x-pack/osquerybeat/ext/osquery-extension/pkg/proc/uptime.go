// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package proc

import (
	"bytes"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
)

var (
	ErrInvalidProcUptimecontent = errors.New("invalid /proc/uptime content")
)

// ReadUptime Reads system uptime from <root>/proc/uptime
func ReadUptime(root string) (secs int64, err error) {
	return ReadUptimeFS(os.DirFS("/"), root)
}

func ReadUptimeFS(fsys fs.FS, root string) (secs int64, err error) {
	fp := filepath.Join(root, "/proc/uptime")
	b, err := fs.ReadFile(fsys, fp)
	if err != nil {
		return
	}
	detail := bytes.Split(b, []byte{' '})
	if len(detail) != 2 {
		return secs, ErrInvalidProcUptimecontent
	}

	num, err := strconv.ParseFloat(bytesToString(detail[0]), 64)
	if err != nil {
		return secs, ErrInvalidProcUptimecontent
	}
	return int64(num), nil
}
