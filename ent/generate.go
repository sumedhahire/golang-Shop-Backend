//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./schema --target=entgen --feature sql/versioned-migration,sql/execquery,sql/upsert,sql/lock
package ent
