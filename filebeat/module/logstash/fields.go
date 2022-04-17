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

package logstash

import (
	"github.com/menderesk/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "logstash", asset.ModuleFieldsPri, AssetLogstash); err != nil {
		panic(err)
	}
}

// AssetLogstash returns asset data.
// This is the base64 encoded zlib format compressed contents of module/logstash.
func AssetLogstash() string {
	return "eJzsVk1v2zAMvedXED23+QE59LKtQIF9ANvuBmPTshZJFPTR1P9+kB0ntiN3TdsVGDDdIlp871FPZG5gR+0GFAsf0DcrgCCDog1cDVtXKwBHitDTBgSuACrypZM2SDYbuF0BwPE8fOEqKloB1JJU5Tdd9AYMapqgpBVam1I6jvawk8k8zTTLdtw70f48on2MnQEtgvXrroOE2rGG0BAMSTsF69Gnc25jfnooxXj1VHbU7tlVs9gThNL62dAhJ7CDUqH3sG/IUUeRHsgEYCeFNBhonaUUGkc4h30FpXtTs9OYwoBbjqGj4qIx0ogD2oijYrHEEEBHFWSRK+hEAT2Gs+Cg4Sw4skrRFSgrnbe/qJyH/qB8Ry2gqeABVSSoaBuFSJrlqSL5KzhSWWOZvnr2ZQwJrLSkpKFCXnKR9IjapgeiUc4xn+G7+4/AdXeHA/xJXd785D2KvPtRSfSziMXQLJ3SUjjs2QUXKV9WeiB1IZpisc6dW8IbsLzi/az5XNpgzlP87yT/RCdZ7iIvkv0d91BFbYfXddCksjh/U5dVUUhTpB9vp+4rajr2jQ7gKewE9IYeb+0MewP3xsbgr+FOqkDOX8O3GNJOegUfuKLSL5ideVdIU2iplJz3kp6jYiMuI/jpkcrY+T1ITVCzG3EFaaBHo5JNtcDrUDiLDnWe1otK9yO49Oj60TYpIZRsaili3x7f35690iI7sF83y29up3onQiF6qmDbjiqxMNzfYQpBxpgGDecNsIzb/wmpFu5yBv47AAD//5HKKPI="
}
