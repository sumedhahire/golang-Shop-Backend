package inventory

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"inventory/ent/entgen"
	"time"
)

type (
	RQInventory struct {
		Name        string   `json:"name" validate:"required,min=3,max=32" example:"test"`
		Tag         []string `json:"tag" `
		Price       float32  `json:"price" validate:"required" example:"10.00"`
		Description string   `json:"description" validate:"required"`
		IsActive    bool     `json:"isActive" validate:"required"`
	}

	Inventory struct {
		Id          string
		Name        string
		Tags        []string
		Price       float32
		Description string
		ImageUrl    string
		IsActive    bool

		UpdatedAt time.Time
		CreatedAt time.Time
		DeletedAt time.Time
	}

	TagData struct {
		Name []string `json:"tags"`
	}

	RSInventory struct {
		Id          string   `json:"id"`
		Name        string   `json:"name"`
		Description string   `json:"description"`
		Tags        []string `json:"tags"`
		Price       float32  `json:"price"`
		ImageUrl    string   `json:"imageUrl"`
	}

	Filter struct {
		Name string `json:"name"`
	}
)

func (inventory *Inventory) MapTo(entInventory *entgen.TblInventory) {
	inventory.Id = entInventory.ID
	inventory.Name = entInventory.Name
	inventory.ImageUrl = entInventory.ImageLink
	inventory.Price = entInventory.Price
	inventory.Description = entInventory.Description
	inventory.IsActive = entInventory.IsActive
	inventory.UpdatedAt = entInventory.UpdatedAt
	inventory.CreatedAt = entInventory.CreatedAt
}

func (rs *RSInventory) MapTo(inventory *Inventory) {
	rs.Id = inventory.Id
	rs.Name = inventory.Name
	rs.Description = inventory.Description
	rs.ImageUrl = inventory.ImageUrl
	rs.Tags = inventory.Tags
	rs.Price = inventory.Price
}

func (rq *RQInventory) MapFrom() Inventory {
	var inventory Inventory

	inventory.Id = uuid.NewString()
	inventory.Name = rq.Name
	inventory.Description = rq.Description
	inventory.Tags = rq.Tag
	inventory.Price = rq.Price
	inventory.IsActive = rq.IsActive
	return inventory
}

func GetFilters(e echo.Context) Filter {
	data := e.QueryParam("filter")

	if data == "" {
		return Filter{} // no filter, return empty filter
	}

	var filter Filter
	err := json.Unmarshal([]byte(data), &filter)
	if err != nil {
		fmt.Println(err)
		return Filter{}
	}
	fmt.Println(filter)
	return filter
}
