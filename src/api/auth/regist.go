package auth

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Regist(c echo.Context) error {

	var userInfo = &UserRegistration{}

	err := json.NewDecoder(c.Request().Body).Decode(&userInfo)

	if err != nil {

		return err

	}

	if userInfo.Login == "" || userInfo.Password == "" || userInfo.Email == "" {

		return c.JSON(403, &Message{message: "Invalid body parameters"})

	}

	return c.JSON(http.StatusOK, &Message{message: "User registered successfully"})

}
