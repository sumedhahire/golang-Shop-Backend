package inventory

import (
	"context"
	"encoding/json"
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

// Get godoc
// @Summary      Get inventory item
// @Description  Returns an inventory item by ID
// @Tags         inventory
// @Produce      json
// @Param        id   path      string  true  "Inventory ID"
// @Success      200  {object}  response.BaseRS{data=inventory.RSInventory,error=interface{}}
// @Failure      404  {object}  response.BaseRS{data=interface{},error=errorhandler.BaseErr}
// @Router       /api/v1/inventory/{id} [get]
func (h SHandler) Get(c echo.Context) error {
	id := c.Param("id")

	ctx := context.Background()
	rs, err := h.Service.Get(ctx, id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusFound, util.ConvertToResponse(rs))
}

// List godoc
// @Summary      List inventory items
// @Description  Returns a list of inventory items with optional filters
// @Tags         inventory
// @Produce      json
// @Success      200  {object}  response.BaseRS{data=[]inventory.RSInventory,error=interface{}}
// @Failure      500  {object}  response.BaseRS{data=interface{},error=errorhandler.BaseErr}
// @Router       /api/v1/inventory [get]
func (h SHandler) List(c echo.Context) error {
	filter := GetFilters(c)
	ctx := context.Background()
	rs, err := h.Service.List(ctx, filter)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, util.ConvertToListResponse(rs))
}

// Add godoc
// @Summary      Add inventory item
// @Description  Adds a new inventory item with image and metadata
// @Tags         inventory
// @Accept       multipart/form-data
// @Produce      json
// @Param        image        formData  file    true   "Product image"
// @Param        name         formData  string  true   "Product name"
// @Param        price        formData  number  true   "Product price"
// @Param        description  formData  string  false  "Product description"
// @Param        isActive     formData  bool    true   "Product active status"
// @Param        tag          formData  string  true   "JSON array of tag IDs"
// @Success      201  {object}  response.BaseRS{data=inventory.RSInventory,error=interface{}}
// @Failure      400  {object}  response.BaseRS{data=interface{},error=errorhandler.BaseErr}
// @Failure      500  {object}  response.BaseRS{data=interface{},error=errorhandler.BaseErr}
// @Router       /api/v1/inventory/admin [post]
// @Security     BearerAuth
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
	tagJson := c.FormValue("tag")

	var tags []string
	err = json.Unmarshal([]byte(tagJson), &tags)
	if err != nil {
		return err
	}

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
		Tag:         tags,
	}

	ctx := context.Background()
	rs, err := h.Service.Add(ctx, rq, file)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, util.ConvertToResponse(rs))
}

// Update godoc
// @Summary      Update inventory item
// @Description  Updates an existing inventory item by ID
// @Tags         inventory
// @Accept       json
// @Produce      json
// @Param        id   path   string       true  "Inventory ID"
// @Param body inventory.RQInventory true "Updated inventory data"
// @Success      200  {object}  response.BaseRS{data=inventory.RSInventory,error=interface{}}
// @Failure      400  {object}  response.BaseRS{data=interface{},error=errorhandler.BaseErr}
// @Failure      404  {object}  response.BaseRS{data=interface{},error=errorhandler.BaseErr}
// @Router       /api/v1/inventory/admin/{id} [put]
// @Security     BearerAuth
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
