package api

import (
	"sync/atomic"

	"github.com/Lucho2027/chirpy/internal/database"
	"github.com/redis/go-redis/v9"
)


type ApiConfig struct {
		FileserverHits atomic.Int32
		Database *database.Queries
		Platform string
		JWT_Secret string
		Redis *redis.Client
		
}

func NewApiConfig(db *database.Queries, platform string, secret string, redisClient *redis.Client) *ApiConfig{
	return &ApiConfig{
		Database: db ,
		Platform: platform,
		JWT_Secret: secret,
		Redis: redisClient,
	}
}