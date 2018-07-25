package handler

import (
	"context"
	"github.com/micro/go-log"
	users "github.com/zergwangj/micro/srv/users/proto/users"
	"database/sql"
	"github.com/zergwangj/micro/srv/users/models"
	"github.com/micro/go-config"
	google_protobuf1 "github.com/golang/protobuf/ptypes/empty"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro"
)

type Users struct {
	Db 			*sql.DB
	Conf		config.Config
	Pub			micro.Publisher
}

// ListUsers
func (u *Users) ListUsers(ctx context.Context, req *users.ListUsersRequest, rsp *users.ListUsersResponse) error {
	log.Log("Received Users.ListUsers request")

	if req.Page <= 0 || req.PageSize <= 0 {
		return errors.InternalServerError("zergwangj.srv.users.ListUsers", "param is invalid")
	}

	db := u.Db
	userList, err := models.ListUsers(db, req.Page, req.PageSize)
	if err != nil {
		return err
	}
	for _, user := range userList {
		userInfo := new(users.UserInfo)
		userInfo.Id = user.Id
		userInfo.Username = user.Username
		userInfo.Password = user.Password
		rsp.Users = append(rsp.Users, userInfo)
	}
	return nil
}

// GetUser
func (u *Users) GetUser(ctx context.Context, req *users.UserId, rsp *users.UserInfo) error {
	log.Log("Received Users.GetUser request")
	db := u.Db
	user, err := models.GetUser(db, req.Id)
	if err != nil {
		return err
	}
	rsp.Id = user.Id
	rsp.Username = user.Username
	rsp.Password = user.Password
	return nil
}

// GetUserByName
func (u *Users) GetUserByName(ctx context.Context, req *users.UserName, rsp *users.UserInfo) error {
	log.Log("Received Users.GetUserByName request")
	db := u.Db
	user, err := models.GetUserByName(db, req.Username)
	if err != nil {
		return err
	}
	rsp.Id = user.Id
	rsp.Username = user.Username
	rsp.Password = user.Password
	return nil
}

// CreateUser
func (u *Users) CreateUser(ctx context.Context, req *users.UserInfo, rsp *users.UserInfo) error {
	log.Log("Received Users.CreateUser request")
	db := u.Db
	user, err := models.CreateUser(db, &models.User{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return err
	}
	rsp.Id = user.Id
	rsp.Username = user.Username
	rsp.Password = user.Password

	pub := u.Pub
	pub.Publish(context.TODO(), &users.UserInfo{
		Id:			rsp.Id,
		Username: 	rsp.Username,
		Password: 	rsp.Password,
	})

	return nil
}

// DeleteUser
func (u *Users) DeleteUser(ctx context.Context, req *users.UserId, rsp *google_protobuf1.Empty) error {
	log.Log("Received Users.DeleteUser request")
	db := u.Db
	err := models.DeleteUser(db, req.Id)
	if err != nil {
		return err
	}
	return nil
}

// UpdateUser
func (u *Users) UpdateUser(ctx context.Context, req *users.UserInfo, rsp *users.UserInfo) error {
	log.Log("Received Users.UpdateUser request")
	db := u.Db
	user, err := models.UpdateUser(db, &models.User{
		Id:		  	req.Id,
		Username: 	req.Username,
		Password: 	req.Password,
	})
	if err != nil {
		return err
	}
	rsp.Id = user.Id
	rsp.Username = user.Username
	rsp.Password = user.Password
	return nil
}