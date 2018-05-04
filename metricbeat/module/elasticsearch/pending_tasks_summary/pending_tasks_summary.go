package pending_tasks_summary

import (
	"errors"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/common/cfgwarn"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/metricbeat/helper"
	"github.com/elastic/beats/metricbeat/mb"
	"github.com/elastic/beats/metricbeat/mb/parse"
	"github.com/elastic/beats/metricbeat/module/elasticsearch"
)

// init registers the MetricSet with the central registry.
// The New method will be called after the setup of the module and before starting to fetch data
func init() {
	mb.Registry.AddMetricSet("elasticsearch", "pending_tasks_summary", New, hostParser)
}

var (
	hostParser = parse.URLHostParserBuilder{
		DefaultScheme: "http",
		PathConfigKey: "path",
		DefaultPath:   "_cluster/pending_tasks",
	}.Build()
)

// MetricSet holds any configuration or state information. It must implement
// the mb.MetricSet interface. And this is best achieved by embedding
// mb.BaseMetricSet because it implements all of the required mb.MetricSet
// interface methods except for Fetch.
type MetricSet struct {
	mb.BaseMetricSet
	http *helper.HTTP
}

// New creates a new instance of the MetricSet. New is responsible for unpacking
// any MetricSet specific configuration options if there are any.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	cfgwarn.Experimental("The elasticsearch pending_tasks metricset is experimental.")

	http, err := helper.NewHTTP(base)
	if err != nil {
		return nil, err
	}

	return &MetricSet{
		base,
		http,
	}, nil
}

// Fetch methods implements the data gathering and data conversion to the right format
func (m *MetricSet) Fetch() (common.MapStr, error) {
	isMaster, err := elasticsearch.IsMaster(m.http, m.HostData().SanitizedURI)
	if err != nil {
		return nil, err
	}

	// Not master, no event sent
	if !isMaster {
		logp.Debug("elasticsearch", "Trying to fetch pending tasks summary from a none master node.")
		return nil, errors.New("None master node")
	}

	content, err := m.http.FetchContent()
	if err != nil {
		return nil, err
	}

	event, _ := eventMapping(content)

	event[mb.NamespaceKey] = "cluster.pending_tasks_summary"

	return event, nil
}
