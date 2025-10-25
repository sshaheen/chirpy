package main

import (
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(v)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"error":"internal server error"}`))
		return
	}
	w.WriteHeader(status)
	w.Write(dat)
}
