package main

import (
	"net/http"
	"strings"
)

func (c *apiConfig) revokeHandler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	headerStrs := strings.Split(authHeader, " ")
	if len(headerStrs) < 2 {
		errResp := struct {
			Error string `json:"error"`
		}{
			Error: "No authorization headers were passed",
		}
		writeJSON(w, 401, errResp)
		return
	}
	refreshToken := strings.TrimSpace(headerStrs[1])
	c.dbQueries.RevokeRefreshToken(r.Context(), refreshToken)
	w.WriteHeader(204)
}
