package home

import (
	"github.com/gofiber/fiber/v2"

	constantsEntity "github.com/fiber-go-pos-app/internal/entity/constants"
)

func WebSISHomeHandler(ctx *fiber.Ctx) error {

	return ctx.Render("sis/pages/home", constantsEntity.WebData{
		Title:        constantsEntity.WebSISHomeTitle,
		BaseURL:      constantsEntity.BaseURL,
		CurrentURL:   constantsEntity.WebSISHomeURL,
		TemplateURL:  constantsEntity.TemplateUrl,
		LinkPageList: constantsEntity.LinkPageList,
	})
}
