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

package mysql

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("packetbeat", "mysql", asset.ModuleFieldsPri, Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "eJy8k0FvGyEQhe/7K55y6aXOD9hDL5UPkdJKrXO3MAw26gI2MzTaf1/BktTbrFtblcJlV8B78wFvVvhBYw8/8mnoAHEyUI+7L+Pm2+NdBxhindxRXAw9PnUAUNdWfCTtrNOgnxQE1tFg+L5D++vr1hWC8vTbvgwZj9Rjn2I+tplzxbnKMaUU0+v8i3YX40AqnM0vULbxYCEHmphxypRGJJKcAkMFVP+PkIPjiQKOwSSQCEmZ7rs3UMpa0kJmm+Izv0EbYthfxTUH09F7FabqWWtitnmYcekYRLnAVfTCMDMM2e8oIVoUsvItWwfFAhYl5CnIwnlcYEqydeYm7Ievm/X3p3ahV1I706Dm2PQ8jI2CTGFfgAzZb6eU3ES5WT+uP/+b8uzND9TucWYXbctoyw6ZC4x/ZOJdCeur/4Wv1rgO7ulAxW5q2wanirkysCn6SiFJBVa66D+UtVMmXgpY7bGtjob+o1sKUTVCMXo9Jnbj1EEX63piVnu6tYoLNqJpF6v9CgAA///kpYOQ"
}
