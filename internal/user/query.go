package user

import (
	"inventory/ent/entgen"
	"inventory/ent/entgen/tbluser"
)

func get(client *entgen.Client, id string) *entgen.TblUserQuery {
	return client.TblUser.Query().
		Where(tbluser.IDEQ(id))
}

func list(client *entgen.Client) *entgen.TblUserQuery {
	return client.TblUser.Query().
		Where(tbluser.DeletedAtIsNil()).
		Where(tbluser.RoleEQ(tbluser.RoleUser))
}

func inActiveUser(client *entgen.Tx, id string) *entgen.TblUserUpdateOne {
	return client.TblUser.
		UpdateOneID(id).
		SetIsActive(false)
}
func activeUser(client *entgen.Tx, id string) *entgen.TblUserUpdateOne {
	return client.TblUser.
		UpdateOneID(id).
		SetIsActive(true)
}
