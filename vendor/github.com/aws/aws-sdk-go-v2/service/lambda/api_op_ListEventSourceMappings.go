// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/ListEventSourceMappingsRequest
type ListEventSourceMappingsInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the event source.
	//
	//    * Amazon Kinesis - The ARN of the data stream or a stream consumer.
	//
	//    * Amazon DynamoDB Streams - The ARN of the stream.
	//
	//    * Amazon Simple Queue Service - The ARN of the queue.
	EventSourceArn *string `location:"querystring" locationName:"EventSourceArn" type:"string"`

	// The name of the Lambda function.
	//
	// Name formats
	//
	//    * Function name - MyFunction.
	//
	//    * Function ARN - arn:aws:lambda:us-west-2:123456789012:function:MyFunction.
	//
	//    * Version or Alias ARN - arn:aws:lambda:us-west-2:123456789012:function:MyFunction:PROD.
	//
	//    * Partial ARN - 123456789012:function:MyFunction.
	//
	// The length constraint applies only to the full ARN. If you specify only the
	// function name, it's limited to 64 characters in length.
	FunctionName *string `location:"querystring" locationName:"FunctionName" min:"1" type:"string"`

	// A pagination token returned by a previous call.
	Marker *string `location:"querystring" locationName:"Marker" type:"string"`

	// The maximum number of event source mappings to return.
	MaxItems *int64 `location:"querystring" locationName:"MaxItems" min:"1" type:"integer"`
}

// String returns the string representation
func (s ListEventSourceMappingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListEventSourceMappingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListEventSourceMappingsInput"}
	if s.FunctionName != nil && len(*s.FunctionName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FunctionName", 1))
	}
	if s.MaxItems != nil && *s.MaxItems < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxItems", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListEventSourceMappingsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.EventSourceArn != nil {
		v := *s.EventSourceArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "EventSourceArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FunctionName != nil {
		v := *s.FunctionName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "FunctionName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "Marker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxItems != nil {
		v := *s.MaxItems

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "MaxItems", protocol.Int64Value(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/ListEventSourceMappingsResponse
type ListEventSourceMappingsOutput struct {
	_ struct{} `type:"structure"`

	// A list of event source mappings.
	EventSourceMappings []EventSourceMappingConfiguration `type:"list"`

	// A pagination token that's returned when the response doesn't contain all
	// event source mappings.
	NextMarker *string `type:"string"`
}

// String returns the string representation
func (s ListEventSourceMappingsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListEventSourceMappingsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.EventSourceMappings) > 0 {
		v := s.EventSourceMappings

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "EventSourceMappings", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextMarker != nil {
		v := *s.NextMarker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextMarker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListEventSourceMappings = "ListEventSourceMappings"

// ListEventSourceMappingsRequest returns a request value for making API operation for
// AWS Lambda.
//
// Lists event source mappings. Specify an EventSourceArn to only show event
// source mappings for a single event source.
//
//    // Example sending a request using ListEventSourceMappingsRequest.
//    req := client.ListEventSourceMappingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/ListEventSourceMappings
func (c *Client) ListEventSourceMappingsRequest(input *ListEventSourceMappingsInput) ListEventSourceMappingsRequest {
	op := &aws.Operation{
		Name:       opListEventSourceMappings,
		HTTPMethod: "GET",
		HTTPPath:   "/2015-03-31/event-source-mappings/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"NextMarker"},
			LimitToken:      "MaxItems",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListEventSourceMappingsInput{}
	}

	req := c.newRequest(op, input, &ListEventSourceMappingsOutput{})
	return ListEventSourceMappingsRequest{Request: req, Input: input, Copy: c.ListEventSourceMappingsRequest}
}

// ListEventSourceMappingsRequest is the request type for the
// ListEventSourceMappings API operation.
type ListEventSourceMappingsRequest struct {
	*aws.Request
	Input *ListEventSourceMappingsInput
	Copy  func(*ListEventSourceMappingsInput) ListEventSourceMappingsRequest
}

// Send marshals and sends the ListEventSourceMappings API request.
func (r ListEventSourceMappingsRequest) Send(ctx context.Context) (*ListEventSourceMappingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListEventSourceMappingsResponse{
		ListEventSourceMappingsOutput: r.Request.Data.(*ListEventSourceMappingsOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListEventSourceMappingsRequestPaginator returns a paginator for ListEventSourceMappings.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListEventSourceMappingsRequest(input)
//   p := lambda.NewListEventSourceMappingsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListEventSourceMappingsPaginator(req ListEventSourceMappingsRequest) ListEventSourceMappingsPaginator {
	return ListEventSourceMappingsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListEventSourceMappingsInput
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

// ListEventSourceMappingsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListEventSourceMappingsPaginator struct {
	aws.Pager
}

func (p *ListEventSourceMappingsPaginator) CurrentPage() *ListEventSourceMappingsOutput {
	return p.Pager.CurrentPage().(*ListEventSourceMappingsOutput)
}

// ListEventSourceMappingsResponse is the response type for the
// ListEventSourceMappings API operation.
type ListEventSourceMappingsResponse struct {
	*ListEventSourceMappingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListEventSourceMappings request.
func (r *ListEventSourceMappingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
