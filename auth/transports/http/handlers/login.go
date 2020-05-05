package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/sergiosegrera/store/auth/endpoints"
	"github.com/sergiosegrera/store/auth/service"
)

func MakeLoginHandler(e endpoint.Endpoint) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var request endpoints.LoginRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(data{"error": "Could not parse input"})
			return
		}

		response, err := e(r.Context(), request)
		if err != nil {
			switch err {
			case service.ErrWrongPassword:
				w.WriteHeader(401)
				json.NewEncoder(w).Encode(data{"error": "Wrong password"})
				return
			default:
				w.WriteHeader(500)
				json.NewEncoder(w).Encode(data{"error": "Internal server error"})
				return
			}
		}

		token := response.(endpoints.LoginResponse)

		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   token.Token,
			Expires: token.ExpirationTime,
		})

		w.WriteHeader(200)
		json.NewEncoder(w).Encode(response)
	}
}
