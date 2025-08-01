package api

import (
	"log"
	"net/http"
)

func (cfg *ApiConfig) HandleReset(w http.ResponseWriter, r *http.Request){
	if cfg.Platform != "dev"{
		RespondWithError(w, http.StatusForbidden, "Forbidden")
		return 
	}

	err := cfg.Database.RemoveAllUsers(r.Context())
	if err != nil {
		log.Printf("Error deleting users %s", err)
		RespondWithError(w, http.StatusInternalServerError, "Error deleting users")
		return
	}
	
	w.WriteHeader(http.StatusOK)
	cfg.FileserverHits.Swap(0)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}