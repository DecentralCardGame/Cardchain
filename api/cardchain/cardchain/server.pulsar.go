// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package cardchain

import (
	fmt "fmt"
	io "io"
	reflect "reflect"
	sync "sync"

	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

var (
	md_Server                protoreflect.MessageDescriptor
	fd_Server_reporter       protoreflect.FieldDescriptor
	fd_Server_invalidReports protoreflect.FieldDescriptor
	fd_Server_validReports   protoreflect.FieldDescriptor
)

func init() {
	file_cardchain_cardchain_server_proto_init()
	md_Server = File_cardchain_cardchain_server_proto.Messages().ByName("Server")
	fd_Server_reporter = md_Server.Fields().ByName("reporter")
	fd_Server_invalidReports = md_Server.Fields().ByName("invalidReports")
	fd_Server_validReports = md_Server.Fields().ByName("validReports")
}

var _ protoreflect.Message = (*fastReflection_Server)(nil)

type fastReflection_Server Server

func (x *Server) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Server)(x)
}

func (x *Server) slowProtoReflect() protoreflect.Message {
	mi := &file_cardchain_cardchain_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Server_messageType fastReflection_Server_messageType
var _ protoreflect.MessageType = fastReflection_Server_messageType{}

type fastReflection_Server_messageType struct{}

func (x fastReflection_Server_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Server)(nil)
}
func (x fastReflection_Server_messageType) New() protoreflect.Message {
	return new(fastReflection_Server)
}
func (x fastReflection_Server_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Server
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Server) Descriptor() protoreflect.MessageDescriptor {
	return md_Server
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Server) Type() protoreflect.MessageType {
	return _fastReflection_Server_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Server) New() protoreflect.Message {
	return new(fastReflection_Server)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Server) Interface() protoreflect.ProtoMessage {
	return (*Server)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Server) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Reporter != "" {
		value := protoreflect.ValueOfString(x.Reporter)
		if !f(fd_Server_reporter, value) {
			return
		}
	}
	if x.InvalidReports != uint64(0) {
		value := protoreflect.ValueOfUint64(x.InvalidReports)
		if !f(fd_Server_invalidReports, value) {
			return
		}
	}
	if x.ValidReports != uint64(0) {
		value := protoreflect.ValueOfUint64(x.ValidReports)
		if !f(fd_Server_validReports, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Server) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cardchain.cardchain.Server.reporter":
		return x.Reporter != ""
	case "cardchain.cardchain.Server.invalidReports":
		return x.InvalidReports != uint64(0)
	case "cardchain.cardchain.Server.validReports":
		return x.ValidReports != uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cardchain.cardchain.Server"))
		}
		panic(fmt.Errorf("message cardchain.cardchain.Server does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Server) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cardchain.cardchain.Server.reporter":
		x.Reporter = ""
	case "cardchain.cardchain.Server.invalidReports":
		x.InvalidReports = uint64(0)
	case "cardchain.cardchain.Server.validReports":
		x.ValidReports = uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cardchain.cardchain.Server"))
		}
		panic(fmt.Errorf("message cardchain.cardchain.Server does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Server) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cardchain.cardchain.Server.reporter":
		value := x.Reporter
		return protoreflect.ValueOfString(value)
	case "cardchain.cardchain.Server.invalidReports":
		value := x.InvalidReports
		return protoreflect.ValueOfUint64(value)
	case "cardchain.cardchain.Server.validReports":
		value := x.ValidReports
		return protoreflect.ValueOfUint64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cardchain.cardchain.Server"))
		}
		panic(fmt.Errorf("message cardchain.cardchain.Server does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Server) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cardchain.cardchain.Server.reporter":
		x.Reporter = value.Interface().(string)
	case "cardchain.cardchain.Server.invalidReports":
		x.InvalidReports = value.Uint()
	case "cardchain.cardchain.Server.validReports":
		x.ValidReports = value.Uint()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cardchain.cardchain.Server"))
		}
		panic(fmt.Errorf("message cardchain.cardchain.Server does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Server) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cardchain.cardchain.Server.reporter":
		panic(fmt.Errorf("field reporter of message cardchain.cardchain.Server is not mutable"))
	case "cardchain.cardchain.Server.invalidReports":
		panic(fmt.Errorf("field invalidReports of message cardchain.cardchain.Server is not mutable"))
	case "cardchain.cardchain.Server.validReports":
		panic(fmt.Errorf("field validReports of message cardchain.cardchain.Server is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cardchain.cardchain.Server"))
		}
		panic(fmt.Errorf("message cardchain.cardchain.Server does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Server) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cardchain.cardchain.Server.reporter":
		return protoreflect.ValueOfString("")
	case "cardchain.cardchain.Server.invalidReports":
		return protoreflect.ValueOfUint64(uint64(0))
	case "cardchain.cardchain.Server.validReports":
		return protoreflect.ValueOfUint64(uint64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cardchain.cardchain.Server"))
		}
		panic(fmt.Errorf("message cardchain.cardchain.Server does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Server) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cardchain.cardchain.Server", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Server) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Server) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Server) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Server) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Server)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Reporter)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.InvalidReports != 0 {
			n += 1 + runtime.Sov(uint64(x.InvalidReports))
		}
		if x.ValidReports != 0 {
			n += 1 + runtime.Sov(uint64(x.ValidReports))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Server)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.ValidReports != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.ValidReports))
			i--
			dAtA[i] = 0x18
		}
		if x.InvalidReports != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.InvalidReports))
			i--
			dAtA[i] = 0x10
		}
		if len(x.Reporter) > 0 {
			i -= len(x.Reporter)
			copy(dAtA[i:], x.Reporter)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Reporter)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Server)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Server: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Server: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Reporter", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Reporter = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field InvalidReports", wireType)
				}
				x.InvalidReports = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.InvalidReports |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ValidReports", wireType)
				}
				x.ValidReports = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.ValidReports |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: cardchain/cardchain/server.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reporter       string `protobuf:"bytes,1,opt,name=reporter,proto3" json:"reporter,omitempty"`
	InvalidReports uint64 `protobuf:"varint,2,opt,name=invalidReports,proto3" json:"invalidReports,omitempty"`
	ValidReports   uint64 `protobuf:"varint,3,opt,name=validReports,proto3" json:"validReports,omitempty"`
}

func (x *Server) Reset() {
	*x = Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cardchain_cardchain_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_cardchain_cardchain_server_proto_rawDescGZIP(), []int{0}
}

func (x *Server) GetReporter() string {
	if x != nil {
		return x.Reporter
	}
	return ""
}

func (x *Server) GetInvalidReports() uint64 {
	if x != nil {
		return x.InvalidReports
	}
	return 0
}

func (x *Server) GetValidReports() uint64 {
	if x != nil {
		return x.ValidReports
	}
	return 0
}

var File_cardchain_cardchain_server_proto protoreflect.FileDescriptor

var file_cardchain_cardchain_server_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x61, 0x72, 0x64, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x63, 0x61, 0x72, 0x64,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x63, 0x61, 0x72, 0x64, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x63, 0x61,
	0x72, 0x64, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x22, 0x70, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x12, 0x26, 0x0a,
	0x0e, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x42, 0xd3, 0x01, 0x0a, 0x17, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x63, 0x61, 0x72, 0x64,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x42, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x44, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x43, 0x61, 0x72, 0x64, 0x47, 0x61,
	0x6d, 0x65, 0x2f, 0x63, 0x61, 0x72, 0x64, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x61, 0x72, 0x64, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x63, 0x61, 0x72, 0x64, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0xa2, 0x02, 0x03, 0x43, 0x43, 0x58, 0xaa, 0x02, 0x13, 0x43, 0x61, 0x72,
	0x64, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0xca, 0x02, 0x13, 0x43, 0x61, 0x72, 0x64, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5c, 0x43, 0x61, 0x72,
	0x64, 0x63, 0x68, 0x61, 0x69, 0x6e, 0xe2, 0x02, 0x1f, 0x43, 0x61, 0x72, 0x64, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x5c, 0x43, 0x61, 0x72, 0x64, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14, 0x43, 0x61, 0x72, 0x64, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x3a, 0x3a, 0x43, 0x61, 0x72, 0x64, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cardchain_cardchain_server_proto_rawDescOnce sync.Once
	file_cardchain_cardchain_server_proto_rawDescData = file_cardchain_cardchain_server_proto_rawDesc
)

func file_cardchain_cardchain_server_proto_rawDescGZIP() []byte {
	file_cardchain_cardchain_server_proto_rawDescOnce.Do(func() {
		file_cardchain_cardchain_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_cardchain_cardchain_server_proto_rawDescData)
	})
	return file_cardchain_cardchain_server_proto_rawDescData
}

var file_cardchain_cardchain_server_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cardchain_cardchain_server_proto_goTypes = []interface{}{
	(*Server)(nil), // 0: cardchain.cardchain.Server
}
var file_cardchain_cardchain_server_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cardchain_cardchain_server_proto_init() }
func file_cardchain_cardchain_server_proto_init() {
	if File_cardchain_cardchain_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cardchain_cardchain_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cardchain_cardchain_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cardchain_cardchain_server_proto_goTypes,
		DependencyIndexes: file_cardchain_cardchain_server_proto_depIdxs,
		MessageInfos:      file_cardchain_cardchain_server_proto_msgTypes,
	}.Build()
	File_cardchain_cardchain_server_proto = out.File
	file_cardchain_cardchain_server_proto_rawDesc = nil
	file_cardchain_cardchain_server_proto_goTypes = nil
	file_cardchain_cardchain_server_proto_depIdxs = nil
}
