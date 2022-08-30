package auth

import (

	"net/http"
	"github.com/labstack/echo/v4"
)

func Vkontakte(c echo.Context) error {

	code := c.QueryParam("code")
	_, err := http.Get("" + code)

	if err != nil {
		
		return err

	}

	return nil

}
