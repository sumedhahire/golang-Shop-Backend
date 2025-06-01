package inventory

import (
	"context"
	"fmt"
	"inventory/ent/entgen"
	"inventory/internal/util"
)

type IStorage interface {
	Get(ctx context.Context, id string) (*Inventory, error)
	List(ctx context.Context, filter Filter) ([]Inventory, error)
	Add(ctx context.Context, inventory *Inventory) (string, error)
	Update(ctx context.Context, updateId string, inventory RQInventory) error
}

type sStorage struct {
	client *entgen.Client
}

func NewStorage(client *entgen.Client) IStorage {
	return &sStorage{client: client}
}

func (s sStorage) Get(ctx context.Context, id string) (*Inventory, error) {
	entInventory, err := get(s.client, id).First(ctx)
	if err != nil {
		return nil, util.WrapperForDatabaseError("get", err)
	}

	var inventory Inventory
	inventory.MapTo(entInventory)
	for _, invTag := range entInventory.Edges.InventoryTag {
		inventory.Tags = append(inventory.Tags, invTag.Edges.Tag.Name)

	}
	return &inventory, nil
}

func (s sStorage) List(ctx context.Context, filter Filter) ([]Inventory, error) {
	entInventory, err := list(s.client, filter).All(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, util.WrapperForDatabaseError("list", err)
	}

	inventories := make([]Inventory, len(entInventory))
	for i, inventoryTag := range entInventory {
		entTags, err := getTags(s.client, inventoryTag.ID).All(ctx)
		if err != nil {
			return nil, err
		}
		var inventory Inventory
		inventory.MapTo(inventoryTag)
		for _, tags := range entTags {
			tag := tags.Edges.Tag
			inventory.Tags = append(inventory.Tags, tag.Name)
		}
		inventories[i] = inventory
	}

	return inventories, nil
}

func (s sStorage) Add(ctx context.Context, inventory *Inventory) (string, error) {
	var id string
	err := util.ExecTx(ctx, s.client, func(tx *entgen.Tx) error {
		inv, err := add(tx, *inventory).Save(ctx)
		if err != nil {
			fmt.Println("dfdsffsdf", err)
			return util.WrapperForDatabaseError("add", err)
		}
		id = inv.ID

		for _, tag := range inventory.Tags {
			err := addTag(tx, id, tag).Exec(ctx)
			if err != nil {
				fmt.Println("$$$$$$$$$", err)
				return err
			}
		}

		return nil
	})
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return id, nil
}

func (s sStorage) Update(ctx context.Context, updateId string, inventory RQInventory) error {
	err := util.ExecTx(ctx, s.client, func(tx *entgen.Tx) error {
		_, err := getLock(tx, updateId).First(ctx)
		if err != nil {
			return err
		}

		err = update(tx, inventory, updateId).Exec(ctx)
		if err != nil {
			return err
		}
		return nil

	})
	if err != nil {
		return err
	}
	return nil
}
