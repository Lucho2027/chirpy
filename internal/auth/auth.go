package auth

import "github.com/x/crypto/bcrypt"


func HashPassword(password string){
	bcrypt.GenerateFromPassword(password)
}