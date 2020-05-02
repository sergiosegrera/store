package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/sergiosegrera/store/cart/models"
	"github.com/sergiosegrera/store/cart/service"
)

type PostCartRequest struct {
	Cart models.Cart `json:"cart"`
}

type PostCartResponse struct {
	Cart models.Cart `json:"cart"`
}

func MakePostCartEndpoint(svc service.CartService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PostCartRequest)
		cart := svc.PostCart(req.Cart)
		return PostCartResponse{Cart: cart}, nil
	}
}
