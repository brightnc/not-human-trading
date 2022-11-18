package httphdl

import (
	"github.com/brightnc/not-human-trading/internal/handler/httphdl/dto"
	"github.com/brightnc/not-human-trading/pkg/apperror"
	"github.com/brightnc/not-human-trading/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func (hdl *HTTPHandler) UpdateOrder(c *fiber.Ctx) error {
	var request dto.BotOrderRequest

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

	err = hdl.svc.UpdateBotOrder(request.ToBotOrderDomain())
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.SendStatus(fiber.StatusOK)
}
