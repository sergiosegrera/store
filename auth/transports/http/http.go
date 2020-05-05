package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sergiosegrera/store/auth/config"
	"github.com/sergiosegrera/store/auth/endpoints"
	"github.com/sergiosegrera/store/auth/service"
	"github.com/sergiosegrera/store/auth/transports/http/handlers"
)

func Serve(svc service.AuthService, conf *config.Config) error {
	router := chi.NewRouter()
	router.Use(middleware.Compress(5, "gzip"))

	login := handlers.MakeLoginHandler(endpoints.MakeLoginEndpoint(svc))
	refresh := handlers.MakeRefreshHandler(endpoints.MakeRefreshEndpoint(svc))

	router.Post("/login", login)
	router.Post("/refresh", refresh)

	return http.ListenAndServe(fmt.Sprintf(":%v", conf.HttpPort), router)
}
