// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftmngr/v063/action.proto

package v063

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type AllowedActioner int32

const (
	AllowedActioner_ALLOWED_ACTIONER_ALL         AllowedActioner = 0
	AllowedActioner_ALLOWED_ACTIONER_SYSTEM_ONLY AllowedActioner = 1
	AllowedActioner_ALLOWED_ACTIONER_USER_ONLY   AllowedActioner = 2
)

var AllowedActioner_name = map[int32]string{
	0: "ALLOWED_ACTIONER_ALL",
	1: "ALLOWED_ACTIONER_SYSTEM_ONLY",
	2: "ALLOWED_ACTIONER_USER_ONLY",
}

var AllowedActioner_value = map[string]int32{
	"ALLOWED_ACTIONER_ALL":         0,
	"ALLOWED_ACTIONER_SYSTEM_ONLY": 1,
	"ALLOWED_ACTIONER_USER_ONLY":   2,
}

func (x AllowedActioner) String() string {
	return proto.EnumName(AllowedActioner_name, int32(x))
}

func (AllowedActioner) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6723d1598ed5cf62, []int{0}
}

type Action struct {
	Name            string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Desc            string          `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	Disable         bool            `protobuf:"varint,3,opt,name=disable,proto3" json:"disable,omitempty"`
	When            string          `protobuf:"bytes,4,opt,name=when,proto3" json:"when,omitempty"`
	Then            []string        `protobuf:"bytes,5,rep,name=then,proto3" json:"then,omitempty"`
	AllowedActioner AllowedActioner `protobuf:"varint,6,opt,name=allowed_actioner,json=allowedActioner,proto3,enum=thesixnetwork.sixnft.nftmngr.v063.AllowedActioner" json:"allowed_actioner,omitempty"`
}

func (m *Action) Reset()         { *m = Action{} }
func (m *Action) String() string { return proto.CompactTextString(m) }
func (*Action) ProtoMessage()    {}
func (*Action) Descriptor() ([]byte, []int) {
	return fileDescriptor_6723d1598ed5cf62, []int{0}
}
func (m *Action) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Action) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Action.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Action) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Action.Merge(m, src)
}
func (m *Action) XXX_Size() int {
	return m.Size()
}
func (m *Action) XXX_DiscardUnknown() {
	xxx_messageInfo_Action.DiscardUnknown(m)
}

var xxx_messageInfo_Action proto.InternalMessageInfo

func (m *Action) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Action) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *Action) GetDisable() bool {
	if m != nil {
		return m.Disable
	}
	return false
}

func (m *Action) GetWhen() string {
	if m != nil {
		return m.When
	}
	return ""
}

func (m *Action) GetThen() []string {
	if m != nil {
		return m.Then
	}
	return nil
}

func (m *Action) GetAllowedActioner() AllowedActioner {
	if m != nil {
		return m.AllowedActioner
	}
	return AllowedActioner_ALLOWED_ACTIONER_ALL
}

func init() {
	proto.RegisterEnum("thesixnetwork.sixnft.nftmngr.v063.AllowedActioner", AllowedActioner_name, AllowedActioner_value)
	proto.RegisterType((*Action)(nil), "thesixnetwork.sixnft.nftmngr.v063.Action")
}

func init() { proto.RegisterFile("nftmngr/v063/action.proto", fileDescriptor_6723d1598ed5cf62) }

var fileDescriptor_6723d1598ed5cf62 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x41, 0x4f, 0xe2, 0x40,
	0x18, 0x86, 0x3b, 0xc0, 0xb2, 0xbb, 0x73, 0x58, 0xc8, 0x84, 0xc3, 0x2c, 0x31, 0x63, 0xf5, 0x44,
	0x3c, 0xcc, 0x18, 0x30, 0xde, 0x8b, 0xf6, 0x60, 0x52, 0x21, 0x29, 0x18, 0x83, 0x97, 0xa6, 0x94,
	0xa1, 0x34, 0xd2, 0x0e, 0x69, 0x47, 0xc1, 0x7f, 0xe1, 0xcf, 0xe2, 0xc8, 0x4d, 0x4f, 0xc6, 0xc0,
	0x1f, 0x31, 0xd3, 0x81, 0x03, 0x72, 0xf0, 0xf6, 0xce, 0xf3, 0x3d, 0xdf, 0x24, 0x6f, 0x3e, 0xf8,
	0x3f, 0x19, 0xcb, 0x38, 0x09, 0x53, 0xf6, 0x7c, 0x7e, 0xd9, 0x62, 0x7e, 0x20, 0x23, 0x91, 0xd0,
	0x59, 0x2a, 0xa4, 0x40, 0x27, 0x72, 0xc2, 0xb3, 0x68, 0x91, 0x70, 0x39, 0x17, 0xe9, 0x23, 0x55,
	0x71, 0x2c, 0xe9, 0xd6, 0xa7, 0xca, 0xaf, 0xd7, 0x42, 0x11, 0x8a, 0xdc, 0x66, 0x2a, 0xe9, 0xc5,
	0xd3, 0x37, 0x00, 0xcb, 0x56, 0xfe, 0x13, 0x42, 0xb0, 0x94, 0xf8, 0x31, 0xc7, 0xc0, 0x04, 0x8d,
	0xbf, 0x6e, 0x9e, 0x15, 0x1b, 0xf1, 0x2c, 0xc0, 0x05, 0xcd, 0x54, 0x46, 0x18, 0xfe, 0x1e, 0x45,
	0x99, 0x3f, 0x9c, 0x72, 0x5c, 0x34, 0x41, 0xe3, 0x8f, 0xbb, 0x7b, 0x2a, 0x7b, 0x3e, 0xe1, 0x09,
	0x2e, 0x69, 0x5b, 0x65, 0xc5, 0xa4, 0x62, 0xbf, 0xcc, 0xa2, 0x62, 0x2a, 0xa3, 0x00, 0x56, 0xfd,
	0xe9, 0x54, 0xcc, 0xf9, 0xc8, 0xd3, 0x2d, 0x78, 0x8a, 0xcb, 0x26, 0x68, 0xfc, 0x6b, 0x36, 0xe9,
	0x8f, 0x45, 0xa8, 0xa5, 0x57, 0xad, 0xed, 0x66, 0xbb, 0xb4, 0xfc, 0x38, 0x06, 0x6e, 0xc5, 0xdf,
	0xc7, 0x67, 0x31, 0xac, 0x7c, 0x33, 0x11, 0x86, 0x35, 0xcb, 0x71, 0xba, 0xf7, 0xf6, 0xb5, 0x67,
	0x5d, 0xf5, 0x6f, 0xba, 0x1d, 0xdb, 0xf5, 0x2c, 0xc7, 0xa9, 0x1a, 0xc8, 0x84, 0x47, 0x07, 0x93,
	0xde, 0xa0, 0xd7, 0xb7, 0x6f, 0xbd, 0x6e, 0xc7, 0x19, 0x54, 0x01, 0x22, 0xb0, 0x7e, 0x60, 0xdc,
	0xf5, 0x6c, 0x57, 0xcf, 0x0b, 0xed, 0xce, 0x72, 0x4d, 0xc0, 0x6a, 0x4d, 0xc0, 0xe7, 0x9a, 0x80,
	0xd7, 0x0d, 0x31, 0x56, 0x1b, 0x62, 0xbc, 0x6f, 0x88, 0xf1, 0x70, 0x11, 0x46, 0x72, 0xf2, 0x34,
	0xa4, 0x81, 0x88, 0xd9, 0x5e, 0x3b, 0xa6, 0xdb, 0xb1, 0x05, 0xdb, 0x1d, 0x56, 0xbe, 0xcc, 0x78,
	0x96, 0x9f, 0x77, 0x58, 0xce, 0xef, 0xd3, 0xfa, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x74, 0xb7,
	0x37, 0xf5, 0x01, 0x00, 0x00,
}

func (m *Action) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Action) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Action) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AllowedActioner != 0 {
		i = encodeVarintAction(dAtA, i, uint64(m.AllowedActioner))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Then) > 0 {
		for iNdEx := len(m.Then) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Then[iNdEx])
			copy(dAtA[i:], m.Then[iNdEx])
			i = encodeVarintAction(dAtA, i, uint64(len(m.Then[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.When) > 0 {
		i -= len(m.When)
		copy(dAtA[i:], m.When)
		i = encodeVarintAction(dAtA, i, uint64(len(m.When)))
		i--
		dAtA[i] = 0x22
	}
	if m.Disable {
		i--
		if m.Disable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Desc) > 0 {
		i -= len(m.Desc)
		copy(dAtA[i:], m.Desc)
		i = encodeVarintAction(dAtA, i, uint64(len(m.Desc)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintAction(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAction(dAtA []byte, offset int, v uint64) int {
	offset -= sovAction(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Action) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovAction(uint64(l))
	}
	l = len(m.Desc)
	if l > 0 {
		n += 1 + l + sovAction(uint64(l))
	}
	if m.Disable {
		n += 2
	}
	l = len(m.When)
	if l > 0 {
		n += 1 + l + sovAction(uint64(l))
	}
	if len(m.Then) > 0 {
		for _, s := range m.Then {
			l = len(s)
			n += 1 + l + sovAction(uint64(l))
		}
	}
	if m.AllowedActioner != 0 {
		n += 1 + sovAction(uint64(m.AllowedActioner))
	}
	return n
}

func sovAction(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAction(x uint64) (n int) {
	return sovAction(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Action) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAction
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
			return fmt.Errorf("proto: Action: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Action: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAction
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
				return ErrInvalidLengthAction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Desc", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAction
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
				return ErrInvalidLengthAction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Desc = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Disable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Disable = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field When", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAction
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
				return ErrInvalidLengthAction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.When = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Then", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAction
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
				return ErrInvalidLengthAction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Then = append(m.Then, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedActioner", wireType)
			}
			m.AllowedActioner = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AllowedActioner |= AllowedActioner(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAction
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
func skipAction(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAction
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
					return 0, ErrIntOverflowAction
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
					return 0, ErrIntOverflowAction
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
				return 0, ErrInvalidLengthAction
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAction
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAction
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAction        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAction          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAction = fmt.Errorf("proto: unexpected end of group")
)