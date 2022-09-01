package auth

import (
	"fmt"
	"os"

	"github.com/Cthulhu-tech/store/src/utils/database"
	"github.com/Cthulhu-tech/store/src/utils/jwt"
	"github.com/Cthulhu-tech/store/src/utils/message"
	"github.com/labstack/echo/v4"
)

func Refresh(c echo.Context) error {

	refreshDuration := 10080

	cookie, err := jwt.GetCookie(c)

	if err != nil {

		return message.JSON(c, 403, "Need refresh cookie")

	}

	value := jwt.CheckToken(cookie)

	if value.Id == 0 || value.User == "" {

		return message.JSON(c, 403, "Need refresh cookie")

	}
	fmt.Println(value.Id, value.User)
	db := database.GetDB()

	rows, err := db.Query("SELECT * FROM users WHERE id = ? AND login = ? AND confirm = 1", value.Id, value.User)

	defer rows.Close()

	if err != nil {

		return message.JSON(c, 500, "Server error")

	}

	userInfo := UserAllData{}

	for rows.Next() {

		userCheck := UserAllData{}

		if err := rows.Scan(&userCheck.Id, &userCheck.Login, &userCheck.Email, &userCheck.Password, &userCheck.Confirm); err != nil {

			return message.JSON(c, 500, "Server error")

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

	rows, err = db.Query("INSERT INTO token (user_id, jwt) VALUES (?, ?)", userInfo.Id, refresh)

	defer rows.Close()

	if err != nil {

		return message.JSON(c, 500, "Server error")

	}

	jwt.SetCookie(c, refreshDuration, refresh)

	return c.JSON(201, &MessageToken{Message: "You login", Token: access})

}
