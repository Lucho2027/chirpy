package auth

import (
	"time"

	"github.com/google/uuid"
	"github.com/golang-jwt/jwt/v5"
)

func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {
	signingKey := []byte("2342")
	currentTime := time.Now()

	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(currentTime.Add(expiresIn)),
		IssuedAt: jwt.NewNumericDate(time.Now()),
		Issuer:    "chirpy",
		Subject: string(userID),
	}

	return "we", nil
}
