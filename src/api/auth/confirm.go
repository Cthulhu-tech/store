package auth

import (
	"encoding/json"
	"log"

	"github.com/Cthulhu-tech/store/src/utils/database"
	"github.com/Cthulhu-tech/store/src/utils/message"
	"github.com/labstack/echo/v4"
)

func Confirm(c echo.Context) error {

	var data = &UserConfirm{}

	err := json.NewDecoder(c.Request().Body).Decode(&data)

	if err != nil {

		return err

	}

	if data.Code == "" || data.Secret == 0 {

		return message.JSON(c, 403, "Invalid body parameters")

	}

	db := database.GetDB()

	rows, err := db.Query("SELECT `sp_confirm`(?, ?) AS `sp_confirm`", data.Code, data.Secret)

	defer rows.Close()

	if err != nil {

		return invalidQueryParam(c)

	}

	value := &ConfirmFunc{}

	for rows.Next() {

		var confirm ConfirmFunc

		if err := rows.Scan(&confirm.Value); err != nil {

			log.Println(err.Error())

		}

		value.Value = confirm.Value

	}

	if value.Value == 1 {

		return message.JSON(c, 201, "User confirm")

	}

	return message.JSON(c, 403, "Confirm error")

}
