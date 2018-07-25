package handler

import (
	"context"
	"github.com/micro/go-log"
	"github.com/zergwangj/micro/api/auth/client"
	"github.com/micro/go-micro/errors"
	authsrv "github.com/zergwangj/micro/srv/auth/proto/auth"
	authapi "github.com/zergwangj/micro/api/auth/proto/micro/auth"
)

type Auth struct{}

// Auth.CreateSession is called by the API as /auth/createsession with post body {"username": "","password":""}
func (a *Auth) CreateSession(ctx context.Context, req *authapi.CreateSessionRequest, rsp *authapi.CreateSessionResponse) error {
	log.Log("Received Auth.CreateSession request")

	// extract the client from the context
	authClient, ok := client.AuthFromContext(ctx)
	if !ok {
		return errors.InternalServerError("zergwangj.api.auth.CreateSession", "auth client not found")
	}

	// make request
	response, err := authClient.CreateSession(ctx, &authsrv.CreateSessionRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return errors.InternalServerError("zergwangj.api.auth.CreateSession", err.Error())
	}

	rsp.Token = response.Token

	return nil
}

// Auth.ValidateToken is called by the API as /auth/validatetoken with post body {"token": ""}
func (a *Auth) ValidateToken(ctx context.Context, req *authapi.ValidateTokenRequest, rsp *authapi.ValidateTokenResponse) error {
	log.Log("Received Auth.ValidateToken request")

	// extract the client from the context
	authClient, ok := client.AuthFromContext(ctx)
	if !ok {
		return errors.InternalServerError("zergwangj.api.auth.ValidateToken", "auth client not found")
	}

	// make request
	response, err := authClient.ValidateToken(ctx, &authsrv.ValidateTokenRequest{
		Token: req.Token,
	})
	if err != nil {
		return errors.InternalServerError("zergwangj.api.auth.ValidateToken", err.Error())
	}

	rsp.Valid = response.Valid
	rsp.Error = response.Error

	return nil
}

// Auth.DeleteSession is called by the API as /auth/deletesession with post body {"username": "","token":""}
func (a *Auth) DeleteSession(ctx context.Context, req *authapi.DeleteSessionRequest, rsp *authapi.DeleteSessionResponse) error {
	log.Log("Received Auth.DeleteSession request")

	// extract the client from the context
	authClient, ok := client.AuthFromContext(ctx)
	if !ok {
		return errors.InternalServerError("zergwangj.api.auth.DeleteSession", "auth client not found")
	}

	// make request
	response, err := authClient.DeleteSession(ctx, &authsrv.DeleteSessionRequest{
		Username: req.Username,
		Token: req.Token,
	})
	if err != nil {
		return errors.InternalServerError("zergwangj.api.auth.DeleteSession", err.Error())
	}

	rsp.Ok = response.Ok

	return nil
}
