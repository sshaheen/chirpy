package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sshaheen/chirpy/internal/database"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	dbQueries      *database.Queries
	platform       string
	secret         string
}

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Println("Couldn't open DB")
		os.Exit(1)
	}
	platform := os.Getenv("PLATFORM")
	secret := os.Getenv("SECRET")
	dbQueries := database.New(db)
	apiCfg := &apiConfig{}
	apiCfg.dbQueries = dbQueries
	apiCfg.platform = platform
	apiCfg.secret = secret
	port := "8080"
	filepathRoot := "."
	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	prefix := "/app"
	mux.Handle("/app", apiCfg.middlewareMetricsInc(http.StripPrefix(prefix, http.FileServer(http.Dir(filepathRoot)))))
	mux.Handle("/app/assets/", apiCfg.middlewareMetricsInc(http.StripPrefix(prefix, http.FileServer(http.Dir(filepathRoot)))))
	mux.HandleFunc("GET /admin/healthz", OKHealth)
	mux.HandleFunc("GET /admin/metrics", apiCfg.MetricsHandler)
	mux.HandleFunc("POST /admin/reset", apiCfg.ResetMetricsHandler)
	mux.HandleFunc("POST /api/users", apiCfg.createUserHandler)
	mux.HandleFunc("POST /api/chirps", apiCfg.createChirpHandler)
	mux.HandleFunc("GET /api/chirps", apiCfg.getAllChirpsHandler)
	mux.HandleFunc("GET /api/chirps/{chirpID}", apiCfg.getChirpHandler)
	mux.HandleFunc("POST /api/login", apiCfg.loginHandler)
	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
