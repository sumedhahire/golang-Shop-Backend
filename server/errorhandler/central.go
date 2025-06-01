package errorhandler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"inventory/internal/util"
	"net/http"
)

func HttpEchoCustomError(err error, e echo.Context) {
	var baseErr BaseErr
	log.Error(err.Error())

	baseErr.Code = http.StatusInternalServerError
	baseErr.CodeText = http.StatusText(http.StatusInternalServerError)
	baseErr.Message = "internal server"
	baseErr.Err = err
	if echoErr, ok := err.(*echo.HTTPError); ok {
		baseErr = BaseErr{
			Code:       uint64(echoErr.Code),
			CodeText:   http.StatusText(echoErr.Code),
			Message:    http.StatusText(echoErr.Code),
			Validation: nil,
			Err:        echoErr,
		}
	}
	switch errType := err.(type) {

	case *util.DatabaseError:
		baseErr = HandleDatabaseError(errType)

	case *util.CommonError:

		baseErr = handleCommonError(errType)

	case *util.OauthError:
		baseErr = handleOauthError(errType)

	default:
		baseErr = BaseErrResponse(errType)

	}
	errX := e.JSON(int(baseErr.Code), ConvertToRSErr(baseErr))
	if errX != nil {
		log.Error(errX)
	}
}
