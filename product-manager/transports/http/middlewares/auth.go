package middlewares

import (
	"context"
	"encoding/json"
	"net/http"
)

type payload map[string]interface{}

func Auth() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c, err := r.Cookie("token")
			if err != nil {
				switch err {
				case http.ErrNoCookie:
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(401)
					json.NewEncoder(w).Encode(payload{"error": "Invalid token"})
					return
				default:
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(400)
					json.NewEncoder(w).Encode(payload{"error": "Bad token"})
					return
				}
			}

			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "token", c.Value)))
		})
	}
}
