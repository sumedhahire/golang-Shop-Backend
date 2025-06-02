package user

import (
	"context"
	"errors"
	"fmt"
	"inventory/ent/entgen"
	"inventory/internal/util"
)

type IUserStorage interface {
	Get(ctx context.Context, id string) (*User, error)
	List(ctx context.Context) ([]User, error)
	Add(ctx context.Context, user *User) error
}

type sStorage struct {
	client *entgen.Client
}

func NewStorage(client *entgen.Client) IUserStorage {
	return &sStorage{client: client}
}

func (s sStorage) Get(ctx context.Context, id string) (*User, error) {
	entUser, err := get(s.client, id).Only(ctx)
	if err != nil {
		return nil, err
	}

	var user User
	user.MapFrom(entUser)
	return &user, nil
}

func (s sStorage) List(ctx context.Context) ([]User, error) {
	entUser, err := list(s.client).All(ctx)
	if err != nil {
		return nil, err
	}

	users := make([]User, len(entUser))
	for index, value := range entUser {
		users[index].MapFrom(value)
	}

	return users, nil
}

func (s sStorage) ChangeActive(ctx context.Context, id string) error {
	err := util.ExecTx(ctx, s.client, func(tx *entgen.Tx) error {
		err := activeUser(tx, id).Exec(ctx)
		if err != nil {
			return util.WrapperForDatabaseError("add", err)
		}

		return nil
	})
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (s sStorage) Add(ctx context.Context, user *User) error {
	err := util.ExecTx(ctx, s.client, func(tx *entgen.Tx) error {
		count, err := checkByMail(tx, user.Email).Count(ctx)
		if err != nil {
			return err
		}
		if count != 0 {
			return errors.New("user already exists")
		}

		err = add(tx, *user).Exec(ctx)
		if err != nil {
			return util.WrapperForDatabaseError("add", err)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
