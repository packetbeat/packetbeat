// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for CreateVpcEndpoint.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateVpcEndpointRequest
type CreateVpcEndpointInput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier you provide to ensure the idempotency of
	// the request. For more information, see How to Ensure Idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	ClientToken *string `type:"string"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// A policy to attach to the endpoint that controls access to the service. The
	// policy must be in valid JSON format. If this parameter is not specified,
	// we attach a default policy that allows full access to the service.
	PolicyDocument *string `type:"string"`

	// (Interface endpoint) Indicate whether to associate a private hosted zone
	// with the specified VPC. The private hosted zone contains a record set for
	// the default public DNS name for the service for the Region (for example,
	// kinesis.us-east-1.amazonaws.com) which resolves to the private IP addresses
	// of the endpoint network interfaces in the VPC. This enables you to make requests
	// to the default public DNS name for the service instead of the public DNS
	// names that are automatically generated by the VPC endpoint service.
	//
	// To use a private hosted zone, you must set the following VPC attributes to
	// true: enableDnsHostnames and enableDnsSupport. Use ModifyVpcAttribute to
	// set the VPC attributes.
	//
	// Default: true
	PrivateDnsEnabled *bool `type:"boolean"`

	// (Gateway endpoint) One or more route table IDs.
	RouteTableIds []string `locationName:"RouteTableId" locationNameList:"item" type:"list"`

	// (Interface endpoint) The ID of one or more security groups to associate with
	// the endpoint network interface.
	SecurityGroupIds []string `locationName:"SecurityGroupId" locationNameList:"item" type:"list"`

	// The service name. To get a list of available services, use the DescribeVpcEndpointServices
	// request, or get the name from the service provider.
	//
	// ServiceName is a required field
	ServiceName *string `type:"string" required:"true"`

	// (Interface endpoint) The ID of one or more subnets in which to create an
	// endpoint network interface.
	SubnetIds []string `locationName:"SubnetId" locationNameList:"item" type:"list"`

	// The type of endpoint.
	//
	// Default: Gateway
	VpcEndpointType VpcEndpointType `type:"string" enum:"true"`

	// The ID of the VPC in which the endpoint will be used.
	//
	// VpcId is a required field
	VpcId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateVpcEndpointInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateVpcEndpointInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateVpcEndpointInput"}

	if s.ServiceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServiceName"))
	}

	if s.VpcId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpcId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of CreateVpcEndpoint.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateVpcEndpointResult
type CreateVpcEndpointOutput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier you provide to ensure the idempotency of
	// the request.
	ClientToken *string `locationName:"clientToken" type:"string"`

	// Information about the endpoint.
	VpcEndpoint *VpcEndpoint `locationName:"vpcEndpoint" type:"structure"`
}

// String returns the string representation
func (s CreateVpcEndpointOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateVpcEndpoint = "CreateVpcEndpoint"

// CreateVpcEndpointRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates a VPC endpoint for a specified service. An endpoint enables you to
// create a private connection between your VPC and the service. The service
// may be provided by AWS, an AWS Marketplace partner, or another AWS account.
// For more information, see VPC Endpoints (https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/vpc-endpoints.html)
// in the Amazon Virtual Private Cloud User Guide.
//
// A gateway endpoint serves as a target for a route in your route table for
// traffic destined for the AWS service. You can specify an endpoint policy
// to attach to the endpoint that will control access to the service from your
// VPC. You can also specify the VPC route tables that use the endpoint.
//
// An interface endpoint is a network interface in your subnet that serves as
// an endpoint for communicating with the specified service. You can specify
// the subnets in which to create an endpoint, and the security groups to associate
// with the endpoint network interface.
//
// Use DescribeVpcEndpointServices to get a list of supported services.
//
//    // Example sending a request using CreateVpcEndpointRequest.
//    req := client.CreateVpcEndpointRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateVpcEndpoint
func (c *Client) CreateVpcEndpointRequest(input *CreateVpcEndpointInput) CreateVpcEndpointRequest {
	op := &aws.Operation{
		Name:       opCreateVpcEndpoint,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateVpcEndpointInput{}
	}

	req := c.newRequest(op, input, &CreateVpcEndpointOutput{})
	return CreateVpcEndpointRequest{Request: req, Input: input, Copy: c.CreateVpcEndpointRequest}
}

// CreateVpcEndpointRequest is the request type for the
// CreateVpcEndpoint API operation.
type CreateVpcEndpointRequest struct {
	*aws.Request
	Input *CreateVpcEndpointInput
	Copy  func(*CreateVpcEndpointInput) CreateVpcEndpointRequest
}

// Send marshals and sends the CreateVpcEndpoint API request.
func (r CreateVpcEndpointRequest) Send(ctx context.Context) (*CreateVpcEndpointResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateVpcEndpointResponse{
		CreateVpcEndpointOutput: r.Request.Data.(*CreateVpcEndpointOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateVpcEndpointResponse is the response type for the
// CreateVpcEndpoint API operation.
type CreateVpcEndpointResponse struct {
	*CreateVpcEndpointOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateVpcEndpoint request.
func (r *CreateVpcEndpointResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
