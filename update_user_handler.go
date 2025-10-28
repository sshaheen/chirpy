package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sshaheen/chirpy/internal/auth"
	"github.com/sshaheen/chirpy/internal/database"
)

func (c *apiConfig) updateUserHandler(w http.ResponseWriter, r *http.Request) {
	type reqData struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	accessToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		errResp := struct {
			Error string `json:"error"`
		}{
			Error: "encpuntered problem with access token",
		}
		writeJSON(w, 401, errResp)
		return
	}

	userId, err := auth.ValidateJWT(accessToken, c.secret)
	if err != nil {
		errResp := struct {
			Error string `json:"error"`
		}{
			Error: "invalid jwt",
		}
		writeJSON(w, 401, errResp)
		return
	}

	req := reqData{}

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&req)

	hashedPass, err := auth.HashPassword(req.Password)
	if err != nil {
		errResp := struct {
			Error string `json:"error"`
		}{
			Error: "couldn't hash password",
		}
		writeJSON(w, 500, errResp)
		return
	}

	params := database.UpdateUserParams{
		ID:             userId,
		HashedPassword: hashedPass,
		Email:          req.Email,
	}

	updatedUser, err := c.dbQueries.UpdateUser(r.Context(), params)
	if err != nil {
		errResp := struct {
			Error string `json:"error"`
		}{
			Error: "couldn't update user",
		}
		writeJSON(w, 500, errResp)
		return
	}

	type resultData struct {
		ID          uuid.UUID `json:"id"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		Email       string    `json:"email"`
		IsChirpyRed bool      `json:"is_chirpy_red"`
	}

	result := resultData{
		ID:          updatedUser.ID,
		CreatedAt:   updatedUser.CreatedAt,
		UpdatedAt:   updatedUser.UpdatedAt,
		Email:       updatedUser.Email,
		IsChirpyRed: updatedUser.IsChirpyRed,
	}

	writeJSON(w, http.StatusOK, result)
}
