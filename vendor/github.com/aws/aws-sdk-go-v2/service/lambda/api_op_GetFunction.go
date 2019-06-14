// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/GetFunctionRequest
type GetFunctionInput struct {
	_ struct{} `type:"structure"`

	// The name of the Lambda function, version, or alias.
	//
	// Name formats
	//
	//    * Function name - my-function (name-only), my-function:v1 (with alias).
	//
	//    * Function ARN - arn:aws:lambda:us-west-2:123456789012:function:my-function.
	//
	//    * Partial ARN - 123456789012:function:my-function.
	//
	// You can append a version number or alias to any of the formats. The length
	// constraint applies only to the full ARN. If you specify only the function
	// name, it is limited to 64 characters in length.
	//
	// FunctionName is a required field
	FunctionName *string `location:"uri" locationName:"FunctionName" min:"1" type:"string" required:"true"`

	// Specify a version or alias to get details about a published version of the
	// function.
	Qualifier *string `location:"querystring" locationName:"Qualifier" min:"1" type:"string"`
}

// String returns the string representation
func (s GetFunctionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetFunctionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetFunctionInput"}

	if s.FunctionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("FunctionName"))
	}
	if s.FunctionName != nil && len(*s.FunctionName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FunctionName", 1))
	}
	if s.Qualifier != nil && len(*s.Qualifier) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Qualifier", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetFunctionInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.FunctionName != nil {
		v := *s.FunctionName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "FunctionName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Qualifier != nil {
		v := *s.Qualifier

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "Qualifier", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/GetFunctionResponse
type GetFunctionOutput struct {
	_ struct{} `type:"structure"`

	// The deployment package of the function or version.
	Code *FunctionCodeLocation `type:"structure"`

	// The function's reserved concurrency (https://docs.aws.amazon.com/lambda/latest/dg/concurrent-executions.html).
	Concurrency *Concurrency `type:"structure"`

	// The configuration of the function or version.
	Configuration *FunctionConfiguration `type:"structure"`

	// The function's tags (https://docs.aws.amazon.com/lambda/latest/dg/tagging.html).
	Tags map[string]string `type:"map"`
}

// String returns the string representation
func (s GetFunctionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetFunctionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Code != nil {
		v := s.Code

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Code", v, metadata)
	}
	if s.Concurrency != nil {
		v := s.Concurrency

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Concurrency", v, metadata)
	}
	if s.Configuration != nil {
		v := s.Configuration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Configuration", v, metadata)
	}
	if len(s.Tags) > 0 {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "Tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

const opGetFunction = "GetFunction"

// GetFunctionRequest returns a request value for making API operation for
// AWS Lambda.
//
// Returns information about the function or function version, with a link to
// download the deployment package that's valid for 10 minutes. If you specify
// a function version, only details that are specific to that version are returned.
//
//    // Example sending a request using GetFunctionRequest.
//    req := client.GetFunctionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/GetFunction
func (c *Client) GetFunctionRequest(input *GetFunctionInput) GetFunctionRequest {
	op := &aws.Operation{
		Name:       opGetFunction,
		HTTPMethod: "GET",
		HTTPPath:   "/2015-03-31/functions/{FunctionName}",
	}

	if input == nil {
		input = &GetFunctionInput{}
	}

	req := c.newRequest(op, input, &GetFunctionOutput{})
	return GetFunctionRequest{Request: req, Input: input, Copy: c.GetFunctionRequest}
}

// GetFunctionRequest is the request type for the
// GetFunction API operation.
type GetFunctionRequest struct {
	*aws.Request
	Input *GetFunctionInput
	Copy  func(*GetFunctionInput) GetFunctionRequest
}

// Send marshals and sends the GetFunction API request.
func (r GetFunctionRequest) Send(ctx context.Context) (*GetFunctionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetFunctionResponse{
		GetFunctionOutput: r.Request.Data.(*GetFunctionOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetFunctionResponse is the response type for the
// GetFunction API operation.
type GetFunctionResponse struct {
	*GetFunctionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetFunction request.
func (r *GetFunctionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
