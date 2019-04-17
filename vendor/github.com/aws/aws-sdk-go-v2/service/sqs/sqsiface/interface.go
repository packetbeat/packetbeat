// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package sqsiface provides an interface to enable mocking the Amazon Simple Queue Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package sqsiface

import (
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

// SQSAPI provides an interface to enable mocking the
// sqs.SQS service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Simple Queue Service.
//    func myFunc(svc sqsiface.SQSAPI) bool {
//        // Make svc.AddPermission request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := sqs.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockSQSClient struct {
//        sqsiface.SQSAPI
//    }
//    func (m *mockSQSClient) AddPermission(input *sqs.AddPermissionInput) (*sqs.AddPermissionOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockSQSClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type SQSAPI interface {
	AddPermissionRequest(*sqs.AddPermissionInput) sqs.AddPermissionRequest

	ChangeMessageVisibilityRequest(*sqs.ChangeMessageVisibilityInput) sqs.ChangeMessageVisibilityRequest

	ChangeMessageVisibilityBatchRequest(*sqs.ChangeMessageVisibilityBatchInput) sqs.ChangeMessageVisibilityBatchRequest

	CreateQueueRequest(*sqs.CreateQueueInput) sqs.CreateQueueRequest

	DeleteMessageRequest(*sqs.DeleteMessageInput) sqs.DeleteMessageRequest

	DeleteMessageBatchRequest(*sqs.DeleteMessageBatchInput) sqs.DeleteMessageBatchRequest

	DeleteQueueRequest(*sqs.DeleteQueueInput) sqs.DeleteQueueRequest

	GetQueueAttributesRequest(*sqs.GetQueueAttributesInput) sqs.GetQueueAttributesRequest

	GetQueueUrlRequest(*sqs.GetQueueUrlInput) sqs.GetQueueUrlRequest

	ListDeadLetterSourceQueuesRequest(*sqs.ListDeadLetterSourceQueuesInput) sqs.ListDeadLetterSourceQueuesRequest

	ListQueueTagsRequest(*sqs.ListQueueTagsInput) sqs.ListQueueTagsRequest

	ListQueuesRequest(*sqs.ListQueuesInput) sqs.ListQueuesRequest

	PurgeQueueRequest(*sqs.PurgeQueueInput) sqs.PurgeQueueRequest

	ReceiveMessageRequest(*sqs.ReceiveMessageInput) sqs.ReceiveMessageRequest

	RemovePermissionRequest(*sqs.RemovePermissionInput) sqs.RemovePermissionRequest

	SendMessageRequest(*sqs.SendMessageInput) sqs.SendMessageRequest

	SendMessageBatchRequest(*sqs.SendMessageBatchInput) sqs.SendMessageBatchRequest

	SetQueueAttributesRequest(*sqs.SetQueueAttributesInput) sqs.SetQueueAttributesRequest

	TagQueueRequest(*sqs.TagQueueInput) sqs.TagQueueRequest

	UntagQueueRequest(*sqs.UntagQueueInput) sqs.UntagQueueRequest
}

var _ SQSAPI = (*sqs.SQS)(nil)
