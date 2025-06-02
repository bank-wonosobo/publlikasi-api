package routes

import (
	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/handlers"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/services"
	"github.com/bank-wonosobo/publlikasi-api.git/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

func RegisterAnnouncementRouter(router fiber.Router, service services.AnnouncementService, validator validator.Validator) {
	handler := handlers.NewAnnouncement(service, validator)

	// public routes
	announcements := router.Group("/announcements")
	announcements.Get("/", handler.Index)

	// admin routes
	admin := router.Group("/admin/announcements")
	admin.Get("/", handler.Index)
}
