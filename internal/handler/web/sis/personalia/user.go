package personalia

import (
	"github.com/gofiber/fiber/v2"

	constantsEntity "github.com/fiber-go-pos-app/internal/entity/constants"
)

func WebSISUserHandler(ctx *fiber.Ctx) error {

	return ctx.Render("sis/pages/user", constantsEntity.WebData{
		Title:        constantsEntity.WebSISUserTitle,
		BaseURL:      constantsEntity.BaseURL,
		TemplateURL:  constantsEntity.TemplateUrl,
		CurrentURL:   constantsEntity.WebSISUserURL,
		LinkPageList: constantsEntity.LinkPageList,
	})
}
