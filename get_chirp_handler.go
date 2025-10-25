package main

import (
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (c *apiConfig) getChirpHandler(w http.ResponseWriter, r *http.Request) {
	type errorResponse struct {
		Error string `json:"error"`
	}
	type cleanedResponse struct {
		ID        uuid.UUID     `json:"id"`
		CreatedAt time.Time     `json:"created_at"`
		UpdatedAt time.Time     `json:"updated_at"`
		Body      string        `json:"body"`
		UserID    uuid.NullUUID `json:"user_id"`
	}
	chirpId, err := uuid.Parse(r.PathValue("chirpID"))
	if err != nil {
		errResp := errorResponse{
			Error: "Couldn't parse ChirpID",
		}
		writeJSON(w, 500, errResp)
		return
	}

	chirp, err := c.dbQueries.GetChirpById(r.Context(), chirpId)
	if err != nil {
		errResp := errorResponse{
			Error: "That Chirp does not exist",
		}
		writeJSON(w, 404, errResp)
		return
	}
	res := cleanedResponse{
		chirp.ID,
		chirp.CreatedAt,
		chirp.UpdatedAt,
		chirp.Body,
		chirp.UserID,
	}

	writeJSON(w, 200, res)
}
