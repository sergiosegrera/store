package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/sergiosegrera/store/checkout/service"
)

type PostConfirmRequest struct {
	Id string `json:"id"`
}

type PostConfirmResponse struct {
	Error string `json:"error"`
}

func MakePostConfirmEndpoint(svc service.CheckoutService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PostConfirmRequest)
		err := svc.PostConfirm(ctx, req.Id)
		if err != nil {
			return PostConfirmResponse{Error: err.Error()}, err
		}
		return PostConfirmResponse{Error: ""}, err
	}
}
