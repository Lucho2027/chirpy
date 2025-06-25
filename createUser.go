package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)
type User struct {
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email string `json:"email"`
}

type paramsUser struct {
	Email string `json:"email"`
}


func (cfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request ) {
	decoder := json.NewDecoder(r.Body);
	params:= paramsUser{}
	err := decoder.Decode(&params)
	if err != nil{
		log.Printf("Error decoding parameters: %s", err)
	}
	
	user, err := cfg.database.CreateUser(r.Context(), params.Email)
	if err != nil{
		log.Printf("Error saving user on db  %s:", err)
		respondWithError(w, 500, "Not able to create user")
		return
	}	
	respBody := User{
		ID: user.ID,
		CreatedAt: user.CreatedAt.Time,
		UpdatedAt: user.UpdatedAt.Time,
		Email: user.Email,
	}
	resp, err := json.Marshal(respBody)
	if err != nil {
		log.Printf("Error marshaling resp handleCreateUser : %s", err)
		respondWithError(w, 500, "Not able to marshal json create user")
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
	 
}