// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/db.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Node struct {
	Key                  string                     `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Type                 string                     `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Attributes           map[string]*AttributeValue `protobuf:"bytes,3,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_db_f9e8e82a2f13d96c, []int{0}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (dst *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(dst, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Node) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Node) GetAttributes() map[string]*AttributeValue {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type AttributeValue struct {
	// Types that are valid to be assigned to Value:
	//	*AttributeValue_StringValue
	//	*AttributeValue_IntValue
	//	*AttributeValue_FloatValue
	//	*AttributeValue_DoubleValue
	//	*AttributeValue_BooleanValue
	Value                isAttributeValue_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *AttributeValue) Reset()         { *m = AttributeValue{} }
func (m *AttributeValue) String() string { return proto.CompactTextString(m) }
func (*AttributeValue) ProtoMessage()    {}
func (*AttributeValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_db_f9e8e82a2f13d96c, []int{1}
}
func (m *AttributeValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttributeValue.Unmarshal(m, b)
}
func (m *AttributeValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttributeValue.Marshal(b, m, deterministic)
}
func (dst *AttributeValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttributeValue.Merge(dst, src)
}
func (m *AttributeValue) XXX_Size() int {
	return xxx_messageInfo_AttributeValue.Size(m)
}
func (m *AttributeValue) XXX_DiscardUnknown() {
	xxx_messageInfo_AttributeValue.DiscardUnknown(m)
}

var xxx_messageInfo_AttributeValue proto.InternalMessageInfo

type isAttributeValue_Value interface {
	isAttributeValue_Value()
}

type AttributeValue_StringValue struct {
	StringValue string `protobuf:"bytes,1,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type AttributeValue_IntValue struct {
	IntValue int64 `protobuf:"varint,2,opt,name=int_value,json=intValue,proto3,oneof"`
}

type AttributeValue_FloatValue struct {
	FloatValue float32 `protobuf:"fixed32,3,opt,name=float_value,json=floatValue,proto3,oneof"`
}

type AttributeValue_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,4,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type AttributeValue_BooleanValue struct {
	BooleanValue bool `protobuf:"varint,5,opt,name=boolean_value,json=booleanValue,proto3,oneof"`
}

func (*AttributeValue_StringValue) isAttributeValue_Value() {}

func (*AttributeValue_IntValue) isAttributeValue_Value() {}

func (*AttributeValue_FloatValue) isAttributeValue_Value() {}

func (*AttributeValue_DoubleValue) isAttributeValue_Value() {}

func (*AttributeValue_BooleanValue) isAttributeValue_Value() {}

func (m *AttributeValue) GetValue() isAttributeValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *AttributeValue) GetStringValue() string {
	if x, ok := m.GetValue().(*AttributeValue_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *AttributeValue) GetIntValue() int64 {
	if x, ok := m.GetValue().(*AttributeValue_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (m *AttributeValue) GetFloatValue() float32 {
	if x, ok := m.GetValue().(*AttributeValue_FloatValue); ok {
		return x.FloatValue
	}
	return 0
}

func (m *AttributeValue) GetDoubleValue() float64 {
	if x, ok := m.GetValue().(*AttributeValue_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (m *AttributeValue) GetBooleanValue() bool {
	if x, ok := m.GetValue().(*AttributeValue_BooleanValue); ok {
		return x.BooleanValue
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AttributeValue) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AttributeValue_OneofMarshaler, _AttributeValue_OneofUnmarshaler, _AttributeValue_OneofSizer, []interface{}{
		(*AttributeValue_StringValue)(nil),
		(*AttributeValue_IntValue)(nil),
		(*AttributeValue_FloatValue)(nil),
		(*AttributeValue_DoubleValue)(nil),
		(*AttributeValue_BooleanValue)(nil),
	}
}

func _AttributeValue_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AttributeValue)
	// value
	switch x := m.Value.(type) {
	case *AttributeValue_StringValue:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.StringValue)
	case *AttributeValue_IntValue:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.IntValue))
	case *AttributeValue_FloatValue:
		b.EncodeVarint(3<<3 | proto.WireFixed32)
		b.EncodeFixed32(uint64(math.Float32bits(x.FloatValue)))
	case *AttributeValue_DoubleValue:
		b.EncodeVarint(4<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.DoubleValue))
	case *AttributeValue_BooleanValue:
		t := uint64(0)
		if x.BooleanValue {
			t = 1
		}
		b.EncodeVarint(5<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case nil:
	default:
		return fmt.Errorf("AttributeValue.Value has unexpected type %T", x)
	}
	return nil
}

func _AttributeValue_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AttributeValue)
	switch tag {
	case 1: // value.string_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &AttributeValue_StringValue{x}
		return true, err
	case 2: // value.int_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &AttributeValue_IntValue{int64(x)}
		return true, err
	case 3: // value.float_value
		if wire != proto.WireFixed32 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed32()
		m.Value = &AttributeValue_FloatValue{math.Float32frombits(uint32(x))}
		return true, err
	case 4: // value.double_value
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Value = &AttributeValue_DoubleValue{math.Float64frombits(x)}
		return true, err
	case 5: // value.boolean_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &AttributeValue_BooleanValue{x != 0}
		return true, err
	default:
		return false, nil
	}
}

func _AttributeValue_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AttributeValue)
	// value
	switch x := m.Value.(type) {
	case *AttributeValue_StringValue:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.StringValue)))
		n += len(x.StringValue)
	case *AttributeValue_IntValue:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.IntValue))
	case *AttributeValue_FloatValue:
		n += 1 // tag and wire
		n += 4
	case *AttributeValue_DoubleValue:
		n += 1 // tag and wire
		n += 8
	case *AttributeValue_BooleanValue:
		n += 1 // tag and wire
		n += 1
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Relation struct {
	Key                  string                     `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Type                 string                     `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	From                 string                     `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	To                   string                     `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	Attributes           map[string]*AttributeValue `protobuf:"bytes,5,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *Relation) Reset()         { *m = Relation{} }
func (m *Relation) String() string { return proto.CompactTextString(m) }
func (*Relation) ProtoMessage()    {}
func (*Relation) Descriptor() ([]byte, []int) {
	return fileDescriptor_db_f9e8e82a2f13d96c, []int{2}
}
func (m *Relation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Relation.Unmarshal(m, b)
}
func (m *Relation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Relation.Marshal(b, m, deterministic)
}
func (dst *Relation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Relation.Merge(dst, src)
}
func (m *Relation) XXX_Size() int {
	return xxx_messageInfo_Relation.Size(m)
}
func (m *Relation) XXX_DiscardUnknown() {
	xxx_messageInfo_Relation.DiscardUnknown(m)
}

var xxx_messageInfo_Relation proto.InternalMessageInfo

func (m *Relation) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Relation) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Relation) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Relation) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Relation) GetAttributes() map[string]*AttributeValue {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func init() {
	proto.RegisterType((*Node)(nil), "proto.Node")
	proto.RegisterMapType((map[string]*AttributeValue)(nil), "proto.Node.AttributesEntry")
	proto.RegisterType((*AttributeValue)(nil), "proto.AttributeValue")
	proto.RegisterType((*Relation)(nil), "proto.Relation")
	proto.RegisterMapType((map[string]*AttributeValue)(nil), "proto.Relation.AttributesEntry")
}

func init() { proto.RegisterFile("proto/db.proto", fileDescriptor_db_f9e8e82a2f13d96c) }

var fileDescriptor_db_f9e8e82a2f13d96c = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x91, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0xb3, 0xf9, 0xa3, 0xcd, 0xa4, 0x56, 0x59, 0x10, 0x8a, 0x22, 0xc6, 0x8a, 0x10, 0x10,
	0x2a, 0xd4, 0x8b, 0xe8, 0x41, 0x14, 0x84, 0x9c, 0x3c, 0x2c, 0xe2, 0xb5, 0x24, 0x74, 0x2b, 0xc1,
	0xb8, 0x5b, 0xd2, 0x89, 0x90, 0xa7, 0xf3, 0x05, 0x7c, 0x13, 0x5f, 0x42, 0x32, 0xbb, 0x8d, 0x56,
	0x3c, 0x78, 0xf2, 0x94, 0xe1, 0xfb, 0x7e, 0xf9, 0xbe, 0x49, 0x06, 0x06, 0x8b, 0x4a, 0xa3, 0x3e,
	0x9b, 0xe5, 0x63, 0x1a, 0x78, 0x40, 0x8f, 0xd1, 0x1b, 0x03, 0xff, 0x5e, 0xcf, 0x24, 0xdf, 0x01,
	0xef, 0x59, 0x36, 0x43, 0x16, 0xb3, 0x24, 0x14, 0xed, 0xc8, 0x39, 0xf8, 0xd8, 0x2c, 0xe4, 0xd0,
	0x25, 0x89, 0x66, 0x7e, 0x05, 0x90, 0x21, 0x56, 0x45, 0x5e, 0xa3, 0x5c, 0x0e, 0xbd, 0xd8, 0x4b,
	0xa2, 0xc9, 0xbe, 0x49, 0x1c, 0xb7, 0x31, 0xe3, 0x9b, 0xce, 0xbd, 0x53, 0x58, 0x35, 0xe2, 0x1b,
	0xbe, 0xf7, 0x00, 0xdb, 0x3f, 0xec, 0x5f, 0x5a, 0x4f, 0x21, 0x78, 0xcd, 0xca, 0xda, 0xd4, 0x46,
	0x93, 0x5d, 0x1b, 0xde, 0xbd, 0xf8, 0xd8, 0x9a, 0xc2, 0x30, 0x97, 0xee, 0x05, 0x1b, 0xbd, 0x33,
	0x18, 0xac, 0xbb, 0xfc, 0x18, 0xfa, 0x4b, 0xac, 0x0a, 0xf5, 0x34, 0x35, 0x51, 0x14, 0x9f, 0x3a,
	0x22, 0x32, 0xaa, 0x81, 0x0e, 0x20, 0x2c, 0x14, 0x4e, 0xbf, 0xca, 0xbc, 0xd4, 0x11, 0xbd, 0x42,
	0xa1, 0xb1, 0x8f, 0x20, 0x9a, 0x97, 0x3a, 0x5b, 0x01, 0x5e, 0xcc, 0x12, 0x37, 0x75, 0x04, 0x90,
	0xd8, 0xd5, 0xcc, 0x74, 0x9d, 0x97, 0xd2, 0x32, 0x7e, 0xcc, 0x12, 0xd6, 0xd6, 0x18, 0xd5, 0x40,
	0x27, 0xb0, 0x95, 0x6b, 0x5d, 0xca, 0x4c, 0x59, 0x2a, 0x88, 0x59, 0xd2, 0x4b, 0x1d, 0xd1, 0xb7,
	0x32, 0x61, 0xb7, 0x9b, 0xf6, 0xb3, 0x47, 0x1f, 0x0c, 0x7a, 0x42, 0x96, 0x19, 0x16, 0x5a, 0xfd,
	0xf1, 0x28, 0x1c, 0xfc, 0x79, 0xa5, 0x5f, 0x68, 0xc7, 0x50, 0xd0, 0xcc, 0x07, 0xe0, 0xa2, 0xa6,
	0x8d, 0x42, 0xe1, 0xa2, 0xe6, 0xd7, 0x6b, 0x87, 0x0b, 0xe8, 0x70, 0x87, 0xf6, 0xdf, 0xae, 0xea,
	0xfe, 0xff, 0x78, 0xf9, 0x06, 0x01, 0xe7, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x70, 0xef, 0x32,
	0x5f, 0x9e, 0x02, 0x00, 0x00,
}
