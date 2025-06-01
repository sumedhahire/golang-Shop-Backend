package v1

import (
	"github.com/labstack/echo/v4"
	"inventory/config"
	"inventory/internal/inventory"
	customMiddleware "inventory/server/middleware"
)

func InitInventory(v1 *echo.Group, client *config.AppConfig) {

	inventoryRoute := v1.Group("/inventory")

	handler := inventory.NewSHandler(client)

	//USER
	inventorySecure := inventoryRoute.Group("")
	inventorySecure.Use(customMiddleware.OauthValidationUser())

	inventoryRoute.GET("/:id", handler.Get)
	inventoryRoute.GET("", handler.List)

	//ADMIN
	inventoryAdmin := inventoryRoute.Group("/admin")
	inventoryAdmin.Use(customMiddleware.OauthValidationAdmin())

	inventoryAdmin.GET("", handler.List)
	inventoryAdmin.POST("", handler.Add)
	inventoryAdmin.PUT("/:id", handler.Update)
}
