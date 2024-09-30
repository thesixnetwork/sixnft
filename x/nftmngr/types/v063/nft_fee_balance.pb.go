// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftmngr/v063/nft_fee_balance.proto

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

type NFTFeeBalance struct {
	// map<int32, string> fee_balances = 1;
	FeeBalances []string `protobuf:"bytes,1,rep,name=fee_balances,json=feeBalances,proto3" json:"fee_balances,omitempty"`
}

func (m *NFTFeeBalance) Reset()         { *m = NFTFeeBalance{} }
func (m *NFTFeeBalance) String() string { return proto.CompactTextString(m) }
func (*NFTFeeBalance) ProtoMessage()    {}
func (*NFTFeeBalance) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d8c071a73f52019, []int{0}
}
func (m *NFTFeeBalance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NFTFeeBalance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NFTFeeBalance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NFTFeeBalance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NFTFeeBalance.Merge(m, src)
}
func (m *NFTFeeBalance) XXX_Size() int {
	return m.Size()
}
func (m *NFTFeeBalance) XXX_DiscardUnknown() {
	xxx_messageInfo_NFTFeeBalance.DiscardUnknown(m)
}

var xxx_messageInfo_NFTFeeBalance proto.InternalMessageInfo

func (m *NFTFeeBalance) GetFeeBalances() []string {
	if m != nil {
		return m.FeeBalances
	}
	return nil
}

func init() {
	proto.RegisterType((*NFTFeeBalance)(nil), "thesixnetwork.sixnft.nftmngr.v063.NFTFeeBalance")
}

func init() {
	proto.RegisterFile("nftmngr/v063/nft_fee_balance.proto", fileDescriptor_3d8c071a73f52019)
}

var fileDescriptor_3d8c071a73f52019 = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xca, 0x4b, 0x2b, 0xc9,
	0xcd, 0x4b, 0x2f, 0xd2, 0x2f, 0x33, 0x30, 0x33, 0xd6, 0xcf, 0x4b, 0x2b, 0x89, 0x4f, 0x4b, 0x4d,
	0x8d, 0x4f, 0x4a, 0xcc, 0x49, 0xcc, 0x4b, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52,
	0x2c, 0xc9, 0x48, 0x2d, 0xce, 0xac, 0xc8, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x03, 0x31,
	0xd3, 0x4a, 0xf4, 0xa0, 0x1a, 0xf5, 0x40, 0x1a, 0x95, 0x8c, 0xb8, 0x78, 0xfd, 0xdc, 0x42, 0xdc,
	0x52, 0x53, 0x9d, 0x20, 0x3a, 0x85, 0x14, 0xb9, 0x78, 0x90, 0x0c, 0x2a, 0x96, 0x60, 0x54, 0x60,
	0xd6, 0xe0, 0x0c, 0xe2, 0x4e, 0x83, 0xab, 0x28, 0x76, 0xf2, 0x3b, 0xf1, 0x48, 0x8e, 0xf1, 0xc2,
	0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1,
	0xc6, 0x63, 0x39, 0x86, 0x28, 0x93, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c,
	0x7d, 0x14, 0xbb, 0xf5, 0x21, 0x76, 0xeb, 0x57, 0xe8, 0xc3, 0x9c, 0x5d, 0x52, 0x59, 0x90, 0x5a,
	0x0c, 0x76, 0x7c, 0x12, 0x1b, 0xd8, 0xb5, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x87, 0x94,
	0x66, 0x7e, 0xd3, 0x00, 0x00, 0x00,
}

func (m *NFTFeeBalance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NFTFeeBalance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NFTFeeBalance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FeeBalances) > 0 {
		for iNdEx := len(m.FeeBalances) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.FeeBalances[iNdEx])
			copy(dAtA[i:], m.FeeBalances[iNdEx])
			i = encodeVarintNftFeeBalance(dAtA, i, uint64(len(m.FeeBalances[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintNftFeeBalance(dAtA []byte, offset int, v uint64) int {
	offset -= sovNftFeeBalance(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NFTFeeBalance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.FeeBalances) > 0 {
		for _, s := range m.FeeBalances {
			l = len(s)
			n += 1 + l + sovNftFeeBalance(uint64(l))
		}
	}
	return n
}

func sovNftFeeBalance(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNftFeeBalance(x uint64) (n int) {
	return sovNftFeeBalance(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NFTFeeBalance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNftFeeBalance
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
			return fmt.Errorf("proto: NFTFeeBalance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NFTFeeBalance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeBalances", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftFeeBalance
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
				return ErrInvalidLengthNftFeeBalance
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftFeeBalance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeBalances = append(m.FeeBalances, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNftFeeBalance(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNftFeeBalance
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
func skipNftFeeBalance(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNftFeeBalance
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
					return 0, ErrIntOverflowNftFeeBalance
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
					return 0, ErrIntOverflowNftFeeBalance
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
				return 0, ErrInvalidLengthNftFeeBalance
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNftFeeBalance
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNftFeeBalance
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNftFeeBalance        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNftFeeBalance          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNftFeeBalance = fmt.Errorf("proto: unexpected end of group")
)