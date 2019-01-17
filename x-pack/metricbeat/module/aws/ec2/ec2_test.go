// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// +build !integration

package ec2

import (
	"fmt"
	"testing"
	"time"

	"github.com/elastic/beats/metricbeat/mb"
	"github.com/elastic/beats/x-pack/metricbeat/module/aws"

	awssdk "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/cloudwatchiface"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/ec2iface"
	"github.com/stretchr/testify/assert"

	"github.com/elastic/beats/libbeat/common"
)

// MockEC2Client struct is used for unit tests.
type MockEC2Client struct {
	ec2iface.EC2API
}

// MockCloudWatchClient struct is used for unit tests.
type MockCloudWatchClient struct {
	cloudwatchiface.CloudWatchAPI
}

var regionName = "us-west-1"

func (m *MockEC2Client) DescribeRegionsRequest(input *ec2.DescribeRegionsInput) ec2.DescribeRegionsRequest {
	return ec2.DescribeRegionsRequest{
		Request: &awssdk.Request{
			Data: &ec2.DescribeRegionsOutput{
				Regions: []ec2.Region{
					{
						RegionName: &regionName,
					},
				},
			},
		},
	}
}

func (m *MockEC2Client) DescribeInstancesRequest(input *ec2.DescribeInstancesInput) ec2.DescribeInstancesRequest {
	monitoringState := ec2.Monitoring{
		State: ec2.MonitoringState(ec2.MonitoringStateDisabled),
	}
	currentTime := time.Now()
	launchTime := currentTime.Add(-time.Minute * 60)

	instance := ec2.Instance{
		InstanceId:   awssdk.String("i-123"),
		InstanceType: ec2.InstanceTypeT2Medium,
		Placement: &ec2.Placement{
			AvailabilityZone: awssdk.String("us-west-1a"),
		},
		ImageId:    awssdk.String("image-123"),
		Monitoring: &monitoringState,
		State: &ec2.InstanceState{
			Name: ec2.InstanceStateNameRunning,
		},
		LaunchTime: &launchTime,
	}
	return ec2.DescribeInstancesRequest{
		Request: &awssdk.Request{
			Data: &ec2.DescribeInstancesOutput{
				Reservations: []ec2.RunInstancesOutput{
					{Instances: []ec2.Instance{instance}},
				},
			},
		},
	}
}

func (m *MockCloudWatchClient) GetMetricDataRequest(input *cloudwatch.GetMetricDataInput) cloudwatch.GetMetricDataRequest {
	id := "cpu1"
	label := "CPUUtilization"
	value := 0.25
	return cloudwatch.GetMetricDataRequest{
		Request: &awssdk.Request{
			Data: &cloudwatch.GetMetricDataOutput{
				MetricDataResults: []cloudwatch.MetricDataResult{
					{
						Id:     &id,
						Label:  &label,
						Values: []float64{value},
					},
				},
			},
		},
	}
}

func TestGetRegions(t *testing.T) {
	mockSvc := &MockEC2Client{}
	regionsList, err := getRegions(mockSvc)
	if err != nil {
		fmt.Println("failed getRegions: ", err)
		t.FailNow()
	}
	assert.Equal(t, 1, len(regionsList))
	assert.Equal(t, regionName, regionsList[0])
}

func TestGetInstanceIDs(t *testing.T) {
	mockSvc := &MockEC2Client{}
	instanceIDs, instancesOutputs, err := getInstancesPerRegion(mockSvc)
	if err != nil {
		fmt.Println("failed getInstancesPerRegion: ", err)
		t.FailNow()
	}

	assert.Equal(t, 1, len(instanceIDs))
	assert.Equal(t, 1, len(instancesOutputs))

	assert.Equal(t, "i-123", instanceIDs[0])
	assert.Equal(t, ec2.InstanceType("t2.medium"), instancesOutputs["i-123"].InstanceType)
	assert.Equal(t, awssdk.String("image-123"), instancesOutputs["i-123"].ImageId)
	assert.Equal(t, awssdk.String("us-west-1a"), instancesOutputs["i-123"].Placement.AvailabilityZone)
}

func TestGetMetricDataPerRegion(t *testing.T) {
	mockSvc := &MockCloudWatchClient{}
	getMetricDataOutput, err := getMetricDataPerRegion("-10m", 300, "i-123", nil, mockSvc)
	if err != nil {
		fmt.Println("failed getMetricDataPerRegion: ", err)
		t.FailNow()
	}
	assert.Equal(t, 1, len(getMetricDataOutput.MetricDataResults))
	assert.Equal(t, "cpu1", *getMetricDataOutput.MetricDataResults[0].Id)
	assert.Equal(t, "CPUUtilization", *getMetricDataOutput.MetricDataResults[0].Label)
	assert.Equal(t, 0.25, getMetricDataOutput.MetricDataResults[0].Values[0])
}

func TestConvertPeriodToDurationWithDetailedMonitoring(t *testing.T) {
	period1 := "300s"
	duration1, periodSec1 := convertPeriodToDuration(period1, "enabled")
	assert.Equal(t, "-600s", duration1)
	assert.Equal(t, 300, periodSec1)

	period2 := "30ss"
	duration2, periodSec2 := convertPeriodToDuration(period2, "enabled")
	assert.Equal(t, "-120s", duration2)
	assert.Equal(t, 60, periodSec2)

	period3 := "10m"
	duration3, periodSec3 := convertPeriodToDuration(period3, "enabled")
	assert.Equal(t, "-20m", duration3)
	assert.Equal(t, 600, periodSec3)

	period4 := "30s"
	duration4, periodSec4 := convertPeriodToDuration(period4, "enabled")
	assert.Equal(t, "-120s", duration4)
	assert.Equal(t, 60, periodSec4)

	period5 := "60s"
	duration5, periodSec5 := convertPeriodToDuration(period5, "enabled")
	assert.Equal(t, "-120s", duration5)
	assert.Equal(t, 60, periodSec5)
}

func TestConvertPeriodToDurationWithBasicMonitoring(t *testing.T) {
	period1 := "300s"
	duration1, periodSec1 := convertPeriodToDuration(period1, "disabled")
	assert.Equal(t, "-600s", duration1)
	assert.Equal(t, 300, periodSec1)

	period2 := "30ss"
	duration2, periodSec2 := convertPeriodToDuration(period2, "disabled")
	assert.Equal(t, "-600s", duration2)
	assert.Equal(t, 300, periodSec2)

	period3 := "10m"
	duration3, periodSec3 := convertPeriodToDuration(period3, "disabled")
	assert.Equal(t, "-20m", duration3)
	assert.Equal(t, 600, periodSec3)

	period5 := "60s"
	duration5, periodSec5 := convertPeriodToDuration(period5, "disabled")
	assert.Equal(t, "-600s", duration5)
	assert.Equal(t, 300, periodSec5)
}

func TestCreateCloudWatchEvents(t *testing.T) {
	mockModuleConfig := aws.Config{
		Period:        "300s",
		DefaultRegion: regionName,
	}

	expectedEvent := mb.Event{
		RootFields: common.MapStr{
			"service": common.MapStr{"name": "ec2"},
			"cloud": common.MapStr{
				"image":             common.MapStr{"id": "image-123"},
				"region":            regionName,
				"provider":          "ec2",
				"instance":          common.MapStr{"id": "i-123"},
				"machine":           common.MapStr{"type": "t2.medium"},
				"availability_zone": "us-west-1a",
			},
		},
		MetricSetFields: common.MapStr{
			"cpu": common.MapStr{
				"total": common.MapStr{"pct": 0.25},
			},
		},
	}
	svcEC2Mock := &MockEC2Client{}
	instanceIDs, instancesOutputs, err := getInstancesPerRegion(svcEC2Mock)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(instanceIDs))
	instanceID := instanceIDs[0]
	assert.Equal(t, "i-123", instanceID)

	svcCloudwatchMock := &MockCloudWatchClient{}
	detailedMonitoring := instancesOutputs[instanceID].Monitoring.State
	//Calculate duration based on period
	durationString, periodSec := convertPeriodToDuration(mockModuleConfig.Period, detailedMonitoring)
	assert.Equal(t, "-600s", durationString)
	assert.Equal(t, 300, periodSec)

	getMetricDataOutput, err := getMetricDataPerRegion(durationString, periodSec, instanceID, nil, svcCloudwatchMock)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(getMetricDataOutput.MetricDataResults))
	assert.Equal(t, "cpu1", *getMetricDataOutput.MetricDataResults[0].Id)
	assert.Equal(t, "CPUUtilization", *getMetricDataOutput.MetricDataResults[0].Label)
	assert.Equal(t, 0.25, getMetricDataOutput.MetricDataResults[0].Values[0])

	event, err := createCloudWatchEvents(getMetricDataOutput, instanceID, instancesOutputs[instanceID], mockModuleConfig.DefaultRegion)
	assert.NoError(t, err)
	assert.Equal(t, expectedEvent.RootFields, event.RootFields)
	assert.Equal(t, expectedEvent.MetricSetFields["cpu"], event.MetricSetFields["cpu"])
}

func TestCheckInstanceStatus(t *testing.T) {
	currentTime := time.Now()
	launchTime1 := currentTime.Add(-time.Minute * 60)
	checkPassed1 := checkInstanceStatus(ec2.InstanceStateNameRunning, &launchTime1, "i-123")
	assert.Equal(t, true, checkPassed1)

	launchTime2 := currentTime.Add(-time.Minute * 9)
	checkPassed2 := checkInstanceStatus(ec2.InstanceStateNameRunning, &launchTime2, "i-123")
	assert.Equal(t, false, checkPassed2)
}
