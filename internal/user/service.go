package user

import (
	"context"
	"inventory/ent/entgen"
)

type IUserService interface {
	Get(ctx context.Context, id string) (RSUser, error)
	List(ctx context.Context) ([]RSUser, error)
}

type SService struct {
	storage IUserStorage
}

func NewService(client *entgen.Client) IUserService {
	return SService{storage: NewStorage(client)}
}

func (s SService) Get(ctx context.Context, id string) (RSUser, error) {
	user, err := s.storage.Get(ctx, id)
	if err != nil {
		return RSUser{}, err
	}

	var rs RSUser
	rs.MapFrom(user)
	return rs, nil
}

func (s SService) List(ctx context.Context) ([]RSUser, error) {
	user, err := s.storage.List(ctx)
	if err != nil {
		return nil, err
	}

	rs := make([]RSUser, len(user))
	for index, value := range user {
		rs[index].MapFrom(&value)
	}
	return rs, nil
}
