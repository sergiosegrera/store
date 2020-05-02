package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/sergiosegrera/store/cart/endpoints"
)

func MakePostCartHandler(e endpoint.Endpoint) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var request endpoints.PostCartRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			JSON(w, 400, payload{"error": "Could not parse input"})
			return
		}

		response, err := e(r.Context(), request)
		if err != nil {
			JSON(w, 500, payload{"error": "Internal server error"})
			return
		}

		JSON(w, 200, response)
	}
}
