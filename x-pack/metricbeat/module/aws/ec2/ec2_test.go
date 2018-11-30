// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package ec2

import (
	"fmt"
	"os"
	"testing"

	"github.com/elastic/beats/x-pack/metricbeat/module/aws"

	"github.com/elastic/beats/libbeat/common"

	awssdk "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/stretchr/testify/assert"

	"github.com/elastic/beats/metricbeat/mb"
	mbtest "github.com/elastic/beats/metricbeat/mb/testing"
)

var regionName = "us-west-1"

func (m *MockEC2Client) DescribeRegionsRequest(input *ec2.DescribeRegionsInput) ec2.DescribeRegionsRequest {
	return ec2.DescribeRegionsRequest{
		Request: &awssdk.Request{
			Data: &ec2.DescribeRegionsOutput{
				Regions: []ec2.Region{
					ec2.Region{
						RegionName: &regionName,
					},
				},
			},
		},
	}
}

func (m *MockEC2Client) DescribeInstancesRequest(input *ec2.DescribeInstancesInput) ec2.DescribeInstancesRequest {
	instance1 := ec2.Instance{
		InstanceId:   awssdk.String("i-123"),
		InstanceType: ec2.InstanceTypeT2Medium,
		Placement:    &ec2.Placement{AvailabilityZone: awssdk.String("us-west-1a")},
		ImageId:      awssdk.String("image-123"),
	}
	instance2 := ec2.Instance{
		InstanceId:   awssdk.String("i-456"),
		InstanceType: ec2.InstanceTypeT2Micro,
		Placement:    &ec2.Placement{AvailabilityZone: awssdk.String("us-west-1b")},
		ImageId:      awssdk.String("image-456"),
	}
	return ec2.DescribeInstancesRequest{
		Request: &awssdk.Request{
			Data: &ec2.DescribeInstancesOutput{
				Reservations: []ec2.RunInstancesOutput{
					ec2.RunInstancesOutput{Instances: []ec2.Instance{instance1, instance2}},
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
					cloudwatch.MetricDataResult{
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

	assert.Equal(t, 2, len(instanceIDs))
	assert.Equal(t, 2, len(instancesOutputs))

	assert.Equal(t, "i-123", instanceIDs[0])
	assert.Equal(t, ec2.InstanceType("t2.medium"), instancesOutputs["i-123"].InstanceType)
	assert.Equal(t, awssdk.String("image-123"), instancesOutputs["i-123"].ImageId)
	assert.Equal(t, awssdk.String("us-west-1a"), instancesOutputs["i-123"].Placement.AvailabilityZone)

	assert.Equal(t, "i-456", instanceIDs[1])
	assert.Equal(t, ec2.InstanceType("t2.micro"), instancesOutputs["i-456"].InstanceType)
	assert.Equal(t, awssdk.String("us-west-1b"), instancesOutputs["i-456"].Placement.AvailabilityZone)
}

func TestGetMetricDataPerRegion(t *testing.T) {
	mockSvc := &MockCloudWatchClient{}
	getMetricDataOutput, err := getMetricDataPerRegion("i-123", nil, mockSvc)
	if err != nil {
		fmt.Println("failed getMetricDataPerRegion: ", err)
		t.FailNow()
	}
	assert.Equal(t, 1, len(getMetricDataOutput.MetricDataResults))
	assert.Equal(t, "cpu1", *getMetricDataOutput.MetricDataResults[0].Id)
	assert.Equal(t, "CPUUtilization", *getMetricDataOutput.MetricDataResults[0].Label)
	assert.Equal(t, 0.25, getMetricDataOutput.MetricDataResults[0].Values[0])
}

func TestMockFetch(t *testing.T) {
	mockCreds := map[string]interface{}{
		"module":     "aws",
		"metricsets": []string{"ec2"},
		"mock":       "true",
	}

	awsMetricSet := mbtest.NewReportingMetricSetV2(t, mockCreds)
	events, errs := mbtest.ReportingFetchV2(awsMetricSet)
	assert.Empty(t, errs)
	if !assert.NotEmpty(t, events) {
		t.FailNow()
	}
	t.Logf("Module: %s Metricset: %s", awsMetricSet.Module().Name(), awsMetricSet.Name())

	assert.Equal(t, 2, len(events))
	for _, event := range events {
		// RootField
		checkRootField("service", event, t)
		// ModuleField
		checkModuleField("cloud.availability_zone", event, t)
		checkModuleField("cloud.provider", event, t)
		checkModuleField("cloud.image.id", event, t)
		checkModuleField("cloud.instance.id", event, t)
		checkModuleField("cloud.machine.type", event, t)
		checkModuleField("cloud.provider", event, t)
		checkModuleField("cloud.region", event, t)
		// MetricSetField
		cpuTotalPct, err := event.MetricSetFields.GetValue("ec2.cpu.total.pct")
		assert.NoError(t, err)
		assert.Equal(t, 0.25, cpuTotalPct)
	}
}

func TestFetch(t *testing.T) {
	accessKeyID, okAccessKeyID := os.LookupEnv("AWS_ACCESS_KEY_ID")
	secretAccessKey, okSecretAccessKey := os.LookupEnv("AWS_SECRET_ACCESS_KEY")
	sessionToken, okSessionToken := os.LookupEnv("AWS_SESSION_TOKEN")
	if !okAccessKeyID || accessKeyID == "" {
		t.Skip("Skipping TestFetch; $AWS_ACCESS_KEY_ID not set or set to empty")
	} else if !okSecretAccessKey || secretAccessKey == "" {
		t.Skip("Skipping TestFetch; $AWS_SECRET_ACCESS_KEY not set or set to empty")
	} else {
		tempCreds := map[string]interface{}{
			"module":     "aws",
			"metricsets": []string{"ec2"},
			"mock":       "false",
		}
		tempCreds["access_key_id"] = accessKeyID
		tempCreds["secret_access_key"] = secretAccessKey
		if okSessionToken && sessionToken != "" {
			tempCreds["session_token"] = sessionToken
		}

		awsMetricSet := mbtest.NewReportingMetricSetV2(t, tempCreds)
		events, errs := mbtest.ReportingFetchV2(awsMetricSet)
		assert.Empty(t, errs)
		if !assert.NotEmpty(t, events) {
			t.FailNow()
		}
		t.Logf("Module: %s Metricset: %s", awsMetricSet.Module().Name(), awsMetricSet.Name())

		for _, event := range events {
			// RootField
			checkRootField("service", event, t)
			// ModuleField
			checkModuleField("cloud.availability_zone", event, t)
			checkModuleField("cloud.provider", event, t)
			checkModuleField("cloud.image.id", event, t)
			checkModuleField("cloud.instance.id", event, t)
			checkModuleField("cloud.machine.type", event, t)
			checkModuleField("cloud.provider", event, t)
			checkModuleField("cloud.region", event, t)
			// MetricSetField
			checkMetricSetField("ec2.cpu.total.pct", event, t)
			checkMetricSetField("ec2.cpu.credit_usage", event, t)
			checkMetricSetField("ec2.cpu.credit_balance", event, t)
			checkMetricSetField("ec2.cpu.surplus_credit_balance", event, t)
			checkMetricSetField("ec2.cpu.surplus_credits_charged", event, t)
			checkMetricSetField("network.in.packets", event, t)
			checkMetricSetField("network.out.packets", event, t)
			checkMetricSetField("network.in.bytes", event, t)
			checkMetricSetField("network.out.bytes", event, t)
			checkMetricSetField("diskio.read.bytes", event, t)
			checkMetricSetField("diskio.write.bytes", event, t)
			checkMetricSetField("diskio.read.ops", event, t)
			checkMetricSetField("diskio.write.ops", event, t)
			checkMetricSetField("status.check_failed", event, t)
			checkMetricSetField("status.check_failed_system", event, t)
			checkMetricSetField("status.check_failed_instance", event, t)
		}
	}
}

func checkRootField(fieldName string, event mb.Event, t *testing.T) {
	expectedKeyValue := common.MapStr{"name": aws.ModuleName}
	if ok, err := event.RootFields.HasKey(fieldName); ok {
		assert.NoError(t, err)
		fieldValue, err := event.RootFields.GetValue(fieldName)
		assert.NoError(t, err)
		assert.Equal(t, expectedKeyValue, fieldValue)
	}
}

func checkModuleField(fieldName string, event mb.Event, t *testing.T) {
	metricValue, err := event.ModuleFields.GetValue(fieldName)
	assert.NoError(t, err)
	if userString, ok := metricValue.(string); !ok {
		fmt.Println("Field "+fieldName+" is not a string: ", userString)
		t.Fail()
	}
}

func checkMetricSetField(metricName string, event mb.Event, t *testing.T) {
	if ok, err := event.MetricSetFields.HasKey(metricName); ok {
		assert.NoError(t, err)
		metricValue, err := event.MetricSetFields.GetValue(metricName)
		assert.NoError(t, err)
		if userPercentFloat, ok := metricValue.(float64); !ok {
			fmt.Println("failed: userPercentFloat = ", userPercentFloat)
			t.Fail()
		} else {
			assert.True(t, userPercentFloat >= 0)
			fmt.Println("succeed: userPercentFloat = ", userPercentFloat)
		}
	}
}
