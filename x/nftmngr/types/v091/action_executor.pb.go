// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftmngr/v091/action_executor.proto

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

type ActionExecutor struct {
	NftSchemaCode   string `protobuf:"bytes,1,opt,name=nftSchemaCode,proto3" json:"nftSchemaCode,omitempty"`
	ExecutorAddress string `protobuf:"bytes,2,opt,name=executorAddress,proto3" json:"executorAddress,omitempty"`
	Creator         string `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *ActionExecutor) Reset()         { *m = ActionExecutor{} }
func (m *ActionExecutor) String() string { return proto.CompactTextString(m) }
func (*ActionExecutor) ProtoMessage()    {}
func (*ActionExecutor) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3eef538ad390545, []int{0}
}
func (m *ActionExecutor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ActionExecutor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ActionExecutor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ActionExecutor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionExecutor.Merge(m, src)
}
func (m *ActionExecutor) XXX_Size() int {
	return m.Size()
}
func (m *ActionExecutor) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionExecutor.DiscardUnknown(m)
}

var xxx_messageInfo_ActionExecutor proto.InternalMessageInfo

func (m *ActionExecutor) GetNftSchemaCode() string {
	if m != nil {
		return m.NftSchemaCode
	}
	return ""
}

func (m *ActionExecutor) GetExecutorAddress() string {
	if m != nil {
		return m.ExecutorAddress
	}
	return ""
}

func (m *ActionExecutor) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

type ActionExecutorBySchema struct {
	NftSchemaCode   string   `protobuf:"bytes,1,opt,name=nftSchemaCode,proto3" json:"nftSchemaCode,omitempty"`
	ExecutorAddress []string `protobuf:"bytes,2,rep,name=executorAddress,proto3" json:"executorAddress,omitempty"`
}

func (m *ActionExecutorBySchema) Reset()         { *m = ActionExecutorBySchema{} }
func (m *ActionExecutorBySchema) String() string { return proto.CompactTextString(m) }
func (*ActionExecutorBySchema) ProtoMessage()    {}
func (*ActionExecutorBySchema) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3eef538ad390545, []int{1}
}
func (m *ActionExecutorBySchema) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ActionExecutorBySchema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ActionExecutorBySchema.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ActionExecutorBySchema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionExecutorBySchema.Merge(m, src)
}
func (m *ActionExecutorBySchema) XXX_Size() int {
	return m.Size()
}
func (m *ActionExecutorBySchema) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionExecutorBySchema.DiscardUnknown(m)
}

var xxx_messageInfo_ActionExecutorBySchema proto.InternalMessageInfo

func (m *ActionExecutorBySchema) GetNftSchemaCode() string {
	if m != nil {
		return m.NftSchemaCode
	}
	return ""
}

func (m *ActionExecutorBySchema) GetExecutorAddress() []string {
	if m != nil {
		return m.ExecutorAddress
	}
	return nil
}

func init() {
	proto.RegisterType((*ActionExecutor)(nil), "thesixnetwork.sixnft.nftmngr.v091.ActionExecutor")
	proto.RegisterType((*ActionExecutorBySchema)(nil), "thesixnetwork.sixnft.nftmngr.v091.ActionExecutorBySchema")
}

func init() {
	proto.RegisterFile("nftmngr/v091/action_executor.proto", fileDescriptor_c3eef538ad390545)
}

var fileDescriptor_c3eef538ad390545 = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xca, 0x4b, 0x2b, 0xc9,
	0xcd, 0x4b, 0x2f, 0xd2, 0x2f, 0x33, 0xb0, 0x34, 0xd4, 0x4f, 0x4c, 0x2e, 0xc9, 0xcc, 0xcf, 0x8b,
	0x4f, 0xad, 0x48, 0x4d, 0x2e, 0x2d, 0xc9, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52,
	0x2c, 0xc9, 0x48, 0x2d, 0xce, 0xac, 0xc8, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x03, 0x31,
	0xd3, 0x4a, 0xf4, 0xa0, 0x1a, 0xf5, 0x40, 0x1a, 0x95, 0xaa, 0xb8, 0xf8, 0x1c, 0xc1, 0x7a, 0x5d,
	0xa1, 0x5a, 0x85, 0x54, 0xb8, 0x78, 0xf3, 0xd2, 0x4a, 0x82, 0x93, 0x33, 0x52, 0x73, 0x13, 0x9d,
	0xf3, 0x53, 0x52, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x50, 0x05, 0x85, 0x34, 0xb8, 0xf8,
	0x61, 0x96, 0x39, 0xa6, 0xa4, 0x14, 0xa5, 0x16, 0x17, 0x4b, 0x30, 0x81, 0xd5, 0xa1, 0x0b, 0x0b,
	0x49, 0x70, 0xb1, 0x27, 0x17, 0xa5, 0x26, 0x96, 0xe4, 0x17, 0x49, 0x30, 0x83, 0x55, 0xc0, 0xb8,
	0x4a, 0x19, 0x5c, 0x62, 0xa8, 0x76, 0x3b, 0x55, 0x42, 0x6c, 0xa0, 0xc4, 0x0d, 0xcc, 0x58, 0xdc,
	0xe0, 0xe4, 0x77, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e,
	0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x26, 0xe9, 0x99,
	0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x28, 0xa1, 0xa5, 0x0f, 0x09, 0x2d, 0xfd,
	0x0a, 0x7d, 0x58, 0x40, 0x97, 0x54, 0x16, 0xa4, 0x16, 0x83, 0x83, 0x3b, 0x89, 0x0d, 0x1c, 0xbe,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x83, 0x04, 0x6e, 0xa1, 0x85, 0x01, 0x00, 0x00,
}

func (m *ActionExecutor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ActionExecutor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ActionExecutor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintActionExecutor(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ExecutorAddress) > 0 {
		i -= len(m.ExecutorAddress)
		copy(dAtA[i:], m.ExecutorAddress)
		i = encodeVarintActionExecutor(dAtA, i, uint64(len(m.ExecutorAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.NftSchemaCode) > 0 {
		i -= len(m.NftSchemaCode)
		copy(dAtA[i:], m.NftSchemaCode)
		i = encodeVarintActionExecutor(dAtA, i, uint64(len(m.NftSchemaCode)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ActionExecutorBySchema) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ActionExecutorBySchema) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ActionExecutorBySchema) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ExecutorAddress) > 0 {
		for iNdEx := len(m.ExecutorAddress) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ExecutorAddress[iNdEx])
			copy(dAtA[i:], m.ExecutorAddress[iNdEx])
			i = encodeVarintActionExecutor(dAtA, i, uint64(len(m.ExecutorAddress[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.NftSchemaCode) > 0 {
		i -= len(m.NftSchemaCode)
		copy(dAtA[i:], m.NftSchemaCode)
		i = encodeVarintActionExecutor(dAtA, i, uint64(len(m.NftSchemaCode)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintActionExecutor(dAtA []byte, offset int, v uint64) int {
	offset -= sovActionExecutor(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ActionExecutor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NftSchemaCode)
	if l > 0 {
		n += 1 + l + sovActionExecutor(uint64(l))
	}
	l = len(m.ExecutorAddress)
	if l > 0 {
		n += 1 + l + sovActionExecutor(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovActionExecutor(uint64(l))
	}
	return n
}

func (m *ActionExecutorBySchema) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NftSchemaCode)
	if l > 0 {
		n += 1 + l + sovActionExecutor(uint64(l))
	}
	if len(m.ExecutorAddress) > 0 {
		for _, s := range m.ExecutorAddress {
			l = len(s)
			n += 1 + l + sovActionExecutor(uint64(l))
		}
	}
	return n
}

func sovActionExecutor(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozActionExecutor(x uint64) (n int) {
	return sovActionExecutor(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ActionExecutor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowActionExecutor
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
			return fmt.Errorf("proto: ActionExecutor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ActionExecutor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftSchemaCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionExecutor
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
				return ErrInvalidLengthActionExecutor
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActionExecutor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NftSchemaCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecutorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionExecutor
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
				return ErrInvalidLengthActionExecutor
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActionExecutor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExecutorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionExecutor
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
				return ErrInvalidLengthActionExecutor
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActionExecutor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipActionExecutor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthActionExecutor
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
func (m *ActionExecutorBySchema) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowActionExecutor
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
			return fmt.Errorf("proto: ActionExecutorBySchema: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ActionExecutorBySchema: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftSchemaCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionExecutor
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
				return ErrInvalidLengthActionExecutor
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActionExecutor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NftSchemaCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecutorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionExecutor
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
				return ErrInvalidLengthActionExecutor
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActionExecutor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExecutorAddress = append(m.ExecutorAddress, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipActionExecutor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthActionExecutor
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
func skipActionExecutor(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowActionExecutor
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
					return 0, ErrIntOverflowActionExecutor
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
					return 0, ErrIntOverflowActionExecutor
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
				return 0, ErrInvalidLengthActionExecutor
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupActionExecutor
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthActionExecutor
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthActionExecutor        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowActionExecutor          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupActionExecutor = fmt.Errorf("proto: unexpected end of group")
)
