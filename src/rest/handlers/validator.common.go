package handlers

import (
	"github.com/gofiber/fiber/v3"
	"goboilerplate.com/src/pkg/utils"
)

func ValidateStruct(ctx fiber.Ctx, s interface{}) bool {
	if err := ctx.Bind().Body(s); err != nil {
		_ = ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
		})
		return false
	}

	validator := utils.GetValidator()
	errors := validator.ValidateStruct(s)
	if len(errors) > 0 {
		_ = ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Validation failed",
			"errors":  errors,
		})
		return false
	}
	return true
}
