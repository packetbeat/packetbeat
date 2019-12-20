// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutBucketEncryptionRequest
type PutBucketEncryptionInput struct {
	_ struct{} `type:"structure" payload:"ServerSideEncryptionConfiguration"`

	// The name of the bucket for which the server-side encryption configuration
	// is set.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// Container for server-side encryption configuration rules. Currently S3 supports
	// one rule only.
	//
	// ServerSideEncryptionConfiguration is a required field
	ServerSideEncryptionConfiguration *ServerSideEncryptionConfiguration `locationName:"ServerSideEncryptionConfiguration" type:"structure" required:"true" xmlURI:"http://s3.amazonaws.com/doc/2006-03-01/"`
}

// String returns the string representation
func (s PutBucketEncryptionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutBucketEncryptionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutBucketEncryptionInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.ServerSideEncryptionConfiguration == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServerSideEncryptionConfiguration"))
	}
	if s.ServerSideEncryptionConfiguration != nil {
		if err := s.ServerSideEncryptionConfiguration.Validate(); err != nil {
			invalidParams.AddNested("ServerSideEncryptionConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *PutBucketEncryptionInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutBucketEncryptionInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.ServerSideEncryptionConfiguration != nil {
		v := s.ServerSideEncryptionConfiguration

		metadata := protocol.Metadata{XMLNamespaceURI: "http://s3.amazonaws.com/doc/2006-03-01/"}
		e.SetFields(protocol.PayloadTarget, "ServerSideEncryptionConfiguration", v, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutBucketEncryptionOutput
type PutBucketEncryptionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutBucketEncryptionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutBucketEncryptionOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opPutBucketEncryption = "PutBucketEncryption"

// PutBucketEncryptionRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Creates a new server-side encryption configuration (or replaces an existing
// one, if present).
//
//    // Example sending a request using PutBucketEncryptionRequest.
//    req := client.PutBucketEncryptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutBucketEncryption
func (c *Client) PutBucketEncryptionRequest(input *PutBucketEncryptionInput) PutBucketEncryptionRequest {
	op := &aws.Operation{
		Name:       opPutBucketEncryption,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}?encryption",
	}

	if input == nil {
		input = &PutBucketEncryptionInput{}
	}

	req := c.newRequest(op, input, &PutBucketEncryptionOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return PutBucketEncryptionRequest{Request: req, Input: input, Copy: c.PutBucketEncryptionRequest}
}

// PutBucketEncryptionRequest is the request type for the
// PutBucketEncryption API operation.
type PutBucketEncryptionRequest struct {
	*aws.Request
	Input *PutBucketEncryptionInput
	Copy  func(*PutBucketEncryptionInput) PutBucketEncryptionRequest
}

// Send marshals and sends the PutBucketEncryption API request.
func (r PutBucketEncryptionRequest) Send(ctx context.Context) (*PutBucketEncryptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutBucketEncryptionResponse{
		PutBucketEncryptionOutput: r.Request.Data.(*PutBucketEncryptionOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutBucketEncryptionResponse is the response type for the
// PutBucketEncryption API operation.
type PutBucketEncryptionResponse struct {
	*PutBucketEncryptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutBucketEncryption request.
func (r *PutBucketEncryptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
