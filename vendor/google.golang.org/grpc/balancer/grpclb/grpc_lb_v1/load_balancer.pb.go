// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/lb/v1/load_balancer.proto

package grpc_lb_v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type LoadBalanceRequest struct {
	// Types that are valid to be assigned to LoadBalanceRequestType:
	//	*LoadBalanceRequest_InitialRequest
	//	*LoadBalanceRequest_ClientStats
	LoadBalanceRequestType isLoadBalanceRequest_LoadBalanceRequestType `protobuf_oneof:"load_balance_request_type"`
	XXX_NoUnkeyedLiteral   struct{}                                    `json:"-"`
	XXX_unrecognized       []byte                                      `json:"-"`
	XXX_sizecache          int32                                       `json:"-"`
}

func (m *LoadBalanceRequest) Reset()         { *m = LoadBalanceRequest{} }
func (m *LoadBalanceRequest) String() string { return proto.CompactTextString(m) }
func (*LoadBalanceRequest) ProtoMessage()    {}
func (*LoadBalanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cd3f6d792743fdf, []int{0}
}

func (m *LoadBalanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadBalanceRequest.Unmarshal(m, b)
}
func (m *LoadBalanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadBalanceRequest.Marshal(b, m, deterministic)
}
func (m *LoadBalanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadBalanceRequest.Merge(m, src)
}
func (m *LoadBalanceRequest) XXX_Size() int {
	return xxx_messageInfo_LoadBalanceRequest.Size(m)
}
func (m *LoadBalanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadBalanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoadBalanceRequest proto.InternalMessageInfo

type isLoadBalanceRequest_LoadBalanceRequestType interface {
	isLoadBalanceRequest_LoadBalanceRequestType()
}

type LoadBalanceRequest_InitialRequest struct {
	InitialRequest *InitialLoadBalanceRequest `protobuf:"bytes,1,opt,name=initial_request,json=initialRequest,proto3,oneof"`
}

type LoadBalanceRequest_ClientStats struct {
	ClientStats *ClientStats `protobuf:"bytes,2,opt,name=client_stats,json=clientStats,proto3,oneof"`
}

func (*LoadBalanceRequest_InitialRequest) isLoadBalanceRequest_LoadBalanceRequestType() {}

func (*LoadBalanceRequest_ClientStats) isLoadBalanceRequest_LoadBalanceRequestType() {}

func (m *LoadBalanceRequest) GetLoadBalanceRequestType() isLoadBalanceRequest_LoadBalanceRequestType {
	if m != nil {
		return m.LoadBalanceRequestType
	}
	return nil
}

func (m *LoadBalanceRequest) GetInitialRequest() *InitialLoadBalanceRequest {
	if x, ok := m.GetLoadBalanceRequestType().(*LoadBalanceRequest_InitialRequest); ok {
		return x.InitialRequest
	}
	return nil
}

func (m *LoadBalanceRequest) GetClientStats() *ClientStats {
	if x, ok := m.GetLoadBalanceRequestType().(*LoadBalanceRequest_ClientStats); ok {
		return x.ClientStats
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*LoadBalanceRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*LoadBalanceRequest_InitialRequest)(nil),
		(*LoadBalanceRequest_ClientStats)(nil),
	}
}

type InitialLoadBalanceRequest struct {
	// The name of the load balanced service (e.g., service.googleapis.com). Its
	// length should be less than 256 bytes.
	// The name might include a port number. How to handle the port number is up
	// to the balancer.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InitialLoadBalanceRequest) Reset()         { *m = InitialLoadBalanceRequest{} }
func (m *InitialLoadBalanceRequest) String() string { return proto.CompactTextString(m) }
func (*InitialLoadBalanceRequest) ProtoMessage()    {}
func (*InitialLoadBalanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cd3f6d792743fdf, []int{1}
}

func (m *InitialLoadBalanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitialLoadBalanceRequest.Unmarshal(m, b)
}
func (m *InitialLoadBalanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitialLoadBalanceRequest.Marshal(b, m, deterministic)
}
func (m *InitialLoadBalanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitialLoadBalanceRequest.Merge(m, src)
}
func (m *InitialLoadBalanceRequest) XXX_Size() int {
	return xxx_messageInfo_InitialLoadBalanceRequest.Size(m)
}
func (m *InitialLoadBalanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InitialLoadBalanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InitialLoadBalanceRequest proto.InternalMessageInfo

func (m *InitialLoadBalanceRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Contains the number of calls finished for a particular load balance token.
type ClientStatsPerToken struct {
	// See Server.load_balance_token.
	LoadBalanceToken string `protobuf:"bytes,1,opt,name=load_balance_token,json=loadBalanceToken,proto3" json:"load_balance_token,omitempty"`
	// The total number of RPCs that finished associated with the token.
	NumCalls             int64    `protobuf:"varint,2,opt,name=num_calls,json=numCalls,proto3" json:"num_calls,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientStatsPerToken) Reset()         { *m = ClientStatsPerToken{} }
func (m *ClientStatsPerToken) String() string { return proto.CompactTextString(m) }
func (*ClientStatsPerToken) ProtoMessage()    {}
func (*ClientStatsPerToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cd3f6d792743fdf, []int{2}
}

func (m *ClientStatsPerToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientStatsPerToken.Unmarshal(m, b)
}
func (m *ClientStatsPerToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientStatsPerToken.Marshal(b, m, deterministic)
}
func (m *ClientStatsPerToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientStatsPerToken.Merge(m, src)
}
func (m *ClientStatsPerToken) XXX_Size() int {
	return xxx_messageInfo_ClientStatsPerToken.Size(m)
}
func (m *ClientStatsPerToken) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientStatsPerToken.DiscardUnknown(m)
}

var xxx_messageInfo_ClientStatsPerToken proto.InternalMessageInfo

func (m *ClientStatsPerToken) GetLoadBalanceToken() string {
	if m != nil {
		return m.LoadBalanceToken
	}
	return ""
}

func (m *ClientStatsPerToken) GetNumCalls() int64 {
	if m != nil {
		return m.NumCalls
	}
	return 0
}

// Contains client level statistics that are useful to load balancing. Each
// count except the timestamp should be reset to zero after reporting the stats.
type ClientStats struct {
	// The timestamp of generating the report.
	Timestamp *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// The total number of RPCs that started.
	NumCallsStarted int64 `protobuf:"varint,2,opt,name=num_calls_started,json=numCallsStarted,proto3" json:"num_calls_started,omitempty"`
	// The total number of RPCs that finished.
	NumCallsFinished int64 `protobuf:"varint,3,opt,name=num_calls_finished,json=numCallsFinished,proto3" json:"num_calls_finished,omitempty"`
	// The total number of RPCs that failed to reach a server except dropped RPCs.
	NumCallsFinishedWithClientFailedToSend int64 `protobuf:"varint,6,opt,name=num_calls_finished_with_client_failed_to_send,json=numCallsFinishedWithClientFailedToSend,proto3" json:"num_calls_finished_with_client_failed_to_send,omitempty"`
	// The total number of RPCs that finished and are known to have been received
	// by a server.
	NumCallsFinishedKnownReceived int64 `protobuf:"varint,7,opt,name=num_calls_finished_known_received,json=numCallsFinishedKnownReceived,proto3" json:"num_calls_finished_known_received,omitempty"`
	// The list of dropped calls.
	CallsFinishedWithDrop []*ClientStatsPerToken `protobuf:"bytes,8,rep,name=calls_finished_with_drop,json=callsFinishedWithDrop,proto3" json:"calls_finished_with_drop,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}               `json:"-"`
	XXX_unrecognized      []byte                 `json:"-"`
	XXX_sizecache         int32                  `json:"-"`
}

func (m *ClientStats) Reset()         { *m = ClientStats{} }
func (m *ClientStats) String() string { return proto.CompactTextString(m) }
func (*ClientStats) ProtoMessage()    {}
func (*ClientStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cd3f6d792743fdf, []int{3}
}

func (m *ClientStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientStats.Unmarshal(m, b)
}
func (m *ClientStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientStats.Marshal(b, m, deterministic)
}
func (m *ClientStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientStats.Merge(m, src)
}
func (m *ClientStats) XXX_Size() int {
	return xxx_messageInfo_ClientStats.Size(m)
}
func (m *ClientStats) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientStats.DiscardUnknown(m)
}

var xxx_messageInfo_ClientStats proto.InternalMessageInfo

func (m *ClientStats) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *ClientStats) GetNumCallsStarted() int64 {
	if m != nil {
		return m.NumCallsStarted
	}
	return 0
}

func (m *ClientStats) GetNumCallsFinished() int64 {
	if m != nil {
		return m.NumCallsFinished
	}
	return 0
}

func (m *ClientStats) GetNumCallsFinishedWithClientFailedToSend() int64 {
	if m != nil {
		return m.NumCallsFinishedWithClientFailedToSend
	}
	return 0
}

func (m *ClientStats) GetNumCallsFinishedKnownReceived() int64 {
	if m != nil {
		return m.NumCallsFinishedKnownReceived
	}
	return 0
}

func (m *ClientStats) GetCallsFinishedWithDrop() []*ClientStatsPerToken {
	if m != nil {
		return m.CallsFinishedWithDrop
	}
	return nil
}

type LoadBalanceResponse struct {
	// Types that are valid to be assigned to LoadBalanceResponseType:
	//	*LoadBalanceResponse_InitialResponse
	//	*LoadBalanceResponse_ServerList
	//	*LoadBalanceResponse_FallbackResponse
	LoadBalanceResponseType isLoadBalanceResponse_LoadBalanceResponseType `protobuf_oneof:"load_balance_response_type"`
	XXX_NoUnkeyedLiteral    struct{}                                      `json:"-"`
	XXX_unrecognized        []byte                                        `json:"-"`
	XXX_sizecache           int32                                         `json:"-"`
}

func (m *LoadBalanceResponse) Reset()         { *m = LoadBalanceResponse{} }
func (m *LoadBalanceResponse) String() string { return proto.CompactTextString(m) }
func (*LoadBalanceResponse) ProtoMessage()    {}
func (*LoadBalanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cd3f6d792743fdf, []int{4}
}

func (m *LoadBalanceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadBalanceResponse.Unmarshal(m, b)
}
func (m *LoadBalanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadBalanceResponse.Marshal(b, m, deterministic)
}
func (m *LoadBalanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadBalanceResponse.Merge(m, src)
}
func (m *LoadBalanceResponse) XXX_Size() int {
	return xxx_messageInfo_LoadBalanceResponse.Size(m)
}
func (m *LoadBalanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadBalanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoadBalanceResponse proto.InternalMessageInfo

type isLoadBalanceResponse_LoadBalanceResponseType interface {
	isLoadBalanceResponse_LoadBalanceResponseType()
}

type LoadBalanceResponse_InitialResponse struct {
	InitialResponse *InitialLoadBalanceResponse `protobuf:"bytes,1,opt,name=initial_response,json=initialResponse,proto3,oneof"`
}

type LoadBalanceResponse_ServerList struct {
	ServerList *ServerList `protobuf:"bytes,2,opt,name=server_list,json=serverList,proto3,oneof"`
}

type LoadBalanceResponse_FallbackResponse struct {
	FallbackResponse *FallbackResponse `protobuf:"bytes,3,opt,name=fallback_response,json=fallbackResponse,proto3,oneof"`
}

func (*LoadBalanceResponse_InitialResponse) isLoadBalanceResponse_LoadBalanceResponseType() {}

func (*LoadBalanceResponse_ServerList) isLoadBalanceResponse_LoadBalanceResponseType() {}

func (*LoadBalanceResponse_FallbackResponse) isLoadBalanceResponse_LoadBalanceResponseType() {}

func (m *LoadBalanceResponse) GetLoadBalanceResponseType() isLoadBalanceResponse_LoadBalanceResponseType {
	if m != nil {
		return m.LoadBalanceResponseType
	}
	return nil
}

func (m *LoadBalanceResponse) GetInitialResponse() *InitialLoadBalanceResponse {
	if x, ok := m.GetLoadBalanceResponseType().(*LoadBalanceResponse_InitialResponse); ok {
		return x.InitialResponse
	}
	return nil
}

func (m *LoadBalanceResponse) GetServerList() *ServerList {
	if x, ok := m.GetLoadBalanceResponseType().(*LoadBalanceResponse_ServerList); ok {
		return x.ServerList
	}
	return nil
}

func (m *LoadBalanceResponse) GetFallbackResponse() *FallbackResponse {
	if x, ok := m.GetLoadBalanceResponseType().(*LoadBalanceResponse_FallbackResponse); ok {
		return x.FallbackResponse
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*LoadBalanceResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*LoadBalanceResponse_InitialResponse)(nil),
		(*LoadBalanceResponse_ServerList)(nil),
		(*LoadBalanceResponse_FallbackResponse)(nil),
	}
}

type FallbackResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FallbackResponse) Reset()         { *m = FallbackResponse{} }
func (m *FallbackResponse) String() string { return proto.CompactTextString(m) }
func (*FallbackResponse) ProtoMessage()    {}
func (*FallbackResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cd3f6d792743fdf, []int{5}
}

func (m *FallbackResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FallbackResponse.Unmarshal(m, b)
}
func (m *FallbackResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FallbackResponse.Marshal(b, m, deterministic)
}
func (m *FallbackResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FallbackResponse.Merge(m, src)
}
func (m *FallbackResponse) XXX_Size() int {
	return xxx_messageInfo_FallbackResponse.Size(m)
}
func (m *FallbackResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FallbackResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FallbackResponse proto.InternalMessageInfo

type InitialLoadBalanceResponse struct {
	// This is an application layer redirect that indicates the client should use
	// the specified server for load balancing. When this field is non-empty in
	// the response, the client should open a separate connection to the
	// load_balancer_delegate and call the BalanceLoad method. Its length should
	// be less than 64 bytes.
	LoadBalancerDelegate string `protobuf:"bytes,1,opt,name=load_balancer_delegate,json=loadBalancerDelegate,proto3" json:"load_balancer_delegate,omitempty"`
	// This interval defines how often the client should send the client stats
	// to the load balancer. Stats should only be reported when the duration is
	// positive.
	ClientStatsReportInterval *duration.Duration `protobuf:"bytes,2,opt,name=client_stats_report_interval,json=clientStatsReportInterval,proto3" json:"client_stats_report_interval,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}           `json:"-"`
	XXX_unrecognized          []byte             `json:"-"`
	XXX_sizecache             int32              `json:"-"`
}

func (m *InitialLoadBalanceResponse) Reset()         { *m = InitialLoadBalanceResponse{} }
func (m *InitialLoadBalanceResponse) String() string { return proto.CompactTextString(m) }
func (*InitialLoadBalanceResponse) ProtoMessage()    {}
func (*InitialLoadBalanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cd3f6d792743fdf, []int{6}
}

func (m *InitialLoadBalanceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitialLoadBalanceResponse.Unmarshal(m, b)
}
func (m *InitialLoadBalanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitialLoadBalanceResponse.Marshal(b, m, deterministic)
}
func (m *InitialLoadBalanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitialLoadBalanceResponse.Merge(m, src)
}
func (m *InitialLoadBalanceResponse) XXX_Size() int {
	return xxx_messageInfo_InitialLoadBalanceResponse.Size(m)
}
func (m *InitialLoadBalanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InitialLoadBalanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InitialLoadBalanceResponse proto.InternalMessageInfo

func (m *InitialLoadBalanceResponse) GetLoadBalancerDelegate() string {
	if m != nil {
		return m.LoadBalancerDelegate
	}
	return ""
}

func (m *InitialLoadBalanceResponse) GetClientStatsReportInterval() *duration.Duration {
	if m != nil {
		return m.ClientStatsReportInterval
	}
	return nil
}

type ServerList struct {
	// Contains a list of servers selected by the load balancer. The list will
	// be updated when server resolutions change or as needed to balance load
	// across more servers. The client should consume the server list in order
	// unless instructed otherwise via the client_config.
	Servers              []*Server `protobuf:"bytes,1,rep,name=servers,proto3" json:"servers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ServerList) Reset()         { *m = ServerList{} }
func (m *ServerList) String() string { return proto.CompactTextString(m) }
func (*ServerList) ProtoMessage()    {}
func (*ServerList) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cd3f6d792743fdf, []int{7}
}

func (m *ServerList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerList.Unmarshal(m, b)
}
func (m *ServerList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerList.Marshal(b, m, deterministic)
}
func (m *ServerList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerList.Merge(m, src)
}
func (m *ServerList) XXX_Size() int {
	return xxx_messageInfo_ServerList.Size(m)
}
func (m *ServerList) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerList.DiscardUnknown(m)
}

var xxx_messageInfo_ServerList proto.InternalMessageInfo

func (m *ServerList) GetServers() []*Server {
	if m != nil {
		return m.Servers
	}
	return nil
}

// Contains server information. When the drop field is not true, use the other
// fields.
type Server struct {
	// A resolved address for the server, serialized in network-byte-order. It may
	// either be an IPv4 or IPv6 address.
	IpAddress []byte `protobuf:"bytes,1,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	// A resolved port number for the server.
	Port int32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	// An opaque but printable token for load reporting. The client must include
	// the token of the picked server into the initial metadata when it starts a
	// call to that server. The token is used by the server to verify the request
	// and to allow the server to report load to the gRPC LB system. The token is
	// also used in client stats for reporting dropped calls.
	//
	// Its length can be variable but must be less than 50 bytes.
	LoadBalanceToken string `protobuf:"bytes,3,opt,name=load_balance_token,json=loadBalanceToken,proto3" json:"load_balance_token,omitempty"`
	// Indicates whether this particular request should be dropped by the client.
	// If the request is dropped, there will be a corresponding entry in
	// ClientStats.calls_finished_with_drop.
	Drop                 bool     `protobuf:"varint,4,opt,name=drop,proto3" json:"drop,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Server) Reset()         { *m = Server{} }
func (m *Server) String() string { return proto.CompactTextString(m) }
func (*Server) ProtoMessage()    {}
func (*Server) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cd3f6d792743fdf, []int{8}
}

func (m *Server) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Server.Unmarshal(m, b)
}
func (m *Server) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Server.Marshal(b, m, deterministic)
}
func (m *Server) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Server.Merge(m, src)
}
func (m *Server) XXX_Size() int {
	return xxx_messageInfo_Server.Size(m)
}
func (m *Server) XXX_DiscardUnknown() {
	xxx_messageInfo_Server.DiscardUnknown(m)
}

var xxx_messageInfo_Server proto.InternalMessageInfo

func (m *Server) GetIpAddress() []byte {
	if m != nil {
		return m.IpAddress
	}
	return nil
}

func (m *Server) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Server) GetLoadBalanceToken() string {
	if m != nil {
		return m.LoadBalanceToken
	}
	return ""
}

func (m *Server) GetDrop() bool {
	if m != nil {
		return m.Drop
	}
	return false
}

func init() {
	proto.RegisterType((*LoadBalanceRequest)(nil), "grpc.lb.v1.LoadBalanceRequest")
	proto.RegisterType((*InitialLoadBalanceRequest)(nil), "grpc.lb.v1.InitialLoadBalanceRequest")
	proto.RegisterType((*ClientStatsPerToken)(nil), "grpc.lb.v1.ClientStatsPerToken")
	proto.RegisterType((*ClientStats)(nil), "grpc.lb.v1.ClientStats")
	proto.RegisterType((*LoadBalanceResponse)(nil), "grpc.lb.v1.LoadBalanceResponse")
	proto.RegisterType((*FallbackResponse)(nil), "grpc.lb.v1.FallbackResponse")
	proto.RegisterType((*InitialLoadBalanceResponse)(nil), "grpc.lb.v1.InitialLoadBalanceResponse")
	proto.RegisterType((*ServerList)(nil), "grpc.lb.v1.ServerList")
	proto.RegisterType((*Server)(nil), "grpc.lb.v1.Server")
}

func init() { proto.RegisterFile("grpc/lb/v1/load_balancer.proto", fileDescriptor_7cd3f6d792743fdf) }

var fileDescriptor_7cd3f6d792743fdf = []byte{
	// 785 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0x8e, 0x6a, 0x27, 0x75, 0x8e, 0xb3, 0xc5, 0x61, 0xb7, 0x4e, 0x71, 0xd3, 0x24, 0x13, 0xb0,
	0x22, 0x18, 0x3a, 0x79, 0xc9, 0x76, 0xb1, 0x01, 0xbb, 0xd8, 0xdc, 0x20, 0x48, 0xd3, 0x5e, 0x04,
	0x74, 0x80, 0x0e, 0x05, 0x06, 0x8e, 0x92, 0x68, 0x87, 0x08, 0x4d, 0x6a, 0x14, 0xed, 0x62, 0xd7,
	0x7b, 0x81, 0x3d, 0xc9, 0xb0, 0x57, 0xd8, 0x9b, 0x0d, 0x22, 0x29, 0x4b, 0xb1, 0x63, 0xf4, 0x4a,
	0xe4, 0x39, 0x1f, 0xbf, 0xf3, 0x7f, 0x04, 0x87, 0x13, 0x9d, 0xa7, 0x03, 0x91, 0x0c, 0xe6, 0xa7,
	0x03, 0xa1, 0x68, 0x46, 0x12, 0x2a, 0xa8, 0x4c, 0x99, 0x8e, 0x73, 0xad, 0x8c, 0x42, 0x50, 0xea,
	0x63, 0x91, 0xc4, 0xf3, 0xd3, 0xfe, 0xe1, 0x44, 0xa9, 0x89, 0x60, 0x03, 0xab, 0x49, 0x66, 0xe3,
	0x41, 0x36, 0xd3, 0xd4, 0x70, 0x25, 0x1d, 0xb6, 0x7f, 0xb4, 0xac, 0x37, 0x7c, 0xca, 0x0a, 0x43,
	0xa7, 0xb9, 0x03, 0x44, 0xff, 0x05, 0x80, 0xde, 0x2a, 0x9a, 0x0d, 0x9d, 0x0d, 0xcc, 0xfe, 0x98,
	0xb1, 0xc2, 0xa0, 0x6b, 0xd8, 0xe5, 0x92, 0x1b, 0x4e, 0x05, 0xd1, 0x4e, 0x14, 0x06, 0xc7, 0xc1,
	0x49, 0xf7, 0xec, 0xab, 0xb8, 0xb6, 0x1e, 0xbf, 0x76, 0x90, 0xd5, 0xf7, 0x97, 0x1b, 0xf8, 0x53,
	0xff, 0xbe, 0x62, 0xfc, 0x09, 0x76, 0x52, 0xc1, 0x99, 0x34, 0xa4, 0x30, 0xd4, 0x14, 0xe1, 0x23,
	0x4b, 0xf7, 0x45, 0x93, 0xee, 0x95, 0xd5, 0x8f, 0x4a, 0xf5, 0xe5, 0x06, 0xee, 0xa6, 0xf5, 0x75,
	0xf8, 0x0c, 0xf6, 0x9b, 0xa9, 0xa8, 0x9c, 0x22, 0xe6, 0xcf, 0x9c, 0x45, 0x03, 0xd8, 0x5f, 0xeb,
	0x09, 0x42, 0xd0, 0x96, 0x74, 0xca, 0xac, 0xfb, 0xdb, 0xd8, 0x9e, 0xa3, 0xdf, 0xe1, 0x49, 0xc3,
	0xd6, 0x35, 0xd3, 0x37, 0xea, 0x8e, 0x49, 0xf4, 0x12, 0xd0, 0x3d, 0x23, 0xa6, 0x94, 0xfa, 0x87,
	0x3d, 0x51, 0x53, 0x3b, 0xf4, 0x33, 0xd8, 0x96, 0xb3, 0x29, 0x49, 0xa9, 0x10, 0x2e, 0x9a, 0x16,
	0xee, 0xc8, 0xd9, 0xf4, 0x55, 0x79, 0x8f, 0xfe, 0x6d, 0x41, 0xb7, 0x61, 0x02, 0xfd, 0x00, 0xdb,
	0x8b, 0xcc, 0xfb, 0x4c, 0xf6, 0x63, 0x57, 0x9b, 0xb8, 0xaa, 0x4d, 0x7c, 0x53, 0x21, 0x70, 0x0d,
	0x46, 0x5f, 0xc3, 0xde, 0xc2, 0x4c, 0x99, 0x3a, 0x6d, 0x58, 0xe6, 0xcd, 0xed, 0x56, 0xe6, 0x46,
	0x4e, 0x5c, 0x06, 0x50, 0x63, 0xc7, 0x5c, 0xf2, 0xe2, 0x96, 0x65, 0x61, 0xcb, 0x82, 0x7b, 0x15,
	0xf8, 0xc2, 0xcb, 0xd1, 0x6f, 0xf0, 0xcd, 0x2a, 0x9a, 0x7c, 0xe0, 0xe6, 0x96, 0xf8, 0x4a, 0x8d,
	0x29, 0x17, 0x2c, 0x23, 0x46, 0x91, 0x82, 0xc9, 0x2c, 0xdc, 0xb2, 0x44, 0x2f, 0x96, 0x89, 0xde,
	0x71, 0x73, 0xeb, 0x62, 0xbd, 0xb0, 0xf8, 0x1b, 0x35, 0x62, 0x32, 0x43, 0x97, 0xf0, 0xe5, 0x03,
	0xf4, 0x77, 0x52, 0x7d, 0x90, 0x44, 0xb3, 0x94, 0xf1, 0x39, 0xcb, 0xc2, 0xc7, 0x96, 0xf2, 0xf9,
	0x32, 0xe5, 0x9b, 0x12, 0x85, 0x3d, 0x08, 0xfd, 0x0a, 0xe1, 0x43, 0x4e, 0x66, 0x5a, 0xe5, 0x61,
	0xe7, 0xb8, 0x75, 0xd2, 0x3d, 0x3b, 0x5a, 0xd3, 0x46, 0x55, 0x69, 0xf1, 0xe7, 0xe9, 0xb2, 0xc7,
	0xe7, 0x5a, 0xe5, 0x57, 0xed, 0x4e, 0xbb, 0xb7, 0x79, 0xd5, 0xee, 0x6c, 0xf6, 0xb6, 0xa2, 0xbf,
	0x1f, 0xc1, 0x93, 0x7b, 0xfd, 0x53, 0xe4, 0x4a, 0x16, 0x0c, 0x8d, 0xa0, 0x57, 0x8f, 0x82, 0x93,
	0xf9, 0x0a, 0xbe, 0xf8, 0xd8, 0x2c, 0x38, 0xf4, 0xe5, 0x06, 0xde, 0x5d, 0x0c, 0x83, 0x27, 0xfd,
	0x11, 0xba, 0x05, 0xd3, 0x73, 0xa6, 0x89, 0xe0, 0x85, 0xf1, 0xc3, 0xf0, 0xb4, 0xc9, 0x37, 0xb2,
	0xea, 0xb7, 0xdc, 0x0e, 0x13, 0x14, 0x8b, 0x1b, 0x7a, 0x03, 0x7b, 0x63, 0x2a, 0x44, 0x42, 0xd3,
	0xbb, 0xda, 0xa1, 0x96, 0x25, 0x38, 0x68, 0x12, 0x5c, 0x78, 0x50, 0xc3, 0x8d, 0xde, 0x78, 0x49,
	0x36, 0x3c, 0x80, 0xfe, 0xd2, 0x5c, 0x39, 0x85, 0x1b, 0x2c, 0x04, 0xbd, 0x65, 0x96, 0xe8, 0x9f,
	0x00, 0xfa, 0xeb, 0x63, 0x45, 0xdf, 0xc3, 0xd3, 0x7b, 0x3b, 0x8b, 0x64, 0x4c, 0xb0, 0x09, 0x35,
	0xd5, 0x00, 0x7e, 0xd6, 0x98, 0x23, 0x7d, 0xee, 0x75, 0xe8, 0x3d, 0x1c, 0x34, 0x97, 0x03, 0xd1,
	0x2c, 0x57, 0xda, 0x10, 0x2e, 0x0d, 0xd3, 0x73, 0x2a, 0x7c, 0x7e, 0xf6, 0x57, 0x26, 0xe6, 0xdc,
	0x6f, 0x3b, 0xbc, 0xdf, 0x58, 0x16, 0xd8, 0x3e, 0x7e, 0xed, 0xdf, 0x46, 0x3f, 0x03, 0xd4, 0xb9,
	0x44, 0x2f, 0xe1, 0xb1, 0xcb, 0x65, 0x11, 0x06, 0xb6, 0x75, 0xd0, 0x6a, 0xd2, 0x71, 0x05, 0xb9,
	0x6a, 0x77, 0x5a, 0xbd, 0x76, 0xf4, 0x57, 0x00, 0x5b, 0x4e, 0x83, 0x9e, 0x03, 0xf0, 0x9c, 0xd0,
	0x2c, 0xd3, 0xac, 0x28, 0x6c, 0x48, 0x3b, 0x78, 0x9b, 0xe7, 0xbf, 0x38, 0x41, 0xb9, 0x6c, 0x4a,
	0xdb, 0xd6, 0xdf, 0x4d, 0x6c, 0xcf, 0x6b, 0xb6, 0x4a, 0x6b, 0xcd, 0x56, 0x41, 0xd0, 0xb6, 0x7d,
	0xdd, 0x3e, 0x0e, 0x4e, 0x3a, 0xd8, 0x9e, 0x5d, 0x7f, 0x9e, 0x25, 0xb0, 0xd3, 0x48, 0xb8, 0x46,
	0x18, 0xba, 0xfe, 0x5c, 0x8a, 0xd1, 0x61, 0x33, 0x8e, 0xd5, 0x3d, 0xd8, 0x3f, 0x5a, 0xab, 0x77,
	0x95, 0x3b, 0x09, 0xbe, 0x0d, 0x86, 0xef, 0xe0, 0x13, 0xae, 0x1a, 0xc0, 0xe1, 0x5e, 0xd3, 0xe4,
	0x75, 0x99, 0xf6, 0xeb, 0xe0, 0xfd, 0xa9, 0x2f, 0xc3, 0x44, 0x09, 0x2a, 0x27, 0xb1, 0xd2, 0x93,
	0x81, 0xfd, 0x65, 0x55, 0x35, 0xb7, 0x37, 0x91, 0xd8, 0x0f, 0x11, 0x09, 0x99, 0x9f, 0x26, 0x5b,
	0xb6, 0x64, 0xdf, 0xfd, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x09, 0x7b, 0x39, 0x1e, 0xdc, 0x06, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LoadBalancerClient is the client API for LoadBalancer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoadBalancerClient interface {
	// Bidirectional rpc to get a list of servers.
	BalanceLoad(ctx context.Context, opts ...grpc.CallOption) (LoadBalancer_BalanceLoadClient, error)
}

type loadBalancerClient struct {
	cc *grpc.ClientConn
}

func NewLoadBalancerClient(cc *grpc.ClientConn) LoadBalancerClient {
	return &loadBalancerClient{cc}
}

func (c *loadBalancerClient) BalanceLoad(ctx context.Context, opts ...grpc.CallOption) (LoadBalancer_BalanceLoadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LoadBalancer_serviceDesc.Streams[0], "/grpc.lb.v1.LoadBalancer/BalanceLoad", opts...)
	if err != nil {
		return nil, err
	}
	x := &loadBalancerBalanceLoadClient{stream}
	return x, nil
}

type LoadBalancer_BalanceLoadClient interface {
	Send(*LoadBalanceRequest) error
	Recv() (*LoadBalanceResponse, error)
	grpc.ClientStream
}

type loadBalancerBalanceLoadClient struct {
	grpc.ClientStream
}

func (x *loadBalancerBalanceLoadClient) Send(m *LoadBalanceRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *loadBalancerBalanceLoadClient) Recv() (*LoadBalanceResponse, error) {
	m := new(LoadBalanceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LoadBalancerServer is the server API for LoadBalancer service.
type LoadBalancerServer interface {
	// Bidirectional rpc to get a list of servers.
	BalanceLoad(LoadBalancer_BalanceLoadServer) error
}

// UnimplementedLoadBalancerServer can be embedded to have forward compatible implementations.
type UnimplementedLoadBalancerServer struct {
}

func (*UnimplementedLoadBalancerServer) BalanceLoad(srv LoadBalancer_BalanceLoadServer) error {
	return status.Errorf(codes.Unimplemented, "method BalanceLoad not implemented")
}

func RegisterLoadBalancerServer(s *grpc.Server, srv LoadBalancerServer) {
	s.RegisterService(&_LoadBalancer_serviceDesc, srv)
}

func _LoadBalancer_BalanceLoad_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LoadBalancerServer).BalanceLoad(&loadBalancerBalanceLoadServer{stream})
}

type LoadBalancer_BalanceLoadServer interface {
	Send(*LoadBalanceResponse) error
	Recv() (*LoadBalanceRequest, error)
	grpc.ServerStream
}

type loadBalancerBalanceLoadServer struct {
	grpc.ServerStream
}

func (x *loadBalancerBalanceLoadServer) Send(m *LoadBalanceResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *loadBalancerBalanceLoadServer) Recv() (*LoadBalanceRequest, error) {
	m := new(LoadBalanceRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _LoadBalancer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.lb.v1.LoadBalancer",
	HandlerType: (*LoadBalancerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BalanceLoad",
			Handler:       _LoadBalancer_BalanceLoad_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc/lb/v1/load_balancer.proto",
}
