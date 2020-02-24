// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchlogs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/DeleteDestinationRequest
type DeleteDestinationInput struct {
	_ struct{} `type:"structure"`

	// The name of the destination.
	//
	// DestinationName is a required field
	DestinationName *string `locationName:"destinationName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDestinationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDestinationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDestinationInput"}

	if s.DestinationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DestinationName"))
	}
	if s.DestinationName != nil && len(*s.DestinationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DestinationName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/DeleteDestinationOutput
type DeleteDestinationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteDestinationOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteDestination = "DeleteDestination"

// DeleteDestinationRequest returns a request value for making API operation for
// Amazon CloudWatch Logs.
//
// Deletes the specified destination, and eventually disables all the subscription
// filters that publish to it. This operation does not delete the physical resource
// encapsulated by the destination.
//
//    // Example sending a request using DeleteDestinationRequest.
//    req := client.DeleteDestinationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/DeleteDestination
func (c *Client) DeleteDestinationRequest(input *DeleteDestinationInput) DeleteDestinationRequest {
	op := &aws.Operation{
		Name:       opDeleteDestination,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDestinationInput{}
	}

	req := c.newRequest(op, input, &DeleteDestinationOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteDestinationRequest{Request: req, Input: input, Copy: c.DeleteDestinationRequest}
}

// DeleteDestinationRequest is the request type for the
// DeleteDestination API operation.
type DeleteDestinationRequest struct {
	*aws.Request
	Input *DeleteDestinationInput
	Copy  func(*DeleteDestinationInput) DeleteDestinationRequest
}

// Send marshals and sends the DeleteDestination API request.
func (r DeleteDestinationRequest) Send(ctx context.Context) (*DeleteDestinationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDestinationResponse{
		DeleteDestinationOutput: r.Request.Data.(*DeleteDestinationOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDestinationResponse is the response type for the
// DeleteDestination API operation.
type DeleteDestinationResponse struct {
	*DeleteDestinationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDestination request.
func (r *DeleteDestinationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
