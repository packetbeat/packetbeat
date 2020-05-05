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

package nats

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "nats", asset.ModuleFieldsPri, AssetNats); err != nil {
		panic(err)
	}
}

// AssetNats returns asset data.
// This is the base64 encoded gzipped contents of module/nats.
func AssetNats() string {
	return "eJzUmM+S2zYMxu9+CkxO7SH7AD50ptNecmgO3fS8oSlY5kQkFAD0xnn6DiXLlmjJlhtlt6ujKeP76QP4B3wPX/CwhmBUVgDqtMI1vPv4+6fHdyuAAsWyq9VRWMNvKwBo3oS/qIgVrgAYKzSCayjNCmDrsCpk3bz3HoLxeIqcHj3U6U2mWB9/GYmfns/pT5/BUlDjgoCoUSfqrIDujMIzMgKjKWDL5OHjWaJP0KcQ5D3ygytOIx3OFzw8E/d/n4BKz6cdHkPBhz+nRNR5vJApjOI8jcc2fooCtAWPys6CZTTp7QtRSyGgTUNyIdq3+obqH53XTX7PQVOGjWJx5OhrDHPfPXkG+qxKaqrBSEfqgmKJnI1d4e2yEaLfICejbGTGoNUBjFW3R7CVw6By4RhTVFzQrCbe2/GJsXSiyFjkTpzqWI0u6M/e8KLuxDqbYWfGikKZDWyJvdE1FJGHE2i2dTWyowL0PPedQKzhF0H76yihR/+wOQyLbBbk2J9mEB5LP+kSHyCKKZu1Iy3jUDNZFBkFtcQTkAsUWkWls6ZqRRr7+jzAMQhk+Thx1XGUSqypsHjaVmR0wsIa2WLIR+8w0dZx4KBcdbCZqk/jizBcS/ldVjYq1aHdBbC4WNr6RIyepkpv4eVjRKmjcPlEG19EZsobTzFoknfBknehTBuqyUtgZNGAwbSUlNfcGriWphmAVyAnFTuksQl/k+faavFjvJcRO1CK+ZRaKKEUtaT/e0JPkG8koSfe6YRKRc9p3ZLokX/OqpUk4CTRO6BRaLaCUa6dal5PP1RpKV537Djtie0+fme1MX59yo9GtxFnYHaojF8jivb7nbFimcDso0Z2o+O3QGfCZsBQOFF2m9g0DBTAU3BKnArwn78/PE7EuPYhkJ/Xv0++dmNu3flZcFHFO6fpnHLE6H/btMvD9uz12RuKe9H3hl+fPEHcCy5xI69P3lDci85E+T778uQJYgq873GnsmCvOAj7dnrq443DlCn9Y7EgX+wgy2y3bWygGttOV8CFzNDKiU42Dfv5/eqdHUOK/d+4vFG7+0lc6jwKmFYj9fRbiqFIhzEwA77xFtXYHT6I+z5+E7FIpyWxUhDUxrHUQDeiV3B2Tp94eNMIL9I8J9EuWWAYYYNp5eA0cXGPx4vaW/hbEyjqgzfflvXUm2/ORw9t/Pb0V8DmMBfI7PNCO17qUtxUudlzup89silxiuffAAAA//+DHv/E"
}
