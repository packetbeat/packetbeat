// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build freebsd openbsd netbsd

package imports

import "syscall"

func direntInode(dirent *syscall.Dirent) uint64 {
	return uint64(dirent.Fileno)
}
