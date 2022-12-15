// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftoracle/action_request.proto

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

type ActionParam struct {
	NftSchemaCode string    `protobuf:"bytes,1,opt,name=nft_schema_code,json=nftSchemaCode,proto3" json:"nft_schema_code,omitempty"`
	TokenId       string    `protobuf:"bytes,2,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	Action        string    `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	RefId         string    `protobuf:"bytes,4,opt,name=ref_id,json=refId,proto3" json:"ref_id,omitempty"`
	ExpiredAt     time.Time `protobuf:"bytes,5,opt,name=expired_at,json=expiredAt,proto3,stdtime" json:"expired_at"`
	OnBehalfOf    string    `protobuf:"bytes,6,opt,name=on_behalf_of,json=onBehalfOf,proto3" json:"on_behalf_of,omitempty"`
}

func (m *ActionParam) Reset()         { *m = ActionParam{} }
func (m *ActionParam) String() string { return proto.CompactTextString(m) }
func (*ActionParam) ProtoMessage()    {}
func (*ActionParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_36b7260720dc8e70, []int{0}
}
func (m *ActionParam) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ActionParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ActionParam.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ActionParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionParam.Merge(m, src)
}
func (m *ActionParam) XXX_Size() int {
	return m.Size()
}
func (m *ActionParam) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionParam.DiscardUnknown(m)
}

var xxx_messageInfo_ActionParam proto.InternalMessageInfo

func (m *ActionParam) GetNftSchemaCode() string {
	if m != nil {
		return m.NftSchemaCode
	}
	return ""
}

func (m *ActionParam) GetTokenId() string {
	if m != nil {
		return m.TokenId
	}
	return ""
}

func (m *ActionParam) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *ActionParam) GetRefId() string {
	if m != nil {
		return m.RefId
	}
	return ""
}

func (m *ActionParam) GetExpiredAt() time.Time {
	if m != nil {
		return m.ExpiredAt
	}
	return time.Time{}
}

func (m *ActionParam) GetOnBehalfOf() string {
	if m != nil {
		return m.OnBehalfOf
	}
	return ""
}

type ActionRequest struct {
	Id                    uint64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	NftSchemaCode         string        `protobuf:"bytes,2,opt,name=nft_schema_code,json=nftSchemaCode,proto3" json:"nft_schema_code,omitempty"`
	TokenId               string        `protobuf:"bytes,3,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	Action                string        `protobuf:"bytes,4,opt,name=action,proto3" json:"action,omitempty"`
	Caller                string        `protobuf:"bytes,5,opt,name=caller,proto3" json:"caller,omitempty"`
	RefId                 string        `protobuf:"bytes,6,opt,name=ref_id,json=refId,proto3" json:"ref_id,omitempty"`
	RequiredConfirm       uint64        `protobuf:"varint,7,opt,name=required_confirm,json=requiredConfirm,proto3" json:"required_confirm,omitempty"`
	Status                RequestStatus `protobuf:"varint,8,opt,name=status,proto3,enum=thesixnetwork.sixnft.nftoracle.RequestStatus" json:"status,omitempty"`
	CurrentConfirm        uint64        `protobuf:"varint,9,opt,name=current_confirm,json=currentConfirm,proto3" json:"current_confirm,omitempty"`
	Confirmers            []string      `protobuf:"bytes,10,rep,name=confirmers,proto3" json:"confirmers,omitempty"`
	CreatedAt             time.Time     `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at"`
	ValidUntil            time.Time     `protobuf:"bytes,12,opt,name=valid_until,json=validUntil,proto3,stdtime" json:"valid_until"`
	DataHashes            []*DataHash   `protobuf:"bytes,13,rep,name=data_hashes,json=dataHashes,proto3" json:"data_hashes,omitempty"`
	ExpiredHeight         int64         `protobuf:"varint,14,opt,name=expired_height,json=expiredHeight,proto3" json:"expired_height,omitempty"`
	ExecutionErrorMessage string        `protobuf:"bytes,15,opt,name=execution_error_message,json=executionErrorMessage,proto3" json:"execution_error_message,omitempty"`
}

func (m *ActionRequest) Reset()         { *m = ActionRequest{} }
func (m *ActionRequest) String() string { return proto.CompactTextString(m) }
func (*ActionRequest) ProtoMessage()    {}
func (*ActionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_36b7260720dc8e70, []int{1}
}
func (m *ActionRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ActionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ActionRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ActionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionRequest.Merge(m, src)
}
func (m *ActionRequest) XXX_Size() int {
	return m.Size()
}
func (m *ActionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ActionRequest proto.InternalMessageInfo

func (m *ActionRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ActionRequest) GetNftSchemaCode() string {
	if m != nil {
		return m.NftSchemaCode
	}
	return ""
}

func (m *ActionRequest) GetTokenId() string {
	if m != nil {
		return m.TokenId
	}
	return ""
}

func (m *ActionRequest) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *ActionRequest) GetCaller() string {
	if m != nil {
		return m.Caller
	}
	return ""
}

func (m *ActionRequest) GetRefId() string {
	if m != nil {
		return m.RefId
	}
	return ""
}

func (m *ActionRequest) GetRequiredConfirm() uint64 {
	if m != nil {
		return m.RequiredConfirm
	}
	return 0
}

func (m *ActionRequest) GetStatus() RequestStatus {
	if m != nil {
		return m.Status
	}
	return RequestStatus_PENDING
}

func (m *ActionRequest) GetCurrentConfirm() uint64 {
	if m != nil {
		return m.CurrentConfirm
	}
	return 0
}

func (m *ActionRequest) GetConfirmers() []string {
	if m != nil {
		return m.Confirmers
	}
	return nil
}

func (m *ActionRequest) GetCreatedAt() time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return time.Time{}
}

func (m *ActionRequest) GetValidUntil() time.Time {
	if m != nil {
		return m.ValidUntil
	}
	return time.Time{}
}

func (m *ActionRequest) GetDataHashes() []*DataHash {
	if m != nil {
		return m.DataHashes
	}
	return nil
}

func (m *ActionRequest) GetExpiredHeight() int64 {
	if m != nil {
		return m.ExpiredHeight
	}
	return 0
}

func (m *ActionRequest) GetExecutionErrorMessage() string {
	if m != nil {
		return m.ExecutionErrorMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*ActionParam)(nil), "thesixnetwork.sixnft.nftoracle.ActionParam")
	proto.RegisterType((*ActionRequest)(nil), "thesixnetwork.sixnft.nftoracle.ActionRequest")
}

func init() { proto.RegisterFile("nftoracle/action_request.proto", fileDescriptor_36b7260720dc8e70) }

var fileDescriptor_36b7260720dc8e70 = []byte{
	// 609 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4d, 0x4f, 0xdc, 0x48,
	0x10, 0x1d, 0xcf, 0x80, 0x61, 0x6a, 0x98, 0x99, 0x55, 0x6b, 0x01, 0x2f, 0x07, 0x63, 0x21, 0xed,
	0xee, 0xe4, 0x10, 0x5b, 0x21, 0x52, 0xee, 0x40, 0x90, 0x40, 0x51, 0x94, 0xc8, 0x24, 0x97, 0x5c,
	0xac, 0x1e, 0xbb, 0xfc, 0x21, 0x6c, 0xf7, 0xa4, 0xbb, 0x9c, 0x4c, 0xfe, 0x05, 0xe7, 0xfc, 0x22,
	0x8e, 0x1c, 0x73, 0x4a, 0x22, 0x90, 0xf2, 0x3b, 0x22, 0xb7, 0x3d, 0x7c, 0x48, 0x49, 0x10, 0xb7,
	0xaa, 0x57, 0xfd, 0xba, 0xea, 0xbd, 0x6a, 0x1b, 0xec, 0x32, 0x26, 0x21, 0x79, 0x98, 0xa3, 0xc7,
	0x43, 0xca, 0x44, 0x19, 0x48, 0x7c, 0x5f, 0xa1, 0x22, 0x77, 0x26, 0x05, 0x09, 0x66, 0x53, 0x8a,
	0x2a, 0x9b, 0x97, 0x48, 0x1f, 0x85, 0x3c, 0x75, 0xeb, 0x30, 0x26, 0xf7, 0x9a, 0xb4, 0xb5, 0x79,
	0xc3, 0x17, 0x33, 0x2c, 0x15, 0xf2, 0x86, 0x78, 0xbb, 0x70, 0xe7, 0xc6, 0xad, 0xbf, 0x13, 0x91,
	0x08, 0x1d, 0x7a, 0x75, 0xd4, 0xa2, 0xdb, 0x89, 0x10, 0x49, 0x8e, 0x9e, 0xce, 0xa6, 0x55, 0xec,
	0x51, 0x56, 0xa0, 0x22, 0x5e, 0xcc, 0x9a, 0x03, 0x3b, 0x3f, 0x0c, 0x18, 0xec, 0xe9, 0x09, 0x5f,
	0x73, 0xc9, 0x0b, 0xf6, 0x1f, 0x8c, 0xcb, 0x98, 0x02, 0x15, 0xa6, 0x58, 0xf0, 0x20, 0x14, 0x11,
	0x5a, 0x86, 0x63, 0x4c, 0xfa, 0xfe, 0xb0, 0x8c, 0xe9, 0x44, 0xa3, 0x07, 0x22, 0x42, 0xf6, 0x0f,
	0xac, 0x92, 0x38, 0xc5, 0x32, 0xc8, 0x22, 0xab, 0xab, 0x0f, 0xac, 0xe8, 0xfc, 0x38, 0x62, 0x1b,
	0x60, 0x36, 0x9a, 0xad, 0x9e, 0x2e, 0xb4, 0x19, 0x5b, 0x07, 0x53, 0x62, 0x5c, 0x13, 0x96, 0x34,
	0xbe, 0x2c, 0x31, 0x3e, 0x8e, 0xd8, 0x01, 0x00, 0xce, 0x67, 0x99, 0xc4, 0x28, 0xe0, 0x64, 0x2d,
	0x3b, 0xc6, 0x64, 0xb0, 0xbb, 0xe5, 0x36, 0x73, 0xbb, 0x8b, 0xb9, 0xdd, 0x37, 0x8b, 0xb9, 0xf7,
	0x57, 0xcf, 0xbf, 0x6e, 0x77, 0xce, 0xbe, 0x6d, 0x1b, 0x7e, 0xbf, 0xe5, 0xed, 0x11, 0x73, 0x60,
	0x4d, 0x94, 0xc1, 0x14, 0x53, 0x9e, 0xc7, 0x81, 0x88, 0x2d, 0x53, 0x77, 0x00, 0x51, 0xee, 0x6b,
	0xe8, 0x55, 0xbc, 0xf3, 0x79, 0x19, 0x86, 0x8d, 0x50, 0xbf, 0xf1, 0x8d, 0x8d, 0xa0, 0x9b, 0x45,
	0x5a, 0xdd, 0x92, 0xdf, 0xcd, 0xa2, 0x5f, 0x49, 0xef, 0xde, 0x27, 0xbd, 0xf7, 0x3b, 0xe9, 0x4b,
	0x77, 0xa4, 0x6f, 0x80, 0x19, 0xf2, 0x3c, 0x47, 0xa9, 0xf5, 0xf5, 0xfd, 0x36, 0xbb, 0x65, 0x89,
	0x79, 0xdb, 0x92, 0x47, 0xf0, 0x57, 0xbd, 0x5c, 0xed, 0x49, 0x28, 0xca, 0x38, 0x93, 0x85, 0xb5,
	0xa2, 0xe7, 0x1c, 0x2f, 0xf0, 0x83, 0x06, 0x66, 0x87, 0x60, 0x2a, 0xe2, 0x54, 0x29, 0x6b, 0xd5,
	0x31, 0x26, 0xa3, 0xdd, 0xc7, 0xee, 0x9f, 0x5f, 0x96, 0xdb, 0xaa, 0x3f, 0xd1, 0x24, 0xbf, 0x25,
	0xb3, 0xff, 0x61, 0x1c, 0x56, 0x52, 0x62, 0x49, 0xd7, 0x0d, 0xfb, 0xba, 0xe1, 0xa8, 0x85, 0x17,
	0xfd, 0x6c, 0x80, 0xf6, 0x00, 0x4a, 0x65, 0x81, 0xd3, 0xab, 0x6d, 0xbe, 0x41, 0xea, 0x6d, 0x86,
	0x12, 0x39, 0x35, 0xdb, 0x1c, 0x3c, 0x64, 0x9b, 0x2d, 0x6f, 0x8f, 0xd8, 0x21, 0x0c, 0x3e, 0xf0,
	0x3c, 0x8b, 0x82, 0xaa, 0xa4, 0x2c, 0xb7, 0xd6, 0x1e, 0x70, 0x0b, 0x68, 0xe2, 0xdb, 0x9a, 0xc7,
	0x8e, 0x61, 0x10, 0x71, 0xe2, 0x41, 0xca, 0x55, 0x8a, 0xca, 0x1a, 0x3a, 0xbd, 0xc9, 0x60, 0x77,
	0x72, 0x9f, 0x41, 0xcf, 0x39, 0xf1, 0x23, 0xae, 0x52, 0x1f, 0xa2, 0x36, 0x42, 0xc5, 0xfe, 0x85,
	0xd1, 0xe2, 0x91, 0xa6, 0x98, 0x25, 0x29, 0x59, 0x23, 0xc7, 0x98, 0xf4, 0xfc, 0x61, 0x8b, 0x1e,
	0x69, 0x90, 0x3d, 0x83, 0x4d, 0x9c, 0x63, 0x58, 0xe9, 0x2f, 0x1e, 0xa5, 0x14, 0x32, 0x28, 0x50,
	0x29, 0x9e, 0xa0, 0x35, 0xd6, 0x0b, 0x5e, 0xbf, 0x2e, 0x1f, 0xd6, 0xd5, 0x97, 0x4d, 0x71, 0xff,
	0xc5, 0xf9, 0xa5, 0x6d, 0x5c, 0x5c, 0xda, 0xc6, 0xf7, 0x4b, 0xdb, 0x38, 0xbb, 0xb2, 0x3b, 0x17,
	0x57, 0x76, 0xe7, 0xcb, 0x95, 0xdd, 0x79, 0xf7, 0x24, 0xc9, 0x28, 0xad, 0xa6, 0x6e, 0x28, 0x0a,
	0xef, 0xce, 0xe0, 0x5e, 0x33, 0xb8, 0x37, 0xf7, 0x6e, 0xfe, 0x08, 0xf4, 0x69, 0x86, 0x6a, 0x6a,
	0x6a, 0x83, 0x9e, 0xfe, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xe9, 0x2b, 0x8f, 0x92, 0x84, 0x04, 0x00,
	0x00,
}

func (m *ActionParam) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ActionParam) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ActionParam) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.OnBehalfOf) > 0 {
		i -= len(m.OnBehalfOf)
		copy(dAtA[i:], m.OnBehalfOf)
		i = encodeVarintActionRequest(dAtA, i, uint64(len(m.OnBehalfOf)))
		i--
		dAtA[i] = 0x32
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.ExpiredAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.ExpiredAt):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintActionRequest(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x2a
	if len(m.RefId) > 0 {
		i -= len(m.RefId)
		copy(dAtA[i:], m.RefId)
		i = encodeVarintActionRequest(dAtA, i, uint64(len(m.RefId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Action) > 0 {
		i -= len(m.Action)
		copy(dAtA[i:], m.Action)
		i = encodeVarintActionRequest(dAtA, i, uint64(len(m.Action)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.TokenId) > 0 {
		i -= len(m.TokenId)
		copy(dAtA[i:], m.TokenId)
		i = encodeVarintActionRequest(dAtA, i, uint64(len(m.TokenId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.NftSchemaCode) > 0 {
		i -= len(m.NftSchemaCode)
		copy(dAtA[i:], m.NftSchemaCode)
		i = encodeVarintActionRequest(dAtA, i, uint64(len(m.NftSchemaCode)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ActionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ActionRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ActionRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ExecutionErrorMessage) > 0 {
		i -= len(m.ExecutionErrorMessage)
		copy(dAtA[i:], m.ExecutionErrorMessage)
		i = encodeVarintActionRequest(dAtA, i, uint64(len(m.ExecutionErrorMessage)))
		i--
		dAtA[i] = 0x7a
	}
	if m.ExpiredHeight != 0 {
		i = encodeVarintActionRequest(dAtA, i, uint64(m.ExpiredHeight))
		i--
		dAtA[i] = 0x70
	}
	if len(m.DataHashes) > 0 {
		for iNdEx := len(m.DataHashes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DataHashes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintActionRequest(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x6a
		}
	}
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.ValidUntil, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.ValidUntil):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintActionRequest(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x62
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CreatedAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintActionRequest(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x5a
	if len(m.Confirmers) > 0 {
		for iNdEx := len(m.Confirmers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Confirmers[iNdEx])
			copy(dAtA[i:], m.Confirmers[iNdEx])
			i = encodeVarintActionRequest(dAtA, i, uint64(len(m.Confirmers[iNdEx])))
			i--
			dAtA[i] = 0x52
		}
	}
	if m.CurrentConfirm != 0 {
		i = encodeVarintActionRequest(dAtA, i, uint64(m.CurrentConfirm))
		i--
		dAtA[i] = 0x48
	}
	if m.Status != 0 {
		i = encodeVarintActionRequest(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x40
	}
	if m.RequiredConfirm != 0 {
		i = encodeVarintActionRequest(dAtA, i, uint64(m.RequiredConfirm))
		i--
		dAtA[i] = 0x38
	}
	if len(m.RefId) > 0 {
		i -= len(m.RefId)
		copy(dAtA[i:], m.RefId)
		i = encodeVarintActionRequest(dAtA, i, uint64(len(m.RefId)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Caller) > 0 {
		i -= len(m.Caller)
		copy(dAtA[i:], m.Caller)
		i = encodeVarintActionRequest(dAtA, i, uint64(len(m.Caller)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Action) > 0 {
		i -= len(m.Action)
		copy(dAtA[i:], m.Action)
		i = encodeVarintActionRequest(dAtA, i, uint64(len(m.Action)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.TokenId) > 0 {
		i -= len(m.TokenId)
		copy(dAtA[i:], m.TokenId)
		i = encodeVarintActionRequest(dAtA, i, uint64(len(m.TokenId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.NftSchemaCode) > 0 {
		i -= len(m.NftSchemaCode)
		copy(dAtA[i:], m.NftSchemaCode)
		i = encodeVarintActionRequest(dAtA, i, uint64(len(m.NftSchemaCode)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintActionRequest(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintActionRequest(dAtA []byte, offset int, v uint64) int {
	offset -= sovActionRequest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ActionParam) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NftSchemaCode)
	if l > 0 {
		n += 1 + l + sovActionRequest(uint64(l))
	}
	l = len(m.TokenId)
	if l > 0 {
		n += 1 + l + sovActionRequest(uint64(l))
	}
	l = len(m.Action)
	if l > 0 {
		n += 1 + l + sovActionRequest(uint64(l))
	}
	l = len(m.RefId)
	if l > 0 {
		n += 1 + l + sovActionRequest(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.ExpiredAt)
	n += 1 + l + sovActionRequest(uint64(l))
	l = len(m.OnBehalfOf)
	if l > 0 {
		n += 1 + l + sovActionRequest(uint64(l))
	}
	return n
}

func (m *ActionRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovActionRequest(uint64(m.Id))
	}
	l = len(m.NftSchemaCode)
	if l > 0 {
		n += 1 + l + sovActionRequest(uint64(l))
	}
	l = len(m.TokenId)
	if l > 0 {
		n += 1 + l + sovActionRequest(uint64(l))
	}
	l = len(m.Action)
	if l > 0 {
		n += 1 + l + sovActionRequest(uint64(l))
	}
	l = len(m.Caller)
	if l > 0 {
		n += 1 + l + sovActionRequest(uint64(l))
	}
	l = len(m.RefId)
	if l > 0 {
		n += 1 + l + sovActionRequest(uint64(l))
	}
	if m.RequiredConfirm != 0 {
		n += 1 + sovActionRequest(uint64(m.RequiredConfirm))
	}
	if m.Status != 0 {
		n += 1 + sovActionRequest(uint64(m.Status))
	}
	if m.CurrentConfirm != 0 {
		n += 1 + sovActionRequest(uint64(m.CurrentConfirm))
	}
	if len(m.Confirmers) > 0 {
		for _, s := range m.Confirmers {
			l = len(s)
			n += 1 + l + sovActionRequest(uint64(l))
		}
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt)
	n += 1 + l + sovActionRequest(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.ValidUntil)
	n += 1 + l + sovActionRequest(uint64(l))
	if len(m.DataHashes) > 0 {
		for _, e := range m.DataHashes {
			l = e.Size()
			n += 1 + l + sovActionRequest(uint64(l))
		}
	}
	if m.ExpiredHeight != 0 {
		n += 1 + sovActionRequest(uint64(m.ExpiredHeight))
	}
	l = len(m.ExecutionErrorMessage)
	if l > 0 {
		n += 1 + l + sovActionRequest(uint64(l))
	}
	return n
}

func sovActionRequest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozActionRequest(x uint64) (n int) {
	return sovActionRequest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ActionParam) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowActionRequest
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
			return fmt.Errorf("proto: ActionParam: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ActionParam: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftSchemaCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionRequest
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
				return ErrInvalidLengthActionRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActionRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NftSchemaCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionRequest
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
				return ErrInvalidLengthActionRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActionRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionRequest
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
				return ErrInvalidLengthActionRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActionRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Action = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RefId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionRequest
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
				return ErrInvalidLengthActionRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActionRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RefId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiredAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionRequest
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
				return ErrInvalidLengthActionRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthActionRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.ExpiredAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OnBehalfOf", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionRequest
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
				return ErrInvalidLengthActionRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActionRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OnBehalfOf = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipActionRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthActionRequest
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
func (m *ActionRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowActionRequest
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
			return fmt.Errorf("proto: ActionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ActionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionRequest
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
					return ErrIntOverflowActionRequest
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
				return ErrInvalidLengthActionRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActionRequest
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
					return ErrIntOverflowActionRequest
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
				return ErrInvalidLengthActionRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActionRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionRequest
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
				return ErrInvalidLengthActionRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActionRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Action = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Caller", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionRequest
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
				return ErrInvalidLengthActionRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActionRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Caller = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RefId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionRequest
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
				return ErrInvalidLengthActionRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActionRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RefId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequiredConfirm", wireType)
			}
			m.RequiredConfirm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionRequest
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
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionRequest
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
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentConfirm", wireType)
			}
			m.CurrentConfirm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionRequest
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
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Confirmers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionRequest
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
				return ErrInvalidLengthActionRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActionRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Confirmers = append(m.Confirmers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionRequest
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
				return ErrInvalidLengthActionRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthActionRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidUntil", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionRequest
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
				return ErrInvalidLengthActionRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthActionRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.ValidUntil, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataHashes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionRequest
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
				return ErrInvalidLengthActionRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthActionRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataHashes = append(m.DataHashes, &DataHash{})
			if err := m.DataHashes[len(m.DataHashes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiredHeight", wireType)
			}
			m.ExpiredHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionRequest
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
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecutionErrorMessage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionRequest
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
				return ErrInvalidLengthActionRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActionRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExecutionErrorMessage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipActionRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthActionRequest
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
func skipActionRequest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowActionRequest
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
					return 0, ErrIntOverflowActionRequest
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
					return 0, ErrIntOverflowActionRequest
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
				return 0, ErrInvalidLengthActionRequest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupActionRequest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthActionRequest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthActionRequest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowActionRequest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupActionRequest = fmt.Errorf("proto: unexpected end of group")
)
