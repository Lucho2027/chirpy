package api

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/Lucho2027/chirpy/internal/auth"
	"github.com/Lucho2027/chirpy/internal/database"
	"github.com/google/uuid"
)

type Chirp struct {
	ID        uuid.UUID `json:"id"`
	Message   string    `json:"body"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type paramsChirp struct {
	Message string `json:"body"`
}

func (cfg *ApiConfig) HandleCreateChirp(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := paramsChirp{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("Error decoding parameters handleCreateChirp : %s", err)
		RespondWithError(w, http.StatusInternalServerError, "Not able to create chirps")
		return
	}
	token, err := auth.GetAuthFromHeader(r.Header, "Bearer")
	if err != nil {
		log.Printf("Error saving chirp - token invalid %s", err)
		RespondWithError(w, http.StatusUnauthorized, "Not able to create chirps")
		return
	}
	userID, err := auth.ValidateJWT(token, cfg.JWT_Secret)
	if err != nil {
		log.Printf("Error saving chirp - token invalid %s", err)
		RespondWithError(w, http.StatusUnauthorized, "Not able to create chirps")
		return
	}

	if !ValidateChirp(params.Message) {
		RespondWithError(w, http.StatusBadRequest, "Chirp is too long!")
		return
	}
	chirp, err := cfg.Database.CreateChirp(r.Context(), database.CreateChirpParams{
		Message: params.Message,
		UserID:  userID,
	})
	if err != nil {
		log.Printf("Error saving chirp on db  %s:", err)
		RespondWithError(w, http.StatusInternalServerError, "Not able to create chirps")
		return
	}
	respBody := Chirp{
		ID:        chirp.ID,
		Message:   chirp.Message,
		UserID:    chirp.UserID,
		CreatedAt: chirp.CreatedAt.Time,
		UpdatedAt: chirp.UpdatedAt.Time,
	}

	RespondWithJson(w, http.StatusCreated, respBody)
}

func (cfg *ApiConfig) HandleGetAll(w http.ResponseWriter, r *http.Request) {
	authorID := r.URL.Query().Get("author_id")
	var (
		chirps []database.Chirp
		err    error
	)

	if authorID == "" {
		chirps, err = cfg.Database.GetAllChirps(r.Context())
	} else {
		parsedAuthorID, parseErr := uuid.Parse(authorID)
		if parseErr != nil {
			log.Printf("Error parsing uuid: %s", err)
			RespondWithError(w, http.StatusInternalServerError, "Error parsing AuthorId")
			return
		}
		chirps, err = cfg.Database.GetAllChirpsByAuthor(r.Context(), parsedAuthorID)
	}

	if err != nil {
		log.Printf("Error getting all chirps: %s", err)
		RespondWithError(w, http.StatusInternalServerError, "Not able to retrieve chirps")
		return
	}
	var respBody = []Chirp{}

	for _, c := range chirps {
		respBody = append(respBody, Chirp{
			ID:        c.ID,
			Message:   c.Message,
			UserID:    c.UserID,
			CreatedAt: c.CreatedAt.Time,
			UpdatedAt: c.UpdatedAt.Time,
		})
	}
	querySort := r.URL.Query().Get("sort")

	switch querySort {
	case "asc":
		sort.Slice(respBody, func(i, j int) bool { return respBody[i].CreatedAt.Before(respBody[j].CreatedAt) })
	case "desc":
		sort.Slice(respBody, func(i, j int) bool { return respBody[i].CreatedAt.After(respBody[j].CreatedAt) })
	}

	RespondWithJson(w, http.StatusOK, respBody)
}
func (cfg *ApiConfig) HandleGetChirpById(w http.ResponseWriter, r *http.Request) {
	chirpId := r.PathValue("chirpID")
	parsedChirpId, err := uuid.Parse(chirpId)
	if err != nil {
		log.Printf("Error parsing uuid: %s", err)
		RespondWithError(w, http.StatusInternalServerError, "Error parsing ChirpId")
		return
	}
	cDb, err := cfg.Database.GetChirpById(r.Context(), parsedChirpId)
	if err != nil {
		log.Printf("Error getting chirp from db %s", err)
		RespondWithError(w, http.StatusNotFound, "Error getting Chirp from db")
		return
	}

	respBody := Chirp{
		ID:        cDb.ID,
		Message:   cDb.Message,
		UserID:    cDb.UserID,
		CreatedAt: cDb.CreatedAt.Time,
		UpdatedAt: cDb.UpdatedAt.Time,
	}
	RespondWithJson(w, http.StatusOK, respBody)
}
func (cfg *ApiConfig) HandleDeleteChirpById(w http.ResponseWriter, r *http.Request) {
	chirpId := r.PathValue("chirpID")
	parsedChirpId, err := uuid.Parse(chirpId)
	if err != nil {
		log.Printf("Error parsing uuid: %s", err)
		RespondWithError(w, http.StatusInternalServerError, "Error parsing ChirpId")
		return
	}
	token, err := auth.GetAuthFromHeader(r.Header, "Bearer")
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, "Not authorized")
		return
	}
	userID, err := auth.ValidateJWT(token, cfg.JWT_Secret)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, "Not authorized")
		return
	}
	chirpToDel, err := cfg.Database.GetChirpById(r.Context(), parsedChirpId)
	if err != nil {
		RespondWithError(w, http.StatusNotFound, "Error getting Chirp from db")
		return
	}
	if chirpToDel.UserID != userID {
		RespondWithError(w, http.StatusForbidden, "Not authorized")
		return
	}
	if err := cfg.Database.DeleteChirpById(r.Context(), database.DeleteChirpByIdParams{
		ChirpID: parsedChirpId,
		UserID:  userID,
	}); err != nil {
		RespondWithError(w, http.StatusNotFound, "Not able to delete chirp")
		return
	}
	RespondWithJson(w, http.StatusNoContent, "")
}
