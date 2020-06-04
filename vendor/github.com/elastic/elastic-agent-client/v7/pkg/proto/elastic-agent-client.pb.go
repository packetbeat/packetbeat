// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: elastic-agent-client.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Status codes for the current state.
type StateObserved_Status int32

const (
	// Application is starting.
	StateObserved_STARTING StateObserved_Status = 0
	// Application is currently configuring.
	StateObserved_CONFIGURING StateObserved_Status = 1
	// Application is in healthy state.
	StateObserved_HEALTHY StateObserved_Status = 2
	// Application is working but in a degraded state.
	StateObserved_DEGRADED StateObserved_Status = 3
	// Application is failing completely.
	StateObserved_FAILED StateObserved_Status = 4
	// Application is stopping.
	StateObserved_STOPPING StateObserved_Status = 5
)

// Enum value maps for StateObserved_Status.
var (
	StateObserved_Status_name = map[int32]string{
		0: "STARTING",
		1: "CONFIGURING",
		2: "HEALTHY",
		3: "DEGRADED",
		4: "FAILED",
		5: "STOPPING",
	}
	StateObserved_Status_value = map[string]int32{
		"STARTING":    0,
		"CONFIGURING": 1,
		"HEALTHY":     2,
		"DEGRADED":    3,
		"FAILED":      4,
		"STOPPING":    5,
	}
)

func (x StateObserved_Status) Enum() *StateObserved_Status {
	p := new(StateObserved_Status)
	*p = x
	return p
}

func (x StateObserved_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StateObserved_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_elastic_agent_client_proto_enumTypes[0].Descriptor()
}

func (StateObserved_Status) Type() protoreflect.EnumType {
	return &file_elastic_agent_client_proto_enumTypes[0]
}

func (x StateObserved_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StateObserved_Status.Descriptor instead.
func (StateObserved_Status) EnumDescriptor() ([]byte, []int) {
	return file_elastic_agent_client_proto_rawDescGZIP(), []int{1, 0}
}

type StateExpected_State int32

const (
	// Expects that the application is running.
	StateExpected_RUNNING StateExpected_State = 0
	// Expects that the application is stopping.
	StateExpected_STOPPING StateExpected_State = 1
)

// Enum value maps for StateExpected_State.
var (
	StateExpected_State_name = map[int32]string{
		0: "RUNNING",
		1: "STOPPING",
	}
	StateExpected_State_value = map[string]int32{
		"RUNNING":  0,
		"STOPPING": 1,
	}
)

func (x StateExpected_State) Enum() *StateExpected_State {
	p := new(StateExpected_State)
	*p = x
	return p
}

func (x StateExpected_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StateExpected_State) Descriptor() protoreflect.EnumDescriptor {
	return file_elastic_agent_client_proto_enumTypes[1].Descriptor()
}

func (StateExpected_State) Type() protoreflect.EnumType {
	return &file_elastic_agent_client_proto_enumTypes[1]
}

func (x StateExpected_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StateExpected_State.Descriptor instead.
func (StateExpected_State) EnumDescriptor() ([]byte, []int) {
	return file_elastic_agent_client_proto_rawDescGZIP(), []int{2, 0}
}

// Status result of the action.
type ActionResponse_Status int32

const (
	// Action was successful.
	ActionResponse_SUCCESS ActionResponse_Status = 0
	// Action has failed.
	ActionResponse_FAILED ActionResponse_Status = 1
)

// Enum value maps for ActionResponse_Status.
var (
	ActionResponse_Status_name = map[int32]string{
		0: "SUCCESS",
		1: "FAILED",
	}
	ActionResponse_Status_value = map[string]int32{
		"SUCCESS": 0,
		"FAILED":  1,
	}
)

func (x ActionResponse_Status) Enum() *ActionResponse_Status {
	p := new(ActionResponse_Status)
	*p = x
	return p
}

func (x ActionResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActionResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_elastic_agent_client_proto_enumTypes[2].Descriptor()
}

func (ActionResponse_Status) Type() protoreflect.EnumType {
	return &file_elastic_agent_client_proto_enumTypes[2]
}

func (x ActionResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActionResponse_Status.Descriptor instead.
func (ActionResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_elastic_agent_client_proto_rawDescGZIP(), []int{4, 0}
}

// Connection information sent to the application on startup so it knows how to connected back to the Elastic Agent.
//
// This is normally sent through stdin and should never be sent across a network un-encrypted.
type ConnInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// GRPC connection address.
	Addr string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	// Server name to use when connecting over TLS.
	ServerName string `protobuf:"bytes,2,opt,name=server_name,json=serverName,proto3" json:"server_name,omitempty"`
	// Token that the application should send as the unique identifier when connecting over the GRPC.
	Token string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	// CA certificate.
	CaCert []byte `protobuf:"bytes,4,opt,name=ca_cert,json=caCert,proto3" json:"ca_cert,omitempty"`
	// Peer certificate.
	PeerCert []byte `protobuf:"bytes,5,opt,name=peer_cert,json=peerCert,proto3" json:"peer_cert,omitempty"`
	// Peer private key.
	PeerKey []byte `protobuf:"bytes,6,opt,name=peer_key,json=peerKey,proto3" json:"peer_key,omitempty"`
}

func (x *ConnInfo) Reset() {
	*x = ConnInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_elastic_agent_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnInfo) ProtoMessage() {}

func (x *ConnInfo) ProtoReflect() protoreflect.Message {
	mi := &file_elastic_agent_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnInfo.ProtoReflect.Descriptor instead.
func (*ConnInfo) Descriptor() ([]byte, []int) {
	return file_elastic_agent_client_proto_rawDescGZIP(), []int{0}
}

func (x *ConnInfo) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *ConnInfo) GetServerName() string {
	if x != nil {
		return x.ServerName
	}
	return ""
}

func (x *ConnInfo) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ConnInfo) GetCaCert() []byte {
	if x != nil {
		return x.CaCert
	}
	return nil
}

func (x *ConnInfo) GetPeerCert() []byte {
	if x != nil {
		return x.PeerCert
	}
	return nil
}

func (x *ConnInfo) GetPeerKey() []byte {
	if x != nil {
		return x.PeerKey
	}
	return nil
}

// A status observed message is streamed from the application to Elastic Agent.
//
// This message contains the currently applied `config_state_idx` (0 in the case of initial start, 1 is the first
// applied config index) along with the status of the application. In the case that the sent `config_state_idx`
// doesn't match the expected `config_state_idx` that Elastic Agent expects, the application is always marked as
// `CONFIGURING`.
type StateObserved struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Token that is used to unique identify the application to agent. When agent started this
	// application it would have provided it this token.
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// Current index of the applied configuration.
	ConfigStateIdx uint64 `protobuf:"varint,2,opt,name=config_state_idx,json=configStateIdx,proto3" json:"config_state_idx,omitempty"`
	// Status code.
	Status StateObserved_Status `protobuf:"varint,3,opt,name=status,proto3,enum=proto.StateObserved_Status" json:"status,omitempty"`
	// Message for the health status.
	Message string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *StateObserved) Reset() {
	*x = StateObserved{}
	if protoimpl.UnsafeEnabled {
		mi := &file_elastic_agent_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateObserved) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateObserved) ProtoMessage() {}

func (x *StateObserved) ProtoReflect() protoreflect.Message {
	mi := &file_elastic_agent_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateObserved.ProtoReflect.Descriptor instead.
func (*StateObserved) Descriptor() ([]byte, []int) {
	return file_elastic_agent_client_proto_rawDescGZIP(), []int{1}
}

func (x *StateObserved) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *StateObserved) GetConfigStateIdx() uint64 {
	if x != nil {
		return x.ConfigStateIdx
	}
	return 0
}

func (x *StateObserved) GetStatus() StateObserved_Status {
	if x != nil {
		return x.Status
	}
	return StateObserved_STARTING
}

func (x *StateObserved) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// A state expected message is streamed from the Elastic Agent to the application informing the application
// what Elastic Agent expects the applications state to be.
type StateExpected struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Expected state of the application.
	State StateExpected_State `protobuf:"varint,1,opt,name=state,proto3,enum=proto.StateExpected_State" json:"state,omitempty"`
	// Index of the either current configuration or new configuration provided.
	ConfigStateIdx uint64 `protobuf:"varint,2,opt,name=config_state_idx,json=configStateIdx,proto3" json:"config_state_idx,omitempty"`
	// Resulting configuration. (If the application already has the current `config_state_idx` this
	// will be empty.)
	Config string `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *StateExpected) Reset() {
	*x = StateExpected{}
	if protoimpl.UnsafeEnabled {
		mi := &file_elastic_agent_client_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateExpected) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateExpected) ProtoMessage() {}

func (x *StateExpected) ProtoReflect() protoreflect.Message {
	mi := &file_elastic_agent_client_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateExpected.ProtoReflect.Descriptor instead.
func (*StateExpected) Descriptor() ([]byte, []int) {
	return file_elastic_agent_client_proto_rawDescGZIP(), []int{2}
}

func (x *StateExpected) GetState() StateExpected_State {
	if x != nil {
		return x.State
	}
	return StateExpected_RUNNING
}

func (x *StateExpected) GetConfigStateIdx() uint64 {
	if x != nil {
		return x.ConfigStateIdx
	}
	return 0
}

func (x *StateExpected) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

// A action request is streamed from the Elastic Agent to the application so an action can be performed
// by the connected application.
type ActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique ID of the action.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name of the action.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// JSON encoded parameters for the action.
	Params []byte `protobuf:"bytes,3,opt,name=params,proto3" json:"params,omitempty"`
}

func (x *ActionRequest) Reset() {
	*x = ActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_elastic_agent_client_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionRequest) ProtoMessage() {}

func (x *ActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_elastic_agent_client_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionRequest.ProtoReflect.Descriptor instead.
func (*ActionRequest) Descriptor() ([]byte, []int) {
	return file_elastic_agent_client_proto_rawDescGZIP(), []int{3}
}

func (x *ActionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ActionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ActionRequest) GetParams() []byte {
	if x != nil {
		return x.Params
	}
	return nil
}

// An action response is streamed from the application back to the Elastic Agent to provide a result to
// an action request.
type ActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Token that is used to unique identify the application to agent. When agent started this
	// application it would have provided it this token.
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// Unique ID of the action.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Status of the action.
	Status ActionResponse_Status `protobuf:"varint,3,opt,name=status,proto3,enum=proto.ActionResponse_Status" json:"status,omitempty"`
	// JSON encoded result for the action.
	Result []byte `protobuf:"bytes,4,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ActionResponse) Reset() {
	*x = ActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_elastic_agent_client_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionResponse) ProtoMessage() {}

func (x *ActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_elastic_agent_client_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionResponse.ProtoReflect.Descriptor instead.
func (*ActionResponse) Descriptor() ([]byte, []int) {
	return file_elastic_agent_client_proto_rawDescGZIP(), []int{4}
}

func (x *ActionResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ActionResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ActionResponse) GetStatus() ActionResponse_Status {
	if x != nil {
		return x.Status
	}
	return ActionResponse_SUCCESS
}

func (x *ActionResponse) GetResult() []byte {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_elastic_agent_client_proto protoreflect.FileDescriptor

var file_elastic_agent_client_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2d,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x01, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x61, 0x64, 0x64, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x63,
	0x61, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x63, 0x61,
	0x43, 0x65, 0x72, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x72,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x70, 0x65, 0x65, 0x72, 0x43, 0x65, 0x72,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x65, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x22, 0xfc, 0x01, 0x0a,
	0x0d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x64, 0x78, 0x12, 0x33,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x64, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x5c, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x54, 0x41, 0x52, 0x54,
	0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x55,
	0x52, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48,
	0x59, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x47, 0x52, 0x41, 0x44, 0x45, 0x44, 0x10,
	0x03, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x12, 0x0c, 0x0a,
	0x08, 0x53, 0x54, 0x4f, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x22, 0xa7, 0x01, 0x0a, 0x0d,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x30, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x28, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f,
	0x69, 0x64, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x64, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x22, 0x22, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55,
	0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x54, 0x4f, 0x50, 0x50,
	0x49, 0x4e, 0x47, 0x10, 0x01, 0x22, 0x4b, 0x0a, 0x0d, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x22, 0xa7, 0x01, 0x0a, 0x0e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x34, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x21, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x32, 0x85, 0x01, 0x0a,
	0x0c, 0x45, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x39, 0x0a,
	0x07, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x1a, 0x14,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x78, 0x70, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x28, 0x01, 0x30, 0x01, 0x12, 0x3a, 0x0a, 0x07, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x28, 0x01, 0x30, 0x01, 0x42, 0x14, 0x5a, 0x0f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_elastic_agent_client_proto_rawDescOnce sync.Once
	file_elastic_agent_client_proto_rawDescData = file_elastic_agent_client_proto_rawDesc
)

func file_elastic_agent_client_proto_rawDescGZIP() []byte {
	file_elastic_agent_client_proto_rawDescOnce.Do(func() {
		file_elastic_agent_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_elastic_agent_client_proto_rawDescData)
	})
	return file_elastic_agent_client_proto_rawDescData
}

var file_elastic_agent_client_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_elastic_agent_client_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_elastic_agent_client_proto_goTypes = []interface{}{
	(StateObserved_Status)(0),  // 0: proto.StateObserved.Status
	(StateExpected_State)(0),   // 1: proto.StateExpected.State
	(ActionResponse_Status)(0), // 2: proto.ActionResponse.Status
	(*ConnInfo)(nil),           // 3: proto.ConnInfo
	(*StateObserved)(nil),      // 4: proto.StateObserved
	(*StateExpected)(nil),      // 5: proto.StateExpected
	(*ActionRequest)(nil),      // 6: proto.ActionRequest
	(*ActionResponse)(nil),     // 7: proto.ActionResponse
}
var file_elastic_agent_client_proto_depIdxs = []int32{
	0, // 0: proto.StateObserved.status:type_name -> proto.StateObserved.Status
	1, // 1: proto.StateExpected.state:type_name -> proto.StateExpected.State
	2, // 2: proto.ActionResponse.status:type_name -> proto.ActionResponse.Status
	4, // 3: proto.ElasticAgent.Checkin:input_type -> proto.StateObserved
	7, // 4: proto.ElasticAgent.Actions:input_type -> proto.ActionResponse
	5, // 5: proto.ElasticAgent.Checkin:output_type -> proto.StateExpected
	6, // 6: proto.ElasticAgent.Actions:output_type -> proto.ActionRequest
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_elastic_agent_client_proto_init() }
func file_elastic_agent_client_proto_init() {
	if File_elastic_agent_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_elastic_agent_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_elastic_agent_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateObserved); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_elastic_agent_client_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateExpected); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_elastic_agent_client_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_elastic_agent_client_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_elastic_agent_client_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_elastic_agent_client_proto_goTypes,
		DependencyIndexes: file_elastic_agent_client_proto_depIdxs,
		EnumInfos:         file_elastic_agent_client_proto_enumTypes,
		MessageInfos:      file_elastic_agent_client_proto_msgTypes,
	}.Build()
	File_elastic_agent_client_proto = out.File
	file_elastic_agent_client_proto_rawDesc = nil
	file_elastic_agent_client_proto_goTypes = nil
	file_elastic_agent_client_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ElasticAgentClient is the client API for ElasticAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ElasticAgentClient interface {
	// Called by the client to provide the Elastic Agent the state of the application.
	//
	// A `StateObserved` must be streamed at least every 30 seconds or it will result in the
	// application be automatically marked as FAILED, and after 60 seconds it will be force killed and
	// restarted.
	Checkin(ctx context.Context, opts ...grpc.CallOption) (ElasticAgent_CheckinClient, error)
	// Called by the client on connection to the GRPC allowing the Elastic Agent to stream action
	// requests to the application and the application stream back responses to those requests.
	Actions(ctx context.Context, opts ...grpc.CallOption) (ElasticAgent_ActionsClient, error)
}

type elasticAgentClient struct {
	cc grpc.ClientConnInterface
}

func NewElasticAgentClient(cc grpc.ClientConnInterface) ElasticAgentClient {
	return &elasticAgentClient{cc}
}

func (c *elasticAgentClient) Checkin(ctx context.Context, opts ...grpc.CallOption) (ElasticAgent_CheckinClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ElasticAgent_serviceDesc.Streams[0], "/proto.ElasticAgent/Checkin", opts...)
	if err != nil {
		return nil, err
	}
	x := &elasticAgentCheckinClient{stream}
	return x, nil
}

type ElasticAgent_CheckinClient interface {
	Send(*StateObserved) error
	Recv() (*StateExpected, error)
	grpc.ClientStream
}

type elasticAgentCheckinClient struct {
	grpc.ClientStream
}

func (x *elasticAgentCheckinClient) Send(m *StateObserved) error {
	return x.ClientStream.SendMsg(m)
}

func (x *elasticAgentCheckinClient) Recv() (*StateExpected, error) {
	m := new(StateExpected)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *elasticAgentClient) Actions(ctx context.Context, opts ...grpc.CallOption) (ElasticAgent_ActionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ElasticAgent_serviceDesc.Streams[1], "/proto.ElasticAgent/Actions", opts...)
	if err != nil {
		return nil, err
	}
	x := &elasticAgentActionsClient{stream}
	return x, nil
}

type ElasticAgent_ActionsClient interface {
	Send(*ActionResponse) error
	Recv() (*ActionRequest, error)
	grpc.ClientStream
}

type elasticAgentActionsClient struct {
	grpc.ClientStream
}

func (x *elasticAgentActionsClient) Send(m *ActionResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *elasticAgentActionsClient) Recv() (*ActionRequest, error) {
	m := new(ActionRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ElasticAgentServer is the server API for ElasticAgent service.
type ElasticAgentServer interface {
	// Called by the client to provide the Elastic Agent the state of the application.
	//
	// A `StateObserved` must be streamed at least every 30 seconds or it will result in the
	// application be automatically marked as FAILED, and after 60 seconds it will be force killed and
	// restarted.
	Checkin(ElasticAgent_CheckinServer) error
	// Called by the client on connection to the GRPC allowing the Elastic Agent to stream action
	// requests to the application and the application stream back responses to those requests.
	Actions(ElasticAgent_ActionsServer) error
}

// UnimplementedElasticAgentServer can be embedded to have forward compatible implementations.
type UnimplementedElasticAgentServer struct {
}

func (*UnimplementedElasticAgentServer) Checkin(ElasticAgent_CheckinServer) error {
	return status.Errorf(codes.Unimplemented, "method Checkin not implemented")
}
func (*UnimplementedElasticAgentServer) Actions(ElasticAgent_ActionsServer) error {
	return status.Errorf(codes.Unimplemented, "method Actions not implemented")
}

func RegisterElasticAgentServer(s *grpc.Server, srv ElasticAgentServer) {
	s.RegisterService(&_ElasticAgent_serviceDesc, srv)
}

func _ElasticAgent_Checkin_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ElasticAgentServer).Checkin(&elasticAgentCheckinServer{stream})
}

type ElasticAgent_CheckinServer interface {
	Send(*StateExpected) error
	Recv() (*StateObserved, error)
	grpc.ServerStream
}

type elasticAgentCheckinServer struct {
	grpc.ServerStream
}

func (x *elasticAgentCheckinServer) Send(m *StateExpected) error {
	return x.ServerStream.SendMsg(m)
}

func (x *elasticAgentCheckinServer) Recv() (*StateObserved, error) {
	m := new(StateObserved)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ElasticAgent_Actions_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ElasticAgentServer).Actions(&elasticAgentActionsServer{stream})
}

type ElasticAgent_ActionsServer interface {
	Send(*ActionRequest) error
	Recv() (*ActionResponse, error)
	grpc.ServerStream
}

type elasticAgentActionsServer struct {
	grpc.ServerStream
}

func (x *elasticAgentActionsServer) Send(m *ActionRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *elasticAgentActionsServer) Recv() (*ActionResponse, error) {
	m := new(ActionResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ElasticAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ElasticAgent",
	HandlerType: (*ElasticAgentServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Checkin",
			Handler:       _ElasticAgent_Checkin_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Actions",
			Handler:       _ElasticAgent_Actions_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "elastic-agent-client.proto",
}
