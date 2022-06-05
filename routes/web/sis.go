package web

import (
	"github.com/fiber-go-pos-app/internal/handler/web/sis/personalia"
	productWeb "github.com/fiber-go-pos-app/internal/handler/web/sis/product"
	"github.com/gofiber/fiber/v2"

	homeWeb "github.com/fiber-go-pos-app/internal/handler/web/sis/home"
)

func BuildSISRoutes(service fiber.Router) {
	service.Get("/", homeWeb.WebSISHomeHandler)
	service.Get("/user", personalia.WebSISUserHandler)
	service.Get("/member", personalia.WebSISMemberHandler)
	service.Get("/product", productWeb.WebProductHandler)
}
