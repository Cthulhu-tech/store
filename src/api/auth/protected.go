package auth

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
)

func Protected(c echo.Context) error {

	return c.String(http.StatusOK, "protected")

}
