package personalia

import (
	"github.com/fiber-go-pos-app/utils/pkg/custom"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	personaliaEntity "github.com/fiber-go-pos-app/internal/entity/personalia"

	personaliaUC "github.com/fiber-go-pos-app/internal/usecase/services/personalia"
)

func GetAllMemberHandler(ctx *fiber.Ctx) error {
	members, err := personaliaUC.GetAllMember(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return custom.BuildJSONRes(ctx, members)
}

func GetAllDTMemberHandler(ctx *fiber.Ctx) error {
	members, err := personaliaUC.GetAllMember(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return custom.BuildDatatableRes(ctx, int64(len(members)), members)
}

func InsertMemberHandler(ctx *fiber.Ctx) error {
	var member personaliaEntity.Member

	if err := ctx.BodyParser(&member); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := validator.New().Struct(&member); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := personaliaUC.InsertMember(ctx, member); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendString("Data member berhasil disimpan")
}

func UpdateMemberHandler(ctx *fiber.Ctx) error {
	var member personaliaEntity.Member

	if err := ctx.BodyParser(&member); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := validator.New().Struct(&member); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := personaliaUC.UpdateMember(ctx, member); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendString("Data member berhasil diubah")
}

func DeleteMemberHandler(ctx *fiber.Ctx) error {
	var member personaliaEntity.Member

	if err := ctx.BodyParser(&member); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := validator.New().Struct(&member); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := personaliaUC.DeleteMember(ctx, member.ID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.SendString("Data Member berhasil dihapus")
}

func UpsertMemberHandler(ctx *fiber.Ctx) error {
	var member personaliaEntity.Member

	if err := ctx.BodyParser(&member); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := validator.New().Struct(&member); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := personaliaUC.UpsertMember(ctx, member); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return custom.BuildJSONRes(ctx, member)
}
