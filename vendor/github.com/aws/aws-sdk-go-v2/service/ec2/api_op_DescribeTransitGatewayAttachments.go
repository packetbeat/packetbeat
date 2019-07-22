// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeTransitGatewayAttachmentsRequest
type DescribeTransitGatewayAttachmentsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// One or more filters. The possible values are:
	//
	//    * association.state - The state of the association (associating | associated
	//    | disassociating).
	//
	//    * association.transit-gateway-route-table-id - The ID of the route table
	//    for the transit gateway.
	//
	//    * resource-id - The ID of the resource.
	//
	//    * resource-owner-id - The ID of the AWS account that owns the resource.
	//
	//    * resource-type - The resource type (vpc | vpn).
	//
	//    * state - The state of the attachment (available | deleted | deleting
	//    | failed | modifying | pendingAcceptance | pending | rollingBack | rejected
	//    | rejecting).
	//
	//    * transit-gateway-attachment-id - The ID of the attachment.
	//
	//    * transit-gateway-id - The ID of the transit gateway.
	//
	//    * transit-gateway-owner-id - The ID of the AWS account that owns the transit
	//    gateway.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int64 `min:"5" type:"integer"`

	// The token for the next page of results.
	NextToken *string `type:"string"`

	// The IDs of the attachments.
	TransitGatewayAttachmentIds []string `type:"list"`
}

// String returns the string representation
func (s DescribeTransitGatewayAttachmentsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeTransitGatewayAttachmentsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeTransitGatewayAttachmentsInput"}
	if s.MaxResults != nil && *s.MaxResults < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 5))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeTransitGatewayAttachmentsResult
type DescribeTransitGatewayAttachmentsOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Information about the attachments.
	TransitGatewayAttachments []TransitGatewayAttachment `locationName:"transitGatewayAttachments" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeTransitGatewayAttachmentsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeTransitGatewayAttachments = "DescribeTransitGatewayAttachments"

// DescribeTransitGatewayAttachmentsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes one or more attachments between resources and transit gateways.
// By default, all attachments are described. Alternatively, you can filter
// the results by attachment ID, attachment state, resource ID, or resource
// owner.
//
//    // Example sending a request using DescribeTransitGatewayAttachmentsRequest.
//    req := client.DescribeTransitGatewayAttachmentsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeTransitGatewayAttachments
func (c *Client) DescribeTransitGatewayAttachmentsRequest(input *DescribeTransitGatewayAttachmentsInput) DescribeTransitGatewayAttachmentsRequest {
	op := &aws.Operation{
		Name:       opDescribeTransitGatewayAttachments,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeTransitGatewayAttachmentsInput{}
	}

	req := c.newRequest(op, input, &DescribeTransitGatewayAttachmentsOutput{})
	return DescribeTransitGatewayAttachmentsRequest{Request: req, Input: input, Copy: c.DescribeTransitGatewayAttachmentsRequest}
}

// DescribeTransitGatewayAttachmentsRequest is the request type for the
// DescribeTransitGatewayAttachments API operation.
type DescribeTransitGatewayAttachmentsRequest struct {
	*aws.Request
	Input *DescribeTransitGatewayAttachmentsInput
	Copy  func(*DescribeTransitGatewayAttachmentsInput) DescribeTransitGatewayAttachmentsRequest
}

// Send marshals and sends the DescribeTransitGatewayAttachments API request.
func (r DescribeTransitGatewayAttachmentsRequest) Send(ctx context.Context) (*DescribeTransitGatewayAttachmentsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeTransitGatewayAttachmentsResponse{
		DescribeTransitGatewayAttachmentsOutput: r.Request.Data.(*DescribeTransitGatewayAttachmentsOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeTransitGatewayAttachmentsRequestPaginator returns a paginator for DescribeTransitGatewayAttachments.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeTransitGatewayAttachmentsRequest(input)
//   p := ec2.NewDescribeTransitGatewayAttachmentsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeTransitGatewayAttachmentsPaginator(req DescribeTransitGatewayAttachmentsRequest) DescribeTransitGatewayAttachmentsPaginator {
	return DescribeTransitGatewayAttachmentsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeTransitGatewayAttachmentsInput
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

// DescribeTransitGatewayAttachmentsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeTransitGatewayAttachmentsPaginator struct {
	aws.Pager
}

func (p *DescribeTransitGatewayAttachmentsPaginator) CurrentPage() *DescribeTransitGatewayAttachmentsOutput {
	return p.Pager.CurrentPage().(*DescribeTransitGatewayAttachmentsOutput)
}

// DescribeTransitGatewayAttachmentsResponse is the response type for the
// DescribeTransitGatewayAttachments API operation.
type DescribeTransitGatewayAttachmentsResponse struct {
	*DescribeTransitGatewayAttachmentsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeTransitGatewayAttachments request.
func (r *DescribeTransitGatewayAttachmentsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
