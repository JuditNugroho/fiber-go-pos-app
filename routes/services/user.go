package services

import (
	"github.com/gofiber/fiber/v2"

	personaliaSvc "github.com/fiber-go-pos-app/internal/handler/services/personalia"
)

func BuildUserRoutes(service fiber.Router) {
	service.Get("/users", personaliaSvc.GetAllUserHandler)
	service.Get("/dt_users", personaliaSvc.GetAllDTUserHandler)
	service.Post("/user/insert", personaliaSvc.InsertUserHandler)
	service.Post("/user/update", personaliaSvc.UpdateUserHandler)
	service.Post("/user/delete", personaliaSvc.DeleteUserHandler)
	service.Post("/user/upsert", personaliaSvc.UpsertUserHandler)
}
