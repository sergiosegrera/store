package middlewares

import (
	"context"
	"time"

	"github.com/sergiosegrera/store/cart/pb"
	"github.com/sergiosegrera/store/cart/service"
	"go.uber.org/zap"
)

type Logging struct {
	Logger *zap.Logger
	Next   service.CartService
}

func (mw Logging) PostCart(ctx context.Context, cart pb.Cart) (output pb.Cart, err error) {
	defer func(begin time.Time) {
		mw.Logger.Info("cart",
			zap.String("method", "postcart"),
			zap.NamedError("err", err),
			zap.Duration("took", time.Since(begin)),
		)
	}(time.Now())

	output, err = mw.Next.PostCart(ctx, cart)
	return
}
