package errors

import "net/http"

type CustomError struct {
	Message string `json:"message"`
	Status  int    `json:"-"`
}

func (s CustomError) Error() string {
	return s.Message
}

func NewNotFoundError(message string) CustomError {
	return CustomError{
		Message: message,
		Status:  http.StatusNotFound,
	}
}

func NewBadRequestError(message string) CustomError {
	return CustomError{
		Message: message,
		Status:  http.StatusBadRequest,
	}
}

func NewConflictError(message string) CustomError {
	return CustomError{
		Message: message,
		Status:  http.StatusConflict,
	}
}

func NewUnauthorizedError(message string) CustomError {
	return CustomError{
		Message: message,
		Status:  http.StatusUnauthorized,
	}
}

func NewForbiddenError(message string) CustomError {
	return CustomError{
		Message: message,
		Status:  http.StatusForbidden,
	}
}

func NewInternalServerError(message string) CustomError {
	return CustomError{
		Message: message,
		Status:  http.StatusInternalServerError,
	}
}
