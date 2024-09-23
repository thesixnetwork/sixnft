// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftmngr/v090/nft_attribute_value.proto

package v090

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
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

type NftAttributeValue struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*NftAttributeValue_NumberAttributeValue
	//	*NftAttributeValue_StringAttributeValue
	//	*NftAttributeValue_BooleanAttributeValue
	//	*NftAttributeValue_FloatAttributeValue
	Value               isNftAttributeValue_Value `protobuf_oneof:"value"`
	HiddenToMarketplace bool                      `protobuf:"varint,6,opt,name=hidden_to_marketplace,json=hiddenToMarketplace,proto3" json:"hidden_to_marketplace,omitempty"`
}

func (m *NftAttributeValue) Reset()         { *m = NftAttributeValue{} }
func (m *NftAttributeValue) String() string { return proto.CompactTextString(m) }
func (*NftAttributeValue) ProtoMessage()    {}
func (*NftAttributeValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_85809ba742864c2a, []int{0}
}
func (m *NftAttributeValue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NftAttributeValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NftAttributeValue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NftAttributeValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NftAttributeValue.Merge(m, src)
}
func (m *NftAttributeValue) XXX_Size() int {
	return m.Size()
}
func (m *NftAttributeValue) XXX_DiscardUnknown() {
	xxx_messageInfo_NftAttributeValue.DiscardUnknown(m)
}

var xxx_messageInfo_NftAttributeValue proto.InternalMessageInfo

type isNftAttributeValue_Value interface {
	isNftAttributeValue_Value()
	MarshalTo([]byte) (int, error)
	Size() int
}

type NftAttributeValue_NumberAttributeValue struct {
	NumberAttributeValue *NumberAttributeValue `protobuf:"bytes,2,opt,name=number_attribute_value,json=numberAttributeValue,proto3,oneof" json:"number_attribute_value,omitempty"`
}
type NftAttributeValue_StringAttributeValue struct {
	StringAttributeValue *StringAttributeValue `protobuf:"bytes,3,opt,name=string_attribute_value,json=stringAttributeValue,proto3,oneof" json:"string_attribute_value,omitempty"`
}
type NftAttributeValue_BooleanAttributeValue struct {
	BooleanAttributeValue *BooleanAttributeValue `protobuf:"bytes,4,opt,name=boolean_attribute_value,json=booleanAttributeValue,proto3,oneof" json:"boolean_attribute_value,omitempty"`
}
type NftAttributeValue_FloatAttributeValue struct {
	FloatAttributeValue *FloatAttributeValue `protobuf:"bytes,5,opt,name=float_attribute_value,json=floatAttributeValue,proto3,oneof" json:"float_attribute_value,omitempty"`
}

func (*NftAttributeValue_NumberAttributeValue) isNftAttributeValue_Value()  {}
func (*NftAttributeValue_StringAttributeValue) isNftAttributeValue_Value()  {}
func (*NftAttributeValue_BooleanAttributeValue) isNftAttributeValue_Value() {}
func (*NftAttributeValue_FloatAttributeValue) isNftAttributeValue_Value()   {}

func (m *NftAttributeValue) GetValue() isNftAttributeValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *NftAttributeValue) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NftAttributeValue) GetNumberAttributeValue() *NumberAttributeValue {
	if x, ok := m.GetValue().(*NftAttributeValue_NumberAttributeValue); ok {
		return x.NumberAttributeValue
	}
	return nil
}

func (m *NftAttributeValue) GetStringAttributeValue() *StringAttributeValue {
	if x, ok := m.GetValue().(*NftAttributeValue_StringAttributeValue); ok {
		return x.StringAttributeValue
	}
	return nil
}

func (m *NftAttributeValue) GetBooleanAttributeValue() *BooleanAttributeValue {
	if x, ok := m.GetValue().(*NftAttributeValue_BooleanAttributeValue); ok {
		return x.BooleanAttributeValue
	}
	return nil
}

func (m *NftAttributeValue) GetFloatAttributeValue() *FloatAttributeValue {
	if x, ok := m.GetValue().(*NftAttributeValue_FloatAttributeValue); ok {
		return x.FloatAttributeValue
	}
	return nil
}

func (m *NftAttributeValue) GetHiddenToMarketplace() bool {
	if m != nil {
		return m.HiddenToMarketplace
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*NftAttributeValue) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*NftAttributeValue_NumberAttributeValue)(nil),
		(*NftAttributeValue_StringAttributeValue)(nil),
		(*NftAttributeValue_BooleanAttributeValue)(nil),
		(*NftAttributeValue_FloatAttributeValue)(nil),
	}
}

type NumberAttributeValue struct {
	Value uint64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *NumberAttributeValue) Reset()         { *m = NumberAttributeValue{} }
func (m *NumberAttributeValue) String() string { return proto.CompactTextString(m) }
func (*NumberAttributeValue) ProtoMessage()    {}
func (*NumberAttributeValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_85809ba742864c2a, []int{1}
}
func (m *NumberAttributeValue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NumberAttributeValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NumberAttributeValue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NumberAttributeValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NumberAttributeValue.Merge(m, src)
}
func (m *NumberAttributeValue) XXX_Size() int {
	return m.Size()
}
func (m *NumberAttributeValue) XXX_DiscardUnknown() {
	xxx_messageInfo_NumberAttributeValue.DiscardUnknown(m)
}

var xxx_messageInfo_NumberAttributeValue proto.InternalMessageInfo

func (m *NumberAttributeValue) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type StringAttributeValue struct {
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *StringAttributeValue) Reset()         { *m = StringAttributeValue{} }
func (m *StringAttributeValue) String() string { return proto.CompactTextString(m) }
func (*StringAttributeValue) ProtoMessage()    {}
func (*StringAttributeValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_85809ba742864c2a, []int{2}
}
func (m *StringAttributeValue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StringAttributeValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StringAttributeValue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StringAttributeValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringAttributeValue.Merge(m, src)
}
func (m *StringAttributeValue) XXX_Size() int {
	return m.Size()
}
func (m *StringAttributeValue) XXX_DiscardUnknown() {
	xxx_messageInfo_StringAttributeValue.DiscardUnknown(m)
}

var xxx_messageInfo_StringAttributeValue proto.InternalMessageInfo

func (m *StringAttributeValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type BooleanAttributeValue struct {
	Value bool `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *BooleanAttributeValue) Reset()         { *m = BooleanAttributeValue{} }
func (m *BooleanAttributeValue) String() string { return proto.CompactTextString(m) }
func (*BooleanAttributeValue) ProtoMessage()    {}
func (*BooleanAttributeValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_85809ba742864c2a, []int{3}
}
func (m *BooleanAttributeValue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BooleanAttributeValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BooleanAttributeValue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BooleanAttributeValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BooleanAttributeValue.Merge(m, src)
}
func (m *BooleanAttributeValue) XXX_Size() int {
	return m.Size()
}
func (m *BooleanAttributeValue) XXX_DiscardUnknown() {
	xxx_messageInfo_BooleanAttributeValue.DiscardUnknown(m)
}

var xxx_messageInfo_BooleanAttributeValue proto.InternalMessageInfo

func (m *BooleanAttributeValue) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

type FloatAttributeValue struct {
	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *FloatAttributeValue) Reset()         { *m = FloatAttributeValue{} }
func (m *FloatAttributeValue) String() string { return proto.CompactTextString(m) }
func (*FloatAttributeValue) ProtoMessage()    {}
func (*FloatAttributeValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_85809ba742864c2a, []int{4}
}
func (m *FloatAttributeValue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FloatAttributeValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FloatAttributeValue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FloatAttributeValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FloatAttributeValue.Merge(m, src)
}
func (m *FloatAttributeValue) XXX_Size() int {
	return m.Size()
}
func (m *FloatAttributeValue) XXX_DiscardUnknown() {
	xxx_messageInfo_FloatAttributeValue.DiscardUnknown(m)
}

var xxx_messageInfo_FloatAttributeValue proto.InternalMessageInfo

func (m *FloatAttributeValue) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*NftAttributeValue)(nil), "thesixnetwork.sixnft.nftmngr.v090.NftAttributeValue")
	proto.RegisterType((*NumberAttributeValue)(nil), "thesixnetwork.sixnft.nftmngr.v090.NumberAttributeValue")
	proto.RegisterType((*StringAttributeValue)(nil), "thesixnetwork.sixnft.nftmngr.v090.StringAttributeValue")
	proto.RegisterType((*BooleanAttributeValue)(nil), "thesixnetwork.sixnft.nftmngr.v090.BooleanAttributeValue")
	proto.RegisterType((*FloatAttributeValue)(nil), "thesixnetwork.sixnft.nftmngr.v090.FloatAttributeValue")
}

func init() {
	proto.RegisterFile("nftmngr/v090/nft_attribute_value.proto", fileDescriptor_85809ba742864c2a)
}

var fileDescriptor_85809ba742864c2a = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x4f, 0x6f, 0xda, 0x30,
	0x18, 0xc6, 0xe3, 0x2d, 0x30, 0xf0, 0x4e, 0x33, 0x64, 0xcb, 0x76, 0x88, 0x18, 0x87, 0x09, 0x69,
	0x5b, 0x8c, 0xd8, 0xb4, 0x3f, 0xc7, 0x71, 0x98, 0x7a, 0x29, 0x87, 0xb4, 0xea, 0xa1, 0x97, 0xc8,
	0x01, 0x27, 0x44, 0x24, 0x36, 0x72, 0x1c, 0x0a, 0xdf, 0xa2, 0x1f, 0xa3, 0x1f, 0xa5, 0x47, 0x8e,
	0x3d, 0x56, 0xf0, 0x45, 0xaa, 0x38, 0x44, 0x28, 0x4d, 0xa4, 0x72, 0xb3, 0xfc, 0x3e, 0xcf, 0xf3,
	0x7b, 0xf3, 0x3a, 0x2f, 0xfc, 0xc2, 0x7c, 0x19, 0xb3, 0x40, 0xe0, 0xd5, 0xf0, 0xef, 0x10, 0x33,
	0x5f, 0xba, 0x44, 0x4a, 0x11, 0x7a, 0xa9, 0xa4, 0xee, 0x8a, 0x44, 0x29, 0xb5, 0x97, 0x82, 0x4b,
	0x8e, 0x3e, 0xcb, 0x39, 0x4d, 0xc2, 0x35, 0xa3, 0xf2, 0x86, 0x8b, 0x85, 0x9d, 0x1d, 0x7d, 0x69,
	0x1f, 0xcc, 0x76, 0x66, 0xfe, 0xf4, 0x31, 0xe0, 0x3c, 0x88, 0x28, 0x56, 0x06, 0x2f, 0xf5, 0x31,
	0x61, 0x9b, 0xdc, 0xdd, 0xbf, 0xd3, 0xe1, 0xbb, 0x89, 0x2f, 0xff, 0x15, 0xd1, 0x57, 0x59, 0x32,
	0x42, 0x50, 0x67, 0x24, 0xa6, 0x26, 0xe8, 0x81, 0x41, 0xdb, 0x51, 0x67, 0xc4, 0xe1, 0x7b, 0x96,
	0xc6, 0x1e, 0x15, 0xcf, 0xfb, 0x30, 0x5f, 0xf5, 0xc0, 0xe0, 0xed, 0xe8, 0xb7, 0xfd, 0x62, 0x23,
	0xf6, 0x44, 0x05, 0x94, 0x61, 0x67, 0x9a, 0xd3, 0x65, 0x35, 0xf7, 0x19, 0x30, 0x91, 0x22, 0x64,
	0x41, 0x05, 0xf8, 0xfa, 0x64, 0xe0, 0x85, 0x0a, 0xa8, 0x02, 0x93, 0x9a, 0x7b, 0x24, 0xe0, 0x07,
	0x8f, 0xf3, 0x88, 0x12, 0x56, 0x21, 0xea, 0x8a, 0xf8, 0xe7, 0x04, 0xe2, 0x38, 0x4f, 0xa8, 0x20,
	0x0d, 0xaf, 0xae, 0x80, 0x22, 0x68, 0xf8, 0x11, 0x27, 0x95, 0xc7, 0x35, 0x1b, 0x8a, 0xf8, 0xeb,
	0x04, 0xe2, 0xff, 0xcc, 0x5f, 0xe1, 0x75, 0xfc, 0xea, 0x35, 0x1a, 0x41, 0x63, 0x1e, 0xce, 0x66,
	0x94, 0xb9, 0x92, 0xbb, 0x31, 0x11, 0x0b, 0x2a, 0x97, 0x11, 0x99, 0x52, 0xb3, 0xd9, 0x03, 0x83,
	0x96, 0xd3, 0xc9, 0x8b, 0x97, 0xfc, 0xfc, 0x58, 0x1a, 0xbf, 0x81, 0x0d, 0xd5, 0x51, 0xff, 0x1b,
	0xec, 0xd6, 0xbd, 0x1f, 0xea, 0x1e, 0x04, 0xea, 0x6f, 0xd1, 0x9d, 0xa3, 0xba, 0x6e, 0xf8, 0x65,
	0x75, 0xbb, 0x50, 0x7f, 0x87, 0x46, 0xed, 0xe0, 0xca, 0xf2, 0x56, 0x21, 0xff, 0x0a, 0x3b, 0x35,
	0x5f, 0x5d, 0x16, 0x83, 0x83, 0x78, 0x3c, 0xb9, 0xdf, 0x59, 0x60, 0xbb, 0xb3, 0xc0, 0xe3, 0xce,
	0x02, 0xb7, 0x7b, 0x4b, 0xdb, 0xee, 0x2d, 0xed, 0x61, 0x6f, 0x69, 0xd7, 0x3f, 0x83, 0x50, 0xce,
	0x53, 0xcf, 0x9e, 0xf2, 0x18, 0x97, 0xe6, 0x8c, 0xf3, 0x39, 0xe3, 0x35, 0x2e, 0x96, 0x50, 0x6e,
	0x96, 0x34, 0x51, 0xab, 0xe8, 0x35, 0xd5, 0xe6, 0xfc, 0x78, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x7b,
	0x35, 0xd6, 0xb9, 0xa1, 0x03, 0x00, 0x00,
}

func (m *NftAttributeValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NftAttributeValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NftAttributeValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.HiddenToMarketplace {
		i--
		if m.HiddenToMarketplace {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.Value != nil {
		{
			size := m.Value.Size()
			i -= size
			if _, err := m.Value.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintNftAttributeValue(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *NftAttributeValue_NumberAttributeValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NftAttributeValue_NumberAttributeValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.NumberAttributeValue != nil {
		{
			size, err := m.NumberAttributeValue.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintNftAttributeValue(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *NftAttributeValue_StringAttributeValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NftAttributeValue_StringAttributeValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.StringAttributeValue != nil {
		{
			size, err := m.StringAttributeValue.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintNftAttributeValue(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *NftAttributeValue_BooleanAttributeValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NftAttributeValue_BooleanAttributeValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.BooleanAttributeValue != nil {
		{
			size, err := m.BooleanAttributeValue.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintNftAttributeValue(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *NftAttributeValue_FloatAttributeValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NftAttributeValue_FloatAttributeValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.FloatAttributeValue != nil {
		{
			size, err := m.FloatAttributeValue.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintNftAttributeValue(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	return len(dAtA) - i, nil
}
func (m *NumberAttributeValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NumberAttributeValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NumberAttributeValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		i = encodeVarintNftAttributeValue(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *StringAttributeValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StringAttributeValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StringAttributeValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintNftAttributeValue(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BooleanAttributeValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BooleanAttributeValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BooleanAttributeValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value {
		i--
		if m.Value {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *FloatAttributeValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FloatAttributeValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FloatAttributeValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Value))))
		i--
		dAtA[i] = 0x9
	}
	return len(dAtA) - i, nil
}

func encodeVarintNftAttributeValue(dAtA []byte, offset int, v uint64) int {
	offset -= sovNftAttributeValue(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NftAttributeValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovNftAttributeValue(uint64(l))
	}
	if m.Value != nil {
		n += m.Value.Size()
	}
	if m.HiddenToMarketplace {
		n += 2
	}
	return n
}

func (m *NftAttributeValue_NumberAttributeValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NumberAttributeValue != nil {
		l = m.NumberAttributeValue.Size()
		n += 1 + l + sovNftAttributeValue(uint64(l))
	}
	return n
}
func (m *NftAttributeValue_StringAttributeValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StringAttributeValue != nil {
		l = m.StringAttributeValue.Size()
		n += 1 + l + sovNftAttributeValue(uint64(l))
	}
	return n
}
func (m *NftAttributeValue_BooleanAttributeValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BooleanAttributeValue != nil {
		l = m.BooleanAttributeValue.Size()
		n += 1 + l + sovNftAttributeValue(uint64(l))
	}
	return n
}
func (m *NftAttributeValue_FloatAttributeValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FloatAttributeValue != nil {
		l = m.FloatAttributeValue.Size()
		n += 1 + l + sovNftAttributeValue(uint64(l))
	}
	return n
}
func (m *NumberAttributeValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != 0 {
		n += 1 + sovNftAttributeValue(uint64(m.Value))
	}
	return n
}

func (m *StringAttributeValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovNftAttributeValue(uint64(l))
	}
	return n
}

func (m *BooleanAttributeValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value {
		n += 2
	}
	return n
}

func (m *FloatAttributeValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != 0 {
		n += 9
	}
	return n
}

func sovNftAttributeValue(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNftAttributeValue(x uint64) (n int) {
	return sovNftAttributeValue(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NftAttributeValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNftAttributeValue
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
			return fmt.Errorf("proto: NftAttributeValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NftAttributeValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftAttributeValue
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
				return ErrInvalidLengthNftAttributeValue
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftAttributeValue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumberAttributeValue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftAttributeValue
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
				return ErrInvalidLengthNftAttributeValue
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNftAttributeValue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &NumberAttributeValue{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &NftAttributeValue_NumberAttributeValue{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StringAttributeValue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftAttributeValue
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
				return ErrInvalidLengthNftAttributeValue
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNftAttributeValue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &StringAttributeValue{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &NftAttributeValue_StringAttributeValue{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BooleanAttributeValue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftAttributeValue
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
				return ErrInvalidLengthNftAttributeValue
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNftAttributeValue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &BooleanAttributeValue{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &NftAttributeValue_BooleanAttributeValue{v}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FloatAttributeValue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftAttributeValue
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
				return ErrInvalidLengthNftAttributeValue
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNftAttributeValue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &FloatAttributeValue{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &NftAttributeValue_FloatAttributeValue{v}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HiddenToMarketplace", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftAttributeValue
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
			m.HiddenToMarketplace = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipNftAttributeValue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNftAttributeValue
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
func (m *NumberAttributeValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNftAttributeValue
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
			return fmt.Errorf("proto: NumberAttributeValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NumberAttributeValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftAttributeValue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNftAttributeValue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNftAttributeValue
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
func (m *StringAttributeValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNftAttributeValue
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
			return fmt.Errorf("proto: StringAttributeValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StringAttributeValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftAttributeValue
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
				return ErrInvalidLengthNftAttributeValue
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftAttributeValue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNftAttributeValue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNftAttributeValue
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
func (m *BooleanAttributeValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNftAttributeValue
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
			return fmt.Errorf("proto: BooleanAttributeValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BooleanAttributeValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftAttributeValue
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
			m.Value = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipNftAttributeValue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNftAttributeValue
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
func (m *FloatAttributeValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNftAttributeValue
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
			return fmt.Errorf("proto: FloatAttributeValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FloatAttributeValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Value = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipNftAttributeValue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNftAttributeValue
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
func skipNftAttributeValue(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNftAttributeValue
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
					return 0, ErrIntOverflowNftAttributeValue
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
					return 0, ErrIntOverflowNftAttributeValue
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
				return 0, ErrInvalidLengthNftAttributeValue
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNftAttributeValue
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNftAttributeValue
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNftAttributeValue        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNftAttributeValue          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNftAttributeValue = fmt.Errorf("proto: unexpected end of group")
)
