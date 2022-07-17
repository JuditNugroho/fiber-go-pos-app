package form

import (
	"github.com/fiber-go-pos-app/utils/pkg/custom"
	"github.com/gofiber/fiber/v2"

	constantsEntity "github.com/fiber-go-pos-app/internal/entity/constants"
	formEntity "github.com/fiber-go-pos-app/internal/entity/form"

	formRepo "github.com/fiber-go-pos-app/internal/repo/form"
)

func LoginForm(ctx *fiber.Ctx, req formEntity.LoginRequest) error {
	data, err := formRepo.LoginFormRepo(ctx, req.UserName)
	if err != nil {
		return err
	}

	// check hash password
	if !custom.CheckPasswordHash(req.Password, data.Password) {
		return constantsEntity.ErrWrongPassword
	}

	return nil
}
