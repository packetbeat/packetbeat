// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package billing

import (
	"time"

	"github.com/Azure/azure-sdk-for-go/services/consumption/mgmt/2019-10-01/consumption"
	"github.com/shopspring/decimal"

	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/elastic-agent-libs/mapstr"
)

func EventsMapping(subscriptionId string, results Usage, startTime time.Time, endTime time.Time) []mb.Event {
	var events []mb.Event
	if len(results.UsageDetails) > 0 {
		for _, ud := range results.UsageDetails {
			event := mb.Event{Timestamp: time.Now().UTC()}

			// shared fields
			event.RootFields = mapstr.M{
				"cloud.provider": "azure",
			}

			//
			// legacy data format
			//
			if legacy, isLegacy := ud.AsLegacyUsageDetail(); isLegacy {
				event.ModuleFields = mapstr.M{
					"resource": mapstr.M{
						"type":  legacy.ConsumedService,
						"group": legacy.ResourceGroup,
						"name":  legacy.ResourceName,
					},
					"subscription_id": legacy.SubscriptionID,
				}
				event.MetricSetFields = mapstr.M{
					"pretax_cost":          legacy.Cost,
					"department_name":      legacy.InvoiceSection,
					"product":              legacy.Product,
					"usage_start":          startTime,
					"usage_end":            endTime,
					"billing_period_start": legacy.BillingPeriodStartDate.ToTime(),
					"billing_period_end":   legacy.BillingPeriodEndDate.ToTime(),
					"currency":             legacy.BillingCurrency,
					"effective_price":      legacy.EffectivePrice,
					"account_name":         legacy.BillingAccountName,
					"account_id":           legacy.BillingAccountID,
					"subscription_name":    legacy.SubscriptionName,
					"unit_price":           legacy.UnitPrice,
					"quantity":             legacy.Quantity,
				}
				_, _ = event.RootFields.Put("cloud.region", legacy.ResourceLocation)
				_, _ = event.RootFields.Put("cloud.instance.name", legacy.ResourceName)
				_, _ = event.RootFields.Put("cloud.instance.id", legacy.ResourceID)
			}

			//
			// modern data format
			//
			if modern, isModern := ud.AsModernUsageDetail(); isModern {
				event.ModuleFields = mapstr.M{
					"resource": mapstr.M{
						"type":  modern.ConsumedService,
						"group": modern.ResourceGroup,
						"name":  modern.InstanceName,
					},
					"subscription_id": modern.SubscriptionGUID,
				}
				event.MetricSetFields = mapstr.M{
					"department_name":      modern.InvoiceSectionName,
					"product":              modern.Product,
					"usage_start":          startTime,
					"usage_end":            endTime,
					"billing_period_start": modern.BillingPeriodStartDate.ToTime(),
					"billing_period_end":   modern.BillingPeriodEndDate.ToTime(),
					"currency":             modern.BillingCurrencyCode,
					"account_id":           modern.BillingAccountID,
					"billing_account_name": modern.BillingAccountName,
					"subscription_name":    modern.SubscriptionName,
					"unit_price":           modern.UnitPrice,
				}
				_, _ = event.RootFields.Put("cloud.region", modern.ResourceLocation)
			}

			events = append(events, event)
		}
	}

	//
	// Forecasts
	//
	groupedCosts := make(map[*string][]consumption.Forecast)

	for _, forecast := range results.ForecastCosts {
		groupedCosts[forecast.UsageDate] = append(groupedCosts[forecast.UsageDate], forecast)
	}

	for _, forecast := range results.ActualCosts {
		groupedCosts[forecast.UsageDate] = append(groupedCosts[forecast.UsageDate], forecast)
	}

	for usageDate, items := range groupedCosts {
		var actualCost *decimal.Decimal
		var forecastCost *decimal.Decimal
		for _, item := range items {
			if item.ChargeType == consumption.ChargeTypeActual {
				actualCost = item.Charge
			} else {
				forecastCost = item.Charge
			}
		}

		parsedDate, err := time.Parse("2006-01-02", *usageDate)
		if err != nil {
			parsedDate = time.Now().UTC()
		}

		event := mb.Event{
			RootFields: mapstr.M{
				"cloud.provider": "azure",
			},
			ModuleFields: mapstr.M{
				"subscription_id": subscriptionId,
			},
			MetricSetFields: mapstr.M{
				"actual_cost":   actualCost,
				"forecast_cost": forecastCost,
				"usage_date":    parsedDate,
				"currency":      items[0].Currency,
			},
			Timestamp: time.Now().UTC(),
		}

		//event.ID = generateEventID(parsedDate)
		events = append(events, event)
	}

	return events
}
