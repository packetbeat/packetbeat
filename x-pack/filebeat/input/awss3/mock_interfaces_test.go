// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package awss3 is a generated GoMock package.
package awss3

import (
	context "context"
	reflect "reflect"
	time "time"

	s3 "github.com/aws/aws-sdk-go-v2/service/s3"
	sqs "github.com/aws/aws-sdk-go-v2/service/sqs"
	logp "github.com/elastic/beats/v7/libbeat/logp"
	gomock "github.com/golang/mock/gomock"
)

// MockSQSAPI is a mock of sqsAPI interface.
type MockSQSAPI struct {
	ctrl     *gomock.Controller
	recorder *MockSQSAPIMockRecorder
}

// MockSQSAPIMockRecorder is the mock recorder for MockSQSAPI.
type MockSQSAPIMockRecorder struct {
	mock *MockSQSAPI
}

// NewMockSQSAPI creates a new mock instance.
func NewMockSQSAPI(ctrl *gomock.Controller) *MockSQSAPI {
	mock := &MockSQSAPI{ctrl: ctrl}
	mock.recorder = &MockSQSAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSQSAPI) EXPECT() *MockSQSAPIMockRecorder {
	return m.recorder
}

// ChangeMessageVisibility mocks base method.
func (m *MockSQSAPI) ChangeMessageVisibility(ctx context.Context, msg *sqs.Message, timeout time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeMessageVisibility", ctx, msg, timeout)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeMessageVisibility indicates an expected call of ChangeMessageVisibility.
func (mr *MockSQSAPIMockRecorder) ChangeMessageVisibility(ctx, msg, timeout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeMessageVisibility", reflect.TypeOf((*MockSQSAPI)(nil).ChangeMessageVisibility), ctx, msg, timeout)
}

// DeleteMessage mocks base method.
func (m *MockSQSAPI) DeleteMessage(ctx context.Context, msg *sqs.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMessage", ctx, msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMessage indicates an expected call of DeleteMessage.
func (mr *MockSQSAPIMockRecorder) DeleteMessage(ctx, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMessage", reflect.TypeOf((*MockSQSAPI)(nil).DeleteMessage), ctx, msg)
}

// ReceiveMessage mocks base method.
func (m *MockSQSAPI) ReceiveMessage(ctx context.Context, maxMessages int) ([]sqs.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReceiveMessage", ctx, maxMessages)
	ret0, _ := ret[0].([]sqs.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReceiveMessage indicates an expected call of ReceiveMessage.
func (mr *MockSQSAPIMockRecorder) ReceiveMessage(ctx, maxMessages interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceiveMessage", reflect.TypeOf((*MockSQSAPI)(nil).ReceiveMessage), ctx, maxMessages)
}

// MocksqsReceiver is a mock of sqsReceiver interface.
type MocksqsReceiver struct {
	ctrl     *gomock.Controller
	recorder *MocksqsReceiverMockRecorder
}

// MocksqsReceiverMockRecorder is the mock recorder for MocksqsReceiver.
type MocksqsReceiverMockRecorder struct {
	mock *MocksqsReceiver
}

// NewMocksqsReceiver creates a new mock instance.
func NewMocksqsReceiver(ctrl *gomock.Controller) *MocksqsReceiver {
	mock := &MocksqsReceiver{ctrl: ctrl}
	mock.recorder = &MocksqsReceiverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocksqsReceiver) EXPECT() *MocksqsReceiverMockRecorder {
	return m.recorder
}

// ReceiveMessage mocks base method.
func (m *MocksqsReceiver) ReceiveMessage(ctx context.Context, maxMessages int) ([]sqs.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReceiveMessage", ctx, maxMessages)
	ret0, _ := ret[0].([]sqs.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReceiveMessage indicates an expected call of ReceiveMessage.
func (mr *MocksqsReceiverMockRecorder) ReceiveMessage(ctx, maxMessages interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceiveMessage", reflect.TypeOf((*MocksqsReceiver)(nil).ReceiveMessage), ctx, maxMessages)
}

// MocksqsDeleter is a mock of sqsDeleter interface.
type MocksqsDeleter struct {
	ctrl     *gomock.Controller
	recorder *MocksqsDeleterMockRecorder
}

// MocksqsDeleterMockRecorder is the mock recorder for MocksqsDeleter.
type MocksqsDeleterMockRecorder struct {
	mock *MocksqsDeleter
}

// NewMocksqsDeleter creates a new mock instance.
func NewMocksqsDeleter(ctrl *gomock.Controller) *MocksqsDeleter {
	mock := &MocksqsDeleter{ctrl: ctrl}
	mock.recorder = &MocksqsDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocksqsDeleter) EXPECT() *MocksqsDeleterMockRecorder {
	return m.recorder
}

// DeleteMessage mocks base method.
func (m *MocksqsDeleter) DeleteMessage(ctx context.Context, msg *sqs.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMessage", ctx, msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMessage indicates an expected call of DeleteMessage.
func (mr *MocksqsDeleterMockRecorder) DeleteMessage(ctx, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMessage", reflect.TypeOf((*MocksqsDeleter)(nil).DeleteMessage), ctx, msg)
}

// MocksqsVisibilityChanger is a mock of sqsVisibilityChanger interface.
type MocksqsVisibilityChanger struct {
	ctrl     *gomock.Controller
	recorder *MocksqsVisibilityChangerMockRecorder
}

// MocksqsVisibilityChangerMockRecorder is the mock recorder for MocksqsVisibilityChanger.
type MocksqsVisibilityChangerMockRecorder struct {
	mock *MocksqsVisibilityChanger
}

// NewMocksqsVisibilityChanger creates a new mock instance.
func NewMocksqsVisibilityChanger(ctrl *gomock.Controller) *MocksqsVisibilityChanger {
	mock := &MocksqsVisibilityChanger{ctrl: ctrl}
	mock.recorder = &MocksqsVisibilityChangerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocksqsVisibilityChanger) EXPECT() *MocksqsVisibilityChangerMockRecorder {
	return m.recorder
}

// ChangeMessageVisibility mocks base method.
func (m *MocksqsVisibilityChanger) ChangeMessageVisibility(ctx context.Context, msg *sqs.Message, timeout time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeMessageVisibility", ctx, msg, timeout)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeMessageVisibility indicates an expected call of ChangeMessageVisibility.
func (mr *MocksqsVisibilityChangerMockRecorder) ChangeMessageVisibility(ctx, msg, timeout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeMessageVisibility", reflect.TypeOf((*MocksqsVisibilityChanger)(nil).ChangeMessageVisibility), ctx, msg, timeout)
}

// MockSQSProcessor is a mock of sqsProcessor interface.
type MockSQSProcessor struct {
	ctrl     *gomock.Controller
	recorder *MockSQSProcessorMockRecorder
}

// MockSQSProcessorMockRecorder is the mock recorder for MockSQSProcessor.
type MockSQSProcessorMockRecorder struct {
	mock *MockSQSProcessor
}

// NewMockSQSProcessor creates a new mock instance.
func NewMockSQSProcessor(ctrl *gomock.Controller) *MockSQSProcessor {
	mock := &MockSQSProcessor{ctrl: ctrl}
	mock.recorder = &MockSQSProcessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSQSProcessor) EXPECT() *MockSQSProcessorMockRecorder {
	return m.recorder
}

// ProcessSQS mocks base method.
func (m *MockSQSProcessor) ProcessSQS(ctx context.Context, msg *sqs.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessSQS", ctx, msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessSQS indicates an expected call of ProcessSQS.
func (mr *MockSQSProcessorMockRecorder) ProcessSQS(ctx, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessSQS", reflect.TypeOf((*MockSQSProcessor)(nil).ProcessSQS), ctx, msg)
}

// MockS3API is a mock of s3API interface.
type MockS3API struct {
	ctrl     *gomock.Controller
	recorder *MockS3APIMockRecorder
}

// MockS3APIMockRecorder is the mock recorder for MockS3API.
type MockS3APIMockRecorder struct {
	mock *MockS3API
}

// NewMockS3API creates a new mock instance.
func NewMockS3API(ctrl *gomock.Controller) *MockS3API {
	mock := &MockS3API{ctrl: ctrl}
	mock.recorder = &MockS3APIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockS3API) EXPECT() *MockS3APIMockRecorder {
	return m.recorder
}

// GetObject mocks base method.
func (m *MockS3API) GetObject(ctx context.Context, bucket, key string) (*s3.GetObjectResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObject", ctx, bucket, key)
	ret0, _ := ret[0].(*s3.GetObjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObject indicates an expected call of GetObject.
func (mr *MockS3APIMockRecorder) GetObject(ctx, bucket, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObject", reflect.TypeOf((*MockS3API)(nil).GetObject), ctx, bucket, key)
}

// ListObjectsPaginator mocks base method.
func (m *MockS3API) ListObjectsPaginator(bucket, prefix string) s3Pager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListObjectsPaginator", bucket, prefix)
	ret0, _ := ret[0].(s3Pager)
	return ret0
}

// ListObjectsPaginator indicates an expected call of ListObjectsPaginator.
func (mr *MockS3APIMockRecorder) ListObjectsPaginator(bucket, prefix interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListObjectsPaginator", reflect.TypeOf((*MockS3API)(nil).ListObjectsPaginator), bucket, prefix)
}

// Mocks3Getter is a mock of s3Getter interface.
type Mocks3Getter struct {
	ctrl     *gomock.Controller
	recorder *Mocks3GetterMockRecorder
}

// Mocks3GetterMockRecorder is the mock recorder for Mocks3Getter.
type Mocks3GetterMockRecorder struct {
	mock *Mocks3Getter
}

// NewMocks3Getter creates a new mock instance.
func NewMocks3Getter(ctrl *gomock.Controller) *Mocks3Getter {
	mock := &Mocks3Getter{ctrl: ctrl}
	mock.recorder = &Mocks3GetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mocks3Getter) EXPECT() *Mocks3GetterMockRecorder {
	return m.recorder
}

// GetObject mocks base method.
func (m *Mocks3Getter) GetObject(ctx context.Context, bucket, key string) (*s3.GetObjectResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObject", ctx, bucket, key)
	ret0, _ := ret[0].(*s3.GetObjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObject indicates an expected call of GetObject.
func (mr *Mocks3GetterMockRecorder) GetObject(ctx, bucket, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObject", reflect.TypeOf((*Mocks3Getter)(nil).GetObject), ctx, bucket, key)
}

// Mocks3Lister is a mock of s3Lister interface.
type Mocks3Lister struct {
	ctrl     *gomock.Controller
	recorder *Mocks3ListerMockRecorder
}

// Mocks3ListerMockRecorder is the mock recorder for Mocks3Lister.
type Mocks3ListerMockRecorder struct {
	mock *Mocks3Lister
}

// NewMocks3Lister creates a new mock instance.
func NewMocks3Lister(ctrl *gomock.Controller) *Mocks3Lister {
	mock := &Mocks3Lister{ctrl: ctrl}
	mock.recorder = &Mocks3ListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mocks3Lister) EXPECT() *Mocks3ListerMockRecorder {
	return m.recorder
}

// ListObjectsPaginator mocks base method.
func (m *Mocks3Lister) ListObjectsPaginator(bucket, prefix string) s3Pager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListObjectsPaginator", bucket, prefix)
	ret0, _ := ret[0].(s3Pager)
	return ret0
}

// ListObjectsPaginator indicates an expected call of ListObjectsPaginator.
func (mr *Mocks3ListerMockRecorder) ListObjectsPaginator(bucket, prefix interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListObjectsPaginator", reflect.TypeOf((*Mocks3Lister)(nil).ListObjectsPaginator), bucket, prefix)
}

// MockS3Pager is a mock of s3Pager interface.
type MockS3Pager struct {
	ctrl     *gomock.Controller
	recorder *MockS3PagerMockRecorder
}

// MockS3PagerMockRecorder is the mock recorder for MockS3Pager.
type MockS3PagerMockRecorder struct {
	mock *MockS3Pager
}

// NewMockS3Pager creates a new mock instance.
func NewMockS3Pager(ctrl *gomock.Controller) *MockS3Pager {
	mock := &MockS3Pager{ctrl: ctrl}
	mock.recorder = &MockS3PagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockS3Pager) EXPECT() *MockS3PagerMockRecorder {
	return m.recorder
}

// CurrentPage mocks base method.
func (m *MockS3Pager) CurrentPage() *s3.ListObjectsOutput {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentPage")
	ret0, _ := ret[0].(*s3.ListObjectsOutput)
	return ret0
}

// CurrentPage indicates an expected call of CurrentPage.
func (mr *MockS3PagerMockRecorder) CurrentPage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentPage", reflect.TypeOf((*MockS3Pager)(nil).CurrentPage))
}

// Err mocks base method.
func (m *MockS3Pager) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err.
func (mr *MockS3PagerMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockS3Pager)(nil).Err))
}

// Next mocks base method.
func (m *MockS3Pager) Next(ctx context.Context) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next", ctx)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next.
func (mr *MockS3PagerMockRecorder) Next(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockS3Pager)(nil).Next), ctx)
}

// MockS3ObjectHandlerFactory is a mock of s3ObjectHandlerFactory interface.
type MockS3ObjectHandlerFactory struct {
	ctrl     *gomock.Controller
	recorder *MockS3ObjectHandlerFactoryMockRecorder
}

// MockS3ObjectHandlerFactoryMockRecorder is the mock recorder for MockS3ObjectHandlerFactory.
type MockS3ObjectHandlerFactoryMockRecorder struct {
	mock *MockS3ObjectHandlerFactory
}

// NewMockS3ObjectHandlerFactory creates a new mock instance.
func NewMockS3ObjectHandlerFactory(ctrl *gomock.Controller) *MockS3ObjectHandlerFactory {
	mock := &MockS3ObjectHandlerFactory{ctrl: ctrl}
	mock.recorder = &MockS3ObjectHandlerFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockS3ObjectHandlerFactory) EXPECT() *MockS3ObjectHandlerFactoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockS3ObjectHandlerFactory) Create(ctx context.Context, log *logp.Logger, acker *eventACKTracker, obj s3EventV2) s3ObjectHandler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, log, acker, obj)
	ret0, _ := ret[0].(s3ObjectHandler)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockS3ObjectHandlerFactoryMockRecorder) Create(ctx, log, acker, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockS3ObjectHandlerFactory)(nil).Create), ctx, log, acker, obj)
}

// MockS3ObjectHandler is a mock of s3ObjectHandler interface.
type MockS3ObjectHandler struct {
	ctrl     *gomock.Controller
	recorder *MockS3ObjectHandlerMockRecorder
}

// MockS3ObjectHandlerMockRecorder is the mock recorder for MockS3ObjectHandler.
type MockS3ObjectHandlerMockRecorder struct {
	mock *MockS3ObjectHandler
}

// NewMockS3ObjectHandler creates a new mock instance.
func NewMockS3ObjectHandler(ctrl *gomock.Controller) *MockS3ObjectHandler {
	mock := &MockS3ObjectHandler{ctrl: ctrl}
	mock.recorder = &MockS3ObjectHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockS3ObjectHandler) EXPECT() *MockS3ObjectHandlerMockRecorder {
	return m.recorder
}

// ProcessS3Object mocks base method.
func (m *MockS3ObjectHandler) ProcessS3Object() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessS3Object")
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessS3Object indicates an expected call of ProcessS3Object.
func (mr *MockS3ObjectHandlerMockRecorder) ProcessS3Object() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessS3Object", reflect.TypeOf((*MockS3ObjectHandler)(nil).ProcessS3Object))
}

// Wait mocks base method.
func (m *MockS3ObjectHandler) Wait() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Wait")
}

// Wait indicates an expected call of Wait.
func (mr *MockS3ObjectHandlerMockRecorder) Wait() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockS3ObjectHandler)(nil).Wait))
}
