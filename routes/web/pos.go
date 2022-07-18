package web

import (
	"github.com/fiber-go-pos-app/utils/pkg/middleware"
	"github.com/gofiber/fiber/v2"

	homeWeb "github.com/fiber-go-pos-app/internal/handler/web/pos/home"
)

func BuildPOSRoutes(service fiber.Router) {

	service.Get("/", middleware.Protected(), homeWeb.WebPOSHomeHandler)
}
