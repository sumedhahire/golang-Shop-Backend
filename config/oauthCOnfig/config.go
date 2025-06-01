package oauthConfig

import (
	"context"
	"fmt"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"inventory/config/dbConfig"
	"inventory/ent/entgen"
	"inventory/ent/entgen/tbluser"
	"inventory/internal/auth/storage"
	"inventory/internal/util"
	"sync"
	"time"
)

var (
	oauth2 *server.Server
	once   sync.Once
)

func userAuth(ctx context.Context, clientID, username, password string) (userId string, err error) {
	userObj, err := getUser(ctx, username)
	if err != nil {
		return "", util.WrapperForDatabaseError("user auth", err)
	}

	if userObj.Password != password {
		return "", util.WrapperForCommonError("user auth", fmt.Errorf("invalid password"))
	}

	return userObj.ID, nil
}

func getUser(ctx context.Context, email string) (*entgen.TblUser, error) {
	db := dbConfig.InitDB()

	userData, err := db.TblUser.Query().
		Where(
			tbluser.EmailEQ(email),
			tbluser.DeletedAtIsNil(),
		).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	fmt.Println(userData)
	return userData, nil
}

func newOauth(client *entgen.Client) {
	manager := manage.NewDefaultManager()

	cfg := manage.RefreshingConfig{
		AccessTokenExp:     time.Minute * 60,
		RefreshTokenExp:    time.Hour * 24,
		IsGenerateRefresh:  true,
		IsResetRefreshTime: false,
		IsRemoveAccess:     false,
		IsRemoveRefreshing: false,
	}

	manager.SetRefreshTokenCfg(&cfg)

	clientObj := storage.NewClientStore(client)
	storeObj := storage.NewTokenStore(client)

	manager.MapClientStorage(clientObj)
	manager.MapTokenStorage(storeObj)

	//set up auth
	oauth2 = server.NewServer(server.NewConfig(), manager)
	//get requests for the /storage
	oauth2.SetAllowGetAccessRequest(true)
	//get client info,extract client id and secret
	oauth2.SetClientInfoHandler(server.ClientFormHandler)

	//authentication func
	oauth2.SetPasswordAuthorizationHandler(userAuth)

}

func GetOauthInstance(client *entgen.Client) *server.Server {
	once.Do(func() {
		newOauth(client)
	})
	return oauth2
}
