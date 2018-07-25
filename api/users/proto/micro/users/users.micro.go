// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/micro/users/users.proto

/*
Package users is a generated protocol buffer package.

It is generated from these files:
	proto/micro/users/users.proto

It has these top-level messages:
	UserId
	UserName
	UserInfo
	ListUsersRequest
	ListUsersResponse
*/
package users

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = google_protobuf1.Empty{}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Users service

type UsersService interface {
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...client.CallOption) (*ListUsersResponse, error)
	GetUser(ctx context.Context, in *UserId, opts ...client.CallOption) (*UserInfo, error)
	GetUserByName(ctx context.Context, in *UserName, opts ...client.CallOption) (*UserInfo, error)
	CreateUser(ctx context.Context, in *UserInfo, opts ...client.CallOption) (*UserInfo, error)
	DeleteUser(ctx context.Context, in *UserId, opts ...client.CallOption) (*google_protobuf1.Empty, error)
	UpdateUser(ctx context.Context, in *UserInfo, opts ...client.CallOption) (*UserInfo, error)
}

type usersService struct {
	c    client.Client
	name string
}

func NewUsersService(name string, c client.Client) UsersService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "users"
	}
	return &usersService{
		c:    c,
		name: name,
	}
}

func (c *usersService) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...client.CallOption) (*ListUsersResponse, error) {
	req := c.c.NewRequest(c.name, "Users.ListUsers", in)
	out := new(ListUsersResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersService) GetUser(ctx context.Context, in *UserId, opts ...client.CallOption) (*UserInfo, error) {
	req := c.c.NewRequest(c.name, "Users.GetUser", in)
	out := new(UserInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersService) GetUserByName(ctx context.Context, in *UserName, opts ...client.CallOption) (*UserInfo, error) {
	req := c.c.NewRequest(c.name, "Users.GetUserByName", in)
	out := new(UserInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersService) CreateUser(ctx context.Context, in *UserInfo, opts ...client.CallOption) (*UserInfo, error) {
	req := c.c.NewRequest(c.name, "Users.CreateUser", in)
	out := new(UserInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersService) DeleteUser(ctx context.Context, in *UserId, opts ...client.CallOption) (*google_protobuf1.Empty, error) {
	req := c.c.NewRequest(c.name, "Users.DeleteUser", in)
	out := new(google_protobuf1.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersService) UpdateUser(ctx context.Context, in *UserInfo, opts ...client.CallOption) (*UserInfo, error) {
	req := c.c.NewRequest(c.name, "Users.UpdateUser", in)
	out := new(UserInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Users service

type UsersHandler interface {
	ListUsers(context.Context, *ListUsersRequest, *ListUsersResponse) error
	GetUser(context.Context, *UserId, *UserInfo) error
	GetUserByName(context.Context, *UserName, *UserInfo) error
	CreateUser(context.Context, *UserInfo, *UserInfo) error
	DeleteUser(context.Context, *UserId, *google_protobuf1.Empty) error
	UpdateUser(context.Context, *UserInfo, *UserInfo) error
}

func RegisterUsersHandler(s server.Server, hdlr UsersHandler, opts ...server.HandlerOption) {
	type users interface {
		ListUsers(ctx context.Context, in *ListUsersRequest, out *ListUsersResponse) error
		GetUser(ctx context.Context, in *UserId, out *UserInfo) error
		GetUserByName(ctx context.Context, in *UserName, out *UserInfo) error
		CreateUser(ctx context.Context, in *UserInfo, out *UserInfo) error
		DeleteUser(ctx context.Context, in *UserId, out *google_protobuf1.Empty) error
		UpdateUser(ctx context.Context, in *UserInfo, out *UserInfo) error
	}
	type Users struct {
		users
	}
	h := &usersHandler{hdlr}
	s.Handle(s.NewHandler(&Users{h}, opts...))
}

type usersHandler struct {
	UsersHandler
}

func (h *usersHandler) ListUsers(ctx context.Context, in *ListUsersRequest, out *ListUsersResponse) error {
	return h.UsersHandler.ListUsers(ctx, in, out)
}

func (h *usersHandler) GetUser(ctx context.Context, in *UserId, out *UserInfo) error {
	return h.UsersHandler.GetUser(ctx, in, out)
}

func (h *usersHandler) GetUserByName(ctx context.Context, in *UserName, out *UserInfo) error {
	return h.UsersHandler.GetUserByName(ctx, in, out)
}

func (h *usersHandler) CreateUser(ctx context.Context, in *UserInfo, out *UserInfo) error {
	return h.UsersHandler.CreateUser(ctx, in, out)
}

func (h *usersHandler) DeleteUser(ctx context.Context, in *UserId, out *google_protobuf1.Empty) error {
	return h.UsersHandler.DeleteUser(ctx, in, out)
}

func (h *usersHandler) UpdateUser(ctx context.Context, in *UserInfo, out *UserInfo) error {
	return h.UsersHandler.UpdateUser(ctx, in, out)
}