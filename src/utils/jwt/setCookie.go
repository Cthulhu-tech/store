package jwt 

import (
	"time"
	"net/http"

	"github.com/labstack/echo/v4"

)

func SetCookie(c echo.Context, duration int, refreshToken string) {
	
	cookie := new(http.Cookie)

	cookie.Name = "refresh"

	cookie.Value = refreshToken

	cookie.Expires = time.Now().Add(time.Duration(duration) * time.Hour)
	
	c.SetCookie(cookie)

}
