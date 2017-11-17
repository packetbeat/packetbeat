// Code generated by protoc-gen-gogo.
// source: k8s.io/kubernetes/pkg/apis/authentication/v1beta1/generated.proto
// DO NOT EDIT!

/*
	Package v1beta1 is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/kubernetes/pkg/apis/authentication/v1beta1/generated.proto

	It has these top-level messages:
		ExtraValue
		TokenReview
		TokenReviewSpec
		TokenReviewStatus
		UserInfo
*/
package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import k8s_io_kubernetes_pkg_apis_meta_v1 "github.com/ericchiang/k8s/apis/meta/v1"
import _ "github.com/ericchiang/k8s/runtime"
import _ "github.com/ericchiang/k8s/runtime/schema"
import _ "github.com/ericchiang/k8s/util/intstr"
import _ "github.com/ericchiang/k8s/api/v1"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// ExtraValue masks the value so protobuf can generate
// +protobuf.nullable=true
// +protobuf.options.(gogoproto.goproto_stringer)=false
type ExtraValue struct {
	Items            []string `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ExtraValue) Reset()                    { *m = ExtraValue{} }
func (m *ExtraValue) String() string            { return proto.CompactTextString(m) }
func (*ExtraValue) ProtoMessage()               {}
func (*ExtraValue) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *ExtraValue) GetItems() []string {
	if m != nil {
		return m.Items
	}
	return nil
}

// TokenReview attempts to authenticate a token to a known user.
// Note: TokenReview requests may be cached by the webhook token authenticator
// plugin in the kube-apiserver.
type TokenReview struct {
	// +optional
	Metadata *k8s_io_kubernetes_pkg_apis_meta_v1.ObjectMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// Spec holds information about the request being evaluated
	Spec *TokenReviewSpec `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	// Status is filled in by the server and indicates whether the request can be authenticated.
	// +optional
	Status           *TokenReviewStatus `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *TokenReview) Reset()                    { *m = TokenReview{} }
func (m *TokenReview) String() string            { return proto.CompactTextString(m) }
func (*TokenReview) ProtoMessage()               {}
func (*TokenReview) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func (m *TokenReview) GetMetadata() *k8s_io_kubernetes_pkg_apis_meta_v1.ObjectMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *TokenReview) GetSpec() *TokenReviewSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *TokenReview) GetStatus() *TokenReviewStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

// TokenReviewSpec is a description of the token authentication request.
type TokenReviewSpec struct {
	// Token is the opaque bearer token.
	// +optional
	Token            *string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TokenReviewSpec) Reset()                    { *m = TokenReviewSpec{} }
func (m *TokenReviewSpec) String() string            { return proto.CompactTextString(m) }
func (*TokenReviewSpec) ProtoMessage()               {}
func (*TokenReviewSpec) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func (m *TokenReviewSpec) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

// TokenReviewStatus is the result of the token authentication request.
type TokenReviewStatus struct {
	// Authenticated indicates that the token was associated with a known user.
	// +optional
	Authenticated *bool `protobuf:"varint,1,opt,name=authenticated" json:"authenticated,omitempty"`
	// User is the UserInfo associated with the provided token.
	// +optional
	User *UserInfo `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	// Error indicates that the token couldn't be checked
	// +optional
	Error            *string `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TokenReviewStatus) Reset()                    { *m = TokenReviewStatus{} }
func (m *TokenReviewStatus) String() string            { return proto.CompactTextString(m) }
func (*TokenReviewStatus) ProtoMessage()               {}
func (*TokenReviewStatus) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{3} }

func (m *TokenReviewStatus) GetAuthenticated() bool {
	if m != nil && m.Authenticated != nil {
		return *m.Authenticated
	}
	return false
}

func (m *TokenReviewStatus) GetUser() *UserInfo {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *TokenReviewStatus) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

// UserInfo holds the information about the user needed to implement the
// user.Info interface.
type UserInfo struct {
	// The name that uniquely identifies this user among all active users.
	// +optional
	Username *string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	// A unique value that identifies this user across time. If this user is
	// deleted and another user by the same name is added, they will have
	// different UIDs.
	// +optional
	Uid *string `protobuf:"bytes,2,opt,name=uid" json:"uid,omitempty"`
	// The names of groups this user is a part of.
	// +optional
	Groups []string `protobuf:"bytes,3,rep,name=groups" json:"groups,omitempty"`
	// Any additional information provided by the authenticator.
	// +optional
	Extra            map[string]*ExtraValue `protobuf:"bytes,4,rep,name=extra" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *UserInfo) Reset()                    { *m = UserInfo{} }
func (m *UserInfo) String() string            { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()               {}
func (*UserInfo) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{4} }

func (m *UserInfo) GetUsername() string {
	if m != nil && m.Username != nil {
		return *m.Username
	}
	return ""
}

func (m *UserInfo) GetUid() string {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return ""
}

func (m *UserInfo) GetGroups() []string {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *UserInfo) GetExtra() map[string]*ExtraValue {
	if m != nil {
		return m.Extra
	}
	return nil
}

func init() {
	proto.RegisterType((*ExtraValue)(nil), "github.com/ericchiang.k8s.apis.authentication.v1beta1.ExtraValue")
	proto.RegisterType((*TokenReview)(nil), "github.com/ericchiang.k8s.apis.authentication.v1beta1.TokenReview")
	proto.RegisterType((*TokenReviewSpec)(nil), "github.com/ericchiang.k8s.apis.authentication.v1beta1.TokenReviewSpec")
	proto.RegisterType((*TokenReviewStatus)(nil), "github.com/ericchiang.k8s.apis.authentication.v1beta1.TokenReviewStatus")
	proto.RegisterType((*UserInfo)(nil), "github.com/ericchiang.k8s.apis.authentication.v1beta1.UserInfo")
}
func (m *ExtraValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExtraValue) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, s := range m.Items {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *TokenReview) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenReview) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Metadata.Size()))
		n1, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Spec != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Spec.Size()))
		n2, err := m.Spec.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Status != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Status.Size()))
		n3, err := m.Status.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *TokenReviewSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenReviewSpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Token != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(*m.Token)))
		i += copy(dAtA[i:], *m.Token)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *TokenReviewStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenReviewStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Authenticated != nil {
		dAtA[i] = 0x8
		i++
		if *m.Authenticated {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.User != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.User.Size()))
		n4, err := m.User.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.Error != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(*m.Error)))
		i += copy(dAtA[i:], *m.Error)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *UserInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Username != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(*m.Username)))
		i += copy(dAtA[i:], *m.Username)
	}
	if m.Uid != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(*m.Uid)))
		i += copy(dAtA[i:], *m.Uid)
	}
	if len(m.Groups) > 0 {
		for _, s := range m.Groups {
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Extra) > 0 {
		for k, _ := range m.Extra {
			dAtA[i] = 0x22
			i++
			v := m.Extra[k]
			msgSize := 0
			if v != nil {
				msgSize = v.Size()
				msgSize += 1 + sovGenerated(uint64(msgSize))
			}
			mapSize := 1 + len(k) + sovGenerated(uint64(len(k))) + msgSize
			i = encodeVarintGenerated(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			if v != nil {
				dAtA[i] = 0x12
				i++
				i = encodeVarintGenerated(dAtA, i, uint64(v.Size()))
				n5, err := v.MarshalTo(dAtA[i:])
				if err != nil {
					return 0, err
				}
				i += n5
			}
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64Generated(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Generated(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ExtraValue) Size() (n int) {
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, s := range m.Items {
			l = len(s)
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TokenReview) Size() (n int) {
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Spec != nil {
		l = m.Spec.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Status != nil {
		l = m.Status.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TokenReviewSpec) Size() (n int) {
	var l int
	_ = l
	if m.Token != nil {
		l = len(*m.Token)
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TokenReviewStatus) Size() (n int) {
	var l int
	_ = l
	if m.Authenticated != nil {
		n += 2
	}
	if m.User != nil {
		l = m.User.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Error != nil {
		l = len(*m.Error)
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *UserInfo) Size() (n int) {
	var l int
	_ = l
	if m.Username != nil {
		l = len(*m.Username)
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Uid != nil {
		l = len(*m.Uid)
		n += 1 + l + sovGenerated(uint64(l))
	}
	if len(m.Groups) > 0 {
		for _, s := range m.Groups {
			l = len(s)
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if len(m.Extra) > 0 {
		for k, v := range m.Extra {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovGenerated(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovGenerated(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovGenerated(uint64(mapEntrySize))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ExtraValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ExtraValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExtraValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TokenReview) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TokenReview: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenReview: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &k8s_io_kubernetes_pkg_apis_meta_v1.ObjectMeta{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Spec == nil {
				m.Spec = &TokenReviewSpec{}
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Status == nil {
				m.Status = &TokenReviewStatus{}
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TokenReviewSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TokenReviewSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenReviewSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Token = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TokenReviewStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TokenReviewStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenReviewStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authenticated", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Authenticated = &b
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.User == nil {
				m.User = &UserInfo{}
			}
			if err := m.User.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Error = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UserInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UserInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Username = &s
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Uid = &s
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Groups", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Groups = append(m.Groups, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Extra", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthGenerated
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(dAtA[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			if m.Extra == nil {
				m.Extra = make(map[string]*ExtraValue)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					valuekey |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				var mapmsglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					mapmsglen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if mapmsglen < 0 {
					return ErrInvalidLengthGenerated
				}
				postmsgIndex := iNdEx + mapmsglen
				if mapmsglen < 0 {
					return ErrInvalidLengthGenerated
				}
				if postmsgIndex > l {
					return io.ErrUnexpectedEOF
				}
				mapvalue := &ExtraValue{}
				if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
					return err
				}
				iNdEx = postmsgIndex
				m.Extra[mapkey] = mapvalue
			} else {
				var mapvalue *ExtraValue
				m.Extra[mapkey] = mapvalue
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGenerated(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/ericchiang/k8s/apis/authentication/v1beta1/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 498 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x93, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0x71, 0x3e, 0x4a, 0xb2, 0x11, 0x02, 0x56, 0x08, 0x99, 0x1c, 0xa2, 0xc8, 0x42, 0x22,
	0x07, 0x58, 0xcb, 0x11, 0x87, 0x0a, 0xc4, 0x81, 0x8a, 0x22, 0x81, 0x84, 0x2a, 0x6d, 0xa1, 0x07,
	0xd4, 0xcb, 0xc6, 0x19, 0xd2, 0xc5, 0xf5, 0xda, 0xda, 0x1d, 0xbb, 0xf4, 0x29, 0xb8, 0x72, 0xe4,
	0xc6, 0xab, 0x70, 0xe4, 0x11, 0x50, 0x78, 0x11, 0xb4, 0x6b, 0xd3, 0x0f, 0x1c, 0x22, 0xd1, 0xdc,
	0x76, 0x46, 0xf3, 0xff, 0xcd, 0xcc, 0xdf, 0x63, 0xf2, 0x3c, 0xd9, 0x36, 0x4c, 0x66, 0x61, 0x52,
	0xcc, 0x40, 0x2b, 0x40, 0x30, 0x61, 0x9e, 0x2c, 0x42, 0x91, 0x4b, 0x13, 0x8a, 0x02, 0x8f, 0x40,
	0xa1, 0x8c, 0x05, 0xca, 0x4c, 0x85, 0x65, 0x34, 0x03, 0x14, 0x51, 0xb8, 0x00, 0x05, 0x5a, 0x20,
	0xcc, 0x59, 0xae, 0x33, 0xcc, 0x68, 0x54, 0x21, 0xd8, 0x39, 0x82, 0xe5, 0xc9, 0x82, 0x59, 0x04,
	0xbb, 0x8c, 0x60, 0x35, 0x62, 0x38, 0x5d, 0xd3, 0x35, 0x05, 0x14, 0x61, 0xd9, 0x68, 0x33, 0x7c,
	0xb4, 0x5a, 0xa3, 0x0b, 0x85, 0x32, 0x85, 0x46, 0xf9, 0xe3, 0xf5, 0xe5, 0x26, 0x3e, 0x82, 0x54,
	0x34, 0x54, 0xd1, 0x6a, 0x55, 0x81, 0xf2, 0x38, 0x94, 0x0a, 0x0d, 0xea, 0x86, 0xe4, 0xe1, 0x3f,
	0x77, 0x59, 0xb1, 0x45, 0x10, 0x10, 0xb2, 0xfb, 0x09, 0xb5, 0x38, 0x10, 0xc7, 0x05, 0xd0, 0x3b,
	0xa4, 0x2b, 0x11, 0x52, 0xe3, 0x7b, 0xe3, 0xf6, 0xa4, 0xcf, 0xab, 0x20, 0xf8, 0xdc, 0x22, 0x83,
	0xb7, 0x59, 0x02, 0x8a, 0x43, 0x29, 0xe1, 0x84, 0xbe, 0x26, 0x3d, 0x6b, 0xca, 0x5c, 0xa0, 0xf0,
	0xbd, 0xb1, 0x37, 0x19, 0x4c, 0x19, 0x5b, 0xe3, 0xb9, 0xad, 0x65, 0x65, 0xc4, 0xf6, 0x66, 0x1f,
	0x21, 0xc6, 0x37, 0x80, 0x82, 0x9f, 0xe9, 0xe9, 0x01, 0xe9, 0x98, 0x1c, 0x62, 0xbf, 0xe5, 0x38,
	0x3b, 0xec, 0xbf, 0xbf, 0x1d, 0xbb, 0x30, 0xd9, 0x7e, 0x0e, 0x31, 0x77, 0x3c, 0x7a, 0x48, 0xb6,
	0x0c, 0x0a, 0x2c, 0x8c, 0xdf, 0x76, 0xe4, 0x17, 0x1b, 0x92, 0x1d, 0x8b, 0xd7, 0xcc, 0xe0, 0x01,
	0xb9, 0xf9, 0x57, 0x5b, 0x6b, 0x1d, 0xda, 0x94, 0x73, 0xa4, 0xcf, 0xab, 0x20, 0xf8, 0xea, 0x91,
	0xdb, 0x0d, 0x0c, 0xbd, 0x4f, 0x6e, 0x5c, 0x68, 0x09, 0x73, 0xa7, 0xe9, 0xf1, 0xcb, 0x49, 0xba,
	0x47, 0x3a, 0x85, 0x01, 0x5d, 0x5b, 0xf3, 0xf4, 0x0a, 0x0b, 0xbc, 0x33, 0xa0, 0x5f, 0xa9, 0x0f,
	0x19, 0x77, 0x20, 0x3b, 0x22, 0x68, 0x9d, 0x69, 0x67, 0x49, 0x9f, 0x57, 0x41, 0xf0, 0xad, 0x45,
	0x7a, 0x7f, 0x0a, 0xe9, 0x90, 0xf4, 0x6c, 0xa9, 0x12, 0x29, 0xd4, 0x8b, 0x9c, 0xc5, 0xf4, 0x16,
	0x69, 0x17, 0x72, 0xee, 0xc6, 0xe9, 0x73, 0xfb, 0xa4, 0x77, 0xc9, 0xd6, 0x42, 0x67, 0x45, 0x6e,
	0x4d, 0xb6, 0xf7, 0x52, 0x47, 0xf4, 0x90, 0x74, 0xc1, 0x1e, 0x95, 0xdf, 0x19, 0xb7, 0x27, 0x83,
	0xe9, 0xcb, 0x0d, 0x46, 0x67, 0xee, 0x3a, 0x77, 0x15, 0xea, 0x53, 0x5e, 0x41, 0x87, 0x27, 0xf5,
	0xc9, 0xba, 0xa4, 0x9d, 0x2a, 0x81, 0xd3, 0x7a, 0x58, 0xfb, 0xa4, 0xfb, 0xa4, 0x5b, 0xda, 0x6b,
	0xae, 0x8d, 0x7b, 0x76, 0x85, 0xee, 0xe7, 0xbf, 0x04, 0xaf, 0x58, 0x4f, 0x5a, 0xdb, 0xde, 0xce,
	0xbd, 0xef, 0xcb, 0x91, 0xf7, 0x63, 0x39, 0xf2, 0x7e, 0x2e, 0x47, 0xde, 0x97, 0x5f, 0xa3, 0x6b,
	0xef, 0xaf, 0xd7, 0x82, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x43, 0x77, 0xd5, 0x8a, 0xb7, 0x04,
	0x00, 0x00,
}
