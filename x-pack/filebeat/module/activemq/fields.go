// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package activemq

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "activemq", asset.ModuleFieldsPri, AssetActivemq); err != nil {
		panic(err)
	}
}

// AssetActivemq returns asset data.
// This is the base64 encoded zlib format compressed contents of module/activemq.
func AssetActivemq() string {
	return "eJysksFO80AMhO95ilFP/39oHyAHJC7cioQEZ2RtnO2qm2xqe4v69mhDCwktqkTxKZrY801iL7HlQw1yFvbc7SrAgkWusbgfpfXTogKEI5NyDU8V0LA6CYOF1Ne4qwBgnZocGW0SDCQaeo/TPGLyaENkXVVAGzg2Wo9DS/TU8Qxeyg5DIUnKw1G5AJw7Td0cxcjyKZ/8tnx4S9JM9IuuH/VIHSO1sA0f/RBUc/msIsXkfXkW3mVWwz8XSRVJIKwpi+P/q7NcthGm5rZcz6MHbEMGzz0LGTezSLzn3s7hWW/9JS/K8gsw5SbYGXm63Cvch3HNaCV1Xyc1uha0ribNP11ETP7vEwxDDI5K/9Uc0yxq5LavJuR49v77Tt4DAAD//1NHAoY="
}
