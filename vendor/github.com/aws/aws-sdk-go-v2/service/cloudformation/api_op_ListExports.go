// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/ListExportsInput
type ListExportsInput struct {
	_ struct{} `type:"structure"`

	// A string (provided by the ListExports response output) that identifies the
	// next page of exported output values that you asked to retrieve.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListExportsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListExportsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListExportsInput"}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/ListExportsOutput
type ListExportsOutput struct {
	_ struct{} `type:"structure"`

	// The output for the ListExports action.
	Exports []Export `type:"list"`

	// If the output exceeds 100 exported output values, a string that identifies
	// the next page of exports. If there is no additional page, this value is null.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListExportsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListExports = "ListExports"

// ListExportsRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Lists all exported output values in the account and region in which you call
// this action. Use this action to see the exported output values that you can
// import into other stacks. To import values, use the Fn::ImportValue (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-importvalue.html)
// function.
//
// For more information, see AWS CloudFormation Export Stack Output Values (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-stack-exports.html).
//
//    // Example sending a request using ListExportsRequest.
//    req := client.ListExportsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/ListExports
func (c *Client) ListExportsRequest(input *ListExportsInput) ListExportsRequest {
	op := &aws.Operation{
		Name:       opListExports,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListExportsInput{}
	}

	req := c.newRequest(op, input, &ListExportsOutput{})
	return ListExportsRequest{Request: req, Input: input, Copy: c.ListExportsRequest}
}

// ListExportsRequest is the request type for the
// ListExports API operation.
type ListExportsRequest struct {
	*aws.Request
	Input *ListExportsInput
	Copy  func(*ListExportsInput) ListExportsRequest
}

// Send marshals and sends the ListExports API request.
func (r ListExportsRequest) Send(ctx context.Context) (*ListExportsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListExportsResponse{
		ListExportsOutput: r.Request.Data.(*ListExportsOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListExportsRequestPaginator returns a paginator for ListExports.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListExportsRequest(input)
//   p := cloudformation.NewListExportsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListExportsPaginator(req ListExportsRequest) ListExportsPaginator {
	return ListExportsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListExportsInput
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

// ListExportsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListExportsPaginator struct {
	aws.Pager
}

func (p *ListExportsPaginator) CurrentPage() *ListExportsOutput {
	return p.Pager.CurrentPage().(*ListExportsOutput)
}

// ListExportsResponse is the response type for the
// ListExports API operation.
type ListExportsResponse struct {
	*ListExportsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListExports request.
func (r *ListExportsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
