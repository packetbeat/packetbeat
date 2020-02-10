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

package nginx

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "nginx", asset.ModuleFieldsPri, AssetNginx); err != nil {
		panic(err)
	}
}

// AssetNginx returns asset data.
// This is the base64 encoded gzipped contents of module/nginx.
func AssetNginx() string {
	return "eJy0mM+O4zYMxu95CmJ73fED5FCgWGCBPbQoih56yygS7RCriC5JTydvX9j5MxlHythJRqfEDr/vR5GS5TzBT9wtITWUXhcARhZxCV/+6L9/WQAEVC/UGnFawq8LAIDfOXQRoWaB1olSasA2CEMIRG6gpohaLQB0w2Irz6mmZgkmHS4AasIYdDlIPUFyW3yz74ftWlxCI9y1hysZhn58H4SgFt6WAPpx7nfu6bxH1dPlnPEV835842SOkh4shhl5A9nr9zwnlBzOOZLglg1X1K4iqb37yRHPibjd6M4VxH78lvZRwPXBAX78CS4EQVXUCn4YkIKD3hTW6F2nCDRc9LzdcgJjoORjF/ArrFEpoA6Z+kiYxqBwJv/1ndW+Vht0AUUh0k+E53+evrP85yRg6D89Vxdqf6GLoNyJH8BJQVCNBUPP9by/U1E7Cv0FKDW975PnZMIxooC26KkmfyhDtgpdqybotpWgtpwUq4ipsU22HpFTM68cf28QjsqQUT5ibMzaSvDfDtUeR5BVN9piVjtwt47jWxPUW+HXXXWayP5iVt/wddw989VdNJTkjF5w9WCnT2yFskWxGo8yUHPW6cpzeKTPu46i8BlVoPYzVFuW/F47fR4uP5w7rTnsVorJqvXOUPP7eiQ3vtM625zm9VC7Xisrs6VG3J7r8KDNpKwo5SVSRujjcot4iucWbcP5bvgw530vZRUmpStxbqISKxZqKLlx6BTDHnv1gqLE6ZaM86FTnI/9UV7UUxusvDlM46hRBOWeehc0pti75vI4MqW5V0Pg3NKX11iZY3y8hMKR8Fwy4Av5cTWup5ZNb6+TW8jXshzTFMJnstwJwTmTmQisVd3FmNsQ56GUFObz3IuS799bWMpKs1ZDg1x4aN+yEPqDPCVMds+MH94XGuTqQ72pE++5Sya7FSnnds6b0D5UnAoX2Q8/ux/qitJUGMGGOD2oftfFJhePbPeohroiNXOGHtdKHwuW0I5IKMLnj+MH/1MyyM/6o8RzSuhtSCt/tpz/Xv7tpAkUMBnVhFJd394ivuDcM2bkpsrFTTnftIVsy26tsEfV6jJyip/d7GcbQRcu3wOnvTKoumbuMTYfNfb7PwAA//9DGWAT"
}
