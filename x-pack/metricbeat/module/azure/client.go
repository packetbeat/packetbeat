// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package azure

import (
	"fmt"
	"strings"
	"time"

	"github.com/elastic/beats/libbeat/common"

	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2019-06-01/insights"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-03-01/resources"
	"github.com/pkg/errors"

	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/metricbeat/mb"
)

// Client represents the azure client which will make use of the azure sdk go metrics related clients
type Client struct {
	AzureMonitorService Service
	Config              Config
	Resources           ResourceConfiguration
	Log                 *logp.Logger
}

// mapMetric function type will map the configuration options to client metrics (depending on the metricset)
type mapMetric func(client *Client, metric MetricConfig, resource resources.GenericResource) ([]Metric, error)

// mapMetric function type will map the configuration options to client metrics (depending on the metricset)
type initEvent func(event *mb.Event, str common.MapStr) error

// NewClient instantiates the an Azure monitoring client
func NewClient(config Config) (*Client, error) {
	azureMonitorService, err := NewService(config.ClientID, config.ClientSecret, config.TenantID, config.SubscriptionID)
	if err != nil {
		return nil, err
	}
	client := &Client{
		AzureMonitorService: azureMonitorService,
		Config:              config,
		Log:                 logp.NewLogger("azure monitor client"),
	}
	client.Resources.RefreshInterval = config.RefreshListInterval
	return client, nil
}

// InitResources function will retrieve and validate the resources configured by the users and then map the information configured to client metrics.
// the mapMetric function sent in this case will handle the mapping part as different metric and aggregation options work for different metricsets
func (client *Client) InitResources(fn mapMetric, report mb.ReporterV2) error {
	if len(client.Config.Resources) == 0 {
		return errors.New("no resource options defined")
	}
	// check if refresh interval has been set and if it has expired
	if !client.Resources.Expired() {
		return nil
	}
	var metrics []Metric
	for _, resource := range client.Config.Resources {
		// retrieve azure resources information
		resourceList, err := client.AzureMonitorService.GetResourceDefinitions(resource.ID, resource.Group, resource.Type, resource.Query)
		if err != nil {
			err = errors.Wrap(err, "failed to retrieve resources")
			return err
		}
		if len(resourceList.Values()) == 0 {
			err = errors.Errorf("failed to retrieve resources: No resources returned using the configuration options resource ID %s, resource group %s, resource type %s, resource query %s",
				resource.ID, resource.Group, resource.Type, resource.Query)
			client.Log.Error(err)
			continue
		}
		for _, res := range resourceList.Values() {
			for _, metric := range resource.Metrics {
				mapCustomFields(&metric, resource)
				met, err := fn(client, metric, res)
				if err != nil {
					return err
				}
				metrics = append(metrics, met...)
			}
		}
	}
	// users could add or remove resources while metricbeat is running so we could encounter the situation where resources are unavailable we log an error message (see above)
	// we also log a debug message when absolutely no resources are found
	if len(metrics) == 0 {
		client.Log.Debug("no resources were found based on all the configurations options entered")
	}
	client.Resources.Metrics = metrics
	return nil
}

// GetMetricValues returns the specified metric data points for the specified resource ID/namespace.
func (client *Client) GetMetricValues(report mb.ReporterV2) error {
	// loop over the set of metrics
	for i, metric := range client.Resources.Metrics {
		// select period to collect metrics, will double the interval value in order to retrieve any missing values
		endTime := time.Now().UTC()
		startTime := endTime.Add(client.Config.Period * (-2))
		timespan := fmt.Sprintf("%s/%s", startTime.Format(time.RFC3339), endTime.Format(time.RFC3339))

		// build the 'filter' parameter which will contain any dimensions configured
		var filter string
		if len(metric.Dimensions) > 0 {
			var filterList []string
			for _, dim := range metric.Dimensions {
				filterList = append(filterList, dim.Name+" eq '"+dim.Value+"'")
			}
			filter = strings.Join(filterList, " AND ")
		}
		resp, err := client.AzureMonitorService.GetMetricValues(metric.Resource.SubID, metric.Namespace, metric.TimeGrain, timespan, metric.Names,
			metric.Aggregations, filter)
		if err != nil {
			err = errors.Wrapf(err, "error while listing metric values by resource ID %s and namespace  %s", metric.Resource.SubID, metric.Namespace)
			client.LogError(report, err)
		} else {
			current := mapMetricValues(resp, client.Resources.Metrics[i].Values, endTime.Truncate(time.Minute).Add(client.Config.Period*(-1)), endTime.Truncate(time.Minute))
			client.Resources.Metrics[i].Values = current
		}
	}
	return nil
}

// LogError is used to reduce the number of lines written when logging errors
func (client *Client) LogError(report mb.ReporterV2, err error) {
	client.Log.Error(err)
	report.Error(err)
}

// CreateMetric function will create a client metric based on the resource and metrics configured
func (client *Client) CreateMetric(selectedResourceID string, resource resources.GenericResource, namespace string, metrics []string, aggregations string, dimensions []Dimension, timegrain string) Metric {
	met := Metric{
		Resource: Resource{
			SubID:        selectedResourceID,
			ID:           *resource.ID,
			Name:         *resource.Name,
			Location:     *resource.Location,
			Type:         *resource.Type,
			Group:        getResourceGroupFromID(*resource.ID),
			Tags:         mapTags(resource.Tags),
			Subscription: client.Config.SubscriptionID},
		Namespace:    namespace,
		Names:        metrics,
		Dimensions:   dimensions,
		Aggregations: aggregations,
		TimeGrain:    timegrain,
	}
	for _, prevMet := range client.Resources.Metrics {
		if len(prevMet.Values) != 0 && matchMetrics(prevMet, met) {
			met.Values = prevMet.Values
		}
	}
	return met
}

// MapMetricByPrimaryAggregation will map the primary aggregation of the metric definition to the client metric
func MapMetricByPrimaryAggregation(client *Client, metrics []insights.MetricDefinition, resource resources.GenericResource, selectedResourceID string, namespace string, dim []Dimension, timegrain string) []Metric {
	var clientMetrics []Metric
	metricGroups := make(map[string][]insights.MetricDefinition)

	for _, met := range metrics {
		metricGroups[string(met.PrimaryAggregationType)] = append(metricGroups[string(met.PrimaryAggregationType)], met)
	}
	for key, metricGroup := range metricGroups {
		var metricNames []string
		for _, metricName := range metricGroup {
			metricNames = append(metricNames, *metricName.Name.Value)
		}
		if selectedResourceID == "" {
			selectedResourceID = *resource.ID
		}
		clientMetrics = append(clientMetrics, client.CreateMetric(selectedResourceID, resource, namespace, metricNames, key, dim, timegrain))
	}
	return clientMetrics
}

// GroupMetricsByAllDimensions will group metrics by dimension names in order to reduce the number of api calls
func GroupMetricsByAllDimensions(metrics []insights.MetricDefinition) map[string][]insights.MetricDefinition {
	var groupedMetrics = make(map[string][]insights.MetricDefinition)
	for _, metric := range metrics {
		if metric.Dimensions != nil {
			for _, dimension := range *metric.Dimensions {
				if _, ok := groupedMetrics[*dimension.Value]; !ok {
					groupedMetrics[*dimension.Value] = make([]insights.MetricDefinition, 0)
				}
				groupedMetrics[*dimension.Value] = append(groupedMetrics[*dimension.Value], metric)
			}
		} else {
			if _, ok := groupedMetrics[NoDimension]; !ok {
				groupedMetrics[NoDimension] = make([]insights.MetricDefinition, 0)
			}
			groupedMetrics[NoDimension] = append(groupedMetrics[NoDimension], metric)
		}
	}
	return groupedMetrics
}
