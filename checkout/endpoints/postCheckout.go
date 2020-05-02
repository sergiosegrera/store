package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/sergiosegrera/store/cart/models"
	"github.com/sergiosegrera/store/checkout/service"
)

type PostCheckoutRequest struct {
	Cart models.Cart `json:"cart"`
}

type PostCheckoutResponse struct {
	Token string `json:"link"`
}

func MakePostCheckoutEndpoint(svc service.CartService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PostCheckoutRequest)
		token := svc.PostCheckout(req.Cart)
		return PostCheckoutResponse{Token: token}, nil
	}
}
