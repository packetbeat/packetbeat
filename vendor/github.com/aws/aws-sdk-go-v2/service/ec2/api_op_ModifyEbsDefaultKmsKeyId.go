// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyEbsDefaultKmsKeyIdRequest
type ModifyEbsDefaultKmsKeyIdInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// An identifier for the AWS Key Management Service (AWS KMS) customer master
	// key (CMK) to use to encrypt the volume. This parameter is only required if
	// you want to use a non-default CMK; if this parameter is not specified, the
	// default CMK for EBS is used. If a KmsKeyId is specified, the Encrypted flag
	// must also be set.
	//
	// The CMK identifier may be provided in any of the following formats:
	//
	//    * Key ID
	//
	//    * Key alias
	//
	//    * ARN using key ID. The ID ARN contains the arn:aws:kms namespace, followed
	//    by the Region of the CMK, the AWS account ID of the CMK owner, the key
	//    namespace, and then the CMK ID. For example, arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef.
	//
	//    * ARN using key alias. The alias ARN contains the arn:aws:kms namespace,
	//    followed by the Region of the CMK, the AWS account ID of the CMK owner,
	//    the alias namespace, and then the CMK alias. For example, arn:aws:kms:us-east-1:012345678910:alias/ExampleAlias.
	//
	// KmsKeyId is a required field
	KmsKeyId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ModifyEbsDefaultKmsKeyIdInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyEbsDefaultKmsKeyIdInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyEbsDefaultKmsKeyIdInput"}

	if s.KmsKeyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("KmsKeyId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyEbsDefaultKmsKeyIdResult
type ModifyEbsDefaultKmsKeyIdOutput struct {
	_ struct{} `type:"structure"`

	// The full ARN of the default CMK that your account uses to encrypt an EBS
	// volume when no CMK is specified in the API call that creates the volume.
	KmsKeyId *string `locationName:"kmsKeyId" type:"string"`
}

// String returns the string representation
func (s ModifyEbsDefaultKmsKeyIdOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyEbsDefaultKmsKeyId = "ModifyEbsDefaultKmsKeyId"

// ModifyEbsDefaultKmsKeyIdRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Changes the default customer master key (CMK) that your account uses to encrypt
// EBS volumes if you don’t specify a CMK in the API call.
//
// Your account has an AWS-managed default CMK that is used for encrypting an
// EBS volume when no CMK is specified in the API call that creates the volume.
// By calling this API, you can specify a customer-managed CMK to use in place
// of the AWS-managed default CMK.
//
// Note: Deleting or disabling the custom CMK that you have specified to act
// as your default CMK will result in instance-launch failures.
//
//    // Example sending a request using ModifyEbsDefaultKmsKeyIdRequest.
//    req := client.ModifyEbsDefaultKmsKeyIdRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyEbsDefaultKmsKeyId
func (c *Client) ModifyEbsDefaultKmsKeyIdRequest(input *ModifyEbsDefaultKmsKeyIdInput) ModifyEbsDefaultKmsKeyIdRequest {
	op := &aws.Operation{
		Name:       opModifyEbsDefaultKmsKeyId,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyEbsDefaultKmsKeyIdInput{}
	}

	req := c.newRequest(op, input, &ModifyEbsDefaultKmsKeyIdOutput{})
	return ModifyEbsDefaultKmsKeyIdRequest{Request: req, Input: input, Copy: c.ModifyEbsDefaultKmsKeyIdRequest}
}

// ModifyEbsDefaultKmsKeyIdRequest is the request type for the
// ModifyEbsDefaultKmsKeyId API operation.
type ModifyEbsDefaultKmsKeyIdRequest struct {
	*aws.Request
	Input *ModifyEbsDefaultKmsKeyIdInput
	Copy  func(*ModifyEbsDefaultKmsKeyIdInput) ModifyEbsDefaultKmsKeyIdRequest
}

// Send marshals and sends the ModifyEbsDefaultKmsKeyId API request.
func (r ModifyEbsDefaultKmsKeyIdRequest) Send(ctx context.Context) (*ModifyEbsDefaultKmsKeyIdResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyEbsDefaultKmsKeyIdResponse{
		ModifyEbsDefaultKmsKeyIdOutput: r.Request.Data.(*ModifyEbsDefaultKmsKeyIdOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyEbsDefaultKmsKeyIdResponse is the response type for the
// ModifyEbsDefaultKmsKeyId API operation.
type ModifyEbsDefaultKmsKeyIdResponse struct {
	*ModifyEbsDefaultKmsKeyIdOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyEbsDefaultKmsKeyId request.
func (r *ModifyEbsDefaultKmsKeyIdResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
