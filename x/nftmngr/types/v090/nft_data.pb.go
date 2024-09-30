// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftmngr/v090/nft_data.proto

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

type OwnerAddressType int32

const (
	OwnerAddressType_ORIGIN_ADDRESS   OwnerAddressType = 0
	OwnerAddressType_INTERNAL_ADDRESS OwnerAddressType = 1
)

var OwnerAddressType_name = map[int32]string{
	0: "ORIGIN_ADDRESS",
	1: "INTERNAL_ADDRESS",
}

var OwnerAddressType_value = map[string]int32{
	"ORIGIN_ADDRESS":   0,
	"INTERNAL_ADDRESS": 1,
}

func (x OwnerAddressType) String() string {
	return proto.EnumName(OwnerAddressType_name, int32(x))
}

func (OwnerAddressType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_24e4d687432a2df6, []int{0}
}

type NftData struct {
	NftSchemaCode     string               `protobuf:"bytes,1,opt,name=nft_schema_code,json=nftSchemaCode,proto3" json:"nft_schema_code,omitempty"`
	TokenId           string               `protobuf:"bytes,2,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	TokenOwner        string               `protobuf:"bytes,3,opt,name=token_owner,json=tokenOwner,proto3" json:"token_owner,omitempty"`
	OwnerAddressType  OwnerAddressType     `protobuf:"varint,4,opt,name=owner_address_type,json=ownerAddressType,proto3,enum=thesixnetwork.sixnft.nftmngr.v090.OwnerAddressType" json:"owner_address_type,omitempty"`
	OriginImage       string               `protobuf:"bytes,5,opt,name=origin_image,json=originImage,proto3" json:"origin_image,omitempty"`
	OnchainImage      string               `protobuf:"bytes,6,opt,name=onchain_image,json=onchainImage,proto3" json:"onchain_image,omitempty"`
	TokenUri          string               `protobuf:"bytes,7,opt,name=token_uri,json=tokenUri,proto3" json:"token_uri,omitempty"`
	OriginAttributes  []*NftAttributeValue `protobuf:"bytes,8,rep,name=origin_attributes,json=originAttributes,proto3" json:"origin_attributes,omitempty"`
	OnchainAttributes []*NftAttributeValue `protobuf:"bytes,9,rep,name=onchain_attributes,json=onchainAttributes,proto3" json:"onchain_attributes,omitempty"`
}

func (m *NftData) Reset()         { *m = NftData{} }
func (m *NftData) String() string { return proto.CompactTextString(m) }
func (*NftData) ProtoMessage()    {}
func (*NftData) Descriptor() ([]byte, []int) {
	return fileDescriptor_24e4d687432a2df6, []int{0}
}
func (m *NftData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NftData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NftData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NftData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NftData.Merge(m, src)
}
func (m *NftData) XXX_Size() int {
	return m.Size()
}
func (m *NftData) XXX_DiscardUnknown() {
	xxx_messageInfo_NftData.DiscardUnknown(m)
}

var xxx_messageInfo_NftData proto.InternalMessageInfo

func (m *NftData) GetNftSchemaCode() string {
	if m != nil {
		return m.NftSchemaCode
	}
	return ""
}

func (m *NftData) GetTokenId() string {
	if m != nil {
		return m.TokenId
	}
	return ""
}

func (m *NftData) GetTokenOwner() string {
	if m != nil {
		return m.TokenOwner
	}
	return ""
}

func (m *NftData) GetOwnerAddressType() OwnerAddressType {
	if m != nil {
		return m.OwnerAddressType
	}
	return OwnerAddressType_ORIGIN_ADDRESS
}

func (m *NftData) GetOriginImage() string {
	if m != nil {
		return m.OriginImage
	}
	return ""
}

func (m *NftData) GetOnchainImage() string {
	if m != nil {
		return m.OnchainImage
	}
	return ""
}

func (m *NftData) GetTokenUri() string {
	if m != nil {
		return m.TokenUri
	}
	return ""
}

func (m *NftData) GetOriginAttributes() []*NftAttributeValue {
	if m != nil {
		return m.OriginAttributes
	}
	return nil
}

func (m *NftData) GetOnchainAttributes() []*NftAttributeValue {
	if m != nil {
		return m.OnchainAttributes
	}
	return nil
}

func init() {
	proto.RegisterEnum("thesixnetwork.sixnft.nftmngr.v090.OwnerAddressType", OwnerAddressType_name, OwnerAddressType_value)
	proto.RegisterType((*NftData)(nil), "thesixnetwork.sixnft.nftmngr.v090.NftData")
}

func init() { proto.RegisterFile("nftmngr/v090/nft_data.proto", fileDescriptor_24e4d687432a2df6) }

var fileDescriptor_24e4d687432a2df6 = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x4d, 0x6b, 0xd4, 0x40,
	0x18, 0xc7, 0x37, 0xae, 0x76, 0xbb, 0xb3, 0x6d, 0x4d, 0x07, 0x0f, 0xd1, 0x42, 0xdc, 0x2a, 0x94,
	0xc5, 0x43, 0x52, 0xda, 0x5e, 0x04, 0x2f, 0xd1, 0x2d, 0x12, 0x90, 0x14, 0xb2, 0xd5, 0x83, 0x97,
	0x61, 0x36, 0x99, 0x24, 0x43, 0xcd, 0xcc, 0x32, 0x79, 0xd2, 0x97, 0x6f, 0xe1, 0xc7, 0xf2, 0xd8,
	0xa3, 0x47, 0xd9, 0xfd, 0x1c, 0x82, 0x64, 0x26, 0x1b, 0xb5, 0x1e, 0x14, 0xbc, 0x25, 0xbf, 0xe7,
	0xff, 0x7f, 0xde, 0xe6, 0x41, 0x7b, 0x22, 0x83, 0x52, 0xe4, 0xca, 0xbf, 0x3c, 0x7c, 0x79, 0xe8,
	0x8b, 0x0c, 0x48, 0x4a, 0x81, 0x7a, 0x0b, 0x25, 0x41, 0xe2, 0x7d, 0x28, 0x58, 0xc5, 0xaf, 0x05,
	0x83, 0x2b, 0xa9, 0x2e, 0xbc, 0xe6, 0x33, 0x03, 0xaf, 0x75, 0x78, 0x8d, 0xe3, 0xc9, 0xc1, 0x1f,
	0x7e, 0x0a, 0xa0, 0xf8, 0xbc, 0x06, 0x46, 0x2e, 0xe9, 0xa7, 0x9a, 0x99, 0x54, 0xcf, 0xbe, 0xf7,
	0xd1, 0x20, 0xca, 0x60, 0x4a, 0x81, 0xe2, 0x03, 0xf4, 0xb0, 0x11, 0x56, 0x49, 0xc1, 0x4a, 0x4a,
	0x12, 0x99, 0x32, 0xc7, 0x1a, 0x5b, 0x93, 0x61, 0xbc, 0x2d, 0x32, 0x98, 0x69, 0xfa, 0x46, 0xa6,
	0x0c, 0x3f, 0x46, 0x9b, 0x20, 0x2f, 0x98, 0x20, 0x3c, 0x75, 0xee, 0x69, 0xc1, 0x40, 0xff, 0x87,
	0x29, 0x7e, 0x8a, 0x46, 0x26, 0x24, 0xaf, 0x04, 0x53, 0x4e, 0x5f, 0x47, 0x91, 0x46, 0x67, 0x0d,
	0xc1, 0x14, 0x61, 0x1d, 0x22, 0x34, 0x4d, 0x15, 0xab, 0x2a, 0x02, 0x37, 0x0b, 0xe6, 0xdc, 0x1f,
	0x5b, 0x93, 0x9d, 0xa3, 0x63, 0xef, 0xaf, 0x73, 0x79, 0x3a, 0x4b, 0x60, 0xbc, 0xe7, 0x37, 0x0b,
	0x16, 0xdb, 0xf2, 0x0e, 0xc1, 0xfb, 0x68, 0x4b, 0x2a, 0x9e, 0x73, 0x41, 0x78, 0x49, 0x73, 0xe6,
	0x3c, 0xd0, 0x4d, 0x8c, 0x0c, 0x0b, 0x1b, 0x84, 0x9f, 0xa3, 0x6d, 0x29, 0x92, 0x82, 0x76, 0x9a,
	0x0d, 0xad, 0xd9, 0x6a, 0xa1, 0x11, 0xed, 0xa1, 0xa1, 0x99, 0xa5, 0x56, 0xdc, 0x19, 0x68, 0x81,
	0x99, 0xfb, 0xbd, 0xe2, 0x98, 0xa2, 0xdd, 0xb6, 0x48, 0xb7, 0xd7, 0xca, 0xd9, 0x1c, 0xf7, 0x27,
	0xa3, 0xa3, 0x93, 0x7f, 0x18, 0x23, 0xca, 0x20, 0x58, 0xfb, 0x3e, 0x34, 0xcf, 0x11, 0xdb, 0x26,
	0x5d, 0x47, 0x2b, 0x9c, 0x20, 0xbc, 0x6e, 0xf2, 0x97, 0x1a, 0xc3, 0xff, 0xa8, 0xb1, 0xdb, 0xe6,
	0xfb, 0x59, 0xe4, 0xc5, 0x2b, 0x64, 0xdf, 0x5d, 0x29, 0xc6, 0x68, 0xe7, 0x2c, 0x0e, 0xdf, 0x86,
	0x11, 0x09, 0xa6, 0xd3, 0xf8, 0x74, 0x36, 0xb3, 0x7b, 0xf8, 0x11, 0xb2, 0xc3, 0xe8, 0xfc, 0x34,
	0x8e, 0x82, 0x77, 0x1d, 0xb5, 0x5e, 0x47, 0x5f, 0x96, 0xae, 0x75, 0xbb, 0x74, 0xad, 0x6f, 0x4b,
	0xd7, 0xfa, 0xbc, 0x72, 0x7b, 0xb7, 0x2b, 0xb7, 0xf7, 0x75, 0xe5, 0xf6, 0x3e, 0x9e, 0xe4, 0x1c,
	0x8a, 0x7a, 0xee, 0x25, 0xb2, 0xf4, 0x7f, 0x6b, 0xd5, 0x37, 0xad, 0xfa, 0xd7, 0xfe, 0xfa, 0x42,
	0x9b, 0x0b, 0xa8, 0xf4, 0x9d, 0xce, 0x37, 0xf4, 0x51, 0x1e, 0xff, 0x08, 0x00, 0x00, 0xff, 0xff,
	0x87, 0x96, 0xcf, 0xa1, 0xfe, 0x02, 0x00, 0x00,
}

func (m *NftData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NftData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NftData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.OnchainAttributes) > 0 {
		for iNdEx := len(m.OnchainAttributes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OnchainAttributes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintNftData(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.OriginAttributes) > 0 {
		for iNdEx := len(m.OriginAttributes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OriginAttributes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintNftData(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.TokenUri) > 0 {
		i -= len(m.TokenUri)
		copy(dAtA[i:], m.TokenUri)
		i = encodeVarintNftData(dAtA, i, uint64(len(m.TokenUri)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.OnchainImage) > 0 {
		i -= len(m.OnchainImage)
		copy(dAtA[i:], m.OnchainImage)
		i = encodeVarintNftData(dAtA, i, uint64(len(m.OnchainImage)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.OriginImage) > 0 {
		i -= len(m.OriginImage)
		copy(dAtA[i:], m.OriginImage)
		i = encodeVarintNftData(dAtA, i, uint64(len(m.OriginImage)))
		i--
		dAtA[i] = 0x2a
	}
	if m.OwnerAddressType != 0 {
		i = encodeVarintNftData(dAtA, i, uint64(m.OwnerAddressType))
		i--
		dAtA[i] = 0x20
	}
	if len(m.TokenOwner) > 0 {
		i -= len(m.TokenOwner)
		copy(dAtA[i:], m.TokenOwner)
		i = encodeVarintNftData(dAtA, i, uint64(len(m.TokenOwner)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.TokenId) > 0 {
		i -= len(m.TokenId)
		copy(dAtA[i:], m.TokenId)
		i = encodeVarintNftData(dAtA, i, uint64(len(m.TokenId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.NftSchemaCode) > 0 {
		i -= len(m.NftSchemaCode)
		copy(dAtA[i:], m.NftSchemaCode)
		i = encodeVarintNftData(dAtA, i, uint64(len(m.NftSchemaCode)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNftData(dAtA []byte, offset int, v uint64) int {
	offset -= sovNftData(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NftData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NftSchemaCode)
	if l > 0 {
		n += 1 + l + sovNftData(uint64(l))
	}
	l = len(m.TokenId)
	if l > 0 {
		n += 1 + l + sovNftData(uint64(l))
	}
	l = len(m.TokenOwner)
	if l > 0 {
		n += 1 + l + sovNftData(uint64(l))
	}
	if m.OwnerAddressType != 0 {
		n += 1 + sovNftData(uint64(m.OwnerAddressType))
	}
	l = len(m.OriginImage)
	if l > 0 {
		n += 1 + l + sovNftData(uint64(l))
	}
	l = len(m.OnchainImage)
	if l > 0 {
		n += 1 + l + sovNftData(uint64(l))
	}
	l = len(m.TokenUri)
	if l > 0 {
		n += 1 + l + sovNftData(uint64(l))
	}
	if len(m.OriginAttributes) > 0 {
		for _, e := range m.OriginAttributes {
			l = e.Size()
			n += 1 + l + sovNftData(uint64(l))
		}
	}
	if len(m.OnchainAttributes) > 0 {
		for _, e := range m.OnchainAttributes {
			l = e.Size()
			n += 1 + l + sovNftData(uint64(l))
		}
	}
	return n
}

func sovNftData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNftData(x uint64) (n int) {
	return sovNftData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NftData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNftData
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
			return fmt.Errorf("proto: NftData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NftData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftSchemaCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftData
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
				return ErrInvalidLengthNftData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NftSchemaCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftData
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
				return ErrInvalidLengthNftData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenOwner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftData
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
				return ErrInvalidLengthNftData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenOwner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerAddressType", wireType)
			}
			m.OwnerAddressType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OwnerAddressType |= OwnerAddressType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginImage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftData
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
				return ErrInvalidLengthNftData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OriginImage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OnchainImage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftData
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
				return ErrInvalidLengthNftData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OnchainImage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenUri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftData
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
				return ErrInvalidLengthNftData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenUri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginAttributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftData
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
				return ErrInvalidLengthNftData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNftData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OriginAttributes = append(m.OriginAttributes, &NftAttributeValue{})
			if err := m.OriginAttributes[len(m.OriginAttributes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OnchainAttributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftData
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
				return ErrInvalidLengthNftData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNftData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OnchainAttributes = append(m.OnchainAttributes, &NftAttributeValue{})
			if err := m.OnchainAttributes[len(m.OnchainAttributes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNftData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNftData
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
func skipNftData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNftData
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
					return 0, ErrIntOverflowNftData
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
					return 0, ErrIntOverflowNftData
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
				return 0, ErrInvalidLengthNftData
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNftData
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNftData
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNftData        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNftData          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNftData = fmt.Errorf("proto: unexpected end of group")
)