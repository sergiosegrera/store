package middlewares

import (
	"context"
	"time"

	"github.com/sergiosegrera/store/product-manager/models"
	"github.com/sergiosegrera/store/product-manager/service"
	"go.uber.org/zap"
)

type Logging struct {
	Logger *zap.Logger
	Next   service.ProductManagerService
}

func (mw Logging) GetProducts(ctx context.Context) (output []*models.Product, err error) {
	defer func(begin time.Time) {
		mw.Logger.Info(
			"product-manager",
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
			"product-manager",
			zap.String("method", "getproduct"),
			zap.Int64("id", id),
			zap.NamedError("err", err),
			zap.Duration("took", time.Since(begin)),
		)
	}(time.Now())

	output, err = mw.Next.GetProduct(ctx, id)
	return
}

func (mw Logging) PostProduct(ctx context.Context, product models.Product) (err error) {
	defer func(begin time.Time) {
		mw.Logger.Info(
			"product-manager",
			zap.String("method", "postproduct"),
			zap.NamedError("err", err),
			zap.Duration("took", time.Since(begin)),
		)
	}(time.Now())

	err = mw.Next.PostProduct(ctx, product)
	return
}

func (mw Logging) DeleteProduct(ctx context.Context, id int64) (err error) {
	defer func(begin time.Time) {
		mw.Logger.Info(
			"product-manager",
			zap.String("method", "deleteproduct"),
			zap.Int64("id", id),
			zap.NamedError("err", err),
			zap.Duration("took", time.Since(begin)),
		)
	}(time.Now())

	err = mw.Next.DeleteProduct(ctx, id)
	return
}

func (mw Logging) PostOption(ctx context.Context, option models.Option) (err error) {
	defer func(begin time.Time) {
		mw.Logger.Info(
			"product-manager",
			zap.String("method", "postoption"),
			zap.NamedError("err", err),
			zap.Duration("took", time.Since(begin)),
		)
	}(time.Now())

	err = mw.Next.PostOption(ctx, option)
	return
}

func (mw Logging) DeleteOption(ctx context.Context, id int64) (err error) {
	defer func(begin time.Time) {
		mw.Logger.Info(
			"product-manager",
			zap.String("method", "deleteoption"),
			zap.Int64("id", id),
			zap.NamedError("err", err),
			zap.Duration("took", time.Since(begin)),
		)
	}(time.Now())

	err = mw.Next.DeleteOption(ctx, id)
	return
}
