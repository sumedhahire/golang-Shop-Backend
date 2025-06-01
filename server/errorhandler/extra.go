package errorhandler

import (
	"database/sql"
	"errors"
	"fmt"
	oauthErr "github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"inventory/internal/util"
	"net/http"
)

func HandleDatabaseError(err *util.DatabaseError) BaseErr {
	var baseErr BaseErr
	log.Error(err.Domain, err.Error())
	if errors.Is(err.Err, sql.ErrNoRows) {
		baseErr.Err = err
		baseErr.Message = "Not found"
		baseErr.Code = http.StatusNotFound
		baseErr.CodeText = http.StatusText(http.StatusNotFound)
	}

	switch errType := err.Err.(type) {

	//case *mysql.MySQLError:
	//	return handleMySqlError(errType)
	//case *pq.Error:
	//	return handlePostgressError(errType)
	default:
		BaseErrResponse(errType)
	}

	return baseErr
}

func handleCommonError(err *util.CommonError) BaseErr {
	var baseErr BaseErr
	log.Error(err.Domain, err.Error())
	switch errType := err.Err.(type) {
	case validator.ValidationErrors:
		return handleValidationError(errType)
	case *echo.HTTPError:
		fmt.Println("hr;fggr")
		return handleEchoHttpError(errType)
	default:
		BaseErrResponse(errType)
	}

	return baseErr
}

func handleEchoHttpError(err *echo.HTTPError) BaseErr {
	var baseErr BaseErr
	baseErr.Err = err
	baseErr.Code = uint64(err.Code)
	baseErr.CodeText = http.StatusText(err.Code)
	if str, ok := err.Message.(string); ok {
		baseErr.Message = str

	} else {
		baseErr.Message = err.Error()

	}
	return baseErr
}

func handleValidationError(err validator.ValidationErrors) BaseErr {
	var baseErr BaseErr
	baseErr.Code = http.StatusBadRequest
	baseErr.CodeText = http.StatusText(http.StatusBadRequest)
	baseErr.Message = "validation error"
	baseErr.Err = err

	var validationErr = make([]ValidationErr, 0)
	for _, rErr := range err {
		var validErr ValidationErr
		switch rErr.Tag() {
		case "required":
			validErr.SetValidation(rErr.Field(), "required")
		case "email":
			validErr.SetValidation(rErr.Field(), "should follow mail structure")
		case "gte":
			validErr.SetValidation(rErr.Field(), "is less")
		default:
			validErr.SetValidation(rErr.Field(), rErr.Tag())
		}
		validationErr = append(validationErr, validErr)
	}
	baseErr.Validation = validationErr

	return baseErr
}

func handleOauthError(err *util.OauthError) BaseErr {
	var baseErr BaseErr

	switch err.Err { // Directly switch on the error value
	case oauthErr.ErrAccessDenied:
		baseErr.Err = err
		baseErr.Message = "Access denied"
		baseErr.Code = http.StatusForbidden
		baseErr.CodeText = http.StatusText(http.StatusForbidden)
	case oauthErr.ErrExpiredRefreshToken:
		baseErr.Err = err
		baseErr.Message = "Refresh token expired"
		baseErr.Code = http.StatusBadRequest
		baseErr.CodeText = http.StatusText(http.StatusBadRequest)
	case oauthErr.ErrInvalidGrant:
		baseErr.Err = err
		baseErr.Message = "Invalid grant"
		baseErr.Code = http.StatusForbidden
		baseErr.CodeText = http.StatusText(http.StatusForbidden)
	case oauthErr.ErrInvalidClient:
		baseErr.Err = err
		baseErr.Message = "Invalid client"
		baseErr.Code = http.StatusForbidden
		baseErr.CodeText = http.StatusText(http.StatusForbidden)
	case oauthErr.ErrInvalidAccessToken:
		baseErr.Err = err
		baseErr.Message = "Invalid access token"
		baseErr.Code = http.StatusUnauthorized
		baseErr.CodeText = http.StatusText(http.StatusUnauthorized)
	case oauthErr.ErrInvalidRefreshToken:
		baseErr.Err = err
		baseErr.Message = "Invalid refresh token"
		baseErr.Code = http.StatusUnauthorized
		baseErr.CodeText = http.StatusText(http.StatusUnauthorized)
	case oauthErr.ErrUnsupportedGrantType:
		baseErr.Err = err
		baseErr.Message = "Unsupported grant type"
		baseErr.Code = http.StatusBadRequest
		baseErr.CodeText = http.StatusText(http.StatusBadRequest)
	default: // Handle other errors, including database errors
		//if mysqlErr, ok := err.Err.(*mysql.MySQLError); ok {
		//	return handleMySqlError(mysqlErr)
		//}
		//if pqErr, ok := err.Err.(*pq.Error); ok {
		//	return handlePostgressError(pqErr)
		//}
		// If it's neither a MySQL nor Postgres error, and not one of the OAuth errors,
		// you might want a default handling here:
		baseErr.Err = err
		baseErr.Message = err.Error()                 // Or a more generic message
		baseErr.Code = http.StatusInternalServerError // Or appropriate code
		baseErr.CodeText = http.StatusText(http.StatusInternalServerError)
	}

	return baseErr
}
