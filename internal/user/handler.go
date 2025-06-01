package user

import (
	"context"
	"github.com/labstack/echo/v4"
	"inventory/config"
	"inventory/internal/util"
	"net/http"
)

type IHandler interface {
	Get(e echo.Context) error
}

type SHandler struct {
	appConfig *config.AppConfig
	Service   IUserService
}

func NewSHandler(client *config.AppConfig) SHandler {
	return SHandler{
		appConfig: client,
		Service:   NewService(client.Client),
	}
}

func (h SHandler) Get(c echo.Context) error {
	id := c.Get("userId").(string)

	rs, err := h.Service.Get(context.Background(), id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, util.ConvertToResponse(rs))
}
