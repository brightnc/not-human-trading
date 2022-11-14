package apperror

// ErrorCode ...
type ErrorCode int

const (
	// ErrCodeInternalServerError ...
	ErrCodeInternalServerError ErrorCode = 5000
	// ErrCodeUnexpectedError ...
	ErrCodeUnexpectedError ErrorCode = 5001
	// ErrCodeServiceUnavailable ...
	ErrCodeServiceUnavailable ErrorCode = 5003

	// ErrCodeInvalidParameter ...
	ErrCodeInvalidParameter ErrorCode = 4000
)
