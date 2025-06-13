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
		respBody := returnVals{
			Error: "Chirp is too long",
		}
		dat, err := json.Marshal(respBody)
		if err != nil {
			log.Printf("Error marshaling respBody inside  len constraint: %s", err)
			return
		}
		w.WriteHeader(http.StatusBadRequest)

		w.Write(dat)
		return
	}
	respBody := returnVals{
		Valid: true,
	}
	dat, err := json.Marshal(respBody)
	if err != nil {
		log.Printf("Error marshaling resp: %s", err)
		respBody := returnVals{
			Error: "Something went wrong",
		}
		dat, err := json.Marshal(respBody)
		if err != nil {
			log.Printf("Error marshaling respBody inside  len constraint: %s", err)
			return
		}
		w.WriteHeader(http.StatusBadRequest)

		w.Write(dat)

		return

	}
	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	w.Write(dat)
}
