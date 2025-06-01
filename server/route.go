package server

import (
	"github.com/labstack/echo/v4"
	"inventory/config"
	"inventory/internal/auth"
	v1Obj "inventory/server/v1"
)

func InitService(e *echo.Echo) {

	appConfig := config.InitServices()

	api := e.Group("/api")

	handler := auth.NewAuthHandler(appConfig.Client)
	api.POST("/login", handler.Login)
	api.GET("/logout", handler.Logout)

	v1 := api.Group("/v1")

	v1Obj.InitInventory(v1, appConfig)
	v1Obj.InitCart(v1, appConfig)
	v1Obj.InitUser(v1, appConfig)
	v1Obj.InitTag(v1, appConfig)
}
