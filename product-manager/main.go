package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/go-pg/pg/v9"
	"github.com/sergiosegrera/store/product-manager/db"
	"github.com/sergiosegrera/store/product-manager/service"
	"github.com/sergiosegrera/store/product-manager/transport/http"
)

func main() {
	// Connect to DB
	options := &pg.Options{
		Addr:     "db:5432",
		User:     "product",
		Database: "product",
		Password: "verysecuremuchwow",
	}

	db, err := db.NewConnection(options)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Add mock data
	// product := models.Product{
	// 	Name:        "White T-Shirt",
	// 	Thumbnail:   "https://imgur.com/qEOvdMp",
	// 	Images:      []string{"https://imgur.com/qEOvdMp", "https://imgur.com/qEOvdMp"},
	// 	Description: "Plain white T-Shirt",
	// 	Price:       30,
	// 	Public:      true,
	// }

	// result, err := db.Model(&product).Returning("id").Insert()
	// if err != nil {
	// 	panic(err)
	// }

	// option := models.Option{
	// 	ProductId: int64(result.RowsReturned()),
	// 	Name:      "Small",
	// 	Stock:     30,
	// }

	// err = db.Insert(&option)
	// if err != nil {
	// 	panic(err)
	// }

	// Start attach db and start http server
	go func() {
		log.Println("Started the http server")
		err := http.Serve(service.NewService(db))
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
