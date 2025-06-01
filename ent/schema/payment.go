package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

type TblPayment struct {
	ent.Schema
}

func (TblPayment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "Tbl_Payment",
		},
	}
}

func (TblPayment) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").StorageKey("Id_uuid"),

		field.String("InventoryId").StorageKey("InventoryId").Optional(),
		field.String("UserId").StorageKey("UserId").Optional(),
		field.String("Status").StorageKey("Status").Default("created"),

		field.String("RazorpayOrderId").StorageKey("RazorpayOrderId").NotEmpty(),
		//field.String("RazorpayPaymentId").StorageKey("RazorpayPaymentId"),
		//field.String("RazorpaySignature").StorageKey("RazorpaySignature"),

		field.Float32("Amount").StorageKey("Amount"),

		field.Time("Created_at").StorageKey("Created_at").Default(time.Now),
		field.Time("Updated_at").StorageKey("Updated_at"),
		field.Time("Deleted_at").StorageKey("Deleted_at").Nillable().Optional(),
	}
}

func (TblPayment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Inventory", TblInventory.Type).
			Ref("InventoryPayment").
			Field("InventoryId").
			Unique(),

		edge.From("User", TblUser.Type).
			Ref("UserPayment").
			Field("UserId").
			Unique(),
	}
}
