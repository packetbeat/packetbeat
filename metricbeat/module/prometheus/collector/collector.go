// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package collector

import (
	"regexp"
	"sync"

	"github.com/pkg/errors"
	dto "github.com/prometheus/client_model/go"

	"github.com/elastic/beats/v7/libbeat/common"
	p "github.com/elastic/beats/v7/metricbeat/helper/prometheus"
	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/beats/v7/metricbeat/mb/parse"
)

const (
	defaultScheme = "http"
	defaultPath   = "/metrics"
)

var (
	// HostParser parses a Prometheus endpoint URL
	HostParser = parse.URLHostParserBuilder{
		DefaultScheme: defaultScheme,
		DefaultPath:   defaultPath,
		PathConfigKey: "metrics_path",
	}.Build()

	upMetricName          = "up"
	upMetricType          = dto.MetricType_GAUGE
	upMetricInstanceLabel = "instance"
	upMetricJobLabel      = "job"
	upMetricJobValue      = "prometheus"
)

func init() {
	mb.Registry.MustAddMetricSet("prometheus", "collector",
		MetricSetBuilder("prometheus", DefaultPromEventsGeneratorFactory),
		mb.WithHostParser(HostParser),
		mb.DefaultMetricSet(),
	)
}

// PromEventsGenerator converts a Prometheus metric family into a PromEvent list
type PromEventsGenerator interface {
	// Start must be called before using the generator
	Start()

	// converts a Prometheus metric family into a list of PromEvents
	GeneratePromEvents(mf *dto.MetricFamily) []PromEvent

	// Stop must be called when the generator won't be used anymore
	Stop()
}

// PromEventsGeneratorFactory creates a PromEventsGenerator when instanciating a metricset
type PromEventsGeneratorFactory func(ms mb.BaseMetricSet) (PromEventsGenerator, error)

// MetricSet for fetching prometheus data
type MetricSet struct {
	mb.BaseMetricSet
	prometheus     p.Prometheus
	includeMetrics []*regexp.Regexp
	excludeMetrics []*regexp.Regexp
	namespace      string
	promEventsGen  PromEventsGenerator
	once           sync.Once
	host           string
}

// MetricSetBuilder returns a builder function for a new Prometheus metricset using
// the given namespace and event generator
func MetricSetBuilder(namespace string, genFactory PromEventsGeneratorFactory) func(base mb.BaseMetricSet) (mb.MetricSet, error) {
	return func(base mb.BaseMetricSet) (mb.MetricSet, error) {
		config := defaultConfig
		if err := base.Module().UnpackConfig(&config); err != nil {
			return nil, err
		}
		prometheus, err := p.NewPrometheusClient(base)
		if err != nil {
			return nil, err
		}

		promEventsGen, err := genFactory(base)
		if err != nil {
			return nil, err
		}

		ms := &MetricSet{
			BaseMetricSet: base,
			prometheus:    prometheus,
			namespace:     namespace,
			promEventsGen: promEventsGen,
		}
		// store host here to use it as a pointer when building `up` metric
		ms.host = ms.Host()
		ms.excludeMetrics, err = compilePatternList(config.MetricsFilters.ExcludeMetrics)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to compile exclude patterns")
		}
		ms.includeMetrics, err = compilePatternList(config.MetricsFilters.IncludeMetrics)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to compile include patterns")
		}

		return ms, nil
	}
}

// Fetch fetches data and reports it
func (m *MetricSet) Fetch(reporter mb.ReporterV2) error {
	m.once.Do(m.promEventsGen.Start)

	families, err := m.prometheus.GetFamilies()
	eventList := map[string]common.MapStr{}
	if err != nil {
		// send up event only
		families = append(families, m.upMetricFamily(0.0))

		// set the error to report it after sending the up event
		err = errors.Wrap(err, "unable to decode response from prometheus endpoint")
	} else {
		// add up event to the list
		families = append(families, m.upMetricFamily(1.0))
	}

	for _, family := range families {
		if m.skipFamily(family) {
			continue
		}
		promEvents := m.promEventsGen.GeneratePromEvents(family)

		for _, promEvent := range promEvents {
			labelsHash := promEvent.LabelsHash()
			if _, ok := eventList[labelsHash]; !ok {
				eventList[labelsHash] = common.MapStr{}

				// Add default instance label if not already there
				if exists, _ := promEvent.Labels.HasKey("instance"); !exists {
					promEvent.Labels.Put("instance", m.Host())
				}
				// Add default job label if not already there
				if exists, _ := promEvent.Labels.HasKey("job"); !exists {
					promEvent.Labels.Put("job", m.Module().Name())
				}
				// Add labels
				if len(promEvent.Labels) > 0 {
					eventList[labelsHash]["labels"] = promEvent.Labels
				}
			}

			// Accumulate metrics in the event
			eventList[labelsHash].Update(promEvent.Data)
		}
	}

	// Report events
	for _, e := range eventList {
		isOpen := reporter.Event(mb.Event{
			RootFields: common.MapStr{m.namespace: e},
		})
		if !isOpen {
			break
		}
	}

	return err
}

// Close stops the metricset
func (m *MetricSet) Close() error {
	m.promEventsGen.Stop()
	return nil
}

func (m *MetricSet) upMetricFamily(value float64) *dto.MetricFamily {
	gauge := dto.Gauge{
		Value: &value,
	}
	label1 := dto.LabelPair{
		Name:  &upMetricInstanceLabel,
		Value: &m.host,
	}
	label2 := dto.LabelPair{
		Name:  &upMetricJobLabel,
		Value: &upMetricJobValue,
	}
	metric := dto.Metric{
		Gauge: &gauge,
		Label: []*dto.LabelPair{&label1, &label2},
	}
	return &dto.MetricFamily{
		Name:   &upMetricName,
		Type:   &upMetricType,
		Metric: []*dto.Metric{&metric},
	}
}

func (m *MetricSet) skipFamily(family *dto.MetricFamily) bool {
	if family == nil {
		return false
	}
	return m.skipFamilyName(*family.Name)
}

func (m *MetricSet) skipFamilyName(family string) bool {
	// example:
	//	include_metrics:
	//		- node_*
	//	exclude_metrics:
	//		- node_disk_*
	//
	// This would mean that we want to keep only the metrics that start with node_ prefix but
	// are not related to disk so we exclude node_disk_* metrics from them.

	// if include_metrics are defined, check if this metric should be included
	if len(m.includeMetrics) > 0 {
		if !matchMetricFamily(family, m.includeMetrics) {
			return true
		}
	}
	// now exclude the metric if it matches any of the given patterns
	if len(m.excludeMetrics) > 0 {
		if matchMetricFamily(family, m.excludeMetrics) {
			return true
		}
	}
	return false
}

func compilePatternList(patterns *[]string) ([]*regexp.Regexp, error) {
	var compiledPatterns []*regexp.Regexp
	compiledPatterns = []*regexp.Regexp{}
	if patterns != nil {
		for _, pattern := range *patterns {
			r, err := regexp.Compile(pattern)
			if err != nil {
				return nil, errors.Wrapf(err, "compiling pattern '%s'", pattern)
			}
			compiledPatterns = append(compiledPatterns, r)
		}
		return compiledPatterns, nil
	}
	return []*regexp.Regexp{}, nil
}

func matchMetricFamily(family string, matchMetrics []*regexp.Regexp) bool {
	for _, checkMetric := range matchMetrics {
		matched := checkMetric.MatchString(family)
		if matched {
			return true
		}
	}
	return false
}
