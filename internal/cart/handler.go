package cart

import (
	"context"
	"github.com/labstack/echo/v4"
	"inventory/config"
	"inventory/internal/util"
	"net/http"
)

type ICartHandler interface {
	Add(e echo.Context) error
	Get(e echo.Context) error
	List(e echo.Context) error
	Buy(e echo.Context) error
	Verify(e echo.Context) error
	BuyCount(e echo.Context) error
}
type SCartHandler struct {
	appConfig *config.AppConfig
	service   ICartService
}

func NewCartHandler(client *config.AppConfig) ICartHandler {
	return &SCartHandler{
		appConfig: client,
		service:   NewCartService(client),
	}
}

func (h *SCartHandler) Add(e echo.Context) error {
	var rq RQCart
	err := util.BindAndValidate(&rq, e)
	if err != nil {
		return err
	}
	rq.UserID = e.Get("userId").(string)

	ctx := context.Background()
	rs, err := h.service.Add(ctx, rq)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusCreated, util.ConvertToResponse(rs))
}

func (h *SCartHandler) Get(e echo.Context) error {
	id := e.Param("id")

	ctx := context.Background()
	rs, err := h.service.Get(ctx, id)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusCreated, util.ConvertToResponse(rs))
}

func (h *SCartHandler) List(e echo.Context) error {
	userId := e.Get("userId").(string)
	status := e.QueryParam("status")

	ctx := context.Background()
	rs, err := h.service.List(ctx, userId, status)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusCreated, util.ConvertToListResponse(rs))
}

func (h *SCartHandler) Buy(e echo.Context) error {
	var rq RQCart
	err := util.BindAndValidate(&rq, e)
	if err != nil {
		return err
	}
	userId := e.Get("userId").(string)
	rq.UserID = userId

	ctx := context.Background()
	rs, err := h.service.Buy(ctx, rq)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusCreated, util.ConvertToResponse(rs))
}

func (h *SCartHandler) Verify(e echo.Context) error {
	var rq RQPayment
	err := util.BindAndValidate(&rq, e)
	if err != nil {
		return err
	}
	//
	ctx := context.Background()

	err = h.service.Verify(ctx, rq)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, util.ConvertToResponse(map[string]string{"message": "payment verified successfully"}))

}

func (h *SCartHandler) BuyCount(e echo.Context) error {
	userId := e.Get("userId").(string)

	ctx := context.Background()
	rs, err := h.service.BuyCount(ctx, userId)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, util.ConvertToResponse(rs))
}
