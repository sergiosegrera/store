package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/sergiosegrera/store/auth/service"
)

type CheckRequest struct {
	Token string `json:"token"`
}

type CheckResponse struct {
	Error string `json:"error"`
}

func MakeCheckEndpoint(svc service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CheckRequest)
		err := svc.Check(ctx, req.Token)
		if err != nil {
			return CheckResponse{Error: err.Error()}, err
		}
		return CheckResponse{Error: ""}, err
	}
}
