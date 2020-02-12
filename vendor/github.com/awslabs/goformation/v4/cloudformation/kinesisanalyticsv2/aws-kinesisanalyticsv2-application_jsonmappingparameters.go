package kinesisanalyticsv2

import (
	"github.com/awslabs/goformation/v4/cloudformation/policies"
)

// Application_JSONMappingParameters AWS CloudFormation Resource (AWS::KinesisAnalyticsV2::Application.JSONMappingParameters)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-jsonmappingparameters.html
type Application_JSONMappingParameters struct {

	// RecordRowPath AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-jsonmappingparameters.html#cfn-kinesisanalyticsv2-application-jsonmappingparameters-recordrowpath
	RecordRowPath string `json:"RecordRowPath,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *Application_JSONMappingParameters) AWSCloudFormationType() string {
	return "AWS::KinesisAnalyticsV2::Application.JSONMappingParameters"
}
