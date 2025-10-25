package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sshaheen/chirpy/internal/database"
)

func (c *apiConfig) createChirpHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body   string        `json:"body"`
		UserID uuid.NullUUID `json:"user_id"`
	}

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

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		errResp := errorResponse{
			Error: "Something went wrong",
		}
		writeJSON(w, 500, errResp)
		return
	}

	chirp := sanitizeChirp(params.Body)

	if len(chirp) > 140 {
		errResp := errorResponse{
			Error: "Chirp is too long",
		}
		writeJSON(w, 400, errResp)
		return
	}

	chirpParams := database.CreateChirpParams{
		Body:   chirp,
		UserID: params.UserID,
	}

	dbResp, err := c.dbQueries.CreateChirp(r.Context(), chirpParams)
	if err != nil {
		errResp := errorResponse{
			Error: "Could not create user",
		}
		writeJSON(w, 500, errResp)
		return
	}

	mappedChirp := cleanedResponse{
		ID:        dbResp.ID,
		CreatedAt: dbResp.CreatedAt,
		UpdatedAt: dbResp.UpdatedAt,
		Body:      dbResp.Body,
		UserID:    dbResp.UserID,
	}

	writeJSON(w, 201, mappedChirp)
}
