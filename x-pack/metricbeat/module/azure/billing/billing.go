// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package billing

import (
	"time"

	"github.com/elastic/beats/v7/x-pack/metricbeat/module/azure"

	"github.com/pkg/errors"

	"github.com/elastic/beats/v7/libbeat/logp"
	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/beats/v7/metricbeat/mb/parse"
)

// init registers the MetricSet with the central registry as soon as the program
// starts. The New function will be called later to instantiate an instance of
// the MetricSet for each host defined in the module's configuration. After the
// MetricSet has been created then Fetch will begin to be called periodically.
func init() {
	mb.Registry.MustAddMetricSet("azure", "billing", New, mb.WithHostParser(parse.EmptyHostParser))
}

// MetricSet holds any configuration or state information. It must implement
// the mb.MetricSet interface. And this is best achieved by embedding
// mb.BaseMetricSet because it implements all of the required mb.MetricSet
// interface methods except for Fetch.
type MetricSet struct {
	mb.BaseMetricSet
	client *Client
	log    *logp.Logger
}

// Config options
type Config struct {
	ClientId                string        `config:"client_id"    validate:"required"`
	ClientSecret            string        `config:"client_secret" validate:"required"`
	TenantId                string        `config:"tenant_id" validate:"required"`
	SubscriptionId          string        `config:"subscription_id" validate:"required"`
	Period                  time.Duration `config:"period" validate:"nonzero,required"`
	ResourceManagerEndpoint string        `config:"resource_manager_endpoint"`
	ActiveDirectoryEndpoint string        `config:"active_directory_endpoint"`
}

func (conf *Config) Validate() error {
	if conf.ResourceManagerEndpoint == "" {
		conf.ResourceManagerEndpoint = azure.DefaultBaseURI
	}
	if conf.ActiveDirectoryEndpoint == "" {
		ok, err := azure.AzureEnvs.HasKey(conf.ResourceManagerEndpoint)
		if err != nil {
			return err
		}
		if ok {
			add, err := azure.AzureEnvs.GetValue(conf.ResourceManagerEndpoint)
			if err != nil {
				return err
			}
			conf.ActiveDirectoryEndpoint = add.(string)
		}
		if conf.ActiveDirectoryEndpoint == "" {
			return errors.New("no active directory endpoint has been configured")
		}
	}
	return nil
}

// New creates a new instance of the MetricSet. New is responsible for unpacking
// any MetricSet specific configuration options if there are any.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	var config Config
	err := base.Module().UnpackConfig(&config)
	if err != nil {
		return nil, errors.Wrap(err, "error unpack raw module config using UnpackConfig")
	}
	if err != nil {
		return nil, err
	}
	// instantiate monitor client
	billingClient, err := NewClient(config)
	if err != nil {
		return nil, errors.Wrap(err, "error initializing the billing client: module azure - billing metricset")
	}
	return &MetricSet{
		BaseMetricSet: base,
		client:        billingClient,
	}, nil
}

// Fetch methods implements the data gathering and data conversion to the right metricset
// It publishes the event which is then forwarded to the output. In case
// of an error set the Error field of mb.Event or simply call report.Error().
func (m *MetricSet) Fetch(report mb.ReporterV2) error {
	results, err := m.client.GetMetrics()
	if err != nil {
		return errors.Wrap(err, "error retrieving usage information")
	}
	events := EventsMapping(results)
	for _, event := range events {
		isOpen := report.Event(event)
		if !isOpen {
			break
		}
	}

	return nil
}
