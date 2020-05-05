package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/sergiosegrera/store/product-manager/endpoints"
)

func MakePostProductHandler(e endpoint.Endpoint) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var request endpoints.PostProductRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			JSON(w, 400, payload{"error": "Could not parse input"})
			return
		}

		_, err = e(r.Context(), request)
		if err != nil {
			JSON(w, 404, payload{"error": "Error Posting item"})
			return
		}

		JSON(w, 200, payload{"message": "Success"})
	}
}
