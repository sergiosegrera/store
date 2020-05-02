package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/sergiosegrera/store/product-manager/service"
)

type DeleteProductRequest struct {
	Id int64 `json:"id"`
}

type DeleteProductResponse struct {
	Error string `json:"error"`
}

func MakeDeleteProductEndpoint(svc service.ProductManagerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteProductRequest)
		err := svc.DeleteProduct(ctx, req.Id)
		if err != nil {
			return DeleteProductResponse{Error: err.Error()}, err
		}

		return DeleteProductResponse{Error: ""}, err
	}
}
