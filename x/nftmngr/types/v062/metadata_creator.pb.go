// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftmngr/v062/metadata_creator.proto

package v062

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

type MapTokenToMinter struct {
	TokenId string `protobuf:"bytes,1,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	Minter  string `protobuf:"bytes,2,opt,name=minter,proto3" json:"minter,omitempty"`
}

func (m *MapTokenToMinter) Reset()         { *m = MapTokenToMinter{} }
func (m *MapTokenToMinter) String() string { return proto.CompactTextString(m) }
func (*MapTokenToMinter) ProtoMessage()    {}
func (*MapTokenToMinter) Descriptor() ([]byte, []int) {
	return fileDescriptor_6baa8ec8c826543a, []int{0}
}
func (m *MapTokenToMinter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MapTokenToMinter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MapTokenToMinter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MapTokenToMinter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MapTokenToMinter.Merge(m, src)
}
func (m *MapTokenToMinter) XXX_Size() int {
	return m.Size()
}
func (m *MapTokenToMinter) XXX_DiscardUnknown() {
	xxx_messageInfo_MapTokenToMinter.DiscardUnknown(m)
}

var xxx_messageInfo_MapTokenToMinter proto.InternalMessageInfo

func (m *MapTokenToMinter) GetTokenId() string {
	if m != nil {
		return m.TokenId
	}
	return ""
}

func (m *MapTokenToMinter) GetMinter() string {
	if m != nil {
		return m.Minter
	}
	return ""
}

type MetadataCreator struct {
	NftSchemaCode    string              `protobuf:"bytes,1,opt,name=nftSchemaCode,proto3" json:"nftSchemaCode,omitempty"`
	MetadataMintedBy []*MapTokenToMinter `protobuf:"bytes,2,rep,name=metadataMintedBy,proto3" json:"metadataMintedBy,omitempty"`
}

func (m *MetadataCreator) Reset()         { *m = MetadataCreator{} }
func (m *MetadataCreator) String() string { return proto.CompactTextString(m) }
func (*MetadataCreator) ProtoMessage()    {}
func (*MetadataCreator) Descriptor() ([]byte, []int) {
	return fileDescriptor_6baa8ec8c826543a, []int{1}
}
func (m *MetadataCreator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MetadataCreator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MetadataCreator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MetadataCreator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataCreator.Merge(m, src)
}
func (m *MetadataCreator) XXX_Size() int {
	return m.Size()
}
func (m *MetadataCreator) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataCreator.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataCreator proto.InternalMessageInfo

func (m *MetadataCreator) GetNftSchemaCode() string {
	if m != nil {
		return m.NftSchemaCode
	}
	return ""
}

func (m *MetadataCreator) GetMetadataMintedBy() []*MapTokenToMinter {
	if m != nil {
		return m.MetadataMintedBy
	}
	return nil
}

func init() {
	proto.RegisterType((*MapTokenToMinter)(nil), "thesixnetwork.sixnft.nftmngr.v062.MapTokenToMinter")
	proto.RegisterType((*MetadataCreator)(nil), "thesixnetwork.sixnft.nftmngr.v062.MetadataCreator")
}

func init() {
	proto.RegisterFile("nftmngr/v062/metadata_creator.proto", fileDescriptor_6baa8ec8c826543a)
}

var fileDescriptor_6baa8ec8c826543a = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xce, 0x4b, 0x2b, 0xc9,
	0xcd, 0x4b, 0x2f, 0xd2, 0x2f, 0x33, 0x30, 0x33, 0xd2, 0xcf, 0x4d, 0x2d, 0x49, 0x4c, 0x49, 0x2c,
	0x49, 0x8c, 0x4f, 0x2e, 0x4a, 0x4d, 0x2c, 0xc9, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x52, 0x2c, 0xc9, 0x48, 0x2d, 0xce, 0xac, 0xc8, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x03,
	0x31, 0xd3, 0x4a, 0xf4, 0xa0, 0x3a, 0xf5, 0x40, 0x3a, 0x95, 0x5c, 0xb9, 0x04, 0x7c, 0x13, 0x0b,
	0x42, 0xf2, 0xb3, 0x53, 0xf3, 0x42, 0xf2, 0x7d, 0x33, 0xf3, 0x4a, 0x52, 0x8b, 0x84, 0x24, 0xb9,
	0x38, 0x4a, 0x40, 0x02, 0xf1, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0xec, 0x60,
	0xbe, 0x67, 0x8a, 0x90, 0x18, 0x17, 0x5b, 0x2e, 0x58, 0x91, 0x04, 0x13, 0x58, 0x02, 0xca, 0x53,
	0x9a, 0xc1, 0xc8, 0xc5, 0xef, 0x0b, 0x75, 0x84, 0x33, 0xc4, 0x0d, 0x42, 0x2a, 0x5c, 0xbc, 0x79,
	0x69, 0x25, 0xc1, 0xc9, 0x19, 0xa9, 0xb9, 0x89, 0xce, 0xf9, 0x29, 0xa9, 0x50, 0xb3, 0x50, 0x05,
	0x85, 0xe2, 0xb9, 0x04, 0x60, 0xae, 0x07, 0x5b, 0x9f, 0xe2, 0x54, 0x29, 0xc1, 0xa4, 0xc0, 0xac,
	0xc1, 0x6d, 0x64, 0xac, 0x47, 0xd0, 0xf9, 0x7a, 0xe8, 0x6e, 0x0f, 0xc2, 0x30, 0xcc, 0xc9, 0xef,
	0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e,
	0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x4c, 0xd2, 0x33, 0x4b, 0x32, 0x4a,
	0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x51, 0xac, 0xd2, 0x87, 0x58, 0xa5, 0x5f, 0xa1, 0x0f, 0x0b,
	0xe5, 0x92, 0xca, 0x82, 0xd4, 0x62, 0x70, 0x58, 0x27, 0xb1, 0x81, 0xc3, 0xd6, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0x8e, 0xd1, 0xa9, 0xc7, 0x82, 0x01, 0x00, 0x00,
}

func (m *MapTokenToMinter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MapTokenToMinter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MapTokenToMinter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Minter) > 0 {
		i -= len(m.Minter)
		copy(dAtA[i:], m.Minter)
		i = encodeVarintMetadataCreator(dAtA, i, uint64(len(m.Minter)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.TokenId) > 0 {
		i -= len(m.TokenId)
		copy(dAtA[i:], m.TokenId)
		i = encodeVarintMetadataCreator(dAtA, i, uint64(len(m.TokenId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MetadataCreator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MetadataCreator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MetadataCreator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MetadataMintedBy) > 0 {
		for iNdEx := len(m.MetadataMintedBy) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MetadataMintedBy[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMetadataCreator(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.NftSchemaCode) > 0 {
		i -= len(m.NftSchemaCode)
		copy(dAtA[i:], m.NftSchemaCode)
		i = encodeVarintMetadataCreator(dAtA, i, uint64(len(m.NftSchemaCode)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMetadataCreator(dAtA []byte, offset int, v uint64) int {
	offset -= sovMetadataCreator(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MapTokenToMinter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TokenId)
	if l > 0 {
		n += 1 + l + sovMetadataCreator(uint64(l))
	}
	l = len(m.Minter)
	if l > 0 {
		n += 1 + l + sovMetadataCreator(uint64(l))
	}
	return n
}

func (m *MetadataCreator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NftSchemaCode)
	if l > 0 {
		n += 1 + l + sovMetadataCreator(uint64(l))
	}
	if len(m.MetadataMintedBy) > 0 {
		for _, e := range m.MetadataMintedBy {
			l = e.Size()
			n += 1 + l + sovMetadataCreator(uint64(l))
		}
	}
	return n
}

func sovMetadataCreator(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMetadataCreator(x uint64) (n int) {
	return sovMetadataCreator(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MapTokenToMinter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetadataCreator
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
			return fmt.Errorf("proto: MapTokenToMinter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MapTokenToMinter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadataCreator
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
				return ErrInvalidLengthMetadataCreator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadataCreator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Minter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadataCreator
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
				return ErrInvalidLengthMetadataCreator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadataCreator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Minter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetadataCreator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMetadataCreator
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
func (m *MetadataCreator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetadataCreator
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
			return fmt.Errorf("proto: MetadataCreator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MetadataCreator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftSchemaCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadataCreator
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
				return ErrInvalidLengthMetadataCreator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadataCreator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NftSchemaCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetadataMintedBy", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadataCreator
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
				return ErrInvalidLengthMetadataCreator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMetadataCreator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MetadataMintedBy = append(m.MetadataMintedBy, &MapTokenToMinter{})
			if err := m.MetadataMintedBy[len(m.MetadataMintedBy)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetadataCreator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMetadataCreator
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
func skipMetadataCreator(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMetadataCreator
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
					return 0, ErrIntOverflowMetadataCreator
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
					return 0, ErrIntOverflowMetadataCreator
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
				return 0, ErrInvalidLengthMetadataCreator
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMetadataCreator
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMetadataCreator
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMetadataCreator        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMetadataCreator          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMetadataCreator = fmt.Errorf("proto: unexpected end of group")
)
