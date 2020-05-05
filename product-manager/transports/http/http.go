package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"google.golang.org/grpc"

	authclient "github.com/sergiosegrera/store/auth/clients/grpc"
	"github.com/sergiosegrera/store/product-manager/config"
	"github.com/sergiosegrera/store/product-manager/endpoints"
	"github.com/sergiosegrera/store/product-manager/middlewares"
	"github.com/sergiosegrera/store/product-manager/service"
	"github.com/sergiosegrera/store/product-manager/transports/http/handlers"
	httpmiddlewares "github.com/sergiosegrera/store/product-manager/transports/http/middlewares"
)

func Serve(svc service.ProductManagerService, conf *config.Config, cc *grpc.ClientConn) error {
	router := chi.NewRouter()
	router.Use(middleware.Compress(5, "gzip"))
	router.Use(httpmiddlewares.Auth())

	authClient := authclient.New(cc)
	// TODO: Better way to create endpoints?
	getProductsEndpoint := middlewares.Auth(authClient)(endpoints.MakeGetProductsEndpoint(svc))
	getProductEndpoint := middlewares.Auth(authClient)(endpoints.MakeGetProductEndpoint(svc))
	postProductEndpoint := middlewares.Auth(authClient)(endpoints.MakePostProductEndpoint(svc))
	deleteProductEndpoint := middlewares.Auth(authClient)(endpoints.MakeDeleteProductEndpoint(svc))

	postOptionEndpoint := middlewares.Auth(authClient)(endpoints.MakePostOptionEndpoint(svc))
	deleteOptionEndpoint := middlewares.Auth(authClient)(endpoints.MakeDeleteOptionEndpoint(svc))

	getProducts := handlers.MakeGetProductsHandler(getProductsEndpoint)
	getProduct := handlers.MakeGetProductHandler(getProductEndpoint)
	postProduct := handlers.MakePostProductHandler(postProductEndpoint)
	deleteProduct := handlers.MakeDeleteProductHandler(deleteProductEndpoint)

	postOption := handlers.MakePostOptionHandler(postOptionEndpoint)
	deleteOption := handlers.MakeDeleteOptionHandler(deleteOptionEndpoint)

	router.Get("/products", getProducts)
	router.Get("/product/{id}", getProduct)
	router.Post("/product", postProduct)
	router.Delete("/product/{id}", deleteProduct)

	router.Post("/option", postOption)
	router.Post("/option/{id}", deleteOption)

	return http.ListenAndServe(fmt.Sprintf(":%v", conf.HttpPort), router)
}
