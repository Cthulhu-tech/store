package jwt 

import (
	"errors"

	"github.com/labstack/echo/v4"
)

func GetAccess(c echo.Context) (string, error) {

	bearer := c.Request().Header.Get("Authorization")

	if(len(bearer) <= 0) {

		return "", errors.New("Access token is required")

	}

	return bearer, nil

}
