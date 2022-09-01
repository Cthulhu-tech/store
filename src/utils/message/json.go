package message

import "github.com/labstack/echo/v4"

func JSON(c echo.Context, status int, message string) error {

	valueMessage := &Message{}

	valueMessage.Message = message

	return c.JSON(status, valueMessage)

}
