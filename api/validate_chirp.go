package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type returnVals struct {
	Valid bool   `json:"valid,omitempty"`
	Error string `json:"error,omitempty"`
}

func ValidateChirp(c string) bool {
	len := len(c)
	return len <= 140
}

func RespondWithJson(w http.ResponseWriter, code int, payload any) {
	resp, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshaling respBody inside: %s", err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(resp)

}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	respBody := returnVals{
		Error: msg,
	}
	RespondWithJson(w, code, respBody)

}
