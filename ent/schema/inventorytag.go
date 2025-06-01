package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

type TblInventoryTag struct {
	ent.Schema
}

func (TblInventoryTag) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "Tbl_Inventory_Tag",
		},
	}
}

func (TblInventoryTag) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").StorageKey("Id_uuid"),
		field.String("InventoryId").StorageKey("InventoryId").Optional(),
		field.String("TagId").StorageKey("TagId").Optional(),

		field.Time("Created_at").StorageKey("Created_at").Default(time.Now),
		field.Time("Updated_at").StorageKey("Updated_at"),
		field.Time("Deleted_at").StorageKey("Deleted_at").Nillable(),
	}
}

func (TblInventoryTag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tag_Id", TblTag.Type).
			Ref("tag").
			Unique(),

		edge.From("inventory_Id", TblInventory.Type).
			Ref("inventory").
			Unique(),
	}
}
