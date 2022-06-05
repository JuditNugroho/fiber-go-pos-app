package web

import (
	"github.com/gofiber/fiber/v2"

	loginSvc "github.com/fiber-go-pos-app/internal/handler/web/login"
)

func BuildLoginRoutes(service fiber.Router) {
	service.Get("/", loginSvc.WebLoginHandler)
}
