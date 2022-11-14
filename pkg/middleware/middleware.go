package middleware

import (
	"net/http"

	"github.com/brightnc/not-human-trading/pkg/apperror"
	"github.com/brightnc/not-human-trading/pkg/appresponse"
	"github.com/brightnc/not-human-trading/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		var (
			debug    = err.Error()
			code     = http.StatusInternalServerError
			response = appresponse.Error(http.StatusInternalServerError, string(apperror.Message[apperror.ErrInternalServerCode]), &debug)
		)

		e, ok := apperror.IsAppError(err)
		if ok {
			debug = e.Debug
			code, ok = appresponse.HTTPErrorStatus[e.Code]
			if !ok {
				code = findHttpStatusCodeFormErrorCode(e.Code)
			}
			response = appresponse.Error(int(e.Code), string(e.Message), &debug)
		}

		logger.ErrorWithContext(ctx.UserContext(), debug)

		return ctx.Status(code).JSON(response)
	}
}
