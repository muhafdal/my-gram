package errs

import "net/http"

type Error interface {
	Status() int
	Message() string
	Error() string
}

type ErrorData struct {
	ErrStatus  int    `json:"errStatus"`
	ErrMessage string `json:"errMessage"`
	ErrErrors  string `json:"errErrors"`
}

func (e *ErrorData) Error() string {
	return e.ErrErrors
}

func (e *ErrorData) Message() string {
	return e.ErrMessage
}

func (e *ErrorData) Status() int {
	return e.ErrStatus
}

func NewUnathorizedError(message string) Error {
	return &ErrorData{
		ErrStatus:  http.StatusForbidden,
		ErrMessage: message,
		ErrErrors:  "FORBIDDEN_ACCESS",
	}
}

func NewUnauthenticatedError(message string) Error {
	return &ErrorData{
		ErrStatus:  http.StatusUnauthorized,
		ErrMessage: message,
		ErrErrors:  "UNAUTHORIZED",
	}
}

func NewBadRequestError(message string) Error {
	return &ErrorData{
		ErrStatus:  http.StatusBadRequest,
		ErrMessage: message,
		ErrErrors:  "BAD_REQUEST",
	}
}

func NewNotFoundError(message string) Error {
	return &ErrorData{
		ErrStatus:  http.StatusNotFound,
		ErrMessage: message,
		ErrErrors:  "NOT_FOUND",
	}
}

func NewUnprocessableEntityError(message string) Error {
	return &ErrorData{
		ErrStatus:  http.StatusUnprocessableEntity,
		ErrMessage: message,
		ErrErrors:  "UNPROCESSABLE_ENTITY",
	}
}

func NewInternalServerError(message string) Error {
	return &ErrorData{
		ErrStatus:  http.StatusInternalServerError,
		ErrMessage: message,
		ErrErrors:  "INTERNAL_SERVER_ERROR",
	}
}

func NewConflictError(message string) Error {
	return &ErrorData{
		ErrStatus:  http.StatusConflict,
		ErrMessage: message,
		ErrErrors:  "CONFLICT_ERROR",
	}
}
