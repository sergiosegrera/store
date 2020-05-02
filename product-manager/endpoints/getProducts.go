package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/sergiosegrera/store/product-manager/models"
	"github.com/sergiosegrera/store/product-manager/service"
)

type GetProductsResponse struct {
	Products []*models.Product `json:"products"`
}

func MakeGetProductsEndpoint(svc service.ProductManagerService) endpoint.Endpoint {
	return func(ctx context.Context, _ interface{}) (interface{}, error) {
		products, err := svc.GetProducts(ctx)
		return GetProductsResponse{Products: products}, err
	}
}
