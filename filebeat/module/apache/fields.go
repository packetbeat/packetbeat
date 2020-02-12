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

package apache

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "apache", asset.ModuleFieldsPri, AssetApache); err != nil {
		panic(err)
	}
}

// AssetApache returns asset data.
// This is the base64 encoded gzipped contents of module/apache.
func AssetApache() string {
	return "eJy0lz9v2zAQxXd/CiJ7OHT0UKBoUXRogQLJLtDkWSJC89TjyYG+faF/qf9QMim7ngIp996PvCeC9yzeoN0KVStdwUYItuxgK56+9A+eNkIYCJpszRb9VnzeCCHE8FL8QtO4rihUSFxo9HtbbgVT0z3cW3AmbPuCZ+HVASabT/0zIbitYStKwqYen0S8ej9nVYAg9khip/TbuyIjNB5qxXZnneVWvFuuBDozWYz2o8QpyxmP1hDCx+MYUqz8VILggAyFrc/eTkqqI794UyuutiJgQxqkMobOEbrfwZakhl0Yd/PaOAQna0JGjS7Te9gjOaxeLgilcmhbV0B3U0RlUhh2aNoigGe5axku7W6BVMy1JAg1+gCy04rKpIA0Aajo/sxE6OpkpC7F8wBcoVm35j8NBJZRhaTlUm70GnISyZbWqzVh67CLI1Cw6NesOF6a4jzlo9Bocrt7HrDAipsQ00nj2ANR9sd21u8ZjRR7VYLnFeEu+sLU1s9/W/P+l4e2mDm4TyUNHK2+7MLykqLLGnRiH3BsdXM0M+WZLHdCYMwkEwGD3DfOxQ7CPJQ5hXyee1HiuV3DMq+U8v2VgDNXjTX51+jZevB8z0aPN5kSUN7US91vjY1nagsbMHZQrkK7qZgK51D3/3Y/1IJSKgxBadE/qH/LYsnNs9w+KlALUpk79Lgo3RacQ5uQgAjprrHDwRFy714OSxmrS7tkhqDK3ItPvCrFr7a5N9qasB8kritT/Hi1H1cEyshVrodpgs4xHuemPkMyqnB1nYoM4dkz+DDzD5mUj5mrZ6y631f0rKwPo0U/+nMFE8aP19ff4gXoCDSaden+4IqxiZzZ+Q3ad6TLli7wdr+Xl59iUhXjpPGPaM3wvBpj0OxN5MrTZ8Hl+9gTwsNpU3rdrD4s5H/V0l8rGDW7RphGW1/2hA7LEsx0Hl1vSQiuGKej/5zab53XyswWj89st2NnuW0CGGEamnZOo/eg+SzKl1iPTvAENaY4BWnzNwAA//9aRU8e"
}
