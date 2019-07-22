// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AssociateClientVpnTargetNetworkRequest
type AssociateClientVpnTargetNetworkInput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. For more information, see How to Ensure Idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	ClientToken *string `type:"string" idempotencyToken:"true"`

	// The ID of the Client VPN endpoint.
	//
	// ClientVpnEndpointId is a required field
	ClientVpnEndpointId *string `type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the subnet to associate with the Client VPN endpoint.
	//
	// SubnetId is a required field
	SubnetId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AssociateClientVpnTargetNetworkInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateClientVpnTargetNetworkInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateClientVpnTargetNetworkInput"}

	if s.ClientVpnEndpointId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientVpnEndpointId"))
	}

	if s.SubnetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubnetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AssociateClientVpnTargetNetworkResult
type AssociateClientVpnTargetNetworkOutput struct {
	_ struct{} `type:"structure"`

	// The unique ID of the target network association.
	AssociationId *string `locationName:"associationId" type:"string"`

	// The current state of the target network association.
	Status *AssociationStatus `locationName:"status" type:"structure"`
}

// String returns the string representation
func (s AssociateClientVpnTargetNetworkOutput) String() string {
	return awsutil.Prettify(s)
}

const opAssociateClientVpnTargetNetwork = "AssociateClientVpnTargetNetwork"

// AssociateClientVpnTargetNetworkRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Associates a target network with a Client VPN endpoint. A target network
// is a subnet in a VPC. You can associate multiple subnets from the same VPC
// with a Client VPN endpoint. You can associate only one subnet in each Availability
// Zone. We recommend that you associate at least two subnets to provide Availability
// Zone redundancy.
//
//    // Example sending a request using AssociateClientVpnTargetNetworkRequest.
//    req := client.AssociateClientVpnTargetNetworkRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AssociateClientVpnTargetNetwork
func (c *Client) AssociateClientVpnTargetNetworkRequest(input *AssociateClientVpnTargetNetworkInput) AssociateClientVpnTargetNetworkRequest {
	op := &aws.Operation{
		Name:       opAssociateClientVpnTargetNetwork,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssociateClientVpnTargetNetworkInput{}
	}

	req := c.newRequest(op, input, &AssociateClientVpnTargetNetworkOutput{})
	return AssociateClientVpnTargetNetworkRequest{Request: req, Input: input, Copy: c.AssociateClientVpnTargetNetworkRequest}
}

// AssociateClientVpnTargetNetworkRequest is the request type for the
// AssociateClientVpnTargetNetwork API operation.
type AssociateClientVpnTargetNetworkRequest struct {
	*aws.Request
	Input *AssociateClientVpnTargetNetworkInput
	Copy  func(*AssociateClientVpnTargetNetworkInput) AssociateClientVpnTargetNetworkRequest
}

// Send marshals and sends the AssociateClientVpnTargetNetwork API request.
func (r AssociateClientVpnTargetNetworkRequest) Send(ctx context.Context) (*AssociateClientVpnTargetNetworkResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateClientVpnTargetNetworkResponse{
		AssociateClientVpnTargetNetworkOutput: r.Request.Data.(*AssociateClientVpnTargetNetworkOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateClientVpnTargetNetworkResponse is the response type for the
// AssociateClientVpnTargetNetwork API operation.
type AssociateClientVpnTargetNetworkResponse struct {
	*AssociateClientVpnTargetNetworkOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateClientVpnTargetNetwork request.
func (r *AssociateClientVpnTargetNetworkResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
