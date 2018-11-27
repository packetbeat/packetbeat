// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package aws

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "aws", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "eJx8j7HOgyAYRXee4sbd5R8Z/q1P0XQgctsQQQhg0LdvlLZRY3vG+4VzQoues4QqSQDZZEuJRpXUCEAzddGEbPwg8S8AgFNgNI5DVvZ6E+umSoLzerQUwN3Q6iTXQ4tBOb7lC3kOlHhEP4bXctLYS7Yidn+f7Uz2VVjZPz9GdqFJubD+Z0sN9pyLj/pw+5FduFRhjYpnAAAA//+4aWAv"
}
