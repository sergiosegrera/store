package bindings

import (
	"context"

	"github.com/sergiosegrera/store/auth/pb"
)

func (b GrpcBinding) Check(ctx context.Context, in *pb.Token) (*pb.Valid, error) {
	err := b.AuthService.Check(ctx, in.Token)
	if err != nil {
		return &pb.Valid{Valid: false}, err
	}

	return &pb.Valid{Valid: true}, err
}
