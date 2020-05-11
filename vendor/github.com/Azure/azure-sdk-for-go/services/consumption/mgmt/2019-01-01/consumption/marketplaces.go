package consumption

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// MarketplacesClient is the consumption management client provides access to consumption resources for Azure
// Enterprise Subscriptions.
type MarketplacesClient struct {
	BaseClient
}

// NewMarketplacesClient creates an instance of the MarketplacesClient client.
func NewMarketplacesClient(subscriptionID string) MarketplacesClient {
	return NewMarketplacesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewMarketplacesClientWithBaseURI creates an instance of the MarketplacesClient client.
func NewMarketplacesClientWithBaseURI(baseURI string, subscriptionID string) MarketplacesClient {
	return MarketplacesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List lists the marketplaces for a scope at the defined scope. Marketplaces are available via this API only for May
// 1, 2014 or later.
// Parameters:
// scope - the scope associated with marketplace operations. This includes '/subscriptions/{subscriptionId}/'
// for subscription scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing
// Account scope, '/providers/Microsoft.Billing/departments/{departmentId}' for Department scope,
// '/providers/Microsoft.Billing/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope and
// '/providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope. For
// subscription, billing account, department, enrollment account and ManagementGroup, you can also add billing
// period to the scope using '/providers/Microsoft.Billing/billingPeriods/{billingPeriodName}'. For e.g. to
// specify billing period at department scope use
// '/providers/Microsoft.Billing/departments/{departmentId}/providers/Microsoft.Billing/billingPeriods/{billingPeriodName}'
// filter - may be used to filter marketplaces by properties/usageEnd (Utc time), properties/usageStart (Utc
// time), properties/resourceGroup, properties/instanceName or properties/instanceId. The filter supports 'eq',
// 'lt', 'gt', 'le', 'ge', and 'and'. It does not currently support 'ne', 'or', or 'not'.
// top - may be used to limit the number of results to the most recent N marketplaces.
// skiptoken - skiptoken is only used if a previous operation returned a partial result. If a previous response
// contains a nextLink element, the value of the nextLink element will include a skiptoken parameter that
// specifies a starting point to use for subsequent calls.
func (client MarketplacesClient) List(ctx context.Context, scope string, filter string, top *int32, skiptoken string) (result MarketplacesListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MarketplacesClient.List")
		defer func() {
			sc := -1
			if result.mlr.Response.Response != nil {
				sc = result.mlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil},
					{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("consumption.MarketplacesClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, scope, filter, top, skiptoken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.MarketplacesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.mlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "consumption.MarketplacesClient", "List", resp, "Failure sending request")
		return
	}

	result.mlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.MarketplacesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client MarketplacesClient) ListPreparer(ctx context.Context, scope string, filter string, top *int32, skiptoken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"scope": scope,
	}

	const APIVersion = "2019-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(skiptoken) > 0 {
		queryParameters["$skiptoken"] = autorest.Encode("query", skiptoken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Consumption/marketplaces", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client MarketplacesClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client MarketplacesClient) ListResponder(resp *http.Response) (result MarketplacesListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client MarketplacesClient) listNextResults(ctx context.Context, lastResults MarketplacesListResult) (result MarketplacesListResult, err error) {
	req, err := lastResults.marketplacesListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "consumption.MarketplacesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "consumption.MarketplacesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.MarketplacesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client MarketplacesClient) ListComplete(ctx context.Context, scope string, filter string, top *int32, skiptoken string) (result MarketplacesListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MarketplacesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, scope, filter, top, skiptoken)
	return
}
