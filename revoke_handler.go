package main

import (
	"net/http"

	"github.com/sshaheen/chirpy/internal/auth"
)

func (c *apiConfig) revokeHandler(w http.ResponseWriter, r *http.Request) {
	refreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		errResp := struct {
			Error string `json:"error"`
		}{
			Error: "Refresh token doesn't exist or is expired",
		}
		writeJSON(w, 401, errResp)
		return
	}
	c.dbQueries.RevokeRefreshToken(r.Context(), refreshToken)
	w.WriteHeader(204)
}
