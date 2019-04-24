// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package envoyproxy

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "envoyproxy", asset.ModuleFieldsPri, AssetEnvoyproxy); err != nil {
		panic(err)
	}
}

// AssetEnvoyproxy returns asset data.
// This is the base64 encoded gzipped contents of module/envoyproxy.
func AssetEnvoyproxy() string {
	return "eJykk0Fu2zAQRfc8xT9AfQEtChSuC3TRTY2uBVYcyYQpDkMOnTCnDyTKdmQYAZwsNeJ/bzQabnCk0oD8iUuI/FIUIFYcNdi9rxlKXbRBLPsG3xUA/GGTHaHniIP2xlk/wPGQECKb3JHB/1LBCugtOZMaBWzg9Ug3SkBKoAZD5Bzm5ztC4NdMQR95rHnMgKrVvVCE5zhqZ1/1lJxjV/VV7nhoJ+PCrfIjlWeOZqndbQB1KlN+Dn1bfK7gx3a72+/VShMpBfaJ2t7pIT0u+7vkUfMrdg5JIumxTRRPtqNW7Lj+Hsd+WAr91KU0MDleBwNYH7K055dee07UsTfp47b+LWosakxqWL8i3AziKVOS1prHh/D7J7iHHOhMWaN1lgNHK+Wz/7Ku0AVT92XtmI98aWGqpK6MdAHTpREJSqnLQfUWAAD//zjwF5k="
}
