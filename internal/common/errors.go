package common

var (
	successCode           = 2000
	badRequestCode        = 4000
	internalErrorCode     = 5000
	notFoundDataErrorCode = 5001
	bizErrorCode          = 5002
	unknownErrorCode      = 5003
)

type SystemError struct {
	Code    int
	Message string
}

func (e SystemError) Error() string {
	return e.Message
}

func NewInternalError(msg string) error {
	return &SystemError{Code: internalErrorCode, Message: msg}
}

func NewNotFoundDataError(msg string) error {
	return &SystemError{Code: notFoundDataErrorCode, Message: msg}
}

func NewBizError(msg string) error {
	return &SystemError{Code: bizErrorCode, Message: msg}
}

func NewBarRequestError(msg string) error {
	return &SystemError{Code: badRequestCode, Message: msg}
}
