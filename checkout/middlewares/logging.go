package middlewares

import (
	"context"
	"time"

	"github.com/sergiosegrera/store/cart/pb"
	"github.com/sergiosegrera/store/checkout/service"
	"go.uber.org/zap"
)

type Logging struct {
	Logger *zap.Logger
	Next   service.CheckoutService
}

func (mw Logging) PostCheckout(ctx context.Context, cart pb.Cart) (output string, err error) {
	defer func(begin time.Time) {
		mw.Logger.Info(
			"checkout",
			zap.String("method", "postcheckout"),
			zap.NamedError("err", err),
			zap.Duration("took", time.Since(begin)),
		)
	}(time.Now())

	output, err = mw.Next.PostCheckout(ctx, cart)
	return
}

func (mw Logging) PostConfirm(ctx context.Context, id string) (err error) {
	defer func(begin time.Time) {
		mw.Logger.Info(
			"checkout",
			zap.String("method", "postconfirm"),
			zap.String("id", id),
			zap.NamedError("err", err),
			zap.Duration("took", time.Since(begin)),
		)
	}(time.Now())

	err = mw.Next.PostConfirm(ctx, id)
	return
}
