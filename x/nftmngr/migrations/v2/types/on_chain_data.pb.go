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
	NftAttributes   []*AttributeDefinition `protobuf:"bytes,1,rep,name=nft_attributes,json=nftAttributes,proto3" json:"nft_attributes,omitempty"`
	TokenAttributes []*AttributeDefinition `protobuf:"bytes,2,rep,name=token_attributes,json=tokenAttributes,proto3" json:"token_attributes,omitempty"`
	Actions         []*Action              `protobuf:"bytes,3,rep,name=actions,proto3" json:"actions,omitempty"`
	Status          []*FlagStatus          `protobuf:"bytes,4,rep,name=status,proto3" json:"status,omitempty"`
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

func (m *OnChainData) GetNftAttributes() []*AttributeDefinition {
	if m != nil {
		return m.NftAttributes
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

type OnChainDataV2 struct {
	TokenAttributes []*AttributeDefinition `protobuf:"bytes,1,rep,name=token_attributes,json=tokenAttributes,proto3" json:"token_attributes,omitempty"`
	Actions         []*Action              `protobuf:"bytes,2,rep,name=actions,proto3" json:"actions,omitempty"`
	Status          []*FlagStatus          `protobuf:"bytes,3,rep,name=status,proto3" json:"status,omitempty"`
}

func (m *OnChainDataV2) Reset()         { *m = OnChainDataV2{} }
func (m *OnChainDataV2) String() string { return proto.CompactTextString(m) }
func (*OnChainDataV2) ProtoMessage()    {}
func (*OnChainDataV2) Descriptor() ([]byte, []int) {
	return fileDescriptor_35d167410338c830, []int{2}
}
func (m *OnChainDataV2) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OnChainDataV2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OnChainDataV2.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OnChainDataV2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OnChainDataV2.Merge(m, src)
}
func (m *OnChainDataV2) XXX_Size() int {
	return m.Size()
}
func (m *OnChainDataV2) XXX_DiscardUnknown() {
	xxx_messageInfo_OnChainDataV2.DiscardUnknown(m)
}

var xxx_messageInfo_OnChainDataV2 proto.InternalMessageInfo

func (m *OnChainDataV2) GetTokenAttributes() []*AttributeDefinition {
	if m != nil {
		return m.TokenAttributes
	}
	return nil
}

func (m *OnChainDataV2) GetActions() []*Action {
	if m != nil {
		return m.Actions
	}
	return nil
}

func (m *OnChainDataV2) GetStatus() []*FlagStatus {
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
	NftAttributesValue []*NftAttributeValue   `protobuf:"bytes,7,rep,name=nft_attributes_value,json=nftAttributesValue,proto3" json:"nft_attributes_value,omitempty"`
}

func (m *OnChainDataV1) Reset()         { *m = OnChainDataV1{} }
func (m *OnChainDataV1) String() string { return proto.CompactTextString(m) }
func (*OnChainDataV1) ProtoMessage()    {}
func (*OnChainDataV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_35d167410338c830, []int{3}
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
	proto.RegisterType((*OnChainDataV2)(nil), "thesixnetwork.sixnft.nftmngr.OnChainDataV2")
	proto.RegisterType((*OnChainDataV1)(nil), "thesixnetwork.sixnft.nftmngr.OnChainDataV1")
}

func init() { proto.RegisterFile("nftmngr/on_chain_data.proto", fileDescriptor_35d167410338c830) }

var fileDescriptor_35d167410338c830 = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcf, 0x6b, 0x13, 0x41,
	0x18, 0xcd, 0x64, 0x6b, 0x5a, 0x27, 0x49, 0x2b, 0x43, 0x0f, 0x4b, 0x95, 0x35, 0x8d, 0x82, 0x7b,
	0xda, 0xa5, 0xf5, 0x2e, 0xfe, 0x28, 0x82, 0x97, 0x2a, 0x53, 0x28, 0x22, 0xc2, 0x32, 0xd9, 0xcc,
	0x26, 0x43, 0xb3, 0xb3, 0x75, 0xe6, 0xdb, 0x5a, 0xff, 0x0b, 0xff, 0x2b, 0x3d, 0xf6, 0xe8, 0x51,
	0x92, 0x3f, 0xc0, 0xbb, 0x27, 0xc9, 0xcc, 0x6e, 0xba, 0x0b, 0x25, 0x42, 0x42, 0x6e, 0xc3, 0xe3,
	0xbd, 0x37, 0xbc, 0xf7, 0x7d, 0x7c, 0xf8, 0xa1, 0x4c, 0x20, 0x95, 0x23, 0x15, 0x66, 0x32, 0x8a,
	0xc7, 0x4c, 0xc8, 0x68, 0xc8, 0x80, 0x05, 0x97, 0x2a, 0x83, 0x8c, 0x3c, 0x82, 0x31, 0xd7, 0xe2,
	0x5a, 0x72, 0xf8, 0x9a, 0xa9, 0x8b, 0x60, 0xfe, 0x4c, 0x20, 0x28, 0x14, 0x07, 0xfd, 0x52, 0xca,
	0x00, 0x94, 0x18, 0xe4, 0xc0, 0xa3, 0x21, 0x4f, 0x84, 0x14, 0x20, 0x32, 0x69, 0x1d, 0x0e, 0xf6,
	0x17, 0x9c, 0xb8, 0x82, 0x1e, 0x96, 0xa8, 0x4c, 0x20, 0xba, 0x55, 0x5f, 0xb1, 0x49, 0xce, 0x2d,
	0xa5, 0xff, 0x01, 0xe3, 0xb7, 0x13, 0x36, 0x3a, 0x03, 0x06, 0xb9, 0x26, 0x8f, 0x71, 0x5b, 0x9b,
	0x57, 0x24, 0x59, 0xca, 0x5d, 0xd4, 0x43, 0xfe, 0x7d, 0x8a, 0x2d, 0x74, 0xca, 0x52, 0x4e, 0x0e,
	0x71, 0xa7, 0x20, 0x18, 0x13, 0xb7, 0xd9, 0x43, 0xfe, 0x0e, 0x2d, 0x44, 0xe7, 0x73, 0xa8, 0xff,
	0xa3, 0x89, 0xdb, 0xef, 0xe5, 0x9b, 0x79, 0xc6, 0x13, 0x06, 0x8c, 0x7c, 0xc4, 0xbb, 0xb5, 0xef,
	0xb5, 0x8b, 0x7a, 0x8e, 0xdf, 0x3e, 0x3e, 0x0a, 0x96, 0xa5, 0x0e, 0x5e, 0x95, 0xfc, 0x93, 0x45,
	0x56, 0xda, 0x95, 0x09, 0x2c, 0x70, 0x4d, 0x3e, 0xe3, 0x07, 0x90, 0x5d, 0x70, 0x59, 0xf5, 0x6e,
	0xae, 0xea, 0xbd, 0x67, 0xac, 0x2a, 0xee, 0x2f, 0xf0, 0xb6, 0x2d, 0x53, 0xbb, 0x8e, 0x31, 0x7d,
	0xfa, 0x1f, 0x53, 0x43, 0xa6, 0xa5, 0x88, 0xbc, 0xc4, 0x2d, 0x5b, 0x8b, 0xbb, 0x65, 0xe4, 0xfe,
	0x72, 0xf9, 0xed, 0x14, 0x68, 0xa1, 0xeb, 0xff, 0x41, 0xb8, 0x5b, 0x69, 0xf2, 0xfc, 0xf8, 0xce,
	0xc4, 0x68, 0x13, 0x89, 0x9b, 0xeb, 0x25, 0x76, 0x56, 0x4c, 0xfc, 0xd7, 0xa9, 0x27, 0x3e, 0x22,
	0xcf, 0xf0, 0x9e, 0xe2, 0x57, 0x9c, 0x4d, 0x22, 0xc5, 0xbf, 0xe4, 0x42, 0xf1, 0xa1, 0xd9, 0xca,
	0x1d, 0xba, 0x6b, 0x61, 0x5a, 0xa0, 0xe4, 0x09, 0xee, 0x16, 0x44, 0xcd, 0x63, 0xc5, 0xc1, 0xac,
	0x66, 0x87, 0x76, 0x2c, 0x78, 0x66, 0xb0, 0x3b, 0x76, 0xd1, 0xd9, 0xe0, 0x2e, 0x6e, 0x6d, 0x62,
	0x32, 0xf7, 0xd6, 0x9b, 0x4c, 0x6b, 0xb5, 0xc9, 0x10, 0x86, 0xf7, 0xeb, 0xcd, 0x15, 0x07, 0x60,
	0xdb, 0xf8, 0x85, 0xcb, 0xfd, 0x4e, 0x2b, 0x55, 0x99, 0x23, 0x41, 0x49, 0xad, 0x3d, 0x83, 0xbd,
	0x7e, 0xf7, 0x73, 0xea, 0xa1, 0x9b, 0xa9, 0x87, 0x7e, 0x4f, 0x3d, 0xf4, 0x7d, 0xe6, 0x35, 0x6e,
	0x66, 0x5e, 0xe3, 0xd7, 0xcc, 0x6b, 0x7c, 0x0a, 0x47, 0x02, 0xc6, 0xf9, 0x20, 0x88, 0xb3, 0x34,
	0xac, 0x7d, 0x14, 0xda, 0x8f, 0xc2, 0xeb, 0xb0, 0xbc, 0x74, 0xf0, 0xed, 0x92, 0xeb, 0x41, 0xcb,
	0x1c, 0xb7, 0xe7, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x77, 0x40, 0x4d, 0x76, 0x05, 0x00,
	0x00,
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
			dAtA[i] = 0x22
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
			dAtA[i] = 0x1a
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
			dAtA[i] = 0x12
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
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *OnChainDataV2) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OnChainDataV2) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OnChainDataV2) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			dAtA[i] = 0x1a
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
			dAtA[i] = 0x12
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
			dAtA[i] = 0xa
		}
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
			dAtA[i] = 0x3a
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
	return n
}

func (m *OnChainDataV2) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
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
		case 2:
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
		case 3:
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
		case 4:
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
func (m *OnChainDataV2) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: OnChainDataV2: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OnChainDataV2: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
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
		case 2:
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
		case 3:
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
		case 7:
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
