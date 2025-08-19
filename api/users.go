package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Lucho2027/chirpy/internal/auth"
	"github.com/Lucho2027/chirpy/internal/database"
	"github.com/google/uuid"
)
type User struct {
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type paramsUser struct {
	Email string `json:"email"`
	Password string `json:"password"`
	ExpiresInSeconds *int `json:"expires_in_seconds,omitempty"`
}


func (cfg *ApiConfig) HandleCreateUser(w http.ResponseWriter, r *http.Request ) {
	decoder := json.NewDecoder(r.Body);
	params:= paramsUser{}
	err := decoder.Decode(&params)
	if err != nil{
		log.Printf("Error decoding parameters: %s", err)
	}

	hashedPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		log.Printf("Error hashing password  %s:", err)
		RespondWithError(w, http.StatusInternalServerError, "Not able to create user")
		return
	}
	
	userParams := database.CreateUserParams{
		Email:    params.Email,
		Password: hashedPassword,
	}
	user, err := cfg.Database.CreateUser(r.Context(), userParams)
	if err != nil{
		log.Printf("Error saving user on db  %s:", err)
		RespondWithError(w, http.StatusInternalServerError, "Not able to create user")
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
		RespondWithError(w, http.StatusInternalServerError, "Not able to marshal json create user")
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
	 
}