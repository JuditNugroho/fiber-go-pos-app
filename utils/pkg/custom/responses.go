package custom

import (
	"github.com/gofiber/fiber/v2"
)

func BuildJSONRes(ctx *fiber.Ctx, response any) error {
	ctx.Set("content-type", "application/json; charset=utf-8")
	return ctx.JSON(response)
}

func BuildDatatableRes(ctx *fiber.Ctx, total int64, data any) error {
	ctx.Set("content-type", "application/json; charset=utf-8")
	return ctx.JSON(map[string]any{
		"total": total,
		"rows":  data,
	})
}
