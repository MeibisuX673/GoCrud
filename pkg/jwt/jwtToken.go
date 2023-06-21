package jwt

import (

	"os"
	"time"
	"github.com/golang-jwt/jwt"

)


func CreateJWToken() (string, error){

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour).Unix()

	tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil{

		return "", err

	}

	return tokenStr, nil

}