package main

import (
	"os"
	"os/signal"

	"github.com/sergiosegrera/store/auth/config"
	"github.com/sergiosegrera/store/auth/middlewares"
	"github.com/sergiosegrera/store/auth/service"
	"github.com/sergiosegrera/store/auth/transports/http"
	"go.uber.org/zap"
)

func main() {
	conf := config.New()

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	authService := service.NewService(conf)
	authService = middlewares.Logging{logger, authService}

	go func() {
		logger.Info("started the http server", zap.String("port", conf.HttpPort))
		err := http.Serve(authService, conf)
		if err != nil {
			logger.Error("the http server panicked", zap.String("err", err.Error()))
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	logger.Info("exited", zap.String("sig", sig.String()))
}
