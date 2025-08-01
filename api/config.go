package api

import (
	"sync/atomic"

	"github.com/Lucho2027/chirpy/internal/database"
)


type ApiConfig struct {
		FileserverHits atomic.Int32
		Database *database.Queries
		Platform string
}

func NewApiConfig(db *database.Queries, platform string) *ApiConfig{
	return &ApiConfig{
		Database: db ,
		Platform: platform,
	}
}