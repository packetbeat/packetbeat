// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package zeek

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "zeek", asset.ModuleFieldsPri, Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "eJzMk89OwzAMxu97Cr/Axr0HpAHiBBe0E5cqTdwuLI2LnWrqnh6l0/6AApTmgm9N/fu+T5a9hB0OBRwQdwuAYIPDAl6PXwZFs+2CJV/A7QIA4JlM7xBqYtgqb5z1DThqBDom02s0UA0jfnPHtACoLTojxcjGWoJXLZ79ThWGDgtomPru6jXhf6rHURdqpvZsdwyi6oAMnrhVzh5UhK/Ir3kuiQRFLPnSmk9Gx2Q7HPbElz8JXpP3qKPdypFWriS2TUKqInKo/J+kGKXLlmqtCJqyGgJKQsyRbyYqSVABs6e0tRKIhwk6/3dpjJdVYOUlvTbWB2yQfxpIVOAQErChvnL4G/ve46QZfodrp2TOOlzgMj5kJIjdcwPE1kx/1mTm+o9spv96Pe+yI7u5n8++PGSwGZmVlz1yauUmDmyzeUrR086F8Q11wNSxnvJ/BAAA//+Ycccf"
}
