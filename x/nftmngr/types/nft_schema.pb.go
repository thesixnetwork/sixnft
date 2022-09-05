// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftmngr/nft_schema.proto

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

type NFTSchema struct {
	Code        string       `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name        string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Owner       string       `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	OriginData  *OriginData  `protobuf:"bytes,4,opt,name=origin_data,json=originData,proto3" json:"origin_data,omitempty"`
	OnchainData *OnChainData `protobuf:"bytes,5,opt,name=onchain_data,json=onchainData,proto3" json:"onchain_data,omitempty"`
}

func (m *NFTSchema) Reset()         { *m = NFTSchema{} }
func (m *NFTSchema) String() string { return proto.CompactTextString(m) }
func (*NFTSchema) ProtoMessage()    {}
func (*NFTSchema) Descriptor() ([]byte, []int) {
	return fileDescriptor_f35d7a245dabb499, []int{0}
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

func init() {
	proto.RegisterType((*NFTSchema)(nil), "sixnft.nftmngr.NFTSchema")
}

func init() { proto.RegisterFile("nftmngr/nft_schema.proto", fileDescriptor_f35d7a245dabb499) }

var fileDescriptor_f35d7a245dabb499 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc8, 0x4b, 0x2b, 0xc9,
	0xcd, 0x4b, 0x2f, 0xd2, 0xcf, 0x4b, 0x2b, 0x89, 0x2f, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0xd4, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2b, 0xce, 0xac, 0xc8, 0x4b, 0x2b, 0xd1, 0x83, 0x2a, 0x90,
	0x92, 0x86, 0xa9, 0xcc, 0xcf, 0x8b, 0x4f, 0xce, 0x48, 0xcc, 0xcc, 0x8b, 0x4f, 0x49, 0x2c, 0x81,
	0x2a, 0x96, 0x92, 0x84, 0x4b, 0x16, 0x65, 0xa6, 0xa3, 0x48, 0x29, 0x1d, 0x63, 0xe4, 0xe2, 0xf4,
	0x73, 0x0b, 0x09, 0x06, 0x9b, 0x2d, 0x24, 0xc4, 0xc5, 0x92, 0x9c, 0x9f, 0x92, 0x2a, 0xc1, 0xa8,
	0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x66, 0x83, 0xc4, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x98, 0x20, 0x62,
	0x20, 0xb6, 0x90, 0x08, 0x17, 0x6b, 0x7e, 0x79, 0x5e, 0x6a, 0x91, 0x04, 0x33, 0x58, 0x10, 0xc2,
	0x11, 0xb2, 0xe6, 0xe2, 0x46, 0xb2, 0x40, 0x82, 0x45, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x4a, 0x0f,
	0xd5, 0xa5, 0x7a, 0xfe, 0x60, 0x25, 0x2e, 0x89, 0x25, 0x89, 0x41, 0x5c, 0xf9, 0x70, 0xb6, 0x90,
	0x1d, 0x17, 0x4f, 0x7e, 0x1e, 0xc2, 0xe5, 0x12, 0xac, 0x60, 0xdd, 0xd2, 0x18, 0xba, 0xf3, 0x9c,
	0x41, 0x6a, 0xc0, 0xda, 0xb9, 0xa1, 0x1a, 0x40, 0x1c, 0x27, 0x83, 0x13, 0x8f, 0xe4, 0x18, 0x2f,
	0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18,
	0x6e, 0x3c, 0x96, 0x63, 0x88, 0x12, 0x83, 0x18, 0xa1, 0x5f, 0xa1, 0x0f, 0x0b, 0x86, 0x92, 0xca,
	0x82, 0xd4, 0xe2, 0x24, 0x36, 0x70, 0x08, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x75, 0x89,
	0xed, 0xfe, 0x65, 0x01, 0x00, 0x00,
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
		dAtA[i] = 0x2a
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
	if m.OriginData != nil {
		l = m.OriginData.Size()
		n += 1 + l + sovNftSchema(uint64(l))
	}
	if m.OnchainData != nil {
		l = m.OnchainData.Size()
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
		case 5:
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
