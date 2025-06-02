package user

import (
	"github.com/google/uuid"
	"github.com/sethvargo/go-password/password"
	"inventory/ent/entgen"
	"time"
)

type RQUser struct {
	Id        string    `json:"-"`
	FirstName string    `json:"firstName" validate:"required"`
	LastName  string    `json:"lastName" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Birthdate time.Time `json:"birthdate" validate:"required"`
	Address   string    `json:"address" validate:"required"`
}

func (r *RQUser) MapTO() *User {
	var user User
	if r.Id == "" {
		r.Id = uuid.NewString()
	}
	user.Id = r.Id
	user.FirstName = r.FirstName
	user.LastName = r.LastName
	user.Email = r.Email
	user.BirthDate = r.Birthdate.UTC()
	user.Password, _ = password.Generate(64, 10, 10, false, false)
	user.Address = r.Address

	return &user
}

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
