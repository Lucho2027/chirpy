package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Lucho2027/chirpy/internal/auth"
)

func (cfg *ApiConfig) HandleLogin(w http.ResponseWriter, r *http.Request ) {
	decoder := json.NewDecoder(r.Body);
	params:= paramsUser{}
	err := decoder.Decode(&params)
	if err != nil{
		log.Printf("Error decoding parameters: %s", err)
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
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
	 
}