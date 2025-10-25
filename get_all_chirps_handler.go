package main

import (
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (c *apiConfig) getAllChirpsHandler(w http.ResponseWriter, r *http.Request) {
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

	chirps, err := c.dbQueries.GetAllChirps(r.Context())
	if err != nil {
		errResp := errorResponse{
			Error: "Something went wrong",
		}
		writeJSON(w, 500, errResp)
		return
	}

	responses := []cleanedResponse{}
	for _, chirp := range chirps {
		res := cleanedResponse{
			chirp.ID,
			chirp.CreatedAt,
			chirp.UpdatedAt,
			chirp.Body,
			chirp.UserID,
		}
		responses = append(responses, res)
	}

	writeJSON(w, 200, responses)
}
