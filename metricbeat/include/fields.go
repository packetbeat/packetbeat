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

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package include

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "../metricbeat/_meta/fields.common.yml", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "eJyUlEFu2zAQRfc+xUf20QG0KFCkBbrqKhcYk+OQCMlRyZFV3T4gJTmxEduwdzQ1//+ZedIz3nnuYSRGSTtAvQbu8bKdLReT/aBeUo8fOwB4kaTkU1mLcPAcbAEdyQfaB4ZPoBDAR04KnQcu3Q7rY/2uaTwjUeQekTV7U1i7KHYM3C6/da2/V8etDnKAOsZSA3WkeOPEmZRtu2ne3TWven7Qaat93MxJ0dtmf6TomRkZ5xPjkCVict64iwwT1eGHwEbZdnh1vpzE2pgRaUYSxZ4xZC51EZPj1HQsKZ1LIIihEOarPWTdWqjr7BEkva1/ZP43+sy2h+bxzlR/NyKyjMlCsx+gPjZcojdZChtJtrGyaAyZTR3yKn1jmWUgw2cJ33meJNvbef5upXXydk4UvflULhfjqMK3mub/FIfwNVu5D5kVM8btPenwM0w0FzTQBE9WzNNFisL56A1/ZfhkzIGKVmfKxt1vfeNtlVxz75kUB1bjuHzyUlm8SLJgb8dMVfwaH1f8f61lW4blY+ETEiXBCYXvDEmpsD6+7YW+tbxJdruPAAAA///0dI4s"
}
