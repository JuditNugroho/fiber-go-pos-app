package login

import (
	"github.com/gofiber/fiber/v2"

	constantsEntity "github.com/fiber-go-pos-app/internal/entity/constants"
)

func WebLoginHandler(ctx *fiber.Ctx) error {

	return ctx.Render("login/index", constantsEntity.WebData{
		Title:        constantsEntity.WebLoginTitle,
		BaseURL:      constantsEntity.BaseURL,
		CurrentURL:   constantsEntity.WebLoginURL,
		StaticUrl:    constantsEntity.StaticUrl,
		LinkPageList: constantsEntity.LinkPageList,
	})
}
