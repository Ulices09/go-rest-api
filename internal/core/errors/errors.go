package errors

import (
	"net/http"
)

type CustomError struct {
	Message string `json:"message"`
	Status  int    `json:"-"`
}

func (s CustomError) Error() string {
	return s.Message
}

func NewBadRequestError(message ...interface{}) CustomError {
	return CustomError{
		Message: getMessage("Bad Request", message...),
		Status:  http.StatusBadRequest,
	}
}

func NewNotFoundError(message ...interface{}) CustomError {
	return CustomError{
		Message: getMessage("Not Found", message...),
		Status:  http.StatusNotFound,
	}
}

func NewConflictError(message ...interface{}) CustomError {
	return CustomError{
		Message: getMessage("Conflict", message...),
		Status:  http.StatusConflict,
	}
}

func NewUnauthorizedError(message ...interface{}) CustomError {
	return CustomError{
		Message: getMessage("Unauthorized", message...),
		Status:  http.StatusUnauthorized,
	}
}

func NewForbiddenError(message ...interface{}) CustomError {
	return CustomError{
		Message: getMessage("Forbidden", message...),
		Status:  http.StatusForbidden,
	}
}

func NewInternalServerError(message ...interface{}) CustomError {
	return CustomError{
		Message: getMessage("Internal Server Error", message...),
		Status:  http.StatusInternalServerError,
	}
}

func getMessage(defaultMessage string, message ...interface{}) string {
	if len(message) > 0 {
		if m, ok := message[0].(string); ok {
			return m
		}
	}

	return defaultMessage
}
