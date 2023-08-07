// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: DecentralCardGame/cardchain/cardchain/set_proposal.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type SetProposal struct {
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	SetId       uint64 `protobuf:"varint,3,opt,name=setId,proto3" json:"setId,omitempty"`
}

func (m *SetProposal) Reset()         { *m = SetProposal{} }
func (m *SetProposal) String() string { return proto.CompactTextString(m) }
func (*SetProposal) ProtoMessage()    {}
func (*SetProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_5408b325fd96949c, []int{0}
}
func (m *SetProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SetProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetProposal.Merge(m, src)
}
func (m *SetProposal) XXX_Size() int {
	return m.Size()
}
func (m *SetProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_SetProposal.DiscardUnknown(m)
}

var xxx_messageInfo_SetProposal proto.InternalMessageInfo

func (m *SetProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SetProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SetProposal) GetSetId() uint64 {
	if m != nil {
		return m.SetId
	}
	return 0
}

func init() {
	proto.RegisterType((*SetProposal)(nil), "DecentralCardGame.cardchain.cardchain.SetProposal")
}

func init() {
	proto.RegisterFile("DecentralCardGame/cardchain/cardchain/set_proposal.proto", fileDescriptor_5408b325fd96949c)
}

var fileDescriptor_5408b325fd96949c = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x70, 0x49, 0x4d, 0x4e,
	0xcd, 0x2b, 0x29, 0x4a, 0xcc, 0x71, 0x4e, 0x2c, 0x4a, 0x71, 0x4f, 0xcc, 0x4d, 0xd5, 0x4f, 0x4e,
	0x2c, 0x4a, 0x49, 0xce, 0x48, 0xcc, 0xcc, 0x43, 0x62, 0x15, 0xa7, 0x96, 0xc4, 0x17, 0x14, 0xe5,
	0x17, 0xe4, 0x17, 0x27, 0xe6, 0xe8, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0xa9, 0x62, 0xe8, 0xd4,
	0x83, 0xab, 0x47, 0xb0, 0x94, 0xa2, 0xb9, 0xb8, 0x83, 0x53, 0x4b, 0x02, 0xa0, 0x7a, 0x85, 0x44,
	0xb8, 0x58, 0x4b, 0x32, 0x4b, 0x72, 0x52, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x20, 0x1c,
	0x21, 0x05, 0x2e, 0xee, 0x94, 0xd4, 0xe2, 0xe4, 0xa2, 0xcc, 0x82, 0x92, 0xcc, 0xfc, 0x3c, 0x09,
	0x26, 0xb0, 0x1c, 0xb2, 0x10, 0x48, 0x5f, 0x71, 0x6a, 0x89, 0x67, 0x8a, 0x04, 0xb3, 0x02, 0xa3,
	0x06, 0x4b, 0x10, 0x84, 0xe3, 0x14, 0x74, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f,
	0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c,
	0x51, 0x16, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x98, 0x5e, 0x74,
	0x86, 0x7b, 0xac, 0x02, 0xc9, 0x93, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0xef, 0x19,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x63, 0x13, 0x12, 0x1a, 0x01, 0x00, 0x00,
}

func (m *SetProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SetProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SetId != 0 {
		i = encodeVarintSetProposal(dAtA, i, uint64(m.SetId))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintSetProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintSetProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSetProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovSetProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SetProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovSetProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovSetProposal(uint64(l))
	}
	if m.SetId != 0 {
		n += 1 + sovSetProposal(uint64(m.SetId))
	}
	return n
}

func sovSetProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSetProposal(x uint64) (n int) {
	return sovSetProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SetProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSetProposal
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
			return fmt.Errorf("proto: SetProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSetProposal
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
				return ErrInvalidLengthSetProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSetProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSetProposal
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
				return ErrInvalidLengthSetProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSetProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SetId", wireType)
			}
			m.SetId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSetProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SetId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSetProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSetProposal
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
func skipSetProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSetProposal
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
					return 0, ErrIntOverflowSetProposal
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
					return 0, ErrIntOverflowSetProposal
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
				return 0, ErrInvalidLengthSetProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSetProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSetProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSetProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSetProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSetProposal = fmt.Errorf("proto: unexpected end of group")
)
