// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftoracle/genesis.proto

package types

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

// GenesisState defines the nftoracle module's genesis state.
type GenesisState struct {
	Params                      Params                   `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	MintRequestList             []MintRequest            `protobuf:"bytes,2,rep,name=mintRequestList,proto3" json:"mintRequestList"`
	MintRequestCount            uint64                   `protobuf:"varint,3,opt,name=mintRequestCount,proto3" json:"mintRequestCount,omitempty"`
	ActionRequestList           []ActionRequest          `protobuf:"bytes,4,rep,name=actionRequestList,proto3" json:"actionRequestList"`
	ActionRequestCount          uint64                   `protobuf:"varint,5,opt,name=actionRequestCount,proto3" json:"actionRequestCount,omitempty"`
	CollectionOwnerRequestList  []CollectionOwnerRequest `protobuf:"bytes,6,rep,name=collectionOwnerRequestList,proto3" json:"collectionOwnerRequestList"`
	CollectionOwnerRequestCount uint64                   `protobuf:"varint,7,opt,name=collectionOwnerRequestCount,proto3" json:"collectionOwnerRequestCount,omitempty"`
	OracleConfig                *OracleConfig            `protobuf:"bytes,8,opt,name=oracle_config,json=oracleConfig,proto3" json:"oracle_config,omitempty"`
	ActionSignerList            []ActionSigner           `protobuf:"bytes,9,rep,name=actionSignerList,proto3" json:"actionSignerList"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a4e710a676a8edd, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetMintRequestList() []MintRequest {
	if m != nil {
		return m.MintRequestList
	}
	return nil
}

func (m *GenesisState) GetMintRequestCount() uint64 {
	if m != nil {
		return m.MintRequestCount
	}
	return 0
}

func (m *GenesisState) GetActionRequestList() []ActionRequest {
	if m != nil {
		return m.ActionRequestList
	}
	return nil
}

func (m *GenesisState) GetActionRequestCount() uint64 {
	if m != nil {
		return m.ActionRequestCount
	}
	return 0
}

func (m *GenesisState) GetCollectionOwnerRequestList() []CollectionOwnerRequest {
	if m != nil {
		return m.CollectionOwnerRequestList
	}
	return nil
}

func (m *GenesisState) GetCollectionOwnerRequestCount() uint64 {
	if m != nil {
		return m.CollectionOwnerRequestCount
	}
	return 0
}

func (m *GenesisState) GetOracleConfig() *OracleConfig {
	if m != nil {
		return m.OracleConfig
	}
	return nil
}

func (m *GenesisState) GetActionSignerList() []ActionSigner {
	if m != nil {
		return m.ActionSignerList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "thesixnetwork.sixnft.nftoracle.GenesisState")
}

func init() { proto.RegisterFile("nftoracle/genesis.proto", fileDescriptor_7a4e710a676a8edd) }

var fileDescriptor_7a4e710a676a8edd = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdf, 0xaf, 0xd2, 0x30,
	0x14, 0xc7, 0x37, 0xf9, 0xa1, 0x16, 0x8c, 0xd8, 0x18, 0x5d, 0xa6, 0x56, 0xe2, 0x83, 0x59, 0x14,
	0xb7, 0x88, 0x89, 0xcf, 0x0a, 0x26, 0x3e, 0xa8, 0x41, 0xc7, 0x9b, 0x26, 0x92, 0xb1, 0x94, 0xd1,
	0x08, 0x2d, 0x6e, 0x25, 0x60, 0xe2, 0x1f, 0x71, 0xff, 0x2c, 0x1e, 0x79, 0xbc, 0x4f, 0x37, 0x37,
	0xec, 0x1f, 0xb9, 0x59, 0xdb, 0xdc, 0x0e, 0x46, 0xd8, 0x7d, 0x5a, 0x77, 0xbe, 0xe7, 0x7c, 0xfb,
	0x39, 0xa7, 0x2d, 0x78, 0x4c, 0x27, 0x9c, 0xc5, 0x41, 0x38, 0xc3, 0x5e, 0x84, 0x29, 0x4e, 0x48,
	0xe2, 0x2e, 0x62, 0xc6, 0x19, 0x44, 0x7c, 0x8a, 0x13, 0xb2, 0xa6, 0x98, 0xaf, 0x58, 0xfc, 0xc7,
	0xcd, 0x96, 0x13, 0xee, 0x5e, 0x67, 0xdb, 0x0f, 0x23, 0x16, 0x31, 0x91, 0xea, 0x65, 0x2b, 0x59,
	0x65, 0x3f, 0xd2, 0x76, 0x8b, 0x20, 0x0e, 0xe6, 0xca, 0xcd, 0x7e, 0xaa, 0xe3, 0x73, 0x42, 0xf9,
	0x28, 0xc6, 0x7f, 0x97, 0x38, 0xe1, 0x4a, 0x45, 0x5a, 0x0d, 0x42, 0x4e, 0x18, 0x3d, 0xd0, 0x1d,
	0xad, 0x87, 0x6c, 0x36, 0xc3, 0x32, 0x87, 0xad, 0x28, 0x8e, 0x0f, 0x32, 0x9f, 0xe9, 0x4c, 0xf9,
	0x19, 0x85, 0x8c, 0x4e, 0x48, 0x54, 0x94, 0xd5, 0x46, 0x09, 0x89, 0x28, 0x8e, 0xa5, 0xfc, 0x22,
	0xad, 0x81, 0xe6, 0x67, 0x39, 0x85, 0x21, 0x0f, 0x38, 0x86, 0x9f, 0x40, 0x5d, 0xb6, 0x61, 0x99,
	0x6d, 0xd3, 0x69, 0x74, 0x5f, 0xba, 0xa7, 0xa7, 0xe2, 0x7e, 0x17, 0xd9, 0xbd, 0xea, 0xe6, 0xe2,
	0xb9, 0xe1, 0xab, 0x5a, 0xf8, 0x0b, 0xdc, 0xcf, 0x9a, 0xf6, 0x25, 0xe9, 0x57, 0x92, 0x70, 0xeb,
	0x56, 0xbb, 0xe2, 0x34, 0xba, 0xaf, 0xcb, 0xec, 0xbe, 0xe9, 0x32, 0xe5, 0x79, 0xe8, 0x04, 0x5f,
	0x81, 0x56, 0x2e, 0xd4, 0x67, 0x4b, 0xca, 0xad, 0x4a, 0xdb, 0x74, 0xaa, 0x7e, 0x21, 0x0e, 0x03,
	0xf0, 0x40, 0xb6, 0x9d, 0x47, 0xa9, 0x0a, 0x94, 0x37, 0x65, 0x28, 0x1f, 0xf3, 0x85, 0x0a, 0xa6,
	0xe8, 0x06, 0x5d, 0x00, 0xf7, 0x82, 0x12, 0xa8, 0x26, 0x80, 0x8e, 0x28, 0xf0, 0x3f, 0xb0, 0xf5,
	0x91, 0x0e, 0xb2, 0x13, 0xcd, 0xb3, 0xd5, 0x05, 0xdb, 0xfb, 0x32, 0xb6, 0xfe, 0x51, 0x07, 0x05,
	0x79, 0xc2, 0x1f, 0x7e, 0x00, 0x4f, 0x8e, 0xab, 0x12, 0xfb, 0xb6, 0xc0, 0x3e, 0x95, 0x02, 0x7f,
	0x80, 0x7b, 0x7b, 0x17, 0xcd, 0xba, 0x23, 0x2e, 0x4a, 0xa7, 0x0c, 0x79, 0x20, 0x3e, 0x7d, 0x51,
	0xe3, 0x37, 0x59, 0xee, 0x0f, 0xfe, 0x06, 0x2d, 0x39, 0xa8, 0xa1, 0xb8, 0x9b, 0x62, 0x10, 0x77,
	0xc5, 0x20, 0x3a, 0x37, 0x3b, 0x24, 0x59, 0xa7, 0xda, 0x2f, 0x78, 0xf5, 0xbe, 0x6c, 0x76, 0xc8,
	0xdc, 0xee, 0x90, 0x79, 0xb9, 0x43, 0xe6, 0x59, 0x8a, 0x8c, 0x6d, 0x8a, 0x8c, 0xf3, 0x14, 0x19,
	0x3f, 0xdf, 0x46, 0x84, 0x4f, 0x97, 0x63, 0x37, 0x64, 0x73, 0x6f, 0x6f, 0x27, 0x4f, 0xee, 0xe4,
	0xad, 0x3d, 0xfd, 0x80, 0xf8, 0xbf, 0x05, 0x4e, 0xc6, 0x75, 0xf1, 0x72, 0xde, 0x5d, 0x05, 0x00,
	0x00, 0xff, 0xff, 0xd7, 0x7d, 0x1a, 0x90, 0x48, 0x04, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ActionSignerList) > 0 {
		for iNdEx := len(m.ActionSignerList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ActionSignerList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if m.OracleConfig != nil {
		{
			size, err := m.OracleConfig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.CollectionOwnerRequestCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.CollectionOwnerRequestCount))
		i--
		dAtA[i] = 0x38
	}
	if len(m.CollectionOwnerRequestList) > 0 {
		for iNdEx := len(m.CollectionOwnerRequestList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CollectionOwnerRequestList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.ActionRequestCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.ActionRequestCount))
		i--
		dAtA[i] = 0x28
	}
	if len(m.ActionRequestList) > 0 {
		for iNdEx := len(m.ActionRequestList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ActionRequestList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.MintRequestCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.MintRequestCount))
		i--
		dAtA[i] = 0x18
	}
	if len(m.MintRequestList) > 0 {
		for iNdEx := len(m.MintRequestList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MintRequestList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.MintRequestList) > 0 {
		for _, e := range m.MintRequestList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.MintRequestCount != 0 {
		n += 1 + sovGenesis(uint64(m.MintRequestCount))
	}
	if len(m.ActionRequestList) > 0 {
		for _, e := range m.ActionRequestList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.ActionRequestCount != 0 {
		n += 1 + sovGenesis(uint64(m.ActionRequestCount))
	}
	if len(m.CollectionOwnerRequestList) > 0 {
		for _, e := range m.CollectionOwnerRequestList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.CollectionOwnerRequestCount != 0 {
		n += 1 + sovGenesis(uint64(m.CollectionOwnerRequestCount))
	}
	if m.OracleConfig != nil {
		l = m.OracleConfig.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.ActionSignerList) > 0 {
		for _, e := range m.ActionSignerList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintRequestList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintRequestList = append(m.MintRequestList, MintRequest{})
			if err := m.MintRequestList[len(m.MintRequestList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintRequestCount", wireType)
			}
			m.MintRequestCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MintRequestCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActionRequestList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ActionRequestList = append(m.ActionRequestList, ActionRequest{})
			if err := m.ActionRequestList[len(m.ActionRequestList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActionRequestCount", wireType)
			}
			m.ActionRequestCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ActionRequestCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollectionOwnerRequestList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollectionOwnerRequestList = append(m.CollectionOwnerRequestList, CollectionOwnerRequest{})
			if err := m.CollectionOwnerRequestList[len(m.CollectionOwnerRequestList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollectionOwnerRequestCount", wireType)
			}
			m.CollectionOwnerRequestCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CollectionOwnerRequestCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OracleConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OracleConfig == nil {
				m.OracleConfig = &OracleConfig{}
			}
			if err := m.OracleConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActionSignerList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ActionSignerList = append(m.ActionSignerList, ActionSigner{})
			if err := m.ActionSignerList[len(m.ActionSignerList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
