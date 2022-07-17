package services

import (
	"github.com/gofiber/fiber/v2"

	formSvc "github.com/fiber-go-pos-app/internal/handler/services/form"
)

// BuildLoginRoutes : Service - service to handle login
func BuildLoginRoutes(service fiber.Router) {
	service.Post("/login", formSvc.LoginHandler)
}
