// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DisassociateAddressRequest
type DisassociateAddressInput struct {
	_ struct{} `type:"structure"`

	// [EC2-VPC] The association ID. Required for EC2-VPC.
	AssociationId *string `type:"string"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// [EC2-Classic] The Elastic IP address. Required for EC2-Classic.
	PublicIp *string `type:"string"`
}

// String returns the string representation
func (s DisassociateAddressInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DisassociateAddressOutput
type DisassociateAddressOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DisassociateAddressOutput) String() string {
	return awsutil.Prettify(s)
}

const opDisassociateAddress = "DisassociateAddress"

// DisassociateAddressRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Disassociates an Elastic IP address from the instance or network interface
// it's associated with.
//
// An Elastic IP address is for use in either the EC2-Classic platform or in
// a VPC. For more information, see Elastic IP Addresses (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/elastic-ip-addresses-eip.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
// This is an idempotent operation. If you perform the operation more than once,
// Amazon EC2 doesn't return an error.
//
//    // Example sending a request using DisassociateAddressRequest.
//    req := client.DisassociateAddressRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DisassociateAddress
func (c *Client) DisassociateAddressRequest(input *DisassociateAddressInput) DisassociateAddressRequest {
	op := &aws.Operation{
		Name:       opDisassociateAddress,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisassociateAddressInput{}
	}

	req := c.newRequest(op, input, &DisassociateAddressOutput{})
	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DisassociateAddressRequest{Request: req, Input: input, Copy: c.DisassociateAddressRequest}
}

// DisassociateAddressRequest is the request type for the
// DisassociateAddress API operation.
type DisassociateAddressRequest struct {
	*aws.Request
	Input *DisassociateAddressInput
	Copy  func(*DisassociateAddressInput) DisassociateAddressRequest
}

// Send marshals and sends the DisassociateAddress API request.
func (r DisassociateAddressRequest) Send(ctx context.Context) (*DisassociateAddressResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateAddressResponse{
		DisassociateAddressOutput: r.Request.Data.(*DisassociateAddressOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateAddressResponse is the response type for the
// DisassociateAddress API operation.
type DisassociateAddressResponse struct {
	*DisassociateAddressOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateAddress request.
func (r *DisassociateAddressResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
