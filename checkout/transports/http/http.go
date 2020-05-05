package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sergiosegrera/store/checkout/config"
	"github.com/sergiosegrera/store/checkout/endpoints"
	"github.com/sergiosegrera/store/checkout/service"
	"github.com/sergiosegrera/store/checkout/transports/http/handlers"
)

func Serve(svc service.CheckoutService, conf *config.Config) error {
	router := chi.NewRouter()
	router.Use(middleware.Compress(5, "gzip"))

	postCart := handlers.MakePostCheckoutHandler(endpoints.MakePostCheckoutEndpoint(svc))

	router.Post("/checkout", postCart)

	return http.ListenAndServe(fmt.Sprintf(":%v", conf.HttpPort), router)
}
