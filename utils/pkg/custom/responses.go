package custom

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func BuildJSONRes(ctx *fiber.Ctx, response any) error {
	err := ctx.JSON(response)
	ctx.Set("content-type", "application/json; charset=utf-8")
	return err
}

func BuildPageAndLimit(ctx *fiber.Ctx) (int, int, error) {
	page, err := strconv.Atoi(ctx.Query("page", "1"))
	if err != nil {
		return 0, 0, fmt.Errorf("page must be number")
	}
	limit, err := strconv.Atoi(ctx.Query("limit", "10"))
	if err != nil {
		return 0, 0, fmt.Errorf("limit must be number")
	}
	return page, limit, nil
}

func BuildOffsetAndLimitES(page int, limit int) (int, int) {
	offset := (page - 1) * limit
	return limit, offset
}

func BuildDatatableRes(ctx *fiber.Ctx, total int64, data any) error {
	err := ctx.JSON(map[string]any{
		"total": total,
		"rows":  data,
	})
	ctx.Set("content-type", "application/json; charset=utf-8")
	return err
}
