package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Lucho2027/chirpy/internal/auth"
)

func (cfg *ApiConfig) HandleLogin(w http.ResponseWriter, r *http.Request ) {
	
	decoder := json.NewDecoder(r.Body);
	params:= paramsUser{}
	err := decoder.Decode(&params)
	if err != nil{
		log.Printf("Error decoding parameters: %s", err)
	} 
	
	expiresIn := time.Hour 
	if params.ExpiresInSeconds != nil {
		expiresIn = time.Duration(*params.ExpiresInSeconds) * time.Second
	}

	user, err := cfg.Database.GetByEmail(r.Context(), params.Email);
	if err != nil {
		log.Printf("Error getting user by email %s:", err)
		RespondWithError(w, http.StatusInternalServerError, "Not able to find user")
		return
	}
 
	if  auth.CheckPasswordHash(params.Password, user.HashedPassword ) != nil {
		RespondWithError(w, http.StatusUnauthorized, "")
		return
	}

	token, err := auth.MakeJWT(user.ID, cfg.JWT_Secret, time.Duration(expiresIn))
	if err != nil {
		 log.Printf("Error creating JWT %s", err)
		 RespondWithError(w, http.StatusInternalServerError, "Not able to auth user")
		 return
	}

	respBody := User{
		ID: user.ID,
		CreatedAt: user.CreatedAt.Time,
		UpdatedAt: user.UpdatedAt.Time,
		Email: user.Email,
		Token: token,
	}
	resp, err := json.Marshal(respBody)
	if err != nil {
		log.Printf("Error marshaling resp handleCreateUser : %s", err)
		RespondWithError(w, http.StatusInternalServerError, "Not able to marshal json create user")
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
	 
}