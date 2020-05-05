package main

import (
	"context"
	"os"
	"os/signal"

	cartclient "github.com/sergiosegrera/store/cart/clients/grpc"
	"github.com/sergiosegrera/store/checkout/config"
	"github.com/sergiosegrera/store/checkout/middlewares"
	"github.com/sergiosegrera/store/checkout/service"
	"github.com/sergiosegrera/store/checkout/transports/http"
	stripeclient "github.com/stripe/stripe-go/client"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	conf := config.New()

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	// Connect to DB
	// options := &pg.Options{
	// 	Addr:     "db:5432",
	// 	User:     "product",
	// 	Database: "product",
	// 	Password: "verysecuremuchwow",
	// }

	// db, err := db.NewConnection(options)
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()
	// Start stripe client connection
	sc := stripeclient.New(conf.StripeSecret, nil)

	// Start attach db and start http server
	conn, err := grpc.Dial(conf.CartGrpcAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	cc := cartclient.New(context.Background(), conn)

	checkoutService := service.NewService(cc, sc)
	checkoutService = middlewares.Logging{logger, checkoutService}

	go func() {
		logger.Info("started the http server", zap.String("port", conf.HttpPort))
		err := http.Serve(checkoutService, conf)
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
