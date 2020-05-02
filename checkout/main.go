package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	cartClient "github.com/sergiosegrera/store/cart/client/grpc"
	"github.com/sergiosegrera/store/checkout/service"
	"github.com/sergiosegrera/store/checkout/transport/http"
	"google.golang.org/grpc"
)

func main() {
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

	// Start attach db and start http server
	s := grpc.Dial("cart:8080", grpc.WithInsecure())
	client := cartClient.New(context.Background(), s)

	go func() {
		log.Println("Started the http server")
		err := http.Serve(service.NewService(client))
		if err != nil {
			log.Println("The http server panicked:", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	log.Println("Got signal:", sig)
}
