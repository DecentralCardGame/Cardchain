// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cardchain/cardchain/council.proto

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

type Response int32

const (
	Response_Yes        Response = 0
	Response_No         Response = 1
	Response_Suggestion Response = 2
)

var Response_name = map[int32]string{
	0: "Yes",
	1: "No",
	2: "Suggestion",
}

var Response_value = map[string]int32{
	"Yes":        0,
	"No":         1,
	"Suggestion": 2,
}

func (x Response) String() string {
	return proto.EnumName(Response_name, int32(x))
}

func (Response) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fcd39b22a840d084, []int{0}
}

type CouncelingStatus int32

const (
	CouncelingStatus_councilOpen     CouncelingStatus = 0
	CouncelingStatus_councilCreated  CouncelingStatus = 1
	CouncelingStatus_councilClosed   CouncelingStatus = 2
	CouncelingStatus_commited        CouncelingStatus = 3
	CouncelingStatus_revealed        CouncelingStatus = 4
	CouncelingStatus_suggestionsMade CouncelingStatus = 5
)

var CouncelingStatus_name = map[int32]string{
	0: "councilOpen",
	1: "councilCreated",
	2: "councilClosed",
	3: "commited",
	4: "revealed",
	5: "suggestionsMade",
}

var CouncelingStatus_value = map[string]int32{
	"councilOpen":     0,
	"councilCreated":  1,
	"councilClosed":   2,
	"commited":        3,
	"revealed":        4,
	"suggestionsMade": 5,
}

func (x CouncelingStatus) String() string {
	return proto.EnumName(CouncelingStatus_name, int32(x))
}

func (CouncelingStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fcd39b22a840d084, []int{1}
}

type Council struct {
	CardId         uint64               `protobuf:"varint,1,opt,name=cardId,proto3" json:"cardId,omitempty"`
	Voters         []string             `protobuf:"bytes,2,rep,name=voters,proto3" json:"voters,omitempty"`
	HashResponses  []*WrapHashResponse  `protobuf:"bytes,3,rep,name=hashResponses,proto3" json:"hashResponses,omitempty"`
	ClearResponses []*WrapClearResponse `protobuf:"bytes,4,rep,name=clearResponses,proto3" json:"clearResponses,omitempty"`
	Treasury       types.Coin           `protobuf:"bytes,8,opt,name=treasury,proto3" json:"treasury"`
	Status         CouncelingStatus     `protobuf:"varint,6,opt,name=status,proto3,enum=cardchain.cardchain.CouncelingStatus" json:"status,omitempty"`
	TrialStart     uint64               `protobuf:"varint,7,opt,name=trialStart,proto3" json:"trialStart,omitempty"`
}

func (m *Council) Reset()         { *m = Council{} }
func (m *Council) String() string { return proto.CompactTextString(m) }
func (*Council) ProtoMessage()    {}
func (*Council) Descriptor() ([]byte, []int) {
	return fileDescriptor_fcd39b22a840d084, []int{0}
}
func (m *Council) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Council) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Council.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Council) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Council.Merge(m, src)
}
func (m *Council) XXX_Size() int {
	return m.Size()
}
func (m *Council) XXX_DiscardUnknown() {
	xxx_messageInfo_Council.DiscardUnknown(m)
}

var xxx_messageInfo_Council proto.InternalMessageInfo

func (m *Council) GetCardId() uint64 {
	if m != nil {
		return m.CardId
	}
	return 0
}

func (m *Council) GetVoters() []string {
	if m != nil {
		return m.Voters
	}
	return nil
}

func (m *Council) GetHashResponses() []*WrapHashResponse {
	if m != nil {
		return m.HashResponses
	}
	return nil
}

func (m *Council) GetClearResponses() []*WrapClearResponse {
	if m != nil {
		return m.ClearResponses
	}
	return nil
}

func (m *Council) GetTreasury() types.Coin {
	if m != nil {
		return m.Treasury
	}
	return types.Coin{}
}

func (m *Council) GetStatus() CouncelingStatus {
	if m != nil {
		return m.Status
	}
	return CouncelingStatus_councilOpen
}

func (m *Council) GetTrialStart() uint64 {
	if m != nil {
		return m.TrialStart
	}
	return 0
}

type WrapClearResponse struct {
	User       string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Response   Response `protobuf:"varint,2,opt,name=response,proto3,enum=cardchain.cardchain.Response" json:"response,omitempty"`
	Suggestion string   `protobuf:"bytes,3,opt,name=suggestion,proto3" json:"suggestion,omitempty"`
}

func (m *WrapClearResponse) Reset()         { *m = WrapClearResponse{} }
func (m *WrapClearResponse) String() string { return proto.CompactTextString(m) }
func (*WrapClearResponse) ProtoMessage()    {}
func (*WrapClearResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fcd39b22a840d084, []int{1}
}
func (m *WrapClearResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WrapClearResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WrapClearResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WrapClearResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WrapClearResponse.Merge(m, src)
}
func (m *WrapClearResponse) XXX_Size() int {
	return m.Size()
}
func (m *WrapClearResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WrapClearResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WrapClearResponse proto.InternalMessageInfo

func (m *WrapClearResponse) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *WrapClearResponse) GetResponse() Response {
	if m != nil {
		return m.Response
	}
	return Response_Yes
}

func (m *WrapClearResponse) GetSuggestion() string {
	if m != nil {
		return m.Suggestion
	}
	return ""
}

type WrapHashResponse struct {
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Hash string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *WrapHashResponse) Reset()         { *m = WrapHashResponse{} }
func (m *WrapHashResponse) String() string { return proto.CompactTextString(m) }
func (*WrapHashResponse) ProtoMessage()    {}
func (*WrapHashResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fcd39b22a840d084, []int{2}
}
func (m *WrapHashResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WrapHashResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WrapHashResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WrapHashResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WrapHashResponse.Merge(m, src)
}
func (m *WrapHashResponse) XXX_Size() int {
	return m.Size()
}
func (m *WrapHashResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WrapHashResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WrapHashResponse proto.InternalMessageInfo

func (m *WrapHashResponse) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *WrapHashResponse) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func init() {
	proto.RegisterEnum("cardchain.cardchain.Response", Response_name, Response_value)
	proto.RegisterEnum("cardchain.cardchain.CouncelingStatus", CouncelingStatus_name, CouncelingStatus_value)
	proto.RegisterType((*Council)(nil), "cardchain.cardchain.Council")
	proto.RegisterType((*WrapClearResponse)(nil), "cardchain.cardchain.WrapClearResponse")
	proto.RegisterType((*WrapHashResponse)(nil), "cardchain.cardchain.WrapHashResponse")
}

func init() { proto.RegisterFile("cardchain/cardchain/council.proto", fileDescriptor_fcd39b22a840d084) }

var fileDescriptor_fcd39b22a840d084 = []byte{
	// 525 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0xf6, 0xd7, 0x9b, 0x8f, 0xc9, 0xdb, 0xd4, 0xdd, 0x22, 0x64, 0x2a, 0x61, 0x42, 0x24, 0x50,
	0x54, 0x24, 0x5b, 0x2d, 0x17, 0x3e, 0xc4, 0xa5, 0x41, 0x02, 0x84, 0x28, 0xd2, 0xe6, 0x80, 0xe0,
	0xb6, 0xb1, 0x47, 0x8e, 0x25, 0xc7, 0x1b, 0xed, 0xae, 0x23, 0x2a, 0x6e, 0xfc, 0x02, 0x7e, 0x56,
	0x8f, 0x3d, 0x72, 0x42, 0x28, 0xf9, 0x07, 0xfc, 0x02, 0xe4, 0x8d, 0x71, 0x4d, 0x48, 0x6f, 0xcf,
	0xcc, 0x3c, 0xf3, 0xec, 0xcc, 0xb3, 0xbb, 0x70, 0x3f, 0x62, 0x22, 0x8e, 0x66, 0x2c, 0xcd, 0xc3,
	0x06, 0xe2, 0x45, 0x1e, 0xa5, 0x59, 0xb0, 0x10, 0x5c, 0x71, 0x72, 0x58, 0x17, 0x82, 0x1a, 0x1d,
	0xdd, 0x4a, 0x78, 0xc2, 0x75, 0x3d, 0x2c, 0xd1, 0x86, 0x7a, 0xe4, 0x47, 0x5c, 0xce, 0xb9, 0x0c,
	0xa7, 0x4c, 0x62, 0xb8, 0x3c, 0x99, 0xa2, 0x62, 0x27, 0x61, 0xc4, 0xd3, 0x7c, 0x53, 0x1f, 0xfe,
	0xb2, 0xa0, 0x3d, 0xde, 0x88, 0x93, 0xdb, 0xd0, 0x2a, 0xe5, 0xde, 0xc4, 0x9e, 0x39, 0x30, 0x47,
	0x0e, 0xad, 0xa2, 0x32, 0xbf, 0xe4, 0x0a, 0x85, 0xf4, 0xac, 0x81, 0x3d, 0xea, 0xd2, 0x2a, 0x22,
	0x6f, 0x61, 0x6f, 0xc6, 0xe4, 0x8c, 0xa2, 0x5c, 0xf0, 0x5c, 0xa2, 0xf4, 0xec, 0x81, 0x3d, 0xea,
	0x9d, 0x3e, 0x08, 0x76, 0x8c, 0x17, 0x7c, 0x10, 0x6c, 0xf1, 0xba, 0xc1, 0xa6, 0x7f, 0xf7, 0x92,
	0x73, 0xe8, 0x47, 0x19, 0x32, 0x71, 0xad, 0xe6, 0x68, 0xb5, 0x87, 0x37, 0xaa, 0x8d, 0x9b, 0x74,
	0xba, 0xd5, 0x4d, 0x9e, 0x43, 0x47, 0x09, 0x64, 0xb2, 0x10, 0x17, 0x5e, 0x67, 0x60, 0x8e, 0x7a,
	0xa7, 0x77, 0x82, 0x8d, 0x17, 0x41, 0xe9, 0x45, 0x50, 0x79, 0x11, 0x8c, 0x79, 0x9a, 0x9f, 0x39,
	0x97, 0x3f, 0xee, 0x19, 0xb4, 0x6e, 0x20, 0x2f, 0xa0, 0x25, 0x15, 0x53, 0x85, 0xf4, 0x5a, 0x03,
	0x73, 0xd4, 0xbf, 0x61, 0x25, 0xed, 0x1b, 0x66, 0x69, 0x9e, 0x4c, 0x34, 0x99, 0x56, 0x4d, 0xc4,
	0x07, 0x50, 0x22, 0x65, 0xd9, 0x44, 0x31, 0xa1, 0xbc, 0xb6, 0x36, 0xb3, 0x91, 0x19, 0x7e, 0x35,
	0xe1, 0xe0, 0x9f, 0x0d, 0x08, 0x01, 0xa7, 0x90, 0x28, 0xb4, 0xf9, 0x5d, 0xaa, 0x31, 0x79, 0x0a,
	0x1d, 0x51, 0xd5, 0x3d, 0x4b, 0x8f, 0x72, 0x77, 0xe7, 0x28, 0xb5, 0x0d, 0x35, 0xbd, 0x1c, 0x42,
	0x16, 0x49, 0x82, 0x52, 0xa5, 0x3c, 0xf7, 0x6c, 0x2d, 0xda, 0xc8, 0x0c, 0x9f, 0x81, 0xbb, 0x7d,
	0x27, 0x3b, 0x47, 0x20, 0xe0, 0x94, 0x37, 0xa5, 0x8f, 0xef, 0x52, 0x8d, 0x8f, 0x1f, 0x41, 0xa7,
	0xee, 0x69, 0x83, 0xfd, 0x11, 0xa5, 0x6b, 0x90, 0x16, 0x58, 0xe7, 0xdc, 0x35, 0x49, 0x1f, 0x60,
	0x52, 0x1f, 0xe3, 0x5a, 0xc7, 0x5f, 0xc0, 0xdd, 0x76, 0x8a, 0xec, 0x43, 0xaf, 0x7a, 0xd2, 0xef,
	0x17, 0x98, 0xbb, 0x06, 0x21, 0xd0, 0xaf, 0x12, 0x63, 0x81, 0x4c, 0x61, 0xec, 0x9a, 0xe4, 0x00,
	0xf6, 0xfe, 0xe4, 0x32, 0x2e, 0x31, 0x76, 0x2d, 0xf2, 0x3f, 0x74, 0x22, 0x3e, 0x9f, 0xa7, 0x25,
	0xc1, 0x2e, 0x23, 0x81, 0x4b, 0x64, 0x19, 0xc6, 0xae, 0x43, 0x0e, 0x61, 0xff, 0x7a, 0x3d, 0xf9,
	0x8e, 0xc5, 0xe8, 0xfe, 0x77, 0x46, 0x2f, 0x57, 0xbe, 0x79, 0xb5, 0xf2, 0xcd, 0x9f, 0x2b, 0xdf,
	0xfc, 0xb6, 0xf6, 0x8d, 0xab, 0xb5, 0x6f, 0x7c, 0x5f, 0xfb, 0xc6, 0xa7, 0x27, 0x49, 0xaa, 0x66,
	0xc5, 0x34, 0x88, 0xf8, 0x3c, 0x7c, 0x89, 0x11, 0xe6, 0x4a, 0xb0, 0x6c, 0xcc, 0x44, 0xfc, 0x8a,
	0xcd, 0xb1, 0xf1, 0xf5, 0x3e, 0x37, 0xb0, 0xba, 0x58, 0xa0, 0x9c, 0xb6, 0xf4, 0xd7, 0x79, 0xfc,
	0x3b, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x55, 0xf5, 0x1e, 0xaa, 0x03, 0x00, 0x00,
}

func (m *Council) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Council) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Council) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Treasury.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCouncil(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	if m.TrialStart != 0 {
		i = encodeVarintCouncil(dAtA, i, uint64(m.TrialStart))
		i--
		dAtA[i] = 0x38
	}
	if m.Status != 0 {
		i = encodeVarintCouncil(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x30
	}
	if len(m.ClearResponses) > 0 {
		for iNdEx := len(m.ClearResponses) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClearResponses[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCouncil(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.HashResponses) > 0 {
		for iNdEx := len(m.HashResponses) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.HashResponses[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCouncil(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Voters) > 0 {
		for iNdEx := len(m.Voters) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Voters[iNdEx])
			copy(dAtA[i:], m.Voters[iNdEx])
			i = encodeVarintCouncil(dAtA, i, uint64(len(m.Voters[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.CardId != 0 {
		i = encodeVarintCouncil(dAtA, i, uint64(m.CardId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *WrapClearResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WrapClearResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WrapClearResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Suggestion) > 0 {
		i -= len(m.Suggestion)
		copy(dAtA[i:], m.Suggestion)
		i = encodeVarintCouncil(dAtA, i, uint64(len(m.Suggestion)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Response != 0 {
		i = encodeVarintCouncil(dAtA, i, uint64(m.Response))
		i--
		dAtA[i] = 0x10
	}
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintCouncil(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *WrapHashResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WrapHashResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WrapHashResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintCouncil(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintCouncil(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCouncil(dAtA []byte, offset int, v uint64) int {
	offset -= sovCouncil(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Council) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CardId != 0 {
		n += 1 + sovCouncil(uint64(m.CardId))
	}
	if len(m.Voters) > 0 {
		for _, s := range m.Voters {
			l = len(s)
			n += 1 + l + sovCouncil(uint64(l))
		}
	}
	if len(m.HashResponses) > 0 {
		for _, e := range m.HashResponses {
			l = e.Size()
			n += 1 + l + sovCouncil(uint64(l))
		}
	}
	if len(m.ClearResponses) > 0 {
		for _, e := range m.ClearResponses {
			l = e.Size()
			n += 1 + l + sovCouncil(uint64(l))
		}
	}
	if m.Status != 0 {
		n += 1 + sovCouncil(uint64(m.Status))
	}
	if m.TrialStart != 0 {
		n += 1 + sovCouncil(uint64(m.TrialStart))
	}
	l = m.Treasury.Size()
	n += 1 + l + sovCouncil(uint64(l))
	return n
}

func (m *WrapClearResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovCouncil(uint64(l))
	}
	if m.Response != 0 {
		n += 1 + sovCouncil(uint64(m.Response))
	}
	l = len(m.Suggestion)
	if l > 0 {
		n += 1 + l + sovCouncil(uint64(l))
	}
	return n
}

func (m *WrapHashResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovCouncil(uint64(l))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovCouncil(uint64(l))
	}
	return n
}

func sovCouncil(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCouncil(x uint64) (n int) {
	return sovCouncil(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Council) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCouncil
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
			return fmt.Errorf("proto: Council: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Council: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CardId", wireType)
			}
			m.CardId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCouncil
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
				return fmt.Errorf("proto: wrong wireType = %d for field Voters", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCouncil
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
				return ErrInvalidLengthCouncil
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCouncil
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Voters = append(m.Voters, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HashResponses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCouncil
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
				return ErrInvalidLengthCouncil
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCouncil
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HashResponses = append(m.HashResponses, &WrapHashResponse{})
			if err := m.HashResponses[len(m.HashResponses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClearResponses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCouncil
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
				return ErrInvalidLengthCouncil
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCouncil
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClearResponses = append(m.ClearResponses, &WrapClearResponse{})
			if err := m.ClearResponses[len(m.ClearResponses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCouncil
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= CouncelingStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrialStart", wireType)
			}
			m.TrialStart = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCouncil
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TrialStart |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Treasury", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCouncil
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
				return ErrInvalidLengthCouncil
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCouncil
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Treasury.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCouncil(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCouncil
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
func (m *WrapClearResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCouncil
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
			return fmt.Errorf("proto: WrapClearResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WrapClearResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCouncil
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
				return ErrInvalidLengthCouncil
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCouncil
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Response", wireType)
			}
			m.Response = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCouncil
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Response |= Response(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Suggestion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCouncil
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
				return ErrInvalidLengthCouncil
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCouncil
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Suggestion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCouncil(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCouncil
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
func (m *WrapHashResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCouncil
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
			return fmt.Errorf("proto: WrapHashResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WrapHashResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCouncil
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
				return ErrInvalidLengthCouncil
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCouncil
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCouncil
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
				return ErrInvalidLengthCouncil
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCouncil
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCouncil(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCouncil
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
func skipCouncil(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCouncil
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
					return 0, ErrIntOverflowCouncil
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
					return 0, ErrIntOverflowCouncil
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
				return 0, ErrInvalidLengthCouncil
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCouncil
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCouncil
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCouncil        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCouncil          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCouncil = fmt.Errorf("proto: unexpected end of group")
)
