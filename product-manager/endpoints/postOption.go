package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/sergiosegrera/store/product-manager/models"
	"github.com/sergiosegrera/store/product-manager/service"
)

type PostOptionRequest struct {
	Option models.Option `json:"option"`
}

type PostOptionResponse struct {
	Error string `json:"error"`
}

func MakePostOptionEndpoint(svc service.ProductManagerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PostOptionRequest)
		err := svc.PostOption(ctx, req.Option)
		if err != nil {
			return PostOptionResponse{Error: err.Error()}, err
		}
		return PostOptionResponse{Error: ""}, err
	}
}
