package easyops

import (
	"math"

	"github.com/elastic/beats/v7/libbeat/common"
)

type divMetricBuilder struct {
	baseBuilderFields
}

func newDivMetricBuilder(field string, originMetric []string, groupKeys []string, defaultValues map[string]interface{}) AggregateMetricBuilder {
	return &divMetricBuilder{
		baseBuilderFields{
			field:         field,
			originMetrics: originMetric,
			groupKeys:     groupKeys,
			defaultValues: defaultValues,
		},
	}
}

func (builder *divMetricBuilder) Build(events []common.MapStr) []common.MapStr {
	var result []common.MapStr
	eventMap := GroupEventsByKeys(events, builder.groupKeys)
	for _, es := range eventMap {
		if len(es) == 0 {
			continue
		}
		rs := common.MapStr{}
		for _, groupKey := range builder.groupKeys {
			// GetValue success in GroupEventsByKeys
			val, _ := es[0].GetValue(groupKey)
			_, _ = rs.Put(groupKey, val)
		}
		_, _ = rs.Put(builder.field, builder.div(es, builder.originMetrics))
		result = append(result, rs)
	}
	return result
}

func (builder *divMetricBuilder) div(events []common.MapStr, originMetric []string) interface{} {
	var floatResult float64 = 0
MetricLoop:
	for index, metric := range originMetric {
		metricSum := 0.0
		for _, event := range events {
			value, err := event.GetValue(metric)
			if err == nil {
				metricSum += ConvertNumericValue(value)
			}
		}
		if index == 0 {
			floatResult = metricSum
		} else if metricSum != 0 {
			floatResult /= metricSum
		} else {
			floatResult = math.MaxFloat64
			break MetricLoop
		}
	}
	return floatResult
}
