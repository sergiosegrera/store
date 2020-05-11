package main

import (
	"os"
	"os/signal"

	"github.com/go-pg/pg/v9"
	"github.com/sergiosegrera/store/product/config"
	"github.com/sergiosegrera/store/product/db"
	"github.com/sergiosegrera/store/product/middlewares"
	"github.com/sergiosegrera/store/product/service"
	"github.com/sergiosegrera/store/product/transports/http"
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

	productService := service.NewService(db)
	productService = middlewares.Logging{logger, productService}

	// Start attach db and start http server
	go func() {
		logger.Info("started the http server", zap.String("port", conf.HttpPort))
		err := http.Serve(productService, conf)
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
