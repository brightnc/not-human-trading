package appsuccess

import "fmt"

type AppSuccess struct {
	Code    string
	Message string
}

const (
	svcPrefix string = "NHT"
)

func code(code successCode) string {
	return fmt.Sprintf("%s-%s", svcPrefix, code)
}

// NewInternalServerError ...
func OK() AppSuccess {
	return AppSuccess{
		Code:    code(ok),
		Message: ErrMessageInternalServerError,
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
