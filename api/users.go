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
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Email       string    `json:"email"`
	Token       string    `json:"token"`
	RefresToken string    `json:"refresh_token"`
}

type paramsUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (cfg *ApiConfig) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := paramsUser{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		return
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
	if err != nil {
		log.Printf("Error saving user on db  %s:", err)
		RespondWithError(w, http.StatusInternalServerError, "Not able to create user")
		return
	}

	respBody := User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt.Time,
		UpdatedAt: user.UpdatedAt.Time,
		Email:     user.Email,
	}
	RespondWithJson(w, http.StatusCreated, respBody)
}

func (cfg *ApiConfig) HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := paramsUser{}
	err := decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, "Not able to update user")
		return
	}
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, "Not able to update users")
		return
	}
	userID, err := auth.ValidateJWT(token, cfg.JWT_Secret)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, "Not able to update ")
		return
	}
	hashedPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		log.Printf("Error hashing password  %s:", err)
		RespondWithError(w, http.StatusInternalServerError, "Not able to update user")
		return
	}
	user ,err := cfg.Database.UpdateUser(r.Context(), database.UpdateUserParams{
		Email:    params.Email,
		Password: hashedPassword,
		ID:       userID,
	})
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Not able to update user")
		return
	}
	respBody := User{
		ID: user.ID,
		CreatedAt: user.CreatedAt.Time,
		UpdatedAt: user.UpdatedAt.Time,
		Email: params.Email,
	}
	
	RespondWithJson(w, http.StatusOK, respBody)
}
