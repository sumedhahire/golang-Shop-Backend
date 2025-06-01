package tag

import (
	"context"
	"github.com/labstack/echo/v4"
	"inventory/config"
	"inventory/internal/util"
	"net/http"
)

type IHandler interface {
	List(c echo.Context) error
}

type SHandler struct {
	service IService
}

func NewHandler(config *config.AppConfig) IHandler {
	return &SHandler{service: NewService(config.Client)}
}

// List godoc
// @Summary      List tags
// @Description  Returns a user based on the authenticated user ID from the context
// @Tags         user
// @Produce      json
// @Success      200  {object} response.BaseRS(data=([]tag.RSTag),error={})
// @Failure 400 {object} response.BaseRS(data={},error=(errorhandler.BaseErr))  // Return error in case of failure (adjust this to match your actual error handler)
// @Failure 404 {object} response.BaseRS(data={},error=(errorhandler.BaseErr))  // Return error if the user to delete is not found
// @Router       /api/v1/tag [get]
func (h *SHandler) List(c echo.Context) error {
	ctx := context.Background()
	rs, err := h.service.List(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, util.ConvertToListResponse(rs))
}
