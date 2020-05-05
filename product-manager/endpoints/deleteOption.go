package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/sergiosegrera/store/product-manager/service"
)

type DeleteOptionRequest struct {
	Id int64 `json:"id"`
}

type DeleteOptionResponse struct {
	Error string `json:"error"`
}

func MakeDeleteOptionEndpoint(svc service.ProductManagerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteOptionRequest)
		err := svc.DeleteOption(ctx, req.Id)
		if err != nil {
			return DeleteOptionResponse{Error: err.Error()}, err
		}

		return DeleteOptionResponse{Error: ""}, err
	}
}
