package main

import (
	"os"
	"os/signal"

	"github.com/go-pg/pg/v9"
	"github.com/sergiosegrera/store/product-manager/config"
	"github.com/sergiosegrera/store/product-manager/db"
	"github.com/sergiosegrera/store/product-manager/middlewares"
	"github.com/sergiosegrera/store/product-manager/service"
	"github.com/sergiosegrera/store/product-manager/transports/http"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	conf := config.New()

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

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

	// Auth service conenction
	conn, err := grpc.Dial(conf.AuthGrpcAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	productManagerService := service.NewService(db)
	productManagerService = middlewares.Logging{logger, productManagerService}

	// Start attach db and start http server
	go func() {
		logger.Info("started the http server", zap.String("port", conf.HttpPort))
		err := http.Serve(productManagerService, conf, conn)
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
