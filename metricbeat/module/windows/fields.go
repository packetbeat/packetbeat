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

package windows

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "windows", asset.ModuleFieldsPri, AssetWindows); err != nil {
		panic(err)
	}
}

// AssetWindows returns asset data.
// This is the base64 encoded gzipped contents of module/windows.
func AssetWindows() string {
	return "eJysVl1v2zYUffevuMhLgSIx1vZpfiiQxcvmYUmLJoUxIIDFkFfWXShS5b20Y2A/fqAkO5Y/YscoHwKFpM49555D0RfwhIsBzMkZP+cegJBYHMDZuJk56wEYZB2oEvJuAJ97AAA33kSLkPsA49WrXPggE+1dTtMB5Moy9gACWlSMA5iqHkBOaA0PapALcKrE9eJpyKJKm4OPVTuzo34XaB2swpCX3q3mdwGm0aW1HDtLNWMbeJPCCwlyLMpp7AA0PJ5wMffBdFY6Vf/rLAGMWiyYKRuxv6NaiRJIc/99//2Ogv7xX9TSWWimJs16br3auzwpVVWRm7Z7z96f7Sf+eYP4TU2roc0QUGJwaF4EbFnHGGbUadpu616pmrUYGWjvRJFjkAKBRUnk9bgui3H/YCa2bV4z2myI3m8yAD6rskqHqxh/uv3rWn+4rDZ2vNpPgEuIjn5EhNGw1lJLa3T0YSRADAoKxQX4vF4slS7I4TuGP76PhqCcSdNbuC1GrWmHP+uC099TJI9R80y/Te79i7xjqBniyqrF5GSKbTJ+n6ETuPLWohYf3s65JVLTWjqxtOl1CSwqNAfvDQKO6WLCjVUNs8mo3lB5Znq0uDytKiBkl1F8qYR0dr6Fmv3mvWTnkA2J1aNFk55vlIvKZud10LK7BQuW2VGST/Xs9h4uv9//+eXb6P6fh7+9VvZu6yNyRI8utfbRSeNYdAYDzAvSBahVAEN0fEBKpaQ4WcnV4GE8uh1+Gd89cN24Tx8feKYLz9LHZ4SLJ1jXBxdv/HZcR2sX8CMqSzmhqcmC+DoKOVkEKZQAJTIlOuH1jGzbT07baMhNQYVprF847LP85FQrLVHZBvn4VF95J+QiuemuWH9Vkeul5rHJ9bfoXDt5l+K6evb15bgKfPofzaHE4zOlX0nmJ3bj2oeX8qurDaQgTpdCagyG4AOkso3Tq1xj5YPwFuS8QNcczuSyeOBWbWotMczJWnjEGnuKDtNVv3G3bmGucYjOIndCBlXwMzLJpuXUBVeoKSe99uahM7jnRrbeTfcdvQ+//PrxhH4vU7G734rZa1KSjlrwOon9OhoeYB8roRL75aYdezXkPpRKBmBiUInsxjK5Kspkuakka4lRe2c2Cxx/E7/jliW05qABch3sfu//AAAA//8t/Vii"
}
