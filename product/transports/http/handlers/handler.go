package handlers

import (
	"encoding/json"
	"net/http"
)

type payload map[string]interface{}

func JSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
