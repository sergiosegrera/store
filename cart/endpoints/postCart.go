package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/sergiosegrera/store/cart/pb"
	"github.com/sergiosegrera/store/cart/service"
)

type PostCartRequest struct {
	Cart pb.Cart `json:"cart"`
}

type PostCartResponse struct {
	Cart pb.Cart `json:"cart"`
}

func MakePostCartEndpoint(svc service.CartService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PostCartRequest)
		cart, err := svc.PostCart(ctx, req.Cart)
		return PostCartResponse{Cart: cart}, err
	}
}
