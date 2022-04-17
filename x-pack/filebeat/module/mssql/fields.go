// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package mssql

import (
	"github.com/menderesk/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "mssql", asset.ModuleFieldsPri, AssetMssql); err != nil {
		panic(err)
	}
}

// AssetMssql returns asset data.
// This is the base64 encoded zlib format compressed contents of module/mssql.
func AssetMssql() string {
	return "eJxsj0FugzAQRfec4ivr5gJsK2XVqKo4gYEPtTJm6IxJxe0rKGoJrXf+nv/e+Iwb5xLJ/UMKIMcsLHFa76cCaOmNxTFHHUpcK1RvL7hEYc2QcdV2EhZAFymtlwUAnDGExF/kcvI8skRvOo1b8sC9rHV0pgn5nbhWi0a0RxeFvlX2kr1ItP/JDuBnTUmHjbS0d4N/d/pPshepxT4OD08H3+s6Ae3WbyS6h55PmHwKIvMaOu1OQz1lxIwmDAjiipoIMDZ6p80YTRu6H1TfG984f6q1xVcAAAD//2bhgvE="
}
