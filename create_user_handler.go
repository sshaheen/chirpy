package main

import (
	"encoding/json"
	"net/http"

	"github.com/sshaheen/chirpy/internal/auth"
	"github.com/sshaheen/chirpy/internal/database"
)

func (c *apiConfig) createUserHandler(w http.ResponseWriter, r *http.Request) {
	type errorResponse struct {
		Error string `json:"error"`
	}

	params := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		errResp := errorResponse{
			Error: "Something went wrong",
		}
		writeJSON(w, 500, errResp)
		return
	}

	params.Password, err = auth.HashPassword(params.Password)
	if err != nil {
		errResp := errorResponse{
			Error: "Could not hash user",
		}
		writeJSON(w, 500, errResp)
		return
	}

	user, err := c.dbQueries.CreateUser(r.Context(), database.CreateUserParams{Email: params.Email, HashedPassword: params.Password})
	if err != nil {
		errResp := errorResponse{
			Error: "Could not create user",
		}
		writeJSON(w, 500, errResp)
		return
	}

	mappedUser := User{
		ID:             user.ID,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
	}

	writeJSON(w, 201, mappedUser)
}
