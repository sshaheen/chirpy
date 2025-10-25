package main

import (
	"encoding/json"
	"net/http"
)

func (c *apiConfig) createUserHandler(w http.ResponseWriter, r *http.Request) {
	type errorResponse struct {
		Error string `json:"error"`
	}

	params := struct {
		Email string `json:"email"`
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

	user, err := c.dbQueries.CreateUser(r.Context(), params.Email)
	if err != nil {
		errResp := errorResponse{
			Error: "Could not create user",
		}
		writeJSON(w, 500, errResp)
		return
	}

	mappedUser := User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	}

	writeJSON(w, 201, mappedUser)
}
