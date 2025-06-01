package tag

import (
	"inventory/ent/entgen"
	"inventory/ent/entgen/tbltag"
)

func get(client *entgen.Client, id string) *entgen.TblTagQuery {
	return client.TblTag.Query().Where(tbltag.ID(id))
}

func list(client *entgen.Client) *entgen.TblTagQuery {
	return client.TblTag.Query()
}
