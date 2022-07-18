package form

import (
	"fmt"
	"github.com/fiber-go-pos-app/utils/pkg/custom"
	"github.com/fiber-go-pos-app/utils/pkg/jwt"
	"github.com/gofiber/fiber/v2"

	constantsEntity "github.com/fiber-go-pos-app/internal/entity/constants"
	formEntity "github.com/fiber-go-pos-app/internal/entity/form"

	personaliaRepo "github.com/fiber-go-pos-app/internal/repo/personalia"
)

func LoginForm(ctx *fiber.Ctx, req formEntity.LoginRequest) error {
	data, err := personaliaRepo.GetUserByUserName(ctx, req.UserName)
	if err != nil {
		return err
	}

	// check hash password
	if !custom.CheckPasswordHash(req.Password, data.Password) {
		return constantsEntity.ErrWrongPassword
	}

	fmt.Println("ASDDASDSASASDA")

	token, err := jwt.CreateJWTToken(formEntity.JWTRequest{
		UserID: data.UserID,
		Name:   data.UserName,
		Admin:  data.IsAdmin,
	})

	fmt.Println("TOKEN : ", token)

	if err != nil {
		return err
	}

	return nil
}
