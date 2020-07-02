// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package gsuite

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "gsuite", asset.ModuleFieldsPri, AssetGsuite); err != nil {
		panic(err)
	}
}

// AssetGsuite returns asset data.
// This is the base64 encoded gzipped contents of module/gsuite.
func AssetGsuite() string {
	return "eJy0VMFu4zYQvfsrBttDgQUioT3qUMAo3KLYpFkk2aJ7EmhyZA1McVhy6F316wuKVqA4NgpsvDxyZt68N+S8G9jj2MAuJhJcAQiJxQbelYt3KwCDUQfyQuwa+GUFAMdsuGOTbC7qCK2JzRS7AacGXCDmY7BTyUo7JTbQKRvnkIw+ZwdO/jn5VcN8fi9No0dNHelj02r1nHDHAYFcx2FQuRjUlpOcFoBWDrYIHSdnQAn0Ij42dW3wgJY9hljtmHcWK81DrcxA7iaafR3Qc5BYH36qA3YY0GmslRY6kBDG2lKUI5flPJYzUVo4VHscnyOz/j2OXziYxf2FKeRz7+wIPmBEJ/ClRwdaWYvhafQIFOHD5nMFvxaZ0iNodjENGNo9jsDddBfwn4RROEDHAe7XSXr4+fYe1h//mGMROIByQAadUEdYcgNvWUBpzclJrC6IzLLepvKpx6ksMy6Yqxfxv5RNOD9n8yIE8P7T4+bhfQNrx9JjgBQxALlJelQDguFBkatOyzZ/P20e/lzftnN9qeQkkUyZ5oXKD5vPU75jd9OnQbmZ9Ol88IBOrjyfgrkI41c1+LzIKhmSH47fdHxFZk/OXI9G+TyRU9D4LWw47JSjf6ftrcqc30au0CtIIL2SvB6q61ALGtiOx1XIi/1jPJ3izCqqwb6isTQsOLPySwDlvSU9qWrzzQnFS7r+VxvAoxosPH5cdph6Vmd5dIpsCtie/L03crjlHbkZewKq4LfsHJAdMX8LzzHS1iIcyspO9gnC32K9ynt0hr7OzjvW+XnO6yVHQkrQtNvxenofijtmV+jgcX13CypJny2yPMB5Lhx2yZG0Xkl/PS4AnyZ7KuDnO0dRkmKr2Vzz42XdBRky8oXWqNmZ1ubnbb8jj6kNTG1ekvovAAD//43JglU="
}
