// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/DeleteAccessKeyRequest
type DeleteAccessKeyInput struct {
	_ struct{} `type:"structure"`

	// The access key ID for the access key ID and secret access key you want to
	// delete.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters that can consist of any upper or lowercased letter
	// or digit.
	//
	// AccessKeyId is a required field
	AccessKeyId *string `min:"16" type:"string" required:"true"`

	// The name of the user whose access key pair you want to delete.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	UserName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DeleteAccessKeyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteAccessKeyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteAccessKeyInput"}

	if s.AccessKeyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccessKeyId"))
	}
	if s.AccessKeyId != nil && len(*s.AccessKeyId) < 16 {
		invalidParams.Add(aws.NewErrParamMinLen("AccessKeyId", 16))
	}
	if s.UserName != nil && len(*s.UserName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/DeleteAccessKeyOutput
type DeleteAccessKeyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteAccessKeyOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteAccessKey = "DeleteAccessKey"

// DeleteAccessKeyRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Deletes the access key pair associated with the specified IAM user.
//
// If you do not specify a user name, IAM determines the user name implicitly
// based on the AWS access key ID signing the request. This operation works
// for access keys under the AWS account. Consequently, you can use this operation
// to manage AWS account root user credentials even if the AWS account has no
// associated users.
//
//    // Example sending a request using DeleteAccessKeyRequest.
//    req := client.DeleteAccessKeyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/DeleteAccessKey
func (c *Client) DeleteAccessKeyRequest(input *DeleteAccessKeyInput) DeleteAccessKeyRequest {
	op := &aws.Operation{
		Name:       opDeleteAccessKey,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteAccessKeyInput{}
	}

	req := c.newRequest(op, input, &DeleteAccessKeyOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteAccessKeyRequest{Request: req, Input: input, Copy: c.DeleteAccessKeyRequest}
}

// DeleteAccessKeyRequest is the request type for the
// DeleteAccessKey API operation.
type DeleteAccessKeyRequest struct {
	*aws.Request
	Input *DeleteAccessKeyInput
	Copy  func(*DeleteAccessKeyInput) DeleteAccessKeyRequest
}

// Send marshals and sends the DeleteAccessKey API request.
func (r DeleteAccessKeyRequest) Send(ctx context.Context) (*DeleteAccessKeyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteAccessKeyResponse{
		DeleteAccessKeyOutput: r.Request.Data.(*DeleteAccessKeyOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteAccessKeyResponse is the response type for the
// DeleteAccessKey API operation.
type DeleteAccessKeyResponse struct {
	*DeleteAccessKeyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteAccessKey request.
func (r *DeleteAccessKeyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
