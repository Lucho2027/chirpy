package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type parameters struct {
	Body string `json:"body"`
}
type returnVals struct {
	Valid bool   `json:"valid,omitempty"`
	Error string `json:"error,omitempty"`
}

func validateChirp(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		return
	}

	len := len(params.Body)
	if len > 140 {
		log.Printf("Chirp is too long")
		respondWithError(w, http.StatusBadRequest, "Chirp is too long")
		return
	}
	respBody := returnVals{
		Valid: true,
	}
	respondWithJson(w, http.StatusOK, respBody)

}

func respondWithJson(w http.ResponseWriter, code int, payload returnVals) {
	resp, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshaling respBody inside: %s", err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(resp)

}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respBody := returnVals{
		Error: msg,
	}
	respondWithJson(w, code, respBody)

}
