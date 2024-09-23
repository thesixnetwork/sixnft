// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftmngr/v063/origin_data.proto

package v063

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

type AttributeOverriding int32

const (
	AttributeOverriding_ORIGIN AttributeOverriding = 0
	AttributeOverriding_CHAIN  AttributeOverriding = 1
)

var AttributeOverriding_name = map[int32]string{
	0: "ORIGIN",
	1: "CHAIN",
}

var AttributeOverriding_value = map[string]int32{
	"ORIGIN": 0,
	"CHAIN":  1,
}

func (x AttributeOverriding) String() string {
	return proto.EnumName(AttributeOverriding_name, int32(x))
}

func (AttributeOverriding) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_666311fb5d1f88a0, []int{0}
}

type URIRetrievalMethod int32

const (
	URIRetrievalMethod_BASE  URIRetrievalMethod = 0
	URIRetrievalMethod_TOKEN URIRetrievalMethod = 1
)

var URIRetrievalMethod_name = map[int32]string{
	0: "BASE",
	1: "TOKEN",
}

var URIRetrievalMethod_value = map[string]int32{
	"BASE":  0,
	"TOKEN": 1,
}

func (x URIRetrievalMethod) String() string {
	return proto.EnumName(URIRetrievalMethod_name, int32(x))
}

func (URIRetrievalMethod) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_666311fb5d1f88a0, []int{1}
}

type OriginData struct {
	OriginChain           string                 `protobuf:"bytes,1,opt,name=origin_chain,json=originChain,proto3" json:"origin_chain,omitempty"`
	OriginContractAddress string                 `protobuf:"bytes,2,opt,name=origin_contract_address,json=originContractAddress,proto3" json:"origin_contract_address,omitempty"`
	OriginBaseUri         string                 `protobuf:"bytes,3,opt,name=origin_base_uri,json=originBaseUri,proto3" json:"origin_base_uri,omitempty"`
	AttributeOverriding   AttributeOverriding    `protobuf:"varint,4,opt,name=attribute_overriding,json=attributeOverriding,proto3,enum=thesixnetwork.sixnft.nftmngr.v063.AttributeOverriding" json:"attribute_overriding,omitempty"`
	MetadataFormat        string                 `protobuf:"bytes,5,opt,name=metadata_format,json=metadataFormat,proto3" json:"metadata_format,omitempty"`
	OriginAttributes      []*AttributeDefinition `protobuf:"bytes,6,rep,name=origin_attributes,json=originAttributes,proto3" json:"origin_attributes,omitempty"`
	UriRetrievalMethod    URIRetrievalMethod     `protobuf:"varint,7,opt,name=uri_retrieval_method,json=uriRetrievalMethod,proto3,enum=thesixnetwork.sixnft.nftmngr.v063.URIRetrievalMethod" json:"uri_retrieval_method,omitempty"`
}

func (m *OriginData) Reset()         { *m = OriginData{} }
func (m *OriginData) String() string { return proto.CompactTextString(m) }
func (*OriginData) ProtoMessage()    {}
func (*OriginData) Descriptor() ([]byte, []int) {
	return fileDescriptor_666311fb5d1f88a0, []int{0}
}
func (m *OriginData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OriginData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OriginData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OriginData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OriginData.Merge(m, src)
}
func (m *OriginData) XXX_Size() int {
	return m.Size()
}
func (m *OriginData) XXX_DiscardUnknown() {
	xxx_messageInfo_OriginData.DiscardUnknown(m)
}

var xxx_messageInfo_OriginData proto.InternalMessageInfo

func (m *OriginData) GetOriginChain() string {
	if m != nil {
		return m.OriginChain
	}
	return ""
}

func (m *OriginData) GetOriginContractAddress() string {
	if m != nil {
		return m.OriginContractAddress
	}
	return ""
}

func (m *OriginData) GetOriginBaseUri() string {
	if m != nil {
		return m.OriginBaseUri
	}
	return ""
}

func (m *OriginData) GetAttributeOverriding() AttributeOverriding {
	if m != nil {
		return m.AttributeOverriding
	}
	return AttributeOverriding_ORIGIN
}

func (m *OriginData) GetMetadataFormat() string {
	if m != nil {
		return m.MetadataFormat
	}
	return ""
}

func (m *OriginData) GetOriginAttributes() []*AttributeDefinition {
	if m != nil {
		return m.OriginAttributes
	}
	return nil
}

func (m *OriginData) GetUriRetrievalMethod() URIRetrievalMethod {
	if m != nil {
		return m.UriRetrievalMethod
	}
	return URIRetrievalMethod_BASE
}

func init() {
	proto.RegisterEnum("thesixnetwork.sixnft.nftmngr.v063.AttributeOverriding", AttributeOverriding_name, AttributeOverriding_value)
	proto.RegisterEnum("thesixnetwork.sixnft.nftmngr.v063.URIRetrievalMethod", URIRetrievalMethod_name, URIRetrievalMethod_value)
	proto.RegisterType((*OriginData)(nil), "thesixnetwork.sixnft.nftmngr.v063.OriginData")
}

func init() { proto.RegisterFile("nftmngr/v063/origin_data.proto", fileDescriptor_666311fb5d1f88a0) }

var fileDescriptor_666311fb5d1f88a0 = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0x49, 0x1b, 0xe8, 0x14, 0xda, 0xb0, 0x2d, 0xc2, 0xe2, 0x60, 0xa5, 0x1c, 0x68, 0xa8,
	0x90, 0x8d, 0x28, 0xe4, 0x9e, 0xb4, 0x05, 0x22, 0x44, 0x22, 0x19, 0x7a, 0xe1, 0x62, 0x6d, 0xe2,
	0x8d, 0x33, 0x02, 0xef, 0x56, 0xe3, 0x71, 0x28, 0x7f, 0xc1, 0x67, 0x71, 0xec, 0x91, 0x23, 0x4a,
	0xbe, 0x80, 0x3f, 0x40, 0x59, 0xdb, 0x41, 0xa5, 0x48, 0xd0, 0xdb, 0xea, 0xed, 0x7b, 0x6f, 0xde,
	0x1b, 0x0d, 0x78, 0x7a, 0xc2, 0xa9, 0x4e, 0x28, 0x98, 0x3d, 0xed, 0x1c, 0x06, 0x86, 0x30, 0x41,
	0x1d, 0xc5, 0x92, 0xa5, 0x7f, 0x46, 0x86, 0x8d, 0xd8, 0xe3, 0xa9, 0xca, 0xf0, 0x5c, 0x2b, 0xfe,
	0x6c, 0xe8, 0xa3, 0xbf, 0x7c, 0x4e, 0xd8, 0x2f, 0x45, 0xfe, 0x52, 0xf4, 0x60, 0xff, 0x92, 0x85,
	0x64, 0x26, 0x1c, 0xe5, 0xac, 0xa2, 0x58, 0x4d, 0x50, 0x23, 0xa3, 0xd1, 0x85, 0xd7, 0xc3, 0x9f,
	0x75, 0x80, 0xa1, 0x9d, 0x70, 0x2c, 0x59, 0x8a, 0x3d, 0xb8, 0x5d, 0xce, 0x1b, 0x4f, 0x25, 0x6a,
	0xd7, 0x69, 0x39, 0xed, 0x8d, 0x70, 0xb3, 0xc0, 0x8e, 0x96, 0x90, 0xe8, 0xc0, 0xfd, 0x8a, 0x62,
	0x34, 0x93, 0x1c, 0x73, 0x24, 0xe3, 0x98, 0x54, 0x96, 0xb9, 0x37, 0x2c, 0xfb, 0x5e, 0xc9, 0x2e,
	0x7f, 0xbb, 0xc5, 0xa7, 0x78, 0x04, 0xdb, 0xa5, 0x6e, 0x24, 0x33, 0x15, 0xe5, 0x84, 0x6e, 0xdd,
	0xf2, 0xef, 0x14, 0x70, 0x4f, 0x66, 0xea, 0x94, 0x50, 0x20, 0xec, 0xfe, 0xce, 0x6b, 0x66, 0x8a,
	0x08, 0x63, 0xd4, 0x89, 0xbb, 0xd6, 0x72, 0xda, 0x5b, 0xcf, 0x3a, 0xfe, 0x3f, 0xcb, 0xfb, 0xdd,
	0x4a, 0x3e, 0x5c, 0xa9, 0xc3, 0x1d, 0x79, 0x15, 0x14, 0xfb, 0xb0, 0x9d, 0x2a, 0x96, 0xcb, 0xd5,
	0x46, 0x13, 0x43, 0xa9, 0x64, 0x77, 0xdd, 0x46, 0xda, 0xaa, 0xe0, 0x97, 0x16, 0x15, 0x63, 0xb8,
	0x5b, 0x66, 0x5f, 0xd9, 0x64, 0x6e, 0xa3, 0x55, 0x6f, 0x6f, 0x5e, 0x2f, 0xd0, 0xf1, 0x6a, 0xfd,
	0x61, 0xb3, 0x30, 0x5c, 0x7d, 0x65, 0x22, 0x81, 0xdd, 0x9c, 0x30, 0x22, 0xc5, 0x84, 0x6a, 0x26,
	0x3f, 0x45, 0xa9, 0xe2, 0xa9, 0x89, 0xdd, 0x9b, 0xb6, 0xf8, 0x8b, 0xff, 0x98, 0x73, 0x1a, 0xf6,
	0xc3, 0x4a, 0xfd, 0xd6, 0x8a, 0x43, 0x91, 0x13, 0xfe, 0x81, 0x1d, 0x3c, 0x81, 0x9d, 0xbf, 0xac,
	0x48, 0x00, 0x34, 0x86, 0x61, 0xff, 0x55, 0x7f, 0xd0, 0xac, 0x89, 0x0d, 0x58, 0x3f, 0x7a, 0xdd,
	0xed, 0x0f, 0x9a, 0xce, 0xc1, 0x63, 0x10, 0x57, 0x7d, 0xc5, 0x2d, 0x58, 0xeb, 0x75, 0xdf, 0x9d,
	0x14, 0xd4, 0xf7, 0xc3, 0x37, 0x27, 0x83, 0xa6, 0xd3, 0x1b, 0x7c, 0x9b, 0x7b, 0xce, 0xc5, 0xdc,
	0x73, 0x7e, 0xcc, 0x3d, 0xe7, 0xeb, 0xc2, 0xab, 0x5d, 0x2c, 0xbc, 0xda, 0xf7, 0x85, 0x57, 0xfb,
	0xf0, 0x3c, 0x41, 0x9e, 0xe6, 0x23, 0x7f, 0x6c, 0xd2, 0xe0, 0x52, 0x8f, 0xa0, 0xe8, 0x11, 0x9c,
	0x07, 0xd5, 0xc5, 0xf2, 0x97, 0x33, 0x95, 0xd9, 0xbb, 0x1d, 0x35, 0xec, 0x8d, 0x1e, 0xfe, 0x0a,
	0x00, 0x00, 0xff, 0xff, 0xee, 0x55, 0xa2, 0xf1, 0x11, 0x03, 0x00, 0x00,
}

func (m *OriginData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OriginData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OriginData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UriRetrievalMethod != 0 {
		i = encodeVarintOriginData(dAtA, i, uint64(m.UriRetrievalMethod))
		i--
		dAtA[i] = 0x38
	}
	if len(m.OriginAttributes) > 0 {
		for iNdEx := len(m.OriginAttributes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OriginAttributes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOriginData(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.MetadataFormat) > 0 {
		i -= len(m.MetadataFormat)
		copy(dAtA[i:], m.MetadataFormat)
		i = encodeVarintOriginData(dAtA, i, uint64(len(m.MetadataFormat)))
		i--
		dAtA[i] = 0x2a
	}
	if m.AttributeOverriding != 0 {
		i = encodeVarintOriginData(dAtA, i, uint64(m.AttributeOverriding))
		i--
		dAtA[i] = 0x20
	}
	if len(m.OriginBaseUri) > 0 {
		i -= len(m.OriginBaseUri)
		copy(dAtA[i:], m.OriginBaseUri)
		i = encodeVarintOriginData(dAtA, i, uint64(len(m.OriginBaseUri)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.OriginContractAddress) > 0 {
		i -= len(m.OriginContractAddress)
		copy(dAtA[i:], m.OriginContractAddress)
		i = encodeVarintOriginData(dAtA, i, uint64(len(m.OriginContractAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.OriginChain) > 0 {
		i -= len(m.OriginChain)
		copy(dAtA[i:], m.OriginChain)
		i = encodeVarintOriginData(dAtA, i, uint64(len(m.OriginChain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOriginData(dAtA []byte, offset int, v uint64) int {
	offset -= sovOriginData(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OriginData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OriginChain)
	if l > 0 {
		n += 1 + l + sovOriginData(uint64(l))
	}
	l = len(m.OriginContractAddress)
	if l > 0 {
		n += 1 + l + sovOriginData(uint64(l))
	}
	l = len(m.OriginBaseUri)
	if l > 0 {
		n += 1 + l + sovOriginData(uint64(l))
	}
	if m.AttributeOverriding != 0 {
		n += 1 + sovOriginData(uint64(m.AttributeOverriding))
	}
	l = len(m.MetadataFormat)
	if l > 0 {
		n += 1 + l + sovOriginData(uint64(l))
	}
	if len(m.OriginAttributes) > 0 {
		for _, e := range m.OriginAttributes {
			l = e.Size()
			n += 1 + l + sovOriginData(uint64(l))
		}
	}
	if m.UriRetrievalMethod != 0 {
		n += 1 + sovOriginData(uint64(m.UriRetrievalMethod))
	}
	return n
}

func sovOriginData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOriginData(x uint64) (n int) {
	return sovOriginData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OriginData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOriginData
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
			return fmt.Errorf("proto: OriginData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OriginData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginChain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOriginData
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
				return ErrInvalidLengthOriginData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOriginData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OriginChain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOriginData
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
				return ErrInvalidLengthOriginData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOriginData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OriginContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginBaseUri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOriginData
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
				return ErrInvalidLengthOriginData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOriginData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OriginBaseUri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttributeOverriding", wireType)
			}
			m.AttributeOverriding = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOriginData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AttributeOverriding |= AttributeOverriding(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetadataFormat", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOriginData
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
				return ErrInvalidLengthOriginData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOriginData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MetadataFormat = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginAttributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOriginData
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
				return ErrInvalidLengthOriginData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOriginData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OriginAttributes = append(m.OriginAttributes, &AttributeDefinition{})
			if err := m.OriginAttributes[len(m.OriginAttributes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UriRetrievalMethod", wireType)
			}
			m.UriRetrievalMethod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOriginData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UriRetrievalMethod |= URIRetrievalMethod(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOriginData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOriginData
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
func skipOriginData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOriginData
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
					return 0, ErrIntOverflowOriginData
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
					return 0, ErrIntOverflowOriginData
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
				return 0, ErrInvalidLengthOriginData
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOriginData
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOriginData
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOriginData        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOriginData          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOriginData = fmt.Errorf("proto: unexpected end of group")
)
