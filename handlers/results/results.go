package results

import "github.com/gofiber/fiber/v2"

const (
	successStatusName = "success"
	errorStatusName   = "error"
)

func Result(status *fiber.Ctx, statusName string, message string, data interface{}) error {
	return status.JSON(
		fiber.Map{
			"status":  statusName,
			"message": message,
			"data":    data,
		},
	)
}

func OkResult(context *fiber.Ctx, message string, data interface{}) error {
	return Result(
		context.Status(fiber.StatusOK),
		successStatusName,
		message,
		data,
	)
}

func BadRequestResult(context *fiber.Ctx, message string, data interface{}) error {
	return Result(
		context.Status(fiber.StatusBadRequest),
		errorStatusName,
		message,
		data,
	)
}

func ServerErrorResult(context *fiber.Ctx, message string, data interface{}) error {
	return Result(
		context.Status(fiber.StatusInternalServerError),
		errorStatusName,
		message,
		data,
	)
}
