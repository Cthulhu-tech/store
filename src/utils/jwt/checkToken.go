package jwt

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

func CheckToken(tokenValue string) *TokenUserInfo {

	claims := jwt.MapClaims{}

	_token, _ := jwt.ParseWithClaims(tokenValue, claims, func(token *jwt.Token) (interface{}, error) {

		return []byte(os.Getenv("JWT_SECRET_REFRESH")), nil

	})

	if claims, ok := _token.Claims.(jwt.MapClaims); ok {

		id := claims["id"].(int)
		user := claims["user"].(string)

		return &TokenUserInfo{Id: id, User: user}

	}

	return &TokenUserInfo{Id: 0, User: ""}

}
