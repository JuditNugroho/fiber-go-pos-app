package form

import (
	"github.com/fiber-go-pos-app/utils/pkg/custom"
	"github.com/fiber-go-pos-app/utils/pkg/jwt"
	"github.com/gofiber/fiber/v2"

	constantsEntity "github.com/fiber-go-pos-app/internal/entity/constants"
	formEntity "github.com/fiber-go-pos-app/internal/entity/form"

	personaliaRepo "github.com/fiber-go-pos-app/internal/repo/personalia"
)

func LoginForm(ctx *fiber.Ctx, req formEntity.LoginRequest) (formEntity.LoginResponse, error) {
	// Initialization variable
	var res formEntity.LoginResponse

	data, err := personaliaRepo.GetUserByUserName(ctx, req.UserName)
	if err != nil {
		return res, err
	}

	// check hash password
	if !custom.CheckPasswordHash(req.Password, data.Password) {
		return res, constantsEntity.ErrWrongPassword
	}

	token, err := jwt.CreateJWTToken(formEntity.JWTRequest{
		UserID: data.UserID,
		Name:   data.UserName,
		Admin:  data.IsAdmin,
	})
	if err != nil {
		return res, err
	}

	return formEntity.LoginResponse{
		UserID:   data.UserID,
		UserName: data.UserName,
		IsAdmin:  data.IsAdmin,
		JWTToken: token,
	}, nil
}
