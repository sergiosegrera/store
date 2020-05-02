package grpc

import (
	"context"
	"log"

	"github.com/sergiosegrera/store/cart/models"
	"github.com/sergiosegrera/store/cart/service"
	"github.com/sergiosegrera/store/cart/transport/grpc/pb"
	"google.golang.org/grpc"
)

type Client struct {
	ctx    context.Context
	client pb.CartServiceClient
}

func New(ctx context.Context, cc *grpc.ClientConn) service.CartService {
	return Client{ctx, pb.NewCartServiceClient(cc)}
}

func (c Client) PostCart(cart models.Cart) models.Cart {
	// Decode
	var request pb.Cart
	for _, product := range cart.CartProducts {
		request.CartProducts = append(request.CartProducts, &pb.CartProduct{
			Id:       product.Id,
			OptionId: product.OptionId,
			Count:    product.Count,
			Price:    product.Price,
		})
	}

	// Logic
	result, err := c.client.PostCart(c.ctx, &request)
	if err != nil {
		log.Println(err)
	}

	// Encode
	var response models.Cart
	for _, product := range result.CartProducts {
		response.CartProducts = append(response.CartProducts, &models.CartProduct{
			Id:       product.Id,
			OptionId: product.OptionId,
			Count:    product.Count,
			Price:    product.Price,
		})
	}
	return response
}
