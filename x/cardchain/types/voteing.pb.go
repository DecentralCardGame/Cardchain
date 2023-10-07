// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: DecentralCardGame/cardchain/cardchain/voteing.proto

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

type VotingResults struct {
	TotalVotes              uint64          `protobuf:"varint,1,opt,name=totalVotes,proto3" json:"totalVotes,omitempty"`
	TotalFairEnoughVotes    uint64          `protobuf:"varint,2,opt,name=totalFairEnoughVotes,proto3" json:"totalFairEnoughVotes,omitempty"`
	TotalOverpoweredVotes   uint64          `protobuf:"varint,3,opt,name=totalOverpoweredVotes,proto3" json:"totalOverpoweredVotes,omitempty"`
	TotalUnderpoweredVotes  uint64          `protobuf:"varint,4,opt,name=totalUnderpoweredVotes,proto3" json:"totalUnderpoweredVotes,omitempty"`
	TotalInappropriateVotes uint64          `protobuf:"varint,5,opt,name=totalInappropriateVotes,proto3" json:"totalInappropriateVotes,omitempty"`
	CardResults             []*VotingResult `protobuf:"bytes,6,rep,name=cardResults,proto3" json:"cardResults,omitempty"`
	Notes                   string          `protobuf:"bytes,7,opt,name=notes,proto3" json:"notes,omitempty"`
}

func (m *VotingResults) Reset()         { *m = VotingResults{} }
func (m *VotingResults) String() string { return proto.CompactTextString(m) }
func (*VotingResults) ProtoMessage()    {}
func (*VotingResults) Descriptor() ([]byte, []int) {
	return fileDescriptor_693a4ceca03817bb, []int{0}
}
func (m *VotingResults) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VotingResults) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VotingResults.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VotingResults) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VotingResults.Merge(m, src)
}
func (m *VotingResults) XXX_Size() int {
	return m.Size()
}
func (m *VotingResults) XXX_DiscardUnknown() {
	xxx_messageInfo_VotingResults.DiscardUnknown(m)
}

var xxx_messageInfo_VotingResults proto.InternalMessageInfo

func (m *VotingResults) GetTotalVotes() uint64 {
	if m != nil {
		return m.TotalVotes
	}
	return 0
}

func (m *VotingResults) GetTotalFairEnoughVotes() uint64 {
	if m != nil {
		return m.TotalFairEnoughVotes
	}
	return 0
}

func (m *VotingResults) GetTotalOverpoweredVotes() uint64 {
	if m != nil {
		return m.TotalOverpoweredVotes
	}
	return 0
}

func (m *VotingResults) GetTotalUnderpoweredVotes() uint64 {
	if m != nil {
		return m.TotalUnderpoweredVotes
	}
	return 0
}

func (m *VotingResults) GetTotalInappropriateVotes() uint64 {
	if m != nil {
		return m.TotalInappropriateVotes
	}
	return 0
}

func (m *VotingResults) GetCardResults() []*VotingResult {
	if m != nil {
		return m.CardResults
	}
	return nil
}

func (m *VotingResults) GetNotes() string {
	if m != nil {
		return m.Notes
	}
	return ""
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
	return fileDescriptor_693a4ceca03817bb, []int{1}
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
	CardId   uint64 `protobuf:"varint,1,opt,name=cardId,proto3" json:"cardId,omitempty"`
	VoteType string `protobuf:"bytes,2,opt,name=voteType,proto3" json:"voteType,omitempty"`
}

func (m *SingleVote) Reset()         { *m = SingleVote{} }
func (m *SingleVote) String() string { return proto.CompactTextString(m) }
func (*SingleVote) ProtoMessage()    {}
func (*SingleVote) Descriptor() ([]byte, []int) {
	return fileDescriptor_693a4ceca03817bb, []int{2}
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

func (m *SingleVote) GetVoteType() string {
	if m != nil {
		return m.VoteType
	}
	return ""
}

type VoteRight struct {
	CardId      uint64 `protobuf:"varint,1,opt,name=cardId,proto3" json:"cardId,omitempty"`
	ExpireBlock int64  `protobuf:"varint,2,opt,name=expireBlock,proto3" json:"expireBlock,omitempty"`
}

func (m *VoteRight) Reset()         { *m = VoteRight{} }
func (m *VoteRight) String() string { return proto.CompactTextString(m) }
func (*VoteRight) ProtoMessage()    {}
func (*VoteRight) Descriptor() ([]byte, []int) {
	return fileDescriptor_693a4ceca03817bb, []int{3}
}
func (m *VoteRight) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VoteRight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VoteRight.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VoteRight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteRight.Merge(m, src)
}
func (m *VoteRight) XXX_Size() int {
	return m.Size()
}
func (m *VoteRight) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteRight.DiscardUnknown(m)
}

var xxx_messageInfo_VoteRight proto.InternalMessageInfo

func (m *VoteRight) GetCardId() uint64 {
	if m != nil {
		return m.CardId
	}
	return 0
}

func (m *VoteRight) GetExpireBlock() int64 {
	if m != nil {
		return m.ExpireBlock
	}
	return 0
}

func init() {
	proto.RegisterType((*VotingResults)(nil), "DecentralCardGame.cardchain.cardchain.VotingResults")
	proto.RegisterType((*VotingResult)(nil), "DecentralCardGame.cardchain.cardchain.VotingResult")
	proto.RegisterType((*SingleVote)(nil), "DecentralCardGame.cardchain.cardchain.SingleVote")
	proto.RegisterType((*VoteRight)(nil), "DecentralCardGame.cardchain.cardchain.VoteRight")
}

func init() {
	proto.RegisterFile("DecentralCardGame/cardchain/cardchain/voteing.proto", fileDescriptor_693a4ceca03817bb)
}

var fileDescriptor_693a4ceca03817bb = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x9b, 0xed, 0x6e, 0xb4, 0xaf, 0x8a, 0x3a, 0xac, 0x6b, 0xf1, 0x10, 0x42, 0x40, 0x08,
	0x22, 0x59, 0xd8, 0x8a, 0xec, 0x51, 0x76, 0xad, 0xd2, 0x93, 0x30, 0xda, 0x1e, 0xbc, 0x4d, 0x93,
	0x31, 0x19, 0x4c, 0x67, 0x86, 0xc9, 0xa4, 0xb6, 0xdf, 0xc2, 0x8f, 0xe5, 0xb1, 0x47, 0x8f, 0xd2,
	0x7e, 0x03, 0x4f, 0x1e, 0x25, 0x93, 0x50, 0x43, 0x93, 0xc0, 0xde, 0xe6, 0xbd, 0xff, 0xfb, 0xbf,
	0xf0, 0xff, 0x91, 0x07, 0xe3, 0x77, 0x34, 0xa4, 0x5c, 0x2b, 0x92, 0xde, 0x12, 0x15, 0x7d, 0x20,
	0x4b, 0x7a, 0x19, 0x12, 0x15, 0x85, 0x09, 0x61, 0xbc, 0xf6, 0x5a, 0x09, 0x4d, 0x19, 0x8f, 0x03,
	0xa9, 0x84, 0x16, 0xe8, 0x45, 0xc3, 0x14, 0x1c, 0x46, 0xff, 0xbf, 0xbc, 0xbf, 0x27, 0xf0, 0x70,
	0x2e, 0x34, 0xe3, 0x31, 0xa6, 0x59, 0x9e, 0xea, 0x0c, 0x39, 0x00, 0x5a, 0x68, 0x92, 0xce, 0x85,
	0xa6, 0xd9, 0xc8, 0x72, 0x2d, 0xff, 0x14, 0xd7, 0x3a, 0xe8, 0x0a, 0xce, 0x4d, 0xf5, 0x9e, 0x30,
	0x35, 0xe1, 0x22, 0x8f, 0x93, 0x72, 0xf2, 0xc4, 0x4c, 0xb6, 0x6a, 0xe8, 0x35, 0x3c, 0x35, 0xfd,
	0x8f, 0x2b, 0xaa, 0xa4, 0xf8, 0x4e, 0x15, 0x8d, 0x4a, 0x53, 0xdf, 0x98, 0xda, 0x45, 0xf4, 0x06,
	0x2e, 0x8c, 0x30, 0xe3, 0xd1, 0x91, 0xed, 0xd4, 0xd8, 0x3a, 0x54, 0x74, 0x0d, 0xcf, 0x8c, 0x32,
	0xe5, 0x44, 0x4a, 0x25, 0xa4, 0x62, 0x44, 0xd3, 0xd2, 0x78, 0x66, 0x8c, 0x5d, 0x32, 0x9a, 0xc1,
	0xb0, 0x40, 0x53, 0xa1, 0x18, 0xd9, 0x6e, 0xdf, 0x1f, 0x5e, 0x8d, 0x83, 0x3b, 0xa1, 0x0c, 0xea,
	0x18, 0x71, 0x7d, 0x0f, 0x3a, 0x87, 0x33, 0x6e, 0x3e, 0x7f, 0xcf, 0xb5, 0xfc, 0x01, 0x2e, 0x0b,
	0xef, 0x8f, 0x05, 0x0f, 0xea, 0x1e, 0x74, 0x01, 0x76, 0xe1, 0x9a, 0x46, 0x15, 0xf5, 0xaa, 0x42,
	0x3e, 0x3c, 0xfa, 0xda, 0x0a, 0xfb, 0xb8, 0x8d, 0x5e, 0xc2, 0x63, 0xd1, 0x8e, 0xb8, 0xd1, 0x47,
	0xaf, 0xe0, 0x49, 0xde, 0x01, 0xb6, 0x29, 0xa0, 0x00, 0x10, 0xeb, 0xc2, 0xd9, 0xa2, 0x14, 0x59,
	0x94, 0x49, 0x35, 0xb2, 0x4d, 0xe6, 0xaa, 0xf2, 0xde, 0x02, 0x7c, 0x62, 0x3c, 0x4e, 0xcd, 0x58,
	0x67, 0xe2, 0xe7, 0x70, 0xbf, 0xf8, 0x9b, 0x3f, 0x6f, 0x24, 0x35, 0x51, 0x07, 0xf8, 0x50, 0x7b,
	0x13, 0x18, 0x14, 0x5e, 0xcc, 0xe2, 0xa4, 0x1b, 0x99, 0x0b, 0x43, 0xba, 0x96, 0x4c, 0xd1, 0x9b,
	0x54, 0x84, 0xdf, 0xcc, 0x8e, 0x3e, 0xae, 0xb7, 0x6e, 0xf0, 0xcf, 0x9d, 0x63, 0x6d, 0x77, 0x8e,
	0xf5, 0x7b, 0xe7, 0x58, 0x3f, 0xf6, 0x4e, 0x6f, 0xbb, 0x77, 0x7a, 0xbf, 0xf6, 0x4e, 0xef, 0xcb,
	0x75, 0xcc, 0x74, 0x92, 0x2f, 0x82, 0x50, 0x2c, 0x2f, 0x9b, 0x97, 0x77, 0x7b, 0xb8, 0xb7, 0x75,
	0xed, 0xf6, 0xf4, 0x46, 0xd2, 0x6c, 0x61, 0x9b, 0xd3, 0x1b, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff,
	0xfa, 0x54, 0xe5, 0xf3, 0xb1, 0x03, 0x00, 0x00,
}

func (m *VotingResults) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VotingResults) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VotingResults) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Notes) > 0 {
		i -= len(m.Notes)
		copy(dAtA[i:], m.Notes)
		i = encodeVarintVoteing(dAtA, i, uint64(len(m.Notes)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.CardResults) > 0 {
		for iNdEx := len(m.CardResults) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CardResults[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVoteing(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.TotalInappropriateVotes != 0 {
		i = encodeVarintVoteing(dAtA, i, uint64(m.TotalInappropriateVotes))
		i--
		dAtA[i] = 0x28
	}
	if m.TotalUnderpoweredVotes != 0 {
		i = encodeVarintVoteing(dAtA, i, uint64(m.TotalUnderpoweredVotes))
		i--
		dAtA[i] = 0x20
	}
	if m.TotalOverpoweredVotes != 0 {
		i = encodeVarintVoteing(dAtA, i, uint64(m.TotalOverpoweredVotes))
		i--
		dAtA[i] = 0x18
	}
	if m.TotalFairEnoughVotes != 0 {
		i = encodeVarintVoteing(dAtA, i, uint64(m.TotalFairEnoughVotes))
		i--
		dAtA[i] = 0x10
	}
	if m.TotalVotes != 0 {
		i = encodeVarintVoteing(dAtA, i, uint64(m.TotalVotes))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
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
		i = encodeVarintVoteing(dAtA, i, uint64(len(m.Result)))
		i--
		dAtA[i] = 0x32
	}
	if m.InappropriateVotes != 0 {
		i = encodeVarintVoteing(dAtA, i, uint64(m.InappropriateVotes))
		i--
		dAtA[i] = 0x28
	}
	if m.UnderpoweredVotes != 0 {
		i = encodeVarintVoteing(dAtA, i, uint64(m.UnderpoweredVotes))
		i--
		dAtA[i] = 0x20
	}
	if m.OverpoweredVotes != 0 {
		i = encodeVarintVoteing(dAtA, i, uint64(m.OverpoweredVotes))
		i--
		dAtA[i] = 0x18
	}
	if m.FairEnoughVotes != 0 {
		i = encodeVarintVoteing(dAtA, i, uint64(m.FairEnoughVotes))
		i--
		dAtA[i] = 0x10
	}
	if m.CardId != 0 {
		i = encodeVarintVoteing(dAtA, i, uint64(m.CardId))
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
	if len(m.VoteType) > 0 {
		i -= len(m.VoteType)
		copy(dAtA[i:], m.VoteType)
		i = encodeVarintVoteing(dAtA, i, uint64(len(m.VoteType)))
		i--
		dAtA[i] = 0x12
	}
	if m.CardId != 0 {
		i = encodeVarintVoteing(dAtA, i, uint64(m.CardId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *VoteRight) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VoteRight) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VoteRight) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ExpireBlock != 0 {
		i = encodeVarintVoteing(dAtA, i, uint64(m.ExpireBlock))
		i--
		dAtA[i] = 0x10
	}
	if m.CardId != 0 {
		i = encodeVarintVoteing(dAtA, i, uint64(m.CardId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintVoteing(dAtA []byte, offset int, v uint64) int {
	offset -= sovVoteing(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *VotingResults) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TotalVotes != 0 {
		n += 1 + sovVoteing(uint64(m.TotalVotes))
	}
	if m.TotalFairEnoughVotes != 0 {
		n += 1 + sovVoteing(uint64(m.TotalFairEnoughVotes))
	}
	if m.TotalOverpoweredVotes != 0 {
		n += 1 + sovVoteing(uint64(m.TotalOverpoweredVotes))
	}
	if m.TotalUnderpoweredVotes != 0 {
		n += 1 + sovVoteing(uint64(m.TotalUnderpoweredVotes))
	}
	if m.TotalInappropriateVotes != 0 {
		n += 1 + sovVoteing(uint64(m.TotalInappropriateVotes))
	}
	if len(m.CardResults) > 0 {
		for _, e := range m.CardResults {
			l = e.Size()
			n += 1 + l + sovVoteing(uint64(l))
		}
	}
	l = len(m.Notes)
	if l > 0 {
		n += 1 + l + sovVoteing(uint64(l))
	}
	return n
}

func (m *VotingResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CardId != 0 {
		n += 1 + sovVoteing(uint64(m.CardId))
	}
	if m.FairEnoughVotes != 0 {
		n += 1 + sovVoteing(uint64(m.FairEnoughVotes))
	}
	if m.OverpoweredVotes != 0 {
		n += 1 + sovVoteing(uint64(m.OverpoweredVotes))
	}
	if m.UnderpoweredVotes != 0 {
		n += 1 + sovVoteing(uint64(m.UnderpoweredVotes))
	}
	if m.InappropriateVotes != 0 {
		n += 1 + sovVoteing(uint64(m.InappropriateVotes))
	}
	l = len(m.Result)
	if l > 0 {
		n += 1 + l + sovVoteing(uint64(l))
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
		n += 1 + sovVoteing(uint64(m.CardId))
	}
	l = len(m.VoteType)
	if l > 0 {
		n += 1 + l + sovVoteing(uint64(l))
	}
	return n
}

func (m *VoteRight) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CardId != 0 {
		n += 1 + sovVoteing(uint64(m.CardId))
	}
	if m.ExpireBlock != 0 {
		n += 1 + sovVoteing(uint64(m.ExpireBlock))
	}
	return n
}

func sovVoteing(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVoteing(x uint64) (n int) {
	return sovVoteing(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VotingResults) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVoteing
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
			return fmt.Errorf("proto: VotingResults: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VotingResults: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalVotes", wireType)
			}
			m.TotalVotes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoteing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalVotes |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalFairEnoughVotes", wireType)
			}
			m.TotalFairEnoughVotes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoteing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalFairEnoughVotes |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalOverpoweredVotes", wireType)
			}
			m.TotalOverpoweredVotes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoteing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalOverpoweredVotes |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalUnderpoweredVotes", wireType)
			}
			m.TotalUnderpoweredVotes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoteing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalUnderpoweredVotes |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalInappropriateVotes", wireType)
			}
			m.TotalInappropriateVotes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoteing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalInappropriateVotes |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CardResults", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoteing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthVoteing
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVoteing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CardResults = append(m.CardResults, &VotingResult{})
			if err := m.CardResults[len(m.CardResults)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Notes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoteing
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
				return ErrInvalidLengthVoteing
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVoteing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Notes = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVoteing(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVoteing
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
func (m *VotingResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVoteing
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
					return ErrIntOverflowVoteing
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
					return ErrIntOverflowVoteing
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
					return ErrIntOverflowVoteing
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
					return ErrIntOverflowVoteing
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
					return ErrIntOverflowVoteing
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
					return ErrIntOverflowVoteing
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
				return ErrInvalidLengthVoteing
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVoteing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Result = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVoteing(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVoteing
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
				return ErrIntOverflowVoteing
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
					return ErrIntOverflowVoteing
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoteing
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
				return ErrInvalidLengthVoteing
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVoteing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VoteType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVoteing(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVoteing
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
func (m *VoteRight) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVoteing
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
			return fmt.Errorf("proto: VoteRight: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VoteRight: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CardId", wireType)
			}
			m.CardId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoteing
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
				return fmt.Errorf("proto: wrong wireType = %d for field ExpireBlock", wireType)
			}
			m.ExpireBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoteing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpireBlock |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipVoteing(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVoteing
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
func skipVoteing(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVoteing
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
					return 0, ErrIntOverflowVoteing
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
					return 0, ErrIntOverflowVoteing
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
				return 0, ErrInvalidLengthVoteing
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVoteing
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVoteing
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVoteing        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVoteing          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVoteing = fmt.Errorf("proto: unexpected end of group")
)