package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/sergiosegrera/store/checkout/endpoints"
	"github.com/sergiosegrera/store/checkout/service"
)

func MakePostCheckoutHandler(e endpoint.Endpoint) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var request endpoints.PostCheckoutRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			JSON(w, 400, payload{"error": "Could not parse input"})
			return
		}

		response, err := e(r.Context(), request)
		if err != nil {
			switch err {
			case service.ErrCouldNotVerifyCart:
				JSON(w, 500, payload{"error": err.Error()})
				return
			case service.ErrNoProductsInCart:
				JSON(w, 400, payload{"error": err.Error()})
				return
			case service.ErrCreatingToken:
				JSON(w, 500, payload{"error": err.Error()})
				return
			default:
				JSON(w, 500, payload{"error": "Unknown error"})
				return
			}
		}

		JSON(w, 200, response)
	}
}
