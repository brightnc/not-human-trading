package apperror

import "errors"

// AppError ...
// handle application error
type AppError struct {
	Message string
	Code    ErrorCode
}

// IsAppError ...
// check an error, is that error made by AppError?
func IsAppError(err error) (e AppError, ok bool) {
	ok = errors.As(err, &e)
	return
}

func (e AppError) Error() string {
	return e.Message
}

// NewInternalServerError ...
func NewInternalServerError() AppError {
	return AppError{
		Message: ErrMessageInternalServerError,
		Code:    ErrCodeInternalServerError,
	}
}

// NewServiceUnexpectedError ...
func NewServiceUnexpectedError() AppError {
	return AppError{
		Message: ErrMessageUnexpectedError,
		Code:    ErrCodeUnexpectedError,
	}
}

// NewServiceUnavailableError ...
func NewServiceUnavailableError() AppError {
	return AppError{
		Message: ErrMessageServiceUnavailable,
		Code:    ErrCodeServiceUnavailable,
	}
}

// NewInvalidaParameterError ...
func NewInvalidaParameterError() AppError {
	return AppError{
		Message: ErrMessageInvalidParameter,
		Code:    ErrCodeInvalidParameter,
	}
}
