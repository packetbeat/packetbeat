// Code generated by protoc-gen-gogo.
// source: uuid.proto
// DO NOT EDIT!

package events

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// / Type representing a 128-bit UUID.
//
// The bytes of the UUID should be packed in little-endian **byte** (not bit) order. For example, the UUID `f47ac10b-58cc-4372-a567-0e02b2c3d479` should be encoded as `UUID{ low: 0x7243cc580bc17af4, high: 0x79d4c3b2020e67a5 }`
type UUID struct {
	Low              *uint64 `protobuf:"varint,1,req,name=low" json:"low,omitempty"`
	High             *uint64 `protobuf:"varint,2,req,name=high" json:"high,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UUID) Reset()                    { *m = UUID{} }
func (m *UUID) String() string            { return proto.CompactTextString(m) }
func (*UUID) ProtoMessage()               {}
func (*UUID) Descriptor() ([]byte, []int) { return fileDescriptorUuid, []int{0} }

func (m *UUID) GetLow() uint64 {
	if m != nil && m.Low != nil {
		return *m.Low
	}
	return 0
}

func (m *UUID) GetHigh() uint64 {
	if m != nil && m.High != nil {
		return *m.High
	}
	return 0
}

func init() {
	proto.RegisterType((*UUID)(nil), "events.UUID")
}
func (m *UUID) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *UUID) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Low == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("low")
	} else {
		data[i] = 0x8
		i++
		i = encodeVarintUuid(data, i, uint64(*m.Low))
	}
	if m.High == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("high")
	} else {
		data[i] = 0x10
		i++
		i = encodeVarintUuid(data, i, uint64(*m.High))
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64Uuid(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Uuid(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintUuid(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *UUID) Size() (n int) {
	var l int
	_ = l
	if m.Low != nil {
		n += 1 + sovUuid(uint64(*m.Low))
	}
	if m.High != nil {
		n += 1 + sovUuid(uint64(*m.High))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovUuid(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozUuid(x uint64) (n int) {
	return sovUuid(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UUID) Unmarshal(data []byte) error {
	var hasFields [1]uint64
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUuid
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UUID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UUID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Low", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUuid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Low = &v
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field High", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUuid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.High = &v
			hasFields[0] |= uint64(0x00000002)
		default:
			iNdEx = preIndex
			skippy, err := skipUuid(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUuid
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("low")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("high")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipUuid(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUuid
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowUuid
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowUuid
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthUuid
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowUuid
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipUuid(data[start:])
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
	ErrInvalidLengthUuid = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUuid   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("uuid.proto", fileDescriptorUuid) }

var fileDescriptorUuid = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0xcd, 0x4c,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0x2d, 0x4b, 0xcd, 0x2b, 0x29, 0x96, 0xd2,
	0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7,
	0x07, 0x4b, 0x27, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0xa6, 0xa4, 0xc3, 0xc5,
	0x12, 0x1a, 0xea, 0xe9, 0x22, 0x24, 0xc0, 0xc5, 0x9c, 0x93, 0x5f, 0x2e, 0xc1, 0xa8, 0xc0, 0xa4,
	0xc1, 0x12, 0x04, 0x62, 0x0a, 0x09, 0x71, 0xb1, 0x64, 0x64, 0xa6, 0x67, 0x48, 0x30, 0x81, 0x85,
	0xc0, 0x6c, 0x27, 0x9b, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e,
	0x91, 0x4b, 0x31, 0xbf, 0x28, 0x5d, 0x2f, 0x39, 0x27, 0xbf, 0x34, 0x25, 0x2d, 0xbf, 0x34, 0x2f,
	0xa5, 0xa8, 0x52, 0x2f, 0xa5, 0x28, 0xbf, 0xa0, 0x38, 0x3f, 0x2f, 0x25, 0x55, 0x0f, 0xe2, 0x1a,
	0x27, 0xee, 0xd0, 0xd2, 0xcc, 0x14, 0xb7, 0xc4, 0xe4, 0x92, 0xfc, 0xa2, 0x4a, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xc1, 0x49, 0x67, 0x8e, 0xaf, 0x00, 0x00, 0x00,
}
