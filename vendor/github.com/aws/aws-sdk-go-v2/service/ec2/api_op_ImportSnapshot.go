// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for ImportSnapshot.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ImportSnapshotRequest
type ImportSnapshotInput struct {
	_ struct{} `type:"structure"`

	// The client-specific data.
	ClientData *Data `type:"structure"`

	// Token to enable idempotency for VM import requests.
	ClientToken *string `type:"string"`

	// The description string for the import snapshot task.
	Description *string `type:"string"`

	// Information about the disk container.
	DiskContainer *SnapshotDiskContainer `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// Specifies whether the destination snapshot of the imported image should be
	// encrypted. The default CMK for EBS is used unless you specify a non-default
	// AWS Key Management Service (AWS KMS) CMK using KmsKeyId. For more information,
	// see Amazon EBS Encryption (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	Encrypted *bool `type:"boolean"`

	// An identifier for the AWS Key Management Service (AWS KMS) customer master
	// key (CMK) to use when creating the encrypted snapshot. This parameter is
	// only required if you want to use a non-default CMK; if this parameter is
	// not specified, the default CMK for EBS is used. If a KmsKeyId is specified,
	// the Encrypted flag must also be set.
	//
	// The CMK identifier may be provided in any of the following formats:
	//
	//    * Key ID
	//
	//    * Key alias. The alias ARN contains the arn:aws:kms namespace, followed
	//    by the Region of the CMK, the AWS account ID of the CMK owner, the alias
	//    namespace, and then the CMK alias. For example, arn:aws:kms:us-east-1:012345678910:alias/ExampleAlias.
	//
	//    * ARN using key ID. The ID ARN contains the arn:aws:kms namespace, followed
	//    by the Region of the CMK, the AWS account ID of the CMK owner, the key
	//    namespace, and then the CMK ID. For example, arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef.
	//
	//    * ARN using key alias. The alias ARN contains the arn:aws:kms namespace,
	//    followed by the Region of the CMK, the AWS account ID of the CMK owner,
	//    the alias namespace, and then the CMK alias. For example, arn:aws:kms:us-east-1:012345678910:alias/ExampleAlias.
	//
	// AWS parses KmsKeyId asynchronously, meaning that the action you call may
	// appear to complete even though you provided an invalid identifier. This action
	// will eventually report failure.
	//
	// The specified CMK must exist in the Region that the snapshot is being copied
	// to.
	KmsKeyId *string `type:"string"`

	// The name of the role to use when not using the default role, 'vmimport'.
	RoleName *string `type:"string"`
}

// String returns the string representation
func (s ImportSnapshotInput) String() string {
	return awsutil.Prettify(s)
}

// Contains the output for ImportSnapshot.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ImportSnapshotResult
type ImportSnapshotOutput struct {
	_ struct{} `type:"structure"`

	// A description of the import snapshot task.
	Description *string `locationName:"description" type:"string"`

	// The ID of the import snapshot task.
	ImportTaskId *string `locationName:"importTaskId" type:"string"`

	// Information about the import snapshot task.
	SnapshotTaskDetail *SnapshotTaskDetail `locationName:"snapshotTaskDetail" type:"structure"`
}

// String returns the string representation
func (s ImportSnapshotOutput) String() string {
	return awsutil.Prettify(s)
}

const opImportSnapshot = "ImportSnapshot"

// ImportSnapshotRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Imports a disk into an EBS snapshot.
//
//    // Example sending a request using ImportSnapshotRequest.
//    req := client.ImportSnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ImportSnapshot
func (c *Client) ImportSnapshotRequest(input *ImportSnapshotInput) ImportSnapshotRequest {
	op := &aws.Operation{
		Name:       opImportSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ImportSnapshotInput{}
	}

	req := c.newRequest(op, input, &ImportSnapshotOutput{})
	return ImportSnapshotRequest{Request: req, Input: input, Copy: c.ImportSnapshotRequest}
}

// ImportSnapshotRequest is the request type for the
// ImportSnapshot API operation.
type ImportSnapshotRequest struct {
	*aws.Request
	Input *ImportSnapshotInput
	Copy  func(*ImportSnapshotInput) ImportSnapshotRequest
}

// Send marshals and sends the ImportSnapshot API request.
func (r ImportSnapshotRequest) Send(ctx context.Context) (*ImportSnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ImportSnapshotResponse{
		ImportSnapshotOutput: r.Request.Data.(*ImportSnapshotOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ImportSnapshotResponse is the response type for the
// ImportSnapshot API operation.
type ImportSnapshotResponse struct {
	*ImportSnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ImportSnapshot request.
func (r *ImportSnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
