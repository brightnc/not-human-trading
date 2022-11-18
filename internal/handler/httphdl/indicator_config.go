package httphdl

import (
	"net/http"

	"github.com/brightnc/not-human-trading/internal/core/port"
	"github.com/brightnc/not-human-trading/internal/handler/httphdl/dto"
	"github.com/brightnc/not-human-trading/pkg/apperror"
	"github.com/brightnc/not-human-trading/pkg/appresponse"
	"github.com/brightnc/not-human-trading/pkg/logger"
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

func (hdl *HTTPHandler) UpdateIndicator(c *fiber.Ctx) error {
	var request dto.IndicatorConfigRequest

	/*
		|--------------------------------------------------------------------------
		| Request Body Serialization
		|--------------------------------------------------------------------------
		|
		| Here you can parse the body from incoming request into the structure/model
		| to use in your application.
		|
	*/

	err := c.BodyParser(&request)
	if err != nil {
		logger.Error("cannot parese request from payload")
		return c.Status(fiber.StatusBadRequest).JSON(apperror.NewInvalidaParameterError())
	}

	/*
		|--------------------------------------------------------------------------
		| Data Validation
		|--------------------------------------------------------------------------
		|
		| Here you may specify which part of the incoming request body you want to validate
		| before putting them into the business logic.
		|
	*/
	err = hdl.validateBody(&request)
	if err != nil {
		logger.Error("invalid paramter from payload")
		return c.Status(fiber.StatusBadRequest).JSON(apperror.NewInvalidaParameterError())
	}

	err = hdl.svc.UpdateIndicator(request.ToIndicatorConfigDomain())
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.SendStatus(fiber.StatusOK)
}

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

func (hdl *HTTPHandler) UpdateOrder(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func (hdl *HTTPHandler) UpdateBotConfig(c *fiber.Ctx) error {
	var request dto.BotConfigRequest

	err := c.BodyParser(&request)
	if err != nil {
		logger.Error("cannot parese request from payload")
		return c.Status(fiber.StatusBadRequest).JSON(apperror.NewInvalidaParameterError())
	}

	err = hdl.validateBody(&request)
	if err != nil {
		logger.Error("invalid paramter from payload")
		return c.Status(fiber.StatusBadRequest).JSON(apperror.NewInvalidaParameterError())
	}

	err = hdl.svc.UpdateBotConfig(request.ToBotConfigDomain())
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.SendStatus(fiber.StatusOK)
}
