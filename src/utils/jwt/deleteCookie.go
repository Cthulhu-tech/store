package jwt 

import (
	"time"
	"net/http"

	"github.com/labstack/echo/v4"

)

func DeleteCookie(c echo.Context) {
	
	cookie := new(http.Cookie)

	cookie.Name = "refresh"

	cookie.Value = ""

	cookie.Expires = time.Now().Add(-1 * time.Second)
	
	c.SetCookie(cookie)
	
}
