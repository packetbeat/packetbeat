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

package vsphere

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "vsphere", asset.ModuleFieldsPri, AssetVsphere); err != nil {
		panic(err)
	}
}

// AssetVsphere returns asset data.
// This is the base64 encoded zlib format compressed contents of module/vsphere.
func AssetVsphere() string {
	return "eJzsWE9vnD4Qve+nGOWe/QAcftJPqdJetqmUJtfVxAyLG4ORPeyKfvrKBjb8bXaDUS71ASXYvPdgPO/B3sIrVREcbZGSoQ0AS1YUwc3x0Z+52QDEZIWRBUudR/DfBgCgmYVMx6VylxlShJYiOOAGIJGkYhv5pbeQY0ZdCje4Ktxio8uiOTPB0gfqgsXIaFmf4aYhZ2GbqQmQ/n20YyijK8UdexOtkleqTtrEg7m/6HHjS6tpjNsSJtbhh6O8l4psZZkyGAG3nAILFJKrLWtGtX2pmOykAqXzw3X0Px0ieETQCXBKk4VxI9EmQ45gTD/SmRiioDLvDVFwlaWlOKjKJ0vxOioLwZMarUBF8T5RGocLLtFakBGU86Vqm+Xn2ZExpNryEk8YXP/ZdvBNW553AlGUdXGy9HfIDXT34wlkDrsBape29oFwvLULXEDsGzscr2/rd2gzyrRZq1l3HtzRT0G/36WNuLV8OZC8dex4qbic+KTN6979Ny3sQ037vYaFPuz5PUgaLlFlKFKZL3p/mUW63rWc7W3l8D6XGtcAsEfmDit5ZRsmAy9fy6Of6zLArq7DvF3rgHtsSPpQkEGW+QEe61e5f6GxXmj8f0Sp8EVdlRyHkiyvlh86ga+OYLFTe62+Q1eUWrtDkMgL/1h7wRfsufotGF5sNweXaxWlZZ3t66SYVKhfftHoS6A+uV9gZ3eeGCaIPyWru8SWkcuAjA9HMqgUpISKU8A8bihc/bANdRin+puksmA5k2Af2O4pNYBtbj7v3A6yJHQ+U4zzZ9o2cDn6P39YQGu1kMgUw0ly6uXNPaG3P/4EAAD//wn7xB4="
}
