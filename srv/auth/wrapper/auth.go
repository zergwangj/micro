package wrapper

import (
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
	"context"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro"
	auth "github.com/zergwangj/micro/srv/auth/proto/auth"
	"errors"
	"strings"
	"github.com/micro/go-grpc"
)

type clientWrapper struct {
	token string
	client.Client
}

func (c *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	ctx = metadata.NewContext(ctx, map[string]string{
		"Token": c.token,
	})
	return c.Client.Call(ctx, req, rsp, opts...)
}

// NewClientWrapper
func NewClientWrapper(token string) client.Wrapper {
	return func(c client.Client) client.Client {
		return &clientWrapper{token, c}
	}
}

// NewHandlerWrapper
func NewHandlerWrapper() server.HandlerWrapper {
	return func(h server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			service := grpc.NewService(
				micro.Name("zergwangj.srv.auth"),
			)
			service.Init()
			client := auth.NewAuthService("zergwangj.srv.auth", service.Client())
			meta, ok := metadata.FromContext(ctx)
			if !ok {
				return errors.New("no auth meta-data found in request")
			}
			authString := meta["authorization"]
			kv := strings.Split(authString, " ")
			if len(kv) != 2 || kv[0] != "Bearer" {
				return errors.New("no auth token found in request")
			}
			token := kv[1]
			ret, err := client.ValidateToken(context.TODO(), &auth.ValidateTokenRequest{Token: token})
			if err != nil {
				return err
			}
			if !ret.Valid {
				return errors.New(ret.Error)
			}
			return h(ctx, req, rsp)
		}
	}
}