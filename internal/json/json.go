package json

import (
	"encoding/json"
	"net/http"
)

//Response
func Write(w http.ResponseWriter , status int, data any) error {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

//request
func Read(r *http.Request, data any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(data)
}