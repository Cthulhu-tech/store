package auth

import (
	"encoding/json"

	"github.com/Cthulhu-tech/store/src/utils/database"
	"github.com/Cthulhu-tech/store/src/utils/message"
	"github.com/labstack/echo/v4"
)

func ConfirmSocial(c echo.Context) error {

	var data = &UserRegistrationConfirm{}

	err := json.NewDecoder(c.Request().Body).Decode(&data)

	if err != nil {

		return err

	}

	if data.Login == "" || data.Password == "" || data.Code == "" || data.Secret == 0 {

		return message.JSON(c, 403, "Invalid body parameters")

	}

	db := database.GetDB()

	rows, err := db.Query("SELECT `sp_cofnfirm_vk`(?, ?, ?, ?) AS `sp_cofnfirm_vk`", data.Login, data.Password, data.Code, data.Secret)

	defer rows.Close()

	if err != nil {

		return invalidQueryParam(c)

	}

	value := &ConfirmVkFunc{}

	for rows.Next() {

		var confirm ConfirmVkFunc

		if err := rows.Scan(&confirm.Value); err != nil {

			return message.JSON(c, 403, "Invalid body parameters")

		}

		value.Value = confirm.Value

	}

	if value.Value == 1 {

		return message.JSON(c, 201, "User confirm")

	}

	return message.JSON(c, 403, "Confirm error")

}
