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
	ActionRequestList           []ActionOracleRequest    `protobuf:"bytes,4,rep,name=actionRequestList,proto3" json:"actionRequestList"`
	ActionRequestCount          uint64                   `protobuf:"varint,5,opt,name=actionRequestCount,proto3" json:"actionRequestCount,omitempty"`
	CollectionOwnerRequestList  []CollectionOwnerRequest `protobuf:"bytes,6,rep,name=collectionOwnerRequestList,proto3" json:"collectionOwnerRequestList"`
	CollectionOwnerRequestCount uint64                   `protobuf:"varint,7,opt,name=collectionOwnerRequestCount,proto3" json:"collectionOwnerRequestCount,omitempty"`
	OracleConfig                *OracleConfig            `protobuf:"bytes,8,opt,name=oracle_config,json=oracleConfig,proto3" json:"oracle_config,omitempty"`
	ActionSignerList            []ActionSigner           `protobuf:"bytes,9,rep,name=actionSignerList,proto3" json:"actionSignerList"`
	BindedSignerList            []BindedSigner           `protobuf:"bytes,10,rep,name=bindedSignerList,proto3" json:"bindedSignerList"`
	ActionSignerConfigList      []ActionSignerConfig     `protobuf:"bytes,13,rep,name=actionSignerConfigList,proto3" json:"actionSignerConfigList"`
	SyncActionSignerList        []SyncActionSigner       `protobuf:"bytes,14,rep,name=syncActionSignerList,proto3" json:"syncActionSignerList"`
	SyncActionSignerCount       uint64                   `protobuf:"varint,15,opt,name=syncActionSignerCount,proto3" json:"syncActionSignerCount,omitempty"`
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

func (m *GenesisState) GetActionRequestList() []ActionOracleRequest {
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

func (m *GenesisState) GetBindedSignerList() []BindedSigner {
	if m != nil {
		return m.BindedSignerList
	}
	return nil
}

func (m *GenesisState) GetActionSignerConfigList() []ActionSignerConfig {
	if m != nil {
		return m.ActionSignerConfigList
	}
	return nil
}

func (m *GenesisState) GetSyncActionSignerList() []SyncActionSigner {
	if m != nil {
		return m.SyncActionSignerList
	}
	return nil
}

func (m *GenesisState) GetSyncActionSignerCount() uint64 {
	if m != nil {
		return m.SyncActionSignerCount
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "thesixnetwork.sixnft.nftoracle.GenesisState")
}

func init() { proto.RegisterFile("nftoracle/genesis.proto", fileDescriptor_7a4e710a676a8edd) }

var fileDescriptor_7a4e710a676a8edd = []byte{
	// 538 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x8e, 0x69, 0x09, 0x65, 0xdb, 0xd2, 0xb2, 0x2a, 0x25, 0x0a, 0x60, 0xa2, 0x0a, 0xa1, 0x08,
	0x2a, 0x1b, 0x52, 0xc4, 0x99, 0x26, 0x48, 0x1c, 0x00, 0x15, 0x92, 0x1b, 0x48, 0x44, 0x8e, 0xbb,
	0x71, 0x17, 0x92, 0xdd, 0x60, 0x6f, 0xd4, 0x56, 0xe2, 0x21, 0x78, 0xac, 0x1e, 0x7b, 0xe4, 0x02,
	0x42, 0xc9, 0x8b, 0x20, 0xcf, 0x2c, 0xac, 0xff, 0x1a, 0x97, 0x53, 0xec, 0x99, 0xf9, 0x7e, 0xf6,
	0xcb, 0x78, 0xc9, 0x6d, 0x31, 0x54, 0x32, 0xf4, 0xfc, 0x11, 0x73, 0x03, 0x26, 0x58, 0xc4, 0x23,
	0x67, 0x12, 0x4a, 0x25, 0xa9, 0xad, 0x8e, 0x58, 0xc4, 0x4f, 0x04, 0x53, 0xc7, 0x32, 0xfc, 0xe2,
	0xc4, 0x8f, 0x43, 0xe5, 0xfc, 0x9b, 0xae, 0x6f, 0x05, 0x32, 0x90, 0x30, 0xea, 0xc6, 0x4f, 0x88,
	0xaa, 0x6f, 0x1b, 0xba, 0x89, 0x17, 0x7a, 0x63, 0xcd, 0x56, 0xbf, 0x6b, 0xea, 0x63, 0x2e, 0x54,
	0x3f, 0x64, 0x5f, 0xa7, 0x2c, 0x52, 0xba, 0x6b, 0x9b, 0xae, 0xe7, 0x2b, 0x2e, 0x45, 0xa6, 0xdf,
	0x34, 0x7d, 0x5f, 0x8e, 0x46, 0x0c, 0x67, 0xe4, 0xb1, 0x60, 0x61, 0x66, 0xf2, 0x9e, 0x99, 0xc4,
	0x9f, 0xbe, 0x2f, 0xc5, 0x90, 0x07, 0xf9, 0xb6, 0x16, 0x8a, 0x78, 0x20, 0x58, 0x98, 0x6f, 0x0f,
	0xb8, 0x38, 0x64, 0x87, 0xe9, 0xf6, 0x83, 0x0b, 0xd0, 0x69, 0x8d, 0x1d, 0x33, 0x15, 0x9d, 0x0a,
	0xbf, 0x5f, 0x20, 0xb4, 0xf3, 0x73, 0x85, 0xac, 0xbd, 0xc2, 0xb8, 0x7b, 0xca, 0x53, 0x8c, 0xbe,
	0x24, 0x55, 0xcc, 0xab, 0x66, 0x35, 0xac, 0xe6, 0x6a, 0xeb, 0xa1, 0xb3, 0x38, 0x7e, 0xe7, 0x1d,
	0x4c, 0xb7, 0x97, 0xcf, 0x7e, 0xdd, 0xaf, 0x74, 0x35, 0x96, 0x7e, 0x24, 0x1b, 0x71, 0xba, 0x5d,
	0x8c, 0xe4, 0x0d, 0x8f, 0x54, 0xed, 0x4a, 0x63, 0xa9, 0xb9, 0xda, 0x7a, 0x5c, 0x46, 0xf7, 0xd6,
	0xc0, 0x34, 0x67, 0x96, 0x89, 0x3e, 0x22, 0x9b, 0x89, 0x52, 0x47, 0x4e, 0x85, 0xaa, 0x2d, 0x35,
	0xac, 0xe6, 0x72, 0x37, 0x57, 0xa7, 0x01, 0xb9, 0x89, 0xc7, 0x4e, 0x5a, 0x59, 0x06, 0x2b, 0x7b,
	0x65, 0x56, 0xf6, 0x01, 0x78, 0x00, 0x2f, 0x69, 0x4b, 0x79, 0x4e, 0xea, 0x10, 0x9a, 0x2a, 0xa2,
	0xad, 0xab, 0x60, 0xab, 0xa0, 0x43, 0xbf, 0x91, 0xba, 0xd9, 0xa0, 0x83, 0x78, 0x81, 0x92, 0x0e,
	0xab, 0xe0, 0xf0, 0x79, 0x99, 0xc3, 0x4e, 0x21, 0x83, 0x36, 0xb9, 0x80, 0x9f, 0xbe, 0x20, 0x77,
	0x8a, 0xbb, 0x68, 0xfb, 0x1a, 0xd8, 0x5e, 0x34, 0x42, 0xdf, 0x93, 0xf5, 0xd4, 0x5e, 0xd7, 0x56,
	0x60, 0x5d, 0x76, 0xcb, 0x2c, 0x63, 0x9c, 0x1d, 0xc0, 0x74, 0xd7, 0x64, 0xe2, 0x8d, 0x7e, 0x22,
	0x9b, 0x18, 0x54, 0x0f, 0x36, 0x14, 0x82, 0xb8, 0x0e, 0x41, 0xec, 0x5e, 0xee, 0xaf, 0x42, 0x9c,
	0x3e, 0x7e, 0x8e, 0x2b, 0xe6, 0xc7, 0x8f, 0x29, 0xc1, 0x4f, 0x2e, 0xc7, 0xdf, 0x4e, 0xe0, 0xfe,
	0xf2, 0x67, 0xb9, 0xe8, 0x84, 0x6c, 0x27, 0x35, 0xf1, 0x54, 0xa0, 0xb2, 0x0e, 0x2a, 0xad, 0xff,
	0x39, 0x05, 0xa2, 0xb5, 0xd6, 0x05, 0xbc, 0xf4, 0x33, 0xd9, 0x8a, 0xbf, 0xec, 0xfd, 0x6c, 0x6a,
	0x37, 0x40, 0xef, 0x49, 0x99, 0x5e, 0x2f, 0x83, 0xd5, 0x6a, 0x85, 0x9c, 0xf4, 0x19, 0xb9, 0x95,
	0xad, 0xe3, 0xb2, 0x6c, 0xc0, 0xb2, 0x14, 0x37, 0xdb, 0xaf, 0xcf, 0x66, 0xb6, 0x75, 0x3e, 0xb3,
	0xad, 0xdf, 0x33, 0xdb, 0xfa, 0x3e, 0xb7, 0x2b, 0xe7, 0x73, 0xbb, 0xf2, 0x63, 0x6e, 0x57, 0x3e,
	0x3c, 0x0d, 0xb8, 0x3a, 0x9a, 0x0e, 0x1c, 0x5f, 0x8e, 0xdd, 0x94, 0x4f, 0x17, 0x7d, 0xba, 0x27,
	0xae, 0xb9, 0xbf, 0xd4, 0xe9, 0x84, 0x45, 0x83, 0x2a, 0xdc, 0x59, 0x7b, 0x7f, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xd8, 0x43, 0xe5, 0xb3, 0x2b, 0x06, 0x00, 0x00,
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
	if m.SyncActionSignerCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.SyncActionSignerCount))
		i--
		dAtA[i] = 0x78
	}
	if len(m.SyncActionSignerList) > 0 {
		for iNdEx := len(m.SyncActionSignerList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SyncActionSignerList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x72
		}
	}
	if len(m.ActionSignerConfigList) > 0 {
		for iNdEx := len(m.ActionSignerConfigList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ActionSignerConfigList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x6a
		}
	}
	if len(m.BindedSignerList) > 0 {
		for iNdEx := len(m.BindedSignerList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BindedSignerList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
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
	if len(m.BindedSignerList) > 0 {
		for _, e := range m.BindedSignerList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ActionSignerConfigList) > 0 {
		for _, e := range m.ActionSignerConfigList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.SyncActionSignerList) > 0 {
		for _, e := range m.SyncActionSignerList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.SyncActionSignerCount != 0 {
		n += 1 + sovGenesis(uint64(m.SyncActionSignerCount))
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
			m.ActionRequestList = append(m.ActionRequestList, ActionOracleRequest{})
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
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BindedSignerList", wireType)
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
			m.BindedSignerList = append(m.BindedSignerList, BindedSigner{})
			if err := m.BindedSignerList[len(m.BindedSignerList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActionSignerConfigList", wireType)
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
			m.ActionSignerConfigList = append(m.ActionSignerConfigList, ActionSignerConfig{})
			if err := m.ActionSignerConfigList[len(m.ActionSignerConfigList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SyncActionSignerList", wireType)
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
			m.SyncActionSignerList = append(m.SyncActionSignerList, SyncActionSigner{})
			if err := m.SyncActionSignerList[len(m.SyncActionSignerList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SyncActionSignerCount", wireType)
			}
			m.SyncActionSignerCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SyncActionSignerCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
