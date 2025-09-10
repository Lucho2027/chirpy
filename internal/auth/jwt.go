package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {
	if tokenSecret == "" {
		return "", fmt.Errorf("token secret cannot be empty")
	}
	if expiresIn <= 0 {
		return "", fmt.Errorf("expiration time must be positive")
	}
	if userID == uuid.Nil {
		return "", fmt.Errorf("user ID cannot be nil")
	}
	currentTime := time.Now()
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(currentTime.Add(expiresIn)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    "chirpy",
		Subject:   userID.String(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(tokenSecret))

	return ss, err
}

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return uuid.Nil, err
	}

	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok {
		uid, err := uuid.Parse(claims.Subject)
		if err != nil {
			return uuid.Nil, err
		}
		return uid, nil
	}
	return uuid.Nil, err
}
