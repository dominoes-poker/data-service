package results

import (
	"github.com/gofiber/fiber/v2"
)

const (
	successStatusName = "success"
	errorStatusName   = "error"
)

func Result(status *fiber.Ctx, statusName string, data interface{}) error {
	return status.JSON(
		fiber.Map{
			"status": statusName,
			"data":   data,
		},
	)
}

func OkResult(context *fiber.Ctx, data interface{}) error {
	return Result(
		context.Status(fiber.StatusOK),
		successStatusName,
		data,
	)
}

func BadRequestResult(context *fiber.Ctx, data interface{}) error {
	return Result(
		context.Status(fiber.StatusBadRequest),
		errorStatusName,
		data,
	)
}

func NotFoundResult(context *fiber.Ctx, data interface{}) error {
	return Result(
		context.Status(fiber.StatusNotFound),
		errorStatusName,
		data,
	)
}

func ServerErrorResult(context *fiber.Ctx, data interface{}) error {
	return Result(
		context.Status(fiber.StatusInternalServerError),
		errorStatusName,
		data,
	)
}
