// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetObjectLockConfigurationRequest
type GetObjectLockConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The bucket whose Object Lock configuration you want to retrieve.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`
}

// String returns the string representation
func (s GetObjectLockConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetObjectLockConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetObjectLockConfigurationInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *GetObjectLockConfigurationInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetObjectLockConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetObjectLockConfigurationOutput
type GetObjectLockConfigurationOutput struct {
	_ struct{} `type:"structure" payload:"ObjectLockConfiguration"`

	// The specified bucket's Object Lock configuration.
	ObjectLockConfiguration *ObjectLockConfiguration `type:"structure"`
}

// String returns the string representation
func (s GetObjectLockConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetObjectLockConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ObjectLockConfiguration != nil {
		v := s.ObjectLockConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "ObjectLockConfiguration", v, metadata)
	}
	return nil
}

const opGetObjectLockConfiguration = "GetObjectLockConfiguration"

// GetObjectLockConfigurationRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Gets the Object Lock configuration for a bucket. The rule specified in the
// Object Lock configuration will be applied by default to every new object
// placed in the specified bucket.
//
//    // Example sending a request using GetObjectLockConfigurationRequest.
//    req := client.GetObjectLockConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetObjectLockConfiguration
func (c *Client) GetObjectLockConfigurationRequest(input *GetObjectLockConfigurationInput) GetObjectLockConfigurationRequest {
	op := &aws.Operation{
		Name:       opGetObjectLockConfiguration,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?object-lock",
	}

	if input == nil {
		input = &GetObjectLockConfigurationInput{}
	}

	req := c.newRequest(op, input, &GetObjectLockConfigurationOutput{})
	return GetObjectLockConfigurationRequest{Request: req, Input: input, Copy: c.GetObjectLockConfigurationRequest}
}

// GetObjectLockConfigurationRequest is the request type for the
// GetObjectLockConfiguration API operation.
type GetObjectLockConfigurationRequest struct {
	*aws.Request
	Input *GetObjectLockConfigurationInput
	Copy  func(*GetObjectLockConfigurationInput) GetObjectLockConfigurationRequest
}

// Send marshals and sends the GetObjectLockConfiguration API request.
func (r GetObjectLockConfigurationRequest) Send(ctx context.Context) (*GetObjectLockConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetObjectLockConfigurationResponse{
		GetObjectLockConfigurationOutput: r.Request.Data.(*GetObjectLockConfigurationOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetObjectLockConfigurationResponse is the response type for the
// GetObjectLockConfiguration API operation.
type GetObjectLockConfigurationResponse struct {
	*GetObjectLockConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetObjectLockConfiguration request.
func (r *GetObjectLockConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
