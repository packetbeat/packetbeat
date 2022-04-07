// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

//go:build integration && azure
// +build integration,azure

package monitor

import (
	"testing"

	"github.com/elastic/beats/v8/x-pack/metricbeat/module/azure/test"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/beats/v8/libbeat/common"

	mbtest "github.com/elastic/beats/v8/metricbeat/mb/testing"
)

func TestFetchMetricset(t *testing.T) {
	config := test.GetConfig(t, "monitor")
	metricSet := mbtest.NewReportingMetricSetV2Error(t, config)
	events, errs := mbtest.ReportingFetchV2Error(metricSet)
	if len(errs) > 0 {
		t.Fatalf("Expected 0 error, had %d. %v\n", len(errs), errs)
	}
	assert.NotEmpty(t, events)
	mbtest.TestMetricsetFieldsDocumented(t, metricSet, events)
}

func TestData(t *testing.T) {
	config := test.GetConfig(t, "monitor")
	config["resources"] = []map[string]interface{}{{
		"resource_query": "resourceType eq 'Microsoft.DocumentDb/databaseAccounts'",
		"metrics": []map[string]interface{}{{"namespace": "Microsoft.DocumentDb/databaseAccounts",
			"name": []string{"DataUsage", "DocumentCount", "DocumentQuota"}}}}}
	metricSet := mbtest.NewFetcher(t, config)
	metricSet.WriteEvents(t, "/")
}

func TestDataMultipleDimensions(t *testing.T) {
	config := test.GetConfig(t, "monitor")
	config["resources"] = []map[string]interface{}{{
		"resource_query": "resourceType eq 'Microsoft.KeyVault/vaults'",
		"metrics": []map[string]interface{}{{"namespace": "Microsoft.KeyVault/vaults",
			"name": []string{"Availability"}, "dimensions": []map[string]interface{}{{"name": "ActivityName", "value": "*"}}}}}}
	metricSet := mbtest.NewFetcher(t, config)
	metricSet.WriteEventsCond(t, "/", func(m common.MapStr) bool {
		if m["azure"].(common.MapStr)["dimensions"].(common.MapStr)["activity_name"] == "secretget" {
			return true
		}
		return false
	})
}
