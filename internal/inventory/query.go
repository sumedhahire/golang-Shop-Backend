package inventory

import (
	"inventory/ent/entgen"
	"inventory/ent/entgen/tblinventory"
	"time"
)

func getLock(client *entgen.Tx, productId string) *entgen.TblInventoryQuery {
	return client.TblInventory.Query().Where(tblinventory.ID(productId)).ForUpdate()
}

func get(client *entgen.Client, id string) *entgen.TblInventoryQuery {
	return client.TblInventory.Query().Where(tblinventory.ID(id),
		tblinventory.IsActive(true))
}

func list(client *entgen.Client, filter Filter) *entgen.TblInventoryQuery {
	query := client.TblInventory.Query().
		Where(tblinventory.IsActive(true),
			tblinventory.DeletedAtIsNil())

	if filter.Name != "" {
		query.Where(tblinventory.NameContainsFold(filter.Name))
	}
	return query
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

func update(client *entgen.Tx, inventory RQInventory, productId string) *entgen.TblInventoryUpdateOne {
	query := client.TblInventory.UpdateOneID(productId).
		SetUpdatedAt(time.Now().UTC()).
		SetIsActive(inventory.IsActive).
		SetPrice(inventory.Price).
		SetDescription(inventory.Description).
		SetName(inventory.Name)

	return query
}
