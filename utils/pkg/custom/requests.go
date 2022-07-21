package custom

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

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
