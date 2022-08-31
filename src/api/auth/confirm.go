package auth

import (
	"github.com/Cthulhu-tech/store/src/utils/database"
	"github.com/labstack/echo/v4"
)

func Confirm(c echo.Context) error {

	code := c.QueryParam("code")

	if len(code) == 0 {

		return invalidQueryParam(c)

	}

	db := database.GetDB()

	_, err := db.Query("", code)

	if err != nil {

		return invalidQueryParam(c)

	}

	return nil

}
