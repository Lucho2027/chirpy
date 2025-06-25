package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"sync/atomic"

	"github.com/Lucho2027/chirpy/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
		fileserverHits atomic.Int32
		database *database.Queries
		platform string
}


func main() {
    err := godotenv.Load()
	if err != nil {
		log.Printf("Error Loading env file %s\n", err)
	}
	dbURL := os.Getenv("DB_URL")
	env := os.Getenv("PLATFORM")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Printf("Problem initializing db: %s", err)
	}
	const port = "8080"
	const filepathRoot="."


	dbQueries := database.New(db)
	
	apiCfg := apiConfig{
		fileserverHits: atomic.Int32{},
		database: dbQueries,
		platform: env,
	}

	mux := http.NewServeMux()
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot)))))
  	mux.HandleFunc("GET /api/healthz",  handleReadiness)
	mux.HandleFunc("POST /api/users", apiCfg.handleCreateUser )
	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerMetrics)
	mux.HandleFunc("POST /admin/reset", apiCfg.handleReset)
	mux.HandleFunc("POST /api/validate_chirp", validateChirp )	


	srv := &http.Server{
		Addr: ":" + port,
		Handler: mux,
	}
	
	log.Printf("Serving files on port: %s\n",  port)
	log.Fatal(srv.ListenAndServe())
}


func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`<html>
  <body>
    <h1>Welcome, Chirpy Admin</h1>
    <p>Chirpy has been visited %d times!</p>
  </body>
</html>`, cfg.fileserverHits.Load())))
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}