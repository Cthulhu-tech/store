package middlewares

import (

	"net/http"
	"github.com/labstack/echo/v4"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {

		if(true){

			return next(c)

		}

		return c.JSON(http.StatusForbidden, "You are not authorized!")

    }

}
