// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package stan

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "stan", asset.ModuleFieldsPri, AssetStan); err != nil {
		panic(err)
	}
}

// AssetStan returns asset data.
// This is the base64 encoded gzipped contents of module/stan.
func AssetStan() string {
	return "eJy8l0Fv4zYQhe/+FYPtJQGSzd2HLYIkCwRIXbROT0WxS4kjmzA1VDijNdxfX1AyFVmik6CVqkNgyPK8j0+TN+Q17PCwBBZFCwAxYnEJn9ai6NMCQCPn3lRiHC3hywIAmifhF6driwsAjxYV4xIyFLUAKAxazcvmyWsgVWJXO1xyqHAJG+/q6ngnoRCu7+FH3yF3JMoQhxpiWEzOIFslsEeP4FFpKLwrYaWEYS0eVWloA4z+B3q4WD/fri6PNftkJ3TNs5+N7r6JmDs87J3v3z8DG67nLUbZx/uRSG5rlolUjrWSMltFhJZHIn3H35G461lOcAMUrOVo7U27xr8BSVfOkECJ4k3elxz2RLyGb6BPHv6efHHennf4o02hIrgCZIsQ+uB1DdGmJEiJzGqDnISxjjb/gqQuM/SBZcCR1Iog2UHmpGjKg6HGn7cMKYxn+cb4MhHL11APGF9qpLzDYnEe9YDnMzwWr/rwBUpDF3+Gj4aA66zT4b8uQStRYB0zbBVD5ZhNZg/g8rz2HnVybVZNurQn9dGVJWk0VrKdCOW3GmtsK0KmGDXUlSNovKAxpCINW7PZIgt4rJwX1NHhZhGnz49DVJRMkzqO7KFBubb4A206hWLSerQqkE6TQWERE4dQzOqbfvo4Ksym9ir8sBVtJ9X97dOvq4cruHv6Y/388PvD/WUS0zs7IeVjESFRXzW129A0DOQ0dn17XMjFEyqN/gq+OmvdPny6U6SNVoJp3NwaJEknmSHBTa+dPuzrIMxSGt1b7efErBS99jyr+caknsGWlMhMg27VScfCoHIf4lhZ2+K8hEyaf9i9krQzLnfEdYn6IzzjYDvzJmffVvWFJ8o3M4yGqXdYx347Zb9o/zm/GX0mHxITcT60cOMEzzAoZpebZpDszWACR8qmV+ZjXN0+r1uN9njztqVj5iswBSg6dEI//U9bJ5Ip9047cnsabU6OHp1Y0B4CVQiZ3RnCCkmbEcZ/T5Vj3deca06gNyCuj5kN8jtSuaKwhtKtlDlnsTsrfxDskQeyUCq/C3HHUeznc9sdazGdCRORtO9THGQY5U5ZFv8EAAD//z+0hbY="
}
