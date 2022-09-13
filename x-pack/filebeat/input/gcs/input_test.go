// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package gcs

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"cloud.google.com/go/storage"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
	"google.golang.org/api/option"

	v2 "github.com/elastic/beats/v7/filebeat/input/v2"
	beattest "github.com/elastic/beats/v7/libbeat/publisher/testing"
	"github.com/elastic/beats/v7/x-pack/filebeat/input/gcs/mock"
	conf "github.com/elastic/elastic-agent-libs/config"
	"github.com/elastic/elastic-agent-libs/logp"
)

const (
	bucketGcsTestNew    = "gcs-test-new"
	bucketGcsTestLatest = "gcs-test-latest"
)

func Test_StorageClient(t *testing.T) {
	tests := []struct {
		name            string
		baseConfig      map[string]interface{}
		mockHandler     func() http.Handler
		expected        map[string]bool
		checkJSON       bool
		isError         error
		unexpectedError error
	}{
		{
			name: "Test1_SingleBucketWithPoll_NoErr",
			baseConfig: map[string]interface{}{
				"project_id":                 "elastic-sa",
				"auth.credentials_file.path": "/gcs_creds.json",
				"max_workers":                2,
				"poll":                       true,
				"poll_interval":              "5s",
				"parse_json":                 false,
				"buckets": []map[string]interface{}{
					{
						"name": "gcs-test-new",
					},
				},
			},
			mockHandler: mock.GCSServer,
			expected: map[string]bool{
				mock.Gcs_test_new_object_ata_json:      true,
				mock.Gcs_test_new_object_data3_json:    true,
				mock.Gcs_test_new_object_docs_ata_json: true,
			},
			unexpectedError: context.Canceled,
		},
		{
			name: "Test2_SingleBucketWithoutPoll_NoErr",
			baseConfig: map[string]interface{}{
				"project_id":                 "elastic-sa",
				"auth.credentials_file.path": "/gcs_creds.json",
				"max_workers":                2,
				"poll":                       false,
				"poll_interval":              "10s",
				"parse_json":                 false,
				"buckets": []map[string]interface{}{
					{
						"name": bucketGcsTestNew,
					},
				},
			},
			mockHandler: mock.GCSServer,
			expected: map[string]bool{
				mock.Gcs_test_new_object_ata_json:      true,
				mock.Gcs_test_new_object_data3_json:    true,
				mock.Gcs_test_new_object_docs_ata_json: true,
			},
			unexpectedError: nil,
		},
		{
			name: "Test3_TwoBucketsWithPoll_NoErr",
			baseConfig: map[string]interface{}{
				"project_id":                 "elastic-sa",
				"auth.credentials_file.path": "/gcs_creds.json",
				"max_workers":                2,
				"poll":                       true,
				"poll_interval":              "5s",
				"parse_json":                 false,
				"buckets": []map[string]interface{}{
					{
						"name": bucketGcsTestNew,
					},
					{
						"name": bucketGcsTestLatest,
					},
				},
			},
			mockHandler: mock.GCSServer,
			expected: map[string]bool{
				mock.Gcs_test_new_object_ata_json:      true,
				mock.Gcs_test_new_object_data3_json:    true,
				mock.Gcs_test_new_object_docs_ata_json: true,
				mock.Gcs_test_latest_object_ata_json:   true,
				mock.Gcs_test_latest_object_data3_json: true,
			},
			unexpectedError: context.Canceled,
		},
		{
			name: "Test4_TwoBucketsWithoutPoll_NoErr",
			baseConfig: map[string]interface{}{
				"project_id":                 "elastic-sa",
				"auth.credentials_file.path": "/gcs_creds.json",
				"max_workers":                2,
				"poll":                       false,
				"poll_interval":              "10s",
				"parse_json":                 false,
				"buckets": []map[string]interface{}{
					{
						"name": bucketGcsTestNew,
					},
					{
						"name": bucketGcsTestLatest,
					},
				},
			},
			mockHandler: mock.GCSServer,
			expected: map[string]bool{
				mock.Gcs_test_new_object_ata_json:      true,
				mock.Gcs_test_new_object_data3_json:    true,
				mock.Gcs_test_new_object_docs_ata_json: true,
				mock.Gcs_test_latest_object_ata_json:   true,
				mock.Gcs_test_latest_object_data3_json: true,
			},
			unexpectedError: nil,
		},
		{
			name: "Test5_SingleBucketWithPoll_InvalidBucketErr",
			baseConfig: map[string]interface{}{
				"project_id":                 "elastic-sa",
				"auth.credentials_file.path": "/gcs_creds.json",
				"max_workers":                2,
				"poll":                       true,
				"poll_interval":              "5s",
				"parse_json":                 false,
				"buckets": []map[string]interface{}{
					{
						"name": "gcs-test",
					},
				},
			},
			mockHandler:     mock.GCSServer,
			expected:        map[string]bool{},
			isError:         errors.New("storage: bucket doesn't exist"),
			unexpectedError: nil,
		},
		{
			name: "Test6_SingleBucketWithoutPoll_InvalidBucketErr",
			baseConfig: map[string]interface{}{
				"project_id":                 "elastic-sa",
				"auth.credentials_file.path": "/gcs_creds.json",
				"max_workers":                2,
				"poll":                       false,
				"poll_interval":              "5s",
				"parse_json":                 false,
				"buckets": []map[string]interface{}{
					{
						"name": "gcs-test",
					},
				},
			},
			mockHandler:     mock.GCSServer,
			expected:        map[string]bool{},
			isError:         errors.New("storage: bucket doesn't exist"),
			unexpectedError: nil,
		},
		{
			name: "Test7_TwoBucketsWithPoll_InvalidBucketErr",
			baseConfig: map[string]interface{}{
				"project_id":                 "elastic-sa",
				"auth.credentials_file.path": "/gcs_creds.json",
				"max_workers":                2,
				"poll":                       true,
				"poll_interval":              "5s",
				"parse_json":                 false,
				"buckets": []map[string]interface{}{
					{
						"name": "gcs-test",
					},
					{
						"name": "gcs-latest",
					},
				},
			},
			mockHandler:     mock.GCSServer,
			expected:        map[string]bool{},
			isError:         errors.New("storage: bucket doesn't exist"),
			unexpectedError: nil,
		},
		{
			name: "Test8_SingleBucketWithPoll_InvalidConfigValue",
			baseConfig: map[string]interface{}{
				"project_id":                 "elastic-sa",
				"auth.credentials_file.path": "/gcs_creds.json",
				"max_workers":                5100,
				"poll":                       true,
				"poll_interval":              "5s",
				"parse_json":                 false,
				"buckets": []map[string]interface{}{
					{
						"name": "gcs-test",
					},
				},
			},
			mockHandler:     mock.GCSServer,
			expected:        map[string]bool{},
			isError:         errors.New("requires value <= 5000 accessing 'max_workers'"),
			unexpectedError: nil,
		},
		{
			name: "Test9_TwoBucketWithPoll_InvalidConfigValue",
			baseConfig: map[string]interface{}{
				"project_id":                 "elastic-sa",
				"auth.credentials_file.path": "/gcs_creds.json",
				"max_workers":                5100,
				"poll":                       true,
				"poll_interval":              "5s",
				"parse_json":                 false,
				"buckets": []map[string]interface{}{
					{
						"name": "gcs-test",
					},
					{
						"name": "gcs-latest",
					},
				},
			},
			mockHandler:     mock.GCSServer,
			expected:        map[string]bool{},
			isError:         errors.New("requires value <= 5000 accessing 'max_workers'"),
			unexpectedError: nil,
		},
		{
			name: "Test10_SingleBucketWithPoll_parseJSON",
			baseConfig: map[string]interface{}{
				"project_id":                 "elastic-sa",
				"auth.credentials_file.path": "/gcs_creds.json",
				"max_workers":                1,
				"poll":                       true,
				"poll_interval":              "5s",
				"parse_json":                 true,
				"buckets": []map[string]interface{}{
					{
						"name": bucketGcsTestLatest,
					},
				},
			},
			mockHandler: mock.GCSServer,
			checkJSON:   true,
			expected: map[string]bool{
				mock.Gcs_test_latest_object_ata_json_parsed:   true,
				mock.Gcs_test_latest_object_data3_json_parsed: true,
			},
			unexpectedError: context.Canceled,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serv := httptest.NewServer(tt.mockHandler())
			httpclient := http.Client{
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{
						InsecureSkipVerify: true, //nolint:gosec // We can ignore as this is just for testing
					},
				},
			}
			t.Cleanup(serv.Close)

			client, _ := storage.NewClient(context.Background(), option.WithEndpoint(serv.URL), option.WithoutAuthentication(), option.WithHTTPClient(&httpclient))
			cfg := conf.MustNewConfigFrom(tt.baseConfig)
			conf := config{}
			err := cfg.Unpack(&conf)
			if err != nil {
				assert.EqualError(t, err, tt.isError.Error())
				return
			}
			input := newStatelessInput(conf)

			assert.Equal(t, "gcs-stateless", input.Name())
			assert.NoError(t, input.Test(v2.TestContext{}))

			chanClient := beattest.NewChanClient(len(tt.expected))
			t.Cleanup(func() { _ = chanClient.Close() })

			ctx, cancel := newV2Context()
			t.Cleanup(cancel)

			var g errgroup.Group
			g.Go(func() error {
				return input.Run(ctx, chanClient, client)
			})

			var timeout *time.Timer
			if conf.PollInterval != nil {
				timeout = time.NewTimer(1*time.Second + *conf.PollInterval)
			} else {
				timeout = time.NewTimer(5 * time.Second)
			}
			t.Cleanup(func() { _ = timeout.Stop() })

			if len(tt.expected) == 0 {
				if tt.isError != nil && g.Wait() != nil {
					assert.EqualError(t, g.Wait(), tt.isError.Error())
					cancel()
				} else {
					assert.NoError(t, g.Wait())
					cancel()
				}
				return
			}

			var receivedCount int
		wait:
			for {
				select {
				case <-timeout.C:
					t.Errorf("timed out waiting for %d events", len(tt.expected))
					cancel()
					return
				case got := <-chanClient.Channel:
					var val interface{}
					var err error
					if !tt.checkJSON {
						val, err = got.Fields.GetValue("message")
						assert.NoError(t, err)
						assert.True(t, tt.expected[val.(string)])
					} else {
						val, err = got.Fields.GetValue("gcs.storage.object.json_data")
						fVal := fmt.Sprintf("%v", val)
						assert.NoError(t, err)
						assert.True(t, tt.expected[fVal])
					}
					assert.Equal(t, tt.isError, err)
					receivedCount += 1
					if receivedCount == len(tt.expected) {
						cancel()
						break wait
					}
				}
			}
			assert.ErrorIs(t, tt.unexpectedError, g.Wait())
		})
	}
}

func newV2Context() (v2.Context, func()) {
	ctx, cancel := context.WithCancel(context.Background())
	return v2.Context{
		Logger:      logp.NewLogger("gcs_test"),
		ID:          "test_id",
		Cancelation: ctx,
	}, cancel
}
