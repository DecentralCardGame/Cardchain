// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cardchain/cardchain/sell_offer.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type SellOfferStatus int32

const (
	SellOfferStatus_open    SellOfferStatus = 0
	SellOfferStatus_sold    SellOfferStatus = 1
	SellOfferStatus_removed SellOfferStatus = 2
)

var SellOfferStatus_name = map[int32]string{
	0: "open",
	1: "sold",
	2: "removed",
}

var SellOfferStatus_value = map[string]int32{
	"open":    0,
	"sold":    1,
	"removed": 2,
}

func (x SellOfferStatus) String() string {
	return proto.EnumName(SellOfferStatus_name, int32(x))
}

func (SellOfferStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_18dbd9a8240e28bf, []int{0}
}

type SellOffer struct {
	Seller string                                  `protobuf:"bytes,1,opt,name=seller,proto3" json:"seller,omitempty"`
	Buyer  string                                  `protobuf:"bytes,2,opt,name=buyer,proto3" json:"buyer,omitempty"`
	Card   uint64                                  `protobuf:"varint,3,opt,name=card,proto3" json:"card,omitempty"`
	Price  github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,4,opt,name=price,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"price"`
	Status SellOfferStatus                         `protobuf:"varint,5,opt,name=status,proto3,enum=DecentralCardGame.cardchain.cardchain.SellOfferStatus" json:"status,omitempty"`
}

func (m *SellOffer) Reset()         { *m = SellOffer{} }
func (m *SellOffer) String() string { return proto.CompactTextString(m) }
func (*SellOffer) ProtoMessage()    {}
func (*SellOffer) Descriptor() ([]byte, []int) {
	return fileDescriptor_18dbd9a8240e28bf, []int{0}
}
func (m *SellOffer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SellOffer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SellOffer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SellOffer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SellOffer.Merge(m, src)
}
func (m *SellOffer) XXX_Size() int {
	return m.Size()
}
func (m *SellOffer) XXX_DiscardUnknown() {
	xxx_messageInfo_SellOffer.DiscardUnknown(m)
}

var xxx_messageInfo_SellOffer proto.InternalMessageInfo

func (m *SellOffer) GetSeller() string {
	if m != nil {
		return m.Seller
	}
	return ""
}

func (m *SellOffer) GetBuyer() string {
	if m != nil {
		return m.Buyer
	}
	return ""
}

func (m *SellOffer) GetCard() uint64 {
	if m != nil {
		return m.Card
	}
	return 0
}

func (m *SellOffer) GetStatus() SellOfferStatus {
	if m != nil {
		return m.Status
	}
	return SellOfferStatus_open
}

func init() {
	proto.RegisterEnum("DecentralCardGame.cardchain.cardchain.SellOfferStatus", SellOfferStatus_name, SellOfferStatus_value)
	proto.RegisterType((*SellOffer)(nil), "DecentralCardGame.cardchain.cardchain.SellOffer")
}

func init() {
	proto.RegisterFile("cardchain/cardchain/sell_offer.proto", fileDescriptor_18dbd9a8240e28bf)
}

var fileDescriptor_18dbd9a8240e28bf = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x4b, 0x02, 0x41,
	0x14, 0xc7, 0x77, 0x6c, 0xb5, 0x9c, 0xa0, 0x64, 0x90, 0x58, 0x3a, 0xac, 0x12, 0x45, 0x12, 0xb4,
	0x0b, 0x06, 0xd1, 0x59, 0x8b, 0x6e, 0x05, 0xeb, 0xad, 0x4b, 0xac, 0xbb, 0x4f, 0x5d, 0x9a, 0xdd,
	0xb7, 0xcc, 0x8c, 0x91, 0xdf, 0xa2, 0x8f, 0xe5, 0xd1, 0x63, 0x74, 0x90, 0xd0, 0x3e, 0x48, 0xcc,
	0x28, 0xab, 0xd4, 0xa5, 0xd3, 0xfc, 0xde, 0xcc, 0xfb, 0xff, 0x99, 0xff, 0x7b, 0xf4, 0x34, 0x0a,
	0x45, 0x1c, 0x8d, 0xc2, 0x24, 0xf3, 0x37, 0x24, 0x81, 0xf3, 0x67, 0x1c, 0x0c, 0x40, 0x78, 0xb9,
	0x40, 0x85, 0xec, 0xec, 0x16, 0x22, 0xc8, 0x94, 0x08, 0x79, 0x37, 0x14, 0xf1, 0x7d, 0x98, 0x82,
	0x57, 0x74, 0x6f, 0xe8, 0xb8, 0x3e, 0xc4, 0x21, 0x1a, 0x85, 0xaf, 0x69, 0x25, 0x3e, 0xf9, 0x26,
	0xb4, 0xda, 0x03, 0xce, 0x1f, 0xb5, 0x21, 0x3b, 0xa2, 0x15, 0x6d, 0x0f, 0xc2, 0x21, 0x4d, 0xd2,
	0xaa, 0x06, 0xeb, 0x8a, 0xd5, 0x69, 0xb9, 0x3f, 0x9e, 0x80, 0x70, 0x4a, 0xe6, 0x7a, 0x55, 0x30,
	0x46, 0x6d, 0x6d, 0xef, 0xec, 0x34, 0x49, 0xcb, 0x0e, 0x0c, 0xb3, 0x3b, 0x5a, 0xce, 0x45, 0x12,
	0x81, 0x63, 0xeb, 0xce, 0x8e, 0x3f, 0x9d, 0x37, 0xac, 0xcf, 0x79, 0xe3, 0x7c, 0x98, 0xa8, 0xd1,
	0xb8, 0xef, 0x45, 0x98, 0xfa, 0x11, 0xca, 0x14, 0xe5, 0xfa, 0xb8, 0x94, 0xf1, 0x8b, 0xaf, 0x26,
	0x39, 0x48, 0xaf, 0x8b, 0x49, 0x16, 0xac, 0xd4, 0xec, 0x81, 0x56, 0xa4, 0x0a, 0xd5, 0x58, 0x3a,
	0xe5, 0x26, 0x69, 0x1d, 0xb4, 0xaf, 0xbd, 0x7f, 0x85, 0xf4, 0x8a, 0x28, 0x3d, 0xa3, 0x0e, 0xd6,
	0x2e, 0x17, 0x6d, 0x7a, 0xf8, 0xeb, 0x89, 0xed, 0x51, 0x1b, 0x73, 0xc8, 0x6a, 0x96, 0x26, 0x89,
	0x3c, 0xae, 0x11, 0xb6, 0x4f, 0x77, 0x05, 0xa4, 0xf8, 0x0a, 0x71, 0xad, 0xd4, 0x09, 0xa6, 0x0b,
	0x97, 0xcc, 0x16, 0x2e, 0xf9, 0x5a, 0xb8, 0xe4, 0x7d, 0xe9, 0x5a, 0xb3, 0xa5, 0x6b, 0x7d, 0x2c,
	0x5d, 0xeb, 0xe9, 0x66, 0x2b, 0xcd, 0x9f, 0x7f, 0xf9, 0xdd, 0x62, 0x55, 0x6f, 0x5b, 0x6b, 0x33,
	0x19, 0xfb, 0x15, 0x33, 0xf5, 0xab, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x01, 0x47, 0xf9, 0x67,
	0xda, 0x01, 0x00, 0x00,
}

func (m *SellOffer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SellOffer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SellOffer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintSellOffer(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x28
	}
	{
		size := m.Price.Size()
		i -= size
		if _, err := m.Price.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSellOffer(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.Card != 0 {
		i = encodeVarintSellOffer(dAtA, i, uint64(m.Card))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Buyer) > 0 {
		i -= len(m.Buyer)
		copy(dAtA[i:], m.Buyer)
		i = encodeVarintSellOffer(dAtA, i, uint64(len(m.Buyer)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Seller) > 0 {
		i -= len(m.Seller)
		copy(dAtA[i:], m.Seller)
		i = encodeVarintSellOffer(dAtA, i, uint64(len(m.Seller)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSellOffer(dAtA []byte, offset int, v uint64) int {
	offset -= sovSellOffer(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SellOffer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Seller)
	if l > 0 {
		n += 1 + l + sovSellOffer(uint64(l))
	}
	l = len(m.Buyer)
	if l > 0 {
		n += 1 + l + sovSellOffer(uint64(l))
	}
	if m.Card != 0 {
		n += 1 + sovSellOffer(uint64(m.Card))
	}
	l = m.Price.Size()
	n += 1 + l + sovSellOffer(uint64(l))
	if m.Status != 0 {
		n += 1 + sovSellOffer(uint64(m.Status))
	}
	return n
}

func sovSellOffer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSellOffer(x uint64) (n int) {
	return sovSellOffer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SellOffer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSellOffer
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
			return fmt.Errorf("proto: SellOffer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SellOffer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seller", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSellOffer
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
				return ErrInvalidLengthSellOffer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSellOffer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Seller = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Buyer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSellOffer
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
				return ErrInvalidLengthSellOffer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSellOffer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Buyer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Card", wireType)
			}
			m.Card = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSellOffer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Card |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSellOffer
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
				return ErrInvalidLengthSellOffer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSellOffer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSellOffer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= SellOfferStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSellOffer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSellOffer
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
func skipSellOffer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSellOffer
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
					return 0, ErrIntOverflowSellOffer
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
					return 0, ErrIntOverflowSellOffer
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
				return 0, ErrInvalidLengthSellOffer
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSellOffer
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSellOffer
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSellOffer        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSellOffer          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSellOffer = fmt.Errorf("proto: unexpected end of group")
)
