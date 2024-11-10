package error

import (
	"errors"
	"net/http"
)

type CustomError struct {
	Code int
	Err  error
}

func NotFound(e error) *CustomError {
	return &CustomError{
		Code: http.StatusNotFound,
		Err:  e,
	}
}

func UnprocessableEntity(e error) *CustomError {
	return &CustomError{
		Code: http.StatusUnprocessableEntity,
		Err:  e,
	}
}

func InternalServerError() *CustomError {
	e := errors.New("the server was unable to complete your request. Please try again later")

	return &CustomError{
		Code: http.StatusInternalServerError,
		Err:  e,
	}
}

func (c *CustomError) ErrorMessage() string {
	return c.Err.Error()
}
