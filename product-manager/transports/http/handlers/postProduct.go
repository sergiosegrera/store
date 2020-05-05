package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/sergiosegrera/store/product-manager/endpoints"
)

func MakePostProductHandler(e endpoint.Endpoint) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			switch err {
			case http.ErrNoCookie:
				w.WriteHeader(401)
				json.NewEncoder(w).Encode(payload{"error": "Invalid token"})
				return
			default:
				w.WriteHeader(400)
				json.NewEncoder(w).Encode(payload{"error": "Bad token"})
				return
			}
		}

		var request endpoints.PostProductRequest
		err = json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			JSON(w, 400, payload{"error": "Could not parse input"})
			return
		}

		_, err = e(context.WithValue(r.Context(), "token", c.Value), request)
		if err != nil {
			JSON(w, 404, payload{"error": err.Error()})
			return
		}

		JSON(w, 200, payload{"message": "Success"})
	}
}
