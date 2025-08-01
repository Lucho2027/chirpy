package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Lucho2027/chirpy/api"
	"github.com/Lucho2027/chirpy/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)




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

	dbQueries := database.New(db)
	apiCfg := api.NewApiConfig(dbQueries, env)

	mux := http.NewServeMux()
	api.RegisterRoutes(mux, apiCfg)	

	srv := &http.Server{
		Addr: ":" + port,
		Handler: mux,
	}
	
	log.Printf("Serving files on port: %s\n",  port)
	log.Fatal(srv.ListenAndServe())
}

