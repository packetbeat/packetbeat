// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/DeleteAccountPasswordPolicyInput
type DeleteAccountPasswordPolicyInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteAccountPasswordPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/DeleteAccountPasswordPolicyOutput
type DeleteAccountPasswordPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteAccountPasswordPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteAccountPasswordPolicy = "DeleteAccountPasswordPolicy"

// DeleteAccountPasswordPolicyRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Deletes the password policy for the AWS account. There are no parameters.
//
//    // Example sending a request using DeleteAccountPasswordPolicyRequest.
//    req := client.DeleteAccountPasswordPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/DeleteAccountPasswordPolicy
func (c *Client) DeleteAccountPasswordPolicyRequest(input *DeleteAccountPasswordPolicyInput) DeleteAccountPasswordPolicyRequest {
	op := &aws.Operation{
		Name:       opDeleteAccountPasswordPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteAccountPasswordPolicyInput{}
	}

	req := c.newRequest(op, input, &DeleteAccountPasswordPolicyOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteAccountPasswordPolicyRequest{Request: req, Input: input, Copy: c.DeleteAccountPasswordPolicyRequest}
}

// DeleteAccountPasswordPolicyRequest is the request type for the
// DeleteAccountPasswordPolicy API operation.
type DeleteAccountPasswordPolicyRequest struct {
	*aws.Request
	Input *DeleteAccountPasswordPolicyInput
	Copy  func(*DeleteAccountPasswordPolicyInput) DeleteAccountPasswordPolicyRequest
}

// Send marshals and sends the DeleteAccountPasswordPolicy API request.
func (r DeleteAccountPasswordPolicyRequest) Send(ctx context.Context) (*DeleteAccountPasswordPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteAccountPasswordPolicyResponse{
		DeleteAccountPasswordPolicyOutput: r.Request.Data.(*DeleteAccountPasswordPolicyOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteAccountPasswordPolicyResponse is the response type for the
// DeleteAccountPasswordPolicy API operation.
type DeleteAccountPasswordPolicyResponse struct {
	*DeleteAccountPasswordPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteAccountPasswordPolicy request.
func (r *DeleteAccountPasswordPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
