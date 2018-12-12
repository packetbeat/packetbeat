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
	return "eJyUVMtuGzsM3fsrDnLX8QfM4gJBWqCrrvIDtERHRPSYSpy48/eFNBonNvyAvRvJ50HyiM/44HmASSGkuAFU1POA1/XbcjFZRpUUB/y/AYDXFJUklg7CXtjbAvok8bTzDIkg78GfHBU6j1y2G/S/DZvG8YxIgQcE1iymsG5DspPndnlRtf7eHDcc0h7qGAsG6kjxzpEzKdt207S317Tq94NKK/ZxMZeK3hb7lYqeiJFxEhn7nAIOTow783Cg2nzv2SjbLd6clCNZazMCzYhJsWOMmUsdxMFxbDyWlE4p4JMh7+erNWRdS6jjHOBTfO8Hmf9MktkO0Dzd6erPloicpmihWUaohBaXICanwiZF27KycIyZTW1yp74xzDKS4ROHHzwfUra3/fxeobXzdo4UxHwxl7N2VOJbRfNfCqP/7q3cD5lNZgrrO9nixR9oLmhBS3iyyTyduSicP8Xw9wwfhdlT0apM2bj7pa9565Td945JsWc1jstXXmoWz5wssbdTpkp+LR9X9H902OphWRYSESkmHKNwSZCUCuvj017S1+GNsvP/hxcv7T2EkVR2XnTu6+qC/smiWuSpwvvJSOou7LV/AQAA///QnK6i"
}
