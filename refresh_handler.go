package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/sshaheen/chirpy/internal/auth"
)

func (c *apiConfig) refreshHandler(w http.ResponseWriter, r *http.Request) {
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
	dbTokData, err := c.dbQueries.VerifyRefreshToken(r.Context(), refreshToken)
	if err != nil {
		errResp := struct {
			Error string `json:"error"`
		}{
			Error: "Refresh token doesn't exist or is expired",
		}
		writeJSON(w, 401, errResp)
		return
	}
	newAccessToken, err := auth.MakeJWT(dbTokData.UserID, c.secret, time.Duration(time.Hour))
	if err != nil {
		errResp := struct {
			Error string `json:"error"`
		}{
			Error: "couldn't generate new access token",
		}
		writeJSON(w, 401, errResp)
		return
	}
	respJSON := struct {
		Token string `json:"token"`
	}{
		Token: newAccessToken,
	}
	writeJSON(w, http.StatusOK, respJSON)
}
