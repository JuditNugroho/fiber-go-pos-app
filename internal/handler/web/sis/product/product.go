package product

import (
	"github.com/gofiber/fiber/v2"

	constantsEntity "github.com/fiber-go-pos-app/internal/entity/constants"
)

func WebProductHandler(ctx *fiber.Ctx) error {

	return ctx.Render("sis/pages/product", constantsEntity.WebData{
		Title:        constantsEntity.WebSISProductTitle,
		BaseURL:      constantsEntity.BaseURL,
		StaticUrl:    constantsEntity.StaticUrl,
		CurrentURL:   constantsEntity.WebSISProductURL,
		LinkPageList: constantsEntity.LinkPageList,
	})
}
