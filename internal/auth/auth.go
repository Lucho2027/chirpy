package auth

import (
	"log"

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

func CheckPasswordHash(password, hash string) error{
	var err = bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password) )
	if err != nil {
		log.Printf("Wrong Password: %s", err)
		return err 
	}	
	return nil
}