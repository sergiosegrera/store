package bindings

import (
	"context"

	"github.com/sergiosegrera/store/cart/models"
	"github.com/sergiosegrera/store/cart/service"
	"github.com/sergiosegrera/store/cart/transport/grpc/pb"
)

type GRPCBinding struct {
	service.CartService
}

func (b GRPCBinding) PostCart(ctx context.Context, in *pb.Cart) (*pb.Cart, error) {
	// Decode
	var request models.Cart
	for _, product := range in.CartProducts {
		request.CartProducts = append(request.CartProducts, &models.CartProduct{
			Id:       product.Id,
			OptionId: product.OptionId,
			Count:    product.Count,
			Price:    product.Price,
		})
	}

	// Logic
	result := b.CartService.PostCart(request)

	// Encode
	var response pb.Cart
	for _, product := range result.CartProducts {
		response.CartProducts = append(response.CartProducts, &pb.CartProduct{
			Id:       product.Id,
			OptionId: product.OptionId,
			Count:    product.Count,
			Price:    product.Price,
		})
	}
	return &response, nil
}
