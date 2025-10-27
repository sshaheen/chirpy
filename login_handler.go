package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sshaheen/chirpy/internal/auth"
)

func (c *apiConfig) loginHandler(w http.ResponseWriter, r *http.Request) {
	type errorResponse struct {
		Error string `json:"error"`
	}

	type loginData struct {
		Password         string `json:"password"`
		Email            string `json:"email"`
		ExpiresInSeconds int    `json:"expires_in_seconds"`
	}

	userData := loginData{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userData)
	if err != nil {
		errResp := errorResponse{
			Error: "Something went wrong",
		}
		writeJSON(w, 500, errResp)
		return
	}

	if userData.ExpiresInSeconds == 0 || userData.ExpiresInSeconds > 3600 {
		userData.ExpiresInSeconds = 3600
	}

	user, err := c.dbQueries.GetUserByEmail(r.Context(), userData.Email)
	if err != nil {
		errResp := errorResponse{
			Error: "Failed to get user",
		}
		writeJSON(w, 500, errResp)
		return
	}

	token, err := auth.MakeJWT(user.ID, c.secret, time.Duration(userData.ExpiresInSeconds)*time.Second)
	if err != nil {
		errResp := errorResponse{
			Error: "Error making JWT",
		}
		writeJSON(w, 500, errResp)
		return
	}

	valid, err := auth.CheckPasswordHash(userData.Password, user.HashedPassword)
	if err != nil {
		errResp := errorResponse{
			Error: "Error checking password",
		}
		writeJSON(w, 500, errResp)
		return
	}

	if !valid {
		errResp := errorResponse{
			Error: "Invalid password",
		}
		writeJSON(w, 401, errResp)
		return
	}

	resData := struct {
		ID        uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Email     string    `json:"email"`
		Token     string    `json:"token"`
	}{
		user.ID,
		user.CreatedAt,
		user.UpdatedAt,
		user.Email,
		token,
	}

	writeJSON(w, http.StatusOK, resData)
}
