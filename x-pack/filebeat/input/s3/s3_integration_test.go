// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// +build integration
// +build aws

package s3

/*
import (
	"context"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go-v2/aws"
	awssdk "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/s3iface"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/sqsiface"

	"github.com/elastic/beats/v7/filebeat/input"
	v2 "github.com/elastic/beats/v7/filebeat/input/v2"
	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/logp"
	pubtest "github.com/elastic/beats/v7/libbeat/publisher/testing"
	"github.com/elastic/beats/v7/libbeat/tests/resources"
	awscommon "github.com/elastic/beats/v7/x-pack/libbeat/common/aws"
)

const (
	fileName          = "sample1.txt"
	visibilityTimeout = 300 * time.Second
)

var filePath = filepath.Join("ftest", fileName)

// GetConfigForTest function gets aws credentials for integration tests.
func getConfigForTest(t *testing.T) config {
	t.Helper()

	awsConfig := awscommon.ConfigAWS{}
	queueURL := os.Getenv("QUEUE_URL")
	profileName := os.Getenv("AWS_PROFILE_NAME")
	accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	sessionToken := os.Getenv("AWS_SESSION_TOKEN")

	config := config{
		VisibilityTimeout: visibilityTimeout,
	}
	switch {
	case queueURL == "":
		t.Fatal("$QUEUE_URL is not set in environment")
	case profileName == "" && accessKeyID == "":
		t.Fatal("$AWS_ACCESS_KEY_ID or $AWS_PROFILE_NAME not set or set to empty")
	case profileName != "":
		awsConfig.ProfileName = profileName
		config.QueueURL = queueURL
		config.AwsConfig = awsConfig
		return config
	case secretAccessKey == "":
		t.Fatal("$AWS_SECRET_ACCESS_KEY not set or set to empty")
	}

	awsConfig.AccessKeyID = accessKeyID
	awsConfig.SecretAccessKey = secretAccessKey
	if sessionToken != "" {
		awsConfig.SessionToken = sessionToken
	}
	config.AwsConfig = awsConfig
	return config
}

func uploadSampleLogFile(t *testing.T, bucketName string, svcS3 s3iface.ClientAPI) {
	t.Helper()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("Failed to open file %v", filePath)
	}
	defer file.Close()

	s3PutObjectInput := s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filepath.Base(filePath)),
		Body:   file,
	}
	req := svcS3.PutObjectRequest(&s3PutObjectInput)
	output, err := req.Send(ctx)
	if err != nil {
		t.Fatalf("failed to put object into s3 bucket: %v", output)
	}
}

func collectOldMessages(t *testing.T, queueURL string, visibilityTimeout int64, svcSQS sqsiface.ClientAPI) []string {
	t.Helper()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// receive messages from sqs
	req := svcSQS.ReceiveMessageRequest(
		&sqs.ReceiveMessageInput{
			QueueUrl:              &queueURL,
			MessageAttributeNames: []string{"All"},
			MaxNumberOfMessages:   &maxNumberOfMessage,
			VisibilityTimeout:     &visibilityTimeout,
			WaitTimeSeconds:       &waitTimeSecond,
		})

	output, err := req.Send(ctx)
	if err != nil {
		t.Fatalf("failed to receive message from SQS: %v", output)
	}

	var oldMessageHandles []string
	for _, message := range output.Messages {
		oldMessageHandles = append(oldMessageHandles, *message.ReceiptHandle)
	}

	return oldMessageHandles
}

func (input *s3Collector) deleteAllMessages(t *testing.T, awsConfig awssdk.Config, queueURL string, visibilityTimeout int64, svcSQS sqsiface.ClientAPI) error {
	var messageReceiptHandles []string
	init := true
	for init || len(messageReceiptHandles) > 0 {
		init = false
		messageReceiptHandles = collectOldMessages(t, queueURL, visibilityTimeout, svcSQS)
		for _, receiptHandle := range messageReceiptHandles {
			err := input.deleteMessage(queueURL, receiptHandle, svcSQS)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func defaultTestConfig() *common.Config {
	return common.MustNewConfigFrom(map[string]interface{}{
		"queue_url": os.Getenv("QUEUE_URL"),
	})
}

func runTest(t *testing.T, cfg *common.Config, run func(t *testing.T, input *s3Collector, receiver *eventReceiver)) {
	plugin := Plugin()
	inp, err := plugin.Manager.Create(cfg)
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client := pubtest.NewChanClient(0)
	receiver := &eventReceiver{client.Channel}
	pipeline := pubtest.ConstClient(client)
	collector, err := inp.(*s3Input).createCollector(v2.Context{
		Logger:      logp.NewLogger("test"),
		Cancelation: ctx,
	}, pipeline)
	if err != nil {
		t.Fatal(err)
	}

	run(t, collector, receiver)
}

type eventReceiver struct {
	ch chan beat.Event
}

func (r *eventReceiver) waitForEvents(n int) ([]beat.Event, bool) {
	buf := make([]beat.Event, 0, n)
	for event := range r.ch {
		buf = append(buf, event)
		n--
		if n == 0 {
			return buf, true
		}
	}
	return buf, false
}

func TestS3Input(t *testing.T) {
	inputConfig := defaultTestConfig()
	config := getConfigForTest(t)

	runTest(t, inputConfig, func(t *testing.T, input *s3Collector, receiver *eventReceiver) {
		awsConfig, err := awscommon.GetAWSCredentials(config.AwsConfig)
		if err != nil {

		}
		s3BucketRegion := os.Getenv("S3_BUCKET_REGION")
		if s3BucketRegion == "" {
			t.Log("S3_BUCKET_REGION is not set, default to us-west-1")
			s3BucketRegion = "us-west-1"
		}
		awsConfig.Region = s3BucketRegion
		awsConfig = awsConfig.Copy()
		input.svc := sqs.New(awsConfig)

		// remove old messages from SQS
		err = input.deleteAllMessages(t, awsConfig, config.QueueURL, int64(config.VisibilityTimeout.Seconds()), svcSQS)
		if err != nil {
			t.Fatalf("failed to delete message: %v", err.Error())
		}

		// upload a sample log file for testing
		s3BucketNameEnv := os.Getenv("S3_BUCKET_NAME")
		if s3BucketNameEnv == "" {
			t.Fatal("failed to get S3_BUCKET_NAME")
		}

		svcS3 := s3.New(awsConfig)
		uploadSampleLogFile(t, s3BucketNameEnv, svcS3)
		time.Sleep(30 * time.Second)

		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			input.run(svcSQS, svcS3, 300)
		}()

		events, ok := out.waitForEvents(2)
		if !ok {
			t.Fatalf("Expected 2 events, but got %d.", len(events))
		}
		input.Stop()

		// check events
		for i, event := range events {
			bucketName, err := event.GetValue("aws.s3.bucket.name")
			assert.NoError(t, err)
			assert.Equal(t, s3BucketNameEnv, bucketName)

			objectKey, err := event.GetValue("aws.s3.object.key")
			assert.NoError(t, err)
			assert.Equal(t, fileName, objectKey)

			message, err := event.GetValue("message")
			assert.NoError(t, err)
			switch i {
			case 0:
				assert.Equal(t, "logline1\n", message)
			case 1:
				assert.Equal(t, "logline2\n", message)
			}
		}

		// delete messages from the queue
		err = input.deleteAllMessages(t, awsConfig, config.QueueURL, int64(config.VisibilityTimeout.Seconds()), svcSQS)
		if err != nil {
			t.Fatalf("failed to delete message: %v", err.Error())
		}
	})
}

// MockSQSClient struct is used for unit tests.
type MockSQSClient struct {
	sqsiface.ClientAPI
}

var (
	sqsMessageTest = "{\"Records\":[{\"eventSource\":\"aws:s3\",\"awsRegion\":\"ap-southeast-1\"," +
		"\"eventTime\":\"2019-06-21T16:16:54.629Z\",\"eventName\":\"ObjectCreated:Put\"," +
		"\"s3\":{\"configurationId\":\"object-created-event\",\"bucket\":{\"name\":\"test-s3-ks-2\"," +
		"\"arn\":\"arn:aws:s3:::test-s3-ks-2\"},\"object\":{\"key\":\"server-access-logging2019-06-21-16-16-54\"}}}]}"
)

func (m *MockSQSClient) ReceiveMessageRequest(input *sqs.ReceiveMessageInput) sqs.ReceiveMessageRequest {
	httpReq, _ := http.NewRequest("", "", nil)
	return sqs.ReceiveMessageRequest{
		Request: &awssdk.Request{
			Data: &sqs.ReceiveMessageOutput{
				Messages: []sqs.Message{
					{Body: awssdk.String(sqsMessageTest)},
				},
			},
			HTTPRequest: httpReq,
		},
	}
}

func (m *MockSQSClient) DeleteMessageRequest(input *sqs.DeleteMessageInput) sqs.DeleteMessageRequest {
	httpReq, _ := http.NewRequest("", "", nil)
	return sqs.DeleteMessageRequest{
		Request: &awssdk.Request{
			Data:        &sqs.DeleteMessageOutput{},
			HTTPRequest: httpReq,
		},
	}
}

func (m *MockSQSClient) ChangeMessageVisibilityRequest(input *sqs.ChangeMessageVisibilityInput) sqs.ChangeMessageVisibilityRequest {
	httpReq, _ := http.NewRequest("", "", nil)
	return sqs.ChangeMessageVisibilityRequest{
		Request: &awssdk.Request{
			Data:        &sqs.ChangeMessageVisibilityOutput{},
			HTTPRequest: httpReq,
		},
	}
}

func TestMockS3Input(t *testing.T) {
	defer resources.NewGoroutinesChecker().Check(t)
	cfg := common.MustNewConfigFrom(map[string]interface{}{
		"queue_url": "https://sqs.ap-southeast-1.amazonaws.com/123456/test",
	})

	runTest(t, cfg, func(t *testing.T, input *s3Input, out *stubOutleter) {
		svcS3 := &MockS3Client{}
		svcSQS := &MockSQSClient{}

		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			input.run(svcSQS, svcS3, 300)
		}()

		events, ok := out.waitForEvents(2)
		if !ok {
			t.Fatalf("Expected 2 events, but got %d.", len(events))
		}
		input.Wait()

		// check events
		for i, event := range events {
			bucketName, err := event.GetValue("aws.s3.bucket.name")
			assert.NoError(t, err)
			assert.Equal(t, "test-s3-ks-2", bucketName)

			objectKey, err := event.GetValue("aws.s3.object.key")
			assert.NoError(t, err)
			assert.Equal(t, "server-access-logging2019-06-21-16-16-54", objectKey)

			message, err := event.GetValue("message")
			assert.NoError(t, err)
			switch i {
			case 0:
				assert.Equal(t, s3LogString1, message)
			case 1:
				assert.Equal(t, s3LogString2, message)
			}
		}
	})
}
*/
