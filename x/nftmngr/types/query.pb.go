// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftmngr/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1406c74dd8ff3e6a, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1406c74dd8ff3e6a, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryGetNFTSchemaRequest struct {
	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (m *QueryGetNFTSchemaRequest) Reset()         { *m = QueryGetNFTSchemaRequest{} }
func (m *QueryGetNFTSchemaRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetNFTSchemaRequest) ProtoMessage()    {}
func (*QueryGetNFTSchemaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1406c74dd8ff3e6a, []int{2}
}
func (m *QueryGetNFTSchemaRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetNFTSchemaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetNFTSchemaRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetNFTSchemaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetNFTSchemaRequest.Merge(m, src)
}
func (m *QueryGetNFTSchemaRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetNFTSchemaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetNFTSchemaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetNFTSchemaRequest proto.InternalMessageInfo

func (m *QueryGetNFTSchemaRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type QueryGetNFTSchemaResponse struct {
	NFTSchema NFTSchema `protobuf:"bytes,1,opt,name=nFTSchema,proto3" json:"nFTSchema"`
}

func (m *QueryGetNFTSchemaResponse) Reset()         { *m = QueryGetNFTSchemaResponse{} }
func (m *QueryGetNFTSchemaResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetNFTSchemaResponse) ProtoMessage()    {}
func (*QueryGetNFTSchemaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1406c74dd8ff3e6a, []int{3}
}
func (m *QueryGetNFTSchemaResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetNFTSchemaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetNFTSchemaResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetNFTSchemaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetNFTSchemaResponse.Merge(m, src)
}
func (m *QueryGetNFTSchemaResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetNFTSchemaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetNFTSchemaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetNFTSchemaResponse proto.InternalMessageInfo

func (m *QueryGetNFTSchemaResponse) GetNFTSchema() NFTSchema {
	if m != nil {
		return m.NFTSchema
	}
	return NFTSchema{}
}

type QueryAllNFTSchemaRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllNFTSchemaRequest) Reset()         { *m = QueryAllNFTSchemaRequest{} }
func (m *QueryAllNFTSchemaRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllNFTSchemaRequest) ProtoMessage()    {}
func (*QueryAllNFTSchemaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1406c74dd8ff3e6a, []int{4}
}
func (m *QueryAllNFTSchemaRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllNFTSchemaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllNFTSchemaRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllNFTSchemaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllNFTSchemaRequest.Merge(m, src)
}
func (m *QueryAllNFTSchemaRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllNFTSchemaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllNFTSchemaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllNFTSchemaRequest proto.InternalMessageInfo

func (m *QueryAllNFTSchemaRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllNFTSchemaResponse struct {
	NFTSchema  []NFTSchema         `protobuf:"bytes,1,rep,name=nFTSchema,proto3" json:"nFTSchema"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllNFTSchemaResponse) Reset()         { *m = QueryAllNFTSchemaResponse{} }
func (m *QueryAllNFTSchemaResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllNFTSchemaResponse) ProtoMessage()    {}
func (*QueryAllNFTSchemaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1406c74dd8ff3e6a, []int{5}
}
func (m *QueryAllNFTSchemaResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllNFTSchemaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllNFTSchemaResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllNFTSchemaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllNFTSchemaResponse.Merge(m, src)
}
func (m *QueryAllNFTSchemaResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllNFTSchemaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllNFTSchemaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllNFTSchemaResponse proto.InternalMessageInfo

func (m *QueryAllNFTSchemaResponse) GetNFTSchema() []NFTSchema {
	if m != nil {
		return m.NFTSchema
	}
	return nil
}

func (m *QueryAllNFTSchemaResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "sixnft.nftmngr.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "sixnft.nftmngr.QueryParamsResponse")
	proto.RegisterType((*QueryGetNFTSchemaRequest)(nil), "sixnft.nftmngr.QueryGetNFTSchemaRequest")
	proto.RegisterType((*QueryGetNFTSchemaResponse)(nil), "sixnft.nftmngr.QueryGetNFTSchemaResponse")
	proto.RegisterType((*QueryAllNFTSchemaRequest)(nil), "sixnft.nftmngr.QueryAllNFTSchemaRequest")
	proto.RegisterType((*QueryAllNFTSchemaResponse)(nil), "sixnft.nftmngr.QueryAllNFTSchemaResponse")
}

func init() { proto.RegisterFile("nftmngr/query.proto", fileDescriptor_1406c74dd8ff3e6a) }

var fileDescriptor_1406c74dd8ff3e6a = []byte{
	// 485 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xbf, 0x6f, 0xd4, 0x30,
	0x14, 0xc7, 0xcf, 0x6d, 0x39, 0xe9, 0x0c, 0x62, 0x70, 0x4f, 0xa7, 0x34, 0xaa, 0x02, 0xb8, 0x12,
	0xb4, 0x0c, 0x36, 0x2d, 0xac, 0x0c, 0xed, 0xd0, 0x0e, 0x48, 0xa8, 0x04, 0xa6, 0x2e, 0xc8, 0x39,
	0x7c, 0x21, 0x52, 0x62, 0xe7, 0x62, 0x17, 0xb5, 0x42, 0x2c, 0x2c, 0x0c, 0x2c, 0x08, 0xfe, 0x04,
	0xfe, 0x99, 0x8e, 0x95, 0x58, 0x98, 0x10, 0xba, 0xe3, 0x0f, 0x41, 0xb1, 0x9d, 0x5c, 0x73, 0x97,
	0xe3, 0xc7, 0x66, 0xf9, 0x7d, 0xdf, 0xf7, 0x7d, 0xde, 0xf3, 0x33, 0x5c, 0x17, 0x23, 0x9d, 0x89,
	0xb8, 0xa0, 0xe3, 0x53, 0x5e, 0x9c, 0x93, 0xbc, 0x90, 0x5a, 0xa2, 0x9b, 0x2a, 0x39, 0x13, 0x23,
	0x4d, 0x5c, 0xcc, 0xef, 0xc7, 0x32, 0x96, 0x26, 0x44, 0xcb, 0x93, 0x55, 0xf9, 0x9b, 0xb1, 0x94,
	0x71, 0xca, 0x29, 0xcb, 0x13, 0xca, 0x84, 0x90, 0x9a, 0xe9, 0x44, 0x0a, 0xe5, 0xa2, 0xf7, 0x87,
	0x52, 0x65, 0x52, 0xd1, 0x88, 0x29, 0x6e, 0xcd, 0xe9, 0x9b, 0xdd, 0x88, 0x6b, 0xb6, 0x4b, 0x73,
	0x16, 0x27, 0xc2, 0x88, 0x9d, 0xb6, 0x5f, 0x41, 0xe4, 0xac, 0x60, 0x59, 0xe5, 0xe0, 0x55, 0xb7,
	0x62, 0xa4, 0x5f, 0xaa, 0xe1, 0x6b, 0x9e, 0x31, 0x1b, 0xc1, 0x7d, 0x88, 0x9e, 0x95, 0x8e, 0xc7,
	0x46, 0x1e, 0xf2, 0xf1, 0x29, 0x57, 0x1a, 0x3f, 0x81, 0xeb, 0x8d, 0x5b, 0x95, 0x4b, 0xa1, 0x38,
	0x7a, 0x04, 0xbb, 0xd6, 0xd6, 0x03, 0xb7, 0xc1, 0xf6, 0xf5, 0xbd, 0x01, 0x69, 0x76, 0x47, 0xac,
	0xfe, 0x60, 0xed, 0xe2, 0xc7, 0xad, 0x4e, 0xe8, 0xb4, 0x98, 0x40, 0xcf, 0x98, 0x1d, 0x71, 0xfd,
	0xf4, 0xf0, 0xc5, 0x73, 0x53, 0xdd, 0x15, 0x42, 0x08, 0xae, 0x0d, 0xe5, 0x2b, 0x6e, 0xfc, 0x7a,
	0xa1, 0x39, 0xe3, 0x13, 0xb8, 0xd1, 0xa2, 0x77, 0x08, 0x8f, 0x61, 0x4f, 0x54, 0x97, 0x8e, 0x62,
	0x63, 0x9e, 0xa2, 0xce, 0x72, 0x20, 0xb3, 0x0c, 0x1c, 0x39, 0x96, 0xfd, 0x34, 0x5d, 0x60, 0x39,
	0x84, 0x70, 0x36, 0x4e, 0xe7, 0x7d, 0x97, 0xd8, 0xd9, 0x93, 0x72, 0xf6, 0xc4, 0x3e, 0xac, 0x9b,
	0x3d, 0x39, 0x66, 0x31, 0x77, 0xb9, 0xe1, 0x95, 0x4c, 0xfc, 0x15, 0xb8, 0x06, 0x9a, 0x45, 0xda,
	0x1b, 0x58, 0xfd, 0xbf, 0x06, 0xd0, 0x51, 0x03, 0x72, 0xc5, 0x40, 0xde, 0xfb, 0x2b, 0xa4, 0xad,
	0x7d, 0x95, 0x72, 0xef, 0xf3, 0x2a, 0xbc, 0x66, 0x28, 0xd1, 0x18, 0x76, 0xed, 0xbb, 0x21, 0x3c,
	0x0f, 0xb2, 0xb8, 0x1a, 0xfe, 0xd6, 0x1f, 0x35, 0xb6, 0x10, 0x0e, 0xde, 0x7f, 0xfb, 0xf5, 0x65,
	0xc5, 0x43, 0x03, 0x6a, 0xc5, 0xb4, 0xb9, 0x95, 0xe8, 0x23, 0x80, 0xbd, 0xba, 0x49, 0xb4, 0xdd,
	0x6a, 0xd9, 0xb2, 0x2e, 0xfe, 0xce, 0x3f, 0x28, 0x1d, 0xc2, 0x8e, 0x41, 0xd8, 0x42, 0x77, 0xe6,
	0x11, 0x66, 0x5f, 0x80, 0xbe, 0x2d, 0xf7, 0xed, 0x1d, 0xfa, 0x00, 0xe0, 0x8d, 0xda, 0x60, 0x3f,
	0x4d, 0x97, 0x00, 0xb5, 0xec, 0xcc, 0x12, 0xa0, 0xb6, 0x87, 0xc7, 0xd8, 0x00, 0x6d, 0x22, 0x7f,
	0x39, 0xd0, 0xc1, 0x83, 0x8b, 0x49, 0x00, 0x2e, 0x27, 0x01, 0xf8, 0x39, 0x09, 0xc0, 0xa7, 0x69,
	0xd0, 0xb9, 0x9c, 0x06, 0x9d, 0xef, 0xd3, 0xa0, 0x73, 0x32, 0x70, 0x49, 0x67, 0x75, 0x9a, 0x3e,
	0xcf, 0xb9, 0x8a, 0xba, 0xe6, 0x1b, 0x3f, 0xfc, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x91, 0x8e, 0x96,
	0x9b, 0x7d, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a NFTSchema by index.
	NFTSchema(ctx context.Context, in *QueryGetNFTSchemaRequest, opts ...grpc.CallOption) (*QueryGetNFTSchemaResponse, error)
	// Queries a list of NFTSchema items.
	NFTSchemaAll(ctx context.Context, in *QueryAllNFTSchemaRequest, opts ...grpc.CallOption) (*QueryAllNFTSchemaResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/sixnft.nftmngr.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) NFTSchema(ctx context.Context, in *QueryGetNFTSchemaRequest, opts ...grpc.CallOption) (*QueryGetNFTSchemaResponse, error) {
	out := new(QueryGetNFTSchemaResponse)
	err := c.cc.Invoke(ctx, "/sixnft.nftmngr.Query/NFTSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) NFTSchemaAll(ctx context.Context, in *QueryAllNFTSchemaRequest, opts ...grpc.CallOption) (*QueryAllNFTSchemaResponse, error) {
	out := new(QueryAllNFTSchemaResponse)
	err := c.cc.Invoke(ctx, "/sixnft.nftmngr.Query/NFTSchemaAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a NFTSchema by index.
	NFTSchema(context.Context, *QueryGetNFTSchemaRequest) (*QueryGetNFTSchemaResponse, error)
	// Queries a list of NFTSchema items.
	NFTSchemaAll(context.Context, *QueryAllNFTSchemaRequest) (*QueryAllNFTSchemaResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) NFTSchema(ctx context.Context, req *QueryGetNFTSchemaRequest) (*QueryGetNFTSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NFTSchema not implemented")
}
func (*UnimplementedQueryServer) NFTSchemaAll(ctx context.Context, req *QueryAllNFTSchemaRequest) (*QueryAllNFTSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NFTSchemaAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sixnft.nftmngr.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_NFTSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetNFTSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).NFTSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sixnft.nftmngr.Query/NFTSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).NFTSchema(ctx, req.(*QueryGetNFTSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_NFTSchemaAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllNFTSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).NFTSchemaAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sixnft.nftmngr.Query/NFTSchemaAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).NFTSchemaAll(ctx, req.(*QueryAllNFTSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sixnft.nftmngr.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "NFTSchema",
			Handler:    _Query_NFTSchema_Handler,
		},
		{
			MethodName: "NFTSchemaAll",
			Handler:    _Query_NFTSchemaAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nftmngr/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryGetNFTSchemaRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetNFTSchemaRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetNFTSchemaRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Code) > 0 {
		i -= len(m.Code)
		copy(dAtA[i:], m.Code)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Code)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetNFTSchemaResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetNFTSchemaResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetNFTSchemaResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.NFTSchema.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAllNFTSchemaRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllNFTSchemaRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllNFTSchemaRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllNFTSchemaResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllNFTSchemaResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllNFTSchemaResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.NFTSchema) > 0 {
		for iNdEx := len(m.NFTSchema) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NFTSchema[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryGetNFTSchemaRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Code)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetNFTSchemaResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.NFTSchema.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllNFTSchemaRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllNFTSchemaResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.NFTSchema) > 0 {
		for _, e := range m.NFTSchema {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetNFTSchemaRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetNFTSchemaRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetNFTSchemaRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Code = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetNFTSchemaResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetNFTSchemaResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetNFTSchemaResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NFTSchema", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NFTSchema.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllNFTSchemaRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllNFTSchemaRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllNFTSchemaRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllNFTSchemaResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllNFTSchemaResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllNFTSchemaResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NFTSchema", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NFTSchema = append(m.NFTSchema, NFTSchema{})
			if err := m.NFTSchema[len(m.NFTSchema)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
