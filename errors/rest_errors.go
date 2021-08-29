package errors

import (
	"errors"
	"fmt"
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
	error   string        `json:"errors"`
	causes  []interface{} `json:"causes"`
}

func NewRestError(message string, code int, err string, causes []interface{}) RestErr {
	return restErr{
		message: message,
		code:    code,
		error:   err,
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
	return fmt.Sprintf(
		"message: %s - code: %d - error: %s - causes: [ %v ]", r.message, r.code, r.error, r.causes)
}

func NewError(msg string) error {
	return errors.New(msg)
}

func NewBadRequestError(message string) RestErr {
	return restErr{
		message: message,
		code:    http.StatusBadRequest,
		error:   "bad_request",
	}
}

func NewNotFoundError(message string) RestErr {
	return restErr{
		message: message,
		code:    http.StatusNotFound,
		error:   "not_found",
	}
}

func NewInternalServerError(message string, err error) RestErr {
	result := restErr{
		message: message,
		code:    http.StatusInternalServerError,
		error:   "internal_server_error",
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
		error:   "unauthorized",
	}
}
