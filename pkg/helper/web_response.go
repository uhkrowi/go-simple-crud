package helper

import (
	"github.com/gofiber/fiber/v2"
)

type WebResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    any         `json:"meta"`
}

func WebResponseOK(ctx *fiber.Ctx, response interface{}) error {
	return ctx.Status(fiber.StatusOK).JSON(WebResponse{
		Code:    fiber.StatusOK,
		Message: "OK",
		Data:    response,
	})
}

func WebResponseError(ctx *fiber.Ctx, statusCode int, err string) error {
	return ctx.Status(statusCode).JSON(WebResponse{
		Code:    statusCode,
		Message: err,
	})
}

func GetWebResponse(ctx *fiber.Ctx, result interface{}, err error) error {
	if err != nil {
		errorStatusCode := getErrorStatusCode(err.Error())

		return ctx.Status(errorStatusCode).JSON(WebResponse{
			Code:    errorStatusCode,
			Message: err.Error(),
		})
	} else {
		return WebResponseOK(ctx, result)
	}
}
