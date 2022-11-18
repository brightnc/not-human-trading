package httphdl

import (
	"net/http"

	"github.com/brightnc/not-human-trading/internal/core/port"
	"github.com/brightnc/not-human-trading/pkg/apperror"
	"github.com/brightnc/not-human-trading/pkg/appresponse"
	"github.com/brightnc/not-human-trading/pkg/validators"
	"github.com/gofiber/fiber/v2"
)

type HTTPHandler struct {
	svc       port.Service
	validator validators.Validator
}

func NewHTTPHandler(svc port.Service, validators validators.Validator) *HTTPHandler {
	return &HTTPHandler{
		svc:       svc,
		validator: validators,
	}
}

/*
	|--------------------------------------------------------------------------
	| The Handler Adaptor
	|--------------------------------------------------------------------------
	|
	| An Adapter will initiate the interaction with the Application through
	| a Port, using specific technology that means you can choose
	| any technology you want for your application protocol.
	|
*/

func (hdl *HTTPHandler) validateBody(body interface{}) error {
	err := hdl.validator.ValidateStruct(body)
	if err != nil {
		_, ok := apperror.IsAppError(err)
		if ok {
			return err
		}
		return apperror.NewInvalidaParameterError()
	}
	return nil
}

func (hdl *HTTPHandler) responseError(err error, c *fiber.Ctx) error {
	if e, ok := apperror.IsAppError(err); ok {
		return c.Status(appresponse.HTTPErrorStatus[e.Code]).JSON(appresponse.Error(string(e.Message)))
	}
	return c.Status(http.StatusInternalServerError).JSON(appresponse.Error(err.Error()))
}
