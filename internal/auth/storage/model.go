package storage

import (
	"github.com/go-oauth2/oauth2/v4"
	"inventory/ent/entgen"
	"time"
)

// TokenRequest represents the request body for login with OAuth2.
type TokenRequest struct {
	GrantType    string `json:"grant_type" example:"password" enums:"password,refresh_token"`
	Username     string `json:"username" example:"xyz@yahoo.com"`
	Password     string `json:"password" example:"meee"`
	ClientID     string `json:"client_id" example:"xyz"`
	ClientSecret string `json:"client_secret" example:"xyz"`
	RefreshToken string `json:"refresh_token,omitempty" example:"MDBIYZAZYZKTMDZIMI01OTI4LWFKODYTMGYZYMU2ZJA4ZTZH"`
}

type AuthClientInfo struct {
	AuthClient *entgen.TblAuthClient
}

func NewAuthClientInfo(client *entgen.TblAuthClient) AuthClientInfo {
	return AuthClientInfo{AuthClient: client}
}

func (a AuthClientInfo) GetID() string {
	return a.AuthClient.ID
}

func (a AuthClientInfo) GetSecret() string {
	return a.AuthClient.ClientSecret
}

func (a AuthClientInfo) GetDomain() string {
	return *a.AuthClient.Domain
}

func (a AuthClientInfo) IsPublic() bool {
	return a.AuthClient.Public
}

func (a AuthClientInfo) GetUserID() string {
	return a.AuthClient.CreatedBy
}

type AuthTokenInfo struct {
	Client *entgen.TblAuthToken
}

func NewAuthTokenInfo(client *entgen.TblAuthToken) AuthTokenInfo {
	return AuthTokenInfo{Client: client}
}

func (a AuthTokenInfo) New() oauth2.TokenInfo {
	return a
}

func (a AuthTokenInfo) GetClientID() string {
	return a.Client.Clientid
}

func (a AuthTokenInfo) SetClientID(s string) {
	a.Client.Clientid = s
}

func (a AuthTokenInfo) GetUserID() string {
	return a.Client.UserUlid
}

func (a AuthTokenInfo) SetUserID(s string) {
	a.Client.UserUlid = s
}

func (a AuthTokenInfo) GetRedirectURI() string {
	return *a.Client.RedirectURI
}

func (a AuthTokenInfo) SetRedirectURI(s string) {
	*a.Client.RedirectURI = s
}

func (a AuthTokenInfo) GetScope() string {
	return a.Client.Scope
}

func (a AuthTokenInfo) SetScope(s string) {
	a.Client.Scope = s
}

func (a AuthTokenInfo) GetCode() string {
	return a.Client.Code
}

func (a AuthTokenInfo) SetCode(s string) {
	a.Client.Code = s
}

func (a AuthTokenInfo) GetCodeCreateAt() time.Time {
	return a.Client.Codecreatedat
}

func (a AuthTokenInfo) SetCodeCreateAt(time time.Time) {
	a.Client.Codecreatedat = time
}

func (a AuthTokenInfo) GetCodeExpiresIn() time.Duration {
	return time.Duration(a.Client.Codeexpiresin)
}

func (a AuthTokenInfo) SetCodeExpiresIn(duration time.Duration) {
	a.Client.Codeexpiresin = int(duration)
}

func (a AuthTokenInfo) GetCodeChallenge() string {
	return *a.Client.Codechallenge
}

func (a AuthTokenInfo) SetCodeChallenge(s string) {
	*a.Client.Codechallenge = s
}

func (a AuthTokenInfo) GetCodeChallengeMethod() oauth2.CodeChallengeMethod {
	if *a.Client.Codechallenge != "" {
		// If code challenge exists, check if it was generated with SHA-256 or plain
		return oauth2.CodeChallengeS256
	}
	return oauth2.CodeChallengeS256
}

func (a AuthTokenInfo) SetCodeChallengeMethod(method oauth2.CodeChallengeMethod) {
	methodStr := string(method) // Convert CodeChallengeMethod to string
	a.Client.Codechallenge = &methodStr
}

func (a AuthTokenInfo) GetAccess() string {
	return a.Client.Accesstoken
}

func (a AuthTokenInfo) SetAccess(s string) {
	a.Client.Accesstoken = s
}

func (a AuthTokenInfo) GetAccessCreateAt() time.Time {
	return a.Client.Accesstokencreatedat
}

func (a AuthTokenInfo) SetAccessCreateAt(time time.Time) {
	a.Client.Accesstokencreatedat = time
}

func (a AuthTokenInfo) GetAccessExpiresIn() time.Duration {
	return time.Duration(a.Client.Accesstokenexpiresin)
}

func (a AuthTokenInfo) SetAccessExpiresIn(duration time.Duration) {
	a.Client.Accesstokenexpiresin = int(duration)
}

func (a AuthTokenInfo) GetRefresh() string {
	return a.Client.Refreshtoken
}

func (a AuthTokenInfo) SetRefresh(s string) {
	a.Client.Refreshtoken = s
}

func (a AuthTokenInfo) GetRefreshCreateAt() time.Time {
	return a.Client.Refreshtokencreatedat
}

func (a AuthTokenInfo) SetRefreshCreateAt(time time.Time) {
	a.Client.Refreshtokencreatedat = time
}

func (a AuthTokenInfo) GetRefreshExpiresIn() time.Duration {
	return time.Duration(a.Client.Refreshtokenexpiresin)
}

func (a AuthTokenInfo) SetRefreshExpiresIn(duration time.Duration) {
	a.Client.Refreshtokenexpiresin = int(duration)
}
