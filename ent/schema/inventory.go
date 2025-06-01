package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

type TblInventory struct {
	ent.Schema
}

func (TblInventory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "Tbl_Inventory",
		},
	}
}

func (TblInventory) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").StorageKey("Id_uuid"),

		field.String("Name").StorageKey("Name").NotEmpty(),
		field.String("Description").StorageKey("Description").NotEmpty(),
		field.String("ImageLink").StorageKey("ImageLink"),
		field.Float32("Price").StorageKey("Price"),

		field.Bool("Is_Active").StorageKey("Is_Active").Default(true),

		field.Time("Created_at").StorageKey("Created_at").Default(time.Now),
		field.Time("Updated_at").StorageKey("Updated_at"),
		field.Time("Deleted_at").StorageKey("Deleted_at").Nillable().Optional(),
	}
}

func (TblInventory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("inventory", TblInventoryTag.Type),
		edge.To("InventoryCart", TblCart.Type),
		edge.To("InventoryPayment", TblPayment.Type),
	}
}
