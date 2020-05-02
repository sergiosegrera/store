package handlers

import (
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func MakeGetProductsHandler(e endpoint.Endpoint) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// No decoding step, endpoint requires no input
		// Logic step
		response, err := e(r.Context(), nil)

		if err != nil {
			JSON(w, 404, payload{"error": "No products to show"})
			return
		}

		// Encoding
		JSON(w, 200, response)
	}
}
