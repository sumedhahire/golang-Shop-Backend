package v1

import (
	"github.com/labstack/echo/v4"
	"inventory/config"
	"inventory/internal/user"
	customMiddleware "inventory/server/middleware"
)

func InitUser(v1 *echo.Group, client *config.AppConfig) {

	userRoute := v1.Group("/user")
	userRouteSecure := userRoute.Group("")

	userRouteSecure.Use(customMiddleware.OauthValidationUser())

	handler := user.NewSHandler(client)

	userRouteSecure.GET("", handler.Get)
	//userRouteSecure.GET("", handler.List)
	//userRouteSecure.POST("", handler.Add)
	//userRouteSecure.POST("/buy", handler.Buy)
	//userRouteSecure.POST("/verify", handler.Verify)
}
