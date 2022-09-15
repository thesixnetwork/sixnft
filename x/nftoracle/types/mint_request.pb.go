// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftoracle/mint_request.proto

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

type RequestStatus int32

const (
	RequestStatus_PENDING                  RequestStatus = 0
	RequestStatus_SUCCESS_WITH_CONSENSUS   RequestStatus = 1
	RequestStatus_FAILED_WITHOUT_CONCENSUS RequestStatus = 2
)

var RequestStatus_name = map[int32]string{
	0: "PENDING",
	1: "SUCCESS_WITH_CONSENSUS",
	2: "FAILED_WITHOUT_CONCENSUS",
}

var RequestStatus_value = map[string]int32{
	"PENDING":                  0,
	"SUCCESS_WITH_CONSENSUS":   1,
	"FAILED_WITHOUT_CONCENSUS": 2,
}

func (x RequestStatus) String() string {
	return proto.EnumName(RequestStatus_name, int32(x))
}

func (RequestStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3b5fc42e48c76b2e, []int{0}
}

type MintRequest struct {
	Id              uint64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	NftSchemaCode   string        `protobuf:"bytes,2,opt,name=nftSchemaCode,proto3" json:"nftSchemaCode,omitempty"`
	TokenId         string        `protobuf:"bytes,3,opt,name=tokenId,proto3" json:"tokenId,omitempty"`
	RequiredConfirm uint64        `protobuf:"varint,4,opt,name=requiredConfirm,proto3" json:"requiredConfirm,omitempty"`
	Status          RequestStatus `protobuf:"varint,5,opt,name=status,proto3,enum=sixnft.nftoracle.RequestStatus" json:"status,omitempty"`
}

func (m *MintRequest) Reset()         { *m = MintRequest{} }
func (m *MintRequest) String() string { return proto.CompactTextString(m) }
func (*MintRequest) ProtoMessage()    {}
func (*MintRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b5fc42e48c76b2e, []int{0}
}
func (m *MintRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MintRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MintRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MintRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MintRequest.Merge(m, src)
}
func (m *MintRequest) XXX_Size() int {
	return m.Size()
}
func (m *MintRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MintRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MintRequest proto.InternalMessageInfo

func (m *MintRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MintRequest) GetNftSchemaCode() string {
	if m != nil {
		return m.NftSchemaCode
	}
	return ""
}

func (m *MintRequest) GetTokenId() string {
	if m != nil {
		return m.TokenId
	}
	return ""
}

func (m *MintRequest) GetRequiredConfirm() uint64 {
	if m != nil {
		return m.RequiredConfirm
	}
	return 0
}

func (m *MintRequest) GetStatus() RequestStatus {
	if m != nil {
		return m.Status
	}
	return RequestStatus_PENDING
}

func init() {
	proto.RegisterEnum("sixnft.nftoracle.RequestStatus", RequestStatus_name, RequestStatus_value)
	proto.RegisterType((*MintRequest)(nil), "sixnft.nftoracle.MintRequest")
}

func init() { proto.RegisterFile("nftoracle/mint_request.proto", fileDescriptor_3b5fc42e48c76b2e) }

var fileDescriptor_3b5fc42e48c76b2e = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xc1, 0x4a, 0xc3, 0x30,
	0x1c, 0x87, 0x9b, 0x3a, 0x37, 0xcc, 0xd8, 0x1c, 0x39, 0x48, 0x90, 0x11, 0x87, 0x78, 0x28, 0x1e,
	0x3a, 0x98, 0x07, 0xcf, 0x9a, 0x55, 0x2d, 0x68, 0x27, 0xcd, 0xa6, 0xe0, 0x65, 0xd4, 0x35, 0xc5,
	0xa0, 0x4d, 0x66, 0x9b, 0xc1, 0x7c, 0x0b, 0x1f, 0xc9, 0xa3, 0xc7, 0x1d, 0x3d, 0x4a, 0xfb, 0x22,
	0xb2, 0x76, 0x2a, 0xdb, 0x31, 0xff, 0xef, 0x23, 0x7c, 0xfc, 0x60, 0x5b, 0x46, 0x5a, 0x25, 0xc1,
	0xe4, 0x85, 0x77, 0x63, 0x21, 0xf5, 0x38, 0xe1, 0xaf, 0x33, 0x9e, 0x6a, 0x7b, 0x9a, 0x28, 0xad,
	0x50, 0x2b, 0x15, 0x73, 0x19, 0x69, 0xfb, 0x4f, 0x3a, 0xfc, 0x00, 0xb0, 0x7e, 0x23, 0xa4, 0xf6,
	0x4b, 0x0f, 0x35, 0xa1, 0x29, 0x42, 0x0c, 0x3a, 0xc0, 0xaa, 0xf8, 0xa6, 0x08, 0xd1, 0x11, 0x6c,
	0xc8, 0x48, 0xb3, 0xc9, 0x13, 0x8f, 0x03, 0xaa, 0x42, 0x8e, 0xcd, 0x0e, 0xb0, 0x76, 0xfc, 0xf5,
	0x23, 0xc2, 0xb0, 0xa6, 0xd5, 0x33, 0x97, 0x6e, 0x88, 0xb7, 0x0a, 0xfe, 0xfb, 0x44, 0x16, 0xdc,
	0x5d, 0x26, 0x88, 0x84, 0x87, 0x54, 0xc9, 0x48, 0x24, 0x31, 0xae, 0x14, 0x9f, 0x6f, 0x9e, 0xd1,
	0x29, 0xac, 0xa6, 0x3a, 0xd0, 0xb3, 0x14, 0x6f, 0x77, 0x80, 0xd5, 0xec, 0x1d, 0xd8, 0x9b, 0xb1,
	0xf6, 0x2a, 0x92, 0x15, 0x9a, 0xbf, 0xd2, 0x8f, 0xef, 0x60, 0x63, 0x0d, 0xa0, 0x3a, 0xac, 0xdd,
	0x3a, 0x5e, 0xdf, 0xf5, 0x2e, 0x5b, 0x06, 0xda, 0x87, 0x7b, 0x6c, 0x44, 0xa9, 0xc3, 0xd8, 0xf8,
	0xde, 0x1d, 0x5e, 0x8d, 0xe9, 0xc0, 0x63, 0x8e, 0xc7, 0x46, 0xac, 0x05, 0x50, 0x1b, 0xe2, 0x8b,
	0x33, 0xf7, 0xda, 0xe9, 0x17, 0x68, 0x30, 0x1a, 0x2e, 0x29, 0x2d, 0xa9, 0x79, 0xde, 0xfb, 0xcc,
	0x08, 0x58, 0x64, 0x04, 0x7c, 0x67, 0x04, 0xbc, 0xe7, 0xc4, 0x58, 0xe4, 0xc4, 0xf8, 0xca, 0x89,
	0xf1, 0x80, 0xcb, 0xb2, 0xee, 0xbc, 0xfb, 0xbf, 0xb6, 0x7e, 0x9b, 0xf2, 0xf4, 0xb1, 0x5a, 0xec,
	0x7c, 0xf2, 0x13, 0x00, 0x00, 0xff, 0xff, 0x92, 0x58, 0x17, 0x3f, 0x87, 0x01, 0x00, 0x00,
}

func (m *MintRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MintRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MintRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintMintRequest(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x28
	}
	if m.RequiredConfirm != 0 {
		i = encodeVarintMintRequest(dAtA, i, uint64(m.RequiredConfirm))
		i--
		dAtA[i] = 0x20
	}
	if len(m.TokenId) > 0 {
		i -= len(m.TokenId)
		copy(dAtA[i:], m.TokenId)
		i = encodeVarintMintRequest(dAtA, i, uint64(len(m.TokenId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.NftSchemaCode) > 0 {
		i -= len(m.NftSchemaCode)
		copy(dAtA[i:], m.NftSchemaCode)
		i = encodeVarintMintRequest(dAtA, i, uint64(len(m.NftSchemaCode)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintMintRequest(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMintRequest(dAtA []byte, offset int, v uint64) int {
	offset -= sovMintRequest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MintRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovMintRequest(uint64(m.Id))
	}
	l = len(m.NftSchemaCode)
	if l > 0 {
		n += 1 + l + sovMintRequest(uint64(l))
	}
	l = len(m.TokenId)
	if l > 0 {
		n += 1 + l + sovMintRequest(uint64(l))
	}
	if m.RequiredConfirm != 0 {
		n += 1 + sovMintRequest(uint64(m.RequiredConfirm))
	}
	if m.Status != 0 {
		n += 1 + sovMintRequest(uint64(m.Status))
	}
	return n
}

func sovMintRequest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMintRequest(x uint64) (n int) {
	return sovMintRequest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MintRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMintRequest
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
			return fmt.Errorf("proto: MintRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MintRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMintRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftSchemaCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMintRequest
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
				return ErrInvalidLengthMintRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMintRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NftSchemaCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMintRequest
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
				return ErrInvalidLengthMintRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMintRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequiredConfirm", wireType)
			}
			m.RequiredConfirm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMintRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequiredConfirm |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMintRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= RequestStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMintRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMintRequest
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
func skipMintRequest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMintRequest
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
					return 0, ErrIntOverflowMintRequest
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
					return 0, ErrIntOverflowMintRequest
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
				return 0, ErrInvalidLengthMintRequest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMintRequest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMintRequest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMintRequest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMintRequest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMintRequest = fmt.Errorf("proto: unexpected end of group")
)
