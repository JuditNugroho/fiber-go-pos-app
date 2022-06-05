package services

import (
	"github.com/gofiber/fiber/v2"

	productSvc "github.com/fiber-go-pos-app/internal/handler/services/product"
)

// BuildProductRoutes : Service - service to handle product
func BuildProductRoutes(service fiber.Router) {
	service.Get("/product", productSvc.GetProductByIDOrBarcode)
	service.Get("/dt_products", productSvc.GetDTProductHandler)
	service.Post("/product/insert", productSvc.InsertProductHandler)
	service.Post("/product/update", productSvc.UpdateProductHandler)
	service.Post("/product/delete", productSvc.DeleteProductHandler)
	service.Post("/product/upsert", productSvc.UpsertProductHandler)
}
