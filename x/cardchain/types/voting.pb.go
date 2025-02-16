// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cardchain/cardchain/voting.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/cosmos/gogoproto/proto"
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

type VoteType int32

const (
	VoteType_fairEnough    VoteType = 0
	VoteType_inappropriate VoteType = 1
	VoteType_overpowered   VoteType = 2
	VoteType_underpowered  VoteType = 3
)

var VoteType_name = map[int32]string{
	0: "fairEnough",
	1: "inappropriate",
	2: "overpowered",
	3: "underpowered",
}

var VoteType_value = map[string]int32{
	"fairEnough":    0,
	"inappropriate": 1,
	"overpowered":   2,
	"underpowered":  3,
}

func (x VoteType) String() string {
	return proto.EnumName(VoteType_name, int32(x))
}

func (VoteType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_39c6bcdea6f40fbf, []int{0}
}

type VotingResult struct {
	CardId             uint64 `protobuf:"varint,1,opt,name=cardId,proto3" json:"cardId,omitempty"`
	FairEnoughVotes    uint64 `protobuf:"varint,2,opt,name=fairEnoughVotes,proto3" json:"fairEnoughVotes,omitempty"`
	OverpoweredVotes   uint64 `protobuf:"varint,3,opt,name=overpoweredVotes,proto3" json:"overpoweredVotes,omitempty"`
	UnderpoweredVotes  uint64 `protobuf:"varint,4,opt,name=underpoweredVotes,proto3" json:"underpoweredVotes,omitempty"`
	InappropriateVotes uint64 `protobuf:"varint,5,opt,name=inappropriateVotes,proto3" json:"inappropriateVotes,omitempty"`
	Result             string `protobuf:"bytes,6,opt,name=result,proto3" json:"result,omitempty"`
}

func (m *VotingResult) Reset()         { *m = VotingResult{} }
func (m *VotingResult) String() string { return proto.CompactTextString(m) }
func (*VotingResult) ProtoMessage()    {}
func (*VotingResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_39c6bcdea6f40fbf, []int{0}
}
func (m *VotingResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VotingResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VotingResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VotingResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VotingResult.Merge(m, src)
}
func (m *VotingResult) XXX_Size() int {
	return m.Size()
}
func (m *VotingResult) XXX_DiscardUnknown() {
	xxx_messageInfo_VotingResult.DiscardUnknown(m)
}

var xxx_messageInfo_VotingResult proto.InternalMessageInfo

func (m *VotingResult) GetCardId() uint64 {
	if m != nil {
		return m.CardId
	}
	return 0
}

func (m *VotingResult) GetFairEnoughVotes() uint64 {
	if m != nil {
		return m.FairEnoughVotes
	}
	return 0
}

func (m *VotingResult) GetOverpoweredVotes() uint64 {
	if m != nil {
		return m.OverpoweredVotes
	}
	return 0
}

func (m *VotingResult) GetUnderpoweredVotes() uint64 {
	if m != nil {
		return m.UnderpoweredVotes
	}
	return 0
}

func (m *VotingResult) GetInappropriateVotes() uint64 {
	if m != nil {
		return m.InappropriateVotes
	}
	return 0
}

func (m *VotingResult) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type SingleVote struct {
	CardId   uint64   `protobuf:"varint,1,opt,name=cardId,proto3" json:"cardId,omitempty"`
	VoteType VoteType `protobuf:"varint,2,opt,name=voteType,proto3,enum=cardchain.cardchain.VoteType" json:"voteType,omitempty"`
}

func (m *SingleVote) Reset()         { *m = SingleVote{} }
func (m *SingleVote) String() string { return proto.CompactTextString(m) }
func (*SingleVote) ProtoMessage()    {}
func (*SingleVote) Descriptor() ([]byte, []int) {
	return fileDescriptor_39c6bcdea6f40fbf, []int{1}
}
func (m *SingleVote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SingleVote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SingleVote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SingleVote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SingleVote.Merge(m, src)
}
func (m *SingleVote) XXX_Size() int {
	return m.Size()
}
func (m *SingleVote) XXX_DiscardUnknown() {
	xxx_messageInfo_SingleVote.DiscardUnknown(m)
}

var xxx_messageInfo_SingleVote proto.InternalMessageInfo

func (m *SingleVote) GetCardId() uint64 {
	if m != nil {
		return m.CardId
	}
	return 0
}

func (m *SingleVote) GetVoteType() VoteType {
	if m != nil {
		return m.VoteType
	}
	return VoteType_fairEnough
}

func init() {
	proto.RegisterEnum("cardchain.cardchain.VoteType", VoteType_name, VoteType_value)
	proto.RegisterType((*VotingResult)(nil), "cardchain.cardchain.VotingResult")
	proto.RegisterType((*SingleVote)(nil), "cardchain.cardchain.SingleVote")
}

func init() { proto.RegisterFile("cardchain/cardchain/voting.proto", fileDescriptor_39c6bcdea6f40fbf) }

var fileDescriptor_39c6bcdea6f40fbf = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xbf, 0x4e, 0x32, 0x41,
	0x14, 0xc5, 0x77, 0x80, 0x8f, 0xf0, 0x5d, 0x11, 0x96, 0x31, 0x31, 0x34, 0x6e, 0x08, 0x15, 0x21,
	0x66, 0x49, 0xb4, 0xd1, 0xd6, 0x3f, 0x31, 0x76, 0x66, 0x35, 0x14, 0x36, 0x66, 0xd8, 0xbd, 0x2e,
	0x93, 0xc0, 0xcc, 0x64, 0x98, 0x45, 0x79, 0x0b, 0x1f, 0xcb, 0x92, 0xd2, 0xd2, 0xc0, 0x1b, 0xf8,
	0x04, 0x86, 0x01, 0x97, 0x15, 0xb0, 0x3b, 0x73, 0xce, 0xaf, 0x98, 0x73, 0x72, 0xa1, 0x11, 0x32,
	0x1d, 0x85, 0x7d, 0xc6, 0x45, 0x67, 0xad, 0xc6, 0xd2, 0x70, 0x11, 0xfb, 0x4a, 0x4b, 0x23, 0xe9,
	0x41, 0xea, 0xfb, 0xa9, 0x6a, 0x7e, 0x11, 0x28, 0x77, 0x2d, 0x15, 0xe0, 0x28, 0x19, 0x18, 0x7a,
	0x08, 0xc5, 0x45, 0x7a, 0x1b, 0xd5, 0x49, 0x83, 0xb4, 0x0a, 0xc1, 0xea, 0x45, 0x5b, 0x50, 0x7d,
	0x66, 0x5c, 0x5f, 0x0b, 0x99, 0xc4, 0xfd, 0xae, 0x34, 0x38, 0xaa, 0xe7, 0x2c, 0xb0, 0x69, 0xd3,
	0x36, 0xb8, 0x72, 0x8c, 0x5a, 0xc9, 0x17, 0xd4, 0x18, 0x2d, 0xd1, 0xbc, 0x45, 0xb7, 0x7c, 0x7a,
	0x0c, 0xb5, 0x44, 0x44, 0x1b, 0x70, 0xc1, 0xc2, 0xdb, 0x01, 0xf5, 0x81, 0x72, 0xc1, 0x94, 0xd2,
	0x52, 0x69, 0xce, 0x0c, 0x2e, 0xf1, 0x7f, 0x16, 0xdf, 0x91, 0x2c, 0xba, 0x68, 0xdb, 0xaa, 0x5e,
	0x6c, 0x90, 0xd6, 0xff, 0x60, 0xf5, 0x6a, 0x3e, 0x01, 0xdc, 0x73, 0x11, 0x0f, 0x2c, 0xf6, 0x67,
	0xe3, 0x73, 0x28, 0x8d, 0xa5, 0xc1, 0x87, 0x89, 0x42, 0x5b, 0xb5, 0x72, 0x72, 0xe4, 0xef, 0x98,
	0xd0, 0xef, 0xae, 0xa0, 0x20, 0xc5, 0xdb, 0x77, 0x50, 0xfa, 0x71, 0x69, 0x05, 0x60, 0xbd, 0x90,
	0xeb, 0xd0, 0x1a, 0xec, 0xff, 0xfa, 0xaa, 0x4b, 0x68, 0x15, 0xf6, 0x32, 0xcb, 0xb8, 0x39, 0xea,
	0x42, 0x39, 0xdb, 0xde, 0xcd, 0x5f, 0x04, 0xef, 0x33, 0x8f, 0x4c, 0x67, 0x1e, 0xf9, 0x9c, 0x79,
	0xe4, 0x6d, 0xee, 0x39, 0xd3, 0xb9, 0xe7, 0x7c, 0xcc, 0x3d, 0xe7, 0xf1, 0x2c, 0xe6, 0xa6, 0x9f,
	0xf4, 0xfc, 0x50, 0x0e, 0x3b, 0x57, 0x18, 0xa2, 0x30, 0x9a, 0x0d, 0x2e, 0x99, 0x8e, 0x6e, 0xd8,
	0x10, 0x33, 0xb7, 0xf0, 0x9a, 0xd1, 0x66, 0xa2, 0x70, 0xd4, 0x2b, 0xda, 0xbb, 0x38, 0xfd, 0x0e,
	0x00, 0x00, 0xff, 0xff, 0x26, 0xf8, 0x7b, 0x79, 0x3b, 0x02, 0x00, 0x00,
}

func (m *VotingResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VotingResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VotingResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Result) > 0 {
		i -= len(m.Result)
		copy(dAtA[i:], m.Result)
		i = encodeVarintVoting(dAtA, i, uint64(len(m.Result)))
		i--
		dAtA[i] = 0x32
	}
	if m.InappropriateVotes != 0 {
		i = encodeVarintVoting(dAtA, i, uint64(m.InappropriateVotes))
		i--
		dAtA[i] = 0x28
	}
	if m.UnderpoweredVotes != 0 {
		i = encodeVarintVoting(dAtA, i, uint64(m.UnderpoweredVotes))
		i--
		dAtA[i] = 0x20
	}
	if m.OverpoweredVotes != 0 {
		i = encodeVarintVoting(dAtA, i, uint64(m.OverpoweredVotes))
		i--
		dAtA[i] = 0x18
	}
	if m.FairEnoughVotes != 0 {
		i = encodeVarintVoting(dAtA, i, uint64(m.FairEnoughVotes))
		i--
		dAtA[i] = 0x10
	}
	if m.CardId != 0 {
		i = encodeVarintVoting(dAtA, i, uint64(m.CardId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SingleVote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SingleVote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SingleVote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.VoteType != 0 {
		i = encodeVarintVoting(dAtA, i, uint64(m.VoteType))
		i--
		dAtA[i] = 0x10
	}
	if m.CardId != 0 {
		i = encodeVarintVoting(dAtA, i, uint64(m.CardId))
		i--
		dAtA[i] = 0x8
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
func (m *VotingResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CardId != 0 {
		n += 1 + sovVoting(uint64(m.CardId))
	}
	if m.FairEnoughVotes != 0 {
		n += 1 + sovVoting(uint64(m.FairEnoughVotes))
	}
	if m.OverpoweredVotes != 0 {
		n += 1 + sovVoting(uint64(m.OverpoweredVotes))
	}
	if m.UnderpoweredVotes != 0 {
		n += 1 + sovVoting(uint64(m.UnderpoweredVotes))
	}
	if m.InappropriateVotes != 0 {
		n += 1 + sovVoting(uint64(m.InappropriateVotes))
	}
	l = len(m.Result)
	if l > 0 {
		n += 1 + l + sovVoting(uint64(l))
	}
	return n
}

func (m *SingleVote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CardId != 0 {
		n += 1 + sovVoting(uint64(m.CardId))
	}
	if m.VoteType != 0 {
		n += 1 + sovVoting(uint64(m.VoteType))
	}
	return n
}

func sovVoting(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVoting(x uint64) (n int) {
	return sovVoting(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VotingResult) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: VotingResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VotingResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CardId", wireType)
			}
			m.CardId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CardId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FairEnoughVotes", wireType)
			}
			m.FairEnoughVotes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FairEnoughVotes |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OverpoweredVotes", wireType)
			}
			m.OverpoweredVotes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OverpoweredVotes |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnderpoweredVotes", wireType)
			}
			m.UnderpoweredVotes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UnderpoweredVotes |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InappropriateVotes", wireType)
			}
			m.InappropriateVotes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InappropriateVotes |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
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
			m.Result = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
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
func (m *SingleVote) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: SingleVote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SingleVote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CardId", wireType)
			}
			m.CardId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CardId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteType", wireType)
			}
			m.VoteType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VoteType |= VoteType(b&0x7F) << shift
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
