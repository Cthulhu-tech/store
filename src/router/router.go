package router

import (

	"github.com/labstack/echo/v4"

	"github.com/Cthulhu-tech/store/src/api/auth"
	"github.com/Cthulhu-tech/store/src/middleware"

)


func Handler() {

	router := echo.New()

	check := router.Group("/", middlewares.Auth)

	router.GET("/login", auth.Login)
	router.GET("/regist", auth.Regist)

	router.GET("/auth", auth.Vkontakte)

	check.GET("refresh", auth.Refresh)
	check.GET("protected", auth.Protected)

	router.Logger.Fatal(router.Start(":3000"))

}
