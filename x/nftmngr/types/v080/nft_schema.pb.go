// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftmngr/v080/nft_schema.proto

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

type NFTSchema struct {
	Code              string       `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name              string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Owner             string       `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	SystemActioners   []string     `protobuf:"bytes,4,rep,name=system_actioners,json=systemActioners,proto3" json:"system_actioners,omitempty"`
	OriginData        *OriginData  `protobuf:"bytes,5,opt,name=origin_data,json=originData,proto3" json:"origin_data,omitempty"`
	OnchainData       *OnChainData `protobuf:"bytes,6,opt,name=onchain_data,json=onchainData,proto3" json:"onchain_data,omitempty"`
	IsVerified        bool         `protobuf:"varint,7,opt,name=isVerified,proto3" json:"isVerified,omitempty"`
	MintAuthorization string       `protobuf:"bytes,8,opt,name=mint_authorization,json=mintAuthorization,proto3" json:"mint_authorization,omitempty"`
}

func (m *NFTSchema) Reset()         { *m = NFTSchema{} }
func (m *NFTSchema) String() string { return proto.CompactTextString(m) }
func (*NFTSchema) ProtoMessage()    {}
func (*NFTSchema) Descriptor() ([]byte, []int) {
	return fileDescriptor_eab8a847090eb8a2, []int{0}
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

func (m *NFTSchema) GetSystemActioners() []string {
	if m != nil {
		return m.SystemActioners
	}
	return nil
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

func (m *NFTSchema) GetIsVerified() bool {
	if m != nil {
		return m.IsVerified
	}
	return false
}

func (m *NFTSchema) GetMintAuthorization() string {
	if m != nil {
		return m.MintAuthorization
	}
	return ""
}

func init() {
	proto.RegisterType((*NFTSchema)(nil), "thesixnetwork.sixnft.nftmngr.v080.NFTSchema")
}

func init() { proto.RegisterFile("nftmngr/v080/nft_schema.proto", fileDescriptor_eab8a847090eb8a2) }

var fileDescriptor_eab8a847090eb8a2 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xbf, 0x4e, 0xf3, 0x30,
	0x14, 0xc5, 0x9b, 0xfe, 0xfb, 0x5a, 0xf7, 0x93, 0x00, 0x8b, 0xc1, 0xaa, 0x84, 0x15, 0x98, 0xc2,
	0x50, 0xa7, 0x02, 0x06, 0xd6, 0x02, 0x62, 0x2c, 0x22, 0x20, 0x06, 0x96, 0xc8, 0x4d, 0x9d, 0xc6,
	0x42, 0xb1, 0x2b, 0xc7, 0xa5, 0x2d, 0x4f, 0xc1, 0x63, 0x31, 0x76, 0x64, 0x44, 0xad, 0xc4, 0x73,
	0x20, 0x3b, 0x2d, 0x34, 0x13, 0x6c, 0xc7, 0xe7, 0xfa, 0xfc, 0x7c, 0xaf, 0x2f, 0x38, 0x10, 0xb1,
	0x4e, 0xc5, 0x48, 0xf9, 0xcf, 0xdd, 0xf3, 0xae, 0x2f, 0x62, 0x1d, 0x66, 0x51, 0xc2, 0x52, 0x4a,
	0xc6, 0x4a, 0x6a, 0x09, 0x0f, 0x75, 0xc2, 0x32, 0x3e, 0x13, 0x4c, 0x4f, 0xa5, 0x7a, 0x22, 0x46,
	0xc6, 0x9a, 0xac, 0x33, 0xc4, 0x64, 0xda, 0x6e, 0x81, 0x20, 0x45, 0x18, 0x25, 0x94, 0x8b, 0x70,
	0x48, 0xf5, 0x1a, 0xd2, 0xc6, 0xc5, 0x1b, 0x8a, 0x8f, 0x0a, 0xf5, 0xa3, 0xcf, 0x32, 0x68, 0xf6,
	0xaf, 0xef, 0xef, 0xec, 0xc3, 0x10, 0x82, 0x6a, 0x24, 0x87, 0x0c, 0x39, 0xae, 0xe3, 0x35, 0x03,
	0xab, 0x8d, 0x27, 0x68, 0xca, 0x50, 0x39, 0xf7, 0x8c, 0x86, 0xfb, 0xa0, 0x26, 0xa7, 0x82, 0x29,
	0x54, 0xb1, 0x66, 0x7e, 0x80, 0xc7, 0x60, 0x37, 0x9b, 0x67, 0x9a, 0xa5, 0x21, 0x8d, 0x34, 0x97,
	0x82, 0xa9, 0x0c, 0x55, 0xdd, 0x8a, 0xd7, 0x0c, 0x76, 0x72, 0xbf, 0xb7, 0xb1, 0x61, 0x1f, 0xb4,
	0xb6, 0x7a, 0x41, 0x35, 0xd7, 0xf1, 0x5a, 0x27, 0x1d, 0xf2, 0xeb, 0xc4, 0xe4, 0xc6, 0xa6, 0xae,
	0xa8, 0xa6, 0x01, 0x90, 0xdf, 0x1a, 0xde, 0x82, 0xff, 0x52, 0xfc, 0x0c, 0x8f, 0xea, 0x16, 0x48,
	0xfe, 0x02, 0x14, 0x97, 0x26, 0x66, 0x89, 0xad, 0x35, 0xc3, 0x22, 0x31, 0x00, 0x3c, 0x7b, 0x60,
	0x8a, 0xc7, 0x9c, 0x0d, 0xd1, 0x3f, 0xd7, 0xf1, 0x1a, 0xc1, 0x96, 0x03, 0x3b, 0x00, 0xa6, 0x5c,
	0xe8, 0x90, 0x4e, 0x74, 0x22, 0x15, 0x7f, 0xa1, 0x66, 0x36, 0xd4, 0xb0, 0x1f, 0xb2, 0x67, 0x2a,
	0xbd, 0xed, 0xc2, 0x45, 0xff, 0x6d, 0x89, 0x9d, 0xc5, 0x12, 0x3b, 0x1f, 0x4b, 0xec, 0xbc, 0xae,
	0x70, 0x69, 0xb1, 0xc2, 0xa5, 0xf7, 0x15, 0x2e, 0x3d, 0x9e, 0x8d, 0xb8, 0x4e, 0x26, 0x03, 0x12,
	0xc9, 0xd4, 0x2f, 0xf4, 0xeb, 0xe7, 0xfd, 0xfa, 0x33, 0x7f, 0xb3, 0x44, 0x3d, 0x1f, 0xb3, 0xcc,
	0xae, 0x72, 0x50, 0xb7, 0xfb, 0x3b, 0xfd, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x91, 0xe6, 0x25, 0x78,
	0x45, 0x02, 0x00, 0x00,
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
	if len(m.MintAuthorization) > 0 {
		i -= len(m.MintAuthorization)
		copy(dAtA[i:], m.MintAuthorization)
		i = encodeVarintNftSchema(dAtA, i, uint64(len(m.MintAuthorization)))
		i--
		dAtA[i] = 0x42
	}
	if m.IsVerified {
		i--
		if m.IsVerified {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
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
		dAtA[i] = 0x32
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
		dAtA[i] = 0x2a
	}
	if len(m.SystemActioners) > 0 {
		for iNdEx := len(m.SystemActioners) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.SystemActioners[iNdEx])
			copy(dAtA[i:], m.SystemActioners[iNdEx])
			i = encodeVarintNftSchema(dAtA, i, uint64(len(m.SystemActioners[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
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
	if len(m.SystemActioners) > 0 {
		for _, s := range m.SystemActioners {
			l = len(s)
			n += 1 + l + sovNftSchema(uint64(l))
		}
	}
	if m.OriginData != nil {
		l = m.OriginData.Size()
		n += 1 + l + sovNftSchema(uint64(l))
	}
	if m.OnchainData != nil {
		l = m.OnchainData.Size()
		n += 1 + l + sovNftSchema(uint64(l))
	}
	if m.IsVerified {
		n += 2
	}
	l = len(m.MintAuthorization)
	if l > 0 {
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
				return fmt.Errorf("proto: wrong wireType = %d for field SystemActioners", wireType)
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
			m.SystemActioners = append(m.SystemActioners, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
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
		case 6:
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
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsVerified", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftSchema
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
			m.IsVerified = bool(v != 0)
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintAuthorization", wireType)
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
			m.MintAuthorization = string(dAtA[iNdEx:postIndex])
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
