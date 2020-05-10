package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-kit/kit/endpoint"
	"github.com/sergiosegrera/store/checkout/endpoints"
)

func MakePostConfirmHandler(e endpoint.Endpoint) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: verify id?
		id := chi.URLParam(r, "id")

		response, err := e(r.Context(), endpoints.PostConfirmRequest{Id: id})
		if err != nil {
			switch err {
			default:
				JSON(w, 500, payload{"error": "Internal server error"})
			}
		}

		JSON(w, 200, response)
	}
}
