// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutBucketAnalyticsConfigurationRequest
type PutBucketAnalyticsConfigurationInput struct {
	_ struct{} `type:"structure" payload:"AnalyticsConfiguration"`

	// The configuration and any analyses for the analytics filter.
	//
	// AnalyticsConfiguration is a required field
	AnalyticsConfiguration *AnalyticsConfiguration `locationName:"AnalyticsConfiguration" type:"structure" required:"true" xmlURI:"http://s3.amazonaws.com/doc/2006-03-01/"`

	// The name of the bucket to which an analytics configuration is stored.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// The identifier used to represent an analytics configuration.
	//
	// Id is a required field
	Id *string `location:"querystring" locationName:"id" type:"string" required:"true"`
}

// String returns the string representation
func (s PutBucketAnalyticsConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutBucketAnalyticsConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutBucketAnalyticsConfigurationInput"}

	if s.AnalyticsConfiguration == nil {
		invalidParams.Add(aws.NewErrParamRequired("AnalyticsConfiguration"))
	}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.AnalyticsConfiguration != nil {
		if err := s.AnalyticsConfiguration.Validate(); err != nil {
			invalidParams.AddNested("AnalyticsConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *PutBucketAnalyticsConfigurationInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutBucketAnalyticsConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.AnalyticsConfiguration != nil {
		v := s.AnalyticsConfiguration

		metadata := protocol.Metadata{XMLNamespaceURI: "http://s3.amazonaws.com/doc/2006-03-01/"}
		e.SetFields(protocol.PayloadTarget, "AnalyticsConfiguration", v, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "id", protocol.StringValue(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutBucketAnalyticsConfigurationOutput
type PutBucketAnalyticsConfigurationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutBucketAnalyticsConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutBucketAnalyticsConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opPutBucketAnalyticsConfiguration = "PutBucketAnalyticsConfiguration"

// PutBucketAnalyticsConfigurationRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Sets an analytics configuration for the bucket (specified by the analytics
// configuration ID).
//
//    // Example sending a request using PutBucketAnalyticsConfigurationRequest.
//    req := client.PutBucketAnalyticsConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutBucketAnalyticsConfiguration
func (c *Client) PutBucketAnalyticsConfigurationRequest(input *PutBucketAnalyticsConfigurationInput) PutBucketAnalyticsConfigurationRequest {
	op := &aws.Operation{
		Name:       opPutBucketAnalyticsConfiguration,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}?analytics",
	}

	if input == nil {
		input = &PutBucketAnalyticsConfigurationInput{}
	}

	req := c.newRequest(op, input, &PutBucketAnalyticsConfigurationOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return PutBucketAnalyticsConfigurationRequest{Request: req, Input: input, Copy: c.PutBucketAnalyticsConfigurationRequest}
}

// PutBucketAnalyticsConfigurationRequest is the request type for the
// PutBucketAnalyticsConfiguration API operation.
type PutBucketAnalyticsConfigurationRequest struct {
	*aws.Request
	Input *PutBucketAnalyticsConfigurationInput
	Copy  func(*PutBucketAnalyticsConfigurationInput) PutBucketAnalyticsConfigurationRequest
}

// Send marshals and sends the PutBucketAnalyticsConfiguration API request.
func (r PutBucketAnalyticsConfigurationRequest) Send(ctx context.Context) (*PutBucketAnalyticsConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutBucketAnalyticsConfigurationResponse{
		PutBucketAnalyticsConfigurationOutput: r.Request.Data.(*PutBucketAnalyticsConfigurationOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutBucketAnalyticsConfigurationResponse is the response type for the
// PutBucketAnalyticsConfiguration API operation.
type PutBucketAnalyticsConfigurationResponse struct {
	*PutBucketAnalyticsConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutBucketAnalyticsConfiguration request.
func (r *PutBucketAnalyticsConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
