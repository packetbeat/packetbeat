// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package billing

import (
	"time"

	"github.com/stretchr/testify/mock"

<<<<<<< HEAD
=======
	"github.com/Azure/azure-sdk-for-go/services/consumption/mgmt/2019-10-01/consumption"
	"github.com/Azure/azure-sdk-for-go/services/costmanagement/mgmt/2019-11-01/costmanagement"

>>>>>>> 86b111d594 ([Azure Billing] Switch to Cost Management API for forecast data (#32589))
	"github.com/elastic/beats/v7/x-pack/metricbeat/module/azure"

	"github.com/elastic/beats/v7/libbeat/logp"

	prevConsumption "github.com/Azure/azure-sdk-for-go/services/consumption/mgmt/2019-01-01/consumption"
	"github.com/Azure/azure-sdk-for-go/services/consumption/mgmt/2019-10-01/consumption"
)

// MockService mock for the azure monitor services
type MockService struct {
	mock.Mock
}

// NewMockClient instantiates a new client with the mock billing service
func NewMockClient() *Client {
	return &Client{
		new(MockService),
		azure.Config{},
		logp.NewLogger("test azure monitor"),
	}
}

<<<<<<< HEAD
// GetForcast is a mock function for the billing service
func (service *MockService) GetForcast(filter string) (consumption.ForecastsListResult, error) {
	args := service.Called(filter)
	return args.Get(0).(consumption.ForecastsListResult), args.Error(1)
=======
// GetForecast is a mock function for the billing service
func (service *MockService) GetForecast(scope string, startTime, endTime time.Time) (costmanagement.QueryResult, error) {
	args := service.Called(scope, startTime, endTime)
	return args.Get(0).(costmanagement.QueryResult), args.Error(1)
>>>>>>> 86b111d594 ([Azure Billing] Switch to Cost Management API for forecast data (#32589))
}

// GetUsageDetails is a mock function for the billing service
func (service *MockService) GetUsageDetails(scope string, expand string, filter string, skiptoken string, top *int32, apply string) (prevConsumption.UsageDetailsListResultPage, error) {
	args := service.Called(scope, expand, filter, skiptoken, top, apply)
	return args.Get(0).(prevConsumption.UsageDetailsListResultPage), args.Error(1)
}
