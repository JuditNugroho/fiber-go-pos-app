package personalia

import (
	"github.com/gofiber/fiber/v2"

	constantsEntity "github.com/fiber-go-pos-app/internal/entity/constants"
)

func WebSISMemberHandler(ctx *fiber.Ctx) error {

	return ctx.Render("sis/pages/member", constantsEntity.WebData{
		Title:        constantsEntity.WebSISMemberTitle,
		BaseURL:      constantsEntity.BaseURL,
		StaticUrl:    constantsEntity.StaticUrl,
		CurrentURL:   constantsEntity.WebSISMemberURL,
		LinkPageList: constantsEntity.LinkPageList,
	})
}
