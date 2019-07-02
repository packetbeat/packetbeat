// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package oracle

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "oracle", asset.ModuleFieldsPri, AssetOracle); err != nil {
		panic(err)
	}
}

// AssetOracle returns asset data.
// This is the base64 encoded gzipped contents of module/oracle.
func AssetOracle() string {
	return "eJzEVdtu2jAYvucpPvWmm0R5gFxMYitISLRIbVVpV9UP+QNWHTvzoS19+sluoCFxaKdNzBcQ+fB/B/uzL/DI2wza0EryAHDCSc5wtogdZwMgZ7syonJCqwxv3cjJ0ZIso9S5j+vsRhv3sNKqEOsMBUkbeg1LJssZluxoABSCZW6zAQBcQFHJDejQ3LbiDGujfVX3pOD3qKE1azbrOlpKthWteD+Uqt/BSK5rC9m1NniTQPg9GNjBP/L2WZu8NXZA4m5P4rBMEijsxkMhZBqtLbaDdbnbzFDCQqhCm5LCWGtVSmyTh2hLeichtVonBvs0eyV+eYbIWTlRCDa9mAmT8YHRHeCpkBzqQBdwm7fTjYSf76hWvPajpgzvYN6KV25aDVpq7yJ8Ernf/Savkl5Gy61jm5z1wWZEkMgnw7EiBzqu6EWUvoykoy0Q6sjqHdHTkoxm15sbiX6KY2GYT+zmt55JwN2mtrcpg55IyJCagARv2cSzO4qzaeU8ye6iUihv4TbC4olkSJkNS3M4Deu0qacZluQ4R8mOYtH+MDhyPiXvj0KYln4+jecqImQY349n8/H3+QTaYHZ9P57PLvFl91EyqSCM3mME5cslmyBRaRe23VseRrf4hcpK8hC0OxJJAtR4D95qP5NFbnRVcf71vNcUraRQ/PBvvJmTdXhU+lnVdWtHOhfWCAsVd/v25+1iOh2G/7vJ1RCL6XQ+u54MsbgO/8G/m8mPxf3kZnT8bWk/hPj8u9K4z+tb3dKa/+J9ORrIo2H8KIh9vJ12JCNuLWGfuOH+DukPRkjVafnGHMfPz/CL6v6HoeFOahL8HQAA//8FUMIb"
}
