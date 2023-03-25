// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftoracle/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Params defines the parameters for the module.
type Params struct {
	MintRequestActiveDuration      time.Duration `protobuf:"bytes,1,opt,name=mint_request_active_duration,json=mintRequestActiveDuration,proto3,stdduration" json:"mint_request_active_duration" yaml:"mint_request_active_duration"`
	ActionRequestActiveDuration    time.Duration `protobuf:"bytes,2,opt,name=action_request_active_duration,json=actionRequestActiveDuration,proto3,stdduration" json:"action_request_active_duration" yaml:"action_request_active_duration"`
	VerifyRequestActiveDuration    time.Duration `protobuf:"bytes,3,opt,name=verify_request_active_duration,json=verifyRequestActiveDuration,proto3,stdduration" json:"verify_request_active_duration" yaml:"verify_request_active_duration"`
	ActionSignerActiveDuration     time.Duration `protobuf:"bytes,4,opt,name=action_signer_active_duration,json=actionSignerActiveDuration,proto3,stdduration" json:"action_signer_active_duration" yaml:"action_signer_active_duration"`
	SyncActionSignerActiveDuration time.Duration `protobuf:"bytes,5,opt,name=sync_action_signer_active_duration,json=syncActionSignerActiveDuration,proto3,stdduration" json:"sync_action_signer_active_duration" yaml:"sync_action_signer_active_duration"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_83d28596b5eb4570, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMintRequestActiveDuration() time.Duration {
	if m != nil {
		return m.MintRequestActiveDuration
	}
	return 0
}

func (m *Params) GetActionRequestActiveDuration() time.Duration {
	if m != nil {
		return m.ActionRequestActiveDuration
	}
	return 0
}

func (m *Params) GetVerifyRequestActiveDuration() time.Duration {
	if m != nil {
		return m.VerifyRequestActiveDuration
	}
	return 0
}

func (m *Params) GetActionSignerActiveDuration() time.Duration {
	if m != nil {
		return m.ActionSignerActiveDuration
	}
	return 0
}

func (m *Params) GetSyncActionSignerActiveDuration() time.Duration {
	if m != nil {
		return m.SyncActionSignerActiveDuration
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "thesixnetwork.sixnft.nftoracle.Params")
}

func init() { proto.RegisterFile("nftoracle/params.proto", fileDescriptor_83d28596b5eb4570) }

var fileDescriptor_83d28596b5eb4570 = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xd3, 0xb1, 0x4e, 0xfa, 0x40,
	0x1c, 0x07, 0xf0, 0xde, 0xff, 0x0f, 0x0c, 0x75, 0x23, 0xc6, 0x08, 0xea, 0xd5, 0x54, 0x4d, 0x74,
	0xe9, 0x89, 0xc6, 0x85, 0x0d, 0xe2, 0xe6, 0x62, 0x70, 0x73, 0x69, 0x4a, 0xbd, 0x96, 0x46, 0x7a,
	0x87, 0xd7, 0x2b, 0xd2, 0x67, 0x30, 0x31, 0xba, 0x31, 0x38, 0xf8, 0x38, 0x8c, 0x8c, 0x4e, 0x68,
	0xe0, 0x0d, 0x8c, 0x0f, 0x60, 0xae, 0x57, 0x34, 0x28, 0x6d, 0xe3, 0x76, 0x4d, 0xbf, 0xbf, 0x6f,
	0x3f, 0xbd, 0xe4, 0xa7, 0xae, 0x11, 0x87, 0x53, 0x66, 0xd9, 0x5d, 0x8c, 0x7a, 0x16, 0xb3, 0xfc,
	0xc0, 0xe8, 0x31, 0xca, 0x69, 0x19, 0xf2, 0x0e, 0x0e, 0xbc, 0x01, 0xc1, 0xfc, 0x96, 0xb2, 0x6b,
	0x43, 0x1c, 0x1d, 0x6e, 0x7c, 0x85, 0xab, 0xab, 0x2e, 0x75, 0x69, 0x1c, 0x45, 0xe2, 0x24, 0xa7,
	0xaa, 0xd0, 0xa5, 0xd4, 0x15, 0x55, 0xe2, 0xa9, 0x1d, 0x3a, 0xe8, 0x2a, 0x64, 0x16, 0xf7, 0x28,
	0x91, 0xef, 0xf5, 0x8f, 0xa2, 0x5a, 0x3a, 0x8f, 0x3f, 0x53, 0xbe, 0x03, 0xea, 0xa6, 0xef, 0x11,
	0x6e, 0x32, 0x7c, 0x13, 0xe2, 0x80, 0x9b, 0x96, 0xcd, 0xbd, 0x3e, 0x36, 0xe7, 0x13, 0xeb, 0x60,
	0x1b, 0xec, 0xaf, 0x1c, 0x55, 0x0c, 0x59, 0x69, 0xcc, 0x2b, 0x8d, 0xd3, 0x24, 0xd0, 0x44, 0xa3,
	0x89, 0xa6, 0xbc, 0x4f, 0xb4, 0x9d, 0xc8, 0xf2, 0xbb, 0x75, 0x3d, 0xab, 0x4c, 0x1f, 0xbe, 0x6a,
	0xa0, 0x55, 0x11, 0x91, 0x96, 0x4c, 0x34, 0xe2, 0xc0, 0xbc, 0xab, 0xfc, 0x08, 0x54, 0x28, 0x66,
	0x28, 0x49, 0xf5, 0xfc, 0xcb, 0xf3, 0xd4, 0x12, 0xcf, 0x9e, 0xf4, 0x64, 0xd7, 0x49, 0xd1, 0x86,
	0x0c, 0xa5, 0x9b, 0xfa, 0x98, 0x79, 0x4e, 0x94, 0x6a, 0xfa, 0xff, 0x47, 0x53, 0x76, 0x5d, 0x62,
	0x92, 0xa1, 0xe5, 0xa6, 0x7b, 0xa0, 0x6e, 0x25, 0x3f, 0x16, 0x78, 0x2e, 0xc1, 0xec, 0x17, 0xa9,
	0x90, 0x47, 0x3a, 0x4c, 0x48, 0xbb, 0x0b, 0xd7, 0xb4, 0xbc, 0x4d, 0x8a, 0xaa, 0x32, 0x73, 0x11,
	0x47, 0x7e, 0x80, 0x9e, 0x80, 0xaa, 0x07, 0x11, 0xb1, 0xcd, 0x6c, 0x55, 0x31, 0x4f, 0x75, 0x92,
	0xa8, 0x0e, 0xa4, 0x2a, 0xbf, 0x52, 0xd2, 0xa0, 0x08, 0x36, 0x52, 0x79, 0xf5, 0xc2, 0xf0, 0x59,
	0x53, 0x9a, 0x67, 0xa3, 0x29, 0x04, 0xe3, 0x29, 0x04, 0x6f, 0x53, 0x08, 0x1e, 0x66, 0x50, 0x19,
	0xcf, 0xa0, 0xf2, 0x32, 0x83, 0xca, 0x65, 0xcd, 0xf5, 0x78, 0x27, 0x6c, 0x1b, 0x36, 0xf5, 0xd1,
	0xc2, 0xc6, 0x21, 0xb9, 0x71, 0x68, 0x80, 0xbe, 0x17, 0x94, 0x47, 0x3d, 0x1c, 0xb4, 0x4b, 0x31,
	0xfe, 0xf8, 0x33, 0x00, 0x00, 0xff, 0xff, 0xfa, 0xdb, 0xbe, 0x8e, 0xba, 0x03, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.SyncActionSignerActiveDuration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.SyncActionSignerActiveDuration):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintParams(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x2a
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.ActionSignerActiveDuration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.ActionSignerActiveDuration):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintParams(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x22
	n3, err3 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.VerifyRequestActiveDuration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.VerifyRequestActiveDuration):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintParams(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x1a
	n4, err4 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.ActionRequestActiveDuration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.ActionRequestActiveDuration):])
	if err4 != nil {
		return 0, err4
	}
	i -= n4
	i = encodeVarintParams(dAtA, i, uint64(n4))
	i--
	dAtA[i] = 0x12
	n5, err5 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.MintRequestActiveDuration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.MintRequestActiveDuration):])
	if err5 != nil {
		return 0, err5
	}
	i -= n5
	i = encodeVarintParams(dAtA, i, uint64(n5))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.MintRequestActiveDuration)
	n += 1 + l + sovParams(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.ActionRequestActiveDuration)
	n += 1 + l + sovParams(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.VerifyRequestActiveDuration)
	n += 1 + l + sovParams(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.ActionSignerActiveDuration)
	n += 1 + l + sovParams(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.SyncActionSignerActiveDuration)
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintRequestActiveDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.MintRequestActiveDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActionRequestActiveDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.ActionRequestActiveDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VerifyRequestActiveDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.VerifyRequestActiveDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActionSignerActiveDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.ActionSignerActiveDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SyncActionSignerActiveDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.SyncActionSignerActiveDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
