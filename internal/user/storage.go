package user

import (
	"context"
	"inventory/ent/entgen"
)

type IUserStorage interface {
	Get(ctx context.Context, id string) (*User, error)
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
