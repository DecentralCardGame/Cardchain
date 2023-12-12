// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: DecentralCardGame/cardchain/cardchain/match.proto

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

type Outcome int32

const (
	Outcome_AWon    Outcome = 0
	Outcome_BWon    Outcome = 1
	Outcome_Draw    Outcome = 2
	Outcome_Aborted Outcome = 3
)

var Outcome_name = map[int32]string{
	0: "AWon",
	1: "BWon",
	2: "Draw",
	3: "Aborted",
}

var Outcome_value = map[string]int32{
	"AWon":    0,
	"BWon":    1,
	"Draw":    2,
	"Aborted": 3,
}

func (x Outcome) String() string {
	return proto.EnumName(Outcome_name, int32(x))
}

func (Outcome) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6da747430368d591, []int{0}
}

type Match struct {
	Timestamp        uint64       `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Reporter         string       `protobuf:"bytes,2,opt,name=reporter,proto3" json:"reporter,omitempty"`
	PlayerA          *MatchPlayer `protobuf:"bytes,3,opt,name=playerA,proto3" json:"playerA,omitempty"`
	PlayerB          *MatchPlayer `protobuf:"bytes,4,opt,name=playerB,proto3" json:"playerB,omitempty"`
	Outcome          Outcome      `protobuf:"varint,7,opt,name=outcome,proto3,enum=DecentralCardGame.cardchain.cardchain.Outcome" json:"outcome,omitempty"`
	CoinsDistributed bool         `protobuf:"varint,10,opt,name=coinsDistributed,proto3" json:"coinsDistributed,omitempty"`
	ServerConfirmed  bool         `protobuf:"varint,8,opt,name=serverConfirmed,proto3" json:"serverConfirmed,omitempty"`
}

func (m *Match) Reset()         { *m = Match{} }
func (m *Match) String() string { return proto.CompactTextString(m) }
func (*Match) ProtoMessage()    {}
func (*Match) Descriptor() ([]byte, []int) {
	return fileDescriptor_6da747430368d591, []int{0}
}
func (m *Match) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Match) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Match.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Match) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Match.Merge(m, src)
}
func (m *Match) XXX_Size() int {
	return m.Size()
}
func (m *Match) XXX_DiscardUnknown() {
	xxx_messageInfo_Match.DiscardUnknown(m)
}

var xxx_messageInfo_Match proto.InternalMessageInfo

func (m *Match) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Match) GetReporter() string {
	if m != nil {
		return m.Reporter
	}
	return ""
}

func (m *Match) GetPlayerA() *MatchPlayer {
	if m != nil {
		return m.PlayerA
	}
	return nil
}

func (m *Match) GetPlayerB() *MatchPlayer {
	if m != nil {
		return m.PlayerB
	}
	return nil
}

func (m *Match) GetOutcome() Outcome {
	if m != nil {
		return m.Outcome
	}
	return Outcome_AWon
}

func (m *Match) GetCoinsDistributed() bool {
	if m != nil {
		return m.CoinsDistributed
	}
	return false
}

func (m *Match) GetServerConfirmed() bool {
	if m != nil {
		return m.ServerConfirmed
	}
	return false
}

type MatchPlayer struct {
	Addr        string        `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	PlayedCards []uint64      `protobuf:"varint,2,rep,packed,name=playedCards,proto3" json:"playedCards,omitempty"`
	Confirmed   bool          `protobuf:"varint,3,opt,name=confirmed,proto3" json:"confirmed,omitempty"`
	Outcome     Outcome       `protobuf:"varint,4,opt,name=outcome,proto3,enum=DecentralCardGame.cardchain.cardchain.Outcome" json:"outcome,omitempty"`
	Deck        []uint64      `protobuf:"varint,5,rep,packed,name=deck,proto3" json:"deck,omitempty"`
	VotedCards  []*SingleVote `protobuf:"bytes,6,rep,name=votedCards,proto3" json:"votedCards,omitempty"`
}

func (m *MatchPlayer) Reset()         { *m = MatchPlayer{} }
func (m *MatchPlayer) String() string { return proto.CompactTextString(m) }
func (*MatchPlayer) ProtoMessage()    {}
func (*MatchPlayer) Descriptor() ([]byte, []int) {
	return fileDescriptor_6da747430368d591, []int{1}
}
func (m *MatchPlayer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MatchPlayer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MatchPlayer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MatchPlayer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchPlayer.Merge(m, src)
}
func (m *MatchPlayer) XXX_Size() int {
	return m.Size()
}
func (m *MatchPlayer) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchPlayer.DiscardUnknown(m)
}

var xxx_messageInfo_MatchPlayer proto.InternalMessageInfo

func (m *MatchPlayer) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *MatchPlayer) GetPlayedCards() []uint64 {
	if m != nil {
		return m.PlayedCards
	}
	return nil
}

func (m *MatchPlayer) GetConfirmed() bool {
	if m != nil {
		return m.Confirmed
	}
	return false
}

func (m *MatchPlayer) GetOutcome() Outcome {
	if m != nil {
		return m.Outcome
	}
	return Outcome_AWon
}

func (m *MatchPlayer) GetDeck() []uint64 {
	if m != nil {
		return m.Deck
	}
	return nil
}

func (m *MatchPlayer) GetVotedCards() []*SingleVote {
	if m != nil {
		return m.VotedCards
	}
	return nil
}

func init() {
	proto.RegisterEnum("DecentralCardGame.cardchain.cardchain.Outcome", Outcome_name, Outcome_value)
	proto.RegisterType((*Match)(nil), "DecentralCardGame.cardchain.cardchain.Match")
	proto.RegisterType((*MatchPlayer)(nil), "DecentralCardGame.cardchain.cardchain.MatchPlayer")
}

func init() {
	proto.RegisterFile("DecentralCardGame/cardchain/cardchain/match.proto", fileDescriptor_6da747430368d591)
}

var fileDescriptor_6da747430368d591 = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x4f, 0x8b, 0xd3, 0x40,
	0x1c, 0xed, 0x24, 0xd9, 0x4d, 0xfb, 0x0b, 0x68, 0x99, 0xd3, 0xb0, 0x48, 0x08, 0x0b, 0x42, 0xd8,
	0x43, 0x96, 0x8d, 0x1e, 0xbc, 0xf6, 0x0f, 0xe8, 0x41, 0x51, 0x47, 0x50, 0xf0, 0x36, 0x9d, 0x19,
	0xdb, 0xc1, 0x26, 0x13, 0x26, 0xd3, 0xea, 0x7e, 0x0b, 0xf1, 0x53, 0x79, 0xdc, 0xa3, 0x47, 0x69,
	0x3f, 0x87, 0x20, 0x99, 0xb5, 0x6d, 0xd8, 0x5e, 0x02, 0xbd, 0xbd, 0x3c, 0xe6, 0xf7, 0xde, 0xcb,
	0xfc, 0xe6, 0xc1, 0xcd, 0x54, 0x72, 0x59, 0x5a, 0xc3, 0x96, 0x13, 0x66, 0xc4, 0x4b, 0x56, 0xc8,
	0x6b, 0xce, 0x8c, 0xe0, 0x0b, 0xa6, 0xca, 0x16, 0x2a, 0x98, 0xe5, 0x8b, 0xac, 0x32, 0xda, 0x6a,
	0xfc, 0xf4, 0x68, 0x24, 0xdb, 0x1f, 0x3c, 0xa0, 0x8b, 0xbc, 0x9b, 0xf2, 0x5a, 0x5b, 0x55, 0xce,
	0xef, 0xa5, 0x2f, 0xff, 0x7a, 0x70, 0xf6, 0xa6, 0xb1, 0xc2, 0x4f, 0x60, 0x60, 0x55, 0x21, 0x6b,
	0xcb, 0x8a, 0x8a, 0xa0, 0x04, 0xa5, 0x01, 0x3d, 0x10, 0xf8, 0x02, 0xfa, 0x46, 0x56, 0xda, 0x58,
	0x69, 0x88, 0x97, 0xa0, 0x74, 0x40, 0xf7, 0xdf, 0xf8, 0x35, 0x84, 0xd5, 0x92, 0xdd, 0x4a, 0x33,
	0x22, 0x7e, 0x82, 0xd2, 0x28, 0xcf, 0xb3, 0x4e, 0x81, 0x33, 0x67, 0xfc, 0xce, 0x8d, 0xd2, 0x9d,
	0xc4, 0x41, 0x6d, 0x4c, 0x82, 0x53, 0xd5, 0xc6, 0xf8, 0x15, 0x84, 0x7a, 0x65, 0xb9, 0x2e, 0x24,
	0x09, 0x13, 0x94, 0x3e, 0xca, 0xb3, 0x8e, 0x6a, 0x6f, 0xef, 0xa7, 0xe8, 0x6e, 0x1c, 0x5f, 0xc1,
	0x90, 0x6b, 0x55, 0xd6, 0x53, 0x55, 0x5b, 0xa3, 0x66, 0x2b, 0x2b, 0x05, 0x81, 0x04, 0xa5, 0x7d,
	0x7a, 0xc4, 0xe3, 0x14, 0x1e, 0xd7, 0xd2, 0xac, 0xa5, 0x99, 0xe8, 0xf2, 0x8b, 0x32, 0x85, 0x14,
	0xa4, 0xef, 0x8e, 0x3e, 0xa4, 0x2f, 0x7f, 0x7a, 0x10, 0xb5, 0x82, 0x63, 0x0c, 0x01, 0x13, 0xc2,
	0xb8, 0x05, 0x0c, 0xa8, 0xc3, 0x38, 0x81, 0xc8, 0xfd, 0x8e, 0x68, 0x02, 0xd7, 0xc4, 0x4b, 0xfc,
	0x34, 0xa0, 0x6d, 0xaa, 0xd9, 0x1d, 0xdf, 0x3b, 0xf9, 0xce, 0xe9, 0x40, 0xb4, 0xef, 0x20, 0x38,
	0xed, 0x0e, 0x30, 0x04, 0x42, 0xf2, 0xaf, 0xe4, 0xcc, 0x45, 0x70, 0x18, 0xbf, 0x07, 0x58, 0x6b,
	0xbb, 0x0b, 0x77, 0x9e, 0xf8, 0x69, 0x94, 0xdf, 0x74, 0x34, 0xf8, 0xa0, 0xca, 0xf9, 0x52, 0x7e,
	0xd4, 0x56, 0xd2, 0x96, 0xc8, 0xd5, 0x73, 0x08, 0xff, 0x5b, 0xe3, 0x3e, 0x04, 0xa3, 0x4f, 0xba,
	0x1c, 0xf6, 0x1a, 0x34, 0x6e, 0x10, 0x6a, 0xd0, 0xd4, 0xb0, 0x6f, 0x43, 0x0f, 0x47, 0x10, 0x8e,
	0x66, 0xcd, 0x23, 0x14, 0x43, 0x7f, 0x4c, 0x7f, 0x6d, 0x62, 0x74, 0xb7, 0x89, 0xd1, 0x9f, 0x4d,
	0x8c, 0x7e, 0x6c, 0xe3, 0xde, 0xdd, 0x36, 0xee, 0xfd, 0xde, 0xc6, 0xbd, 0xcf, 0x2f, 0xe6, 0xca,
	0x2e, 0x56, 0xb3, 0x8c, 0xeb, 0xe2, 0xfa, 0xb8, 0x23, 0x93, 0x7d, 0x33, 0xbe, 0xb7, 0x5a, 0x62,
	0x6f, 0x2b, 0x59, 0xcf, 0xce, 0x5d, 0x4b, 0x9e, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x78, 0x51,
	0x2f, 0x04, 0xb5, 0x03, 0x00, 0x00,
}

func (m *Match) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Match) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Match) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CoinsDistributed {
		i--
		if m.CoinsDistributed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x50
	}
	if m.ServerConfirmed {
		i--
		if m.ServerConfirmed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if m.Outcome != 0 {
		i = encodeVarintMatch(dAtA, i, uint64(m.Outcome))
		i--
		dAtA[i] = 0x38
	}
	if m.PlayerB != nil {
		{
			size, err := m.PlayerB.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMatch(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.PlayerA != nil {
		{
			size, err := m.PlayerA.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMatch(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Reporter) > 0 {
		i -= len(m.Reporter)
		copy(dAtA[i:], m.Reporter)
		i = encodeVarintMatch(dAtA, i, uint64(len(m.Reporter)))
		i--
		dAtA[i] = 0x12
	}
	if m.Timestamp != 0 {
		i = encodeVarintMatch(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MatchPlayer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MatchPlayer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MatchPlayer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.VotedCards) > 0 {
		for iNdEx := len(m.VotedCards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VotedCards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMatch(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Deck) > 0 {
		dAtA4 := make([]byte, len(m.Deck)*10)
		var j3 int
		for _, num := range m.Deck {
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		i -= j3
		copy(dAtA[i:], dAtA4[:j3])
		i = encodeVarintMatch(dAtA, i, uint64(j3))
		i--
		dAtA[i] = 0x2a
	}
	if m.Outcome != 0 {
		i = encodeVarintMatch(dAtA, i, uint64(m.Outcome))
		i--
		dAtA[i] = 0x20
	}
	if m.Confirmed {
		i--
		if m.Confirmed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.PlayedCards) > 0 {
		dAtA6 := make([]byte, len(m.PlayedCards)*10)
		var j5 int
		for _, num := range m.PlayedCards {
			for num >= 1<<7 {
				dAtA6[j5] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j5++
			}
			dAtA6[j5] = uint8(num)
			j5++
		}
		i -= j5
		copy(dAtA[i:], dAtA6[:j5])
		i = encodeVarintMatch(dAtA, i, uint64(j5))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Addr) > 0 {
		i -= len(m.Addr)
		copy(dAtA[i:], m.Addr)
		i = encodeVarintMatch(dAtA, i, uint64(len(m.Addr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMatch(dAtA []byte, offset int, v uint64) int {
	offset -= sovMatch(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Match) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Timestamp != 0 {
		n += 1 + sovMatch(uint64(m.Timestamp))
	}
	l = len(m.Reporter)
	if l > 0 {
		n += 1 + l + sovMatch(uint64(l))
	}
	if m.PlayerA != nil {
		l = m.PlayerA.Size()
		n += 1 + l + sovMatch(uint64(l))
	}
	if m.PlayerB != nil {
		l = m.PlayerB.Size()
		n += 1 + l + sovMatch(uint64(l))
	}
	if m.Outcome != 0 {
		n += 1 + sovMatch(uint64(m.Outcome))
	}
	if m.ServerConfirmed {
		n += 2
	}
	if m.CoinsDistributed {
		n += 2
	}
	return n
}

func (m *MatchPlayer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Addr)
	if l > 0 {
		n += 1 + l + sovMatch(uint64(l))
	}
	if len(m.PlayedCards) > 0 {
		l = 0
		for _, e := range m.PlayedCards {
			l += sovMatch(uint64(e))
		}
		n += 1 + sovMatch(uint64(l)) + l
	}
	if m.Confirmed {
		n += 2
	}
	if m.Outcome != 0 {
		n += 1 + sovMatch(uint64(m.Outcome))
	}
	if len(m.Deck) > 0 {
		l = 0
		for _, e := range m.Deck {
			l += sovMatch(uint64(e))
		}
		n += 1 + sovMatch(uint64(l)) + l
	}
	if len(m.VotedCards) > 0 {
		for _, e := range m.VotedCards {
			l = e.Size()
			n += 1 + l + sovMatch(uint64(l))
		}
	}
	return n
}

func sovMatch(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMatch(x uint64) (n int) {
	return sovMatch(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Match) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMatch
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
			return fmt.Errorf("proto: Match: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Match: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reporter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMatch
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
				return ErrInvalidLengthMatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reporter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlayerA", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMatch
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
				return ErrInvalidLengthMatch
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PlayerA == nil {
				m.PlayerA = &MatchPlayer{}
			}
			if err := m.PlayerA.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlayerB", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMatch
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
				return ErrInvalidLengthMatch
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PlayerB == nil {
				m.PlayerB = &MatchPlayer{}
			}
			if err := m.PlayerB.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Outcome", wireType)
			}
			m.Outcome = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Outcome |= Outcome(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerConfirmed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ServerConfirmed = bool(v != 0)
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoinsDistributed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CoinsDistributed = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipMatch(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMatch
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
func (m *MatchPlayer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMatch
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
			return fmt.Errorf("proto: MatchPlayer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MatchPlayer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMatch
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
				return ErrInvalidLengthMatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMatch
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.PlayedCards = append(m.PlayedCards, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMatch
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthMatch
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthMatch
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.PlayedCards) == 0 {
					m.PlayedCards = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMatch
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.PlayedCards = append(m.PlayedCards, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field PlayedCards", wireType)
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Confirmed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Confirmed = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Outcome", wireType)
			}
			m.Outcome = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Outcome |= Outcome(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMatch
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Deck = append(m.Deck, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMatch
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthMatch
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthMatch
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Deck) == 0 {
					m.Deck = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMatch
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Deck = append(m.Deck, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Deck", wireType)
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotedCards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMatch
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
				return ErrInvalidLengthMatch
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VotedCards = append(m.VotedCards, &SingleVote{})
			if err := m.VotedCards[len(m.VotedCards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMatch(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMatch
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
func skipMatch(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMatch
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
					return 0, ErrIntOverflowMatch
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
					return 0, ErrIntOverflowMatch
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
				return 0, ErrInvalidLengthMatch
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMatch
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMatch
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMatch        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMatch          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMatch = fmt.Errorf("proto: unexpected end of group")
)
