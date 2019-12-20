// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeFleetInstancesRequest
type DescribeFleetInstancesInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The filters.
	//
	//    * instance-type - The instance type.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The ID of the EC2 Fleet.
	//
	// FleetId is a required field
	FleetId *string `type:"string" required:"true"`

	// The maximum number of results to return in a single call. Specify a value
	// between 1 and 1000. The default value is 1000. To retrieve the remaining
	// results, make another call with the returned NextToken value.
	MaxResults *int64 `type:"integer"`

	// The token for the next set of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeFleetInstancesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeFleetInstancesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeFleetInstancesInput"}

	if s.FleetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FleetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeFleetInstancesResult
type DescribeFleetInstancesOutput struct {
	_ struct{} `type:"structure"`

	// The running instances. This list is refreshed periodically and might be out
	// of date.
	ActiveInstances []ActiveInstance `locationName:"activeInstanceSet" locationNameList:"item" type:"list"`

	// The ID of the EC2 Fleet.
	FleetId *string `locationName:"fleetId" type:"string"`

	// The token for the next set of results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeFleetInstancesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeFleetInstances = "DescribeFleetInstances"

// DescribeFleetInstancesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the running instances for the specified EC2 Fleet.
//
//    // Example sending a request using DescribeFleetInstancesRequest.
//    req := client.DescribeFleetInstancesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeFleetInstances
func (c *Client) DescribeFleetInstancesRequest(input *DescribeFleetInstancesInput) DescribeFleetInstancesRequest {
	op := &aws.Operation{
		Name:       opDescribeFleetInstances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeFleetInstancesInput{}
	}

	req := c.newRequest(op, input, &DescribeFleetInstancesOutput{})
	return DescribeFleetInstancesRequest{Request: req, Input: input, Copy: c.DescribeFleetInstancesRequest}
}

// DescribeFleetInstancesRequest is the request type for the
// DescribeFleetInstances API operation.
type DescribeFleetInstancesRequest struct {
	*aws.Request
	Input *DescribeFleetInstancesInput
	Copy  func(*DescribeFleetInstancesInput) DescribeFleetInstancesRequest
}

// Send marshals and sends the DescribeFleetInstances API request.
func (r DescribeFleetInstancesRequest) Send(ctx context.Context) (*DescribeFleetInstancesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeFleetInstancesResponse{
		DescribeFleetInstancesOutput: r.Request.Data.(*DescribeFleetInstancesOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeFleetInstancesResponse is the response type for the
// DescribeFleetInstances API operation.
type DescribeFleetInstancesResponse struct {
	*DescribeFleetInstancesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeFleetInstances request.
func (r *DescribeFleetInstancesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
