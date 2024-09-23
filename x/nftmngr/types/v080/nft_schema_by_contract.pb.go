// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftmngr/v080/nft_schema_by_contract.proto

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

type NFTSchemaByContract struct {
	OriginContractAddress string   `protobuf:"bytes,1,opt,name=originContractAddress,proto3" json:"originContractAddress,omitempty"`
	SchemaCodes           []string `protobuf:"bytes,2,rep,name=schemaCodes,proto3" json:"schemaCodes,omitempty"`
}

func (m *NFTSchemaByContract) Reset()         { *m = NFTSchemaByContract{} }
func (m *NFTSchemaByContract) String() string { return proto.CompactTextString(m) }
func (*NFTSchemaByContract) ProtoMessage()    {}
func (*NFTSchemaByContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_eef9fad5da3f22bb, []int{0}
}
func (m *NFTSchemaByContract) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NFTSchemaByContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NFTSchemaByContract.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NFTSchemaByContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NFTSchemaByContract.Merge(m, src)
}
func (m *NFTSchemaByContract) XXX_Size() int {
	return m.Size()
}
func (m *NFTSchemaByContract) XXX_DiscardUnknown() {
	xxx_messageInfo_NFTSchemaByContract.DiscardUnknown(m)
}

var xxx_messageInfo_NFTSchemaByContract proto.InternalMessageInfo

func (m *NFTSchemaByContract) GetOriginContractAddress() string {
	if m != nil {
		return m.OriginContractAddress
	}
	return ""
}

func (m *NFTSchemaByContract) GetSchemaCodes() []string {
	if m != nil {
		return m.SchemaCodes
	}
	return nil
}

func init() {
	proto.RegisterType((*NFTSchemaByContract)(nil), "thesixnetwork.sixnft.nftmngr.v080.NFTSchemaByContract")
}

func init() {
	proto.RegisterFile("nftmngr/v080/nft_schema_by_contract.proto", fileDescriptor_eef9fad5da3f22bb)
}

var fileDescriptor_eef9fad5da3f22bb = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xcc, 0x4b, 0x2b, 0xc9,
	0xcd, 0x4b, 0x2f, 0xd2, 0x2f, 0x33, 0xb0, 0x30, 0xd0, 0xcf, 0x4b, 0x2b, 0x89, 0x2f, 0x4e, 0xce,
	0x48, 0xcd, 0x4d, 0x8c, 0x4f, 0xaa, 0x8c, 0x4f, 0xce, 0xcf, 0x2b, 0x29, 0x4a, 0x4c, 0x2e, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x2c, 0xc9, 0x48, 0x2d, 0xce, 0xac, 0xc8, 0x4b, 0x2d,
	0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x03, 0x31, 0xd3, 0x4a, 0xf4, 0xa0, 0xfa, 0xf5, 0x40, 0xfa, 0x95,
	0x72, 0xb9, 0x84, 0xfd, 0xdc, 0x42, 0x82, 0xc1, 0x26, 0x38, 0x55, 0x3a, 0x43, 0xf5, 0x0b, 0x99,
	0x70, 0x89, 0xe6, 0x17, 0x65, 0xa6, 0x67, 0xe6, 0xc1, 0x44, 0x1c, 0x53, 0x52, 0x8a, 0x52, 0x8b,
	0x8b, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xb0, 0x4b, 0x0a, 0x29, 0x70, 0x71, 0x43, 0xdc,
	0xe2, 0x9c, 0x9f, 0x92, 0x5a, 0x2c, 0xc1, 0xa4, 0xc0, 0xac, 0xc1, 0x19, 0x84, 0x2c, 0xe4, 0xe4,
	0x77, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c,
	0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x26, 0xe9, 0x99, 0x25, 0x19,
	0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x28, 0xce, 0xd6, 0x87, 0x38, 0x5b, 0xbf, 0x42, 0x1f,
	0xe6, 0xf1, 0x92, 0xca, 0x82, 0xd4, 0x62, 0xb0, 0xf7, 0x93, 0xd8, 0xc0, 0x1e, 0x35, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x3a, 0xf5, 0x94, 0xd1, 0x15, 0x01, 0x00, 0x00,
}

func (m *NFTSchemaByContract) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NFTSchemaByContract) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NFTSchemaByContract) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SchemaCodes) > 0 {
		for iNdEx := len(m.SchemaCodes) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.SchemaCodes[iNdEx])
			copy(dAtA[i:], m.SchemaCodes[iNdEx])
			i = encodeVarintNftSchemaByContract(dAtA, i, uint64(len(m.SchemaCodes[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.OriginContractAddress) > 0 {
		i -= len(m.OriginContractAddress)
		copy(dAtA[i:], m.OriginContractAddress)
		i = encodeVarintNftSchemaByContract(dAtA, i, uint64(len(m.OriginContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNftSchemaByContract(dAtA []byte, offset int, v uint64) int {
	offset -= sovNftSchemaByContract(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NFTSchemaByContract) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OriginContractAddress)
	if l > 0 {
		n += 1 + l + sovNftSchemaByContract(uint64(l))
	}
	if len(m.SchemaCodes) > 0 {
		for _, s := range m.SchemaCodes {
			l = len(s)
			n += 1 + l + sovNftSchemaByContract(uint64(l))
		}
	}
	return n
}

func sovNftSchemaByContract(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNftSchemaByContract(x uint64) (n int) {
	return sovNftSchemaByContract(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NFTSchemaByContract) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNftSchemaByContract
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
			return fmt.Errorf("proto: NFTSchemaByContract: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NFTSchemaByContract: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftSchemaByContract
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
				return ErrInvalidLengthNftSchemaByContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftSchemaByContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OriginContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SchemaCodes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftSchemaByContract
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
				return ErrInvalidLengthNftSchemaByContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftSchemaByContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SchemaCodes = append(m.SchemaCodes, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNftSchemaByContract(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNftSchemaByContract
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
func skipNftSchemaByContract(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNftSchemaByContract
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
					return 0, ErrIntOverflowNftSchemaByContract
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
					return 0, ErrIntOverflowNftSchemaByContract
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
				return 0, ErrInvalidLengthNftSchemaByContract
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNftSchemaByContract
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNftSchemaByContract
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNftSchemaByContract        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNftSchemaByContract          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNftSchemaByContract = fmt.Errorf("proto: unexpected end of group")
)
