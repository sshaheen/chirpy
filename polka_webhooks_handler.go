package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/sshaheen/chirpy/internal/auth"
)

func (c *apiConfig) polkaWebhooksHandler(w http.ResponseWriter, r *http.Request) {
	type userData struct {
		UserId string `json:"user_id"`
	}

	type reqData struct {
		Event string   `json:"event"`
		Data  userData `json:"data"`
	}

	type errorResponse struct {
		Error string `json:"error"`
	}

	_, err := auth.GetAPIKey(r.Header)
	if err != nil {
		w.WriteHeader(401)
		return
	}

	var requestData reqData

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&requestData)
	if err != nil {
		errResp := errorResponse{
			Error: "error decoding request",
		}
		writeJSON(w, 500, errResp)
		return
	}

	if requestData.Event != "user.upgraded" {
		w.WriteHeader(204)
		return
	}

	userId, err := uuid.Parse(requestData.Data.UserId)
	if err != nil {
		errResp := errorResponse{
			Error: "error parsing user id",
		}
		writeJSON(w, 500, errResp)
		return
	}

	_, err = c.dbQueries.UpgradeUser(r.Context(), userId)
	if err != nil {
		errResp := errorResponse{
			Error: "user id does not exist",
		}
		writeJSON(w, 404, errResp)
		return
	}

	w.WriteHeader(204)
}
