package home

import (
	"github.com/gofiber/fiber/v2"

	constantsEntity "github.com/fiber-go-pos-app/internal/entity/constants"
)

func WebPOSHomeHandler(ctx *fiber.Ctx) error {

	return ctx.Render("pos/index", constantsEntity.WebData{
		Title:        constantsEntity.WebPOSHomeTitle,
		BaseURL:      constantsEntity.BaseURL,
		CurrentURL:   constantsEntity.WebPOSHomeURL,
		StaticUrl:    constantsEntity.StaticUrl,
		LinkPageList: constantsEntity.LinkPageList,
	})
}
