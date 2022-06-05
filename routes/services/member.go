package services

import (
	"github.com/gofiber/fiber/v2"

	personaliaSvc "github.com/fiber-go-pos-app/internal/handler/services/personalia"
)

// BuildMemberRoutes : Service - service to handle member
func BuildMemberRoutes(service fiber.Router) {
	service.Get("/members", personaliaSvc.GetAllMemberHandler)
	service.Get("/dt_members", personaliaSvc.GetAllDTMemberHandler)
	service.Post("/member/insert", personaliaSvc.InsertMemberHandler)
	service.Post("/member/update", personaliaSvc.UpdateMemberHandler)
	service.Post("/member/delete", personaliaSvc.DeleteMemberHandler)
	service.Post("/member/upsert", personaliaSvc.UpsertMemberHandler)
}
