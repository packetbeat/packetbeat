// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package sysmon

import (
	"github.com/menderesk/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("winlogbeat", "sysmon", asset.ModuleFieldsPri, AssetSysmon); err != nil {
		panic(err)
	}
}

// AssetSysmon returns asset data.
// This is the base64 encoded zlib format compressed contents of module/sysmon.
func AssetSysmon() string {
	return "eJysjrFOKzEQRXt/xVX6+ANcvOo1NDRBokSOfVcZ4djBM5uwf49iEqGVIiraGd1zzhbvXAJ00WOrDjCxwoDNbhxwbHku3DggU1OXk0mrAf8cALwcqETshB0InlkNk7BkhZ6YZJIEa+O5wnkHdBZGZcCeFh1uu+AGeIsaj7xX+VzVq0WbdXwBW04M1/BL6/l2W/W9Ss3tovheIbVMdNrcKzOm1kfT/+cdPmb2xT+0TlLoY08HOTOvxPvWCmN9JH6qWVI0KmQaksxCu0qlEJeouBN/kYq+8ZNptrgv/ENzxQ/Wu68AAAD//xH3plw="
}
