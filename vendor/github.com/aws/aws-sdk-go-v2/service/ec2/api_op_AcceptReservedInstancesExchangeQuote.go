// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for accepting the quote.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AcceptReservedInstancesExchangeQuoteRequest
type AcceptReservedInstancesExchangeQuoteInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The IDs of the Convertible Reserved Instances to exchange for another Convertible
	// Reserved Instance of the same or higher value.
	//
	// ReservedInstanceIds is a required field
	ReservedInstanceIds []string `locationName:"ReservedInstanceId" locationNameList:"ReservedInstanceId" type:"list" required:"true"`

	// The configuration of the target Convertible Reserved Instance to exchange
	// for your current Convertible Reserved Instances.
	TargetConfigurations []TargetConfigurationRequest `locationName:"TargetConfiguration" locationNameList:"TargetConfigurationRequest" type:"list"`
}

// String returns the string representation
func (s AcceptReservedInstancesExchangeQuoteInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AcceptReservedInstancesExchangeQuoteInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AcceptReservedInstancesExchangeQuoteInput"}

	if s.ReservedInstanceIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReservedInstanceIds"))
	}
	if s.TargetConfigurations != nil {
		for i, v := range s.TargetConfigurations {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "TargetConfigurations", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The result of the exchange and whether it was successful.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AcceptReservedInstancesExchangeQuoteResult
type AcceptReservedInstancesExchangeQuoteOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the successful exchange.
	ExchangeId *string `locationName:"exchangeId" type:"string"`
}

// String returns the string representation
func (s AcceptReservedInstancesExchangeQuoteOutput) String() string {
	return awsutil.Prettify(s)
}

const opAcceptReservedInstancesExchangeQuote = "AcceptReservedInstancesExchangeQuote"

// AcceptReservedInstancesExchangeQuoteRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Accepts the Convertible Reserved Instance exchange quote described in the
// GetReservedInstancesExchangeQuote call.
//
//    // Example sending a request using AcceptReservedInstancesExchangeQuoteRequest.
//    req := client.AcceptReservedInstancesExchangeQuoteRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AcceptReservedInstancesExchangeQuote
func (c *Client) AcceptReservedInstancesExchangeQuoteRequest(input *AcceptReservedInstancesExchangeQuoteInput) AcceptReservedInstancesExchangeQuoteRequest {
	op := &aws.Operation{
		Name:       opAcceptReservedInstancesExchangeQuote,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AcceptReservedInstancesExchangeQuoteInput{}
	}

	req := c.newRequest(op, input, &AcceptReservedInstancesExchangeQuoteOutput{})
	return AcceptReservedInstancesExchangeQuoteRequest{Request: req, Input: input, Copy: c.AcceptReservedInstancesExchangeQuoteRequest}
}

// AcceptReservedInstancesExchangeQuoteRequest is the request type for the
// AcceptReservedInstancesExchangeQuote API operation.
type AcceptReservedInstancesExchangeQuoteRequest struct {
	*aws.Request
	Input *AcceptReservedInstancesExchangeQuoteInput
	Copy  func(*AcceptReservedInstancesExchangeQuoteInput) AcceptReservedInstancesExchangeQuoteRequest
}

// Send marshals and sends the AcceptReservedInstancesExchangeQuote API request.
func (r AcceptReservedInstancesExchangeQuoteRequest) Send(ctx context.Context) (*AcceptReservedInstancesExchangeQuoteResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AcceptReservedInstancesExchangeQuoteResponse{
		AcceptReservedInstancesExchangeQuoteOutput: r.Request.Data.(*AcceptReservedInstancesExchangeQuoteOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AcceptReservedInstancesExchangeQuoteResponse is the response type for the
// AcceptReservedInstancesExchangeQuote API operation.
type AcceptReservedInstancesExchangeQuoteResponse struct {
	*AcceptReservedInstancesExchangeQuoteOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AcceptReservedInstancesExchangeQuote request.
func (r *AcceptReservedInstancesExchangeQuoteResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
