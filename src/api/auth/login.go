package auth

import (
	"encoding/json"
	"os"

	"github.com/Cthulhu-tech/store/src/utils/database"
	"github.com/Cthulhu-tech/store/src/utils/jwt"
	"github.com/Cthulhu-tech/store/src/utils/message"
	"github.com/labstack/echo/v4"

	_ "github.com/go-sql-driver/mysql"
)

func Login(c echo.Context) error {

	refreshDuration := 10080

	user := UserLogin{}

	err := json.NewDecoder(c.Request().Body).Decode(&user)

	if err != nil {

		return message.JSON(c, 403, "Need all parameters")

	}

	if user.Password == "" || user.Login == "" || !isEmailValid(user.Mail) {

		return message.JSON(c, 403, "Need all parameters")

	}

	db := database.GetDB()

	rows, err := db.Query("SELECT * FROM users WHERE email = ? AND login = ? AND confirm = 1", user.Mail, user.Login)

	if err != nil {

		return message.JSON(c, 403, "User not found")

	}

	userInfo := UserAllData{}

	for rows.Next() {

		userCheck := UserAllData{}

		if err := rows.Scan(&userCheck.Id, &userCheck.Login, &userCheck.Email, &userCheck.Password, &userCheck.Confirm); err != nil {

			return message.JSON(c, 403, "User not found")

		}

		userInfo.Id = userCheck.Id
		userInfo.Login = userCheck.Login
		userInfo.Email = userCheck.Email
		userInfo.Password = userCheck.Password
		userInfo.Confirm = userCheck.Confirm

	}

	refresh, err := jwt.CreateToken(userInfo.Login, userInfo.Id, refreshDuration, os.Getenv("JWT_SECRET_REFRESH"))

	if err != nil {

		return message.JSON(c, 500, "Server error")

	}

	access, err := jwt.CreateToken(userInfo.Login, userInfo.Id, 30, os.Getenv("JWT_SECRET_BEARER"))

	if err != nil {

		return message.JSON(c, 500, "Server error")

	}

	_, err = db.Query("INSERT INTO token (user_id, jwt) VALUES (?, ?)", userInfo.Id, refresh)

	if err != nil {

		return message.JSON(c, 500, "Server error")

	}

	jwt.SetCookie(c, refreshDuration, refresh)

	return c.JSON(201, &MessageToken{Message: "You login", Token: access})

}
