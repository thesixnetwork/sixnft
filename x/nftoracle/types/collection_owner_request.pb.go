// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftoracle/collection_owner_request.proto

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

type OriginContractParam struct {
	Chain           string    `protobuf:"bytes,1,opt,name=chain,proto3" json:"chain,omitempty"`
	ContractAddress string    `protobuf:"bytes,2,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	ContractOwner   string    `protobuf:"bytes,3,opt,name=contract_owner,json=contractOwner,proto3" json:"contract_owner,omitempty"`
	RequestExpire   time.Time `protobuf:"bytes,4,opt,name=request_expire,json=requestExpire,proto3,stdtime" json:"request_expire"`
}

func (m *OriginContractParam) Reset()         { *m = OriginContractParam{} }
func (m *OriginContractParam) String() string { return proto.CompactTextString(m) }
func (*OriginContractParam) ProtoMessage()    {}
func (*OriginContractParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b5e5fd2fa665471, []int{0}
}
func (m *OriginContractParam) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OriginContractParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OriginContractParam.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OriginContractParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OriginContractParam.Merge(m, src)
}
func (m *OriginContractParam) XXX_Size() int {
	return m.Size()
}
func (m *OriginContractParam) XXX_DiscardUnknown() {
	xxx_messageInfo_OriginContractParam.DiscardUnknown(m)
}

var xxx_messageInfo_OriginContractParam proto.InternalMessageInfo

func (m *OriginContractParam) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *OriginContractParam) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *OriginContractParam) GetContractOwner() string {
	if m != nil {
		return m.ContractOwner
	}
	return ""
}

func (m *OriginContractParam) GetRequestExpire() time.Time {
	if m != nil {
		return m.RequestExpire
	}
	return time.Time{}
}

type CollectionOwnerRequest struct {
	Id              uint64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	NftSchemaCode   string                `protobuf:"bytes,2,opt,name=nftSchemaCode,proto3" json:"nftSchemaCode,omitempty"`
	Signer          string                `protobuf:"bytes,3,opt,name=signer,proto3" json:"signer,omitempty"`
	RequiredConfirm uint64                `protobuf:"varint,4,opt,name=required_confirm,json=requiredConfirm,proto3" json:"required_confirm,omitempty"`
	Status          RequestStatus         `protobuf:"varint,5,opt,name=status,proto3,enum=thesixnetwork.sixnft.nftoracle.RequestStatus" json:"status,omitempty"`
	CurrentConfirm  uint64                `protobuf:"varint,6,opt,name=current_confirm,json=currentConfirm,proto3" json:"current_confirm,omitempty"`
	Confirmers      []string              `protobuf:"bytes,7,rep,name=confirmers,proto3" json:"confirmers,omitempty"`
	CreatedAt       time.Time             `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at"`
	ValidUntil      time.Time             `protobuf:"bytes,9,opt,name=valid_until,json=validUntil,proto3,stdtime" json:"valid_until"`
	ContractInfo    []*OriginContractInfo `protobuf:"bytes,10,rep,name=contract_info,json=contractInfo,proto3" json:"contract_info,omitempty"`
	ExpiredHeight   int64                 `protobuf:"varint,11,opt,name=expired_height,json=expiredHeight,proto3" json:"expired_height,omitempty"`
}

func (m *CollectionOwnerRequest) Reset()         { *m = CollectionOwnerRequest{} }
func (m *CollectionOwnerRequest) String() string { return proto.CompactTextString(m) }
func (*CollectionOwnerRequest) ProtoMessage()    {}
func (*CollectionOwnerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b5e5fd2fa665471, []int{1}
}
func (m *CollectionOwnerRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CollectionOwnerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CollectionOwnerRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CollectionOwnerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionOwnerRequest.Merge(m, src)
}
func (m *CollectionOwnerRequest) XXX_Size() int {
	return m.Size()
}
func (m *CollectionOwnerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionOwnerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionOwnerRequest proto.InternalMessageInfo

func (m *CollectionOwnerRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CollectionOwnerRequest) GetNftSchemaCode() string {
	if m != nil {
		return m.NftSchemaCode
	}
	return ""
}

func (m *CollectionOwnerRequest) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

func (m *CollectionOwnerRequest) GetRequiredConfirm() uint64 {
	if m != nil {
		return m.RequiredConfirm
	}
	return 0
}

func (m *CollectionOwnerRequest) GetStatus() RequestStatus {
	if m != nil {
		return m.Status
	}
	return RequestStatus_PENDING
}

func (m *CollectionOwnerRequest) GetCurrentConfirm() uint64 {
	if m != nil {
		return m.CurrentConfirm
	}
	return 0
}

func (m *CollectionOwnerRequest) GetConfirmers() []string {
	if m != nil {
		return m.Confirmers
	}
	return nil
}

func (m *CollectionOwnerRequest) GetCreatedAt() time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return time.Time{}
}

func (m *CollectionOwnerRequest) GetValidUntil() time.Time {
	if m != nil {
		return m.ValidUntil
	}
	return time.Time{}
}

func (m *CollectionOwnerRequest) GetContractInfo() []*OriginContractInfo {
	if m != nil {
		return m.ContractInfo
	}
	return nil
}

func (m *CollectionOwnerRequest) GetExpiredHeight() int64 {
	if m != nil {
		return m.ExpiredHeight
	}
	return 0
}

type OriginContractInfo struct {
	ContractOriginDataInfo *OriginContractParam `protobuf:"bytes,1,opt,name=contractOriginDataInfo,proto3" json:"contractOriginDataInfo,omitempty"`
	Hash                   []byte               `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Confirmers             []string             `protobuf:"bytes,3,rep,name=confirmers,proto3" json:"confirmers,omitempty"`
}

func (m *OriginContractInfo) Reset()         { *m = OriginContractInfo{} }
func (m *OriginContractInfo) String() string { return proto.CompactTextString(m) }
func (*OriginContractInfo) ProtoMessage()    {}
func (*OriginContractInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b5e5fd2fa665471, []int{2}
}
func (m *OriginContractInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OriginContractInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OriginContractInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OriginContractInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OriginContractInfo.Merge(m, src)
}
func (m *OriginContractInfo) XXX_Size() int {
	return m.Size()
}
func (m *OriginContractInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_OriginContractInfo.DiscardUnknown(m)
}

var xxx_messageInfo_OriginContractInfo proto.InternalMessageInfo

func (m *OriginContractInfo) GetContractOriginDataInfo() *OriginContractParam {
	if m != nil {
		return m.ContractOriginDataInfo
	}
	return nil
}

func (m *OriginContractInfo) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *OriginContractInfo) GetConfirmers() []string {
	if m != nil {
		return m.Confirmers
	}
	return nil
}

func init() {
	proto.RegisterType((*OriginContractParam)(nil), "thesixnetwork.sixnft.nftoracle.OriginContractParam")
	proto.RegisterType((*CollectionOwnerRequest)(nil), "thesixnetwork.sixnft.nftoracle.CollectionOwnerRequest")
	proto.RegisterType((*OriginContractInfo)(nil), "thesixnetwork.sixnft.nftoracle.OriginContractInfo")
}

func init() {
	proto.RegisterFile("nftoracle/collection_owner_request.proto", fileDescriptor_5b5e5fd2fa665471)
}

var fileDescriptor_5b5e5fd2fa665471 = []byte{
	// 595 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0x8e, 0x9b, 0x34, 0xbf, 0x66, 0xd3, 0xa4, 0x3f, 0x2d, 0x55, 0xb1, 0x7a, 0x70, 0xad, 0x0a,
	0x84, 0x39, 0x60, 0x8b, 0xf4, 0x09, 0xda, 0x50, 0x09, 0xd4, 0x43, 0x91, 0x0b, 0x42, 0xe2, 0x62,
	0x6d, 0xd6, 0x6b, 0x7b, 0x55, 0x7b, 0x37, 0xac, 0xd7, 0x34, 0xbc, 0x45, 0x5f, 0x84, 0x23, 0xef,
	0x50, 0x71, 0xea, 0x91, 0x13, 0xa0, 0xe4, 0x45, 0x90, 0xd7, 0x6b, 0xa7, 0x29, 0x02, 0x94, 0xdb,
	0xcc, 0xb7, 0xf3, 0xcd, 0x7c, 0xf3, 0xc7, 0x06, 0x0e, 0x8b, 0x24, 0x17, 0x08, 0xa7, 0xc4, 0xc3,
	0x3c, 0x4d, 0x09, 0x96, 0x94, 0xb3, 0x80, 0x5f, 0x31, 0x22, 0x02, 0x41, 0x3e, 0x14, 0x24, 0x97,
	0xee, 0x54, 0x70, 0xc9, 0xa1, 0x25, 0x13, 0x92, 0xd3, 0x19, 0x23, 0xf2, 0x8a, 0x8b, 0x4b, 0xb7,
	0x34, 0x23, 0xe9, 0x36, 0xf4, 0xfd, 0x87, 0xcb, 0x4c, 0x2b, 0xc4, 0xfd, 0xdd, 0x98, 0xc7, 0x5c,
	0x99, 0x5e, 0x69, 0x69, 0xf4, 0x20, 0xe6, 0x3c, 0x4e, 0x89, 0xa7, 0xbc, 0x49, 0x11, 0x79, 0x92,
	0x66, 0x24, 0x97, 0x28, 0x9b, 0x56, 0x01, 0x87, 0x5f, 0x0d, 0xf0, 0xe0, 0x5c, 0xd0, 0x98, 0xb2,
	0x31, 0x67, 0x52, 0x20, 0x2c, 0x5f, 0x23, 0x81, 0x32, 0xb8, 0x0b, 0x36, 0x71, 0x82, 0x28, 0x33,
	0x0d, 0xdb, 0x70, 0x7a, 0x7e, 0xe5, 0xc0, 0xa7, 0xe0, 0x7f, 0xac, 0xc3, 0x02, 0x14, 0x86, 0x82,
	0xe4, 0xb9, 0xb9, 0xa1, 0x02, 0x76, 0x6a, 0xfc, 0xb8, 0x82, 0xe1, 0x63, 0x30, 0x6c, 0x42, 0x55,
	0xa3, 0x66, 0x5b, 0x05, 0x0e, 0x6a, 0xf4, 0xbc, 0x04, 0xe1, 0x19, 0x18, 0xea, 0x3e, 0x02, 0x32,
	0x9b, 0x52, 0x41, 0xcc, 0x8e, 0x6d, 0x38, 0xfd, 0xd1, 0xbe, 0x5b, 0x29, 0x77, 0x6b, 0xe5, 0xee,
	0x9b, 0x5a, 0xf9, 0xc9, 0xd6, 0xcd, 0xf7, 0x83, 0xd6, 0xf5, 0x8f, 0x03, 0xc3, 0x1f, 0x68, 0xee,
	0xa9, 0xa2, 0x1e, 0x7e, 0xee, 0x80, 0xbd, 0x71, 0x33, 0x5f, 0x55, 0xc0, 0xaf, 0x02, 0xe0, 0x10,
	0x6c, 0xd0, 0x50, 0x35, 0xd3, 0xf1, 0x37, 0x68, 0x08, 0x1f, 0x81, 0x01, 0x8b, 0xe4, 0x05, 0x4e,
	0x48, 0x86, 0xc6, 0x3c, 0x24, 0xba, 0x8d, 0x55, 0x10, 0xee, 0x81, 0x6e, 0x4e, 0xe3, 0xa5, 0x78,
	0xed, 0x95, 0x73, 0x28, 0x2b, 0x53, 0x41, 0xc2, 0x00, 0x73, 0x16, 0x51, 0x91, 0x29, 0xdd, 0x1d,
	0x7f, 0xa7, 0xc6, 0xc7, 0x15, 0x0c, 0x4f, 0x41, 0x37, 0x97, 0x48, 0x16, 0xb9, 0xb9, 0x69, 0x1b,
	0xce, 0x70, 0xf4, 0xcc, 0xfd, 0xfb, 0x86, 0x5d, 0xad, 0xf8, 0x42, 0x91, 0x7c, 0x4d, 0x86, 0x4f,
	0xc0, 0x0e, 0x2e, 0x84, 0x20, 0x4c, 0x36, 0x05, 0xbb, 0xaa, 0xe0, 0x50, 0xc3, 0x75, 0x3d, 0x0b,
	0x00, 0x1d, 0x40, 0x44, 0x6e, 0xfe, 0x67, 0xb7, 0x9d, 0x9e, 0x7f, 0x07, 0x81, 0x63, 0x00, 0xb0,
	0x20, 0x48, 0x92, 0x30, 0x40, 0xd2, 0xdc, 0x5a, 0x63, 0xd8, 0x3d, 0xcd, 0x3b, 0x96, 0xf0, 0x14,
	0xf4, 0x3f, 0xa2, 0x94, 0x86, 0x41, 0xc1, 0x24, 0x4d, 0xcd, 0xde, 0x1a, 0x59, 0x80, 0x22, 0xbe,
	0x2d, 0x79, 0xf0, 0x1d, 0x68, 0xae, 0x21, 0xa0, 0x2c, 0xe2, 0x26, 0xb0, 0xdb, 0x4e, 0x7f, 0x34,
	0xfa, 0xd7, 0x88, 0x56, 0x0f, 0xf6, 0x15, 0x8b, 0xb8, 0xbf, 0x8d, 0xef, 0x78, 0xe5, 0xf1, 0x55,
	0xd7, 0x14, 0x06, 0x09, 0xa1, 0x71, 0x22, 0xcd, 0xbe, 0x6d, 0x38, 0x6d, 0x7f, 0xa0, 0xd1, 0x97,
	0x0a, 0x3c, 0xfc, 0x62, 0x00, 0xf8, 0x7b, 0x2e, 0x78, 0x09, 0xf6, 0x9a, 0x23, 0x55, 0xaf, 0x2f,
	0x90, 0x44, 0xe5, 0x8b, 0xba, 0x9f, 0xfe, 0xe8, 0x68, 0x3d, 0x7d, 0xea, 0x83, 0xf2, 0xff, 0x90,
	0x12, 0x42, 0xd0, 0x49, 0x50, 0x9e, 0xa8, 0xfb, 0xdb, 0xf6, 0x95, 0x7d, 0x6f, 0x87, 0xed, 0xfb,
	0x3b, 0x3c, 0x39, 0xbb, 0x99, 0x5b, 0xc6, 0xed, 0xdc, 0x32, 0x7e, 0xce, 0x2d, 0xe3, 0x7a, 0x61,
	0xb5, 0x6e, 0x17, 0x56, 0xeb, 0xdb, 0xc2, 0x6a, 0xbd, 0x7f, 0x1e, 0x53, 0x99, 0x14, 0x13, 0x17,
	0xf3, 0xcc, 0x5b, 0x11, 0xe9, 0x55, 0x22, 0xbd, 0x99, 0xb7, 0xfc, 0x81, 0xc8, 0x4f, 0x53, 0x92,
	0x4f, 0xba, 0x6a, 0x5d, 0x47, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb3, 0xbb, 0x97, 0x78, 0xa4,
	0x04, 0x00, 0x00,
}

func (m *OriginContractParam) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OriginContractParam) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OriginContractParam) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.RequestExpire, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.RequestExpire):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintCollectionOwnerRequest(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	if len(m.ContractOwner) > 0 {
		i -= len(m.ContractOwner)
		copy(dAtA[i:], m.ContractOwner)
		i = encodeVarintCollectionOwnerRequest(dAtA, i, uint64(len(m.ContractOwner)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintCollectionOwnerRequest(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintCollectionOwnerRequest(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CollectionOwnerRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CollectionOwnerRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CollectionOwnerRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ExpiredHeight != 0 {
		i = encodeVarintCollectionOwnerRequest(dAtA, i, uint64(m.ExpiredHeight))
		i--
		dAtA[i] = 0x58
	}
	if len(m.ContractInfo) > 0 {
		for iNdEx := len(m.ContractInfo) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ContractInfo[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCollectionOwnerRequest(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.ValidUntil, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.ValidUntil):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintCollectionOwnerRequest(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x4a
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CreatedAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintCollectionOwnerRequest(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x42
	if len(m.Confirmers) > 0 {
		for iNdEx := len(m.Confirmers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Confirmers[iNdEx])
			copy(dAtA[i:], m.Confirmers[iNdEx])
			i = encodeVarintCollectionOwnerRequest(dAtA, i, uint64(len(m.Confirmers[iNdEx])))
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.CurrentConfirm != 0 {
		i = encodeVarintCollectionOwnerRequest(dAtA, i, uint64(m.CurrentConfirm))
		i--
		dAtA[i] = 0x30
	}
	if m.Status != 0 {
		i = encodeVarintCollectionOwnerRequest(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x28
	}
	if m.RequiredConfirm != 0 {
		i = encodeVarintCollectionOwnerRequest(dAtA, i, uint64(m.RequiredConfirm))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintCollectionOwnerRequest(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.NftSchemaCode) > 0 {
		i -= len(m.NftSchemaCode)
		copy(dAtA[i:], m.NftSchemaCode)
		i = encodeVarintCollectionOwnerRequest(dAtA, i, uint64(len(m.NftSchemaCode)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintCollectionOwnerRequest(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *OriginContractInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OriginContractInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OriginContractInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Confirmers) > 0 {
		for iNdEx := len(m.Confirmers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Confirmers[iNdEx])
			copy(dAtA[i:], m.Confirmers[iNdEx])
			i = encodeVarintCollectionOwnerRequest(dAtA, i, uint64(len(m.Confirmers[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintCollectionOwnerRequest(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x12
	}
	if m.ContractOriginDataInfo != nil {
		{
			size, err := m.ContractOriginDataInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCollectionOwnerRequest(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCollectionOwnerRequest(dAtA []byte, offset int, v uint64) int {
	offset -= sovCollectionOwnerRequest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OriginContractParam) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovCollectionOwnerRequest(uint64(l))
	}
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovCollectionOwnerRequest(uint64(l))
	}
	l = len(m.ContractOwner)
	if l > 0 {
		n += 1 + l + sovCollectionOwnerRequest(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.RequestExpire)
	n += 1 + l + sovCollectionOwnerRequest(uint64(l))
	return n
}

func (m *CollectionOwnerRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovCollectionOwnerRequest(uint64(m.Id))
	}
	l = len(m.NftSchemaCode)
	if l > 0 {
		n += 1 + l + sovCollectionOwnerRequest(uint64(l))
	}
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovCollectionOwnerRequest(uint64(l))
	}
	if m.RequiredConfirm != 0 {
		n += 1 + sovCollectionOwnerRequest(uint64(m.RequiredConfirm))
	}
	if m.Status != 0 {
		n += 1 + sovCollectionOwnerRequest(uint64(m.Status))
	}
	if m.CurrentConfirm != 0 {
		n += 1 + sovCollectionOwnerRequest(uint64(m.CurrentConfirm))
	}
	if len(m.Confirmers) > 0 {
		for _, s := range m.Confirmers {
			l = len(s)
			n += 1 + l + sovCollectionOwnerRequest(uint64(l))
		}
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt)
	n += 1 + l + sovCollectionOwnerRequest(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.ValidUntil)
	n += 1 + l + sovCollectionOwnerRequest(uint64(l))
	if len(m.ContractInfo) > 0 {
		for _, e := range m.ContractInfo {
			l = e.Size()
			n += 1 + l + sovCollectionOwnerRequest(uint64(l))
		}
	}
	if m.ExpiredHeight != 0 {
		n += 1 + sovCollectionOwnerRequest(uint64(m.ExpiredHeight))
	}
	return n
}

func (m *OriginContractInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ContractOriginDataInfo != nil {
		l = m.ContractOriginDataInfo.Size()
		n += 1 + l + sovCollectionOwnerRequest(uint64(l))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovCollectionOwnerRequest(uint64(l))
	}
	if len(m.Confirmers) > 0 {
		for _, s := range m.Confirmers {
			l = len(s)
			n += 1 + l + sovCollectionOwnerRequest(uint64(l))
		}
	}
	return n
}

func sovCollectionOwnerRequest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCollectionOwnerRequest(x uint64) (n int) {
	return sovCollectionOwnerRequest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OriginContractParam) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCollectionOwnerRequest
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
			return fmt.Errorf("proto: OriginContractParam: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OriginContractParam: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectionOwnerRequest
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
				return ErrInvalidLengthCollectionOwnerRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollectionOwnerRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectionOwnerRequest
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
				return ErrInvalidLengthCollectionOwnerRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollectionOwnerRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractOwner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectionOwnerRequest
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
				return ErrInvalidLengthCollectionOwnerRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollectionOwnerRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractOwner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestExpire", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectionOwnerRequest
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
				return ErrInvalidLengthCollectionOwnerRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCollectionOwnerRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.RequestExpire, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCollectionOwnerRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCollectionOwnerRequest
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
func (m *CollectionOwnerRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCollectionOwnerRequest
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
			return fmt.Errorf("proto: CollectionOwnerRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CollectionOwnerRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectionOwnerRequest
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
					return ErrIntOverflowCollectionOwnerRequest
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
				return ErrInvalidLengthCollectionOwnerRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollectionOwnerRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NftSchemaCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectionOwnerRequest
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
				return ErrInvalidLengthCollectionOwnerRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollectionOwnerRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequiredConfirm", wireType)
			}
			m.RequiredConfirm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectionOwnerRequest
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
					return ErrIntOverflowCollectionOwnerRequest
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
					return ErrIntOverflowCollectionOwnerRequest
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
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectionOwnerRequest
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
				return ErrInvalidLengthCollectionOwnerRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollectionOwnerRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Confirmers = append(m.Confirmers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectionOwnerRequest
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
				return ErrInvalidLengthCollectionOwnerRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCollectionOwnerRequest
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
					return ErrIntOverflowCollectionOwnerRequest
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
				return ErrInvalidLengthCollectionOwnerRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCollectionOwnerRequest
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
				return fmt.Errorf("proto: wrong wireType = %d for field ContractInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectionOwnerRequest
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
				return ErrInvalidLengthCollectionOwnerRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCollectionOwnerRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractInfo = append(m.ContractInfo, &OriginContractInfo{})
			if err := m.ContractInfo[len(m.ContractInfo)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
					return ErrIntOverflowCollectionOwnerRequest
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
			skippy, err := skipCollectionOwnerRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCollectionOwnerRequest
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
func (m *OriginContractInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCollectionOwnerRequest
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
			return fmt.Errorf("proto: OriginContractInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OriginContractInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractOriginDataInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectionOwnerRequest
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
				return ErrInvalidLengthCollectionOwnerRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCollectionOwnerRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ContractOriginDataInfo == nil {
				m.ContractOriginDataInfo = &OriginContractParam{}
			}
			if err := m.ContractOriginDataInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
					return ErrIntOverflowCollectionOwnerRequest
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
				return ErrInvalidLengthCollectionOwnerRequest
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCollectionOwnerRequest
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
					return ErrIntOverflowCollectionOwnerRequest
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
				return ErrInvalidLengthCollectionOwnerRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollectionOwnerRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Confirmers = append(m.Confirmers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCollectionOwnerRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCollectionOwnerRequest
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
func skipCollectionOwnerRequest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCollectionOwnerRequest
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
					return 0, ErrIntOverflowCollectionOwnerRequest
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
					return 0, ErrIntOverflowCollectionOwnerRequest
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
				return 0, ErrInvalidLengthCollectionOwnerRequest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCollectionOwnerRequest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCollectionOwnerRequest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCollectionOwnerRequest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCollectionOwnerRequest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCollectionOwnerRequest = fmt.Errorf("proto: unexpected end of group")
)
