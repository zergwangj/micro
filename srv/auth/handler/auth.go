package handler

import (
	"context"
	"github.com/muesli/cache2go"
	"github.com/micro/go-log"
	auth "github.com/zergwangj/micro/srv/auth/proto/auth"
	"time"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"database/sql"
	"github.com/zergwangj/micro/srv/users/models"
	"github.com/micro/go-config"
)

type Auth struct {
	Db			*sql.DB
	Cache		*cache2go.CacheTable
	Conf		config.Config
}

// CreateSession
func (a *Auth) CreateSession(ctx context.Context, req *auth.CreateSessionRequest, rsp *auth.CreateSessionResponse) error {
	log.Log("Received Auth.CreateSession request")

	db := a.Db
	cache := a.Cache
	conf := a.Conf

	jwtSecret := conf.Get("jwt", "secret").String("secret")
	jwtTtl := conf.Get("jwt", "ttl").Int(180000)

	user, err := models.GetUserByName(db, req.GetUsername())
	if err != nil {
		return err
	}

	if user.Password != req.GetPassword() {
		return fmt.Errorf("Password is not valid")
	}

	iat := time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat":      iat,
		"exp":      iat + int64(jwtTtl),
		"username": req.GetUsername(),
		"password": req.GetPassword(),
	})

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return err
	}

	cache.Add(tokenString, time.Duration(jwtTtl * 1000000000), user)

	rsp.Token = tokenString

	return nil
}

// ValidateToken
func (a *Auth) ValidateToken(ctx context.Context, req *auth.ValidateTokenRequest, rsp *auth.ValidateTokenResponse) error {
	log.Log("Received Auth.ValidateSession request")

	cache := a.Cache

	tokenString := req.GetToken()

	_, err := cache.Value(tokenString)
	if err != nil {
		rsp.Valid = false
		rsp.Error = err.Error()
		return nil
	}

	rsp.Valid = true
	rsp.Error = ""

	return nil
}

// DeleteSession
func (a *Auth) DeleteSession(ctx context.Context, req *auth.DeleteSessionRequest, rsp *auth.DeleteSessionResponse) error {
	log.Log("Received Auth.DeleteSession request")

	cache := a.Cache

	tokenString := req.GetToken()
	username := req.GetUsername()

	res, err := cache.Value(tokenString)
	if err != nil {
		rsp.Ok = false
		return nil
	}

	if res.Data().(*models.User).Username != username {
		rsp.Ok = false
		return nil
	}

	cache.Delete(tokenString)
	rsp.Ok = true

	return nil
}
