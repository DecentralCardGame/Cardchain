// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cardchain/cardchain/match_reporter_proposal.proto

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

type MatchReporterProposal struct {
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Reporter    string `protobuf:"bytes,3,opt,name=reporter,proto3" json:"reporter,omitempty"`
}

func (m *MatchReporterProposal) Reset()         { *m = MatchReporterProposal{} }
func (m *MatchReporterProposal) String() string { return proto.CompactTextString(m) }
func (*MatchReporterProposal) ProtoMessage()    {}
func (*MatchReporterProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f42125af27677dc, []int{0}
}
func (m *MatchReporterProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MatchReporterProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MatchReporterProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MatchReporterProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchReporterProposal.Merge(m, src)
}
func (m *MatchReporterProposal) XXX_Size() int {
	return m.Size()
}
func (m *MatchReporterProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchReporterProposal.DiscardUnknown(m)
}

var xxx_messageInfo_MatchReporterProposal proto.InternalMessageInfo

func (m *MatchReporterProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *MatchReporterProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *MatchReporterProposal) GetReporter() string {
	if m != nil {
		return m.Reporter
	}
	return ""
}

func init() {
	proto.RegisterType((*MatchReporterProposal)(nil), "DecentralCardGame.cardchain.cardchain.MatchReporterProposal")
}

func init() {
	proto.RegisterFile("cardchain/cardchain/match_reporter_proposal.proto", fileDescriptor_5f42125af27677dc)
}

var fileDescriptor_5f42125af27677dc = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4c, 0x4e, 0x2c, 0x4a,
	0x49, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x47, 0xb0, 0x72, 0x13, 0x4b, 0x92, 0x33, 0xe2, 0x8b, 0x52,
	0x0b, 0xf2, 0x8b, 0x4a, 0x52, 0x8b, 0xe2, 0x0b, 0x8a, 0xf2, 0x0b, 0xf2, 0x8b, 0x13, 0x73, 0xf4,
	0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x54, 0x5d, 0x52, 0x93, 0x53, 0xf3, 0x4a, 0x8a, 0x12, 0x73,
	0x9c, 0x13, 0x8b, 0x52, 0xdc, 0x13, 0x73, 0x53, 0xf5, 0xe0, 0x5a, 0x11, 0x2c, 0xa5, 0x6c, 0x2e,
	0x51, 0x5f, 0x90, 0x39, 0x41, 0x50, 0x63, 0x02, 0xa0, 0xa6, 0x08, 0x89, 0x70, 0xb1, 0x96, 0x64,
	0x96, 0xe4, 0xa4, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x38, 0x42, 0x0a, 0x5c, 0xdc,
	0x29, 0xa9, 0xc5, 0xc9, 0x45, 0x99, 0x05, 0x25, 0x99, 0xf9, 0x79, 0x12, 0x4c, 0x60, 0x39, 0x64,
	0x21, 0x21, 0x29, 0x2e, 0x0e, 0x98, 0x93, 0x24, 0x98, 0xc1, 0xd2, 0x70, 0xbe, 0x53, 0xd0, 0x89,
	0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3,
	0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x59, 0xa4, 0x67, 0x96, 0x64, 0x94, 0x26,
	0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x63, 0x38, 0x5c, 0xdf, 0x19, 0xee, 0xe7, 0x0a, 0x24, 0xff, 0x97,
	0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0xbd, 0x6b, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xaa,
	0x3c, 0x6d, 0x58, 0x23, 0x01, 0x00, 0x00,
}

func (m *MatchReporterProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MatchReporterProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MatchReporterProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Reporter) > 0 {
		i -= len(m.Reporter)
		copy(dAtA[i:], m.Reporter)
		i = encodeVarintMatchReporterProposal(dAtA, i, uint64(len(m.Reporter)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintMatchReporterProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintMatchReporterProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMatchReporterProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovMatchReporterProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MatchReporterProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovMatchReporterProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovMatchReporterProposal(uint64(l))
	}
	l = len(m.Reporter)
	if l > 0 {
		n += 1 + l + sovMatchReporterProposal(uint64(l))
	}
	return n
}

func sovMatchReporterProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMatchReporterProposal(x uint64) (n int) {
	return sovMatchReporterProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MatchReporterProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMatchReporterProposal
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
			return fmt.Errorf("proto: MatchReporterProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MatchReporterProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMatchReporterProposal
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
				return ErrInvalidLengthMatchReporterProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMatchReporterProposal
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
					return ErrIntOverflowMatchReporterProposal
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
				return ErrInvalidLengthMatchReporterProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMatchReporterProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reporter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMatchReporterProposal
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
				return ErrInvalidLengthMatchReporterProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMatchReporterProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reporter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMatchReporterProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMatchReporterProposal
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
func skipMatchReporterProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMatchReporterProposal
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
					return 0, ErrIntOverflowMatchReporterProposal
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
					return 0, ErrIntOverflowMatchReporterProposal
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
				return 0, ErrInvalidLengthMatchReporterProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMatchReporterProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMatchReporterProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMatchReporterProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMatchReporterProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMatchReporterProposal = fmt.Errorf("proto: unexpected end of group")
)
