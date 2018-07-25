// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/micro/auth/auth.proto

package auth

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateSessionRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSessionRequest) Reset()         { *m = CreateSessionRequest{} }
func (m *CreateSessionRequest) String() string { return proto.CompactTextString(m) }
func (*CreateSessionRequest) ProtoMessage()    {}
func (*CreateSessionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_454de6f864046ef2, []int{0}
}
func (m *CreateSessionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSessionRequest.Unmarshal(m, b)
}
func (m *CreateSessionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSessionRequest.Marshal(b, m, deterministic)
}
func (dst *CreateSessionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSessionRequest.Merge(dst, src)
}
func (m *CreateSessionRequest) XXX_Size() int {
	return xxx_messageInfo_CreateSessionRequest.Size(m)
}
func (m *CreateSessionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSessionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSessionRequest proto.InternalMessageInfo

func (m *CreateSessionRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *CreateSessionRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type CreateSessionResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSessionResponse) Reset()         { *m = CreateSessionResponse{} }
func (m *CreateSessionResponse) String() string { return proto.CompactTextString(m) }
func (*CreateSessionResponse) ProtoMessage()    {}
func (*CreateSessionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_454de6f864046ef2, []int{1}
}
func (m *CreateSessionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSessionResponse.Unmarshal(m, b)
}
func (m *CreateSessionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSessionResponse.Marshal(b, m, deterministic)
}
func (dst *CreateSessionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSessionResponse.Merge(dst, src)
}
func (m *CreateSessionResponse) XXX_Size() int {
	return xxx_messageInfo_CreateSessionResponse.Size(m)
}
func (m *CreateSessionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSessionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSessionResponse proto.InternalMessageInfo

func (m *CreateSessionResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ValidateTokenRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateTokenRequest) Reset()         { *m = ValidateTokenRequest{} }
func (m *ValidateTokenRequest) String() string { return proto.CompactTextString(m) }
func (*ValidateTokenRequest) ProtoMessage()    {}
func (*ValidateTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_454de6f864046ef2, []int{2}
}
func (m *ValidateTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateTokenRequest.Unmarshal(m, b)
}
func (m *ValidateTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateTokenRequest.Marshal(b, m, deterministic)
}
func (dst *ValidateTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateTokenRequest.Merge(dst, src)
}
func (m *ValidateTokenRequest) XXX_Size() int {
	return xxx_messageInfo_ValidateTokenRequest.Size(m)
}
func (m *ValidateTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateTokenRequest proto.InternalMessageInfo

func (m *ValidateTokenRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ValidateTokenResponse struct {
	Valid                bool     `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateTokenResponse) Reset()         { *m = ValidateTokenResponse{} }
func (m *ValidateTokenResponse) String() string { return proto.CompactTextString(m) }
func (*ValidateTokenResponse) ProtoMessage()    {}
func (*ValidateTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_454de6f864046ef2, []int{3}
}
func (m *ValidateTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateTokenResponse.Unmarshal(m, b)
}
func (m *ValidateTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateTokenResponse.Marshal(b, m, deterministic)
}
func (dst *ValidateTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateTokenResponse.Merge(dst, src)
}
func (m *ValidateTokenResponse) XXX_Size() int {
	return xxx_messageInfo_ValidateTokenResponse.Size(m)
}
func (m *ValidateTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateTokenResponse proto.InternalMessageInfo

func (m *ValidateTokenResponse) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *ValidateTokenResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type DeleteSessionRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteSessionRequest) Reset()         { *m = DeleteSessionRequest{} }
func (m *DeleteSessionRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteSessionRequest) ProtoMessage()    {}
func (*DeleteSessionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_454de6f864046ef2, []int{4}
}
func (m *DeleteSessionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteSessionRequest.Unmarshal(m, b)
}
func (m *DeleteSessionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteSessionRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteSessionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteSessionRequest.Merge(dst, src)
}
func (m *DeleteSessionRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteSessionRequest.Size(m)
}
func (m *DeleteSessionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteSessionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteSessionRequest proto.InternalMessageInfo

func (m *DeleteSessionRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *DeleteSessionRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type DeleteSessionResponse struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteSessionResponse) Reset()         { *m = DeleteSessionResponse{} }
func (m *DeleteSessionResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteSessionResponse) ProtoMessage()    {}
func (*DeleteSessionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_454de6f864046ef2, []int{5}
}
func (m *DeleteSessionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteSessionResponse.Unmarshal(m, b)
}
func (m *DeleteSessionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteSessionResponse.Marshal(b, m, deterministic)
}
func (dst *DeleteSessionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteSessionResponse.Merge(dst, src)
}
func (m *DeleteSessionResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteSessionResponse.Size(m)
}
func (m *DeleteSessionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteSessionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteSessionResponse proto.InternalMessageInfo

func (m *DeleteSessionResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func init() {
	proto.RegisterType((*CreateSessionRequest)(nil), "auth.CreateSessionRequest")
	proto.RegisterType((*CreateSessionResponse)(nil), "auth.CreateSessionResponse")
	proto.RegisterType((*ValidateTokenRequest)(nil), "auth.ValidateTokenRequest")
	proto.RegisterType((*ValidateTokenResponse)(nil), "auth.ValidateTokenResponse")
	proto.RegisterType((*DeleteSessionRequest)(nil), "auth.DeleteSessionRequest")
	proto.RegisterType((*DeleteSessionResponse)(nil), "auth.DeleteSessionResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthClient interface {
	CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*CreateSessionResponse, error)
	ValidateToken(ctx context.Context, in *ValidateTokenRequest, opts ...grpc.CallOption) (*ValidateTokenResponse, error)
	DeleteSession(ctx context.Context, in *DeleteSessionRequest, opts ...grpc.CallOption) (*DeleteSessionResponse, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*CreateSessionResponse, error) {
	out := new(CreateSessionResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/CreateSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ValidateToken(ctx context.Context, in *ValidateTokenRequest, opts ...grpc.CallOption) (*ValidateTokenResponse, error) {
	out := new(ValidateTokenResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/ValidateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) DeleteSession(ctx context.Context, in *DeleteSessionRequest, opts ...grpc.CallOption) (*DeleteSessionResponse, error) {
	out := new(DeleteSessionResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/DeleteSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
type AuthServer interface {
	CreateSession(context.Context, *CreateSessionRequest) (*CreateSessionResponse, error)
	ValidateToken(context.Context, *ValidateTokenRequest) (*ValidateTokenResponse, error)
	DeleteSession(context.Context, *DeleteSessionRequest) (*DeleteSessionResponse, error)
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/CreateSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CreateSession(ctx, req.(*CreateSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ValidateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ValidateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/ValidateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ValidateToken(ctx, req.(*ValidateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_DeleteSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).DeleteSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/DeleteSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).DeleteSession(ctx, req.(*DeleteSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSession",
			Handler:    _Auth_CreateSession_Handler,
		},
		{
			MethodName: "ValidateToken",
			Handler:    _Auth_ValidateToken_Handler,
		},
		{
			MethodName: "DeleteSession",
			Handler:    _Auth_DeleteSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/micro/auth/auth.proto",
}

func init() { proto.RegisterFile("proto/micro/auth/auth.proto", fileDescriptor_auth_454de6f864046ef2) }

var fileDescriptor_auth_454de6f864046ef2 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0xd5, 0xa8, 0xa0, 0x72, 0x52, 0x91, 0x88, 0x52, 0x54, 0xd2, 0x0e, 0xe0, 0x05, 0x54,
	0x41, 0x23, 0x60, 0x63, 0x43, 0x65, 0x60, 0x62, 0x28, 0x88, 0xdd, 0xd0, 0xa3, 0x8d, 0xd2, 0xfa,
	0x82, 0xed, 0x94, 0x01, 0xb1, 0xf0, 0x0a, 0xbc, 0x14, 0x3b, 0xaf, 0xc0, 0x83, 0xa0, 0xd8, 0x71,
	0xa1, 0xc5, 0x0c, 0x2c, 0x51, 0xee, 0xee, 0xd7, 0x7d, 0x7f, 0xfe, 0x0b, 0x74, 0x72, 0x49, 0x9a,
	0x92, 0x59, 0x7a, 0x2f, 0x29, 0xe1, 0x85, 0x9e, 0x98, 0x47, 0xdf, 0x74, 0xc3, 0x7a, 0xf9, 0x1e,
	0x77, 0xc7, 0x44, 0xe3, 0x29, 0x26, 0x3c, 0x4f, 0x13, 0x2e, 0x04, 0x69, 0xae, 0x53, 0x12, 0xca,
	0x6a, 0xd8, 0x15, 0x44, 0x03, 0x89, 0x5c, 0xe3, 0x35, 0x2a, 0x95, 0x92, 0x18, 0xe2, 0x63, 0x81,
	0x4a, 0x87, 0x31, 0x34, 0x0a, 0x85, 0x52, 0xf0, 0x19, 0xb6, 0x6b, 0xbb, 0xb5, 0x83, 0x8d, 0xe1,
	0xa2, 0x2e, 0x67, 0x39, 0x57, 0xea, 0x89, 0xe4, 0xa8, 0x1d, 0xd8, 0x99, 0xab, 0xd9, 0x11, 0xb4,
	0x56, 0xf6, 0xa9, 0x9c, 0x84, 0xc2, 0x30, 0x82, 0x35, 0x4d, 0x19, 0x8a, 0x6a, 0x9b, 0x2d, 0xd8,
	0x21, 0x44, 0xb7, 0x7c, 0x9a, 0x8e, 0xb8, 0xc6, 0x9b, 0xb2, 0xe1, 0xf0, 0x7e, 0xf5, 0x00, 0x5a,
	0x2b, 0xea, 0xef, 0xe5, 0xf3, 0x72, 0x60, 0xe4, 0x8d, 0xa1, 0x2d, 0xca, 0x2e, 0x4a, 0x49, 0xb2,
	0x32, 0x69, 0x0b, 0x76, 0x09, 0xd1, 0x05, 0x4e, 0xf1, 0x5f, 0x5f, 0xbc, 0xb0, 0x13, 0xfc, 0xb4,
	0xb3, 0x0f, 0xad, 0x95, 0x4d, 0x95, 0x9d, 0x4d, 0x08, 0x28, 0xab, 0xbc, 0x04, 0x94, 0x9d, 0xbc,
	0x07, 0x50, 0x3f, 0x2f, 0xf4, 0x24, 0x7c, 0x80, 0xe6, 0x52, 0x3a, 0x61, 0xdc, 0x37, 0xf7, 0xf2,
	0x9d, 0x20, 0xee, 0x78, 0x67, 0x16, 0xc1, 0xba, 0xaf, 0x1f, 0x9f, 0x6f, 0xc1, 0x36, 0xdb, 0x4a,
	0xe6, 0xc7, 0xf6, 0xf0, 0xca, 0x2a, 0xd4, 0x59, 0xad, 0x17, 0x66, 0xd0, 0x5c, 0x0a, 0xca, 0x71,
	0x7c, 0x59, 0x3b, 0x8e, 0x37, 0x59, 0xb6, 0x67, 0x38, 0x9d, 0x70, 0xe7, 0x17, 0x27, 0x79, 0x36,
	0x29, 0xbc, 0x94, 0xb0, 0xa5, 0x18, 0x1c, 0xcc, 0x97, 0xb2, 0x83, 0x79, 0x73, 0x73, 0xb0, 0xde,
	0xdf, 0xb0, 0xbb, 0x75, 0xf3, 0xdb, 0x9e, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x59, 0x57, 0x6a,
	0xfc, 0xf9, 0x02, 0x00, 0x00,
}
