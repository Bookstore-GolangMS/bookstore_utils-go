package errors

import (
	"errors"
	"net/http"
)

type RestErr interface {
	Message() string
	Code() int
	Error() string
	Causes() []interface{}
}

type restErr struct {
	message string        `json:"message"`
	code    int           `json:"code"`
	err     string        `json:"error"`
	causes  []interface{} `json:"causes"`
}

func NewRestError(message string, code int, err string, causes []interface{}) RestErr {
	return restErr{
		message: message,
		code:    code,
		err:     err,
		causes:  causes,
	}
}

func (r restErr) Message() string {
	return r.message
}

func (r restErr) Code() int {
	return r.code
}

func (r restErr) Causes() []interface{} {
	return r.causes
}

func (r restErr) Error() string {
	return r.err
}

func NewError(msg string) error {
	return errors.New(msg)
}

func NewBadRequestError(message string) RestErr {
	return restErr{
		message: message,
		code:    http.StatusBadRequest,
		err:     "bad_request",
	}
}

func NewNotFoundError(message string) RestErr {
	return restErr{
		message: message,
		code:    http.StatusNotFound,
		err:     "not_found",
	}
}

func NewInternalServerError(message string, err error) RestErr {
	result := restErr{
		message: message,
		code:    http.StatusInternalServerError,
		err:     "internal_server_error",
	}

	if err != nil {
		result.causes = append(result.causes, err.Error())
	}

	return result
}

func NewUnauthorizedError(message string) RestErr {
	return restErr{
		message: message,
		code:    http.StatusUnauthorized,
		err:     "unauthorized",
	}
}
