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

package tls

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("packetbeat", "tls", asset.ModuleFieldsPri, AssetTls); err != nil {
		panic(err)
	}
}

// AssetTls returns asset data.
// This is the base64 encoded gzipped contents of protos/tls.
func AssetTls() string {
	return "eJzsWk9v27gSv+dTDHJ5LZA47+FhD5vDAoGTQ4CiLTYtdm8CLY6t2VCkSo7sqJ9+QUqyaf1x3MTJtl37FFvSzG/+/4bROdxjdQmsXCKRBSmUJwBMrPASTq+bn+DTu7vTEwCJLrVUMBl9Cb+dAADEt5y7AlOaUwq4RM0wJ1TSTU6g+esyPHEOWuQYdIbvAFwVeAkLa8qi+SW+P34mwth++g/7j8S5KBUnjSCYC+Uwut5VsFGxROvI6OhKq+Meq5WxcuvKgEc2n08ZtuLAzIEz9F6Cwho2qVFQOpSTrWfwQeRFcL2/8X+T/5+eDEC06Mo86Exy5MzIQ4C9rQE6dAFvJhzMEHWtDOVZuFpqiVZVpBdQa65tgA8awcw7Ek9JnsLc2GB0K/f2GoyFU6b0Hnlzuf4O+MCo/X2TIbtTRag5SdGyzzLBmFj8UqJjHHLBzBiFQu/vgj8y5Axt4we7RBvcsNYRLtQggA2IkjPUHIAAsUPV9UDpvKdE+0wEfJd9GSplBuzp5vhQFsMjmbw7QR71EOyR1bMKVhmlWeytFbkMHbAZkJeaPC917UVZWu8xzsi1GbPlqdi85npCfRueaeFnTV9KBF3mM58NBkj6OM+rrQrxqRtMNNaiK4yWpBeD9mmNqdcGK+LYLeOmlUVhLKNMUpMXtjG0rjj3EvFU5NgHM1LXFLiLw9jgcpMBMXeIkDEX7vLiYrVaTUhoMTF2cSGco4XOUbO78PLPveBzkp1vk4eMczXmkXVjGDO/Xxw9472hobYiaT5tlyRR+rSNI9OVNFxqW0ELHSPxXxLS0if0UPU9HrEe8HdNdDLj2It3XS/FMERRqEZ3okSFNmkrM9G4MEwHgzWUTP7Two2QnAck6x6xlVTkYEVKhbo3fp70ygIGqr4eGC9qBuoFZ22La6u+1nsGNF+nzpmfaEID5gVX4NgO94EAzYCQSz8GHLbVVDeSINbtNn3dFZr226+FQ5rfRNH39lYfcCb4GbFbGxCq9VXg3yjlb01hWtolwtRWBZuFFUVWwZub6fQtpOHCDkSwAd7tEruMpYUWXFpMhFoYS5zlr2LwWi9s9NZxy0UFM/RRAtIgaUEs1JjFrZTdKYlpUhjS7JK5sbngfyKkb26mbyGggAbEBG7r/ouh0YyZiEHa1pO9BE+FhkLYRxI7KXW4SSavYX80vQLcFVoEhXOGFoZP04/Cd5QZCh7kms20+sG5ZrCfXNhCGkIWM8e+wNtwu5eT0SJDx2vxvRpvNgBt/FqSIo7QO44gbjYF0nEWBS/vILIK02Gy9xL+7HO8NYKO6XNr8vDd08MBaWvu9G3W/rS0/dVIah2fp5DU12OH7xtxKGOlsE0Evw+m99l3D0/NtDalTrFuK6LD+QLT8QO0jcaItO0YwZVaicp1+eF3RPSiQEVcz/ujJgvyX8oAmib40zKA/mnaXjxgC+h08/DIJh36bXRS1pvLbc96ErFQZqB3byH885f//trEtpW1YzRZEiqpB8hzplMYsxvX/Mc1opvZNApAG05mODe2n7G1dtmNUk/1tQ9FLSM+hoviRC5wmqVQ1C/tGImY86gb9gQSRIzgwIeChvaLFkJRzhSlyT1Wmw3qJfjQWniHRW4HsEbjNQ3Ryvr8G36/uzqD67srMBZuptd3V/sY5+jrWLgfT+87+ootM44hjmV4fyl9NZf2SqLFMopWKEarBdMSk/rc6xlY78rZX5gyXG2EwnsvdDDiO45mg5inE7wWh0XPwlGvaXcXQFfK4+QuNaVmWx1iAk1rUZAauXvwGbsQmr4ejDZ+iOQFFXtrFyopNR2EKH7WxIH1k95SsRNLmH86HSYb3wjgYyPLNxKLi3YJId3GeCeS1OS50aFkDsLkRR6AZMZx0NAuVhQN+x2lE0NTJhWK+CA5+q6RNVqr5Fw5OsH2KNUbzcRVe8jgSr/KaBkaV/hn4LFijxV7rNjnVmznRPLF9pH2UOu4jxz3keM+Mm7ccR857iPHfeTIbo7s5riPPMMbx4o9VuwPWbHj+0iSZoKe9Gb41D8YeJEtHa8JSPsvpv3eC35JJHu+oSwUWk68KvcU5V1+dqVBWCt8k9AsSNfv/tbveARVQXAAiUu0VfOjxRRpiXJy8ncAAAD//9V6IOo="
}
