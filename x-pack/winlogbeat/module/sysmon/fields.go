// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package sysmon

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("winlogbeat", "sysmon", asset.ModuleFieldsPri, AssetSysmon); err != nil {
		panic(err)
	}
}

// AssetSysmon returns asset data.
// This is the base64 encoded gzipped contents of module/sysmon.
func AssetSysmon() string {
	return "eJxUzrFuwzAQA9BdX0Fkjz9AQ6fOXVKgs+qjEaG2zr07N9DfF1WTISsJEu+ML/YM775pS0DUWJlxuowAm8qx8pQAoc9W96jaMl4SALxf6UQxIq4Ef9gCS+UqDt8516XOCB3l092UAOPK4sz4ZJSE+y6ncXxGKxsfqkmaTx4lDh8tEH1n/oPf1OSePfk+ahO9Of5XmFUIYxzWKFjUhun17YLvg9an9BsAAP//OQhWnA=="
}
