package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtWare "github.com/gofiber/jwt/v3"

	"github.com/fiber-go-pos-app/utils/pkg/jwt"

	constantsEntity "github.com/fiber-go-pos-app/internal/entity/constants"
)

// Protected protect routes
func Protected() fiber.Handler {
	return jwtWare.New(jwtWare.Config{
		SigningMethod: constantsEntity.JWTMethod,
		SigningKey:    jwt.GetPrivateKey().Public(),
		ErrorHandler:  jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == constantsEntity.ErrMissingOrMalformedJWT {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": constantsEntity.ErrMissingOrMalformedJWT, "data": nil})
	}

	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": constantsEntity.ErrInvalidORExpiredJWT, "data": nil})
}
