package inventory

import (
	"github.com/google/uuid"
	"inventory/ent/entgen"
	"inventory/ent/entgen/tblinventory"
	"inventory/ent/entgen/tblinventorytag"
	"time"
)

func getLock(client *entgen.Tx, productId string) *entgen.TblInventoryQuery {
	return client.TblInventory.Query().Where(tblinventory.ID(productId)).ForUpdate()
}

func get(client *entgen.Client, id string) *entgen.TblInventoryQuery {
	return client.TblInventory.Query().Where(tblinventory.ID(id),
		tblinventory.IsActive(true)).WithInventoryTag(func(query *entgen.TblInventoryTagQuery) {
		query.WithTag()
	})
}

func list(client *entgen.Client, filter Filter) *entgen.TblInventoryQuery {
	query := client.TblInventory.Query().
		Where(tblinventory.IsActive(true), tblinventory.DeletedAtIsNil())

	if filter.Name != "" {
		query.Where(tblinventory.NameContainsFold(filter.Name))
	}

	return query
}

func getTags(client *entgen.Client, inventory string) *entgen.TblInventoryTagQuery {
	return client.TblInventoryTag.Query().Where(tblinventorytag.InventoryId(inventory)).WithTag(func(query *entgen.TblTagQuery) {

	})

}

func add(client *entgen.Tx, inventory Inventory) *entgen.TblInventoryCreate {
	return client.TblInventory.Create().
		SetID(inventory.Id).
		SetName(inventory.Name).
		SetImageLink(inventory.ImageUrl).
		SetIsActive(inventory.IsActive).
		SetPrice(inventory.Price).
		SetDescription(inventory.Description).
		SetCreatedAt(time.Now().UTC()).
		SetUpdatedAt(time.Now().UTC())
}

func addTag(client *entgen.Tx, inventoryId string, tag string) *entgen.TblInventoryTagCreate {
	return client.TblInventoryTag.Create().SetID(uuid.NewString()).SetInventoryId(inventoryId).SetTagId(tag).SetCreatedAt(time.Now().UTC()).SetUpdatedAt(time.Now().UTC())
}

func update(client *entgen.Tx, inventory RQInventory, productId string) *entgen.TblInventoryUpdateOne {
	query := client.TblInventory.UpdateOneID(productId).
		SetUpdatedAt(time.Now().UTC()).
		SetIsActive(inventory.IsActive).
		SetPrice(inventory.Price).
		SetDescription(inventory.Description).
		SetName(inventory.Name)

	return query
}
