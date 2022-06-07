package common

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetIntParam(ctx *fiber.Ctx, paramName string) (int, error) {
	paramValue := ctx.Params(paramName)

	paramValueInt, err := strconv.Atoi(paramValue)

	return paramValueInt, err
}

func GetUintParam(ctx *fiber.Ctx, paramName string) (uint, error) {
	paramValue, err := GetIntParam(ctx, paramName)
	return uint(paramValue), err
}
