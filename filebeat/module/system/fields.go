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

package system

import (
	"github.com/menderesk/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "system", asset.ModuleFieldsPri, AssetSystem); err != nil {
		panic(err)
	}
}

// AssetSystem returns asset data.
// This is the base64 encoded zlib format compressed contents of module/system.
func AssetSystem() string {
	return "eJzsmM2O2zYQx+9+ioEvaYFED+BD0SJA0AD9AnYvPTlcciQRoTgCOVqv+vQF9bHRypRNad3msjoZkmd+fw05H+IH+IrtAXzrGasdAGs2eID9XXdjvwNQ6KXTNWuyB/hpBwDwO6nGIOTkoBbOa1sMDsBQAbk26LMdgC/J8VGSzXVxAHYN7gByjUb5Q+fnA1hR4YQeLm5rPEDhqKmHOxEF4frUeYLcUQVcYlRCuKbEKVU0XD7fjHEvsGP837Rtnjqv5PQ/IpgEMc86YlqmelhX6FlU9YunozBhtPCzJ7Xg8gDvfn62fDf7Q6ULJ3rtQ/jPsSV5Dr9WUoNZtmCbgq0dFU5UK6m1I4neZ1uhWm0Enlum8Cr0XhRrQxu3SuE1Ht1KWDBZE84o1/syip1nEyykwDRiXNI81BfTsL/uS4S7u1+79EPLWvb517vL4KOw8IBAFoFy2NfC+xM5tQdysK+bB6PlV2z32YIqrwsruHHzICUKezYP8FAqpNFoGXpwqMBLYOWorlEd9TyKY4QjD5IkDQo+/9XXrxA6SdaiDGYeuBQMIgiu0YKwCnRVodKC0bSjqiy+Kb6px0e0fCYAn0RVhybzi5RYM75itTsACA85NUGj7aIbii78MHp/D5/tozBavYdPQptwA1lmP14TvxjyWC6N2eSpcRKziO2l9J0UKHLnAUvGRq2TEvmbggJp8c1j6byU0FOfkixri5aPkUJz/R3P3rNAyhJ8Xg75S4GNZdcetaejJHUziQle00Ua6ovabcRd9JYuymGhyd5wXa85XLGomttbbriL7lZH7LZbLcXpuo7eKEpu6Vfq9TAo+xqlzrUEpr5we5AOBaOCh7Yr3F8C9QtIqiph1bwnXh4c0Dmazz6TXhPGHPjjz/vQIwIF3Xlkk9pOxxlHtOBNCo/990ejaNQOeddqlto6c7uNfn//N5xKdBGi9oBPKBtextanjb1WNs6FTqu0Q8nk2leIiMyok3VyROctLEkjC1cg9wvNBKdSyzIq0J80y1LbYknh8N+NoRpAYxjgUYtOwsK0NA2LUPHPk/805wbwxrQrKVIQp9oWwhTsICej0HXnCEGKxVP/ObI0hZdozBaYwlw0hnsHF3AL3GjRvz6ULX1Zpc6CzdknZzI1YpnGLDYxu/25BhrNgM7Ld0mBkbwxBzbujz5qr9gg/9NiTU7pDBW7OWvFidlHsiy09UM8XxzdGZoeIr4dmkWxb4dmrzw0+zcAAP//vfvB9g=="
}
