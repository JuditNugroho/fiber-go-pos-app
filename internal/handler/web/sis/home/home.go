package home

import (
	"github.com/gofiber/fiber/v2"

	constantsEntity "github.com/fiber-go-pos-app/internal/entity/constants"
)

func WebSISHomeHandler(ctx *fiber.Ctx) error {

	return ctx.Render("templates/sis/pages/home", constantsEntity.WebData{
		Title:        constantsEntity.WebSISHomeTitle,
		BaseURL:      constantsEntity.BaseURL,
		CurrentURL:   constantsEntity.WebSISHomeURL,
		StaticUrl:    constantsEntity.StaticUrl,
		LinkPageList: constantsEntity.LinkPageList,
	})
}
