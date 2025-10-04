package api

import (
	"net/http"
	"time"

	"github.com/Lucho2027/chirpy/internal/auth"
)

type Token struct {
	Token string `json:"token"`
}

func (cfg *ApiConfig) RefreshToken(w http.ResponseWriter, r *http.Request) {
	rt, err := auth.GetAuthFromHeader(r.Header, "Bearer")
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, "invalid token")
		return
	}

	userID, err := cfg.Database.GetUserFromRefreshToken(r.Context(), rt)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, "invalid or expired token")
		return
	}
	access, err := auth.MakeJWT(userID, cfg.JWT_Secret, time.Hour)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "could not create access token")
		return
	}
	RespondWithJson(w, http.StatusOK, Token{Token: access})
}
func (cfg *ApiConfig) RevokeToken(w http.ResponseWriter, r *http.Request) {
	rt, err := auth.GetAuthFromHeader(r.Header, "Bearer")
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, "invalid token")
		return
	}
	if err := cfg.Database.RevokeToken(r.Context(), rt); err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
