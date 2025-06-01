package user

import (
	"context"
	"inventory/ent/entgen"
)

type IUserService interface {
	Get(ctx context.Context, id string) (RSUser, error)
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
