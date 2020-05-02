package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sergiosegrera/store/checkout/endpoints"
	"github.com/sergiosegrera/store/checkout/service"
	"github.com/sergiosegrera/store/checkout/transport/http/handlers"
)

func Serve(svc *service.Service) error {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	postCart := handlers.MakePostCheckoutHandler(endpoints.MakePostCheckoutEndpoint(svc))

	router.Post("/checkout", postCart)

	return http.ListenAndServe(":8080", router)
}
