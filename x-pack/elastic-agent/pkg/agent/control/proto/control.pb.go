// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: control.proto

package proto

import (
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

// Status codes for the current state.
type Status int32

const (
	Status_STARTING    Status = 0
	Status_CONFIGURING Status = 1
	Status_HEALTHY     Status = 2
	Status_DEGRADED    Status = 3
	Status_FAILED      Status = 4
	Status_STOPPING    Status = 5
	Status_UPGRADING   Status = 6
	Status_ROLLBACK    Status = 7
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "STARTING",
		1: "CONFIGURING",
		2: "HEALTHY",
		3: "DEGRADED",
		4: "FAILED",
		5: "STOPPING",
		6: "UPGRADING",
		7: "ROLLBACK",
	}
	Status_value = map[string]int32{
		"STARTING":    0,
		"CONFIGURING": 1,
		"HEALTHY":     2,
		"DEGRADED":    3,
		"FAILED":      4,
		"STOPPING":    5,
		"UPGRADING":   6,
		"ROLLBACK":    7,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_control_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_control_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{0}
}

// Action status codes for restart and upgrade response.
type ActionStatus int32

const (
	// Action was successful.
	ActionStatus_SUCCESS ActionStatus = 0
	// Action failed.
	ActionStatus_FAILURE ActionStatus = 1
)

// Enum value maps for ActionStatus.
var (
	ActionStatus_name = map[int32]string{
		0: "SUCCESS",
		1: "FAILURE",
	}
	ActionStatus_value = map[string]int32{
		"SUCCESS": 0,
		"FAILURE": 1,
	}
)

func (x ActionStatus) Enum() *ActionStatus {
	p := new(ActionStatus)
	*p = x
	return p
}

func (x ActionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_control_proto_enumTypes[1].Descriptor()
}

func (ActionStatus) Type() protoreflect.EnumType {
	return &file_control_proto_enumTypes[1]
}

func (x ActionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActionStatus.Descriptor instead.
func (ActionStatus) EnumDescriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{1}
}

// Empty message.
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{0}
}

// Version response message.
type VersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Current running version.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// Current running commit.
	Commit string `protobuf:"bytes,2,opt,name=commit,proto3" json:"commit,omitempty"`
	// Current running build time.
	BuildTime string `protobuf:"bytes,3,opt,name=buildTime,proto3" json:"buildTime,omitempty"`
	// Current running version is a snapshot.
	Snapshot bool `protobuf:"varint,4,opt,name=snapshot,proto3" json:"snapshot,omitempty"`
}

func (x *VersionResponse) Reset() {
	*x = VersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionResponse) ProtoMessage() {}

func (x *VersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionResponse.ProtoReflect.Descriptor instead.
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{1}
}

func (x *VersionResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *VersionResponse) GetCommit() string {
	if x != nil {
		return x.Commit
	}
	return ""
}

func (x *VersionResponse) GetBuildTime() string {
	if x != nil {
		return x.BuildTime
	}
	return ""
}

func (x *VersionResponse) GetSnapshot() bool {
	if x != nil {
		return x.Snapshot
	}
	return false
}

type RestartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Response status.
	Status ActionStatus `protobuf:"varint,1,opt,name=status,proto3,enum=proto.ActionStatus" json:"status,omitempty"`
	// Error message when it fails to trigger restart.
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *RestartResponse) Reset() {
	*x = RestartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestartResponse) ProtoMessage() {}

func (x *RestartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestartResponse.ProtoReflect.Descriptor instead.
func (*RestartResponse) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{2}
}

func (x *RestartResponse) GetStatus() ActionStatus {
	if x != nil {
		return x.Status
	}
	return ActionStatus_SUCCESS
}

func (x *RestartResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

// Upgrade request message.
type UpgradeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Optional) Version to upgrade to.
	//
	// If not provided Elastic Agent will auto discover the latest version in the same major
	// to upgrade to. If wanting to upgrade to a new major that major must be present in the
	// this version field.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// (Optional) Use a different source URI then configured.
	//
	// If provided the upgrade process will use the provided sourceURI instead of the configured
	// sourceURI in the configuration.
	SourceURI string `protobuf:"bytes,2,opt,name=sourceURI,proto3" json:"sourceURI,omitempty"`
	// (Optional) Overrides predefined behavior for agent package verification.
	//
	// If provided Elastic Agent package is not checked for signature during upgrade.
	SkipVerify bool `protobuf:"varint,3,opt,name=skipVerify,proto3" json:"skipVerify,omitempty"`
	// (Optional) Overrides predefined behavior for agent package verification.
	//
	// If provided Elastic Agent package is checked against these pgp keys as well.
	PgpBytes []string `protobuf:"bytes,4,rep,name=pgpBytes,proto3" json:"pgpBytes,omitempty"`
}

func (x *UpgradeRequest) Reset() {
	*x = UpgradeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpgradeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpgradeRequest) ProtoMessage() {}

func (x *UpgradeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpgradeRequest.ProtoReflect.Descriptor instead.
func (*UpgradeRequest) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{3}
}

func (x *UpgradeRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *UpgradeRequest) GetSourceURI() string {
	if x != nil {
		return x.SourceURI
	}
	return ""
}

func (x *UpgradeRequest) GetSkipVerify() bool {
	if x != nil {
		return x.SkipVerify
	}
	return false
}

func (x *UpgradeRequest) GetPgpBytes() []string {
	if x != nil {
		return x.PgpBytes
	}
	return nil
}

// A upgrade response message.
type UpgradeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Response status.
	Status ActionStatus `protobuf:"varint,1,opt,name=status,proto3,enum=proto.ActionStatus" json:"status,omitempty"`
	// Version that is being upgraded to.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// Error message when it fails to trigger upgrade.
	Error string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *UpgradeResponse) Reset() {
	*x = UpgradeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpgradeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpgradeResponse) ProtoMessage() {}

func (x *UpgradeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpgradeResponse.ProtoReflect.Descriptor instead.
func (*UpgradeResponse) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{4}
}

func (x *UpgradeResponse) GetStatus() ActionStatus {
	if x != nil {
		return x.Status
	}
	return ActionStatus_SUCCESS
}

func (x *UpgradeResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *UpgradeResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

// Current status of the application in Elastic Agent.
type ApplicationStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique application ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Application name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Current status.
	Status Status `protobuf:"varint,3,opt,name=status,proto3,enum=proto.Status" json:"status,omitempty"`
	// Current status message.
	Message string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	// Current status payload.
	Payload string `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *ApplicationStatus) Reset() {
	*x = ApplicationStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationStatus) ProtoMessage() {}

func (x *ApplicationStatus) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationStatus.ProtoReflect.Descriptor instead.
func (*ApplicationStatus) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{5}
}

func (x *ApplicationStatus) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ApplicationStatus) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ApplicationStatus) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_STARTING
}

func (x *ApplicationStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ApplicationStatus) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

// Current metadata for a running process.
type ProcMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Process         string `protobuf:"bytes,1,opt,name=process,proto3" json:"process,omitempty"`
	Name            string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Hostname        string `protobuf:"bytes,3,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Id              string `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	EphemeralId     string `protobuf:"bytes,5,opt,name=ephemeral_id,json=ephemeralId,proto3" json:"ephemeral_id,omitempty"`
	Version         string `protobuf:"bytes,6,opt,name=version,proto3" json:"version,omitempty"`
	BuildCommit     string `protobuf:"bytes,7,opt,name=build_commit,json=buildCommit,proto3" json:"build_commit,omitempty"`
	BuildTime       string `protobuf:"bytes,8,opt,name=build_time,json=buildTime,proto3" json:"build_time,omitempty"`
	Username        string `protobuf:"bytes,9,opt,name=username,proto3" json:"username,omitempty"`
	UserId          string `protobuf:"bytes,10,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserGid         string `protobuf:"bytes,11,opt,name=user_gid,json=userGid,proto3" json:"user_gid,omitempty"`
	Architecture    string `protobuf:"bytes,12,opt,name=architecture,proto3" json:"architecture,omitempty"`
	RouteKey        string `protobuf:"bytes,13,opt,name=route_key,json=routeKey,proto3" json:"route_key,omitempty"`
	ElasticLicensed bool   `protobuf:"varint,14,opt,name=elastic_licensed,json=elasticLicensed,proto3" json:"elastic_licensed,omitempty"`
	Error           string `protobuf:"bytes,15,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ProcMeta) Reset() {
	*x = ProcMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcMeta) ProtoMessage() {}

func (x *ProcMeta) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcMeta.ProtoReflect.Descriptor instead.
func (*ProcMeta) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{6}
}

func (x *ProcMeta) GetProcess() string {
	if x != nil {
		return x.Process
	}
	return ""
}

func (x *ProcMeta) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProcMeta) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *ProcMeta) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProcMeta) GetEphemeralId() string {
	if x != nil {
		return x.EphemeralId
	}
	return ""
}

func (x *ProcMeta) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ProcMeta) GetBuildCommit() string {
	if x != nil {
		return x.BuildCommit
	}
	return ""
}

func (x *ProcMeta) GetBuildTime() string {
	if x != nil {
		return x.BuildTime
	}
	return ""
}

func (x *ProcMeta) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ProcMeta) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ProcMeta) GetUserGid() string {
	if x != nil {
		return x.UserGid
	}
	return ""
}

func (x *ProcMeta) GetArchitecture() string {
	if x != nil {
		return x.Architecture
	}
	return ""
}

func (x *ProcMeta) GetRouteKey() string {
	if x != nil {
		return x.RouteKey
	}
	return ""
}

func (x *ProcMeta) GetElasticLicensed() bool {
	if x != nil {
		return x.ElasticLicensed
	}
	return false
}

func (x *ProcMeta) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

// Status is the current status of Elastic Agent.
type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Overall status of Elastic Agent.
	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=proto.Status" json:"status,omitempty"`
	// Overall status message of Elastic Agent.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// Status of each application in Elastic Agent.
	Applications []*ApplicationStatus `protobuf:"bytes,3,rep,name=applications,proto3" json:"applications,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{7}
}

func (x *StatusResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_STARTING
}

func (x *StatusResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *StatusResponse) GetApplications() []*ApplicationStatus {
	if x != nil {
		return x.Applications
	}
	return nil
}

// ProcMetaResponse is the current running version infomation for all processes.
type ProcMetaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Procs []*ProcMeta `protobuf:"bytes,1,rep,name=procs,proto3" json:"procs,omitempty"`
}

func (x *ProcMetaResponse) Reset() {
	*x = ProcMetaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcMetaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcMetaResponse) ProtoMessage() {}

func (x *ProcMetaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcMetaResponse.ProtoReflect.Descriptor instead.
func (*ProcMetaResponse) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{8}
}

func (x *ProcMetaResponse) GetProcs() []*ProcMeta {
	if x != nil {
		return x.Procs
	}
	return nil
}

var File_control_proto protoreflect.FileDescriptor

var file_control_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x7d, 0x0a, 0x0f, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x22, 0x54,
	0x0a, 0x0f, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x22, 0x84, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x52, 0x49, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x52, 0x49, 0x12,
	0x1e, 0x0a, 0x0a, 0x73, 0x6b, 0x69, 0x70, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x6b, 0x69, 0x70, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x67, 0x70, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x67, 0x70, 0x42, 0x79, 0x74, 0x65, 0x73, 0x22, 0x6e, 0x0a, 0x0f, 0x55,
	0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x92, 0x01, 0x0a, 0x11,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x22, 0xb5, 0x03, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68,
	0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68,
	0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x70, 0x68, 0x65, 0x6d,
	0x65, 0x72, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65,
	0x70, 0x68, 0x65, 0x6d, 0x65, 0x72, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x63, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x67, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x47, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74,
	0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x72,
	0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x29, 0x0a, 0x10, 0x65, 0x6c, 0x61, 0x73, 0x74,
	0x69, 0x63, 0x5f, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0f, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73,
	0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x8f, 0x01, 0x0a, 0x0e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3c, 0x0a, 0x0c,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0c, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x39, 0x0a, 0x10, 0x50, 0x72,
	0x6f, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25,
	0x0a, 0x05, 0x70, 0x72, 0x6f, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x05,
	0x70, 0x72, 0x6f, 0x63, 0x73, 0x2a, 0x79, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x0c, 0x0a, 0x08, 0x53, 0x54, 0x41, 0x52, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0f, 0x0a,
	0x0b, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x55, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0b,
	0x0a, 0x07, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x59, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x44,
	0x45, 0x47, 0x52, 0x41, 0x44, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49,
	0x4c, 0x45, 0x44, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x54, 0x4f, 0x50, 0x50, 0x49, 0x4e,
	0x47, 0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x50, 0x47, 0x52, 0x41, 0x44, 0x49, 0x4e, 0x47,
	0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x4f, 0x4c, 0x4c, 0x42, 0x41, 0x43, 0x4b, 0x10, 0x07,
	0x2a, 0x28, 0x0a, 0x0c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x01, 0x32, 0x93, 0x02, 0x0a, 0x13, 0x45,
	0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x12, 0x2f, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x0c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x12, 0x15,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70,
	0x67, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a,
	0x08, 0x50, 0x72, 0x6f, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x72, 0x6f, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x22, 0x5a, 0x1d, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_control_proto_rawDescOnce sync.Once
	file_control_proto_rawDescData = file_control_proto_rawDesc
)

func file_control_proto_rawDescGZIP() []byte {
	file_control_proto_rawDescOnce.Do(func() {
		file_control_proto_rawDescData = protoimpl.X.CompressGZIP(file_control_proto_rawDescData)
	})
	return file_control_proto_rawDescData
}

var file_control_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_control_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_control_proto_goTypes = []interface{}{
	(Status)(0),               // 0: proto.Status
	(ActionStatus)(0),         // 1: proto.ActionStatus
	(*Empty)(nil),             // 2: proto.Empty
	(*VersionResponse)(nil),   // 3: proto.VersionResponse
	(*RestartResponse)(nil),   // 4: proto.RestartResponse
	(*UpgradeRequest)(nil),    // 5: proto.UpgradeRequest
	(*UpgradeResponse)(nil),   // 6: proto.UpgradeResponse
	(*ApplicationStatus)(nil), // 7: proto.ApplicationStatus
	(*ProcMeta)(nil),          // 8: proto.ProcMeta
	(*StatusResponse)(nil),    // 9: proto.StatusResponse
	(*ProcMetaResponse)(nil),  // 10: proto.ProcMetaResponse
}
var file_control_proto_depIdxs = []int32{
	1,  // 0: proto.RestartResponse.status:type_name -> proto.ActionStatus
	1,  // 1: proto.UpgradeResponse.status:type_name -> proto.ActionStatus
	0,  // 2: proto.ApplicationStatus.status:type_name -> proto.Status
	0,  // 3: proto.StatusResponse.status:type_name -> proto.Status
	7,  // 4: proto.StatusResponse.applications:type_name -> proto.ApplicationStatus
	8,  // 5: proto.ProcMetaResponse.procs:type_name -> proto.ProcMeta
	2,  // 6: proto.ElasticAgentControl.Version:input_type -> proto.Empty
	2,  // 7: proto.ElasticAgentControl.Status:input_type -> proto.Empty
	2,  // 8: proto.ElasticAgentControl.Restart:input_type -> proto.Empty
	5,  // 9: proto.ElasticAgentControl.Upgrade:input_type -> proto.UpgradeRequest
	2,  // 10: proto.ElasticAgentControl.ProcMeta:input_type -> proto.Empty
	3,  // 11: proto.ElasticAgentControl.Version:output_type -> proto.VersionResponse
	9,  // 12: proto.ElasticAgentControl.Status:output_type -> proto.StatusResponse
	4,  // 13: proto.ElasticAgentControl.Restart:output_type -> proto.RestartResponse
	6,  // 14: proto.ElasticAgentControl.Upgrade:output_type -> proto.UpgradeResponse
	10, // 15: proto.ElasticAgentControl.ProcMeta:output_type -> proto.ProcMetaResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_control_proto_init() }
func file_control_proto_init() {
	if File_control_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_control_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_control_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionResponse); i {
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
		file_control_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestartResponse); i {
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
		file_control_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpgradeRequest); i {
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
		file_control_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpgradeResponse); i {
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
		file_control_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationStatus); i {
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
		file_control_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcMeta); i {
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
		file_control_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResponse); i {
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
		file_control_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcMetaResponse); i {
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
			RawDescriptor: file_control_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_control_proto_goTypes,
		DependencyIndexes: file_control_proto_depIdxs,
		EnumInfos:         file_control_proto_enumTypes,
		MessageInfos:      file_control_proto_msgTypes,
	}.Build()
	File_control_proto = out.File
	file_control_proto_rawDesc = nil
	file_control_proto_goTypes = nil
	file_control_proto_depIdxs = nil
}
