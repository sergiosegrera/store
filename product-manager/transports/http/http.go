package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/sergiosegrera/store/product-manager/config"
	"github.com/sergiosegrera/store/product-manager/endpoints"
	"github.com/sergiosegrera/store/product-manager/service"
	"github.com/sergiosegrera/store/product-manager/transports/http/handlers"
)

func Serve(svc service.ProductManagerService, conf *config.Config) error {
	router := chi.NewRouter()

	getProducts := handlers.MakeGetProductsHandler(endpoints.MakeGetProductsEndpoint(svc))
	postProduct := handlers.MakePostProductHandler(endpoints.MakePostProductEndpoint(svc))
	deleteProduct := handlers.MakeDeleteProductHandler(endpoints.MakeDeleteProductEndpoint(svc))

	postOption := handlers.MakePostOptionHandler(endpoints.MakePostOptionEndpoint(svc))

	router.Get("/products", getProducts)
	router.Post("/product", postProduct)
	router.Delete("/product/{id}", deleteProduct)

	router.Post("/option", postOption)

	return http.ListenAndServe(fmt.Sprintf(":%v", conf.HttpPort), router)
}
