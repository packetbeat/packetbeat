// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/ListPolicyVersionsRequest
type ListPolicyVersionsInput struct {
	_ struct{} `type:"structure"`

	// Use this parameter only when paginating results and only after you receive
	// a response indicating that the results are truncated. Set it to the value
	// of the Marker element in the response that you received to indicate where
	// the next call should start.
	Marker *string `min:"1" type:"string"`

	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true.
	//
	// If you do not include this parameter, the number of items defaults to 100.
	// Note that IAM might return fewer results, even when there are more results
	// available. In that case, the IsTruncated response element returns true, and
	// Marker contains a value to include in the subsequent call that tells the
	// service where to continue from.
	MaxItems *int64 `min:"1" type:"integer"`

	// The Amazon Resource Name (ARN) of the IAM policy for which you want the versions.
	//
	// For more information about ARNs, see Amazon Resource Names (ARNs) and AWS
	// Service Namespaces (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the AWS General Reference.
	//
	// PolicyArn is a required field
	PolicyArn *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s ListPolicyVersionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListPolicyVersionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListPolicyVersionsInput"}
	if s.Marker != nil && len(*s.Marker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Marker", 1))
	}
	if s.MaxItems != nil && *s.MaxItems < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxItems", 1))
	}

	if s.PolicyArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyArn"))
	}
	if s.PolicyArn != nil && len(*s.PolicyArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("PolicyArn", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a successful ListPolicyVersions request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/ListPolicyVersionsResponse
type ListPolicyVersionsOutput struct {
	_ struct{} `type:"structure"`

	// A flag that indicates whether there are more items to return. If your results
	// were truncated, you can make a subsequent pagination request using the Marker
	// request parameter to retrieve more items. Note that IAM might return fewer
	// than the MaxItems number of results even when there are more results available.
	// We recommend that you check IsTruncated after every call to ensure that you
	// receive all your results.
	IsTruncated *bool `type:"boolean"`

	// When IsTruncated is true, this element is present and contains the value
	// to use for the Marker parameter in a subsequent pagination request.
	Marker *string `type:"string"`

	// A list of policy versions.
	//
	// For more information about managed policy versions, see Versioning for Managed
	// Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-versions.html)
	// in the IAM User Guide.
	Versions []PolicyVersion `type:"list"`
}

// String returns the string representation
func (s ListPolicyVersionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListPolicyVersions = "ListPolicyVersions"

// ListPolicyVersionsRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Lists information about the versions of the specified managed policy, including
// the version that is currently set as the policy's default version.
//
// For more information about managed policies, see Managed Policies and Inline
// Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)
// in the IAM User Guide.
//
//    // Example sending a request using ListPolicyVersionsRequest.
//    req := client.ListPolicyVersionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/ListPolicyVersions
func (c *Client) ListPolicyVersionsRequest(input *ListPolicyVersionsInput) ListPolicyVersionsRequest {
	op := &aws.Operation{
		Name:       opListPolicyVersions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxItems",
			TruncationToken: "IsTruncated",
		},
	}

	if input == nil {
		input = &ListPolicyVersionsInput{}
	}

	req := c.newRequest(op, input, &ListPolicyVersionsOutput{})
	return ListPolicyVersionsRequest{Request: req, Input: input, Copy: c.ListPolicyVersionsRequest}
}

// ListPolicyVersionsRequest is the request type for the
// ListPolicyVersions API operation.
type ListPolicyVersionsRequest struct {
	*aws.Request
	Input *ListPolicyVersionsInput
	Copy  func(*ListPolicyVersionsInput) ListPolicyVersionsRequest
}

// Send marshals and sends the ListPolicyVersions API request.
func (r ListPolicyVersionsRequest) Send(ctx context.Context) (*ListPolicyVersionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPolicyVersionsResponse{
		ListPolicyVersionsOutput: r.Request.Data.(*ListPolicyVersionsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListPolicyVersionsRequestPaginator returns a paginator for ListPolicyVersions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListPolicyVersionsRequest(input)
//   p := iam.NewListPolicyVersionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListPolicyVersionsPaginator(req ListPolicyVersionsRequest) ListPolicyVersionsPaginator {
	return ListPolicyVersionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListPolicyVersionsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListPolicyVersionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListPolicyVersionsPaginator struct {
	aws.Pager
}

func (p *ListPolicyVersionsPaginator) CurrentPage() *ListPolicyVersionsOutput {
	return p.Pager.CurrentPage().(*ListPolicyVersionsOutput)
}

// ListPolicyVersionsResponse is the response type for the
// ListPolicyVersions API operation.
type ListPolicyVersionsResponse struct {
	*ListPolicyVersionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListPolicyVersions request.
func (r *ListPolicyVersionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
