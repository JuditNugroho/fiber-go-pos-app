package product

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	productEntity "github.com/fiber-go-pos-app/internal/entity/product"
	customPkg "github.com/fiber-go-pos-app/utils/pkg/custom"

	productUC "github.com/fiber-go-pos-app/internal/usecase/services/product"
)

// GetDTProductHandler : Get List Of Product for Datatable
func GetDTProductHandler(ctx *fiber.Ctx) error {
	page, limit, err := customPkg.BuildPageAndLimit(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	search := ctx.Query("search", "")

	result, err := productUC.GetDTAllProduct(ctx, page, limit, search)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return customPkg.BuildDatatableRes(ctx, result.Total, result.Data)
}

// GetProductByIDOrBarcode : Get product data from ID or barcode
func GetProductByIDOrBarcode(ctx *fiber.Ctx) error {
	search := ctx.Query("search", "")
	if search == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Inputan tidak boleh kosong!",
		})
	}

	result, err := productUC.GetProductByIDOrBarcode(ctx, search)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return customPkg.BuildJSONRes(ctx, result)
}

func InsertProductHandler(ctx *fiber.Ctx) error {
	var product productEntity.Product

	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := validator.New().Struct(&product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := productUC.InsertProduct(ctx, product); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendString("Data product berhasil disimpan")
}

func UpdateProductHandler(ctx *fiber.Ctx) error {
	var product productEntity.Product

	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := validator.New().Struct(&product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := productUC.UpdateProduct(ctx, product); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendString("Data product berhasil diubah")
}

func DeleteProductHandler(ctx *fiber.Ctx) error {
	var product productEntity.Product

	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := validator.New().Struct(&product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := productUC.DeleteProduct(ctx, product.ProductID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendString("Data Product berhasil dihapus")
}

func UpsertProductHandler(ctx *fiber.Ctx) error {
	var product productEntity.Product

	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := validator.New().Struct(&product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := productUC.UpsertProduct(ctx, product); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return customPkg.BuildJSONRes(ctx, product)
}
