package tag

import (
	"context"
	"inventory/ent/entgen"
)

type IService interface {
	List(ctx context.Context) ([]RSTag, error)
}

type SService struct {
	storage IStorage
}

func NewService(client *entgen.Client) IService {
	return &SService{
		storage: NewStorage(client),
	}
}

func (s *SService) List(ctx context.Context) ([]RSTag, error) {
	tags, err := s.storage.List(ctx)
	if err != nil {
		return nil, err
	}

	rs := make([]RSTag, len(tags))
	for index, value := range tags {
		rs[index].MapFrom(&value)
	}
	return rs, nil
}
