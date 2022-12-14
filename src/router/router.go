package router

import (
	"github.com/labstack/echo/v4"

	"github.com/Cthulhu-tech/store/src/api/auth"
	middlewares "github.com/Cthulhu-tech/store/src/middleware"
	"github.com/Cthulhu-tech/store/src/utils/database"
)

func Handler() {

	router := echo.New()

	check := router.Group("/", middlewares.Auth)

	database.Connecting()

	router.GET("/login", auth.Login)
	router.GET("/regist", auth.Regist)
	router.GET("/refresh", auth.Refresh)
	router.GET("/confirm", auth.Confirm)

	router.GET("/authVk", auth.Vkontakte)
	router.GET("/confirmsocial", auth.ConfirmSocial)

	check.GET("/refresh", auth.Refresh)
	check.GET("/protected", auth.Protected)

	router.Logger.Fatal(router.Start(":3000"))

}
