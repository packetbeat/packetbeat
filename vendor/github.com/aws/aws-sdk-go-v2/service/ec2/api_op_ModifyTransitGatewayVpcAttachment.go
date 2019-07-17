// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyTransitGatewayVpcAttachmentRequest
type ModifyTransitGatewayVpcAttachmentInput struct {
	_ struct{} `type:"structure"`

	// The IDs of one or more subnets to add. You can specify at most one subnet
	// per Availability Zone.
	AddSubnetIds []string `locationNameList:"item" type:"list"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The new VPC attachment options.
	Options *ModifyTransitGatewayVpcAttachmentRequestOptions `type:"structure"`

	// The IDs of one or more subnets to remove.
	RemoveSubnetIds []string `locationNameList:"item" type:"list"`

	// The ID of the attachment.
	//
	// TransitGatewayAttachmentId is a required field
	TransitGatewayAttachmentId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ModifyTransitGatewayVpcAttachmentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyTransitGatewayVpcAttachmentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyTransitGatewayVpcAttachmentInput"}

	if s.TransitGatewayAttachmentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TransitGatewayAttachmentId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyTransitGatewayVpcAttachmentResult
type ModifyTransitGatewayVpcAttachmentOutput struct {
	_ struct{} `type:"structure"`

	// Information about the modified attachment.
	TransitGatewayVpcAttachment *TransitGatewayVpcAttachment `locationName:"transitGatewayVpcAttachment" type:"structure"`
}

// String returns the string representation
func (s ModifyTransitGatewayVpcAttachmentOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyTransitGatewayVpcAttachment = "ModifyTransitGatewayVpcAttachment"

// ModifyTransitGatewayVpcAttachmentRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Modifies the specified VPC attachment.
//
//    // Example sending a request using ModifyTransitGatewayVpcAttachmentRequest.
//    req := client.ModifyTransitGatewayVpcAttachmentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyTransitGatewayVpcAttachment
func (c *Client) ModifyTransitGatewayVpcAttachmentRequest(input *ModifyTransitGatewayVpcAttachmentInput) ModifyTransitGatewayVpcAttachmentRequest {
	op := &aws.Operation{
		Name:       opModifyTransitGatewayVpcAttachment,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyTransitGatewayVpcAttachmentInput{}
	}

	req := c.newRequest(op, input, &ModifyTransitGatewayVpcAttachmentOutput{})
	return ModifyTransitGatewayVpcAttachmentRequest{Request: req, Input: input, Copy: c.ModifyTransitGatewayVpcAttachmentRequest}
}

// ModifyTransitGatewayVpcAttachmentRequest is the request type for the
// ModifyTransitGatewayVpcAttachment API operation.
type ModifyTransitGatewayVpcAttachmentRequest struct {
	*aws.Request
	Input *ModifyTransitGatewayVpcAttachmentInput
	Copy  func(*ModifyTransitGatewayVpcAttachmentInput) ModifyTransitGatewayVpcAttachmentRequest
}

// Send marshals and sends the ModifyTransitGatewayVpcAttachment API request.
func (r ModifyTransitGatewayVpcAttachmentRequest) Send(ctx context.Context) (*ModifyTransitGatewayVpcAttachmentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyTransitGatewayVpcAttachmentResponse{
		ModifyTransitGatewayVpcAttachmentOutput: r.Request.Data.(*ModifyTransitGatewayVpcAttachmentOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyTransitGatewayVpcAttachmentResponse is the response type for the
// ModifyTransitGatewayVpcAttachment API operation.
type ModifyTransitGatewayVpcAttachmentResponse struct {
	*ModifyTransitGatewayVpcAttachmentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyTransitGatewayVpcAttachment request.
func (r *ModifyTransitGatewayVpcAttachmentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
