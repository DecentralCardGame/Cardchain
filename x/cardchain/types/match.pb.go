// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cardchain/cardchain/match.proto

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
	return fileDescriptor_f49b609716516c9f, []int{0}
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
	return fileDescriptor_f49b609716516c9f, []int{0}
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
	return fileDescriptor_f49b609716516c9f, []int{1}
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

func init() { proto.RegisterFile("cardchain/cardchain/match.proto", fileDescriptor_f49b609716516c9f) }

var fileDescriptor_f49b609716516c9f = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xcf, 0x8b, 0xd3, 0x40,
	0x14, 0xee, 0x24, 0xd9, 0x4d, 0xf7, 0x05, 0xb4, 0xcc, 0x69, 0x58, 0x24, 0x0e, 0x0b, 0x42, 0xd8,
	0x43, 0x8a, 0xd5, 0x83, 0xd7, 0xfe, 0x00, 0x3d, 0x28, 0xea, 0x08, 0x0a, 0xde, 0xa6, 0x33, 0x63,
	0x3b, 0xd8, 0x64, 0xc2, 0x64, 0x5a, 0xdd, 0xff, 0x42, 0xfc, 0xab, 0x3c, 0xee, 0xd1, 0xa3, 0xb4,
	0x7f, 0x87, 0x20, 0x99, 0xb5, 0x4d, 0x70, 0x3d, 0x14, 0xf6, 0xf6, 0xe5, 0xe3, 0xbd, 0xef, 0xfb,
	0xf2, 0xde, 0x3c, 0x78, 0x28, 0xb8, 0x95, 0x62, 0xc9, 0x75, 0x39, 0x6c, 0x51, 0xc1, 0x9d, 0x58,
	0xe6, 0x95, 0x35, 0xce, 0xe0, 0x47, 0x33, 0x25, 0x54, 0xe9, 0x2c, 0x5f, 0x4d, 0xb9, 0x95, 0xcf,
	0x79, 0xa1, 0xf2, 0x43, 0x61, 0x8b, 0xce, 0xe9, 0xff, 0x74, 0x36, 0xc6, 0xe9, 0x72, 0x71, 0x23,
	0x74, 0xf1, 0x3b, 0x80, 0x93, 0x57, 0x8d, 0x30, 0x7e, 0x00, 0x67, 0x4e, 0x17, 0xaa, 0x76, 0xbc,
	0xa8, 0x08, 0xa2, 0x28, 0x8b, 0x58, 0x4b, 0xe0, 0x73, 0xe8, 0x5b, 0x55, 0x19, 0xeb, 0x94, 0x25,
	0x01, 0x45, 0xd9, 0x19, 0x3b, 0x7c, 0xe3, 0x97, 0x10, 0x57, 0x2b, 0x7e, 0xa5, 0xec, 0x98, 0x84,
	0x14, 0x65, 0xc9, 0x68, 0x94, 0x1f, 0x15, 0x2f, 0xf7, 0xc6, 0x6f, 0x7c, 0x2b, 0xdb, 0x4b, 0xb4,
	0x6a, 0x13, 0x12, 0xdd, 0x55, 0x6d, 0x82, 0x5f, 0x40, 0x6c, 0xd6, 0x4e, 0x98, 0x42, 0x91, 0x98,
	0xa2, 0xec, 0xde, 0x28, 0x3f, 0x52, 0xed, 0xf5, 0x4d, 0x17, 0xdb, 0xb7, 0xe3, 0x4b, 0x18, 0x08,
	0xa3, 0xcb, 0x7a, 0xa6, 0x6b, 0x67, 0xf5, 0x7c, 0xed, 0x94, 0x24, 0x40, 0x51, 0xd6, 0x67, 0xb7,
	0x78, 0x9c, 0xc1, 0xfd, 0x5a, 0xd9, 0x8d, 0xb2, 0x53, 0x53, 0x7e, 0xd2, 0xb6, 0x50, 0x92, 0xf4,
	0x7d, 0xe9, 0xbf, 0xf4, 0xc5, 0xf7, 0x00, 0x92, 0x4e, 0x70, 0x8c, 0x21, 0xe2, 0x52, 0x5a, 0xbf,
	0x80, 0x33, 0xe6, 0x31, 0xa6, 0x90, 0xf8, 0xdf, 0x91, 0x4d, 0xe0, 0x9a, 0x04, 0x34, 0xcc, 0x22,
	0xd6, 0xa5, 0x9a, 0xdd, 0x89, 0x83, 0x53, 0xe8, 0x9d, 0x5a, 0xa2, 0x3b, 0x83, 0xe8, 0x6e, 0x33,
	0xc0, 0x10, 0x49, 0x25, 0x3e, 0x93, 0x13, 0x1f, 0xc1, 0x63, 0xfc, 0x16, 0x60, 0x63, 0xdc, 0x3e,
	0xdc, 0x29, 0x0d, 0xb3, 0x64, 0xf4, 0xf8, 0x48, 0x83, 0x77, 0xba, 0x5c, 0xac, 0xd4, 0x7b, 0xe3,
	0x14, 0xeb, 0x88, 0x5c, 0x3e, 0x85, 0xf8, 0xaf, 0x35, 0xee, 0x43, 0x34, 0xfe, 0x60, 0xca, 0x41,
	0xaf, 0x41, 0x93, 0x06, 0xa1, 0x06, 0xcd, 0x2c, 0xff, 0x32, 0x08, 0x70, 0x02, 0xf1, 0x78, 0xde,
	0x3c, 0x42, 0x39, 0x08, 0x27, 0xec, 0xc7, 0x36, 0x45, 0xd7, 0xdb, 0x14, 0xfd, 0xda, 0xa6, 0xe8,
	0xdb, 0x2e, 0xed, 0x5d, 0xef, 0xd2, 0xde, 0xcf, 0x5d, 0xda, 0xfb, 0xf8, 0x6c, 0xa1, 0xdd, 0x72,
	0x3d, 0xcf, 0x85, 0x29, 0x86, 0xb7, 0x82, 0x0d, 0xa7, 0x87, 0xcb, 0xf8, 0xda, 0xb9, 0x12, 0x77,
	0x55, 0xa9, 0x7a, 0x7e, 0xea, 0xaf, 0xe4, 0xc9, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xea, 0x66,
	0xb9, 0x36, 0x91, 0x03, 0x00, 0x00,
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
