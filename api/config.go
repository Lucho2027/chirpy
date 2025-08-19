package api

import (
	"sync/atomic"

	"github.com/Lucho2027/chirpy/internal/database"
)


type ApiConfig struct {
		FileserverHits atomic.Int32
		Database *database.Queries
		Platform string
		JWT_Secret string
		
}

func NewApiConfig(db *database.Queries, platform string, secret string) *ApiConfig{
	return &ApiConfig{
		Database: db ,
		Platform: platform,
		JWT_Secret: secret,
	}
}