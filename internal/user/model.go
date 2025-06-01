package user

import (
	"inventory/ent/entgen"
	"time"
)

type User struct {
	Id        string
	FirstName string
	LastName  string
	Email     string
	BirthDate time.Time
	Password  string
	Address   string
	Role      string
	IsActive  bool
}

func (user *User) MapFrom(entUser *entgen.TblUser) {
	user.Id = entUser.ID
	user.Email = entUser.Email
	user.Address = *entUser.Address
	user.FirstName = entUser.Firstname
	user.LastName = entUser.Lastname
	user.BirthDate = *entUser.BirthDate
	user.Role = string(entUser.Role)
	user.IsActive = entUser.IsActive
}

type RSUser struct {
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	BirthDate time.Time `json:"birthDate"`
	Role      string    `json:"role"`
	IsActive  bool      `json:"isActive"`
}

func (rs *RSUser) MapFrom(user *User) {
	rs.Email = user.Email
	rs.FirstName = user.FirstName
	rs.LastName = user.LastName
	rs.Address = user.Address
	rs.BirthDate = user.BirthDate
	rs.Role = user.Role
	rs.IsActive = user.IsActive
}
