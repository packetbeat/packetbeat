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

package kvm

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "kvm", asset.ModuleFieldsPri, AssetKvm); err != nil {
		panic(err)
	}
}

// AssetKvm returns asset data.
// This is the base64 encoded zlib format compressed contents of module/kvm.
func AssetKvm() string {
	return "eJykk71uwyAUhXc/xVH29AEYOnWrsnan5TZCBmMBTuW3r+zY0QUTnChMlhHf+Q4/R7Q0CrQX2wBRR0MCh8+v06EBPBmSgQS+KcoGUBR+vO6jdp3AewNgWgfr1GCoAX41GRXEPHFEJy3N4Det5l9AHHsSMK47Lz8KxGl8OCt1h2UdR01fCayl8c959RDvtpgRE9jZu6GvoXhDDlLOWrIhynibKjErivNUiZIfwjpylVQog9SEdrWu40TW+bEMLrtwH3ZuJaf0FJ80y+zuZa0qF2mGmgu7ni+KlJJWC70tW4nfjU5fzDauuCf1vX80M0Hz+zeEVx7DhvDMQ+Aaee9a653OS+Mr9T8AAP//fuVDEQ=="
}
