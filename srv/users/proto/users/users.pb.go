// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/users/users.proto

package com_hzuet_srv_users

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
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

type UserId struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserId) Reset()         { *m = UserId{} }
func (m *UserId) String() string { return proto.CompactTextString(m) }
func (*UserId) ProtoMessage()    {}
func (*UserId) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_4e25fc69eaf4b17d, []int{0}
}
func (m *UserId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserId.Unmarshal(m, b)
}
func (m *UserId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserId.Marshal(b, m, deterministic)
}
func (dst *UserId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserId.Merge(dst, src)
}
func (m *UserId) XXX_Size() int {
	return xxx_messageInfo_UserId.Size(m)
}
func (m *UserId) XXX_DiscardUnknown() {
	xxx_messageInfo_UserId.DiscardUnknown(m)
}

var xxx_messageInfo_UserId proto.InternalMessageInfo

func (m *UserId) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UserName struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserName) Reset()         { *m = UserName{} }
func (m *UserName) String() string { return proto.CompactTextString(m) }
func (*UserName) ProtoMessage()    {}
func (*UserName) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_4e25fc69eaf4b17d, []int{1}
}
func (m *UserName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserName.Unmarshal(m, b)
}
func (m *UserName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserName.Marshal(b, m, deterministic)
}
func (dst *UserName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserName.Merge(dst, src)
}
func (m *UserName) XXX_Size() int {
	return xxx_messageInfo_UserName.Size(m)
}
func (m *UserName) XXX_DiscardUnknown() {
	xxx_messageInfo_UserName.DiscardUnknown(m)
}

var xxx_messageInfo_UserName proto.InternalMessageInfo

func (m *UserName) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type UserInfo struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_4e25fc69eaf4b17d, []int{2}
}
func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (dst *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(dst, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserInfo) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserInfo) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type ListUsersRequest struct {
	Page                 int64    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize             int64    `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersRequest) Reset()         { *m = ListUsersRequest{} }
func (m *ListUsersRequest) String() string { return proto.CompactTextString(m) }
func (*ListUsersRequest) ProtoMessage()    {}
func (*ListUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_4e25fc69eaf4b17d, []int{3}
}
func (m *ListUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersRequest.Unmarshal(m, b)
}
func (m *ListUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersRequest.Marshal(b, m, deterministic)
}
func (dst *ListUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersRequest.Merge(dst, src)
}
func (m *ListUsersRequest) XXX_Size() int {
	return xxx_messageInfo_ListUsersRequest.Size(m)
}
func (m *ListUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersRequest proto.InternalMessageInfo

func (m *ListUsersRequest) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListUsersRequest) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

type ListUsersResponse struct {
	Users                []*UserInfo `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListUsersResponse) Reset()         { *m = ListUsersResponse{} }
func (m *ListUsersResponse) String() string { return proto.CompactTextString(m) }
func (*ListUsersResponse) ProtoMessage()    {}
func (*ListUsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_4e25fc69eaf4b17d, []int{4}
}
func (m *ListUsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersResponse.Unmarshal(m, b)
}
func (m *ListUsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersResponse.Marshal(b, m, deterministic)
}
func (dst *ListUsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersResponse.Merge(dst, src)
}
func (m *ListUsersResponse) XXX_Size() int {
	return xxx_messageInfo_ListUsersResponse.Size(m)
}
func (m *ListUsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersResponse proto.InternalMessageInfo

func (m *ListUsersResponse) GetUsers() []*UserInfo {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*UserId)(nil), "zergwangj.srv.users.UserId")
	proto.RegisterType((*UserName)(nil), "zergwangj.srv.users.UserName")
	proto.RegisterType((*UserInfo)(nil), "zergwangj.srv.users.UserInfo")
	proto.RegisterType((*ListUsersRequest)(nil), "zergwangj.srv.users.ListUsersRequest")
	proto.RegisterType((*ListUsersResponse)(nil), "zergwangj.srv.users.ListUsersResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UsersClient interface {
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
	GetUser(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*UserInfo, error)
	GetUserByName(ctx context.Context, in *UserName, opts ...grpc.CallOption) (*UserInfo, error)
	CreateUser(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*UserInfo, error)
	DeleteUser(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdateUser(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*UserInfo, error)
}

type usersClient struct {
	cc *grpc.ClientConn
}

func NewUsersClient(cc *grpc.ClientConn) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	out := new(ListUsersResponse)
	err := c.cc.Invoke(ctx, "/zergwangj.srv.users.Users/ListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetUser(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*UserInfo, error) {
	out := new(UserInfo)
	err := c.cc.Invoke(ctx, "/zergwangj.srv.users.Users/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetUserByName(ctx context.Context, in *UserName, opts ...grpc.CallOption) (*UserInfo, error) {
	out := new(UserInfo)
	err := c.cc.Invoke(ctx, "/zergwangj.srv.users.Users/GetUserByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) CreateUser(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*UserInfo, error) {
	out := new(UserInfo)
	err := c.cc.Invoke(ctx, "/zergwangj.srv.users.Users/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) DeleteUser(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/zergwangj.srv.users.Users/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) UpdateUser(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*UserInfo, error) {
	out := new(UserInfo)
	err := c.cc.Invoke(ctx, "/zergwangj.srv.users.Users/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
type UsersServer interface {
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
	GetUser(context.Context, *UserId) (*UserInfo, error)
	GetUserByName(context.Context, *UserName) (*UserInfo, error)
	CreateUser(context.Context, *UserInfo) (*UserInfo, error)
	DeleteUser(context.Context, *UserId) (*empty.Empty, error)
	UpdateUser(context.Context, *UserInfo) (*UserInfo, error)
}

func RegisterUsersServer(s *grpc.Server, srv UsersServer) {
	s.RegisterService(&_Users_serviceDesc, srv)
}

func _Users_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zergwangj.srv.users.Users/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).ListUsers(ctx, req.(*ListUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zergwangj.srv.users.Users/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetUser(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetUserByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetUserByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zergwangj.srv.users.Users/GetUserByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetUserByName(ctx, req.(*UserName))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zergwangj.srv.users.Users/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).CreateUser(ctx, req.(*UserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zergwangj.srv.users.Users/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).DeleteUser(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zergwangj.srv.users.Users/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).UpdateUser(ctx, req.(*UserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _Users_serviceDesc = grpc.ServiceDesc{
	ServiceName: "zergwangj.srv.users.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListUsers",
			Handler:    _Users_ListUsers_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _Users_GetUser_Handler,
		},
		{
			MethodName: "GetUserByName",
			Handler:    _Users_GetUserByName_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _Users_CreateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _Users_DeleteUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _Users_UpdateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/users/users.proto",
}

func init() { proto.RegisterFile("proto/users/users.proto", fileDescriptor_users_4e25fc69eaf4b17d) }

var fileDescriptor_users_4e25fc69eaf4b17d = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x65, 0xbb, 0x2d, 0xcd, 0x54, 0x44, 0xed, 0x80, 0x8a, 0x71, 0x40, 0x2a, 0x2b, 0x51,
	0x55, 0x3d, 0xac, 0x45, 0x7b, 0xe3, 0x58, 0x40, 0x50, 0x09, 0x71, 0x30, 0xea, 0x81, 0x03, 0x52,
	0xb7, 0x78, 0x6a, 0x2c, 0xd5, 0xde, 0xc5, 0xbb, 0x2e, 0x6a, 0xab, 0x5c, 0x78, 0x05, 0x1e, 0x8d,
	0x57, 0x40, 0x3c, 0x07, 0xf2, 0xae, 0xed, 0x90, 0x28, 0x7f, 0x2e, 0x5c, 0xa2, 0x8c, 0xbf, 0x99,
	0xdf, 0x37, 0x3b, 0xa3, 0x81, 0x47, 0xaa, 0x92, 0x46, 0xc6, 0xb5, 0xa6, 0x4a, 0xbb, 0x5f, 0x6e,
	0xbf, 0xe0, 0x83, 0x2f, 0xb2, 0xe0, 0x5f, 0x6f, 0x6b, 0x32, 0x5c, 0x57, 0xd7, 0xdc, 0x4a, 0xd1,
	0x93, 0x4c, 0xca, 0xec, 0x8a, 0x62, 0xa1, 0xf2, 0x58, 0x94, 0xa5, 0x34, 0xc2, 0xe4, 0xb2, 0x6c,
	0x4b, 0xa2, 0x51, 0xab, 0xda, 0xe8, 0xa2, 0xbe, 0x8c, 0xa9, 0x50, 0xe6, 0xc6, 0x89, 0x2c, 0x84,
	0x8d, 0x33, 0x4d, 0xd5, 0x69, 0x8a, 0x43, 0xf0, 0xf3, 0x34, 0xf4, 0xf6, 0xbc, 0x83, 0x20, 0xf1,
	0xf3, 0x94, 0xed, 0xc3, 0x66, 0xa3, 0x7c, 0x10, 0x05, 0x61, 0x04, 0x9b, 0x8d, 0x53, 0x29, 0x0a,
	0xb2, 0x19, 0x83, 0xa4, 0x8f, 0x59, 0xe2, 0xf2, 0x4e, 0xcb, 0x4b, 0x39, 0xcb, 0x98, 0xaa, 0xf3,
	0xa7, 0xeb, 0x1a, 0x4d, 0x09, 0xad, 0xbf, 0xcb, 0x2a, 0x0d, 0x03, 0xa7, 0x75, 0x31, 0x3b, 0x81,
	0xed, 0xf7, 0xb9, 0x36, 0x0d, 0x57, 0x27, 0xf4, 0xad, 0x26, 0x6d, 0x10, 0x61, 0x4d, 0x89, 0x8c,
	0x5a, 0xba, 0xfd, 0xef, 0x18, 0x19, 0x7d, 0xcc, 0x6f, 0x1d, 0x3f, 0x48, 0xfa, 0x98, 0xbd, 0x83,
	0x9d, 0x7f, 0x18, 0x5a, 0xc9, 0x52, 0x13, 0x1e, 0xc3, 0xba, 0x1d, 0x59, 0xe8, 0xed, 0x05, 0x07,
	0x5b, 0x47, 0x4f, 0xf9, 0x9c, 0x71, 0xf2, 0xee, 0x39, 0x89, 0xcb, 0x3d, 0xfa, 0xb3, 0x06, 0xeb,
	0x16, 0x83, 0x05, 0x0c, 0x7a, 0x26, 0x3e, 0x9f, 0x5b, 0x3c, 0xdb, 0x77, 0xb4, 0xbf, 0x2a, 0xcd,
	0xb5, 0xc6, 0x76, 0x7e, 0xfc, 0xfa, 0xfd, 0xd3, 0xdf, 0xc2, 0x41, 0x7c, 0xfd, 0xc2, 0xad, 0x1c,
	0x3f, 0xc3, 0xbd, 0xb7, 0x64, 0xd3, 0x70, 0xb4, 0xb8, 0xd3, 0x34, 0x5a, 0xfe, 0x0c, 0xb6, 0x6b,
	0xc9, 0xdb, 0x38, 0xec, 0xc9, 0xf1, 0x5d, 0x9e, 0x8e, 0x51, 0xc1, 0xfd, 0x16, 0x7f, 0x72, 0x63,
	0xd7, 0xbc, 0x98, 0xd3, 0xc8, 0xab, 0x6c, 0x9e, 0x59, 0x9b, 0x11, 0x3e, 0x9e, 0xd8, 0x34, 0x8b,
	0x8e, 0xef, 0xba, 0x95, 0x8f, 0xf1, 0x1c, 0xe0, 0x55, 0x45, 0xc2, 0x90, 0x7d, 0xd3, 0x72, 0xde,
	0x2a, 0xbb, 0x87, 0xd6, 0x6e, 0xc8, 0x26, 0xf3, 0x7a, 0xe9, 0x1d, 0xe2, 0x27, 0x80, 0xd7, 0x74,
	0x45, 0xad, 0xc3, 0xd2, 0xa9, 0xed, 0x72, 0x77, 0x18, 0xbc, 0x3b, 0x0c, 0xfe, 0xa6, 0x39, 0x8c,
	0x6e, 0x5c, 0x87, 0xb3, 0xe3, 0x3a, 0x07, 0x38, 0x53, 0xe9, 0x7f, 0x6d, 0x3e, 0x9a, 0x6a, 0xfe,
	0x62, 0xc3, 0x76, 0x72, 0xfc, 0x37, 0x00, 0x00, 0xff, 0xff, 0xef, 0x9c, 0x2b, 0xc1, 0xfe, 0x03,
	0x00, 0x00,
}
