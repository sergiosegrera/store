package bindings

import (
	"context"

	"github.com/sergiosegrera/store/cart/pb"
)

func (b GrpcBinding) PostCart(ctx context.Context, in *pb.Cart) (*pb.Cart, error) {
	result, err := b.CartService.PostCart(ctx, *in)

	return &result, err
}
