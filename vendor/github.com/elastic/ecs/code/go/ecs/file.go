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

// Code generated by scripts/gocodegen.go - DO NOT EDIT.

package ecs

import (
	"time"
)

// A file is defined as a set of information that has been created on, or has
// existed on a filesystem. File objects can be associated with host events,
// network events, and/or file events (e.g., those produced by File Integrity
// Monitoring [FIM] products or services). File fields provide details about
// the affected file associated with the event or metric.
type File struct {
	// Path to the file.
	Path string `ecs:"path"`

	// Target path for symlinks.
	TargetPath string `ecs:"target_path"`

	// File extension.
	// This should allow easy filtering by file extensions.
	Extension string `ecs:"extension"`

	// File type (file, dir, or symlink).
	Type string `ecs:"type"`

	// Device that is the source of the file.
	Device string `ecs:"device"`

	// Inode representing the file in the filesystem.
	Inode string `ecs:"inode"`

	// The user ID (UID) or security identifier (SID) of the file owner.
	UID string `ecs:"uid"`

	// File owner's username.
	Owner string `ecs:"owner"`

	// Primary group ID (GID) of the file.
	Gid string `ecs:"gid"`

	// Primary group name of the file.
	Group string `ecs:"group"`

	// Mode of the file in octal representation.
	Mode string `ecs:"mode"`

	// File size in bytes (field is only added when `type` is `file`).
	Size int64 `ecs:"size"`

	// Last time file content was modified.
	Mtime time.Time `ecs:"mtime"`

	// Last time file metadata changed.
	Ctime time.Time `ecs:"ctime"`
}
