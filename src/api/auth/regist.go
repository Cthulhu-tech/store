package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Cthulhu-tech/store/src/utils/database"
	"github.com/Cthulhu-tech/store/src/utils/password"
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

	hash, err := password.HashPassword(userInfo.Password)

	if err != nil {

		return c.JSON(500, &Message{message: "Server Error"})

	}

	db := database.GetDB()

	rows, err := db.Query("SELECT `sp_registation`(?, ?, ?) AS `sp_registation`", userInfo.Login, hash, userInfo.Email)

	if err != nil {

		return c.JSON(500, &Message{message: "Server Error"})

	}

	value := &RegistrationFunc{}

	for rows.Next() {

		var registration RegistrationFunc

		if err := rows.Scan(&registration.Value); err != nil {

			log.Println(err.Error())

		}

		value.Value = registration.Value

	}

	smtpEmail(userInfo.Email, "confirm", c)

	if value.Value == 0 {

		return c.JSON(http.StatusOK, &Message{message: "Username or email already in use"})

	}

	return c.JSON(http.StatusOK, &Message{message: "User registered successfully"})

}
