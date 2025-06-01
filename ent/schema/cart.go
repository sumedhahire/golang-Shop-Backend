package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

type TblCart struct {
	ent.Schema
}

func (TblCart) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "Tbl_Cart",
		},
	}
}

func (TblCart) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").StorageKey("Id_uuid"),

		field.String("ProductId").StorageKey("ProductId").Optional(),
		field.String("UserId").StorageKey("UserId").Optional(),
		field.Enum("Status").
			Values("cart", "canceled", "brought").
			Default("cart"),

		field.Time("Created_at").StorageKey("Created_at").Default(time.Now),
		field.Time("Updated_at").StorageKey("Updated_at"),
		field.Time("Deleted_at").StorageKey("Deleted_at").Nillable().Optional(),
	}
}

func (TblCart) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Inventory", TblInventory.Type).
			Ref("InventoryCart").
			Field("ProductId").
			Unique(),

		edge.From("User", TblUser.Type).
			Ref("UserCart").
			Field("UserId").
			Unique(),
	}
}
