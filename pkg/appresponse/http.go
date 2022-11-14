package appresponse

import (
	"kkp-api/pkg/apperror"
	"net/http"
)

var (
	// HTTPErrorStatus ... an error dictionary for protocol HTTP, which can be accessed by using error code
	// from package apperror
	HTTPErrorStatus = map[apperror.ErrorCode]int{
		apperror.InternalServerError:   http.StatusInternalServerError,
		apperror.DataNotFoundError:     http.StatusNotFound,
		apperror.UnexpectedError:       http.StatusInternalServerError,
		apperror.LookupStatusError:     http.StatusBadRequest,
		apperror.ConfirmStatusError:    http.StatusBadRequest,
		apperror.InvalidFaultCodeError: http.StatusBadRequest,
		apperror.UnheathyError:         http.StatusInternalServerError,
	}
)
