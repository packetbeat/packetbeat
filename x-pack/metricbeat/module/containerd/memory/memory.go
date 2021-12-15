// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package memory

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/elastic/beats/v7/libbeat/common/cfgwarn"

	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/metricbeat/helper/prometheus"
	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/beats/v7/metricbeat/mb/parse"
	"github.com/elastic/beats/v7/x-pack/metricbeat/module/containerd"
)

const (
	defaultScheme = "http"
	defaultPath   = "/v1/metrics"
)

var (
	// HostParser validates Prometheus URLs
	hostParser = parse.URLHostParserBuilder{
		DefaultScheme: defaultScheme,
		DefaultPath:   defaultPath,
	}.Build()

	// Mapping of state metrics
	mapping = &prometheus.MetricsMapping{
		Metrics: map[string]prometheus.MetricMap{
			"container_memory_usage_max_bytes":           prometheus.Metric("usage.max"),
			"container_memory_usage_usage_bytes":         prometheus.Metric("usage.total"),
			"container_memory_usage_limit_bytes":         prometheus.Metric("usage.limit"),
			"container_memory_usage_failcnt_total":       prometheus.Metric("usage.fail.count"),
			"container_memory_kernel_max_bytes":          prometheus.Metric("kernel.max"),
			"container_memory_kernel_usage_bytes":        prometheus.Metric("kernel.total"),
			"container_memory_kernel_limit_bytes":        prometheus.Metric("kernel.limit"),
			"container_memory_kernel_failcnt_total":      prometheus.Metric("kernel.fail.count"),
			"container_memory_swap_max_bytes":            prometheus.Metric("swap.max"),
			"container_memory_swap_usage_bytes":          prometheus.Metric("swap.total"),
			"container_memory_swap_limit_bytes":          prometheus.Metric("swap.limit"),
			"container_memory_swap_failcnt_total":        prometheus.Metric("swap.fail.count"),
			"container_memory_total_inactive_file_bytes": prometheus.Metric("inactiveFiles"),
			"container_memory_total_active_file_bytes":   prometheus.Metric("activeFiles"),
			"container_memory_total_cache_bytes":         prometheus.Metric("cache"),
			"container_memory_total_rss_bytes":           prometheus.Metric("rss"),
		},
		Labels: map[string]prometheus.LabelMap{
			"container_id": prometheus.KeyLabel("id"),
		},
	}
)

// Metricset for containerd memory is a prometheus based metricset
type metricset struct {
	mb.BaseMetricSet
	prometheusClient prometheus.Prometheus
	mod              containerd.Module
	calcPct          bool
}

// init registers the MetricSet with the central registry.
// The New method will be called after the setup of the module and before starting to fetch data
func init() {
	mb.Registry.MustAddMetricSet("containerd", "memory", New,
		mb.WithHostParser(hostParser),
		mb.DefaultMetricSet(),
	)
}

// New creates a new instance of the MetricSet. New is responsible for unpacking
// any MetricSet specific configuration options if there are any.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	cfgwarn.Beta("The containerd cpu metricset is beta.")

	pc, err := prometheus.NewPrometheusClient(base)
	if err != nil {
		return nil, err
	}
	config := containerd.DefaultConfig()
	if err := base.Module().UnpackConfig(&config); err != nil {
		return nil, err
	}

	mod, ok := base.Module().(containerd.Module)
	if !ok {
		return nil, fmt.Errorf("must be child of kubernetes module")
	}
	return &metricset{
		BaseMetricSet:    base,
		prometheusClient: pc,
		mod:              mod,
		calcPct:          config.CalculatePct,
	}, nil
}

// Fetch gathers information from the containerd and reports events with this information.
func (m *metricset) Fetch(reporter mb.ReporterV2) error {
	families, _, err := m.mod.GetContainerdMetricsFamilies(m.prometheusClient)
	if err != nil {
		return errors.Wrap(err, "error getting families")
	}
	events, err := m.prometheusClient.ProcessMetrics(families, mapping)
	if err != nil {
		return errors.Wrap(err, "error getting events")
	}

	for _, event := range events {

		// setting ECS container.id
		rootFields := common.MapStr{}
		containerFields := common.MapStr{}
		var cID string
		if containerID, ok := event["id"]; ok {
			cID = (containerID).(string)
			containerFields.Put("id", cID)
			event.Delete("id")
		}

		if len(containerFields) > 0 {
			rootFields.Put("container", containerFields)
		}

		// Calculate memory total usage percentage
		if m.calcPct {
			inactiveFiles, err := event.GetValue("inactiveFiles")
			if err == nil {
				usageTotal, err := event.GetValue("usage.total")
				if err == nil {
					memoryLimit, err := event.GetValue("usage.limit")
					if err == nil {
						// calculate working set memory usage
						workingSetUsage := usageTotal.(float64) - inactiveFiles.(float64)
						workingSetUsagePct := workingSetUsage / memoryLimit.(float64)
						event.Put("workingset.pct", workingSetUsagePct)

						memoryUsagePct := usageTotal.(float64) / memoryLimit.(float64)
						event.Put("usage.pct", memoryUsagePct)
						m.Logger().Debugf("memoryUsagePct for %+v is %+v", cID, memoryUsagePct)
					}
				}
			}
		}

		reporter.Event(mb.Event{
			RootFields:      rootFields,
			MetricSetFields: event,
			Namespace:       "containerd.memory",
		})
	}
	return nil
}
