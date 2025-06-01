package schema

import (
	"entgo.io/ent/schema/edge"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type TblUser struct {
	ent.Schema
}

func (TblUser) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "tbl_users"},
	}
}

// Fields of the User.
func (TblUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").StorageKey("id_uuid").
			NotEmpty().
			Unique().
			Immutable(), // assuming ULID is assigned once

		field.String("firstname").
			MaxLen(32).
			NotEmpty(),

		field.String("lastname").
			MaxLen(32).
			NotEmpty(),

		field.String("email").
			MaxLen(255).
			NotEmpty().
			Unique(),

		field.Time("birth_date").
			Optional().
			Nillable(),

		field.String("password").
			MaxLen(255).
			NotEmpty(),

		field.Int("is_active").
			Default(1),

		field.Int("zip_code").
			Optional().
			Nillable(),

		field.Text("address").
			Optional().
			Nillable(),

		field.String("ip_address").
			MaxLen(40).
			Optional().
			Nillable(),

		field.Time("created_at").
			Default(time.Now),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),

		field.Time("deleted_at").
			Optional().
			Nillable(),

		field.Enum("role").
			Values("user", "admin").
			Default("user"),
	}
}

// Edges of the User (define relations here if needed).
func (TblUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("UserToken", TblAuthToken.Type),
		edge.To("UserCart", TblCart.Type),
		edge.To("UserPayment", TblPayment.Type),
	}
}
