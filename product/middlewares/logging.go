package middlewares

import (
	"context"
	"time"

	"github.com/sergiosegrera/store/product/models"
	"github.com/sergiosegrera/store/product/service"
	"go.uber.org/zap"
)

type Logging struct {
	Logger *zap.Logger
	Next   service.ProductService
}

func (mw Logging) GetProducts(ctx context.Context) (output []*models.Thumbnail, err error) {
	defer func(begin time.Time) {
		mw.Logger.Info(
			"product",
			zap.String("method", "getproducts"),
			zap.NamedError("err", err),
			zap.Duration("took", time.Since(begin)),
		)
	}(time.Now())

	output, err = mw.Next.GetProducts(ctx)
	return
}

func (mw Logging) GetProduct(ctx context.Context, id int64) (output *models.Product, err error) {
	defer func(begin time.Time) {
		mw.Logger.Info(
			"product",
			zap.String("method", "getproduct"),
			zap.Int64("id", id),
			zap.NamedError("err", err),
			zap.Duration("took", time.Since(begin)),
		)
	}(time.Now())

	output, err = mw.Next.GetProduct(ctx, id)
	return
}
