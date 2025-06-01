package util

import "net/http"

type DatabaseError struct {
	Code    int
	Message string
	Domain  string
	Err     error
}

func (err DatabaseError) Error() string {
	return err.Message
}

func WrapperForDatabaseError(domain string, err error) *DatabaseError {
	return &DatabaseError{
		Code:    http.StatusInternalServerError,
		Message: "Storage error",
		Domain:  domain,
		Err:     err,
	}
}

type CommonError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Domain  string `json:"domain"`
	Err     error  `json:"errId"`
}

func (e *CommonError) Error() string {
	return e.Message
}

func WrapperForCommonError(domain string, err error) error {
	return &CommonError{
		Message: err.Error(),
		Domain:  domain,
		Err:     err,
	}
}

type OauthError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Domain  string `json:"domain"`
	Err     error  `json:"errId"`
}

func (o OauthError) Error() string {
	return o.Message
}

func WrapperForOauthError(domain string, err error) *OauthError {
	return &OauthError{
		Message: err.Error(),
		Domain:  domain,
		Err:     err,
	}

}

type CustomError struct {
	Message string `json:"message"`
	Err     error  `json:"err"`
}

func (e *CustomError) Error() string {
	return e.Message
}
