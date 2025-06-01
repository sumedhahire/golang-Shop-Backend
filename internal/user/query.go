package user

import (
	"inventory/ent/entgen"
	"inventory/ent/entgen/tbluser"
)

func get(client *entgen.Client, id string) *entgen.TblUserQuery {
	return client.TblUser.Query().
		Where(tbluser.IDEQ(id))
}
