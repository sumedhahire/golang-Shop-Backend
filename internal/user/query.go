package user

import (
	"inventory/ent/entgen"
	"inventory/ent/entgen/tbluser"
	"time"
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

func add(client *entgen.Tx, user User) *entgen.TblUserCreate {
	query := client.TblUser.Create().SetID(user.Id).
		SetIsActive(true).
		SetPassword(user.Password).
		SetAddress(user.Address).
		SetFirstname(user.FirstName).
		SetLastname(user.LastName).
		SetEmail(user.Email).
		SetBirthDate(user.BirthDate).
		SetCreatedAt(time.Now().UTC()).
		SetUpdatedAt(time.Now().UTC())

	return query
}

func checkByMail(client *entgen.Tx, mail string) *entgen.TblUserQuery {
	query := client.TblUser.Query().Where(tbluser.EmailEQ(mail))
	return query
}
