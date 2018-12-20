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

package traefik

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "traefik", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "eJy00TGvmzAUBeCdX3GUpVPanaFSxg6tqqqdI2MOxA34Wtcmav59BYGE8FBe9J7eGX0T++PcLY4850hqWLljBiSXGubYjCebDCgZrbqQnPgcXzMA+H2ZQnmiRiKo/DvjCxoxJQrTGG+paJnU2ZgByoYmMkfBZDKgcmzKmA93beFNy7mhTzoH5qhVujCerDDeQukzf35OONA06XA9XlM8kPT5fnkHUiTjPEtUKu1k/BTHF7D7+Q30ZRDn0+z/y5qmLL1zcxeSa/k50t6NJ3sjvl4MHvD7/BkuhFTXap2Pqa8RziPSii/jKkUZg/jIVciyxCckv8b7Ftt71MqcY6W7a/fVYp4g9fnRtQW1L2j64KXsZjCnej/sp3v5m/dLdieqqXl1YNic82idVVlb1dwWk0ld3Fsplxu72aT4S7tW4mWw/4AuEaijDYPtfwAAAP//7KY6tg=="
}
