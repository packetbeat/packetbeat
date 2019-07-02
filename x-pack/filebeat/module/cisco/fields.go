// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package cisco

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "cisco", asset.ModuleFieldsPri, AssetCisco); err != nil {
		panic(err)
	}
}

// AssetCisco returns asset data.
// This is the base64 encoded gzipped contents of module/cisco.
func AssetCisco() string {
	return "eJy8l81u6zYQhfd5itkUaAHF3XtRIHAbwEBzE8DtoitjLjm02FAclRzZ129fkJJtWXUVGVGrRQAz0jkfR5wfPcI7HZegbFT8ACBWHC1h1f3UFFWwtVj2S/jpAQDghXXjCAwHKNFrZ/2uvR08yYHDO2jaW0XgeBcXDwDGktNxmR9+BI8VXezSJcealrAL3NTdyg3XdD1nITCBq87xZJGuvk3fCiOe126ZjRj2TTl0nk+bJ3i2gQ7o3KJ369D/QlBRjLijrdVXyi3KOx0PHK7/M4ID8FtJPZJOG6wmL9ZYChemGyixMcZ+m4hB37Cq02mIFKNlP53xNa+j6/wAjVCA7xLwVFBugqKt9ULBoKI5IrfJmnDWzC9VSgLj+AAcgPbkZRRLUxTrMenPy/bzRfhTgM5GmemgfcGKgE0GeFKKYoQVewns4FcbBaREgQpFlaRBShsn4HWvtYkU0sLcmEm35bIxL7R+XRwnEfbf8P+G2TO9h7XCuia9PeVKfYNzsPhhZZGAPjoU0qfYrd8AtQ4U4x0sNQe5QePY7z7Lk6SnkFyl6syh6b+v++LTp/ovgtQn+zBSUgZC2Trak5unNSU9yHr5HFfoDhgIfoSvLJ4kkRpj1QJefc6GPYXjo+NDAenPQK5iTQGFCijtrkz1L9+efkzZlkKhHYfjHDtbdVrnivzvO3tOdfrUOfc2NLHo7hnuTwL/ib4AEjW6H8Xek2oP8iyV/Xdv/2r6M0PeFuY2M0piVVVvk+kNilgOj/Mow3r18paf/NhQsZ7LMEkt/jEiWo6zjojr183VaAqj4yHm3rpNnXvuRpNqU9u5k/popA0q66xMzZbzCd/8srqvWJ2MQBgOpVVlm0/dHBvIUIjwvblkUQGbLy9vBWz+2BSAPjWCgazhIOUPC3i6iCv08JUAocSgc6a2XyUFINSBhRW7AvKpr9oPGjbD9EzzwzEKVRDZSBJZwFpAk2ehq/miKwoKm3iOffvosKS121w8/B0AAP//Krahjg=="
}
