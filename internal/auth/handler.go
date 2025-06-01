package auth

import (
	"context"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/labstack/echo/v4"
	oauthConfig "inventory/config/oauthCOnfig"
	"inventory/ent/entgen"
	"inventory/internal/util"
	"net/http"
	"strings"
)

type (
	IAuthHandler interface {
		Login(c echo.Context) error
		Logout(c echo.Context) error
	}

	sAuthHandler struct {
		Srv *server.Server
	}
)

func NewAuthHandler(client *entgen.Client) IAuthHandler {
	return sAuthHandler{Srv: oauthConfig.GetOauthInstance(client)}
}

func (s sAuthHandler) Login(c echo.Context) error {
	data, err := handleToken(s.Srv, c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, util.ConvertToResponse(data))
}

func handleToken(srv *server.Server, e echo.Context) (map[string]interface{}, error) {
	grantType, tokenRq, err := srv.ValidationTokenRequest(e.Request())
	if err != nil {
		return nil, util.WrapperForOauthError("handle token", err)
	}

	ctx := context.Background()
	tokenInfo, err := srv.GetAccessToken(ctx, grantType, tokenRq)
	if err != nil {
		return nil, util.WrapperForOauthError("handle token", err)
	}

	return srv.GetTokenData(tokenInfo), nil

}

func (s sAuthHandler) Logout(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "missing token"})
	}

	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix(token, "Bearer ")
	}

	err := s.Srv.Manager.RemoveAccessToken(context.Background(), token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to logout"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "logged out"})
}
