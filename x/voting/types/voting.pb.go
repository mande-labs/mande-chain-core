// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mande/voting/v1beta1/voting.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type VoteBook struct {
	Index    string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Positive uint64 `protobuf:"varint,2,opt,name=positive,proto3" json:"positive,omitempty"`
	Negative uint64 `protobuf:"varint,3,opt,name=negative,proto3" json:"negative,omitempty"`
}

func (m *VoteBook) Reset()         { *m = VoteBook{} }
func (m *VoteBook) String() string { return proto.CompactTextString(m) }
func (*VoteBook) ProtoMessage()    {}
func (*VoteBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4ddf73ad3e922a0, []int{0}
}
func (m *VoteBook) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VoteBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VoteBook.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VoteBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteBook.Merge(m, src)
}
func (m *VoteBook) XXX_Size() int {
	return m.Size()
}
func (m *VoteBook) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteBook.DiscardUnknown(m)
}

var xxx_messageInfo_VoteBook proto.InternalMessageInfo

func (m *VoteBook) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *VoteBook) GetPositive() uint64 {
	if m != nil {
		return m.Positive
	}
	return 0
}

func (m *VoteBook) GetNegative() uint64 {
	if m != nil {
		return m.Negative
	}
	return 0
}

type AggregateVoteCount struct {
	Index                  string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	AggregateVotesCasted   uint64 `protobuf:"varint,2,opt,name=aggregateVotesCasted,proto3" json:"aggregateVotesCasted,omitempty"`
	AggregateVotesReceived uint64 `protobuf:"varint,3,opt,name=aggregateVotesReceived,proto3" json:"aggregateVotesReceived,omitempty"`
}

func (m *AggregateVoteCount) Reset()         { *m = AggregateVoteCount{} }
func (m *AggregateVoteCount) String() string { return proto.CompactTextString(m) }
func (*AggregateVoteCount) ProtoMessage()    {}
func (*AggregateVoteCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4ddf73ad3e922a0, []int{1}
}
func (m *AggregateVoteCount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AggregateVoteCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AggregateVoteCount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AggregateVoteCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AggregateVoteCount.Merge(m, src)
}
func (m *AggregateVoteCount) XXX_Size() int {
	return m.Size()
}
func (m *AggregateVoteCount) XXX_DiscardUnknown() {
	xxx_messageInfo_AggregateVoteCount.DiscardUnknown(m)
}

var xxx_messageInfo_AggregateVoteCount proto.InternalMessageInfo

func (m *AggregateVoteCount) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *AggregateVoteCount) GetAggregateVotesCasted() uint64 {
	if m != nil {
		return m.AggregateVotesCasted
	}
	return 0
}

func (m *AggregateVoteCount) GetAggregateVotesReceived() uint64 {
	if m != nil {
		return m.AggregateVotesReceived
	}
	return 0
}

func init() {
	proto.RegisterType((*VoteBook)(nil), "mande.voting.v1beta1.VoteBook")
	proto.RegisterType((*AggregateVoteCount)(nil), "mande.voting.v1beta1.AggregateVoteCount")
}

func init() { proto.RegisterFile("mande/voting/v1beta1/voting.proto", fileDescriptor_a4ddf73ad3e922a0) }

var fileDescriptor_a4ddf73ad3e922a0 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcc, 0x4d, 0xcc, 0x4b,
	0x49, 0xd5, 0x2f, 0xcb, 0x2f, 0xc9, 0xcc, 0x4b, 0xd7, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34,
	0x84, 0x72, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x44, 0xc0, 0x4a, 0xf4, 0xa0, 0x62, 0x50,
	0x25, 0x4a, 0x11, 0x5c, 0x1c, 0x61, 0xf9, 0x25, 0xa9, 0x4e, 0xf9, 0xf9, 0xd9, 0x42, 0x22, 0x5c,
	0xac, 0x99, 0x79, 0x29, 0xa9, 0x15, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x8e, 0x90,
	0x14, 0x17, 0x47, 0x41, 0x7e, 0x71, 0x66, 0x49, 0x66, 0x59, 0xaa, 0x04, 0x93, 0x02, 0xa3, 0x06,
	0x4b, 0x10, 0x9c, 0x0f, 0x92, 0xcb, 0x4b, 0x4d, 0x4f, 0x04, 0xcb, 0x31, 0x43, 0xe4, 0x60, 0x7c,
	0xa5, 0x69, 0x8c, 0x5c, 0x42, 0x8e, 0xe9, 0xe9, 0x45, 0x20, 0x7e, 0x2a, 0xc8, 0x0e, 0xe7, 0xfc,
	0xd2, 0xbc, 0x12, 0x1c, 0x96, 0x18, 0x71, 0x89, 0x24, 0x22, 0xab, 0x2d, 0x76, 0x4e, 0x2c, 0x2e,
	0x49, 0x4d, 0x81, 0x5a, 0x88, 0x55, 0x4e, 0xc8, 0x8c, 0x4b, 0x0c, 0x55, 0x3c, 0x28, 0x35, 0x39,
	0x35, 0xb3, 0x2c, 0x35, 0x05, 0xea, 0x14, 0x1c, 0xb2, 0x4e, 0xee, 0x27, 0x1e, 0xc9, 0x31, 0x5e,
	0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31,
	0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x9b, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f,
	0xab, 0x0f, 0x0e, 0x2d, 0xdd, 0x9c, 0xc4, 0xa4, 0x62, 0x7d, 0x68, 0xd8, 0x1a, 0xea, 0x57, 0xc0,
	0x02, 0xb8, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0x1c, 0xb0, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xda, 0xdf, 0x2a, 0x2f, 0x7d, 0x01, 0x00, 0x00,
}

func (m *VoteBook) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VoteBook) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VoteBook) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Negative != 0 {
		i = encodeVarintVoting(dAtA, i, uint64(m.Negative))
		i--
		dAtA[i] = 0x18
	}
	if m.Positive != 0 {
		i = encodeVarintVoting(dAtA, i, uint64(m.Positive))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintVoting(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AggregateVoteCount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AggregateVoteCount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AggregateVoteCount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AggregateVotesReceived != 0 {
		i = encodeVarintVoting(dAtA, i, uint64(m.AggregateVotesReceived))
		i--
		dAtA[i] = 0x18
	}
	if m.AggregateVotesCasted != 0 {
		i = encodeVarintVoting(dAtA, i, uint64(m.AggregateVotesCasted))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintVoting(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintVoting(dAtA []byte, offset int, v uint64) int {
	offset -= sovVoting(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *VoteBook) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovVoting(uint64(l))
	}
	if m.Positive != 0 {
		n += 1 + sovVoting(uint64(m.Positive))
	}
	if m.Negative != 0 {
		n += 1 + sovVoting(uint64(m.Negative))
	}
	return n
}

func (m *AggregateVoteCount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovVoting(uint64(l))
	}
	if m.AggregateVotesCasted != 0 {
		n += 1 + sovVoting(uint64(m.AggregateVotesCasted))
	}
	if m.AggregateVotesReceived != 0 {
		n += 1 + sovVoting(uint64(m.AggregateVotesReceived))
	}
	return n
}

func sovVoting(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVoting(x uint64) (n int) {
	return sovVoting(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VoteBook) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVoting
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: VoteBook: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VoteBook: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVoting
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVoting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Positive", wireType)
			}
			m.Positive = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Positive |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Negative", wireType)
			}
			m.Negative = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Negative |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipVoting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVoting
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AggregateVoteCount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVoting
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AggregateVoteCount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AggregateVoteCount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVoting
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVoting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AggregateVotesCasted", wireType)
			}
			m.AggregateVotesCasted = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AggregateVotesCasted |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AggregateVotesReceived", wireType)
			}
			m.AggregateVotesReceived = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AggregateVotesReceived |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipVoting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVoting
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipVoting(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVoting
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
					return 0, ErrIntOverflowVoting
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowVoting
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
			if length < 0 {
				return 0, ErrInvalidLengthVoting
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVoting
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVoting
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVoting        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVoting          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVoting = fmt.Errorf("proto: unexpected end of group")
)