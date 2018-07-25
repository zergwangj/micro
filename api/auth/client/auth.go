package client

import (
	"context"
	"github.com/micro/go-config"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"github.com/zergwangj/micro/srv/auth/proto/auth"
	"github.com/micro/go-grpc"
	"github.com/micro/go-plugins/wrapper/breaker/hystrix"
	"github.com/micro/go-plugins/wrapper/select/roundrobin"
	"github.com/micro/go-plugins/wrapper/ratelimiter/ratelimit"
	jujuratelimit "github.com/juju/ratelimit"
)

type authKey struct {}

// FromContext retrieves the client from the Context
func AuthFromContext(ctx context.Context) (auth.AuthService, bool) {
	c, ok := ctx.Value(authKey{}).(auth.AuthService)
	return c, ok
}

// Client returns a wrapper for the AuthClient
func AuthWrapper(conf config.Config) server.HandlerWrapper {
	service := grpc.NewService(
		micro.Name("zergwangj.srv.auth"),
	)
	limit := conf.Get("ratelimit").Int(100)
	b := jujuratelimit.NewBucketWithRate(float64(limit), int64(limit))
	service.Init(
		micro.WrapClient(ratelimit.NewClientWrapper(b, false)),
		micro.WrapClient(hystrix.NewClientWrapper()),
		micro.WrapClient(roundrobin.NewClientWrapper()),
	)
	client := auth.NewAuthService("zergwangj.srv.auth", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, authKey{}, client)
			return fn(ctx, req, rsp)
		}
	}
}
