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

package query

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/elastic-agent-libs/mapstr"
)

// Response stores the very basic response information to only keep the Status and the ResultType.
type Response struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
	} `json:"data"`
}

// ArrayResponse is for "scalar", "string" type.
// example: {"status":"success","data":{"resultType":"string","result":[1584628642.569,"100"]}}
type ArrayResponse struct {
	Status string    `json:"status"`
	Data   arrayData `json:"data"`
}
type arrayData struct {
	Results []interface{} `json:"result"`
}

// InstantVectorResponse is for "vector" type from Prometheus Query API Request
// instantVectorResult format:
// [
//
//	{
//	  "metric": { "<label_name>": "<label_value>", ... },
//	  "value": [ <unix_time>, "<sample_value>" ]
//	},
//	...
//
// ]
type InstantVectorResponse struct {
	Status string            `json:"status"`
	Data   instantVectorData `json:"data"`
}
type instantVectorData struct {
	Results []instantVectorResult `json:"result"`
}
type instantVectorResult struct {
	Metric map[string]string `json:"metric"`
	Vector []interface{}     `json:"value"`
}

// RangeVectorResponse is for "vector" type from Prometheus Query API Request
// rangeVectorResult format:
// [
//
//	{
//	  "metric": { "<label_name>": "<label_value>", ... },
//	  "values": [ [ <unix_time>, "<sample_value>" ], ... ]
//	},
//	...
//
// ]
type RangeVectorResponse struct {
	Status string          `json:"status"`
	Data   rangeVectorData `json:"data"`
}
type rangeVectorData struct {
	Results []rangeVectorResult `json:"result"`
}
type rangeVectorResult struct {
	Metric  map[string]string `json:"metric"`
	Vectors [][]interface{}   `json:"values"`
}

func parseResponse(body []byte, pathConfig QueryConfig) ([]mb.Event, error) {
	var events []mb.Event

	resultType, err := getResultType(body)
	if err != nil {
		return events, err
	}

	switch resultType {
	case "scalar", "string":
		event, err := getEventFromScalarOrString(body, resultType, pathConfig.Name)
		if err != nil {
			return events, err
		}
		events = append(events, event)
	case "vector":
		evnts, err := getEventsFromVector(body, pathConfig.Name)
		if err != nil {
			return events, err
		}
		events = append(events, evnts...)
	case "matrix":
		evnts, err := getEventsFromMatrix(body, pathConfig.Name)
		if err != nil {
			return events, err
		}
		events = append(events, evnts...)
	default:
		return events, fmt.Errorf("Unknown resultType '%v'", resultType)
	}
	return events, nil
}

func getEventsFromMatrix(body []byte, queryName string) ([]mb.Event, error) {
	events := []mb.Event{}
	convertedMap, err := convertJSONToRangeVectorResponse(body)
	if err != nil {
		return events, err
	}
	results := convertedMap.Data.Results
	for _, result := range results {
		for _, vector := range result.Vectors {
			if vector != nil {
				timestamp, err := getTimestampFromVector(vector)
				if err != nil {
					return []mb.Event{}, err
				}
				val, err := getValueFromVector(vector)
				if err != nil {
					return []mb.Event{}, err
				}
				if math.IsNaN(val) || math.IsInf(val, 0) {
					continue
				}
				events = append(events, mb.Event{
					Timestamp:    getTimestamp(timestamp),
					ModuleFields: mapstr.M{"labels": result.Metric},
					MetricSetFields: mapstr.M{
						queryName: val,
					},
				})
			} else {
				return []mb.Event{}, fmt.Errorf("Could not parse results")
			}
		}
	}
	return events, nil
}

func getEventsFromVector(body []byte, queryName string) ([]mb.Event, error) {
	events := []mb.Event{}
	convertedMap, err := convertJSONToInstantVectorResponse(body)
	if err != nil {
		return events, err
	}
	results := convertedMap.Data.Results
	for _, result := range results {
		if result.Vector != nil {
			timestamp, err := getTimestampFromVector(result.Vector)
			if err != nil {
				return []mb.Event{}, err
			}
			val, err := getValueFromVector(result.Vector)
			if err != nil {
				return []mb.Event{}, err
			}
			if math.IsNaN(val) || math.IsInf(val, 0) {
				continue
			}
			events = append(events, mb.Event{
				Timestamp:    getTimestamp(timestamp),
				ModuleFields: mapstr.M{"labels": result.Metric},
				MetricSetFields: mapstr.M{
					queryName: val,
				},
			})
		} else {
			return []mb.Event{}, fmt.Errorf("Could not parse results")
		}
	}
	return events, nil
}

func getEventFromScalarOrString(body []byte, resultType string, queryName string) (mb.Event, error) {
	convertedArray, err := convertJSONToArrayResponse(body)
	if err != nil {
		return mb.Event{}, err
	}
	if convertedArray.Data.Results != nil {
		timestamp, err := getTimestampFromVector(convertedArray.Data.Results)
		if err != nil {
			return mb.Event{}, err
		}
		if resultType == "scalar" {
			val, err := getValueFromVector(convertedArray.Data.Results)
			if err != nil {
				return mb.Event{}, err
			}
			if math.IsNaN(val) || math.IsInf(val, 0) {
				return mb.Event{}, nil
			}
			return mb.Event{
				Timestamp: getTimestamp(timestamp),
				MetricSetFields: mapstr.M{
					queryName: val,
				},
			}, nil
		} else if resultType == "string" {
			value, ok := convertedArray.Data.Results[1].(string)
			if !ok {
				return mb.Event{}, fmt.Errorf("Could not parse value of result: %v", convertedArray.Data.Results)
			}
			return mb.Event{
				Timestamp: getTimestamp(timestamp),
				ModuleFields: mapstr.M{
					"labels": mapstr.M{
						queryName: value,
					},
				},
				MetricSetFields: mapstr.M{
					queryName: 1,
				},
			}, nil
		}
	}
	return mb.Event{}, fmt.Errorf("could not parse results")
}

func getTimestampFromVector(vector []interface{}) (float64, error) {
	// Example input: [ <unix_time>, "<sample_value>" ]
	if len(vector) != 2 {
		return 0, fmt.Errorf("could not parse results")
	}
	timestamp, ok := vector[0].(float64)
	if !ok {
		return 0, fmt.Errorf("Could not parse timestamp of result: %v", vector)
	}
	return timestamp, nil
}

func getValueFromVector(vector []interface{}) (float64, error) {
	// Example input: [ <unix_time>, "<sample_value>" ]
	if len(vector) != 2 {
		return 0, errors.New("could not parse results")
	}
	value, ok := vector[1].(string)
	if !ok {
		return 0, fmt.Errorf("Could not parse value of result: %v", vector)
	}
	val, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, fmt.Errorf("Could not parse value of result: %v", vector)
	}
	return val, nil
}

func getResultType(body []byte) (string, error) {
	response := Response{}
	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("failed to parse api response: %w", err)
	}
	if response.Status == "error" {
		return "", fmt.Errorf("failed to query")
	}
	return response.Data.ResultType, nil
}

func convertJSONToArrayResponse(body []byte) (ArrayResponse, error) {
	arrayBody := ArrayResponse{}
	if err := json.Unmarshal(body, &arrayBody); err != nil {
		return arrayBody, fmt.Errorf("failed to parse api response: %w", err)
	}
	if arrayBody.Status == "error" {
		return arrayBody, fmt.Errorf("failed to query")
	}
	return arrayBody, nil
}

func convertJSONToRangeVectorResponse(body []byte) (RangeVectorResponse, error) {
	mapBody := RangeVectorResponse{}
	if err := json.Unmarshal(body, &mapBody); err != nil {
		return RangeVectorResponse{}, fmt.Errorf("failed to parse api response: %w", err)
	}
	if mapBody.Status == "error" {
		return mapBody, fmt.Errorf("failed to query")
	}
	return mapBody, nil
}

func convertJSONToInstantVectorResponse(body []byte) (InstantVectorResponse, error) {
	mapBody := InstantVectorResponse{}
	if err := json.Unmarshal(body, &mapBody); err != nil {
		return InstantVectorResponse{}, fmt.Errorf("failed to parse api response: %w", err)
	}
	if mapBody.Status == "error" {
		return mapBody, fmt.Errorf("failed to query")
	}
	return mapBody, nil
}

func getTimestamp(num float64) time.Time {
	sec := int64(num)
	ns := int64((num - float64(sec)) * 1000)
	return time.Unix(sec, ns)
}
