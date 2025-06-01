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
	List(e echo.Context) error
	ChangeActive(e echo.Context) error
}

type SHandler struct {
	appConfig *config.AppConfig
	Service   IUserService
}

func NewSHandler(client *config.AppConfig) IHandler {
	return SHandler{
		appConfig: client,
		Service:   NewService(client.Client),
	}
}

// Get godoc
// @Summary      Get user by ID
// @Description  Returns a user based on the authenticated user ID from the context
// @Tags         user
// @Produce      json
// @Success      200  {object} response.BaseRS(data=(user.RSUser),error={})
// @Failure 400 {object} response.BaseRS(data={},error=(errorhandler.BaseErr))  // Return error in case of failure (adjust this to match your actual error handler)
// @Failure 404 {object} response.BaseRS(data={},error=(errorhandler.BaseErr))  // Return error if the user to delete is not found
// @Router       /api/v1/user [get]
func (h SHandler) Get(c echo.Context) error {
	id := c.Get("userId").(string)

	rs, err := h.Service.Get(context.Background(), id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, util.ConvertToResponse(rs))
}

// List godoc
// @Summary      List user
// @Description  Returns a user based on the authenticated user ID from the context
// @Tags         user
// @Produce      json
// @Success      200  {object} response.BaseRS(data=([]user.RSUser),error={})
// @Failure 400 {object} response.BaseRS(data={},error=(errorhandler.BaseErr))  // Return error in case of failure (adjust this to match your actual error handler)
// @Failure 404 {object} response.BaseRS(data={},error=(errorhandler.BaseErr))  // Return error if the user to delete is not found
// @Router       /api/v1/user/admin [get]
func (h SHandler) List(c echo.Context) error {
	rs, err := h.Service.List(context.Background())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, util.ConvertToListResponse(rs))
}

func (h SHandler) ChangeActive(c echo.Context) error {
	rs, err := h.Service.List(context.Background())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, util.ConvertToListResponse(rs))
}
