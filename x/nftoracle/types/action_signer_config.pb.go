// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftoracle/action_signer_config.proto

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

type ActionSignerConfig struct {
	Chain           string `protobuf:"bytes,1,opt,name=chain,proto3" json:"chain,omitempty"`
	RpcEndpoint     string `protobuf:"bytes,2,opt,name=rpc_endpoint,json=rpcEndpoint,proto3" json:"rpc_endpoint,omitempty"`
	ContractAddress string `protobuf:"bytes,3,opt,name=contractAddress,proto3" json:"contractAddress,omitempty"`
	ContractName    string `protobuf:"bytes,4,opt,name=contractName,proto3" json:"contractName,omitempty"`
	ContractOwner   string `protobuf:"bytes,5,opt,name=contractOwner,proto3" json:"contractOwner,omitempty"`
	Creator         string `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *ActionSignerConfig) Reset()         { *m = ActionSignerConfig{} }
func (m *ActionSignerConfig) String() string { return proto.CompactTextString(m) }
func (*ActionSignerConfig) ProtoMessage()    {}
func (*ActionSignerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f162be7e277ee1b, []int{0}
}
func (m *ActionSignerConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ActionSignerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ActionSignerConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ActionSignerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionSignerConfig.Merge(m, src)
}
func (m *ActionSignerConfig) XXX_Size() int {
	return m.Size()
}
func (m *ActionSignerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionSignerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ActionSignerConfig proto.InternalMessageInfo

func (m *ActionSignerConfig) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *ActionSignerConfig) GetRpcEndpoint() string {
	if m != nil {
		return m.RpcEndpoint
	}
	return ""
}

func (m *ActionSignerConfig) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *ActionSignerConfig) GetContractName() string {
	if m != nil {
		return m.ContractName
	}
	return ""
}

func (m *ActionSignerConfig) GetContractOwner() string {
	if m != nil {
		return m.ContractOwner
	}
	return ""
}

func (m *ActionSignerConfig) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*ActionSignerConfig)(nil), "thesixnetwork.sixnft.nftoracle.ActionSignerConfig")
}

func init() {
	proto.RegisterFile("nftoracle/action_signer_config.proto", fileDescriptor_4f162be7e277ee1b)
}

var fileDescriptor_4f162be7e277ee1b = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xb3, 0x6a, 0x2b, 0xae, 0x15, 0x61, 0xf1, 0xb0, 0xa7, 0x45, 0x4b, 0x0f, 0x3d, 0x25,
	0x88, 0x4f, 0x50, 0xc5, 0x93, 0xa0, 0xa0, 0x37, 0x2f, 0x61, 0xbb, 0xd9, 0x24, 0x8b, 0x76, 0x27,
	0x4c, 0x46, 0x5a, 0xdf, 0xc2, 0xc7, 0xf2, 0xd8, 0xa3, 0xe0, 0x45, 0x92, 0x17, 0x11, 0xb7, 0x46,
	0x49, 0x6f, 0xf3, 0x7f, 0xf3, 0xff, 0x97, 0x8f, 0x4f, 0x7c, 0x4e, 0x80, 0xda, 0x3c, 0xdb, 0x44,
	0x1b, 0x72, 0xe0, 0xd3, 0xda, 0x15, 0xde, 0x62, 0x6a, 0xc0, 0xe7, 0xae, 0x88, 0x2b, 0x04, 0x02,
	0xa1, 0xa8, 0xb4, 0xb5, 0x5b, 0x79, 0x4b, 0x4b, 0xc0, 0xa7, 0xf8, 0xe7, 0xcc, 0x29, 0xfe, 0x9b,
	0x8e, 0x3f, 0x19, 0x17, 0xb3, 0x30, 0x7f, 0x08, 0xeb, 0xab, 0x30, 0x16, 0x27, 0x7c, 0x60, 0x4a,
	0xed, 0xbc, 0x64, 0xa7, 0x6c, 0x7a, 0x70, 0xbf, 0x09, 0xe2, 0x8c, 0x8f, 0xb0, 0x32, 0xa9, 0xf5,
	0x59, 0x05, 0xce, 0x93, 0xdc, 0x09, 0xcf, 0x43, 0xac, 0xcc, 0xf5, 0x2f, 0x12, 0x53, 0x7e, 0x6c,
	0xc0, 0x13, 0x6a, 0x43, 0xb3, 0x2c, 0x43, 0x5b, 0xd7, 0x72, 0x37, 0xb4, 0xb6, 0xb1, 0x18, 0xf3,
	0x51, 0x87, 0x6e, 0xf5, 0xc2, 0xca, 0xbd, 0x50, 0xeb, 0x31, 0x31, 0xe1, 0x47, 0x5d, 0xbe, 0x5b,
	0x7a, 0x8b, 0x72, 0x10, 0x4a, 0x7d, 0x28, 0x24, 0xdf, 0x37, 0x68, 0x35, 0x01, 0xca, 0x61, 0xf8,
	0x77, 0xf1, 0xf2, 0xe6, 0xbd, 0x51, 0x6c, 0xdd, 0x28, 0xf6, 0xd5, 0x28, 0xf6, 0xd6, 0xaa, 0x68,
	0xdd, 0xaa, 0xe8, 0xa3, 0x55, 0xd1, 0xe3, 0x79, 0xe1, 0xa8, 0x7c, 0x99, 0xc7, 0x06, 0x16, 0x49,
	0x4f, 0x51, 0xb2, 0x51, 0x94, 0xac, 0x92, 0x7f, 0xbf, 0xf4, 0x5a, 0xd9, 0x7a, 0x3e, 0x0c, 0x46,
	0x2f, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0xb7, 0x0c, 0x1e, 0xe3, 0x79, 0x01, 0x00, 0x00,
}

func (m *ActionSignerConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ActionSignerConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ActionSignerConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintActionSignerConfig(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.ContractOwner) > 0 {
		i -= len(m.ContractOwner)
		copy(dAtA[i:], m.ContractOwner)
		i = encodeVarintActionSignerConfig(dAtA, i, uint64(len(m.ContractOwner)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ContractName) > 0 {
		i -= len(m.ContractName)
		copy(dAtA[i:], m.ContractName)
		i = encodeVarintActionSignerConfig(dAtA, i, uint64(len(m.ContractName)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintActionSignerConfig(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.RpcEndpoint) > 0 {
		i -= len(m.RpcEndpoint)
		copy(dAtA[i:], m.RpcEndpoint)
		i = encodeVarintActionSignerConfig(dAtA, i, uint64(len(m.RpcEndpoint)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintActionSignerConfig(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintActionSignerConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovActionSignerConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ActionSignerConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovActionSignerConfig(uint64(l))
	}
	l = len(m.RpcEndpoint)
	if l > 0 {
		n += 1 + l + sovActionSignerConfig(uint64(l))
	}
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovActionSignerConfig(uint64(l))
	}
	l = len(m.ContractName)
	if l > 0 {
		n += 1 + l + sovActionSignerConfig(uint64(l))
	}
	l = len(m.ContractOwner)
	if l > 0 {
		n += 1 + l + sovActionSignerConfig(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovActionSignerConfig(uint64(l))
	}
	return n
}

func sovActionSignerConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozActionSignerConfig(x uint64) (n int) {
	return sovActionSignerConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ActionSignerConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowActionSignerConfig
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
			return fmt.Errorf("proto: ActionSignerConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ActionSignerConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionSignerConfig
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
				return ErrInvalidLengthActionSignerConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActionSignerConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RpcEndpoint", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionSignerConfig
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
				return ErrInvalidLengthActionSignerConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActionSignerConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RpcEndpoint = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionSignerConfig
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
				return ErrInvalidLengthActionSignerConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActionSignerConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionSignerConfig
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
				return ErrInvalidLengthActionSignerConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActionSignerConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractOwner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionSignerConfig
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
				return ErrInvalidLengthActionSignerConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActionSignerConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractOwner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionSignerConfig
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
				return ErrInvalidLengthActionSignerConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActionSignerConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipActionSignerConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthActionSignerConfig
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
func skipActionSignerConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowActionSignerConfig
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
					return 0, ErrIntOverflowActionSignerConfig
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
					return 0, ErrIntOverflowActionSignerConfig
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
				return 0, ErrInvalidLengthActionSignerConfig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupActionSignerConfig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthActionSignerConfig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthActionSignerConfig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowActionSignerConfig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupActionSignerConfig = fmt.Errorf("proto: unexpected end of group")
)
