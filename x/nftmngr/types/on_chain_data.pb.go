// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftmngr/on_chain_data.proto

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

type FlagStatus struct {
	StatusName  string `protobuf:"bytes,1,opt,name=status_name,json=statusName,proto3" json:"status_name,omitempty"`
	StatusValue bool   `protobuf:"varint,2,opt,name=status_value,json=statusValue,proto3" json:"status_value,omitempty"`
}

func (m *FlagStatus) Reset()         { *m = FlagStatus{} }
func (m *FlagStatus) String() string { return proto.CompactTextString(m) }
func (*FlagStatus) ProtoMessage()    {}
func (*FlagStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_35d167410338c830, []int{0}
}
func (m *FlagStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FlagStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FlagStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FlagStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlagStatus.Merge(m, src)
}
func (m *FlagStatus) XXX_Size() int {
	return m.Size()
}
func (m *FlagStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_FlagStatus.DiscardUnknown(m)
}

var xxx_messageInfo_FlagStatus proto.InternalMessageInfo

func (m *FlagStatus) GetStatusName() string {
	if m != nil {
		return m.StatusName
	}
	return ""
}

func (m *FlagStatus) GetStatusValue() bool {
	if m != nil {
		return m.StatusValue
	}
	return false
}

type OnChainData struct {
	RevealRequired   bool                   `protobuf:"varint,1,opt,name=reveal_required,json=revealRequired,proto3" json:"reveal_required,omitempty"`
	RevealSecret     []byte                 `protobuf:"bytes,2,opt,name=reveal_secret,json=revealSecret,proto3" json:"reveal_secret,omitempty"`
	SchemaAttributes []*AttributeDefinition `protobuf:"bytes,3,rep,name=schema_attributes,json=schemaAttributes,proto3" json:"schema_attributes,omitempty"`
	TokenAttributes  []*AttributeDefinition `protobuf:"bytes,4,rep,name=token_attributes,json=tokenAttributes,proto3" json:"token_attributes,omitempty"`
	Actions          []*Action              `protobuf:"bytes,5,rep,name=actions,proto3" json:"actions,omitempty"`
	Status           []*FlagStatus          `protobuf:"bytes,6,rep,name=status,proto3" json:"status,omitempty"`
}

func (m *OnChainData) Reset()         { *m = OnChainData{} }
func (m *OnChainData) String() string { return proto.CompactTextString(m) }
func (*OnChainData) ProtoMessage()    {}
func (*OnChainData) Descriptor() ([]byte, []int) {
	return fileDescriptor_35d167410338c830, []int{1}
}
func (m *OnChainData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OnChainData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OnChainData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OnChainData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OnChainData.Merge(m, src)
}
func (m *OnChainData) XXX_Size() int {
	return m.Size()
}
func (m *OnChainData) XXX_DiscardUnknown() {
	xxx_messageInfo_OnChainData.DiscardUnknown(m)
}

var xxx_messageInfo_OnChainData proto.InternalMessageInfo

func (m *OnChainData) GetRevealRequired() bool {
	if m != nil {
		return m.RevealRequired
	}
	return false
}

func (m *OnChainData) GetRevealSecret() []byte {
	if m != nil {
		return m.RevealSecret
	}
	return nil
}

func (m *OnChainData) GetSchemaAttributes() []*AttributeDefinition {
	if m != nil {
		return m.SchemaAttributes
	}
	return nil
}

func (m *OnChainData) GetTokenAttributes() []*AttributeDefinition {
	if m != nil {
		return m.TokenAttributes
	}
	return nil
}

func (m *OnChainData) GetActions() []*Action {
	if m != nil {
		return m.Actions
	}
	return nil
}

func (m *OnChainData) GetStatus() []*FlagStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type OnChainDataV1 struct {
	RevealRequired     bool                   `protobuf:"varint,1,opt,name=reveal_required,json=revealRequired,proto3" json:"reveal_required,omitempty"`
	RevealSecret       []byte                 `protobuf:"bytes,2,opt,name=reveal_secret,json=revealSecret,proto3" json:"reveal_secret,omitempty"`
	NftAttributes      []*AttributeDefinition `protobuf:"bytes,3,rep,name=nft_attributes,json=nftAttributes,proto3" json:"nft_attributes,omitempty"`
	TokenAttributes    []*AttributeDefinition `protobuf:"bytes,4,rep,name=token_attributes,json=tokenAttributes,proto3" json:"token_attributes,omitempty"`
	Actions            []*Action              `protobuf:"bytes,5,rep,name=actions,proto3" json:"actions,omitempty"`
	Status             []*FlagStatus          `protobuf:"bytes,6,rep,name=status,proto3" json:"status,omitempty"`
	NftAttributesValue []*NftAttributeValue   `protobuf:"bytes,8,rep,name=nft_attributes_value,json=nftAttributesValue,proto3" json:"nft_attributes_value,omitempty"`
}

func (m *OnChainDataV1) Reset()         { *m = OnChainDataV1{} }
func (m *OnChainDataV1) String() string { return proto.CompactTextString(m) }
func (*OnChainDataV1) ProtoMessage()    {}
func (*OnChainDataV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_35d167410338c830, []int{2}
}
func (m *OnChainDataV1) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OnChainDataV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OnChainDataV1.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OnChainDataV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OnChainDataV1.Merge(m, src)
}
func (m *OnChainDataV1) XXX_Size() int {
	return m.Size()
}
func (m *OnChainDataV1) XXX_DiscardUnknown() {
	xxx_messageInfo_OnChainDataV1.DiscardUnknown(m)
}

var xxx_messageInfo_OnChainDataV1 proto.InternalMessageInfo

func (m *OnChainDataV1) GetRevealRequired() bool {
	if m != nil {
		return m.RevealRequired
	}
	return false
}

func (m *OnChainDataV1) GetRevealSecret() []byte {
	if m != nil {
		return m.RevealSecret
	}
	return nil
}

func (m *OnChainDataV1) GetNftAttributes() []*AttributeDefinition {
	if m != nil {
		return m.NftAttributes
	}
	return nil
}

func (m *OnChainDataV1) GetTokenAttributes() []*AttributeDefinition {
	if m != nil {
		return m.TokenAttributes
	}
	return nil
}

func (m *OnChainDataV1) GetActions() []*Action {
	if m != nil {
		return m.Actions
	}
	return nil
}

func (m *OnChainDataV1) GetStatus() []*FlagStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *OnChainDataV1) GetNftAttributesValue() []*NftAttributeValue {
	if m != nil {
		return m.NftAttributesValue
	}
	return nil
}

func init() {
	proto.RegisterType((*FlagStatus)(nil), "thesixnetwork.sixnft.nftmngr.FlagStatus")
	proto.RegisterType((*OnChainData)(nil), "thesixnetwork.sixnft.nftmngr.OnChainData")
	proto.RegisterType((*OnChainDataV1)(nil), "thesixnetwork.sixnft.nftmngr.OnChainDataV1")
}

func init() { proto.RegisterFile("nftmngr/on_chain_data.proto", fileDescriptor_35d167410338c830) }

var fileDescriptor_35d167410338c830 = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x94, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x1b, 0x0a, 0xa5, 0xbc, 0x6d, 0xb7, 0x61, 0xed, 0x10, 0x0d, 0x14, 0xba, 0x82, 0x44,
	0x4f, 0x89, 0x06, 0x77, 0xc4, 0x9f, 0x09, 0x89, 0xcb, 0x40, 0x9e, 0x34, 0x21, 0x84, 0x88, 0xdc,
	0xf4, 0x4d, 0x1b, 0xad, 0x71, 0x86, 0xfd, 0x66, 0x8c, 0x6f, 0xc1, 0x9d, 0x2f, 0xc4, 0xb1, 0x47,
	0x8e, 0xa8, 0xfd, 0x16, 0x9c, 0x50, 0xed, 0xa4, 0x4d, 0x2e, 0x45, 0x02, 0x6e, 0xdc, 0xac, 0x27,
	0xef, 0xf3, 0xb3, 0xfc, 0x3c, 0xb1, 0xe1, 0x8e, 0x8c, 0x29, 0x95, 0x13, 0x15, 0x64, 0x32, 0x8c,
	0xa6, 0x22, 0x91, 0xe1, 0x58, 0x90, 0xf0, 0x2f, 0x54, 0x46, 0x19, 0xbb, 0x4b, 0x53, 0xd4, 0xc9,
	0x95, 0x44, 0xfa, 0x94, 0xa9, 0x73, 0x7f, 0xb5, 0x8c, 0xc9, 0x2f, 0x1c, 0x07, 0x83, 0xd2, 0x2a,
	0x88, 0x54, 0x32, 0xca, 0x09, 0xc3, 0x31, 0xc6, 0x89, 0x4c, 0x28, 0xc9, 0xa4, 0x25, 0x1c, 0xec,
	0xaf, 0x67, 0xa2, 0x8a, 0x7a, 0x58, 0xaa, 0x32, 0xa6, 0x70, 0xe3, 0xbe, 0x14, 0xb3, 0x1c, 0xed,
	0xc8, 0xe0, 0x0d, 0xc0, 0xcb, 0x99, 0x98, 0x9c, 0x92, 0xa0, 0x5c, 0xb3, 0x7b, 0xd0, 0xd1, 0x66,
	0x15, 0x4a, 0x91, 0xa2, 0xeb, 0xf4, 0x9d, 0xe1, 0x2d, 0x0e, 0x56, 0x3a, 0x11, 0x29, 0xb2, 0x43,
	0xe8, 0x16, 0x03, 0x06, 0xe2, 0x5e, 0xeb, 0x3b, 0xc3, 0x36, 0x2f, 0x4c, 0x67, 0x2b, 0x69, 0xf0,
	0xb5, 0x09, 0x9d, 0xd7, 0xf2, 0xc5, 0xea, 0x8c, 0xc7, 0x82, 0x04, 0x7b, 0x08, 0xbb, 0x0a, 0x2f,
	0x51, 0xcc, 0x42, 0x85, 0x1f, 0xf3, 0x44, 0xe1, 0xd8, 0x70, 0xdb, 0x7c, 0xc7, 0xca, 0xbc, 0x50,
	0xd9, 0x7d, 0xe8, 0x15, 0x83, 0x1a, 0x23, 0x85, 0x64, 0xe0, 0x5d, 0xde, 0xb5, 0xe2, 0xa9, 0xd1,
	0xd8, 0x07, 0xb8, 0xad, 0xa3, 0x29, 0xa6, 0x62, 0x73, 0x1e, 0xed, 0x36, 0xfb, 0xcd, 0x61, 0xe7,
	0xd1, 0x91, 0xbf, 0x2d, 0x46, 0xff, 0x59, 0x39, 0x7f, 0xbc, 0x0e, 0x8f, 0xef, 0x59, 0xd6, 0xfa,
	0x93, 0x66, 0xef, 0x61, 0x8f, 0xb2, 0x73, 0x94, 0x55, 0xfc, 0xf5, 0x3f, 0xc5, 0xef, 0x1a, 0x54,
	0x85, 0xfe, 0x04, 0x6e, 0xda, 0x82, 0xb4, 0x7b, 0xc3, 0x40, 0x1f, 0xfc, 0x06, 0x6a, 0x86, 0x79,
	0x69, 0x62, 0x4f, 0xa1, 0x65, 0xa3, 0x76, 0x5b, 0xc6, 0x3e, 0xdc, 0x6e, 0xdf, 0x34, 0xcb, 0x0b,
	0xdf, 0xe0, 0x67, 0x13, 0x7a, 0x95, 0x76, 0xce, 0x8e, 0xfe, 0x71, 0x3f, 0x6f, 0x61, 0xa7, 0xf6,
	0xb3, 0xfd, 0x45, 0x39, 0x3d, 0x19, 0xd3, 0xff, 0xd3, 0x0c, 0x13, 0xb0, 0x5f, 0x4f, 0xae, 0xb8,
	0x62, 0x6d, 0xc3, 0x0b, 0xb6, 0xf3, 0x4e, 0x2a, 0x51, 0x99, 0x6b, 0xc8, 0x59, 0x2d, 0x3d, 0xa3,
	0x3d, 0x7f, 0xf5, 0x6d, 0xe1, 0x39, 0xf3, 0x85, 0xe7, 0xfc, 0x58, 0x78, 0xce, 0x97, 0xa5, 0xd7,
	0x98, 0x2f, 0xbd, 0xc6, 0xf7, 0xa5, 0xd7, 0x78, 0x17, 0x4c, 0x12, 0x9a, 0xe6, 0x23, 0x3f, 0xca,
	0xd2, 0xa0, 0xb6, 0x51, 0x60, 0x37, 0x0a, 0xae, 0x82, 0xf2, 0x2d, 0xa1, 0xcf, 0x17, 0xa8, 0x47,
	0x2d, 0xf3, 0x7c, 0x3c, 0xfe, 0x15, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x0d, 0x17, 0x17, 0xd8, 0x04,
	0x00, 0x00,
}

func (m *FlagStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FlagStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FlagStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.StatusValue {
		i--
		if m.StatusValue {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.StatusName) > 0 {
		i -= len(m.StatusName)
		copy(dAtA[i:], m.StatusName)
		i = encodeVarintOnChainData(dAtA, i, uint64(len(m.StatusName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OnChainData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OnChainData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OnChainData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Status) > 0 {
		for iNdEx := len(m.Status) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Status[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOnChainData(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Actions) > 0 {
		for iNdEx := len(m.Actions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Actions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOnChainData(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.TokenAttributes) > 0 {
		for iNdEx := len(m.TokenAttributes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TokenAttributes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOnChainData(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.SchemaAttributes) > 0 {
		for iNdEx := len(m.SchemaAttributes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SchemaAttributes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOnChainData(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.RevealSecret) > 0 {
		i -= len(m.RevealSecret)
		copy(dAtA[i:], m.RevealSecret)
		i = encodeVarintOnChainData(dAtA, i, uint64(len(m.RevealSecret)))
		i--
		dAtA[i] = 0x12
	}
	if m.RevealRequired {
		i--
		if m.RevealRequired {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *OnChainDataV1) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OnChainDataV1) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OnChainDataV1) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NftAttributesValue) > 0 {
		for iNdEx := len(m.NftAttributesValue) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NftAttributesValue[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOnChainData(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.Status) > 0 {
		for iNdEx := len(m.Status) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Status[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOnChainData(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Actions) > 0 {
		for iNdEx := len(m.Actions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Actions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOnChainData(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.TokenAttributes) > 0 {
		for iNdEx := len(m.TokenAttributes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TokenAttributes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOnChainData(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.NftAttributes) > 0 {
		for iNdEx := len(m.NftAttributes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NftAttributes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOnChainData(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.RevealSecret) > 0 {
		i -= len(m.RevealSecret)
		copy(dAtA[i:], m.RevealSecret)
		i = encodeVarintOnChainData(dAtA, i, uint64(len(m.RevealSecret)))
		i--
		dAtA[i] = 0x12
	}
	if m.RevealRequired {
		i--
		if m.RevealRequired {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintOnChainData(dAtA []byte, offset int, v uint64) int {
	offset -= sovOnChainData(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FlagStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StatusName)
	if l > 0 {
		n += 1 + l + sovOnChainData(uint64(l))
	}
	if m.StatusValue {
		n += 2
	}
	return n
}

func (m *OnChainData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RevealRequired {
		n += 2
	}
	l = len(m.RevealSecret)
	if l > 0 {
		n += 1 + l + sovOnChainData(uint64(l))
	}
	if len(m.SchemaAttributes) > 0 {
		for _, e := range m.SchemaAttributes {
			l = e.Size()
			n += 1 + l + sovOnChainData(uint64(l))
		}
	}
	if len(m.TokenAttributes) > 0 {
		for _, e := range m.TokenAttributes {
			l = e.Size()
			n += 1 + l + sovOnChainData(uint64(l))
		}
	}
	if len(m.Actions) > 0 {
		for _, e := range m.Actions {
			l = e.Size()
			n += 1 + l + sovOnChainData(uint64(l))
		}
	}
	if len(m.Status) > 0 {
		for _, e := range m.Status {
			l = e.Size()
			n += 1 + l + sovOnChainData(uint64(l))
		}
	}
	return n
}

func (m *OnChainDataV1) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RevealRequired {
		n += 2
	}
	l = len(m.RevealSecret)
	if l > 0 {
		n += 1 + l + sovOnChainData(uint64(l))
	}
	if len(m.NftAttributes) > 0 {
		for _, e := range m.NftAttributes {
			l = e.Size()
			n += 1 + l + sovOnChainData(uint64(l))
		}
	}
	if len(m.TokenAttributes) > 0 {
		for _, e := range m.TokenAttributes {
			l = e.Size()
			n += 1 + l + sovOnChainData(uint64(l))
		}
	}
	if len(m.Actions) > 0 {
		for _, e := range m.Actions {
			l = e.Size()
			n += 1 + l + sovOnChainData(uint64(l))
		}
	}
	if len(m.Status) > 0 {
		for _, e := range m.Status {
			l = e.Size()
			n += 1 + l + sovOnChainData(uint64(l))
		}
	}
	if len(m.NftAttributesValue) > 0 {
		for _, e := range m.NftAttributesValue {
			l = e.Size()
			n += 1 + l + sovOnChainData(uint64(l))
		}
	}
	return n
}

func sovOnChainData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOnChainData(x uint64) (n int) {
	return sovOnChainData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FlagStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOnChainData
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
			return fmt.Errorf("proto: FlagStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FlagStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatusName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOnChainData
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
				return ErrInvalidLengthOnChainData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOnChainData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StatusName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatusValue", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOnChainData
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
			m.StatusValue = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipOnChainData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOnChainData
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
func (m *OnChainData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOnChainData
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
			return fmt.Errorf("proto: OnChainData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OnChainData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RevealRequired", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOnChainData
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
			m.RevealRequired = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RevealSecret", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOnChainData
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
				return ErrInvalidLengthOnChainData
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthOnChainData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RevealSecret = append(m.RevealSecret[:0], dAtA[iNdEx:postIndex]...)
			if m.RevealSecret == nil {
				m.RevealSecret = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SchemaAttributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOnChainData
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
				return ErrInvalidLengthOnChainData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOnChainData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SchemaAttributes = append(m.SchemaAttributes, &AttributeDefinition{})
			if err := m.SchemaAttributes[len(m.SchemaAttributes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenAttributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOnChainData
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
				return ErrInvalidLengthOnChainData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOnChainData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenAttributes = append(m.TokenAttributes, &AttributeDefinition{})
			if err := m.TokenAttributes[len(m.TokenAttributes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Actions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOnChainData
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
				return ErrInvalidLengthOnChainData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOnChainData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Actions = append(m.Actions, &Action{})
			if err := m.Actions[len(m.Actions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOnChainData
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
				return ErrInvalidLengthOnChainData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOnChainData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = append(m.Status, &FlagStatus{})
			if err := m.Status[len(m.Status)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOnChainData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOnChainData
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
func (m *OnChainDataV1) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOnChainData
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
			return fmt.Errorf("proto: OnChainDataV1: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OnChainDataV1: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RevealRequired", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOnChainData
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
			m.RevealRequired = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RevealSecret", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOnChainData
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
				return ErrInvalidLengthOnChainData
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthOnChainData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RevealSecret = append(m.RevealSecret[:0], dAtA[iNdEx:postIndex]...)
			if m.RevealSecret == nil {
				m.RevealSecret = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftAttributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOnChainData
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
				return ErrInvalidLengthOnChainData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOnChainData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NftAttributes = append(m.NftAttributes, &AttributeDefinition{})
			if err := m.NftAttributes[len(m.NftAttributes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenAttributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOnChainData
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
				return ErrInvalidLengthOnChainData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOnChainData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenAttributes = append(m.TokenAttributes, &AttributeDefinition{})
			if err := m.TokenAttributes[len(m.TokenAttributes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Actions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOnChainData
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
				return ErrInvalidLengthOnChainData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOnChainData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Actions = append(m.Actions, &Action{})
			if err := m.Actions[len(m.Actions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOnChainData
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
				return ErrInvalidLengthOnChainData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOnChainData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = append(m.Status, &FlagStatus{})
			if err := m.Status[len(m.Status)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftAttributesValue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOnChainData
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
				return ErrInvalidLengthOnChainData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOnChainData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NftAttributesValue = append(m.NftAttributesValue, &NftAttributeValue{})
			if err := m.NftAttributesValue[len(m.NftAttributesValue)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOnChainData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOnChainData
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
func skipOnChainData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOnChainData
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
					return 0, ErrIntOverflowOnChainData
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
					return 0, ErrIntOverflowOnChainData
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
				return 0, ErrInvalidLengthOnChainData
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOnChainData
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOnChainData
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOnChainData        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOnChainData          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOnChainData = fmt.Errorf("proto: unexpected end of group")
)
