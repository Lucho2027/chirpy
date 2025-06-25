package main

import (
	"log"
	"net/http"
)

func (cfg *apiConfig) handleReset(w http.ResponseWriter, r *http.Request){
	if cfg.platform != "dev"{
		respondWithError(w, http.StatusForbidden, "Forbidden")
		return 
	}

	err := cfg.database.RemoveAllUsers(r.Context())
	if err != nil {
		log.Printf("Error deleting users %s", err)
		respondWithError(w, http.StatusInternalServerError, "Error deleting users")
		return
	}
	
	w.WriteHeader(http.StatusOK)
	cfg.fileserverHits.Swap(0)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}