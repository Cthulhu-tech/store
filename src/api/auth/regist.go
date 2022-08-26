package auth

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
)

func Regist(c echo.Context) error {

	return c.String(http.StatusOK, "regist")

}
