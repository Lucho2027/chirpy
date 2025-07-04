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

func validateChirp(c string) bool {
	len := len(c)
	if len > 140 {
		return false
	}
    return true

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
