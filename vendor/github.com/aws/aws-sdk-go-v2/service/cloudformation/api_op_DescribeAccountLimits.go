// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The input for the DescribeAccountLimits action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/DescribeAccountLimitsInput
type DescribeAccountLimitsInput struct {
	_ struct{} `type:"structure"`

	// A string that identifies the next page of limits that you want to retrieve.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeAccountLimitsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeAccountLimitsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeAccountLimitsInput"}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The output for the DescribeAccountLimits action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/DescribeAccountLimitsOutput
type DescribeAccountLimitsOutput struct {
	_ struct{} `type:"structure"`

	// An account limit structure that contain a list of AWS CloudFormation account
	// limits and their values.
	AccountLimits []AccountLimit `type:"list"`

	// If the output exceeds 1 MB in size, a string that identifies the next page
	// of limits. If no additional page exists, this value is null.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeAccountLimitsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeAccountLimits = "DescribeAccountLimits"

// DescribeAccountLimitsRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Retrieves your account's AWS CloudFormation limits, such as the maximum number
// of stacks that you can create in your account. For more information about
// account limits, see AWS CloudFormation Limits (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cloudformation-limits.html)
// in the AWS CloudFormation User Guide.
//
//    // Example sending a request using DescribeAccountLimitsRequest.
//    req := client.DescribeAccountLimitsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/DescribeAccountLimits
func (c *Client) DescribeAccountLimitsRequest(input *DescribeAccountLimitsInput) DescribeAccountLimitsRequest {
	op := &aws.Operation{
		Name:       opDescribeAccountLimits,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeAccountLimitsInput{}
	}

	req := c.newRequest(op, input, &DescribeAccountLimitsOutput{})
	return DescribeAccountLimitsRequest{Request: req, Input: input, Copy: c.DescribeAccountLimitsRequest}
}

// DescribeAccountLimitsRequest is the request type for the
// DescribeAccountLimits API operation.
type DescribeAccountLimitsRequest struct {
	*aws.Request
	Input *DescribeAccountLimitsInput
	Copy  func(*DescribeAccountLimitsInput) DescribeAccountLimitsRequest
}

// Send marshals and sends the DescribeAccountLimits API request.
func (r DescribeAccountLimitsRequest) Send(ctx context.Context) (*DescribeAccountLimitsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAccountLimitsResponse{
		DescribeAccountLimitsOutput: r.Request.Data.(*DescribeAccountLimitsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeAccountLimitsResponse is the response type for the
// DescribeAccountLimits API operation.
type DescribeAccountLimitsResponse struct {
	*DescribeAccountLimitsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAccountLimits request.
func (r *DescribeAccountLimitsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
