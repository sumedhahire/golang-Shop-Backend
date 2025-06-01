package inventory

import (
	"context"
	"github.com/labstack/echo/v4"
	"inventory/config"
	"inventory/internal/util"
	"net/http"
	"strconv"
	"strings"
)

type IHandler interface {
	Get(e echo.Context) error
	List(e echo.Context) error
	ListAll(e echo.Context) error
	Add(e echo.Context) error
	Update(e echo.Context) error
}

type SHandler struct {
	appConfig *config.AppConfig
	Service   IService
}

func NewSHandler(client *config.AppConfig) SHandler {
	return SHandler{
		appConfig: client,
		Service:   NewService(client),
	}
}

// /v1/inventory/{id}
func (h SHandler) Get(c echo.Context) error {
	id := c.Param("id")

	ctx := context.Background()
	rs, err := h.Service.Get(ctx, id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusFound, util.ConvertToResponse(rs))
}

// /api/v1/inventory
func (h SHandler) List(c echo.Context) error {
	filter := GetFilters(c)
	ctx := context.Background()
	rs, err := h.Service.List(ctx, filter)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, util.ConvertToListResponse(rs))
}

// /v1/inventory/admin
func (h SHandler) Add(c echo.Context) error {
	file, err := c.FormFile("image")
	if err != nil {
		return err
	}

	// 2. Parse text fields
	name := c.FormValue("name")
	priceStr := c.FormValue("price")
	description := c.FormValue("description")
	isActiveStr := c.FormValue("isActive")

	// 3. Convert price to float64
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return err
	}

	// 4. Convert isActive to bool
	isActive := false
	if strings.ToLower(isActiveStr) == "true" {
		isActive = true
	}

	// 5. Create request model
	rq := RQInventory{
		Name:        name,
		Price:       float32(price),
		Description: description,
		IsActive:    isActive,
	}

	ctx := context.Background()
	rs, err := h.Service.Add(ctx, rq, file)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, util.ConvertToResponse(rs))
}

// /v1/inventory/admin/:id
func (h SHandler) Update(c echo.Context) error {
	var rq RQInventory
	err := util.BindAndValidate(&rq, c)
	if err != nil {
		return err
	}

	productId := c.Param("id")

	ctx := context.Background()
	rs, err := h.Service.Update(ctx, rq, productId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, util.ConvertToResponse(rs))
}
