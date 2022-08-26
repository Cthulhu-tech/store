package jwt 

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(user string, id int, duration int, secret string) (string, error) {
	
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Duration(duration) * time.Minute)

	claims["user"] = user

	claims["id"] = id

	tokenStr, err := token.SignedString([]byte(secret))

	if err != nil {

		fmt.Println(err.Error())

		return "", err


	}

	return tokenStr, nil

}
