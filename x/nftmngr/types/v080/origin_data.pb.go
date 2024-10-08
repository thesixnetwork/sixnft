// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftmngr/v080/origin_data.proto

package v080

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
	return fileDescriptor_4c2c6cc5802a2afd, []int{0}
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
	return fileDescriptor_4c2c6cc5802a2afd, []int{1}
}

type OriginData struct {
	OriginChain           string                 `protobuf:"bytes,1,opt,name=origin_chain,json=originChain,proto3" json:"origin_chain,omitempty"`
	OriginContractAddress string                 `protobuf:"bytes,2,opt,name=origin_contract_address,json=originContractAddress,proto3" json:"origin_contract_address,omitempty"`
	OriginBaseUri         string                 `protobuf:"bytes,3,opt,name=origin_base_uri,json=originBaseUri,proto3" json:"origin_base_uri,omitempty"`
	AttributeOverriding   AttributeOverriding    `protobuf:"varint,4,opt,name=attribute_overriding,json=attributeOverriding,proto3,enum=thesixnetwork.sixnft.nftmngr.v080.AttributeOverriding" json:"attribute_overriding,omitempty"`
	MetadataFormat        string                 `protobuf:"bytes,5,opt,name=metadata_format,json=metadataFormat,proto3" json:"metadata_format,omitempty"`
	OriginAttributes      []*AttributeDefinition `protobuf:"bytes,6,rep,name=origin_attributes,json=originAttributes,proto3" json:"origin_attributes,omitempty"`
	UriRetrievalMethod    URIRetrievalMethod     `protobuf:"varint,7,opt,name=uri_retrieval_method,json=uriRetrievalMethod,proto3,enum=thesixnetwork.sixnft.nftmngr.v080.URIRetrievalMethod" json:"uri_retrieval_method,omitempty"`
}

func (m *OriginData) Reset()         { *m = OriginData{} }
func (m *OriginData) String() string { return proto.CompactTextString(m) }
func (*OriginData) ProtoMessage()    {}
func (*OriginData) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c2c6cc5802a2afd, []int{0}
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
	proto.RegisterEnum("thesixnetwork.sixnft.nftmngr.v080.AttributeOverriding", AttributeOverriding_name, AttributeOverriding_value)
	proto.RegisterEnum("thesixnetwork.sixnft.nftmngr.v080.URIRetrievalMethod", URIRetrievalMethod_name, URIRetrievalMethod_value)
	proto.RegisterType((*OriginData)(nil), "thesixnetwork.sixnft.nftmngr.v080.OriginData")
}

func init() { proto.RegisterFile("nftmngr/v080/origin_data.proto", fileDescriptor_4c2c6cc5802a2afd) }

var fileDescriptor_4c2c6cc5802a2afd = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x8d, 0x49, 0x1b, 0xe8, 0x14, 0xda, 0xb0, 0x2d, 0xc2, 0xe2, 0x60, 0xa5, 0x1c, 0x68, 0xa8,
	0x90, 0x5d, 0xf1, 0x51, 0x71, 0x4d, 0xda, 0x02, 0x11, 0x22, 0x91, 0x0c, 0xbd, 0x70, 0xb1, 0x36,
	0xf1, 0xc6, 0x19, 0x81, 0x77, 0xab, 0xf1, 0x38, 0x94, 0x7f, 0xc1, 0xcf, 0xe2, 0xd8, 0x23, 0x47,
	0x94, 0xfc, 0x02, 0xfe, 0x01, 0xca, 0xda, 0x0e, 0x2a, 0x41, 0x02, 0x6e, 0xab, 0xb7, 0xef, 0xbd,
	0x79, 0x6f, 0x34, 0xe0, 0xe9, 0x31, 0xa7, 0x3a, 0xa1, 0x60, 0x7a, 0xf8, 0xfc, 0x30, 0x30, 0x84,
	0x09, 0xea, 0x28, 0x96, 0x2c, 0xfd, 0x73, 0x32, 0x6c, 0xc4, 0x1e, 0x4f, 0x54, 0x86, 0x17, 0x5a,
	0xf1, 0x27, 0x43, 0x1f, 0xfc, 0xc5, 0x73, 0xcc, 0x7e, 0x29, 0xf2, 0x17, 0xa2, 0x7b, 0xfb, 0x57,
	0x2c, 0x24, 0x33, 0xe1, 0x30, 0x67, 0x15, 0xc5, 0x6a, 0x8c, 0x1a, 0x19, 0x8d, 0x2e, 0xbc, 0xee,
	0xff, 0xa8, 0x03, 0x0c, 0xec, 0x84, 0x13, 0xc9, 0x52, 0xec, 0xc1, 0xcd, 0x72, 0xde, 0x68, 0x22,
	0x51, 0xbb, 0x4e, 0xcb, 0x69, 0x6f, 0x84, 0x9b, 0x05, 0x76, 0xbc, 0x80, 0xc4, 0x11, 0xdc, 0xad,
	0x28, 0x46, 0x33, 0xc9, 0x11, 0x47, 0x32, 0x8e, 0x49, 0x65, 0x99, 0x7b, 0xcd, 0xb2, 0xef, 0x94,
	0xec, 0xf2, 0xb7, 0x53, 0x7c, 0x8a, 0x07, 0xb0, 0x5d, 0xea, 0x86, 0x32, 0x53, 0x51, 0x4e, 0xe8,
	0xd6, 0x2d, 0xff, 0x56, 0x01, 0x77, 0x65, 0xa6, 0xce, 0x08, 0x05, 0xc2, 0xee, 0xaf, 0xbc, 0x66,
	0xaa, 0x88, 0x30, 0x46, 0x9d, 0xb8, 0x6b, 0x2d, 0xa7, 0xbd, 0xf5, 0xf8, 0xc8, 0xff, 0x6b, 0x79,
	0xbf, 0x53, 0xc9, 0x07, 0x4b, 0x75, 0xb8, 0x23, 0x57, 0x41, 0xb1, 0x0f, 0xdb, 0xa9, 0x62, 0xb9,
	0x58, 0x6d, 0x34, 0x36, 0x94, 0x4a, 0x76, 0xd7, 0x6d, 0xa4, 0xad, 0x0a, 0x7e, 0x61, 0x51, 0x31,
	0x82, 0xdb, 0x65, 0xf6, 0xa5, 0x4d, 0xe6, 0x36, 0x5a, 0xf5, 0xf6, 0xe6, 0xff, 0x05, 0x3a, 0x59,
	0xae, 0x3f, 0x6c, 0x16, 0x86, 0xcb, 0xaf, 0x4c, 0x24, 0xb0, 0x9b, 0x13, 0x46, 0xa4, 0x98, 0x50,
	0x4d, 0xe5, 0xc7, 0x28, 0x55, 0x3c, 0x31, 0xb1, 0x7b, 0xdd, 0x16, 0x7f, 0xf6, 0x0f, 0x73, 0xce,
	0xc2, 0x5e, 0x58, 0xa9, 0xdf, 0x58, 0x71, 0x28, 0x72, 0xc2, 0xdf, 0xb0, 0x83, 0x47, 0xb0, 0xf3,
	0x87, 0x15, 0x09, 0x80, 0xc6, 0x20, 0xec, 0xbd, 0xec, 0xf5, 0x9b, 0x35, 0xb1, 0x01, 0xeb, 0xc7,
	0xaf, 0x3a, 0xbd, 0x7e, 0xd3, 0x39, 0x78, 0x08, 0x62, 0xd5, 0x57, 0xdc, 0x80, 0xb5, 0x6e, 0xe7,
	0xed, 0x69, 0x41, 0x7d, 0x37, 0x78, 0x7d, 0xda, 0x6f, 0x3a, 0xdd, 0xfe, 0xd7, 0x99, 0xe7, 0x5c,
	0xce, 0x3c, 0xe7, 0xfb, 0xcc, 0x73, 0xbe, 0xcc, 0xbd, 0xda, 0xe5, 0xdc, 0xab, 0x7d, 0x9b, 0x7b,
	0xb5, 0xf7, 0x4f, 0x13, 0xe4, 0x49, 0x3e, 0xf4, 0x47, 0x26, 0x0d, 0xae, 0xf4, 0x08, 0x8a, 0x1e,
	0xc1, 0x45, 0x50, 0x5d, 0x2c, 0x7f, 0x3e, 0x57, 0x99, 0xbd, 0xdb, 0x61, 0xc3, 0xde, 0xe8, 0x93,
	0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xae, 0xa4, 0x0d, 0x24, 0x11, 0x03, 0x00, 0x00,
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
