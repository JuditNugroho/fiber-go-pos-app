package form

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	formEntity "github.com/fiber-go-pos-app/internal/entity/form"

	formSvc "github.com/fiber-go-pos-app/internal/usecase/services/form"
)

func LoginHandler(ctx *fiber.Ctx) error {
	var loginRequest formEntity.LoginRequest

	if err := ctx.BodyParser(&loginRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := validator.New().Struct(&loginRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := formSvc.LoginForm(ctx, loginRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendString("User berhasil login")
}
