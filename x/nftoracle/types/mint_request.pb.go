// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftoracle/mint_request.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

type MintRequest struct {
	Id              uint64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	NftSchemaCode   string          `protobuf:"bytes,2,opt,name=nft_schema_code,json=nftSchemaCode,proto3" json:"nft_schema_code,omitempty"`
	TokenId         string          `protobuf:"bytes,3,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	RequiredConfirm uint64          `protobuf:"varint,4,opt,name=required_confirm,json=requiredConfirm,proto3" json:"required_confirm,omitempty"`
	Status          RequestStatus   `protobuf:"varint,5,opt,name=status,proto3,enum=sixnft.nftoracle.RequestStatus" json:"status,omitempty"`
	CurrentConfirm  uint64          `protobuf:"varint,6,opt,name=current_confirm,json=currentConfirm,proto3" json:"current_confirm,omitempty"`
	Confirmers      map[string]bool `protobuf:"bytes,7,rep,name=confirmers,proto3" json:"confirmers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// NftOriginData nft_origin_data = 8;
	CreatedAt     time.Time   `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at"`
	ValidUntil    time.Time   `protobuf:"bytes,9,opt,name=valid_until,json=validUntil,proto3,stdtime" json:"valid_until"`
	DataHashes    []*DataHash `protobuf:"bytes,10,rep,name=data_hashes,json=dataHashes,proto3" json:"data_hashes,omitempty"`
	ExpiredHeight int64       `protobuf:"varint,11,opt,name=expired_height,json=expiredHeight,proto3" json:"expired_height,omitempty"`
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

func (m *MintRequest) GetCurrentConfirm() uint64 {
	if m != nil {
		return m.CurrentConfirm
	}
	return 0
}

func (m *MintRequest) GetConfirmers() map[string]bool {
	if m != nil {
		return m.Confirmers
	}
	return nil
}

func (m *MintRequest) GetCreatedAt() time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return time.Time{}
}

func (m *MintRequest) GetValidUntil() time.Time {
	if m != nil {
		return m.ValidUntil
	}
	return time.Time{}
}

func (m *MintRequest) GetDataHashes() []*DataHash {
	if m != nil {
		return m.DataHashes
	}
	return nil
}

func (m *MintRequest) GetExpiredHeight() int64 {
	if m != nil {
		return m.ExpiredHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*MintRequest)(nil), "sixnft.nftoracle.MintRequest")
	proto.RegisterMapType((map[string]bool)(nil), "sixnft.nftoracle.MintRequest.ConfirmersEntry")
}

func init() { proto.RegisterFile("nftoracle/mint_request.proto", fileDescriptor_3b5fc42e48c76b2e) }

var fileDescriptor_3b5fc42e48c76b2e = []byte{
	// 524 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0xc7, 0xe3, 0xa4, 0x2f, 0xc9, 0x5a, 0x4d, 0xaa, 0x55, 0xa5, 0xc7, 0x4f, 0x84, 0x1c, 0x0b,
	0x09, 0x30, 0x07, 0x6c, 0x51, 0x0e, 0x20, 0x10, 0x07, 0x1a, 0x2a, 0x15, 0xa1, 0x5e, 0x5c, 0xb8,
	0x70, 0xb1, 0x36, 0xf6, 0xd8, 0x5e, 0x25, 0xde, 0x0d, 0xde, 0x71, 0x49, 0xbe, 0x45, 0x3f, 0x56,
	0x8f, 0x3d, 0x72, 0x02, 0x94, 0x1c, 0xf9, 0x12, 0x28, 0x6b, 0x27, 0x0d, 0xed, 0x89, 0xdb, 0xcc,
	0xdf, 0xbf, 0x79, 0xd9, 0xbf, 0x87, 0x3c, 0x10, 0x09, 0xca, 0x82, 0x45, 0x13, 0xf0, 0x73, 0x2e,
	0x30, 0x2c, 0xe0, 0x6b, 0x09, 0x0a, 0xbd, 0x69, 0x21, 0x51, 0xd2, 0x43, 0xc5, 0x67, 0x22, 0x41,
	0x6f, 0x03, 0xf5, 0xff, 0xbb, 0xe5, 0xe5, 0x14, 0x84, 0x02, 0x56, 0xa1, 0xdb, 0x1f, 0xfe, 0xea,
	0xd1, 0x3f, 0x4a, 0x65, 0x2a, 0x75, 0xe8, 0xaf, 0xa2, 0x5a, 0x1d, 0xa4, 0x52, 0xa6, 0x13, 0xf0,
	0x75, 0x36, 0x2a, 0x13, 0x1f, 0x79, 0x0e, 0x0a, 0x59, 0x3e, 0xad, 0x80, 0x87, 0xbf, 0x77, 0x88,
	0x79, 0xce, 0x05, 0x06, 0x55, 0x33, 0xda, 0x25, 0x4d, 0x1e, 0x5b, 0x86, 0x63, 0xb8, 0x3b, 0x41,
	0x93, 0xc7, 0xf4, 0x31, 0xe9, 0x89, 0x04, 0x43, 0x15, 0x65, 0x90, 0xb3, 0x30, 0x92, 0x31, 0x58,
	0x4d, 0xc7, 0x70, 0x3b, 0xc1, 0x81, 0x48, 0xf0, 0x42, 0xab, 0x43, 0x19, 0x03, 0xfd, 0x9f, 0xb4,
	0x51, 0x8e, 0x41, 0x84, 0x3c, 0xb6, 0x5a, 0x1a, 0xd8, 0xd7, 0xf9, 0x87, 0x98, 0x3e, 0x25, 0x87,
	0xab, 0x55, 0x79, 0x01, 0x71, 0x18, 0x49, 0x91, 0xf0, 0x22, 0xb7, 0x76, 0xf4, 0x80, 0xde, 0x5a,
	0x1f, 0x56, 0x32, 0x7d, 0x49, 0xf6, 0x14, 0x32, 0x2c, 0x95, 0xb5, 0xeb, 0x18, 0x6e, 0xf7, 0x78,
	0xe0, 0xdd, 0x75, 0xc6, 0xab, 0x17, 0xbd, 0xd0, 0x58, 0x50, 0xe3, 0xf4, 0x09, 0xe9, 0x45, 0x65,
	0x51, 0x80, 0xc0, 0xcd, 0x88, 0x3d, 0x3d, 0xa2, 0x5b, 0xcb, 0xeb, 0x09, 0xe7, 0x84, 0xd4, 0x00,
	0x14, 0xca, 0xda, 0x77, 0x5a, 0xae, 0x79, 0xfc, 0xec, 0xfe, 0x94, 0x2d, 0x4b, 0xbc, 0xe1, 0x86,
	0x3f, 0x15, 0x58, 0xcc, 0x83, 0xad, 0x06, 0x74, 0x48, 0x48, 0x54, 0x00, 0x43, 0x88, 0x43, 0x86,
	0x56, 0xdb, 0x31, 0x5c, 0xf3, 0xb8, 0xef, 0x55, 0xa6, 0x7b, 0x6b, 0xd3, 0xbd, 0x4f, 0x6b, 0xd3,
	0x4f, 0xda, 0xd7, 0x3f, 0x06, 0x8d, 0xab, 0x9f, 0x03, 0x23, 0xe8, 0xd4, 0x75, 0xef, 0x90, 0x9e,
	0x12, 0xf3, 0x92, 0x4d, 0x78, 0x1c, 0x96, 0x02, 0xf9, 0xc4, 0xea, 0xfc, 0x43, 0x17, 0xa2, 0x0b,
	0x3f, 0xaf, 0xea, 0xe8, 0x1b, 0x62, 0xc6, 0x0c, 0x59, 0x98, 0x31, 0x95, 0x81, 0xb2, 0x88, 0x7e,
	0x5b, 0xff, 0xfe, 0xdb, 0xde, 0x33, 0x64, 0x67, 0x4c, 0x65, 0x01, 0x89, 0xeb, 0x08, 0x14, 0x7d,
	0x44, 0xba, 0x30, 0x9b, 0xea, 0x7f, 0x94, 0x01, 0x4f, 0x33, 0xb4, 0x4c, 0xc7, 0x70, 0x5b, 0xc1,
	0x41, 0xad, 0x9e, 0x69, 0xb1, 0xff, 0x96, 0xf4, 0xee, 0xd8, 0x41, 0x0f, 0x49, 0x6b, 0x0c, 0x73,
	0x7d, 0x32, 0x9d, 0x60, 0x15, 0xd2, 0x23, 0xb2, 0x7b, 0xc9, 0x26, 0x65, 0x75, 0x29, 0xed, 0xa0,
	0x4a, 0x5e, 0x37, 0x5f, 0x19, 0x27, 0x1f, 0xaf, 0x17, 0xb6, 0x71, 0xb3, 0xb0, 0x8d, 0x5f, 0x0b,
	0xdb, 0xb8, 0x5a, 0xda, 0x8d, 0x9b, 0xa5, 0xdd, 0xf8, 0xbe, 0xb4, 0x1b, 0x5f, 0x9e, 0xa7, 0x1c,
	0xb3, 0x72, 0xe4, 0x45, 0x32, 0xf7, 0x31, 0x83, 0xd5, 0xd2, 0x80, 0xdf, 0x64, 0x31, 0xf6, 0xab,
	0xfd, 0xfd, 0x99, 0x7f, 0x7b, 0xf9, 0x38, 0x9f, 0x82, 0x1a, 0xed, 0x69, 0x67, 0x5e, 0xfc, 0x09,
	0x00, 0x00, 0xff, 0xff, 0x51, 0x09, 0xb5, 0x44, 0x5c, 0x03, 0x00, 0x00,
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
	if m.ExpiredHeight != 0 {
		i = encodeVarintMintRequest(dAtA, i, uint64(m.ExpiredHeight))
		i--
		dAtA[i] = 0x58
	}
	if len(m.DataHashes) > 0 {
		for iNdEx := len(m.DataHashes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DataHashes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMintRequest(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.ValidUntil, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.ValidUntil):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintMintRequest(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x4a
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CreatedAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintMintRequest(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x42
	if len(m.Confirmers) > 0 {
		for k := range m.Confirmers {
			v := m.Confirmers[k]
			baseI := i
			i--
			if v {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x10
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintMintRequest(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintMintRequest(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.CurrentConfirm != 0 {
		i = encodeVarintMintRequest(dAtA, i, uint64(m.CurrentConfirm))
		i--
		dAtA[i] = 0x30
	}
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
	if m.CurrentConfirm != 0 {
		n += 1 + sovMintRequest(uint64(m.CurrentConfirm))
	}
	if len(m.Confirmers) > 0 {
		for k, v := range m.Confirmers {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovMintRequest(uint64(len(k))) + 1 + 1
			n += mapEntrySize + 1 + sovMintRequest(uint64(mapEntrySize))
		}
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt)
	n += 1 + l + sovMintRequest(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.ValidUntil)
	n += 1 + l + sovMintRequest(uint64(l))
	if len(m.DataHashes) > 0 {
		for _, e := range m.DataHashes {
			l = e.Size()
			n += 1 + l + sovMintRequest(uint64(l))
		}
	}
	if m.ExpiredHeight != 0 {
		n += 1 + sovMintRequest(uint64(m.ExpiredHeight))
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
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentConfirm", wireType)
			}
			m.CurrentConfirm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMintRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentConfirm |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Confirmers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMintRequest
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
				return ErrInvalidLengthMintRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMintRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Confirmers == nil {
				m.Confirmers = make(map[string]bool)
			}
			var mapkey string
			var mapvalue bool
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMintRequest
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthMintRequest
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthMintRequest
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapvaluetemp int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMintRequest
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvaluetemp |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					mapvalue = bool(mapvaluetemp != 0)
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipMintRequest(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthMintRequest
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Confirmers[mapkey] = mapvalue
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMintRequest
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
				return ErrInvalidLengthMintRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMintRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidUntil", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMintRequest
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
				return ErrInvalidLengthMintRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMintRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.ValidUntil, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataHashes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMintRequest
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
				return ErrInvalidLengthMintRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMintRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataHashes = append(m.DataHashes, &DataHash{})
			if err := m.DataHashes[len(m.DataHashes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiredHeight", wireType)
			}
			m.ExpiredHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMintRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpiredHeight |= int64(b&0x7F) << shift
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
