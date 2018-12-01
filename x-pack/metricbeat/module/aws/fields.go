// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package aws

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "aws", asset.ModuleFieldsPri, Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "eJzElk+PGzcMxe/7KYicmzns0YcCjdF70DbIoSi8tER7CGtEQaQ8dT59IXl213a83T+JvXOUxuT7vSeP+BE2tJsBjnoDYGyBZvABR/1wA+BJXeZkLHEGv94AANzhqHcwiC+BwEkI5Ezht69/wiCRTTLHNQxkmZ3CKsvQ9uZBih/RXN+1KvRvoswDRcPw9z83bQ1HnereAKyYgtdZ2/gIEQe6l1gf2yWawTpLSdPKGaXHRQ4Lkbt9WDtX7MmC++eO3O0dOImGHBWspwdg69FgpEygLmMif2LB12oBjD27/rHAGeOUosFy1374+/y2O+h/CnUI5lLpTAxDl5wdvXGPqQ4D+cUqCJ6+8D/E9fmrJ0iUXc1sTSArwBDEoZGvEsHJkIoRlMg2GYGZwJWcKVrYAUcoSiCxOcZRDaOjR7QniVwmz7Yoimu6AFQsw5JyBZp//gL7ZgqapggOxcJKcnurGAf+hrXsywGWGGqR70Q8C/EsxikIYY7kj3j2gcRHlB4V0LlcyINyXWGDERUCluh68iAZ1DAb+RcwaskpFF1cnXVqfAza45ZgSRQfc8QIJQYeuB7YBxfGniLUn80/f5m3Cp/2ymGLoRCwwjfK8moDdOF6zGvy13CgIZ71of4DoxgkZA9exlid+P50/AIY/fRZsr4ocHQlV8vQe646MMAe6BknItkoedNx7BK6DZlewYCpE2RyxNt6cmP9ON2LAY5GeYWO9PQP/UIaKfYOOO0SkGI/HYdjt9wZXYOl9blsMNdF+emheNYNS5cJ/QVRPk054DSRVOkPH0E1yQRbCWUgBdwiB1wGApM3gIyZjS5OUrsYxSrxcigtE0mX4ZjLkALVy6ilIolymyf07QHVEQvrdeB4xeTruMbi69k1Hl4V3+WhW59D6jcm+UPMamhFO9eT2yxWyOHJ+zpIXL8e9w9Kkk3rlGE95WPhdQBLqEoelmL98eZeGTRl7XKuu7pTo+F4j/fjdEA1GDgWewPzYl/4XdAvynXf8N3Izud5ju2/AAAA//8FfHr6"
}
