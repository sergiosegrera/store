package middlewares

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	authclient "github.com/sergiosegrera/store/auth/clients/grpc"
)

func Auth(client *authclient.Client) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			token, ok := ctx.Value("token").(string)

			if !ok {
				return nil, ErrNeedAuth
			}
			// TODO: Add timeout?
			valid, err := client.Check(ctx, token)
			if err != nil || !valid {
				return nil, ErrNeedAuth
			}
			return next(ctx, request)
		}
	}
}

var (
	ErrNeedAuth = errors.New("Authentication needed")
)
