package json

import (
	"encoding/json"
	"net/http"
)

func Write(w http.ResponseWriter , status int, data any) error {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}