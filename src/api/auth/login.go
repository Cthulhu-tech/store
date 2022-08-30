package auth

import (
	"net/http"
	
	"github.com/labstack/echo/v4"

	_ "github.com/go-sql-driver/mysql"
)

func Login(c echo.Context) error {

	return c.String(http.StatusOK, "login")

}
