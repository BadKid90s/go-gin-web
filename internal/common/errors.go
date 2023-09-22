package common

import "errors"

// common errors

var (
	ErrSuccess             = newError(0, "ok")
	ErrBadRequest          = newError(400, "Bad Request")
	ErrUnauthorized        = newError(401, "Unauthorized")
	ErrNotFound            = newError(404, "Not Found")
	ErrInternalServerError = newError(500, "Internal Server Error")
)

type Error struct {
	Code    int
	Message string
}

var errorCodeMap = map[error]int{}

func NewInternalError(msg string) error {
	err := errors.New(msg)
	errorCodeMap[err] = 500
	return err
}

func newError(code int, msg string) error {
	err := errors.New(msg)
	errorCodeMap[err] = code
	return err
}
func (e Error) Error() string {
	return e.Message
}
