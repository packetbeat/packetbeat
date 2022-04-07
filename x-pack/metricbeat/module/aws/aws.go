// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package aws

import (
	"context"
	"fmt"
	"strings"
	"time"

	awssdk "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/ec2iface"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/iamiface"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi"
	"github.com/aws/aws-sdk-go-v2/service/sts"

	"github.com/elastic/beats/v8/libbeat/common"
	"github.com/elastic/beats/v8/metricbeat/mb"
	awscommon "github.com/elastic/beats/v8/x-pack/libbeat/common/aws"
)

// Config defines all required and optional parameters for aws metricsets
type Config struct {
	Period     time.Duration       `config:"period" validate:"nonzero,required"`
	Regions    []string            `config:"regions"`
	Latency    time.Duration       `config:"latency"`
	AWSConfig  awscommon.ConfigAWS `config:",inline"`
	TagsFilter []Tag               `config:"tags_filter"`
}

// MetricSet is the base metricset for all aws metricsets
type MetricSet struct {
	mb.BaseMetricSet
	RegionsList []string
	Endpoint    string
	Period      time.Duration
	Latency     time.Duration
	AwsConfig   *awssdk.Config
	AccountName string
	AccountID   string
	TagsFilter  []Tag
}

// Tag holds a configuration specific for ec2 and cloudwatch metricset.
type Tag struct {
	Key   string `config:"key"`
	Value string `config:"value"`
}

// ModuleName is the name of this module.
const ModuleName = "aws"

func init() {
	if err := mb.Registry.AddModule(ModuleName, newModule); err != nil {
		panic(err)
	}
}

func newModule(base mb.BaseModule) (mb.Module, error) {
	var config Config
	if err := base.UnpackConfig(&config); err != nil {
		return nil, err
	}
	return &base, nil
}

// NewMetricSet creates a base metricset for aws metricsets
func NewMetricSet(base mb.BaseMetricSet) (*MetricSet, error) {
	var config Config
	err := base.Module().UnpackConfig(&config)
	if err != nil {
		return nil, err
	}

	awsConfig, err := awscommon.InitializeAWSConfig(config.AWSConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to get aws credentials, please check AWS credential in config: %w", err)
	}

	_, err = awsConfig.Credentials.Retrieve(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve aws credentials, please check AWS credential in config: %w", err)
	}

	base.Logger().Debug("aws config endpoint = ", config.AWSConfig.Endpoint)
	metricSet := MetricSet{
		BaseMetricSet: base,
		Period:        config.Period,
		Latency:       config.Latency,
		AwsConfig:     &awsConfig,
		TagsFilter:    config.TagsFilter,
		Endpoint:      config.AWSConfig.Endpoint,
	}

	base.Logger().Debug("Metricset level config for period: ", metricSet.Period)
	base.Logger().Debug("Metricset level config for tags filter: ", metricSet.TagsFilter)
	base.Logger().Warn("extra charges on AWS API requests will be generated by this metricset")

	// If regions in config is not empty, then overwrite the awsConfig.Region
	if len(config.Regions) > 0 {
		awsConfig.Region = config.Regions[0]
	}

	stsServiceName := awscommon.CreateServiceName("sts", config.AWSConfig.FIPSEnabled, awsConfig.Region)
	iamServiceName := awscommon.CreateServiceName("iam", config.AWSConfig.FIPSEnabled, awsConfig.Region)

	// Get IAM account id
	svcSts := sts.New(awscommon.EnrichAWSConfigWithEndpoint(
		config.AWSConfig.Endpoint, stsServiceName, "", awsConfig))
	reqIdentity := svcSts.GetCallerIdentityRequest(&sts.GetCallerIdentityInput{})
	outputIdentity, err := reqIdentity.Send(context.TODO())
	if err != nil {
		base.Logger().Warn("failed to get caller identity, please check permission setting: ", err)
	} else {
		metricSet.AccountID = *outputIdentity.Account
		base.Logger().Debug("AWS Credentials belong to account ID: ", metricSet.AccountID)
	}
	iamRegion := ""
	if strings.HasPrefix(awsConfig.Region, "us-gov-") {
		iamRegion = "us-gov"
	}
	// Get account name/alias
	svcIam := iam.New(awscommon.EnrichAWSConfigWithEndpoint(
		config.AWSConfig.Endpoint, iamServiceName, iamRegion, awsConfig))
	metricSet.AccountName = getAccountName(svcIam, base, metricSet)

	// Construct MetricSet with a full regions list
	if config.Regions == nil {
		ec2ServiceName := awscommon.CreateServiceName("ec2", config.AWSConfig.FIPSEnabled, awsConfig.Region)
		svcEC2 := ec2.New(awscommon.EnrichAWSConfigWithEndpoint(
			config.AWSConfig.Endpoint, ec2ServiceName, "", awsConfig))
		completeRegionsList, err := getRegions(svcEC2)
		if err != nil {
			return nil, err
		}

		metricSet.RegionsList = completeRegionsList
		base.Logger().Debug("Metricset level config for regions: ", metricSet.RegionsList)
		return &metricSet, nil
	}

	// Construct MetricSet with specific regions list from config
	metricSet.RegionsList = config.Regions
	base.Logger().Debug("Metricset level config for regions: ", metricSet.RegionsList)
	return &metricSet, nil
}

func getRegions(svc ec2iface.ClientAPI) (completeRegionsList []string, err error) {
	input := &ec2.DescribeRegionsInput{}
	req := svc.DescribeRegionsRequest(input)
	output, err := req.Send(context.TODO())
	if err != nil {
		err = fmt.Errorf("failed DescribeRegions: %w", err)
		return
	}
	for _, region := range output.Regions {
		completeRegionsList = append(completeRegionsList, *region.RegionName)
	}
	return
}

func getAccountName(svc iamiface.ClientAPI, base mb.BaseMetricSet, metricSet MetricSet) string {
	req := svc.ListAccountAliasesRequest(&iam.ListAccountAliasesInput{})
	output, err := req.Send(context.TODO())

	accountName := metricSet.AccountID
	if err != nil {
		base.Logger().Warn("failed to list account aliases, please check permission setting: ", err)
		return accountName
	}

	// When there is no account alias, account ID will be used as cloud.account.name
	if len(output.AccountAliases) == 0 {
		accountName = metricSet.AccountID
		base.Logger().Debug("AWS Credentials belong to account ID: ", metricSet.AccountID)
		return accountName
	}

	// There can be more than one aliases for each account, for now we are only
	// collecting the first one.
	accountName = output.AccountAliases[0]
	base.Logger().Debug("AWS Credentials belong to account name: ", metricSet.AccountName)
	return accountName
}

// StringInSlice checks if a string is already exists in list and its location
func StringInSlice(str string, list []string) (bool, int) {
	for idx, v := range list {
		if v == str {
			return true, idx
		}
	}
	// If this string doesn't exist in given list, then return location to be -1
	return false, -1
}

// InitEvent initialize mb.Event with basic information like service.name, cloud.provider
func InitEvent(regionName string, accountName string, accountID string, timestamp time.Time) mb.Event {
	event := mb.Event{
		Timestamp:       timestamp,
		MetricSetFields: common.MapStr{},
		ModuleFields:    common.MapStr{},
		RootFields:      common.MapStr{},
	}

	event.RootFields.Put("cloud.provider", "aws")
	if regionName != "" {
		event.RootFields.Put("cloud.region", regionName)
	}
	if accountName != "" {
		event.RootFields.Put("cloud.account.name", accountName)
	}
	if accountID != "" {
		event.RootFields.Put("cloud.account.id", accountID)
	}
	return event
}

// CheckTagFiltersExist compare tags filter with a set of tags to see if tags
// filter is a subset of tags
func CheckTagFiltersExist(tagsFilter []Tag, tags interface{}) bool {
	var tagKeys []string
	var tagValues []string

	switch tags.(type) {
	case []resourcegroupstaggingapi.Tag:
		tagsResource := tags.([]resourcegroupstaggingapi.Tag)
		for _, tag := range tagsResource {
			tagKeys = append(tagKeys, *tag.Key)
			tagValues = append(tagValues, *tag.Value)
		}
	case []ec2.Tag:
		tagsEC2 := tags.([]ec2.Tag)
		for _, tag := range tagsEC2 {
			tagKeys = append(tagKeys, *tag.Key)
			tagValues = append(tagValues, *tag.Value)
		}
	case []rds.Tag:
		tagsRDS := tags.([]rds.Tag)
		for _, tag := range tagsRDS {
			tagKeys = append(tagKeys, *tag.Key)
			tagValues = append(tagValues, *tag.Value)
		}
	}

	for _, tagFilter := range tagsFilter {
		if exists, idx := StringInSlice(tagFilter.Key, tagKeys); !exists || tagValues[idx] != tagFilter.Value {
			return false
		}
	}
	return true
}
