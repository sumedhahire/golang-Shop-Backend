package storage

import (
	"context"
	"github.com/go-oauth2/oauth2/v4"
	"inventory/ent/entgen"
	"inventory/ent/entgen/tblauthtoken"
	"inventory/internal/util"
	"time"
)

type TokenStore struct {
	client *entgen.Client
}

func NewTokenStore(client *entgen.Client) TokenStore {
	return TokenStore{
		client: client,
	}
}

func (t TokenStore) Create(ctx context.Context, info oauth2.TokenInfo) error {
	err := util.ExecTx(ctx, t.client, func(tx *entgen.Tx) error {
		_, err := tx.Client().TblAuthToken.Create().SetID(util.GetUuid()).SetAuthUUID(util.GetUuid()).SetAuthXref(util.GetUuid()).
			SetAccesstoken(info.GetAccess()).SetAccesstokencreatedat(info.GetAccessCreateAt()).
			SetAccesstokenexpiresin(int(info.GetAccessExpiresIn())).SetClientid(info.GetClientID()).
			SetUserUlid(info.GetUserID()).SetRefreshtoken(info.GetRefresh()).
			SetRefreshtokencreatedat(info.GetRefreshCreateAt()).SetRefreshtokenexpiresin(int(info.GetRefreshExpiresIn())).
			SetCreatedat(time.Now().UTC()).SetUpdatedat(time.Now().UTC()).SetCode(info.GetCode()).SetCodecreatedat(time.Now().UTC()).
			SetCodeexpiresin(int(info.GetCodeExpiresIn())).SetCodechallenge(info.GetCodeChallenge()).
			Save(ctx)

		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (t TokenStore) RemoveByCode(ctx context.Context, code string) error {
	_, err := t.client.TblAuthToken.Delete().Where(tblauthtoken.Code(code)).Exec(ctx)
	return err
}

func (t TokenStore) RemoveByAccess(ctx context.Context, access string) error {
	_, err := t.client.TblAuthToken.Delete().Where(tblauthtoken.Accesstoken(access)).Exec(ctx)
	return err
}

func (t TokenStore) RemoveByRefresh(ctx context.Context, refresh string) error {
	_, err := t.client.TblAuthToken.Delete().Where(tblauthtoken.Refreshtoken(refresh)).Exec(ctx)
	return err
}

func (t TokenStore) GetByCode(ctx context.Context, code string) (oauth2.TokenInfo, error) {
	token, err := t.client.TblAuthToken.Query().Where(tblauthtoken.Code(code)).First(ctx)
	return NewAuthTokenInfo(token), err
}

func (t TokenStore) GetByAccess(ctx context.Context, access string) (oauth2.TokenInfo, error) {
	token, err := t.client.TblAuthToken.Query().Where(tblauthtoken.Accesstoken(access)).First(ctx)
	return NewAuthTokenInfo(token), err
}

func (t TokenStore) GetByRefresh(ctx context.Context, refresh string) (oauth2.TokenInfo, error) {
	token, err := t.client.TblAuthToken.Query().Where(tblauthtoken.Refreshtoken(refresh)).First(ctx)
	return NewAuthTokenInfo(token), err
}
