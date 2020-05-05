package client

import (
	"context"

	"github.com/sergiosegrera/store/cart/pb"
	"google.golang.org/grpc"
)

type Client struct {
	ctx    context.Context
	client pb.CartServiceClient
}

func New(ctx context.Context, cc *grpc.ClientConn) *Client {
	return &Client{
		ctx,
		pb.NewCartServiceClient(cc),
	}
}

func (c *Client) PostCart(cart *pb.Cart) (*pb.Cart, error) {
	result, err := c.client.PostCart(c.ctx, cart)
	if err != nil {
		return nil, err
	}

	return result, err
}
