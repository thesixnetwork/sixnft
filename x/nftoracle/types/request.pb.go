// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftoracle/request.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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
	RequestStatus_EXPIRED                  RequestStatus = 3
	RequestStatus_FAILED_ON_EXECUTION      RequestStatus = 4
)

var RequestStatus_name = map[int32]string{
	0: "PENDING",
	1: "SUCCESS_WITH_CONSENSUS",
	2: "FAILED_WITHOUT_CONCENSUS",
	3: "EXPIRED",
	4: "FAILED_ON_EXECUTION",
}

var RequestStatus_value = map[string]int32{
	"PENDING":                  0,
	"SUCCESS_WITH_CONSENSUS":   1,
	"FAILED_WITHOUT_CONCENSUS": 2,
	"EXPIRED":                  3,
	"FAILED_ON_EXECUTION":      4,
}

func (x RequestStatus) String() string {
	return proto.EnumName(RequestStatus_name, int32(x))
}

func (RequestStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f60220a52ce49438, []int{0}
}

type NftOriginData struct {
	Image         string   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	HolderAddress string   `protobuf:"bytes,2,opt,name=holder_address,json=holderAddress,proto3" json:"holder_address,omitempty"`
	Traits        []*Trait `protobuf:"bytes,3,rep,name=traits,proto3" json:"traits,omitempty"`
}

func (m *NftOriginData) Reset()         { *m = NftOriginData{} }
func (m *NftOriginData) String() string { return proto.CompactTextString(m) }
func (*NftOriginData) ProtoMessage()    {}
func (*NftOriginData) Descriptor() ([]byte, []int) {
	return fileDescriptor_f60220a52ce49438, []int{0}
}
func (m *NftOriginData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NftOriginData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NftOriginData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NftOriginData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NftOriginData.Merge(m, src)
}
func (m *NftOriginData) XXX_Size() int {
	return m.Size()
}
func (m *NftOriginData) XXX_DiscardUnknown() {
	xxx_messageInfo_NftOriginData.DiscardUnknown(m)
}

var xxx_messageInfo_NftOriginData proto.InternalMessageInfo

func (m *NftOriginData) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *NftOriginData) GetHolderAddress() string {
	if m != nil {
		return m.HolderAddress
	}
	return ""
}

func (m *NftOriginData) GetTraits() []*Trait {
	if m != nil {
		return m.Traits
	}
	return nil
}

type DataHash struct {
	OriginData *NftOriginData `protobuf:"bytes,1,opt,name=origin_data,json=originData,proto3" json:"origin_data,omitempty"`
	Hash       []byte         `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Confirmers []string       `protobuf:"bytes,3,rep,name=confirmers,proto3" json:"confirmers,omitempty"`
}

func (m *DataHash) Reset()         { *m = DataHash{} }
func (m *DataHash) String() string { return proto.CompactTextString(m) }
func (*DataHash) ProtoMessage()    {}
func (*DataHash) Descriptor() ([]byte, []int) {
	return fileDescriptor_f60220a52ce49438, []int{1}
}
func (m *DataHash) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DataHash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DataHash.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DataHash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataHash.Merge(m, src)
}
func (m *DataHash) XXX_Size() int {
	return m.Size()
}
func (m *DataHash) XXX_DiscardUnknown() {
	xxx_messageInfo_DataHash.DiscardUnknown(m)
}

var xxx_messageInfo_DataHash proto.InternalMessageInfo

func (m *DataHash) GetOriginData() *NftOriginData {
	if m != nil {
		return m.OriginData
	}
	return nil
}

func (m *DataHash) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *DataHash) GetConfirmers() []string {
	if m != nil {
		return m.Confirmers
	}
	return nil
}

func init() {
	proto.RegisterEnum("thesixnetwork.sixnft.nftoracle.RequestStatus", RequestStatus_name, RequestStatus_value)
	proto.RegisterType((*NftOriginData)(nil), "thesixnetwork.sixnft.nftoracle.NftOriginData")
	proto.RegisterType((*DataHash)(nil), "thesixnetwork.sixnft.nftoracle.DataHash")
}

func init() { proto.RegisterFile("nftoracle/request.proto", fileDescriptor_f60220a52ce49438) }

var fileDescriptor_f60220a52ce49438 = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xce, 0x36, 0xb5, 0xda, 0x17, 0x23, 0x61, 0x2c, 0x36, 0x04, 0x59, 0x4b, 0xa1, 0x50, 0x04,
	0x77, 0xb1, 0x9e, 0x3d, 0xd4, 0xcd, 0x6a, 0x17, 0x65, 0xb6, 0xec, 0x26, 0x58, 0xbc, 0x2c, 0x93,
	0x64, 0x76, 0x77, 0x30, 0xbb, 0x13, 0x67, 0x5e, 0xb0, 0x82, 0xff, 0x40, 0x10, 0x7f, 0x96, 0xc7,
	0x1e, 0x3d, 0x4a, 0xf2, 0x47, 0x64, 0x67, 0x82, 0x6d, 0x2f, 0x7a, 0xfb, 0xe6, 0xfb, 0xbe, 0xf7,
	0xf8, 0xde, 0x7b, 0x03, 0xfb, 0x75, 0x8e, 0x52, 0xb1, 0xe9, 0x9c, 0xfb, 0x8a, 0x7f, 0x5a, 0x72,
	0x8d, 0xde, 0x42, 0x49, 0x94, 0xc4, 0xc5, 0x92, 0x6b, 0x71, 0x59, 0x73, 0xfc, 0x2c, 0xd5, 0x47,
	0xaf, 0x81, 0x39, 0x7a, 0x7f, 0xdd, 0x83, 0x1b, 0x85, 0x72, 0xc1, 0x6b, 0xcd, 0x99, 0x2d, 0x1c,
	0xec, 0x15, 0xb2, 0x90, 0x06, 0xfa, 0x0d, 0xda, 0xb0, 0x4f, 0x0a, 0x29, 0x8b, 0x39, 0xf7, 0xcd,
	0x6b, 0xb2, 0xcc, 0x7d, 0x14, 0x15, 0xd7, 0xc8, 0xaa, 0x85, 0x35, 0x1c, 0x7e, 0x73, 0xa0, 0x4b,
	0x73, 0x8c, 0x95, 0x28, 0x44, 0x3d, 0x64, 0xc8, 0xc8, 0x1e, 0xdc, 0x11, 0x15, 0x2b, 0x78, 0xdf,
	0x39, 0x70, 0x8e, 0x77, 0x13, 0xfb, 0x20, 0x47, 0xf0, 0xa0, 0x94, 0xf3, 0x19, 0x57, 0x19, 0x9b,
	0xcd, 0x14, 0xd7, 0xba, 0xbf, 0x65, 0xe4, 0xae, 0x65, 0x4f, 0x2d, 0x49, 0x5e, 0xc2, 0x0e, 0x2a,
	0x26, 0x50, 0xf7, 0xdb, 0x07, 0xed, 0xe3, 0xce, 0xc9, 0x91, 0xf7, 0xef, 0x79, 0xbc, 0x51, 0xe3,
	0x4e, 0x36, 0x45, 0x87, 0xdf, 0x1d, 0xb8, 0xd7, 0x84, 0x38, 0x63, 0xba, 0x24, 0x14, 0x3a, 0xd2,
	0xc4, 0xca, 0x66, 0x0c, 0x99, 0x89, 0xd3, 0x39, 0x79, 0xf6, 0xbf, 0x86, 0xb7, 0x86, 0x49, 0x40,
	0x5e, 0x0f, 0x46, 0x60, 0xbb, 0x64, 0xba, 0x34, 0xc1, 0xef, 0x27, 0x06, 0x13, 0x17, 0x60, 0x2a,
	0xeb, 0x5c, 0xa8, 0x8a, 0x2b, 0x9b, 0x79, 0x37, 0xb9, 0xc1, 0x3c, 0xfd, 0x0a, 0xdd, 0xc4, 0xde,
	0x27, 0x45, 0x86, 0x4b, 0x4d, 0x3a, 0x70, 0xf7, 0x3c, 0xa4, 0xc3, 0x88, 0xbe, 0xe9, 0xb5, 0xc8,
	0x00, 0x1e, 0xa5, 0xe3, 0x20, 0x08, 0xd3, 0x34, 0x7b, 0x1f, 0x8d, 0xce, 0xb2, 0x20, 0xa6, 0x69,
	0x48, 0xd3, 0x71, 0xda, 0x73, 0xc8, 0x63, 0xe8, 0xbf, 0x3e, 0x8d, 0xde, 0x85, 0x43, 0x23, 0xc5,
	0xe3, 0x51, 0xa3, 0x06, 0x56, 0xdd, 0x6a, 0xda, 0x84, 0x17, 0xe7, 0x51, 0x12, 0x0e, 0x7b, 0x6d,
	0xb2, 0x0f, 0x0f, 0x37, 0xd6, 0x98, 0x66, 0xe1, 0x45, 0x18, 0x8c, 0x47, 0x51, 0x4c, 0x7b, 0xdb,
	0xaf, 0xde, 0xfe, 0x5c, 0xb9, 0xce, 0xd5, 0xca, 0x75, 0x7e, 0xaf, 0x5c, 0xe7, 0xc7, 0xda, 0x6d,
	0x5d, 0xad, 0xdd, 0xd6, 0xaf, 0xb5, 0xdb, 0xfa, 0xf0, 0xbc, 0x10, 0x58, 0x2e, 0x27, 0xde, 0x54,
	0x56, 0xfe, 0xad, 0x85, 0xf8, 0x76, 0x21, 0xfe, 0xa5, 0x7f, 0xfd, 0x51, 0xf0, 0xcb, 0x82, 0xeb,
	0xc9, 0x8e, 0x39, 0xf8, 0x8b, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x94, 0x18, 0xb1, 0xf3, 0x7b,
	0x02, 0x00, 0x00,
}

func (m *NftOriginData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NftOriginData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NftOriginData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Traits) > 0 {
		for iNdEx := len(m.Traits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Traits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRequest(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.HolderAddress) > 0 {
		i -= len(m.HolderAddress)
		copy(dAtA[i:], m.HolderAddress)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.HolderAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Image) > 0 {
		i -= len(m.Image)
		copy(dAtA[i:], m.Image)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.Image)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DataHash) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DataHash) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DataHash) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Confirmers) > 0 {
		for iNdEx := len(m.Confirmers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Confirmers[iNdEx])
			copy(dAtA[i:], m.Confirmers[iNdEx])
			i = encodeVarintRequest(dAtA, i, uint64(len(m.Confirmers[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x12
	}
	if m.OriginData != nil {
		{
			size, err := m.OriginData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRequest(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRequest(dAtA []byte, offset int, v uint64) int {
	offset -= sovRequest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NftOriginData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Image)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	l = len(m.HolderAddress)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	if len(m.Traits) > 0 {
		for _, e := range m.Traits {
			l = e.Size()
			n += 1 + l + sovRequest(uint64(l))
		}
	}
	return n
}

func (m *DataHash) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OriginData != nil {
		l = m.OriginData.Size()
		n += 1 + l + sovRequest(uint64(l))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	if len(m.Confirmers) > 0 {
		for _, s := range m.Confirmers {
			l = len(s)
			n += 1 + l + sovRequest(uint64(l))
		}
	}
	return n
}

func sovRequest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRequest(x uint64) (n int) {
	return sovRequest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NftOriginData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRequest
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
			return fmt.Errorf("proto: NftOriginData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NftOriginData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Image", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Image = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HolderAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HolderAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Traits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Traits = append(m.Traits, &Trait{})
			if err := m.Traits[len(m.Traits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRequest
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
func (m *DataHash) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRequest
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
			return fmt.Errorf("proto: DataHash: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataHash: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OriginData == nil {
				m.OriginData = &NftOriginData{}
			}
			if err := m.OriginData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Confirmers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Confirmers = append(m.Confirmers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRequest
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
func skipRequest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRequest
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
					return 0, ErrIntOverflowRequest
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
					return 0, ErrIntOverflowRequest
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
				return 0, ErrInvalidLengthRequest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRequest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRequest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRequest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRequest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRequest = fmt.Errorf("proto: unexpected end of group")
)
