package middleware

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"inventory/config"
	"inventory/config/dbConfig"
	oauthConfig "inventory/config/oauthCOnfig"
	"inventory/internal/user"
	"inventory/internal/util"
)

func OauthValidationUser() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(e echo.Context) error {
			srv := oauthConfig.GetOauthInstance(dbConfig.InitDB())
			if srv == nil {
				panic("server not here")
			}
			token, err := srv.ValidationBearerToken(e.Request())
			if err != nil {
				return util.WrapperForOauthError("unauthorized", err)
			}

			userInfo, err := getUserDetails(token.GetUserID())
			if err != nil {
				return util.WrapperForOauthError("unauthorized", err)
			}

			if userInfo.Role != "user" && userInfo.Role != "admin" {
				return util.WrapperForOauthError("unauthorized", fmt.Errorf("not  priveliage"))
			}

			e.Set("token", token)
			e.Set("userId", token.GetUserID())

			return next(e)
		}
	}
}

func OauthValidationAdmin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(e echo.Context) error {
			srv := oauthConfig.GetOauthInstance(dbConfig.InitDB())
			if srv == nil {
				panic("server not here")
			}
			token, err := srv.ValidationBearerToken(e.Request())
			if err != nil {
				return util.WrapperForOauthError("unauthorized", err)
			}

			userInfo, err := getUserDetails(token.GetUserID())
			if err != nil {
				return util.WrapperForOauthError("unauthorized", err)
			}

			if userInfo.Role != "admin" {
				return util.WrapperForOauthError("unauthorized", fmt.Errorf("not admin priveliage"))
			}

			e.Set("token", token)
			e.Set("userId", token.GetUserID())

			return next(e)
		}
	}
}

func getUserDetails(userId string) (user.RSUser, error) {
	ctx := context.Background()
	entClient := config.InitServices().Client

	service := user.NewService(entClient)
	userInfo, err := service.Get(ctx, userId)
	if err != nil {
		return user.RSUser{}, err
	}

	return userInfo, err

}
