package personalia

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	personaliaEntity "github.com/fiber-go-pos-app/internal/entity/personalia"
	customPkg "github.com/fiber-go-pos-app/utils/pkg/custom"

	personaliaUC "github.com/fiber-go-pos-app/internal/usecase/services/personalia"
)

func GetAllUserHandler(ctx *fiber.Ctx) error {
	users, err := personaliaUC.GetAllUser(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return customPkg.BuildJSONRes(ctx, users)
}

func GetAllDTUserHandler(ctx *fiber.Ctx) error {
	users, err := personaliaUC.GetAllUser(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return customPkg.BuildDatatableRes(ctx, int64(len(users)), users)
}

func InsertUserHandler(ctx *fiber.Ctx) error {
	var user personaliaEntity.User

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := validator.New().Struct(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if user.Password != "" {
		user.Password, _ = customPkg.HashPassword(user.Password)
	}

	if err := personaliaUC.InsertUser(ctx, user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendString("Data user berhasil disimpan")
}

func UpdateUserHandler(ctx *fiber.Ctx) error {
	var user personaliaEntity.User
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := validator.New().Struct(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := personaliaUC.UpdateUser(ctx, user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendString("Data user berhasil diubah")
}

func DeleteUserHandler(ctx *fiber.Ctx) error {
	var user personaliaEntity.User

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := validator.New().Struct(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := personaliaUC.DeleteUser(ctx, user.UserID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendString("Data user berhasil dihapus")
}

func UpsertUserHandler(ctx *fiber.Ctx) error {
	var user personaliaEntity.User

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := validator.New().Struct(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := personaliaUC.UpsertUser(ctx, user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return customPkg.BuildJSONRes(ctx, user)
}
