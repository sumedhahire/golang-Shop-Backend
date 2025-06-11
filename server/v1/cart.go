package v1

import (
	"github.com/labstack/echo/v4"
	"inventory/config"
	"inventory/internal/cart"
	customMiddleware "inventory/server/middleware"
)

func InitCart(v1 *echo.Group, client *config.AppConfig) {

	cartRoute := v1.Group("/cart")
	cartRouteSecure := cartRoute.Group("")

	cartRouteSecure.Use(customMiddleware.OauthValidationUser())

	handler := cart.NewCartHandler(client)

	cartRouteSecure.GET("/:id", handler.Get)
	cartRouteSecure.GET("", handler.List)
	cartRouteSecure.POST("", handler.Add)
	cartRouteSecure.POST("/buy", handler.Buy)
	cartRouteSecure.POST("/verify", handler.Verify)
	cartRouteSecure.GET("/count", handler.BuyCount)
	cartRouteSecure.GET("/invoice/:productId", handler.GetInvoice)
}
