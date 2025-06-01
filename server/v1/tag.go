package v1

import (
	"github.com/labstack/echo/v4"
	"inventory/config"
	"inventory/internal/tag"

	customMiddleware "inventory/server/middleware"
)

func InitTag(v1 *echo.Group, client *config.AppConfig) {

	tagRoute := v1.Group("/tag")

	handler := tag.NewHandler(client)

	tagRouteSecure := tagRoute.Group("")
	tagRouteSecure.Use(customMiddleware.OauthValidationUser())

	tagRouteSecure.GET("", handler.List)

	//userRouteSecure.POST("", handler.Add)
	//userRouteSecure.POST("/buy", handler.Buy)
	//userRouteSecure.POST("/verify", handler.Verify)
}
