// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeAggregateIdFormatRequest
type DescribeAggregateIdFormatInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`
}

// String returns the string representation
func (s DescribeAggregateIdFormatInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeAggregateIdFormatResult
type DescribeAggregateIdFormatOutput struct {
	_ struct{} `type:"structure"`

	// Information about each resource's ID format.
	Statuses []IdFormat `locationName:"statusSet" locationNameList:"item" type:"list"`

	// Indicates whether all resource types in the Region are configured to use
	// longer IDs. This value is only true if all users are configured to use longer
	// IDs for all resources types in the Region.
	UseLongIdsAggregated *bool `locationName:"useLongIdsAggregated" type:"boolean"`
}

// String returns the string representation
func (s DescribeAggregateIdFormatOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeAggregateIdFormat = "DescribeAggregateIdFormat"

// DescribeAggregateIdFormatRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the longer ID format settings for all resource types in a specific
// Region. This request is useful for performing a quick audit to determine
// whether a specific Region is fully opted in for longer IDs (17-character
// IDs).
//
// This request only returns information about resource types that support longer
// IDs.
//
// The following resource types support longer IDs: bundle | conversion-task
// | customer-gateway | dhcp-options | elastic-ip-allocation | elastic-ip-association
// | export-task | flow-log | image | import-task | instance | internet-gateway
// | network-acl | network-acl-association | network-interface | network-interface-attachment
// | prefix-list | reservation | route-table | route-table-association | security-group
// | snapshot | subnet | subnet-cidr-block-association | volume | vpc | vpc-cidr-block-association
// | vpc-endpoint | vpc-peering-connection | vpn-connection | vpn-gateway.
//
//    // Example sending a request using DescribeAggregateIdFormatRequest.
//    req := client.DescribeAggregateIdFormatRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeAggregateIdFormat
func (c *Client) DescribeAggregateIdFormatRequest(input *DescribeAggregateIdFormatInput) DescribeAggregateIdFormatRequest {
	op := &aws.Operation{
		Name:       opDescribeAggregateIdFormat,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeAggregateIdFormatInput{}
	}

	req := c.newRequest(op, input, &DescribeAggregateIdFormatOutput{})
	return DescribeAggregateIdFormatRequest{Request: req, Input: input, Copy: c.DescribeAggregateIdFormatRequest}
}

// DescribeAggregateIdFormatRequest is the request type for the
// DescribeAggregateIdFormat API operation.
type DescribeAggregateIdFormatRequest struct {
	*aws.Request
	Input *DescribeAggregateIdFormatInput
	Copy  func(*DescribeAggregateIdFormatInput) DescribeAggregateIdFormatRequest
}

// Send marshals and sends the DescribeAggregateIdFormat API request.
func (r DescribeAggregateIdFormatRequest) Send(ctx context.Context) (*DescribeAggregateIdFormatResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAggregateIdFormatResponse{
		DescribeAggregateIdFormatOutput: r.Request.Data.(*DescribeAggregateIdFormatOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeAggregateIdFormatResponse is the response type for the
// DescribeAggregateIdFormat API operation.
type DescribeAggregateIdFormatResponse struct {
	*DescribeAggregateIdFormatOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAggregateIdFormat request.
func (r *DescribeAggregateIdFormatResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
