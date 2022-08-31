package auth

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"

	_ "github.com/go-sql-driver/mysql"
)

func Login(c echo.Context) error {

	json_map := make(map[string]interface{})

	err := json.NewDecoder(c.Request().Body).Decode(&json_map)

	if err != nil {

		return err

	}

	return c.String(http.StatusOK, "login")

}
