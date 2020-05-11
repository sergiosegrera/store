package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-kit/kit/endpoint"
	"github.com/sergiosegrera/store/product/endpoints"
)

func MakeGetProductHandler(e endpoint.Endpoint) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decoding step
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			JSON(w, 400, payload{"error": "Could not parse id"})
			return
		}

		request := endpoints.GetProductRequest{Id: id}

		// Logic step
		response, err := e(r.Context(), request)

		if err != nil {
			JSON(w, 404, payload{"error": "Product not found"})
			return
		}

		// Encoding step
		JSON(w, 200, response)
	}
}
