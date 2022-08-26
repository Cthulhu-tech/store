package jwt 

import (

	"github.com/labstack/echo/v4"

)

func GetCookie(c echo.Context) (string, error) {

	cookie, err := c.Cookie("refresh")

	if err != nil {

		return "", err

	}

	return cookie.Value, nil

}
