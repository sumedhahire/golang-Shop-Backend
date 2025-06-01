package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

type TblAuthToken struct {
	ent.Schema
}

func (TblAuthToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").StorageKey("id_uuid").Unique(),
		field.String("auth_uuid").StorageKey("auth_uuid"),
		field.String("auth_xref").StorageKey("auth_xref"),

		field.String("accesstoken").StorageKey("accesstoken"),
		field.Time("accesstokencreatedat").StorageKey("accesstokencreatedat").Default(time.Now().UTC()).Immutable(),
		field.Int("accesstokenexpiresin").StorageKey("accesstokenexpiresin"),

		field.String("clientid").StorageKey("clientid").Optional(),
		field.String("user_ulid").StorageKey("user_ulid").Optional(),

		field.String("refreshtoken").StorageKey("refreshtoken"),
		field.Time("refreshtokencreatedat").StorageKey("refreshtokencreatedat").Default(time.Now().UTC()).Immutable(),
		field.Int("refreshtokenexpiresin").StorageKey("refreshtokenexpiresin"),

		field.Time("createdat").StorageKey("createdat").Default(time.Now().UTC()).Immutable(),
		field.Time("updatedat").StorageKey("updatedat").Default(time.Now().UTC()),
		field.Time("deletedat").StorageKey("deletedat").Optional().Nillable(),

		field.String("ip_address").StorageKey("ip_address").Default("::1"),
		field.String("useragent").StorageKey("useragent").Default("test"),

		field.String("createdby").StorageKey("createdby").Default("test"),
		field.String("updatedby").StorageKey("updatedby").Default("test"),
		field.String("deletedby").StorageKey("deletedby").Optional().Nillable(),

		field.String("code").StorageKey("code"),
		field.Time("codecreatedat").StorageKey("codecreatedat").Default(time.Now().UTC()).Immutable(),
		field.Int("codeexpiresin").StorageKey("codeexpiresin"),
		field.String("codechallenge").StorageKey("codechallenge").Optional().Nillable(),

		field.String("redirect_uri").StorageKey("redirect_uri").Optional().Nillable(),
		field.String("scope").StorageKey("scope").Default("test"),
	}
}

func (TblAuthToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("TokenClient", TblAuthClient.Type).
			Ref("ClientToken").
			Field("clientid").
			Unique(),

		edge.From("TokenUser", TblUser.Type).
			Ref("UserToken").
			Field("user_ulid").
			Unique(),
	}
}

func (TblAuthToken) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "authtoken",
		},
	}
}
