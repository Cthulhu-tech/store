package auth

import (
	"fmt"

	"github.com/Cthulhu-tech/store/src/utils/database"
	"github.com/Cthulhu-tech/store/src/utils/random"
	smtpMail "github.com/Cthulhu-tech/store/src/utils/smtp"
	"github.com/labstack/echo/v4"
)

func smtpEmail(email string, method string, c echo.Context) error {

	secret := random.GetNumber(0, 9999)

	url := random.RandomUrl(32)

	db := database.GetDB()

	_, err := db.Query("CALL sp_email(?, ?, ?)", email, url, secret)

	if err != nil {

		fmt.Println(err.Error())

		return c.JSON(500, &Message{message: "Server Error"})

	}

	smtpMail.SendMail(secret, method+url, email)

	return nil

}
