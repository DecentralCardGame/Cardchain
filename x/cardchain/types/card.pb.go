// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cardchain/cardchain/card.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

type Status int32

const (
	Status_scheme         Status = 0
	Status_prototype      Status = 1
	Status_trial          Status = 2
	Status_permanent      Status = 3
	Status_suspended      Status = 4
	Status_banned         Status = 5
	Status_bannedSoon     Status = 6
	Status_bannedVerySoon Status = 7
	Status_none           Status = 8
	Status_adventureItem  Status = 9
)

var Status_name = map[int32]string{
	0: "scheme",
	1: "prototype",
	2: "trial",
	3: "permanent",
	4: "suspended",
	5: "banned",
	6: "bannedSoon",
	7: "bannedVerySoon",
	8: "none",
	9: "adventureItem",
}

var Status_value = map[string]int32{
	"scheme":         0,
	"prototype":      1,
	"trial":          2,
	"permanent":      3,
	"suspended":      4,
	"banned":         5,
	"bannedSoon":     6,
	"bannedVerySoon": 7,
	"none":           8,
	"adventureItem":  9,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a360ffd2377ddc30, []int{0}
}

type CardRarity int32

const (
	CardRarity_common      CardRarity = 0
	CardRarity_uncommon    CardRarity = 1
	CardRarity_rare        CardRarity = 2
	CardRarity_exceptional CardRarity = 3
	CardRarity_unique      CardRarity = 4
)

var CardRarity_name = map[int32]string{
	0: "common",
	1: "uncommon",
	2: "rare",
	3: "exceptional",
	4: "unique",
}

var CardRarity_value = map[string]int32{
	"common":      0,
	"uncommon":    1,
	"rare":        2,
	"exceptional": 3,
	"unique":      4,
}

func (x CardRarity) String() string {
	return proto.EnumName(CardRarity_name, int32(x))
}

func (CardRarity) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a360ffd2377ddc30, []int{1}
}

type CardClass int32

const (
	CardClass_nature     CardClass = 0
	CardClass_culture    CardClass = 1
	CardClass_mysticism  CardClass = 2
	CardClass_technology CardClass = 3
)

var CardClass_name = map[int32]string{
	0: "nature",
	1: "culture",
	2: "mysticism",
	3: "technology",
}

var CardClass_value = map[string]int32{
	"nature":     0,
	"culture":    1,
	"mysticism":  2,
	"technology": 3,
}

func (x CardClass) String() string {
	return proto.EnumName(CardClass_name, int32(x))
}

func (CardClass) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a360ffd2377ddc30, []int{2}
}

type CardType int32

const (
	CardType_place       CardType = 0
	CardType_action      CardType = 1
	CardType_entity      CardType = 2
	CardType_headquarter CardType = 3
)

var CardType_name = map[int32]string{
	0: "place",
	1: "action",
	2: "entity",
	3: "headquarter",
}

var CardType_value = map[string]int32{
	"place":       0,
	"action":      1,
	"entity":      2,
	"headquarter": 3,
}

func (x CardType) String() string {
	return proto.EnumName(CardType_name, int32(x))
}

func (CardType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a360ffd2377ddc30, []int{3}
}

type Card struct {
	Owner              string     `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Artist             string     `protobuf:"bytes,2,opt,name=artist,proto3" json:"artist,omitempty"`
	Content            []byte     `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	ImageId            uint64     `protobuf:"varint,4,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
	FullArt            bool       `protobuf:"varint,5,opt,name=fullArt,proto3" json:"fullArt,omitempty"`
	Notes              string     `protobuf:"bytes,6,opt,name=notes,proto3" json:"notes,omitempty"`
	Status             Status     `protobuf:"varint,7,opt,name=status,proto3,enum=cardchain.cardchain.Status" json:"status,omitempty"`
	VotePool           types.Coin `protobuf:"bytes,8,opt,name=votePool,proto3" json:"votePool"`
	Voters             []string   `protobuf:"bytes,14,rep,name=voters,proto3" json:"voters,omitempty"`
	FairEnoughVotes    uint64     `protobuf:"varint,9,opt,name=fairEnoughVotes,proto3" json:"fairEnoughVotes,omitempty"`
	OverpoweredVotes   uint64     `protobuf:"varint,10,opt,name=overpoweredVotes,proto3" json:"overpoweredVotes,omitempty"`
	UnderpoweredVotes  uint64     `protobuf:"varint,11,opt,name=underpoweredVotes,proto3" json:"underpoweredVotes,omitempty"`
	InappropriateVotes uint64     `protobuf:"varint,12,opt,name=inappropriateVotes,proto3" json:"inappropriateVotes,omitempty"`
	Nerflevel          int64      `protobuf:"varint,13,opt,name=nerflevel,proto3" json:"nerflevel,omitempty"`
	BalanceAnchor      bool       `protobuf:"varint,15,opt,name=balanceAnchor,proto3" json:"balanceAnchor,omitempty"`
	StarterCard        bool       `protobuf:"varint,16,opt,name=starterCard,proto3" json:"starterCard,omitempty"`
	Rarity             CardRarity `protobuf:"varint,17,opt,name=rarity,proto3,enum=cardchain.cardchain.CardRarity" json:"rarity,omitempty"`
}

func (m *Card) Reset()         { *m = Card{} }
func (m *Card) String() string { return proto.CompactTextString(m) }
func (*Card) ProtoMessage()    {}
func (*Card) Descriptor() ([]byte, []int) {
	return fileDescriptor_a360ffd2377ddc30, []int{0}
}
func (m *Card) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Card) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Card.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Card) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Card.Merge(m, src)
}
func (m *Card) XXX_Size() int {
	return m.Size()
}
func (m *Card) XXX_DiscardUnknown() {
	xxx_messageInfo_Card.DiscardUnknown(m)
}

var xxx_messageInfo_Card proto.InternalMessageInfo

func (m *Card) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Card) GetArtist() string {
	if m != nil {
		return m.Artist
	}
	return ""
}

func (m *Card) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *Card) GetImageId() uint64 {
	if m != nil {
		return m.ImageId
	}
	return 0
}

func (m *Card) GetFullArt() bool {
	if m != nil {
		return m.FullArt
	}
	return false
}

func (m *Card) GetNotes() string {
	if m != nil {
		return m.Notes
	}
	return ""
}

func (m *Card) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_scheme
}

func (m *Card) GetVotePool() types.Coin {
	if m != nil {
		return m.VotePool
	}
	return types.Coin{}
}

func (m *Card) GetVoters() []string {
	if m != nil {
		return m.Voters
	}
	return nil
}

func (m *Card) GetFairEnoughVotes() uint64 {
	if m != nil {
		return m.FairEnoughVotes
	}
	return 0
}

func (m *Card) GetOverpoweredVotes() uint64 {
	if m != nil {
		return m.OverpoweredVotes
	}
	return 0
}

func (m *Card) GetUnderpoweredVotes() uint64 {
	if m != nil {
		return m.UnderpoweredVotes
	}
	return 0
}

func (m *Card) GetInappropriateVotes() uint64 {
	if m != nil {
		return m.InappropriateVotes
	}
	return 0
}

func (m *Card) GetNerflevel() int64 {
	if m != nil {
		return m.Nerflevel
	}
	return 0
}

func (m *Card) GetBalanceAnchor() bool {
	if m != nil {
		return m.BalanceAnchor
	}
	return false
}

func (m *Card) GetStarterCard() bool {
	if m != nil {
		return m.StarterCard
	}
	return false
}

func (m *Card) GetRarity() CardRarity {
	if m != nil {
		return m.Rarity
	}
	return CardRarity_common
}

type TimeStamp struct {
	TimeStamp uint64 `protobuf:"varint,1,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
}

func (m *TimeStamp) Reset()         { *m = TimeStamp{} }
func (m *TimeStamp) String() string { return proto.CompactTextString(m) }
func (*TimeStamp) ProtoMessage()    {}
func (*TimeStamp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a360ffd2377ddc30, []int{1}
}
func (m *TimeStamp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TimeStamp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TimeStamp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TimeStamp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeStamp.Merge(m, src)
}
func (m *TimeStamp) XXX_Size() int {
	return m.Size()
}
func (m *TimeStamp) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeStamp.DiscardUnknown(m)
}

var xxx_messageInfo_TimeStamp proto.InternalMessageInfo

func (m *TimeStamp) GetTimeStamp() uint64 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

func init() {
	proto.RegisterEnum("cardchain.cardchain.Status", Status_name, Status_value)
	proto.RegisterEnum("cardchain.cardchain.CardRarity", CardRarity_name, CardRarity_value)
	proto.RegisterEnum("cardchain.cardchain.CardClass", CardClass_name, CardClass_value)
	proto.RegisterEnum("cardchain.cardchain.CardType", CardType_name, CardType_value)
	proto.RegisterType((*Card)(nil), "cardchain.cardchain.Card")
	proto.RegisterType((*TimeStamp)(nil), "cardchain.cardchain.TimeStamp")
}

func init() { proto.RegisterFile("cardchain/cardchain/card.proto", fileDescriptor_a360ffd2377ddc30) }

var fileDescriptor_a360ffd2377ddc30 = []byte{
	// 741 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0x4d, 0x8f, 0xdb, 0x36,
	0x10, 0x35, 0xd7, 0x5a, 0xd9, 0x1a, 0xef, 0x07, 0x97, 0x0d, 0x0a, 0x25, 0x2d, 0x14, 0x21, 0xe8,
	0x41, 0x35, 0x0a, 0x19, 0x49, 0x0e, 0x2d, 0x50, 0xa0, 0x40, 0xb2, 0x2d, 0x8a, 0x1c, 0x0a, 0x14,
	0xda, 0x20, 0x87, 0x5e, 0x0a, 0x9a, 0x9a, 0xb5, 0x09, 0x48, 0xa4, 0x42, 0x52, 0x4e, 0xfc, 0x2f,
	0x7a, 0xed, 0x3f, 0xca, 0x31, 0xc7, 0x9e, 0x8a, 0x62, 0xf7, 0x3f, 0xf4, 0x1c, 0x90, 0xf2, 0x7e,
	0x24, 0xd9, 0xdb, 0xbc, 0x37, 0x8f, 0x33, 0x7c, 0x33, 0x04, 0x21, 0x13, 0xdc, 0xd4, 0x62, 0xcd,
	0xa5, 0x5a, 0x7c, 0x1c, 0x95, 0x9d, 0xd1, 0x4e, 0xb3, 0x2f, 0xae, 0xd9, 0xf2, 0x3a, 0x7a, 0x70,
	0x6f, 0xa5, 0x57, 0x3a, 0xe4, 0x17, 0x3e, 0x1a, 0xa4, 0x0f, 0x32, 0xa1, 0x6d, 0xab, 0xed, 0x62,
	0xc9, 0x2d, 0x2e, 0x36, 0x8f, 0x97, 0xe8, 0xf8, 0xe3, 0x85, 0xd0, 0x52, 0x0d, 0xf9, 0x47, 0xff,
	0x47, 0x10, 0x9d, 0x72, 0x53, 0xb3, 0x7b, 0xb0, 0xaf, 0xdf, 0x28, 0x34, 0x29, 0xc9, 0x49, 0x91,
	0x54, 0x03, 0x60, 0x5f, 0x42, 0xcc, 0x8d, 0x93, 0xd6, 0xa5, 0x7b, 0x81, 0xde, 0x21, 0x96, 0xc2,
	0x44, 0x68, 0xe5, 0x50, 0xb9, 0x74, 0x9c, 0x93, 0xe2, 0xa0, 0xba, 0x82, 0xec, 0x3e, 0x4c, 0x65,
	0xcb, 0x57, 0xf8, 0xa7, 0xac, 0xd3, 0x28, 0x27, 0x45, 0x54, 0x4d, 0x02, 0x7e, 0x51, 0xfb, 0x43,
	0xe7, 0x7d, 0xd3, 0x3c, 0x33, 0x2e, 0xdd, 0xcf, 0x49, 0x31, 0xad, 0xae, 0xa0, 0x6f, 0xae, 0xb4,
	0x43, 0x9b, 0xc6, 0x43, 0xf3, 0x00, 0xd8, 0x53, 0x88, 0xad, 0xe3, 0xae, 0xb7, 0xe9, 0x24, 0x27,
	0xc5, 0xd1, 0x93, 0xaf, 0xca, 0x3b, 0x7c, 0x97, 0x67, 0x41, 0x52, 0xed, 0xa4, 0xec, 0x47, 0x98,
	0x6e, 0xb4, 0xc3, 0xdf, 0xb5, 0x6e, 0xd2, 0x69, 0x4e, 0x8a, 0xd9, 0x93, 0xfb, 0xe5, 0x30, 0x83,
	0xd2, 0xcf, 0xa0, 0xdc, 0xcd, 0xa0, 0x3c, 0xd5, 0x52, 0x3d, 0x8f, 0xde, 0xfd, 0xfb, 0x70, 0x54,
	0x5d, 0x1f, 0xf0, 0x76, 0x7d, 0x6c, 0x6c, 0x7a, 0x94, 0x8f, 0xbd, 0xdd, 0x01, 0xb1, 0x02, 0x8e,
	0xcf, 0xb9, 0x34, 0xbf, 0x28, 0xdd, 0xaf, 0xd6, 0xaf, 0xc2, 0x4d, 0x93, 0xe0, 0xed, 0x53, 0x9a,
	0xcd, 0x81, 0xea, 0x0d, 0x9a, 0x4e, 0xbf, 0x41, 0x83, 0xf5, 0x20, 0x85, 0x20, 0xfd, 0x8c, 0x67,
	0xdf, 0xc1, 0x49, 0xaf, 0xea, 0x4f, 0xc4, 0xb3, 0x20, 0xfe, 0x3c, 0xc1, 0x4a, 0x60, 0x52, 0xf1,
	0xae, 0x33, 0xba, 0x33, 0x92, 0x3b, 0x1c, 0xe4, 0x07, 0x41, 0x7e, 0x47, 0x86, 0x7d, 0x0d, 0x89,
	0x42, 0x73, 0xde, 0xe0, 0x06, 0x9b, 0xf4, 0x30, 0x27, 0xc5, 0xb8, 0xba, 0x21, 0xd8, 0x37, 0x70,
	0xb8, 0xe4, 0x0d, 0x57, 0x02, 0x9f, 0x29, 0xb1, 0xd6, 0x26, 0x3d, 0x0e, 0x1b, 0xf9, 0x98, 0x64,
	0x39, 0xcc, 0xac, 0xe3, 0xc6, 0xa1, 0xf1, 0x6f, 0x24, 0xa5, 0x41, 0x73, 0x9b, 0x62, 0xdf, 0x43,
	0x6c, 0xb8, 0x91, 0x6e, 0x9b, 0x9e, 0x84, 0x1d, 0x3d, 0xbc, 0x73, 0x47, 0x5e, 0x5a, 0x05, 0x59,
	0xb5, 0x93, 0x3f, 0xfa, 0x16, 0x92, 0x97, 0xb2, 0xc5, 0x33, 0xc7, 0xdb, 0xce, 0xdf, 0xd5, 0x5d,
	0x81, 0xf0, 0x00, 0xa3, 0xea, 0x86, 0x98, 0xff, 0x4d, 0x20, 0x1e, 0xb6, 0xcc, 0x00, 0x62, 0x2b,
	0xd6, 0xd8, 0x22, 0x1d, 0xb1, 0x43, 0x48, 0xc2, 0x1b, 0x76, 0xdb, 0x0e, 0x29, 0x61, 0x09, 0xec,
	0x3b, 0x23, 0x79, 0x43, 0xf7, 0x42, 0x06, 0x4d, 0xcb, 0x15, 0x2a, 0x47, 0xc7, 0x1e, 0xda, 0xde,
	0x76, 0xa8, 0x6a, 0xac, 0x69, 0xe4, 0x6b, 0x2c, 0xb9, 0x52, 0x58, 0xd3, 0x7d, 0x76, 0x04, 0x30,
	0xc4, 0x67, 0x5a, 0x2b, 0x1a, 0x33, 0x06, 0x47, 0x03, 0x7e, 0x85, 0x66, 0x1b, 0xb8, 0x09, 0x9b,
	0x42, 0xa4, 0xb4, 0x42, 0x3a, 0x65, 0x27, 0x70, 0xc8, 0xeb, 0x0d, 0x2a, 0xd7, 0x1b, 0x7c, 0xe1,
	0xb0, 0xa5, 0xc9, 0xfc, 0x37, 0x80, 0x1b, 0x73, 0xbe, 0xb4, 0xd0, 0x6d, 0xab, 0x15, 0x1d, 0xb1,
	0x03, 0x98, 0xf6, 0x6a, 0x87, 0x88, 0x2f, 0x62, 0xb8, 0x41, 0xba, 0xc7, 0x8e, 0x61, 0x86, 0x6f,
	0x05, 0x76, 0x4e, 0x6a, 0xc5, 0x1b, 0x3a, 0xf6, 0x87, 0x7a, 0x25, 0x5f, 0xf7, 0x48, 0xa3, 0xf9,
	0x29, 0x24, 0xbe, 0xdc, 0x69, 0xc3, 0x6d, 0x30, 0xab, 0xb8, 0xef, 0x45, 0x47, 0x6c, 0x06, 0x13,
	0xd1, 0x37, 0x01, 0x10, 0x6f, 0xa8, 0xdd, 0x5a, 0x27, 0x85, 0xb4, 0x2d, 0xdd, 0xf3, 0x26, 0x1c,
	0x8a, 0xb5, 0xd2, 0x8d, 0x5e, 0x6d, 0xe9, 0x78, 0xfe, 0x13, 0x4c, 0x7d, 0x91, 0x97, 0xdb, 0x0e,
	0xfd, 0x54, 0xba, 0x86, 0x0b, 0x5f, 0x02, 0x20, 0xe6, 0xc2, 0x77, 0xa5, 0xc4, 0xc7, 0xa8, 0x9c,
	0x74, 0xdb, 0xe1, 0x42, 0x6b, 0xe4, 0xf5, 0xeb, 0x3e, 0x6c, 0x95, 0x8e, 0x9f, 0x57, 0xef, 0x2e,
	0x32, 0xf2, 0xfe, 0x22, 0x23, 0xff, 0x5d, 0x64, 0xe4, 0xaf, 0xcb, 0x6c, 0xf4, 0xfe, 0x32, 0x1b,
	0xfd, 0x73, 0x99, 0x8d, 0xfe, 0xf8, 0x61, 0x25, 0xdd, 0xba, 0x5f, 0x96, 0x42, 0xb7, 0x8b, 0x9f,
	0x51, 0xa0, 0x72, 0x86, 0x37, 0xbe, 0xd7, 0xaf, 0xbc, 0xc5, 0x5b, 0x7f, 0xd5, 0xdb, 0x5b, 0xb1,
	0x5f, 0x8f, 0x5d, 0xc6, 0x61, 0x55, 0x4f, 0x3f, 0x04, 0x00, 0x00, 0xff, 0xff, 0x59, 0x6e, 0x4f,
	0xc4, 0xdb, 0x04, 0x00, 0x00,
}

func (m *Card) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Card) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Card) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Rarity != 0 {
		i = encodeVarintCard(dAtA, i, uint64(m.Rarity))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x88
	}
	if m.StarterCard {
		i--
		if m.StarterCard {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if m.BalanceAnchor {
		i--
		if m.BalanceAnchor {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x78
	}
	if len(m.Voters) > 0 {
		for iNdEx := len(m.Voters) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Voters[iNdEx])
			copy(dAtA[i:], m.Voters[iNdEx])
			i = encodeVarintCard(dAtA, i, uint64(len(m.Voters[iNdEx])))
			i--
			dAtA[i] = 0x72
		}
	}
	if m.Nerflevel != 0 {
		i = encodeVarintCard(dAtA, i, uint64(m.Nerflevel))
		i--
		dAtA[i] = 0x68
	}
	if m.InappropriateVotes != 0 {
		i = encodeVarintCard(dAtA, i, uint64(m.InappropriateVotes))
		i--
		dAtA[i] = 0x60
	}
	if m.UnderpoweredVotes != 0 {
		i = encodeVarintCard(dAtA, i, uint64(m.UnderpoweredVotes))
		i--
		dAtA[i] = 0x58
	}
	if m.OverpoweredVotes != 0 {
		i = encodeVarintCard(dAtA, i, uint64(m.OverpoweredVotes))
		i--
		dAtA[i] = 0x50
	}
	if m.FairEnoughVotes != 0 {
		i = encodeVarintCard(dAtA, i, uint64(m.FairEnoughVotes))
		i--
		dAtA[i] = 0x48
	}
	{
		size, err := m.VotePool.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCard(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	if m.Status != 0 {
		i = encodeVarintCard(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x38
	}
	if len(m.Notes) > 0 {
		i -= len(m.Notes)
		copy(dAtA[i:], m.Notes)
		i = encodeVarintCard(dAtA, i, uint64(len(m.Notes)))
		i--
		dAtA[i] = 0x32
	}
	if m.FullArt {
		i--
		if m.FullArt {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.ImageId != 0 {
		i = encodeVarintCard(dAtA, i, uint64(m.ImageId))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintCard(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Artist) > 0 {
		i -= len(m.Artist)
		copy(dAtA[i:], m.Artist)
		i = encodeVarintCard(dAtA, i, uint64(len(m.Artist)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintCard(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TimeStamp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TimeStamp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TimeStamp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TimeStamp != 0 {
		i = encodeVarintCard(dAtA, i, uint64(m.TimeStamp))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCard(dAtA []byte, offset int, v uint64) int {
	offset -= sovCard(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Card) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovCard(uint64(l))
	}
	l = len(m.Artist)
	if l > 0 {
		n += 1 + l + sovCard(uint64(l))
	}
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovCard(uint64(l))
	}
	if m.ImageId != 0 {
		n += 1 + sovCard(uint64(m.ImageId))
	}
	if m.FullArt {
		n += 2
	}
	l = len(m.Notes)
	if l > 0 {
		n += 1 + l + sovCard(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovCard(uint64(m.Status))
	}
	l = m.VotePool.Size()
	n += 1 + l + sovCard(uint64(l))
	if m.FairEnoughVotes != 0 {
		n += 1 + sovCard(uint64(m.FairEnoughVotes))
	}
	if m.OverpoweredVotes != 0 {
		n += 1 + sovCard(uint64(m.OverpoweredVotes))
	}
	if m.UnderpoweredVotes != 0 {
		n += 1 + sovCard(uint64(m.UnderpoweredVotes))
	}
	if m.InappropriateVotes != 0 {
		n += 1 + sovCard(uint64(m.InappropriateVotes))
	}
	if m.Nerflevel != 0 {
		n += 1 + sovCard(uint64(m.Nerflevel))
	}
	if len(m.Voters) > 0 {
		for _, s := range m.Voters {
			l = len(s)
			n += 1 + l + sovCard(uint64(l))
		}
	}
	if m.BalanceAnchor {
		n += 2
	}
	if m.StarterCard {
		n += 3
	}
	if m.Rarity != 0 {
		n += 2 + sovCard(uint64(m.Rarity))
	}
	return n
}

func (m *TimeStamp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TimeStamp != 0 {
		n += 1 + sovCard(uint64(m.TimeStamp))
	}
	return n
}

func sovCard(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCard(x uint64) (n int) {
	return sovCard(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Card) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCard
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
			return fmt.Errorf("proto: Card: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Card: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCard
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
				return ErrInvalidLengthCard
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCard
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Artist", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCard
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
				return ErrInvalidLengthCard
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCard
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Artist = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCard
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCard
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCard
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = append(m.Content[:0], dAtA[iNdEx:postIndex]...)
			if m.Content == nil {
				m.Content = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImageId", wireType)
			}
			m.ImageId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCard
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ImageId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FullArt", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCard
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
			m.FullArt = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Notes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCard
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
				return ErrInvalidLengthCard
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCard
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Notes = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCard
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotePool", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCard
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
				return ErrInvalidLengthCard
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCard
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.VotePool.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FairEnoughVotes", wireType)
			}
			m.FairEnoughVotes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCard
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
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OverpoweredVotes", wireType)
			}
			m.OverpoweredVotes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCard
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
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnderpoweredVotes", wireType)
			}
			m.UnderpoweredVotes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCard
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
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InappropriateVotes", wireType)
			}
			m.InappropriateVotes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCard
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
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nerflevel", wireType)
			}
			m.Nerflevel = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCard
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nerflevel |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Voters", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCard
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
				return ErrInvalidLengthCard
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCard
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Voters = append(m.Voters, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BalanceAnchor", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCard
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
			m.BalanceAnchor = bool(v != 0)
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StarterCard", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCard
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
			m.StarterCard = bool(v != 0)
		case 17:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rarity", wireType)
			}
			m.Rarity = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCard
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Rarity |= CardRarity(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCard(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCard
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
func (m *TimeStamp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCard
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
			return fmt.Errorf("proto: TimeStamp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TimeStamp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeStamp", wireType)
			}
			m.TimeStamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCard
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeStamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCard(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCard
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
func skipCard(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCard
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
					return 0, ErrIntOverflowCard
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
					return 0, ErrIntOverflowCard
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
				return 0, ErrInvalidLengthCard
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCard
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCard
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCard        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCard          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCard = fmt.Errorf("proto: unexpected end of group")
)
