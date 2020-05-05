package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/sergiosegrera/store/cart/endpoints"
)

type data map[string]interface{}

func MakePostCartHandler(e endpoint.Endpoint) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var request endpoints.PostCartRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(data{"error": "Could not parse input"})
			return
		}

		response, err := e(r.Context(), request)
		// TODO: Check for different errors, change response accordingly
		if err != nil {
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(data{"error": "Internal server error"})
			return
		}

		w.WriteHeader(200)
		json.NewEncoder(w).Encode(response)
	}
}
