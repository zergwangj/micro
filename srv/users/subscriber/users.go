package subscriber

import (
	"github.com/micro/go-log"
	"context"
	users "github.com/zergwangj/micro/srv/users/proto/users"
)

type Users struct{}

func (u *Users) Handle(ctx context.Context, msg *users.UserInfo) error {
	log.Log("Handler Received message: ", msg)
	return nil
}
