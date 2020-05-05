package client

import (
	"context"

	"github.com/sergiosegrera/store/auth/pb"
	"google.golang.org/grpc"
)

type Client struct {
	client pb.AuthServiceClient
}

func New(cc *grpc.ClientConn) *Client {
	return &Client{
		pb.NewAuthServiceClient(cc),
	}
}

func (c *Client) Check(ctx context.Context, token string) (bool, error) {
	result, err := c.client.Check(ctx, &pb.Token{Token: token})
	if err != nil {
		return false, err
	}
	return result.Valid, err
}
