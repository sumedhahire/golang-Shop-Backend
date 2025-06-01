package tag

import (
	"context"
	"inventory/ent/entgen"
)

type IStorage interface {
	List(ctx context.Context) ([]Tag, error)
}

type sStorage struct {
	client *entgen.Client
}

func NewStorage(client *entgen.Client) IStorage {
	return sStorage{client: client}
}

func (s sStorage) List(ctx context.Context) ([]Tag, error) {
	entTags, err := list(s.client).All(ctx)
	if err != nil {
		return nil, err
	}

	tags := make([]Tag, len(entTags))
	for index, value := range entTags {
		tags[index].MapFromEnt(value)
	}
	return tags, nil
}
