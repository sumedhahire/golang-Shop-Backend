package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

type TblAuthClient struct {
	ent.Schema
}

func (TblAuthClient) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").StorageKey("client_uuid").Unique(),
		field.String("Client_secret").StorageKey("Client_secret"),
		field.String("Grant_type").StorageKey("Grant_type").Optional().Nillable(),

		field.Time("CreatedAt").StorageKey("CreatedAt").Default(time.Now().UTC()).Immutable(),
		field.Time("UpdatedAt").StorageKey("UpdatedAt").Default(time.Now().UTC()),
		field.Time("DeletedAt").StorageKey("DeletedAt").Optional().Nillable(),

		field.String("domain").StorageKey("domain").Optional().Nillable(),

		field.Bool("public").StorageKey("public").Default(true),
		field.String("UserAgent").StorageKey("UserAgent").Default("test"),
		field.String("redirect_uri").StorageKey("redirect_uri").Optional().Nillable(),

		field.String("CreatedBy").StorageKey("CreatedBy").Default("test"),
		field.String("UpdatedBy").StorageKey("UpdatedBy").Default("test"),
		field.String("DeletedBy").StorageKey("DeletedBy").Optional().Nillable(),
	}
}

func (TblAuthClient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("ClientToken", TblAuthToken.Type),
	}
}

func (TblAuthClient) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "authclient",
		},
	}
}
