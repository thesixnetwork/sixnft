// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftmngr/v090/attribute_definition.proto

package v090

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

type DefaultMintValue struct {
	// Types that are valid to be assigned to Value:
	//	*DefaultMintValue_NumberAttributeValue
	//	*DefaultMintValue_StringAttributeValue
	//	*DefaultMintValue_BooleanAttributeValue
	//	*DefaultMintValue_FloatAttributeValue
	Value isDefaultMintValue_Value `protobuf_oneof:"value"`
}

func (m *DefaultMintValue) Reset()         { *m = DefaultMintValue{} }
func (m *DefaultMintValue) String() string { return proto.CompactTextString(m) }
func (*DefaultMintValue) ProtoMessage()    {}
func (*DefaultMintValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f86010096f4edfe, []int{0}
}
func (m *DefaultMintValue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DefaultMintValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DefaultMintValue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DefaultMintValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DefaultMintValue.Merge(m, src)
}
func (m *DefaultMintValue) XXX_Size() int {
	return m.Size()
}
func (m *DefaultMintValue) XXX_DiscardUnknown() {
	xxx_messageInfo_DefaultMintValue.DiscardUnknown(m)
}

var xxx_messageInfo_DefaultMintValue proto.InternalMessageInfo

type isDefaultMintValue_Value interface {
	isDefaultMintValue_Value()
	MarshalTo([]byte) (int, error)
	Size() int
}

type DefaultMintValue_NumberAttributeValue struct {
	NumberAttributeValue *NumberAttributeValue `protobuf:"bytes,1,opt,name=number_attribute_value,json=numberAttributeValue,proto3,oneof" json:"number_attribute_value,omitempty"`
}
type DefaultMintValue_StringAttributeValue struct {
	StringAttributeValue *StringAttributeValue `protobuf:"bytes,2,opt,name=string_attribute_value,json=stringAttributeValue,proto3,oneof" json:"string_attribute_value,omitempty"`
}
type DefaultMintValue_BooleanAttributeValue struct {
	BooleanAttributeValue *BooleanAttributeValue `protobuf:"bytes,3,opt,name=boolean_attribute_value,json=booleanAttributeValue,proto3,oneof" json:"boolean_attribute_value,omitempty"`
}
type DefaultMintValue_FloatAttributeValue struct {
	FloatAttributeValue *FloatAttributeValue `protobuf:"bytes,4,opt,name=float_attribute_value,json=floatAttributeValue,proto3,oneof" json:"float_attribute_value,omitempty"`
}

func (*DefaultMintValue_NumberAttributeValue) isDefaultMintValue_Value()  {}
func (*DefaultMintValue_StringAttributeValue) isDefaultMintValue_Value()  {}
func (*DefaultMintValue_BooleanAttributeValue) isDefaultMintValue_Value() {}
func (*DefaultMintValue_FloatAttributeValue) isDefaultMintValue_Value()   {}

func (m *DefaultMintValue) GetValue() isDefaultMintValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *DefaultMintValue) GetNumberAttributeValue() *NumberAttributeValue {
	if x, ok := m.GetValue().(*DefaultMintValue_NumberAttributeValue); ok {
		return x.NumberAttributeValue
	}
	return nil
}

func (m *DefaultMintValue) GetStringAttributeValue() *StringAttributeValue {
	if x, ok := m.GetValue().(*DefaultMintValue_StringAttributeValue); ok {
		return x.StringAttributeValue
	}
	return nil
}

func (m *DefaultMintValue) GetBooleanAttributeValue() *BooleanAttributeValue {
	if x, ok := m.GetValue().(*DefaultMintValue_BooleanAttributeValue); ok {
		return x.BooleanAttributeValue
	}
	return nil
}

func (m *DefaultMintValue) GetFloatAttributeValue() *FloatAttributeValue {
	if x, ok := m.GetValue().(*DefaultMintValue_FloatAttributeValue); ok {
		return x.FloatAttributeValue
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DefaultMintValue) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DefaultMintValue_NumberAttributeValue)(nil),
		(*DefaultMintValue_StringAttributeValue)(nil),
		(*DefaultMintValue_BooleanAttributeValue)(nil),
		(*DefaultMintValue_FloatAttributeValue)(nil),
	}
}

type AttributeDefinition struct {
	Name                string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DataType            string            `protobuf:"bytes,2,opt,name=data_type,json=dataType,proto3" json:"data_type,omitempty"`
	Required            bool              `protobuf:"varint,3,opt,name=required,proto3" json:"required,omitempty"`
	DisplayValueField   string            `protobuf:"bytes,4,opt,name=display_value_field,json=displayValueField,proto3" json:"display_value_field,omitempty"`
	DisplayOption       *DisplayOption    `protobuf:"bytes,5,opt,name=display_option,json=displayOption,proto3" json:"display_option,omitempty"`
	DefaultMintValue    *DefaultMintValue `protobuf:"bytes,6,opt,name=default_mint_value,json=defaultMintValue,proto3" json:"default_mint_value,omitempty"`
	HiddenOveride       bool              `protobuf:"varint,7,opt,name=hidden_overide,json=hiddenOveride,proto3" json:"hidden_overide,omitempty"`
	HiddenToMarketplace bool              `protobuf:"varint,8,opt,name=hidden_to_marketplace,json=hiddenToMarketplace,proto3" json:"hidden_to_marketplace,omitempty"`
	Index               uint64            `protobuf:"varint,9,opt,name=index,proto3" json:"index,omitempty"`
}

func (m *AttributeDefinition) Reset()         { *m = AttributeDefinition{} }
func (m *AttributeDefinition) String() string { return proto.CompactTextString(m) }
func (*AttributeDefinition) ProtoMessage()    {}
func (*AttributeDefinition) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f86010096f4edfe, []int{1}
}
func (m *AttributeDefinition) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AttributeDefinition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AttributeDefinition.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AttributeDefinition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttributeDefinition.Merge(m, src)
}
func (m *AttributeDefinition) XXX_Size() int {
	return m.Size()
}
func (m *AttributeDefinition) XXX_DiscardUnknown() {
	xxx_messageInfo_AttributeDefinition.DiscardUnknown(m)
}

var xxx_messageInfo_AttributeDefinition proto.InternalMessageInfo

func (m *AttributeDefinition) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AttributeDefinition) GetDataType() string {
	if m != nil {
		return m.DataType
	}
	return ""
}

func (m *AttributeDefinition) GetRequired() bool {
	if m != nil {
		return m.Required
	}
	return false
}

func (m *AttributeDefinition) GetDisplayValueField() string {
	if m != nil {
		return m.DisplayValueField
	}
	return ""
}

func (m *AttributeDefinition) GetDisplayOption() *DisplayOption {
	if m != nil {
		return m.DisplayOption
	}
	return nil
}

func (m *AttributeDefinition) GetDefaultMintValue() *DefaultMintValue {
	if m != nil {
		return m.DefaultMintValue
	}
	return nil
}

func (m *AttributeDefinition) GetHiddenOveride() bool {
	if m != nil {
		return m.HiddenOveride
	}
	return false
}

func (m *AttributeDefinition) GetHiddenToMarketplace() bool {
	if m != nil {
		return m.HiddenToMarketplace
	}
	return false
}

func (m *AttributeDefinition) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func init() {
	proto.RegisterType((*DefaultMintValue)(nil), "thesixnetwork.sixnft.nftmngr.v090.DefaultMintValue")
	proto.RegisterType((*AttributeDefinition)(nil), "thesixnetwork.sixnft.nftmngr.v090.AttributeDefinition")
}

func init() {
	proto.RegisterFile("nftmngr/v090/attribute_definition.proto", fileDescriptor_4f86010096f4edfe)
}

var fileDescriptor_4f86010096f4edfe = []byte{
	// 524 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0x6e, 0x68, 0xbb, 0xb5, 0x46, 0x9b, 0x86, 0xbb, 0x42, 0x54, 0xa4, 0x68, 0x9b, 0x04, 0xec,
	0x94, 0x54, 0x1b, 0xe2, 0xc7, 0x91, 0xaa, 0x9a, 0xb8, 0x6c, 0x93, 0xc2, 0x04, 0x12, 0x97, 0xc8,
	0x99, 0x9d, 0xd6, 0x5a, 0x62, 0x07, 0xf7, 0xa5, 0xb4, 0xff, 0x05, 0xff, 0x11, 0x57, 0x8e, 0x3b,
	0x72, 0x03, 0xb5, 0xff, 0x08, 0x8a, 0xdd, 0x1f, 0xb4, 0xa9, 0xb4, 0xde, 0xec, 0xf7, 0xbd, 0xef,
	0xfb, 0xec, 0xe7, 0xf7, 0x8c, 0x5e, 0x89, 0x08, 0x12, 0xd1, 0x53, 0xde, 0xb0, 0xfd, 0xbe, 0xed,
	0x11, 0x00, 0xc5, 0xc3, 0x0c, 0x58, 0x40, 0x59, 0xc4, 0x05, 0x07, 0x2e, 0x85, 0x9b, 0x2a, 0x09,
	0x12, 0x1f, 0x43, 0x9f, 0x0d, 0xf8, 0x48, 0x30, 0xf8, 0x2e, 0xd5, 0x9d, 0x9b, 0x2f, 0x23, 0x70,
	0x67, 0x6c, 0x37, 0x67, 0xb7, 0x8e, 0x57, 0xb4, 0x28, 0x1f, 0xa4, 0x31, 0x19, 0x07, 0x32, 0x5d,
	0xaa, 0xb4, 0x5e, 0xae, 0xa4, 0x88, 0x08, 0x82, 0xa5, 0xe5, 0x90, 0xc4, 0x19, 0x33, 0x79, 0x27,
	0x7f, 0xca, 0xe8, 0xa0, 0xcb, 0x22, 0x92, 0xc5, 0x70, 0xc9, 0x05, 0x7c, 0xce, 0x21, 0x2c, 0xd1,
	0x53, 0x91, 0x25, 0x21, 0x53, 0xeb, 0x24, 0xdb, 0x3a, 0xb2, 0x4e, 0x1f, 0x9f, 0xbd, 0x75, 0x1f,
	0x3c, 0xa3, 0x7b, 0xa5, 0x05, 0x3e, 0xcc, 0xf9, 0x5a, 0xf8, 0x63, 0xc9, 0x3f, 0x14, 0x1b, 0xe2,
	0xb9, 0xe1, 0x00, 0x14, 0x17, 0xbd, 0x82, 0xe1, 0xa3, 0xad, 0x0d, 0x3f, 0x69, 0x81, 0xa2, 0xe1,
	0x60, 0x43, 0x1c, 0x2b, 0xf4, 0x2c, 0x94, 0x32, 0x66, 0x44, 0x14, 0x1c, 0xcb, 0xda, 0xf1, 0xdd,
	0x16, 0x8e, 0x1d, 0xa3, 0x50, 0xb0, 0x6c, 0x86, 0x9b, 0x00, 0x1c, 0xa3, 0x66, 0x14, 0x4b, 0x52,
	0x78, 0x09, 0xbb, 0xa2, 0x1d, 0xdf, 0x6c, 0xe1, 0x78, 0x91, 0xf3, 0x0b, 0x7e, 0x8d, 0xa8, 0x18,
	0xee, 0xec, 0xa2, 0xaa, 0x56, 0x3f, 0xf9, 0x59, 0x46, 0x8d, 0x05, 0xd6, 0x5d, 0x74, 0x1b, 0xc6,
	0xa8, 0x22, 0x48, 0x62, 0x9e, 0xb4, 0xee, 0xeb, 0x35, 0x7e, 0x8e, 0xea, 0x94, 0x00, 0x09, 0x60,
	0x9c, 0x9a, 0xd2, 0xd7, 0xfd, 0x5a, 0x1e, 0xb8, 0x19, 0xa7, 0x0c, 0xb7, 0x50, 0x4d, 0xb1, 0x6f,
	0x19, 0x57, 0x8c, 0xea, 0x22, 0xd5, 0xfc, 0xc5, 0x1e, 0xbb, 0xa8, 0x31, 0x6f, 0x43, 0xed, 0x1a,
	0x44, 0x9c, 0xc5, 0x54, 0xdf, 0xac, 0xee, 0x3f, 0x99, 0x41, 0xfa, 0x60, 0x17, 0x39, 0x80, 0xbf,
	0xa0, 0xfd, 0xd5, 0xb6, 0xb5, 0xab, 0xba, 0x08, 0xed, 0x2d, 0x8a, 0xd0, 0x35, 0xc4, 0x6b, 0xcd,
	0xf3, 0xf7, 0xe8, 0xff, 0x5b, 0x4c, 0x10, 0xa6, 0xa6, 0x9d, 0x83, 0x84, 0x0b, 0x98, 0x55, 0x78,
	0x47, 0x8b, 0x9f, 0x6f, 0x23, 0xbe, 0x36, 0x0b, 0xfe, 0x01, 0x5d, 0x9f, 0x8e, 0x17, 0x68, 0xbf,
	0xcf, 0x29, 0x65, 0x22, 0x90, 0x43, 0xa6, 0x38, 0x65, 0xf6, 0xae, 0xae, 0xc6, 0x9e, 0x89, 0x5e,
	0x9b, 0x20, 0x3e, 0x43, 0xcd, 0x59, 0x1a, 0xc8, 0x20, 0x21, 0xea, 0x8e, 0x41, 0x1a, 0x93, 0x5b,
	0x66, 0xd7, 0x74, 0x76, 0xc3, 0x80, 0x37, 0xf2, 0x72, 0x09, 0xe1, 0x43, 0x54, 0xe5, 0x82, 0xb2,
	0x91, 0x5d, 0x3f, 0xb2, 0x4e, 0x2b, 0xbe, 0xd9, 0x74, 0xae, 0x7e, 0x4d, 0x1c, 0xeb, 0x7e, 0xe2,
	0x58, 0x7f, 0x27, 0x8e, 0xf5, 0x63, 0xea, 0x94, 0xee, 0xa7, 0x4e, 0xe9, 0xf7, 0xd4, 0x29, 0x7d,
	0x7d, 0xdd, 0xe3, 0xd0, 0xcf, 0x42, 0xf7, 0x56, 0x26, 0xde, 0xca, 0xdd, 0x3c, 0x73, 0x37, 0x6f,
	0xe4, 0xcd, 0xff, 0x81, 0xfc, 0x49, 0x07, 0xfa, 0x37, 0x08, 0x77, 0xf4, 0xe8, 0x9f, 0xff, 0x0b,
	0x00, 0x00, 0xff, 0xff, 0x58, 0x72, 0xc7, 0x58, 0x93, 0x04, 0x00, 0x00,
}

func (m *DefaultMintValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DefaultMintValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DefaultMintValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != nil {
		{
			size := m.Value.Size()
			i -= size
			if _, err := m.Value.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *DefaultMintValue_NumberAttributeValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DefaultMintValue_NumberAttributeValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.NumberAttributeValue != nil {
		{
			size, err := m.NumberAttributeValue.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAttributeDefinition(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *DefaultMintValue_StringAttributeValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DefaultMintValue_StringAttributeValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.StringAttributeValue != nil {
		{
			size, err := m.StringAttributeValue.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAttributeDefinition(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *DefaultMintValue_BooleanAttributeValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DefaultMintValue_BooleanAttributeValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.BooleanAttributeValue != nil {
		{
			size, err := m.BooleanAttributeValue.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAttributeDefinition(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *DefaultMintValue_FloatAttributeValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DefaultMintValue_FloatAttributeValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.FloatAttributeValue != nil {
		{
			size, err := m.FloatAttributeValue.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAttributeDefinition(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *AttributeDefinition) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AttributeDefinition) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AttributeDefinition) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Index != 0 {
		i = encodeVarintAttributeDefinition(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x48
	}
	if m.HiddenToMarketplace {
		i--
		if m.HiddenToMarketplace {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if m.HiddenOveride {
		i--
		if m.HiddenOveride {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.DefaultMintValue != nil {
		{
			size, err := m.DefaultMintValue.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAttributeDefinition(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.DisplayOption != nil {
		{
			size, err := m.DisplayOption.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAttributeDefinition(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.DisplayValueField) > 0 {
		i -= len(m.DisplayValueField)
		copy(dAtA[i:], m.DisplayValueField)
		i = encodeVarintAttributeDefinition(dAtA, i, uint64(len(m.DisplayValueField)))
		i--
		dAtA[i] = 0x22
	}
	if m.Required {
		i--
		if m.Required {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.DataType) > 0 {
		i -= len(m.DataType)
		copy(dAtA[i:], m.DataType)
		i = encodeVarintAttributeDefinition(dAtA, i, uint64(len(m.DataType)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintAttributeDefinition(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAttributeDefinition(dAtA []byte, offset int, v uint64) int {
	offset -= sovAttributeDefinition(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DefaultMintValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != nil {
		n += m.Value.Size()
	}
	return n
}

func (m *DefaultMintValue_NumberAttributeValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NumberAttributeValue != nil {
		l = m.NumberAttributeValue.Size()
		n += 1 + l + sovAttributeDefinition(uint64(l))
	}
	return n
}
func (m *DefaultMintValue_StringAttributeValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StringAttributeValue != nil {
		l = m.StringAttributeValue.Size()
		n += 1 + l + sovAttributeDefinition(uint64(l))
	}
	return n
}
func (m *DefaultMintValue_BooleanAttributeValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BooleanAttributeValue != nil {
		l = m.BooleanAttributeValue.Size()
		n += 1 + l + sovAttributeDefinition(uint64(l))
	}
	return n
}
func (m *DefaultMintValue_FloatAttributeValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FloatAttributeValue != nil {
		l = m.FloatAttributeValue.Size()
		n += 1 + l + sovAttributeDefinition(uint64(l))
	}
	return n
}
func (m *AttributeDefinition) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovAttributeDefinition(uint64(l))
	}
	l = len(m.DataType)
	if l > 0 {
		n += 1 + l + sovAttributeDefinition(uint64(l))
	}
	if m.Required {
		n += 2
	}
	l = len(m.DisplayValueField)
	if l > 0 {
		n += 1 + l + sovAttributeDefinition(uint64(l))
	}
	if m.DisplayOption != nil {
		l = m.DisplayOption.Size()
		n += 1 + l + sovAttributeDefinition(uint64(l))
	}
	if m.DefaultMintValue != nil {
		l = m.DefaultMintValue.Size()
		n += 1 + l + sovAttributeDefinition(uint64(l))
	}
	if m.HiddenOveride {
		n += 2
	}
	if m.HiddenToMarketplace {
		n += 2
	}
	if m.Index != 0 {
		n += 1 + sovAttributeDefinition(uint64(m.Index))
	}
	return n
}

func sovAttributeDefinition(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAttributeDefinition(x uint64) (n int) {
	return sovAttributeDefinition(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DefaultMintValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAttributeDefinition
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
			return fmt.Errorf("proto: DefaultMintValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DefaultMintValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumberAttributeValue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttributeDefinition
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
				return ErrInvalidLengthAttributeDefinition
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAttributeDefinition
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &NumberAttributeValue{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &DefaultMintValue_NumberAttributeValue{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StringAttributeValue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttributeDefinition
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
				return ErrInvalidLengthAttributeDefinition
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAttributeDefinition
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &StringAttributeValue{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &DefaultMintValue_StringAttributeValue{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BooleanAttributeValue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttributeDefinition
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
				return ErrInvalidLengthAttributeDefinition
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAttributeDefinition
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &BooleanAttributeValue{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &DefaultMintValue_BooleanAttributeValue{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FloatAttributeValue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttributeDefinition
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
				return ErrInvalidLengthAttributeDefinition
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAttributeDefinition
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &FloatAttributeValue{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &DefaultMintValue_FloatAttributeValue{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAttributeDefinition(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAttributeDefinition
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
func (m *AttributeDefinition) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAttributeDefinition
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
			return fmt.Errorf("proto: AttributeDefinition: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AttributeDefinition: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttributeDefinition
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
				return ErrInvalidLengthAttributeDefinition
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAttributeDefinition
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttributeDefinition
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
				return ErrInvalidLengthAttributeDefinition
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAttributeDefinition
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Required", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttributeDefinition
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
			m.Required = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisplayValueField", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttributeDefinition
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
				return ErrInvalidLengthAttributeDefinition
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAttributeDefinition
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DisplayValueField = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisplayOption", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttributeDefinition
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
				return ErrInvalidLengthAttributeDefinition
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAttributeDefinition
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DisplayOption == nil {
				m.DisplayOption = &DisplayOption{}
			}
			if err := m.DisplayOption.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultMintValue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttributeDefinition
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
				return ErrInvalidLengthAttributeDefinition
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAttributeDefinition
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DefaultMintValue == nil {
				m.DefaultMintValue = &DefaultMintValue{}
			}
			if err := m.DefaultMintValue.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HiddenOveride", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttributeDefinition
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
			m.HiddenOveride = bool(v != 0)
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HiddenToMarketplace", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttributeDefinition
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
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttributeDefinition
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAttributeDefinition(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAttributeDefinition
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
func skipAttributeDefinition(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAttributeDefinition
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
					return 0, ErrIntOverflowAttributeDefinition
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
					return 0, ErrIntOverflowAttributeDefinition
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
				return 0, ErrInvalidLengthAttributeDefinition
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAttributeDefinition
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAttributeDefinition
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAttributeDefinition        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAttributeDefinition          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAttributeDefinition = fmt.Errorf("proto: unexpected end of group")
)
