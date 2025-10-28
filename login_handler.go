package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sshaheen/chirpy/internal/auth"
	"github.com/sshaheen/chirpy/internal/database"
)

func (c *apiConfig) loginHandler(w http.ResponseWriter, r *http.Request) {
	type errorResponse struct {
		Error string `json:"error"`
	}

	type loginData struct {
		Password string `json:"password"`
		Email    string `json:"email"`
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

	user, err := c.dbQueries.GetUserByEmail(r.Context(), userData.Email)
	if err != nil {
		errResp := errorResponse{
			Error: "Failed to get user",
		}
		writeJSON(w, 500, errResp)
		return
	}

	accessToken, err := auth.MakeJWT(user.ID, c.secret, time.Duration(time.Hour))
	if err != nil {
		errResp := errorResponse{
			Error: "Error making JWT",
		}
		writeJSON(w, 500, errResp)
		return
	}

	refreshToken, err := auth.MakeRefreshToken()
	if err != nil {
		errResp := errorResponse{
			Error: "error with refresh token",
		}
		writeJSON(w, 500, errResp)
		return
	}

	refTokParams := database.CreateRefreshTokenParams{
		Token:     refreshToken,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(60 * 24 * time.Hour),
		RevokedAt: sql.NullTime{},
	}

	_, err = c.dbQueries.CreateRefreshToken(r.Context(), refTokParams)
	if err != nil {
		errResp := errorResponse{
			Error: "error creating refresh token in db",
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
		ID           uuid.UUID `json:"id"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
		Email        string    `json:"email"`
		AccessToken  string    `json:"token"`
		RefreshToken string    `json:"refresh_token"`
	}{
		user.ID,
		user.CreatedAt,
		user.UpdatedAt,
		user.Email,
		accessToken,
		refreshToken,
	}

	writeJSON(w, http.StatusOK, resData)
}
