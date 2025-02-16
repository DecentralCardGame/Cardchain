// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cardchain/cardchain/encounter.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

type Encounter struct {
	Id         uint64            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Drawlist   []uint64          `protobuf:"varint,2,rep,packed,name=drawlist,proto3" json:"drawlist,omitempty"`
	Proven     bool              `protobuf:"varint,3,opt,name=proven,proto3" json:"proven,omitempty"`
	Owner      string            `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
	Parameters map[string]string `protobuf:"bytes,5,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ImageId    uint64            `protobuf:"varint,6,opt,name=imageId,proto3" json:"imageId,omitempty"`
	Name       string            `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *Encounter) Reset()         { *m = Encounter{} }
func (m *Encounter) String() string { return proto.CompactTextString(m) }
func (*Encounter) ProtoMessage()    {}
func (*Encounter) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1555abf6f7ee418, []int{0}
}
func (m *Encounter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Encounter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Encounter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Encounter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Encounter.Merge(m, src)
}
func (m *Encounter) XXX_Size() int {
	return m.Size()
}
func (m *Encounter) XXX_DiscardUnknown() {
	xxx_messageInfo_Encounter.DiscardUnknown(m)
}

var xxx_messageInfo_Encounter proto.InternalMessageInfo

func (m *Encounter) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Encounter) GetDrawlist() []uint64 {
	if m != nil {
		return m.Drawlist
	}
	return nil
}

func (m *Encounter) GetProven() bool {
	if m != nil {
		return m.Proven
	}
	return false
}

func (m *Encounter) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Encounter) GetParameters() map[string]string {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *Encounter) GetImageId() uint64 {
	if m != nil {
		return m.ImageId
	}
	return 0
}

func (m *Encounter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type EncounterWithImage struct {
	Encounter *Encounter `protobuf:"bytes,1,opt,name=encounter,proto3" json:"encounter,omitempty"`
	Image     []byte     `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
}

func (m *EncounterWithImage) Reset()         { *m = EncounterWithImage{} }
func (m *EncounterWithImage) String() string { return proto.CompactTextString(m) }
func (*EncounterWithImage) ProtoMessage()    {}
func (*EncounterWithImage) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1555abf6f7ee418, []int{1}
}
func (m *EncounterWithImage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EncounterWithImage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EncounterWithImage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EncounterWithImage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncounterWithImage.Merge(m, src)
}
func (m *EncounterWithImage) XXX_Size() int {
	return m.Size()
}
func (m *EncounterWithImage) XXX_DiscardUnknown() {
	xxx_messageInfo_EncounterWithImage.DiscardUnknown(m)
}

var xxx_messageInfo_EncounterWithImage proto.InternalMessageInfo

func (m *EncounterWithImage) GetEncounter() *Encounter {
	if m != nil {
		return m.Encounter
	}
	return nil
}

func (m *EncounterWithImage) GetImage() []byte {
	if m != nil {
		return m.Image
	}
	return nil
}

func init() {
	proto.RegisterType((*Encounter)(nil), "cardchain.cardchain.Encounter")
	proto.RegisterMapType((map[string]string)(nil), "cardchain.cardchain.Encounter.ParametersEntry")
	proto.RegisterType((*EncounterWithImage)(nil), "cardchain.cardchain.EncounterWithImage")
}

func init() {
	proto.RegisterFile("cardchain/cardchain/encounter.proto", fileDescriptor_d1555abf6f7ee418)
}

var fileDescriptor_d1555abf6f7ee418 = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xbf, 0x4e, 0xf3, 0x30,
	0x14, 0xc5, 0xeb, 0x24, 0xfd, 0x13, 0xf7, 0xd3, 0x07, 0x32, 0x15, 0xb2, 0x3a, 0x44, 0x51, 0x59,
	0x32, 0xa5, 0x52, 0x59, 0x2a, 0x04, 0x0b, 0x50, 0xa1, 0x2e, 0x08, 0x79, 0x41, 0x62, 0x73, 0x13,
	0x2b, 0x89, 0x68, 0xec, 0xc8, 0x75, 0x5b, 0xfa, 0x16, 0x3c, 0x03, 0x4f, 0xc3, 0xd8, 0x91, 0x11,
	0xb5, 0x2f, 0x82, 0xe2, 0x50, 0xb7, 0x42, 0x88, 0xed, 0x1c, 0x5d, 0xdf, 0x7b, 0x7e, 0xd7, 0x36,
	0x3c, 0x8b, 0xa8, 0x8c, 0xa3, 0x94, 0x66, 0xbc, 0xbf, 0x57, 0x8c, 0x47, 0x62, 0xce, 0x15, 0x93,
	0x61, 0x21, 0x85, 0x12, 0xe8, 0xc4, 0x94, 0x42, 0xa3, 0xba, 0x9d, 0x44, 0x24, 0x42, 0xd7, 0xfb,
	0xa5, 0xaa, 0x8e, 0xf6, 0xde, 0x2c, 0xe8, 0x8e, 0x76, 0xed, 0xe8, 0x3f, 0xb4, 0xb2, 0x18, 0x03,
	0x1f, 0x04, 0x0e, 0xb1, 0xb2, 0x18, 0x75, 0x61, 0x2b, 0x96, 0x74, 0x39, 0xcd, 0x66, 0x0a, 0x5b,
	0xbe, 0x1d, 0x38, 0xc4, 0x78, 0x74, 0x0a, 0x1b, 0x85, 0x14, 0x0b, 0xc6, 0xb1, 0xed, 0x83, 0xa0,
	0x45, 0xbe, 0x1d, 0xea, 0xc0, 0xba, 0x58, 0x72, 0x26, 0xb1, 0xe3, 0x83, 0xc0, 0x25, 0x95, 0x41,
	0xf7, 0x10, 0x16, 0x54, 0xd2, 0x9c, 0x29, 0x26, 0x67, 0xb8, 0xee, 0xdb, 0x41, 0x7b, 0x10, 0x86,
	0xbf, 0x70, 0x86, 0x86, 0x26, 0x7c, 0x30, 0x0d, 0x23, 0xae, 0xe4, 0x8a, 0x1c, 0x4c, 0x40, 0x18,
	0x36, 0xb3, 0x9c, 0x26, 0x6c, 0x1c, 0xe3, 0x86, 0xc6, 0xdd, 0x59, 0x84, 0xa0, 0xc3, 0x69, 0xce,
	0x70, 0x53, 0xc7, 0x6b, 0xdd, 0xbd, 0x82, 0x47, 0x3f, 0x86, 0xa1, 0x63, 0x68, 0x3f, 0xb3, 0x95,
	0xde, 0xd5, 0x25, 0xa5, 0x2c, 0xc1, 0x17, 0x74, 0x3a, 0x67, 0xd8, 0xaa, 0xc0, 0xb5, 0xb9, 0xb0,
	0x86, 0xa0, 0x97, 0x42, 0x64, 0xa8, 0x1e, 0x33, 0x95, 0x8e, 0xcb, 0x28, 0x74, 0x09, 0x5d, 0x73,
	0xf1, 0x7a, 0x4e, 0x7b, 0xe0, 0xfd, 0xbd, 0x11, 0xd9, 0x37, 0x94, 0x69, 0x9a, 0x58, 0xa7, 0xfd,
	0x23, 0x95, 0xb9, 0x26, 0xef, 0x1b, 0x0f, 0xac, 0x37, 0x1e, 0xf8, 0xdc, 0x78, 0xe0, 0x75, 0xeb,
	0xd5, 0xd6, 0x5b, 0xaf, 0xf6, 0xb1, 0xf5, 0x6a, 0x4f, 0xc3, 0x24, 0x53, 0xe9, 0x7c, 0x12, 0x46,
	0x22, 0xef, 0xdf, 0xb2, 0x88, 0x71, 0x25, 0xe9, 0xf4, 0x86, 0xca, 0xf8, 0x8e, 0xe6, 0xec, 0xe0,
	0x2f, 0xbc, 0x1c, 0x68, 0xb5, 0x2a, 0xd8, 0x6c, 0xd2, 0xd0, 0x2f, 0x7d, 0xfe, 0x15, 0x00, 0x00,
	0xff, 0xff, 0xc7, 0x31, 0xc4, 0xfa, 0x3b, 0x02, 0x00, 0x00,
}

func (m *Encounter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Encounter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Encounter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintEncounter(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x3a
	}
	if m.ImageId != 0 {
		i = encodeVarintEncounter(dAtA, i, uint64(m.ImageId))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Parameters) > 0 {
		for k := range m.Parameters {
			v := m.Parameters[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintEncounter(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintEncounter(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintEncounter(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintEncounter(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x22
	}
	if m.Proven {
		i--
		if m.Proven {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Drawlist) > 0 {
		dAtA2 := make([]byte, len(m.Drawlist)*10)
		var j1 int
		for _, num := range m.Drawlist {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintEncounter(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintEncounter(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EncounterWithImage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EncounterWithImage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EncounterWithImage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Image) > 0 {
		i -= len(m.Image)
		copy(dAtA[i:], m.Image)
		i = encodeVarintEncounter(dAtA, i, uint64(len(m.Image)))
		i--
		dAtA[i] = 0x12
	}
	if m.Encounter != nil {
		{
			size, err := m.Encounter.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEncounter(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEncounter(dAtA []byte, offset int, v uint64) int {
	offset -= sovEncounter(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Encounter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovEncounter(uint64(m.Id))
	}
	if len(m.Drawlist) > 0 {
		l = 0
		for _, e := range m.Drawlist {
			l += sovEncounter(uint64(e))
		}
		n += 1 + sovEncounter(uint64(l)) + l
	}
	if m.Proven {
		n += 2
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovEncounter(uint64(l))
	}
	if len(m.Parameters) > 0 {
		for k, v := range m.Parameters {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovEncounter(uint64(len(k))) + 1 + len(v) + sovEncounter(uint64(len(v)))
			n += mapEntrySize + 1 + sovEncounter(uint64(mapEntrySize))
		}
	}
	if m.ImageId != 0 {
		n += 1 + sovEncounter(uint64(m.ImageId))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovEncounter(uint64(l))
	}
	return n
}

func (m *EncounterWithImage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Encounter != nil {
		l = m.Encounter.Size()
		n += 1 + l + sovEncounter(uint64(l))
	}
	l = len(m.Image)
	if l > 0 {
		n += 1 + l + sovEncounter(uint64(l))
	}
	return n
}

func sovEncounter(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEncounter(x uint64) (n int) {
	return sovEncounter(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Encounter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEncounter
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
			return fmt.Errorf("proto: Encounter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Encounter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEncounter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowEncounter
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
				m.Drawlist = append(m.Drawlist, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowEncounter
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
					return ErrInvalidLengthEncounter
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthEncounter
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
				if elementCount != 0 && len(m.Drawlist) == 0 {
					m.Drawlist = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowEncounter
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
					m.Drawlist = append(m.Drawlist, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Drawlist", wireType)
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proven", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEncounter
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
			m.Proven = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEncounter
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
				return ErrInvalidLengthEncounter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEncounter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Parameters", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEncounter
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
				return ErrInvalidLengthEncounter
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEncounter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Parameters == nil {
				m.Parameters = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowEncounter
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowEncounter
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthEncounter
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthEncounter
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowEncounter
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthEncounter
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthEncounter
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipEncounter(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthEncounter
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Parameters[mapkey] = mapvalue
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImageId", wireType)
			}
			m.ImageId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEncounter
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
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEncounter
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
				return ErrInvalidLengthEncounter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEncounter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEncounter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEncounter
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
func (m *EncounterWithImage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEncounter
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
			return fmt.Errorf("proto: EncounterWithImage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EncounterWithImage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Encounter", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEncounter
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
				return ErrInvalidLengthEncounter
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEncounter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Encounter == nil {
				m.Encounter = &Encounter{}
			}
			if err := m.Encounter.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Image", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEncounter
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
				return ErrInvalidLengthEncounter
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEncounter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Image = append(m.Image[:0], dAtA[iNdEx:postIndex]...)
			if m.Image == nil {
				m.Image = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEncounter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEncounter
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
func skipEncounter(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEncounter
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
					return 0, ErrIntOverflowEncounter
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
					return 0, ErrIntOverflowEncounter
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
				return 0, ErrInvalidLengthEncounter
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEncounter
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEncounter
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEncounter        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEncounter          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEncounter = fmt.Errorf("proto: unexpected end of group")
)
