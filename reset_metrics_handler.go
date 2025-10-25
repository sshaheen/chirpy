package main

import (
	"fmt"
	"net/http"
)

func (c *apiConfig) ResetMetricsHandler(w http.ResponseWriter, r *http.Request) {
	if c.platform != "dev" {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("You cannot erase users"))
		return
	}
	c.fileserverHits.Store(0)
	c.dbQueries.DeleteAllUsers(r.Context())
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hits: %d", c.fileserverHits.Load())))
}
