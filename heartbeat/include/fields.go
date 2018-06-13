// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package include

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("heartbeat/fields.yml", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "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"
}
