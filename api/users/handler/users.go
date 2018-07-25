package handler

import (
	"context"
	"github.com/micro/go-log"

	"github.com/zergwangj/micro/api/users/client"
	"github.com/micro/go-micro/errors"
	usersapi "github.com/zergwangj/micro/api/users/proto/micro/users"
	userssrv "github.com/zergwangj/micro/srv/users/proto/users"
	google_protobuf1 "github.com/golang/protobuf/ptypes/empty"
)

type Users struct{}

// ListUsers
func (u *Users) ListUsers(ctx context.Context, req *usersapi.ListUsersRequest, rsp *usersapi.ListUsersResponse) error {
	log.Log("Received Users.ListUsers request")

	// extract the client from the context
	usersClient, ok := client.UsersFromContext(ctx)
	if !ok {
		return errors.InternalServerError("zergwangj.api.users.ListUsers", "users client not found")
	}

	// make request
	response, err := usersClient.ListUsers(ctx, &userssrv.ListUsersRequest{
		Page: 		req.Page,
		PageSize:	req.PageSize,
	})
	if err != nil {
		return errors.InternalServerError("zergwangj.api.users.ListUsers", err.Error())
	}

	for _, user := range response.Users {
		userInfo := new(usersapi.UserInfo)
		userInfo.Id = user.Id
		userInfo.Username = user.Username
		userInfo.Password = user.Password
		rsp.Users = append(rsp.Users, userInfo)
	}
	return nil
}

// GetUser
func (u *Users) GetUser(ctx context.Context, req *usersapi.UserId, rsp *usersapi.UserInfo) error {
	log.Log("Received Users.GetUser request")

	// extract the client from the context
	usersClient, ok := client.UsersFromContext(ctx)
	if !ok {
		return errors.InternalServerError("zergwangj.api.users.GetUser", "users client not found")
	}

	// make request
	response, err := usersClient.GetUser(ctx, &userssrv.UserId{
		Id: req.Id,
	})
	if err != nil {
		return errors.InternalServerError("zergwangj.api.users.GetUser", err.Error())
	}

	rsp.Id = response.Id
	rsp.Username = response.Username
	rsp.Password = response.Password

	return nil
}

// GetUserByName
func (u *Users) GetUserByName(ctx context.Context, req *usersapi.UserName, rsp *usersapi.UserInfo) error {
	log.Log("Received Users.GetUserByName request")

	// extract the client from the context
	usersClient, ok := client.UsersFromContext(ctx)
	if !ok {
		return errors.InternalServerError("zergwangj.api.users.GetUserByName", "users client not found")
	}

	// make request
	response, err := usersClient.GetUserByName(ctx, &userssrv.UserName{
		Username: req.Username,
	})
	if err != nil {
		return errors.InternalServerError("zergwangj.api.users.GetUserByName", err.Error())
	}

	rsp.Id = response.Id
	rsp.Username = response.Username
	rsp.Password = response.Password

	return nil
}

// CreateUser
func (u *Users) CreateUser(ctx context.Context, req *usersapi.UserInfo, rsp *usersapi.UserInfo) error {
	log.Log("Received Users.CreateUser request")

	// extract the client from the context
	usersClient, ok := client.UsersFromContext(ctx)
	if !ok {
		return errors.InternalServerError("zergwangj.api.users.CreateUser", "users client not found")
	}

	// make request
	response, err := usersClient.CreateUser(ctx, &userssrv.UserInfo{
		Id:			req.Id,
		Username: 	req.Username,
		Password:	req.Password,
	})
	if err != nil {
		return errors.InternalServerError("zergwangj.api.users.CreateUser", err.Error())
	}

	rsp.Id = response.Id
	rsp.Username = response.Username
	rsp.Password = response.Password

	return nil
}

// DeleteUser
func (u *Users) DeleteUser(ctx context.Context, req *usersapi.UserId, rsp *google_protobuf1.Empty) error {
	log.Log("Received Users.DeleteUser request")

	// extract the client from the context
	usersClient, ok := client.UsersFromContext(ctx)
	if !ok {
		return errors.InternalServerError("zergwangj.api.users.DeleteUser", "users client not found")
	}

	// make request
	_, err := usersClient.DeleteUser(ctx, &userssrv.UserId{
		Id: req.Id,
	})
	if err != nil {
		return errors.InternalServerError("zergwangj.api.users.DeleteUser", err.Error())
	}

	return nil
}

// UpdateUser
func (u *Users) UpdateUser(ctx context.Context, req *usersapi.UserInfo, rsp *usersapi.UserInfo) error {
	log.Log("Received Users.UpdateUser request")

	// extract the client from the context
	usersClient, ok := client.UsersFromContext(ctx)
	if !ok {
		return errors.InternalServerError("zergwangj.api.users.UpdateUser", "users client not found")
	}

	// make request
	response, err := usersClient.UpdateUser(ctx, &userssrv.UserInfo{
		Id:			req.Id,
		Username: 	req.Username,
		Password:	req.Password,
	})
	if err != nil {
		return errors.InternalServerError("zergwangj.api.users.UpdateUser", err.Error())
	}

	rsp.Id = response.Id
	rsp.Username = response.Username
	rsp.Password = response.Password

	return nil
}