package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/sergiosegrera/store/auth/endpoints"
	"github.com/sergiosegrera/store/auth/service"
)

func MakeRefreshHandler(e endpoint.Endpoint) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		c, err := r.Cookie("token")
		if err != nil {
			switch err {
			case http.ErrNoCookie:
				w.WriteHeader(401)
				json.NewEncoder(w).Encode(data{"error": "Invalid token"})
				return
			default:
				w.WriteHeader(400)
				json.NewEncoder(w).Encode(data{"error": "Bad token"})
				return
			}
		}

		response, err := e(r.Context(), endpoints.RefreshRequest{Token: c.Value})
		if err != nil {
			switch err {
			case service.ErrInvalidToken:
				w.WriteHeader(401)
				json.NewEncoder(w).Encode(data{"error": "Invalid token"})
				return
			case service.ErrBadToken:
				w.WriteHeader(400)
				json.NewEncoder(w).Encode(data{"error": "Bad token"})
				return
			default:
				w.WriteHeader(500)
				json.NewEncoder(w).Encode(data{"error": "Internal server error"})
				return
			}
		}

		token := response.(endpoints.RefreshResponse)

		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   token.Token,
			Expires: token.ExpirationTime,
		})

		w.WriteHeader(200)
		json.NewEncoder(w).Encode(response)
	}
}
