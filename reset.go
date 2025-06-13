package main

import "net/http"

func (cfg *apiConfig) handleReset(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	cfg.fileserverHits.Swap(0)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}