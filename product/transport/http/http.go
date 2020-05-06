package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/sergiosegrera/store/product/config"
	"github.com/sergiosegrera/store/product/endpoints"
	"github.com/sergiosegrera/store/product/service"
	"github.com/sergiosegrera/store/product/transport/http/handlers"
)

func Serve(svc *service.Service, conf *config.Config) error {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	getProducts := handlers.MakeGetProductsHandler(endpoints.MakeGetProductsEndpoint(svc))
	getProduct := handlers.MakeGetProductHandler(endpoints.MakeGetProductEndpoint(svc))

	router.Get("/products", getProducts)
	router.Get("/product/{id}", getProduct)

	return http.ListenAndServe(fmt.Sprintf(":%v", conf.HttpPort), router)
}
