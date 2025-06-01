package storage

import (
	"context"
	"github.com/go-oauth2/oauth2/v4"
	"inventory/ent/entgen"
)

type ClientStore struct {
	client *entgen.Client
}

func NewClientStore(client *entgen.Client) *ClientStore {
	return &ClientStore{
		client: client,
	}
}

func (c ClientStore) GetByID(ctx context.Context, id string) (oauth2.ClientInfo, error) {
	authclient, err := c.client.TblAuthClient.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	client := NewAuthClientInfo(authclient)
	return client, nil
}
