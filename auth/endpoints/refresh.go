package endpoints

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/sergiosegrera/store/auth/service"
)

type RefreshRequest struct {
	Token string `json:"token"`
}

type RefreshResponse struct {
	Token          string    `json:"token"`
	ExpirationTime time.Time `json:"expirationTime"`
}

func MakeRefreshEndpoint(svc service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RefreshRequest)
		token, expirationTime, err := svc.Refresh(ctx, req.Token)
		return RefreshResponse{Token: token, ExpirationTime: expirationTime}, err
	}
}
