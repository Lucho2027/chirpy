package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Lucho2027/chirpy/internal/database"
	"github.com/google/uuid"
)

type WebhookData struct {
	UserID uuid.UUID `json:"user_id"`
}
type Webhook struct {
	Event string      `json:"event"`
	Data  WebhookData `json:"data"`
}

func (cfg *ApiConfig) HasRedis() bool {
	return cfg.Redis != nil
}

func (cfg *ApiConfig) HandleWebhook(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := Webhook{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("Error decodign parameters : %s", err)
		return
	}
	if params.Event != "user.upgraded" {
		RespondWithJson(w, http.StatusNoContent, "")
		return
	}

	idempotencyKey := params.Event + ":" + params.Data.UserID.String()
	if cfg.HasRedis() {
		if _, err := cfg.Redis.Get(r.Context(), idempotencyKey).Result(); err == nil {
			RespondWithJson(w, http.StatusNoContent, "")
			return
		}
	}

	id, err := cfg.Database.UpgradeUser(r.Context(), database.UpgradeUserParams{
		IsChirpyRed: bool(true),
		ID:          params.Data.UserID,
	})

	if err != nil {
		RespondWithError(w, http.StatusNotFound, "")
		return
	}
	fmt.Printf("We have upgraded: %s \n", id)
	if cfg.HasRedis() {
		cfg.Redis.Set(r.Context(), idempotencyKey, "processed", 24*time.Hour)
	}

	RespondWithJson(w, http.StatusNoContent, "")

}
