package middlewares

import (
	"context"
	"time"

	"github.com/sergiosegrera/store/auth/service"
	"go.uber.org/zap"
)

type Logging struct {
	Logger *zap.Logger
	Next   service.AuthService
}

func (mw Logging) Login(ctx context.Context, password string) (token string, expires time.Time, err error) {
	defer func(begin time.Time) {
		mw.Logger.Info(
			"auth",
			zap.String("method", "login"),
			zap.NamedError("err", err),
			zap.Duration("took", time.Since(begin)),
		)
	}(time.Now())

	token, expires, err = mw.Next.Login(ctx, password)
	return
}

func (mw Logging) Check(ctx context.Context, token string) (err error) {
	defer func(begin time.Time) {
		mw.Logger.Info(
			"auth",
			zap.String("method", "check"),
			zap.NamedError("err", err),
			zap.Duration("took", time.Since(begin)),
		)
	}(time.Now())

	err = mw.Next.Check(ctx, token)
	return
}

func (mw Logging) Refresh(ctx context.Context, token string) (output string, expires time.Time, err error) {
	defer func(begin time.Time) {
		mw.Logger.Info(
			"auth",
			zap.String("method", "refresh"),
			zap.NamedError("err", err),
			zap.Duration("took", time.Since(begin)),
		)
	}(time.Now())

	output, expires, err = mw.Next.Refresh(ctx, token)
	return
}
