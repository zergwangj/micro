package client

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	users "github.com/zergwangj/micro/srv/users/proto/users"
	"github.com/micro/go-grpc"
	"github.com/micro/go-config"
	"github.com/micro/go-plugins/wrapper/breaker/hystrix"
	"github.com/micro/go-plugins/wrapper/select/roundrobin"
	"github.com/micro/go-plugins/wrapper/ratelimiter/ratelimit"
	jujuratelimit "github.com/juju/ratelimit"
)

type usersKey struct {}

// FromContext retrieves the client from the Context
func UsersFromContext(ctx context.Context) (users.UsersService, bool) {
	c, ok := ctx.Value(usersKey{}).(users.UsersService)
	return c, ok
}

// Client returns a wrapper for the UsersClient
func UsersWrapper(conf config.Config) server.HandlerWrapper {
	service := grpc.NewService(
		micro.Name("zergwangj.srv.users"),
	)
	limit := conf.Get("ratelimit").Int(100)
	b := jujuratelimit.NewBucketWithRate(float64(limit), int64(limit))
	service.Init(
		micro.WrapClient(ratelimit.NewClientWrapper(b, false)),
		micro.WrapClient(hystrix.NewClientWrapper()),
		micro.WrapClient(roundrobin.NewClientWrapper()),
	)
	client := users.NewUsersService("zergwangj.srv.users", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, usersKey{}, client)
			return fn(ctx, req, rsp)
		}
	}
}
