package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/sshaheen/chirpy/internal/auth"
	"github.com/sshaheen/chirpy/internal/database"
)

func (c *apiConfig) deleteChirpHandler(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		errResp := struct {
			Error string `json:"error"`
		}{
			Error: "encountered problem with access token",
		}
		writeJSON(w, 401, errResp)
		return
	}

	authUserId, err := auth.ValidateJWT(token, c.secret)
	if err != nil {
		errResp := struct {
			Error string `json:"error"`
		}{
			Error: "auth token bad or non existent",
		}
		writeJSON(w, 401, errResp)
		return
	}

	chirpId, err := uuid.Parse(r.PathValue("chirpID"))
	if err != nil {
		errResp := struct {
			Error string `json:"error"`
		}{
			Error: "chouldn't parse chirp ID",
		}
		writeJSON(w, 400, errResp)
		return
	}

	chirp, err := c.dbQueries.GetChirpById(r.Context(), chirpId)
	if err != nil {
		errResp := struct {
			Error string `json:"error"`
		}{
			Error: "Chirp doesn't exist",
		}
		writeJSON(w, 404, errResp)
		return
	}

	if chirp.UserID.UUID != authUserId {
		errResp := struct {
			Error string `json:"error"`
		}{
			Error: "Chirp id doesn't match user id",
		}
		writeJSON(w, 403, errResp)
		return
	}

	params := database.DeleteChirpParams{
		ID:     chirpId,
		UserID: uuid.NullUUID{UUID: authUserId, Valid: true},
	}

	err = c.dbQueries.DeleteChirp(r.Context(), params)
	if err != nil {
		errResp := struct {
			Error string `json:"error"`
		}{
			Error: "encountered problem deleting chirp",
		}
		writeJSON(w, 500, errResp)
		return
	}

	w.WriteHeader(204)
}
