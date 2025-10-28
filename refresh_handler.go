package main

import (
	"net/http"
	"time"

	"github.com/sshaheen/chirpy/internal/auth"
)

func (c *apiConfig) refreshHandler(w http.ResponseWriter, r *http.Request) {
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
