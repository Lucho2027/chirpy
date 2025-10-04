package auth

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	pswd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error encoding pswd : %s", err)
		return "", err
	}
	return string(pswd), nil
}

func CheckPasswordHash(password, hash string) error {
	var err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Printf("Wrong Password: %s", err)
		return err
	}
	return nil
}

func MakeRefreshToken() (string, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return "", fmt.Errorf("refresh token - problem creating")
	}
	refreshed := hex.EncodeToString(key)
	return refreshed, nil
}

func GetAuthFromHeader(headers http.Header, authPrefix string) (string, error) {
	infoToGet := headers.Get("Authorization")
	if infoToGet == "" {
		return "", fmt.Errorf("authorization header is missing")
	}
	prefixToLookFor := authPrefix + " "
	if !strings.HasPrefix(infoToGet, prefixToLookFor) {
		fmt.Printf("authorization header must start with %s", prefixToLookFor)
		return "", fmt.Errorf("authorization header must start with %s", prefixToLookFor)
	}
	token := strings.TrimPrefix(infoToGet, prefixToLookFor)
	if token == "" {
		return "", fmt.Errorf("get value from header returned value is empty")
	}

	return token, nil

}
