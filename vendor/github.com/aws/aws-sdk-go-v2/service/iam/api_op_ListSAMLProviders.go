// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/ListSAMLProvidersRequest
type ListSAMLProvidersInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ListSAMLProvidersInput) String() string {
	return awsutil.Prettify(s)
}

// Contains the response to a successful ListSAMLProviders request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/ListSAMLProvidersResponse
type ListSAMLProvidersOutput struct {
	_ struct{} `type:"structure"`

	// The list of SAML provider resource objects defined in IAM for this AWS account.
	SAMLProviderList []SAMLProviderListEntry `type:"list"`
}

// String returns the string representation
func (s ListSAMLProvidersOutput) String() string {
	return awsutil.Prettify(s)
}

const opListSAMLProviders = "ListSAMLProviders"

// ListSAMLProvidersRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Lists the SAML provider resource objects defined in IAM in the account.
//
// This operation requires Signature Version 4 (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html).
//
//    // Example sending a request using ListSAMLProvidersRequest.
//    req := client.ListSAMLProvidersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/ListSAMLProviders
func (c *Client) ListSAMLProvidersRequest(input *ListSAMLProvidersInput) ListSAMLProvidersRequest {
	op := &aws.Operation{
		Name:       opListSAMLProviders,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListSAMLProvidersInput{}
	}

	req := c.newRequest(op, input, &ListSAMLProvidersOutput{})
	return ListSAMLProvidersRequest{Request: req, Input: input, Copy: c.ListSAMLProvidersRequest}
}

// ListSAMLProvidersRequest is the request type for the
// ListSAMLProviders API operation.
type ListSAMLProvidersRequest struct {
	*aws.Request
	Input *ListSAMLProvidersInput
	Copy  func(*ListSAMLProvidersInput) ListSAMLProvidersRequest
}

// Send marshals and sends the ListSAMLProviders API request.
func (r ListSAMLProvidersRequest) Send(ctx context.Context) (*ListSAMLProvidersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListSAMLProvidersResponse{
		ListSAMLProvidersOutput: r.Request.Data.(*ListSAMLProvidersOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListSAMLProvidersResponse is the response type for the
// ListSAMLProviders API operation.
type ListSAMLProvidersResponse struct {
	*ListSAMLProvidersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListSAMLProviders request.
func (r *ListSAMLProvidersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
