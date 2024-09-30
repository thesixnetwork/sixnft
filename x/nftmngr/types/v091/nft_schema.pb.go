// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftmngr/v091/nft_schema.proto

package v091

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

type NFTSchema struct {
	Code              string       `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name              string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Owner             string       `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	Description       string       `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	OriginData        *OriginData  `protobuf:"bytes,5,opt,name=origin_data,json=originData,proto3" json:"origin_data,omitempty"`
	OnchainData       *OnChainData `protobuf:"bytes,6,opt,name=onchain_data,json=onchainData,proto3" json:"onchain_data,omitempty"`
	IsVerified        bool         `protobuf:"varint,7,opt,name=isVerified,proto3" json:"isVerified,omitempty"`
	MintAuthorization string       `protobuf:"bytes,8,opt,name=mint_authorization,json=mintAuthorization,proto3" json:"mint_authorization,omitempty"`
}

func (m *NFTSchema) Reset()         { *m = NFTSchema{} }
func (m *NFTSchema) String() string { return proto.CompactTextString(m) }
func (*NFTSchema) ProtoMessage()    {}
func (*NFTSchema) Descriptor() ([]byte, []int) {
	return fileDescriptor_520b7beb8c1bd742, []int{0}
}
func (m *NFTSchema) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NFTSchema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NFTSchema.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NFTSchema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NFTSchema.Merge(m, src)
}
func (m *NFTSchema) XXX_Size() int {
	return m.Size()
}
func (m *NFTSchema) XXX_DiscardUnknown() {
	xxx_messageInfo_NFTSchema.DiscardUnknown(m)
}

var xxx_messageInfo_NFTSchema proto.InternalMessageInfo

func (m *NFTSchema) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *NFTSchema) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NFTSchema) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *NFTSchema) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *NFTSchema) GetOriginData() *OriginData {
	if m != nil {
		return m.OriginData
	}
	return nil
}

func (m *NFTSchema) GetOnchainData() *OnChainData {
	if m != nil {
		return m.OnchainData
	}
	return nil
}

func (m *NFTSchema) GetIsVerified() bool {
	if m != nil {
		return m.IsVerified
	}
	return false
}

func (m *NFTSchema) GetMintAuthorization() string {
	if m != nil {
		return m.MintAuthorization
	}
	return ""
}

type NFTSchemaINPUT struct {
	Code              string            `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name              string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Owner             string            `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	Description       string            `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	SystemActioners   []string          `protobuf:"bytes,5,rep,name=system_actioners,json=systemActioners,proto3" json:"system_actioners,omitempty"`
	OriginData        *OriginData       `protobuf:"bytes,6,opt,name=origin_data,json=originData,proto3" json:"origin_data,omitempty"`
	OnchainData       *OnChainDataInput `protobuf:"bytes,7,opt,name=onchain_data,json=onchainData,proto3" json:"onchain_data,omitempty"`
	IsVerified        bool              `protobuf:"varint,8,opt,name=isVerified,proto3" json:"isVerified,omitempty"`
	MintAuthorization string            `protobuf:"bytes,9,opt,name=mint_authorization,json=mintAuthorization,proto3" json:"mint_authorization,omitempty"`
}

func (m *NFTSchemaINPUT) Reset()         { *m = NFTSchemaINPUT{} }
func (m *NFTSchemaINPUT) String() string { return proto.CompactTextString(m) }
func (*NFTSchemaINPUT) ProtoMessage()    {}
func (*NFTSchemaINPUT) Descriptor() ([]byte, []int) {
	return fileDescriptor_520b7beb8c1bd742, []int{1}
}
func (m *NFTSchemaINPUT) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NFTSchemaINPUT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NFTSchemaINPUT.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NFTSchemaINPUT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NFTSchemaINPUT.Merge(m, src)
}
func (m *NFTSchemaINPUT) XXX_Size() int {
	return m.Size()
}
func (m *NFTSchemaINPUT) XXX_DiscardUnknown() {
	xxx_messageInfo_NFTSchemaINPUT.DiscardUnknown(m)
}

var xxx_messageInfo_NFTSchemaINPUT proto.InternalMessageInfo

func (m *NFTSchemaINPUT) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *NFTSchemaINPUT) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NFTSchemaINPUT) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *NFTSchemaINPUT) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *NFTSchemaINPUT) GetSystemActioners() []string {
	if m != nil {
		return m.SystemActioners
	}
	return nil
}

func (m *NFTSchemaINPUT) GetOriginData() *OriginData {
	if m != nil {
		return m.OriginData
	}
	return nil
}

func (m *NFTSchemaINPUT) GetOnchainData() *OnChainDataInput {
	if m != nil {
		return m.OnchainData
	}
	return nil
}

func (m *NFTSchemaINPUT) GetIsVerified() bool {
	if m != nil {
		return m.IsVerified
	}
	return false
}

func (m *NFTSchemaINPUT) GetMintAuthorization() string {
	if m != nil {
		return m.MintAuthorization
	}
	return ""
}

func init() {
	proto.RegisterType((*NFTSchema)(nil), "thesixnetwork.sixnft.nftmngr.v091.NFTSchema")
	proto.RegisterType((*NFTSchemaINPUT)(nil), "thesixnetwork.sixnft.nftmngr.v091.NFTSchemaINPUT")
}

func init() { proto.RegisterFile("nftmngr/v091/nft_schema.proto", fileDescriptor_520b7beb8c1bd742) }

var fileDescriptor_520b7beb8c1bd742 = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0x9b, 0xed, 0xb6, 0xdb, 0x38, 0x88, 0x3f, 0x16, 0x07, 0x6b, 0x25, 0xac, 0xb0, 0xa7,
	0x70, 0x58, 0x07, 0x58, 0x2e, 0x1c, 0x17, 0x10, 0xd2, 0x5e, 0x02, 0x84, 0x65, 0x0f, 0x5c, 0x22,
	0x6f, 0xe2, 0x34, 0x16, 0x8a, 0x1d, 0xd9, 0x2e, 0xdb, 0xf2, 0x14, 0x88, 0xa7, 0xe2, 0xd8, 0x23,
	0x17, 0x24, 0xd4, 0xbe, 0x08, 0x8a, 0xd3, 0x96, 0xa4, 0x12, 0x50, 0x84, 0xf6, 0x36, 0xf9, 0x26,
	0xf3, 0xf3, 0xcc, 0x7c, 0x1a, 0x70, 0x4f, 0xe4, 0xa6, 0x14, 0x63, 0x15, 0x7e, 0x7c, 0xf8, 0xf4,
	0x51, 0x28, 0x72, 0x93, 0xe8, 0xb4, 0x60, 0x25, 0x25, 0x95, 0x92, 0x46, 0xc2, 0xfb, 0xa6, 0x60,
	0x9a, 0x4f, 0x05, 0x33, 0x57, 0x52, 0x7d, 0x20, 0x75, 0x98, 0x1b, 0xb2, 0xaa, 0x21, 0x75, 0xcd,
	0xa1, 0xdf, 0x21, 0x48, 0x91, 0xa4, 0x05, 0xe5, 0x22, 0xc9, 0xa8, 0x59, 0x41, 0x0e, 0x71, 0xf7,
	0x0f, 0xc5, 0xc7, 0x9d, 0xfc, 0xd1, 0xf7, 0x3d, 0xe0, 0x46, 0x2f, 0xcf, 0xdf, 0xda, 0x87, 0x21,
	0x04, 0xfb, 0xa9, 0xcc, 0x18, 0x72, 0x7c, 0x27, 0x70, 0x63, 0x1b, 0xd7, 0x9a, 0xa0, 0x25, 0x43,
	0x7b, 0x8d, 0x56, 0xc7, 0xf0, 0x2e, 0x18, 0xc8, 0x2b, 0xc1, 0x14, 0xea, 0x5b, 0xb1, 0xf9, 0x80,
	0x3e, 0xf0, 0x32, 0xa6, 0x53, 0xc5, 0x2b, 0xc3, 0xa5, 0x40, 0xfb, 0x36, 0xd7, 0x96, 0x60, 0x04,
	0xbc, 0x56, 0x0b, 0x68, 0xe0, 0x3b, 0x81, 0xf7, 0xf8, 0x98, 0xfc, 0x75, 0x50, 0xf2, 0xca, 0x56,
	0xbd, 0xa0, 0x86, 0xc6, 0x40, 0x6e, 0x62, 0xf8, 0x06, 0xdc, 0x90, 0xe2, 0xd7, 0xcc, 0x68, 0x68,
	0x81, 0x64, 0x17, 0xa0, 0x78, 0x5e, 0x97, 0x59, 0xa2, 0xb7, 0x62, 0x58, 0x24, 0x06, 0x80, 0xeb,
	0x0b, 0xa6, 0x78, 0xce, 0x59, 0x86, 0x0e, 0x7c, 0x27, 0x18, 0xc5, 0x2d, 0x05, 0x1e, 0x03, 0x58,
	0x72, 0x61, 0x12, 0x3a, 0x31, 0x85, 0x54, 0xfc, 0x13, 0xb5, 0xb3, 0x8e, 0xec, 0xac, 0x77, 0xea,
	0xcc, 0x69, 0x3b, 0x71, 0xf4, 0xa5, 0x0f, 0x6e, 0x6e, 0xf6, 0x7b, 0x16, 0xbd, 0x7e, 0x77, 0x7e,
	0xed, 0x4b, 0x7e, 0x00, 0x6e, 0xeb, 0x99, 0x36, 0xac, 0x4c, 0x68, 0x5a, 0x0b, 0x4c, 0x69, 0x34,
	0xf0, 0xfb, 0x81, 0x1b, 0xdf, 0x6a, 0xf4, 0xd3, 0xb5, 0xbc, 0xed, 0xc7, 0xf0, 0x7f, 0xfd, 0xb8,
	0xd8, 0xf2, 0xe3, 0xc0, 0x02, 0x4f, 0xfe, 0xcd, 0x8f, 0x33, 0x51, 0x4d, 0xcc, 0x9f, 0x4c, 0x19,
	0xed, 0x68, 0x8a, 0xfb, 0x1b, 0x53, 0x9e, 0x45, 0x5f, 0x17, 0xd8, 0x99, 0x2f, 0xb0, 0xf3, 0x63,
	0x81, 0x9d, 0xcf, 0x4b, 0xdc, 0x9b, 0x2f, 0x71, 0xef, 0xdb, 0x12, 0xf7, 0xde, 0x3f, 0x19, 0x73,
	0x53, 0x4c, 0x2e, 0x49, 0x2a, 0xcb, 0xb0, 0xd3, 0x74, 0xd8, 0x34, 0x1d, 0x4e, 0xc3, 0xf5, 0x41,
	0x99, 0x59, 0xc5, 0xb4, 0x3d, 0xab, 0xcb, 0xa1, 0xbd, 0xa5, 0x93, 0x9f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x3b, 0x24, 0x20, 0x64, 0xd1, 0x03, 0x00, 0x00,
}

func (m *NFTSchema) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NFTSchema) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NFTSchema) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MintAuthorization) > 0 {
		i -= len(m.MintAuthorization)
		copy(dAtA[i:], m.MintAuthorization)
		i = encodeVarintNftSchema(dAtA, i, uint64(len(m.MintAuthorization)))
		i--
		dAtA[i] = 0x42
	}
	if m.IsVerified {
		i--
		if m.IsVerified {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.OnchainData != nil {
		{
			size, err := m.OnchainData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintNftSchema(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.OriginData != nil {
		{
			size, err := m.OriginData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintNftSchema(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintNftSchema(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintNftSchema(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintNftSchema(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Code) > 0 {
		i -= len(m.Code)
		copy(dAtA[i:], m.Code)
		i = encodeVarintNftSchema(dAtA, i, uint64(len(m.Code)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *NFTSchemaINPUT) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NFTSchemaINPUT) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NFTSchemaINPUT) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MintAuthorization) > 0 {
		i -= len(m.MintAuthorization)
		copy(dAtA[i:], m.MintAuthorization)
		i = encodeVarintNftSchema(dAtA, i, uint64(len(m.MintAuthorization)))
		i--
		dAtA[i] = 0x4a
	}
	if m.IsVerified {
		i--
		if m.IsVerified {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if m.OnchainData != nil {
		{
			size, err := m.OnchainData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintNftSchema(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.OriginData != nil {
		{
			size, err := m.OriginData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintNftSchema(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if len(m.SystemActioners) > 0 {
		for iNdEx := len(m.SystemActioners) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.SystemActioners[iNdEx])
			copy(dAtA[i:], m.SystemActioners[iNdEx])
			i = encodeVarintNftSchema(dAtA, i, uint64(len(m.SystemActioners[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintNftSchema(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintNftSchema(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintNftSchema(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Code) > 0 {
		i -= len(m.Code)
		copy(dAtA[i:], m.Code)
		i = encodeVarintNftSchema(dAtA, i, uint64(len(m.Code)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNftSchema(dAtA []byte, offset int, v uint64) int {
	offset -= sovNftSchema(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NFTSchema) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Code)
	if l > 0 {
		n += 1 + l + sovNftSchema(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovNftSchema(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovNftSchema(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovNftSchema(uint64(l))
	}
	if m.OriginData != nil {
		l = m.OriginData.Size()
		n += 1 + l + sovNftSchema(uint64(l))
	}
	if m.OnchainData != nil {
		l = m.OnchainData.Size()
		n += 1 + l + sovNftSchema(uint64(l))
	}
	if m.IsVerified {
		n += 2
	}
	l = len(m.MintAuthorization)
	if l > 0 {
		n += 1 + l + sovNftSchema(uint64(l))
	}
	return n
}

func (m *NFTSchemaINPUT) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Code)
	if l > 0 {
		n += 1 + l + sovNftSchema(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovNftSchema(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovNftSchema(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovNftSchema(uint64(l))
	}
	if len(m.SystemActioners) > 0 {
		for _, s := range m.SystemActioners {
			l = len(s)
			n += 1 + l + sovNftSchema(uint64(l))
		}
	}
	if m.OriginData != nil {
		l = m.OriginData.Size()
		n += 1 + l + sovNftSchema(uint64(l))
	}
	if m.OnchainData != nil {
		l = m.OnchainData.Size()
		n += 1 + l + sovNftSchema(uint64(l))
	}
	if m.IsVerified {
		n += 2
	}
	l = len(m.MintAuthorization)
	if l > 0 {
		n += 1 + l + sovNftSchema(uint64(l))
	}
	return n
}

func sovNftSchema(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNftSchema(x uint64) (n int) {
	return sovNftSchema(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NFTSchema) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNftSchema
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
			return fmt.Errorf("proto: NFTSchema: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NFTSchema: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftSchema
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
				return ErrInvalidLengthNftSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Code = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftSchema
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
				return ErrInvalidLengthNftSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftSchema
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
				return ErrInvalidLengthNftSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftSchema
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
				return ErrInvalidLengthNftSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftSchema
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
				return ErrInvalidLengthNftSchema
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNftSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OriginData == nil {
				m.OriginData = &OriginData{}
			}
			if err := m.OriginData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OnchainData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftSchema
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
				return ErrInvalidLengthNftSchema
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNftSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OnchainData == nil {
				m.OnchainData = &OnChainData{}
			}
			if err := m.OnchainData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsVerified", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftSchema
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
			m.IsVerified = bool(v != 0)
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintAuthorization", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftSchema
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
				return ErrInvalidLengthNftSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintAuthorization = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNftSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNftSchema
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
func (m *NFTSchemaINPUT) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNftSchema
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
			return fmt.Errorf("proto: NFTSchemaINPUT: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NFTSchemaINPUT: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftSchema
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
				return ErrInvalidLengthNftSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Code = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftSchema
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
				return ErrInvalidLengthNftSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftSchema
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
				return ErrInvalidLengthNftSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftSchema
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
				return ErrInvalidLengthNftSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SystemActioners", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftSchema
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
				return ErrInvalidLengthNftSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SystemActioners = append(m.SystemActioners, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftSchema
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
				return ErrInvalidLengthNftSchema
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNftSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OriginData == nil {
				m.OriginData = &OriginData{}
			}
			if err := m.OriginData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OnchainData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftSchema
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
				return ErrInvalidLengthNftSchema
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNftSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OnchainData == nil {
				m.OnchainData = &OnChainDataInput{}
			}
			if err := m.OnchainData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsVerified", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftSchema
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
			m.IsVerified = bool(v != 0)
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintAuthorization", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftSchema
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
				return ErrInvalidLengthNftSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintAuthorization = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNftSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNftSchema
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
func skipNftSchema(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNftSchema
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
					return 0, ErrIntOverflowNftSchema
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
					return 0, ErrIntOverflowNftSchema
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
				return 0, ErrInvalidLengthNftSchema
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNftSchema
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNftSchema
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNftSchema        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNftSchema          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNftSchema = fmt.Errorf("proto: unexpected end of group")
)