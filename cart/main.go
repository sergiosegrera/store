package main

import (
	"os"
	"os/signal"

	"github.com/go-pg/pg/v9"
	"github.com/sergiosegrera/store/cart/config"
	"github.com/sergiosegrera/store/cart/db"
	"github.com/sergiosegrera/store/cart/middlewares"
	"github.com/sergiosegrera/store/cart/service"
	grpctransport "github.com/sergiosegrera/store/cart/transports/grpc"
	httptransport "github.com/sergiosegrera/store/cart/transports/http"
	"go.uber.org/zap"
)

func main() {
	conf := config.New()

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	// Connect to DB
	options := &pg.Options{
		Addr:     conf.DatabaseAddress,
		User:     "product",
		Database: "product",
		Password: "verysecuremuchwow",
	}

	db, err := db.NewConnection(options)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	cartService := service.NewService(db, conf)
	cartService = middlewares.Logging{logger, cartService}

	go func() {
		logger.Info("started the http server", zap.String("port", conf.HttpPort))
		err := httptransport.Serve(cartService, conf)
		if err != nil {
			logger.Error("the http server panicked", zap.String("err", err.Error()))
			os.Exit(1)
		}
	}()

	go func() {
		logger.Info("started the grpc server", zap.String("port", conf.GrpcPort))
		err := grpctransport.Serve(cartService, conf)
		if err != nil {
			logger.Error("the grpc server panicked", zap.String("err", err.Error()))
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	logger.Info("exited", zap.String("sig", sig.String()))
}
