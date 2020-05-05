package endpoints

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/sergiosegrera/store/auth/service"
)

type LoginRequest struct {
	Password string `json:"password"`
}

type LoginResponse struct {
	Token          string    `json:"token"`
	ExpirationTime time.Time `json:"expirationTime"`
}

func MakeLoginEndpoint(svc service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginRequest)
		token, expiration, err := svc.Login(ctx, req.Password)
		return LoginResponse{Token: token, ExpirationTime: expiration}, err
	}
}
