package auth

import (
	"fmt"
	"regexp"

	"github.com/Cthulhu-tech/store/src/utils/database"
	"github.com/Cthulhu-tech/store/src/utils/message"
	"github.com/Cthulhu-tech/store/src/utils/random"
	smtpMail "github.com/Cthulhu-tech/store/src/utils/smtp"
	"github.com/labstack/echo/v4"
)

func smtpEmail(email string, method string, c echo.Context) error {

	url := random.RandomUrl(32)

	secret := random.GetNumber(0, 9999)

	db := database.GetDB()

	_, err := db.Query("CALL sp_email(?, ?, ?)", email, url, secret)

	if err != nil {

		fmt.Println(err.Error())

		return message.JSON(c, 500, "Server error")

	}

	err = smtpMail.SendMail(secret, method+url, email)

	if err != nil {

		return err

	}

	return nil

}

func isEmailValid(e string) bool {

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	return emailRegex.MatchString(e)

}
