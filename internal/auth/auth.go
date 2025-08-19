package auth

import (
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

func GetBearerToken(headers http.Header)(string, error){

	bearerToken := headers.Get("Authorization")
	if(bearerToken == ""){
			return "", fmt.Errorf("authorization header is missing")
	}
	if !strings.HasPrefix(bearerToken, "Bearer "){
		return "", fmt.Errorf("authorization header must start with 'Bearer '")
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	if token == "" {
		return "", fmt.Errorf("bearer token is empty")
	}

	return token, nil
}
