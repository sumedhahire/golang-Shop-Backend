package v1

import (
	"github.com/labstack/echo/v4"
	"inventory/config"
	"inventory/internal/user"
	customMiddleware "inventory/server/middleware"
)

func InitUser(v1 *echo.Group, client *config.AppConfig) {

	userRoute := v1.Group("/user")

	handler := user.NewSHandler(client)

	userRouteSecure := userRoute.Group("")
	userRouteSecure.Use(customMiddleware.OauthValidationUser())

	//USER
	userRouteSecure.GET("", handler.Get)

	//ADMIN
	userRouteAdmin := userRoute.Group("/admin")
	userRouteAdmin.Use(customMiddleware.OauthValidationAdmin())
	userRouteAdmin.GET("", handler.List)
	userRouteAdmin.POST("", handler.Add)
	//userRouteSecure.POST("", handler.Add)
	//userRouteSecure.POST("/buy", handler.Buy)
	//userRouteSecure.POST("/verify", handler.Verify)
}
