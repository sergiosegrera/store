package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/sergiosegrera/store/cart/pb"
	"github.com/sergiosegrera/store/checkout/service"
)

type PostCheckoutRequest struct {
	Cart pb.Cart `json:"cart"`
}

type PostCheckoutResponse struct {
	Token string `json:"link"`
}

func MakePostCheckoutEndpoint(svc service.CheckoutService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PostCheckoutRequest)
		token, err := svc.PostCheckout(ctx, req.Cart)
		return PostCheckoutResponse{Token: token}, err
	}
}
