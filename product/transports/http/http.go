package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/sergiosegrera/store/product/config"
	"github.com/sergiosegrera/store/product/endpoints"
	"github.com/sergiosegrera/store/product/service"
	"github.com/sergiosegrera/store/product/transports/http/handlers"
)

func Serve(svc service.ProductService, conf *config.Config) error {
	router := chi.NewRouter()
	router.Use(middleware.Compress(5, "gzip"))

	getProducts := handlers.MakeGetProductsHandler(endpoints.MakeGetProductsEndpoint(svc))
	getProduct := handlers.MakeGetProductHandler(endpoints.MakeGetProductEndpoint(svc))

	router.Get("/products", getProducts)
	router.Get("/product/{id}", getProduct)

	return http.ListenAndServe(fmt.Sprintf(":%v", conf.HttpPort), router)
}
