package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Lucho2027/chirpy/internal/auth"
	"github.com/Lucho2027/chirpy/internal/database"
)

func (cfg *ApiConfig) HandleLogin(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	params := paramsUser{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
	}

	user, err := cfg.Database.GetByEmail(r.Context(), params.Email)
	if err != nil {
		log.Printf("Error getting user by email %s:", err)
		RespondWithError(w, http.StatusInternalServerError, "Not able to find user")
		return
	}

	if auth.CheckPasswordHash(params.Password, user.HashedPassword) != nil {
		RespondWithError(w, http.StatusUnauthorized, "")
		return
	}

	token, err := auth.MakeJWT(user.ID, cfg.JWT_Secret, time.Duration(time.Hour))
	if err != nil {
		log.Printf("Error creating JWT %s", err)
		RespondWithError(w, http.StatusInternalServerError, "Not able to auth user")
		return
	}
	tkn, err := auth.MakeRefreshToken()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Not able to create rfrsh token")
		return
	}
	var nullTime sql.NullTime
	nullTime.Time = time.Now().Add(time.Hour * 24 * 60)
	nullTime.Valid = true
	tokenParams := database.CreateRefreshTokenParams{
		Token:     tkn,
		UserID:    user.ID,
		ExpiresAt: nullTime,
	}
	refreshToken, err := cfg.Database.CreateRefreshToken(r.Context(), tokenParams)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "unable to create refresh token")
		return
	}
	respBody := User{
		ID:          user.ID,
		CreatedAt:   user.CreatedAt.Time,
		UpdatedAt:   user.UpdatedAt.Time,
		Email:       user.Email,
		Token:       token,
		RefresToken: refreshToken.Token,
		IsChirpyRed: user.IsChirpyRed,
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
