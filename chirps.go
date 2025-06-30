package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

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
	Message string    `json:"body"`
	UserID  uuid.UUID `json:"user_id"`
}

func (cfg *apiConfig) handleCreateChirp(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := paramsChirp{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("Error decoding parameters handleCreateChirp : %s", err)
	}
	if !validateChirp(params.Message) {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long!")
		return
	}
	chirp, err := cfg.database.CreateChirp(r.Context(), database.CreateChirpParams{
		Message: params.Message,
		UserID:  params.UserID,
	})
	if err != nil {
		log.Printf("Error saving msg on db  %s:", err)
		respondWithError(w, http.StatusInternalServerError, "Not able to create user")
		return
	}
	respBody := Chirp{
		ID:        chirp.ID,
		Message:   chirp.Message,
		UserID:    chirp.UserID,
		CreatedAt: chirp.CreatedAt.Time,
		UpdatedAt: chirp.UpdatedAt.Time,
	}
	resp, err := json.Marshal(respBody)
	if err != nil {
		log.Printf("Error marshaling resp handleCreateChirp : %s", err)
		respondWithError(w, 500, "Not able to marshal json create chirp")
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}

func (cfg *apiConfig) handleGetAll(w http.ResponseWriter, r *http.Request) {

	chirps, err := cfg.database.GetAllChirps(r.Context())
	if err != nil {
		log.Printf("Error getting all chirps: %s", err)
		respondWithError(w, http.StatusInternalServerError, "Not able to retrieve chirps")
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
	resp, err := json.Marshal(respBody)
	if err != nil {
		log.Printf("Error marshaling resp handleGetAllChirps : %s", err)
		respondWithError(w, 500, "Not able to marshal json get chirp")
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
func (cfg *apiConfig) handleGetChirpById(w http.ResponseWriter, r *http.Request){
	chirpId := r.PathValue("chirpID")
	log.Printf("HERE IS THE PATH VALUE %s", chirpId)
	parsedChirpId, err := uuid.Parse(chirpId)
	if err != nil {
		log.Printf("Error parsing uuid: %s", err)
		respondWithError(w, http.StatusInternalServerError, "Error parsing ChirpId")
		return
	}
	log.Printf("here is uuid obj %v", parsedChirpId)
	cDb, err := cfg.database.GetChirpById(r.Context(), parsedChirpId)
	if err != nil {
		log.Printf("Error getting chirp from db %s", err)
		respondWithError(w, http.StatusInternalServerError, "Error getting Chirp from db")
		return
	}
	
	respBody := Chirp{
		ID: cDb.ID,
		Message: cDb.Message,
		UserID: cDb.UserID,
		CreatedAt: cDb.CreatedAt.Time,
		UpdatedAt: cDb.UpdatedAt.Time,
	}
	resp, err := json.Marshal(respBody)
	if err != nil {
		log.Printf("Error marshaling resp handleCreateChirp : %s", err)
		respondWithError(w, 500, "Not able to marshal json create chirp")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}