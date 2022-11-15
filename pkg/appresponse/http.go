package appresponse

import (
	"net/http"

	"github.com/brightnc/not-human-trading/pkg/apperror"
)

var (
	// HTTPErrorStatus ... an error dictionary for protocol HTTP, which can be accessed by using error code
	// from package apperror
	HTTPErrorStatus = map[apperror.ErrorCode]int{
		apperror.ErrCodeInternalServerError: http.StatusInternalServerError,
		apperror.ErrCodeInvalidParameter:    http.StatusBadRequest,
		apperror.ErrCodeServiceUnavailable:  http.StatusServiceUnavailable,
		apperror.ErrCodeUnexpectedError:     http.StatusInternalServerError,
	}
)
