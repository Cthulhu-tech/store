package auth

import (
	"encoding/json"

	"github.com/Cthulhu-tech/store/src/utils/database"
	"github.com/Cthulhu-tech/store/src/utils/message"
	"github.com/Cthulhu-tech/store/src/utils/password"
	"github.com/labstack/echo/v4"
)

func Regist(c echo.Context) error {

	var userInfo = &UserRegistration{}

	err := json.NewDecoder(c.Request().Body).Decode(&userInfo)

	if err != nil {

		return err

	}

	if userInfo.Login == "" || userInfo.Password == "" || !isEmailValid(userInfo.Email) {

		return message.JSON(c, 403, "Need all parameters")

	}

	hash, err := password.HashPassword(userInfo.Password)

	if err != nil {

		return message.JSON(c, 500, "Server error")

	}
	db := database.GetDB()

	rows, err := db.Query("SELECT `sp_registation`(?, ?, ?) AS `sp_registation`", userInfo.Login, hash, userInfo.Email)

	if err != nil {

		return message.JSON(c, 500, "Server error")

	}

	value := &RegistrationFunc{}

	for rows.Next() {

		var registration RegistrationFunc

		if err := rows.Scan(&registration.Value); err != nil {

			return message.JSON(c, 500, "Server error")

		}

		value.Value = registration.Value

	}

	if value.Value == 0 {

		return message.JSON(c, 403, "Username or email already in use")

	}

	defer func() {

		go smtpEmail(userInfo.Email, "confirm", c)

	}()

	return message.JSON(c, 200, "User registered successfully")

}
