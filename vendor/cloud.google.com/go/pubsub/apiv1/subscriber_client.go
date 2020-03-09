// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by gapic-generator. DO NOT EDIT.

package pubsub

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"github.com/golang/protobuf/proto"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/transport"
	pubsubpb "google.golang.org/genproto/googleapis/pubsub/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

// SubscriberCallOptions contains the retry settings for each method of SubscriberClient.
type SubscriberCallOptions struct {
	CreateSubscription []gax.CallOption
	GetSubscription    []gax.CallOption
	UpdateSubscription []gax.CallOption
	ListSubscriptions  []gax.CallOption
	DeleteSubscription []gax.CallOption
	ModifyAckDeadline  []gax.CallOption
	Acknowledge        []gax.CallOption
	Pull               []gax.CallOption
	StreamingPull      []gax.CallOption
	ModifyPushConfig   []gax.CallOption
	ListSnapshots      []gax.CallOption
	CreateSnapshot     []gax.CallOption
	UpdateSnapshot     []gax.CallOption
	DeleteSnapshot     []gax.CallOption
	Seek               []gax.CallOption
}

func defaultSubscriberClientOptions() []option.ClientOption {
	return []option.ClientOption{
		option.WithEndpoint("pubsub.googleapis.com:443"),
		option.WithScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultSubscriberCallOptions() *SubscriberCallOptions {
	retry := map[[2]string][]gax.CallOption{
		{"default", "idempotent"}: {
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Aborted,
					codes.Unavailable,
					codes.Unknown,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.3,
				})
			}),
		},
		{"default", "non_idempotent"}: {
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.3,
				})
			}),
		},
		{"messaging", "idempotent"}: {
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Aborted,
					codes.Unavailable,
					codes.Unknown,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.3,
				})
			}),
		},
		{"messaging", "non_idempotent"}: {
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.3,
				})
			}),
		},
	}
	return &SubscriberCallOptions{
		CreateSubscription: retry[[2]string{"default", "idempotent"}],
		GetSubscription:    retry[[2]string{"default", "idempotent"}],
		UpdateSubscription: retry[[2]string{"default", "non_idempotent"}],
		ListSubscriptions:  retry[[2]string{"default", "idempotent"}],
		DeleteSubscription: retry[[2]string{"default", "non_idempotent"}],
		ModifyAckDeadline:  retry[[2]string{"default", "non_idempotent"}],
		Acknowledge:        retry[[2]string{"messaging", "non_idempotent"}],
		Pull:               retry[[2]string{"messaging", "idempotent"}],
		StreamingPull:      retry[[2]string{"streaming_messaging", "none"}],
		ModifyPushConfig:   retry[[2]string{"default", "non_idempotent"}],
		ListSnapshots:      retry[[2]string{"default", "idempotent"}],
		CreateSnapshot:     retry[[2]string{"default", "non_idempotent"}],
		UpdateSnapshot:     retry[[2]string{"default", "non_idempotent"}],
		DeleteSnapshot:     retry[[2]string{"default", "non_idempotent"}],
		Seek:               retry[[2]string{"default", "idempotent"}],
	}
}

// SubscriberClient is a client for interacting with Google Cloud Pub/Sub API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type SubscriberClient struct {
	// The connection to the service.
	conn *grpc.ClientConn

	// The gRPC API client.
	subscriberClient pubsubpb.SubscriberClient

	// The call options for this service.
	CallOptions *SubscriberCallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewSubscriberClient creates a new subscriber client.
//
// The service that an application uses to manipulate subscriptions and to
// consume messages from a subscription via the Pull method or by
// establishing a bi-directional stream using the StreamingPull method.
func NewSubscriberClient(ctx context.Context, opts ...option.ClientOption) (*SubscriberClient, error) {
	conn, err := transport.DialGRPC(ctx, append(defaultSubscriberClientOptions(), opts...)...)
	if err != nil {
		return nil, err
	}
	c := &SubscriberClient{
		conn:        conn,
		CallOptions: defaultSubscriberCallOptions(),

		subscriberClient: pubsubpb.NewSubscriberClient(conn),
	}
	c.SetGoogleClientInfo()
	return c, nil
}

// Connection returns the client's connection to the API service.
func (c *SubscriberClient) Connection() *grpc.ClientConn {
	return c.conn
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *SubscriberClient) Close() error {
	return c.conn.Close()
}

// SetGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *SubscriberClient) SetGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// CreateSubscription creates a subscription to a given topic. See the
// <a href="https://cloud.google.com/pubsub/docs/admin#resource_names">
// resource name rules</a>.
// If the subscription already exists, returns ALREADY_EXISTS.
// If the corresponding topic doesn't exist, returns NOT_FOUND.
//
// If the name is not provided in the request, the server will assign a random
// name for this subscription on the same project as the topic, conforming
// to the
// resource name
// format (at https://cloud.google.com/pubsub/docs/admin#resource_names). The
// generated name is populated in the returned Subscription object. Note that
// for REST API requests, you must specify a name in the request.
func (c *SubscriberClient) CreateSubscription(ctx context.Context, req *pubsubpb.Subscription, opts ...gax.CallOption) (*pubsubpb.Subscription, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.CreateSubscription[0:len(c.CallOptions.CreateSubscription):len(c.CallOptions.CreateSubscription)], opts...)
	var resp *pubsubpb.Subscription
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.subscriberClient.CreateSubscription(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetSubscription gets the configuration details of a subscription.
func (c *SubscriberClient) GetSubscription(ctx context.Context, req *pubsubpb.GetSubscriptionRequest, opts ...gax.CallOption) (*pubsubpb.Subscription, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "subscription", url.QueryEscape(req.GetSubscription())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.GetSubscription[0:len(c.CallOptions.GetSubscription):len(c.CallOptions.GetSubscription)], opts...)
	var resp *pubsubpb.Subscription
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.subscriberClient.GetSubscription(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateSubscription updates an existing subscription. Note that certain properties of a
// subscription, such as its topic, are not modifiable.
func (c *SubscriberClient) UpdateSubscription(ctx context.Context, req *pubsubpb.UpdateSubscriptionRequest, opts ...gax.CallOption) (*pubsubpb.Subscription, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "subscription.name", url.QueryEscape(req.GetSubscription().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.UpdateSubscription[0:len(c.CallOptions.UpdateSubscription):len(c.CallOptions.UpdateSubscription)], opts...)
	var resp *pubsubpb.Subscription
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.subscriberClient.UpdateSubscription(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListSubscriptions lists matching subscriptions.
func (c *SubscriberClient) ListSubscriptions(ctx context.Context, req *pubsubpb.ListSubscriptionsRequest, opts ...gax.CallOption) *SubscriptionIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "project", url.QueryEscape(req.GetProject())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ListSubscriptions[0:len(c.CallOptions.ListSubscriptions):len(c.CallOptions.ListSubscriptions)], opts...)
	it := &SubscriptionIterator{}
	req = proto.Clone(req).(*pubsubpb.ListSubscriptionsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*pubsubpb.Subscription, string, error) {
		var resp *pubsubpb.ListSubscriptionsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.subscriberClient.ListSubscriptions(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}
		return resp.Subscriptions, resp.NextPageToken, nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.PageSize)
	it.pageInfo.Token = req.PageToken
	return it
}

// DeleteSubscription deletes an existing subscription. All messages retained in the subscription
// are immediately dropped. Calls to Pull after deletion will return
// NOT_FOUND. After a subscription is deleted, a new one may be created with
// the same name, but the new one has no association with the old
// subscription or its topic unless the same topic is specified.
func (c *SubscriberClient) DeleteSubscription(ctx context.Context, req *pubsubpb.DeleteSubscriptionRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "subscription", url.QueryEscape(req.GetSubscription())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.DeleteSubscription[0:len(c.CallOptions.DeleteSubscription):len(c.CallOptions.DeleteSubscription)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.subscriberClient.DeleteSubscription(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// ModifyAckDeadline modifies the ack deadline for a specific message. This method is useful
// to indicate that more time is needed to process a message by the
// subscriber, or to make the message available for redelivery if the
// processing was interrupted. Note that this does not modify the
// subscription-level ackDeadlineSeconds used for subsequent messages.
func (c *SubscriberClient) ModifyAckDeadline(ctx context.Context, req *pubsubpb.ModifyAckDeadlineRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "subscription", url.QueryEscape(req.GetSubscription())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ModifyAckDeadline[0:len(c.CallOptions.ModifyAckDeadline):len(c.CallOptions.ModifyAckDeadline)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.subscriberClient.ModifyAckDeadline(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// Acknowledge acknowledges the messages associated with the ack_ids in the
// AcknowledgeRequest. The Pub/Sub system can remove the relevant messages
// from the subscription.
//
// Acknowledging a message whose ack deadline has expired may succeed,
// but such a message may be redelivered later. Acknowledging a message more
// than once will not result in an error.
func (c *SubscriberClient) Acknowledge(ctx context.Context, req *pubsubpb.AcknowledgeRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "subscription", url.QueryEscape(req.GetSubscription())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.Acknowledge[0:len(c.CallOptions.Acknowledge):len(c.CallOptions.Acknowledge)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.subscriberClient.Acknowledge(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// Pull pulls messages from the server. The server may return UNAVAILABLE if
// there are too many concurrent pull requests pending for the given
// subscription.
func (c *SubscriberClient) Pull(ctx context.Context, req *pubsubpb.PullRequest, opts ...gax.CallOption) (*pubsubpb.PullResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "subscription", url.QueryEscape(req.GetSubscription())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.Pull[0:len(c.CallOptions.Pull):len(c.CallOptions.Pull)], opts...)
	var resp *pubsubpb.PullResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.subscriberClient.Pull(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// StreamingPull establishes a stream with the server, which sends messages down to the
// client. The client streams acknowledgements and ack deadline modifications
// back to the server. The server will close the stream and return the status
// on any error. The server may close the stream with status UNAVAILABLE to
// reassign server-side resources, in which case, the client should
// re-establish the stream. Flow control can be achieved by configuring the
// underlying RPC channel.
func (c *SubscriberClient) StreamingPull(ctx context.Context, opts ...gax.CallOption) (pubsubpb.Subscriber_StreamingPullClient, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.StreamingPull[0:len(c.CallOptions.StreamingPull):len(c.CallOptions.StreamingPull)], opts...)
	var resp pubsubpb.Subscriber_StreamingPullClient
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.subscriberClient.StreamingPull(ctx, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ModifyPushConfig modifies the PushConfig for a specified subscription.
//
// This may be used to change a push subscription to a pull one (signified by
// an empty PushConfig) or vice versa, or change the endpoint URL and other
// attributes of a push subscription. Messages will accumulate for delivery
// continuously through the call regardless of changes to the PushConfig.
func (c *SubscriberClient) ModifyPushConfig(ctx context.Context, req *pubsubpb.ModifyPushConfigRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "subscription", url.QueryEscape(req.GetSubscription())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ModifyPushConfig[0:len(c.CallOptions.ModifyPushConfig):len(c.CallOptions.ModifyPushConfig)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.subscriberClient.ModifyPushConfig(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// ListSnapshots lists the existing snapshots. Snapshots are used in
// <a href="https://cloud.google.com/pubsub/docs/replay-overview">Seek</a>
// operations, which allow
// you to manage message acknowledgments in bulk. That is, you can set the
// acknowledgment state of messages in an existing subscription to the state
// captured by a snapshot.
func (c *SubscriberClient) ListSnapshots(ctx context.Context, req *pubsubpb.ListSnapshotsRequest, opts ...gax.CallOption) *SnapshotIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "project", url.QueryEscape(req.GetProject())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ListSnapshots[0:len(c.CallOptions.ListSnapshots):len(c.CallOptions.ListSnapshots)], opts...)
	it := &SnapshotIterator{}
	req = proto.Clone(req).(*pubsubpb.ListSnapshotsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*pubsubpb.Snapshot, string, error) {
		var resp *pubsubpb.ListSnapshotsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.subscriberClient.ListSnapshots(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}
		return resp.Snapshots, resp.NextPageToken, nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.PageSize)
	it.pageInfo.Token = req.PageToken
	return it
}

// CreateSnapshot creates a snapshot from the requested subscription. Snapshots are used in
// <a href="https://cloud.google.com/pubsub/docs/replay-overview">Seek</a>
// operations, which allow
// you to manage message acknowledgments in bulk. That is, you can set the
// acknowledgment state of messages in an existing subscription to the state
// captured by a snapshot.
// <br><br>If the snapshot already exists, returns ALREADY_EXISTS.
// If the requested subscription doesn't exist, returns NOT_FOUND.
// If the backlog in the subscription is too old -- and the resulting snapshot
// would expire in less than 1 hour -- then FAILED_PRECONDITION is returned.
// See also the Snapshot.expire_time field. If the name is not provided in
// the request, the server will assign a random
// name for this snapshot on the same project as the subscription, conforming
// to the
// resource name
// format (at https://cloud.google.com/pubsub/docs/admin#resource_names). The
// generated name is populated in the returned Snapshot object. Note that for
// REST API requests, you must specify a name in the request.
func (c *SubscriberClient) CreateSnapshot(ctx context.Context, req *pubsubpb.CreateSnapshotRequest, opts ...gax.CallOption) (*pubsubpb.Snapshot, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.CreateSnapshot[0:len(c.CallOptions.CreateSnapshot):len(c.CallOptions.CreateSnapshot)], opts...)
	var resp *pubsubpb.Snapshot
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.subscriberClient.CreateSnapshot(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateSnapshot updates an existing snapshot. Snapshots are used in
// <a href="https://cloud.google.com/pubsub/docs/replay-overview">Seek</a>
// operations, which allow
// you to manage message acknowledgments in bulk. That is, you can set the
// acknowledgment state of messages in an existing subscription to the state
// captured by a snapshot.
func (c *SubscriberClient) UpdateSnapshot(ctx context.Context, req *pubsubpb.UpdateSnapshotRequest, opts ...gax.CallOption) (*pubsubpb.Snapshot, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "snapshot.name", url.QueryEscape(req.GetSnapshot().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.UpdateSnapshot[0:len(c.CallOptions.UpdateSnapshot):len(c.CallOptions.UpdateSnapshot)], opts...)
	var resp *pubsubpb.Snapshot
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.subscriberClient.UpdateSnapshot(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteSnapshot removes an existing snapshot. Snapshots are used in
// <a href="https://cloud.google.com/pubsub/docs/replay-overview">Seek</a>
// operations, which allow
// you to manage message acknowledgments in bulk. That is, you can set the
// acknowledgment state of messages in an existing subscription to the state
// captured by a snapshot.<br><br>
// When the snapshot is deleted, all messages retained in the snapshot
// are immediately dropped. After a snapshot is deleted, a new one may be
// created with the same name, but the new one has no association with the old
// snapshot or its subscription, unless the same subscription is specified.
func (c *SubscriberClient) DeleteSnapshot(ctx context.Context, req *pubsubpb.DeleteSnapshotRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "snapshot", url.QueryEscape(req.GetSnapshot())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.DeleteSnapshot[0:len(c.CallOptions.DeleteSnapshot):len(c.CallOptions.DeleteSnapshot)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.subscriberClient.DeleteSnapshot(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// Seek seeks an existing subscription to a point in time or to a given snapshot,
// whichever is provided in the request. Snapshots are used in
// <a href="https://cloud.google.com/pubsub/docs/replay-overview">Seek</a>
// operations, which allow
// you to manage message acknowledgments in bulk. That is, you can set the
// acknowledgment state of messages in an existing subscription to the state
// captured by a snapshot. Note that both the subscription and the snapshot
// must be on the same topic.
func (c *SubscriberClient) Seek(ctx context.Context, req *pubsubpb.SeekRequest, opts ...gax.CallOption) (*pubsubpb.SeekResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "subscription", url.QueryEscape(req.GetSubscription())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.Seek[0:len(c.CallOptions.Seek):len(c.CallOptions.Seek)], opts...)
	var resp *pubsubpb.SeekResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.subscriberClient.Seek(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SnapshotIterator manages a stream of *pubsubpb.Snapshot.
type SnapshotIterator struct {
	items    []*pubsubpb.Snapshot
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*pubsubpb.Snapshot, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *SnapshotIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *SnapshotIterator) Next() (*pubsubpb.Snapshot, error) {
	var item *pubsubpb.Snapshot
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *SnapshotIterator) bufLen() int {
	return len(it.items)
}

func (it *SnapshotIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// SubscriptionIterator manages a stream of *pubsubpb.Subscription.
type SubscriptionIterator struct {
	items    []*pubsubpb.Subscription
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*pubsubpb.Subscription, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *SubscriptionIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *SubscriptionIterator) Next() (*pubsubpb.Subscription, error) {
	var item *pubsubpb.Subscription
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *SubscriptionIterator) bufLen() int {
	return len(it.items)
}

func (it *SubscriptionIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
