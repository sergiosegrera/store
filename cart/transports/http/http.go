package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sergiosegrera/store/cart/config"
	"github.com/sergiosegrera/store/cart/endpoints"
	"github.com/sergiosegrera/store/cart/service"
	"github.com/sergiosegrera/store/cart/transports/http/handlers"
)

func Serve(svc service.CartService, conf *config.Config) error {
	router := chi.NewRouter()

	postCart := handlers.MakePostCartHandler(endpoints.MakePostCartEndpoint(svc))

	router.Post("/cart", postCart)

	return http.ListenAndServe(
		fmt.Sprintf(":%v", conf.HttpPort),
		router,
	)
}
