package errorhandler

import (
	"fmt"
	"inventory/internal/common/response"
	"net/http"
)

type BaseErr struct {
	Code       uint64          `json:"code"`
	CodeText   string          `json:"codeText"`
	Message    string          `json:"message"`
	Validation []ValidationErr `json:"validationErrors"`
	Err        interface{}     `json:"-"`
}

func (e *BaseErr) Error() string {
	return fmt.Sprintf("code:%v message:%s", e.Code, e.Message)
}

type ValidationErr struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}

func (v *ValidationErr) SetValidation(field, reason string) {
	v.Field = field
	v.Reason = reason
}

func ConvertToRSErr(baseErr BaseErr) response.BaseRS {

	return response.BaseRS{
		APIVersion: "",
		Data:       nil,
		Error:      baseErr,
	}

}

func BaseErrResponse(err error) BaseErr {
	return BaseErr{
		Code:     http.StatusInternalServerError,
		CodeText: http.StatusText(http.StatusInternalServerError),
		Message:  "Internal server error",
		Err:      err.Error(),
	}
}
